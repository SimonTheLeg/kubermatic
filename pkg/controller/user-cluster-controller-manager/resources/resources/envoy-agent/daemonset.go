/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package envoyagent

import (
	"fmt"
	"net"

	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"
	"k8c.io/kubermatic/v2/pkg/resources/registry"
	"k8c.io/kubermatic/v2/pkg/version/kubermatic"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilpointer "k8s.io/utils/pointer"
)

var (
	defaultResourceRequirements = map[string]*corev1.ResourceRequirements{
		resources.EnvoyAgentDaemonSetName: {
			Requests: corev1.ResourceList{
				corev1.ResourceMemory: resource.MustParse("32Mi"),
				corev1.ResourceCPU:    resource.MustParse("50m"),
			},
			Limits: corev1.ResourceList{
				corev1.ResourceMemory: resource.MustParse("64Mi"),
				corev1.ResourceCPU:    resource.MustParse("1"),
			},
		},
	}
)

const (
	envoyImageName = "envoyproxy/envoy"
)

// DaemonSetCreator returns the function to create and update the Envoy DaemonSet.
func DaemonSetCreator(agentIP net.IP, versions kubermatic.Versions, configHash string, imageRewriter registry.ImageRewriter) reconciling.NamedDaemonSetCreatorGetter {
	return func() (string, reconciling.DaemonSetCreator) {
		return resources.EnvoyAgentDaemonSetName, func(ds *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {
			ds.Name = resources.EnvoyAgentDaemonSetName
			ds.Namespace = metav1.NamespaceSystem
			ds.Labels = resources.BaseAppLabels(resources.EnvoyAgentDaemonSetName, nil)

			ds.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: resources.BaseAppLabels(resources.EnvoyAgentDaemonSetName,
					map[string]string{"app.kubernetes.io/name": "envoy-agent"}),
			}

			// has to be the same as the selector
			ds.Spec.Template.ObjectMeta = metav1.ObjectMeta{
				Labels: resources.BaseAppLabels(resources.EnvoyAgentDaemonSetName,
					map[string]string{"app.kubernetes.io/name": "envoy-agent"}),

				// Used to force the restart of the envoy-agent to re-read its configuration
				// from the configMap when it changes. Necessary to support switching to/from Konnectivity.
				Annotations: map[string]string{"checksum/config": configHash},
			}

			initContainers, err := getInitContainers(agentIP, versions, imageRewriter)
			if err != nil {
				return nil, err
			}

			containers, err := getContainers(versions, imageRewriter)
			if err != nil {
				return nil, err
			}

			ds.Spec.Template.Spec = corev1.PodSpec{
				InitContainers: initContainers,
				Containers:     containers,
				// TODO(youssefazrak) needed?
				PriorityClassName:             "system-cluster-critical",
				DNSPolicy:                     corev1.DNSClusterFirst,
				HostNetwork:                   true,
				Volumes:                       getVolumes(),
				RestartPolicy:                 corev1.RestartPolicyAlways,
				TerminationGracePeriodSeconds: utilpointer.Int64Ptr(30),
				SecurityContext: &corev1.PodSecurityContext{
					SeccompProfile: &corev1.SeccompProfile{
						Type: corev1.SeccompProfileTypeRuntimeDefault,
					},
				},
				SchedulerName: corev1.DefaultSchedulerName,
			}
			ds.Spec.Template.Spec.Tolerations = []corev1.Toleration{
				{
					Effect:   corev1.TaintEffectNoSchedule,
					Operator: corev1.TolerationOpExists,
				},
				{
					Effect:   corev1.TaintEffectNoExecute,
					Operator: corev1.TolerationOpExists,
				},
			}
			if err := resources.SetResourceRequirements(ds.Spec.Template.Spec.Containers, defaultResourceRequirements, nil, ds.Annotations); err != nil {
				return nil, fmt.Errorf("failed to set resource requirements: %w", err)
			}

			return ds, nil
		}
	}
}

func getInitContainers(ip net.IP, versions kubermatic.Versions, imageRewriter registry.ImageRewriter) ([]corev1.Container, error) {
	image := registry.Must(imageRewriter(fmt.Sprintf("%s/%s:%s", resources.RegistryQuay, resources.EnvoyAgentDeviceSetupImage, versions.Kubermatic)))

	// TODO: we are creating and configuring the a dummy interface
	// using init containers. This approach is good enough for the tech preview
	// but it is definitely not production ready. This should be replaced with
	// a binary that properly handles error conditions and reconciles the
	// interface in a loop.
	return []corev1.Container{
		{
			Name:    resources.EnvoyAgentCreateInterfaceInitContainerName,
			Image:   image,
			Command: []string{"sh", "-c", "ip link add envoyagent type dummy || true"},
			SecurityContext: &corev1.SecurityContext{
				Capabilities: &corev1.Capabilities{
					Add: []corev1.Capability{
						"NET_ADMIN",
					},
					Drop: []corev1.Capability{
						"all",
					},
				},
			},
		},
		{
			Name:    resources.EnvoyAgentAssignAddressInitContainerName,
			Image:   image,
			Command: []string{"sh", "-c", fmt.Sprintf("ip addr add %s/32 dev envoyagent scope host || true", ip.String())},
			SecurityContext: &corev1.SecurityContext{
				Capabilities: &corev1.Capabilities{
					Add: []corev1.Capability{
						"NET_ADMIN",
					},
					Drop: []corev1.Capability{
						"all",
					},
				},
			},
		},
	}, nil
}

func getContainers(versions kubermatic.Versions, imageRewriter registry.ImageRewriter) ([]corev1.Container, error) {
	return []corev1.Container{
		{
			Name:            resources.EnvoyAgentDaemonSetName,
			Image:           registry.Must(imageRewriter(fmt.Sprintf("%s:%s", envoyImageName, versions.Envoy))),
			ImagePullPolicy: corev1.PullIfNotPresent,

			// This amount of logs will be kept for the Tech Preview of
			// the new expose strategy
			Args: []string{"--config-path", "etc/envoy/envoy.yaml", "--component-log-level", "upstream:trace,connection:trace,http:trace,router:trace,filter:trace"},
			VolumeMounts: []corev1.VolumeMount{
				{
					Name:      "config-volume",
					MountPath: "/etc/envoy/envoy.yaml",
					SubPath:   resources.EnvoyAgentConfigFileName,
				},
			},
			SecurityContext: &corev1.SecurityContext{
				Capabilities: &corev1.Capabilities{
					Add: []corev1.Capability{
						"CHOWN",
						"SETGID",
						"SETUID",
						"NET_BIND_SERVICE",
					},
					Drop: []corev1.Capability{
						"all",
					},
				},
			},
		},
	}, nil
}

func getVolumes() []corev1.Volume {
	return []corev1.Volume{
		{
			Name: "config-volume",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					DefaultMode: utilpointer.Int32Ptr(corev1.ConfigMapVolumeSourceDefaultMode),
					LocalObjectReference: corev1.LocalObjectReference{
						Name: resources.EnvoyAgentConfigMapName,
					},
				},
			},
		},
	}
}
