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

package nodeportproxy

import (
	"fmt"
	"strconv"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/controller/operator/common"
	"k8c.io/kubermatic/v2/pkg/features"
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"
	"k8c.io/kubermatic/v2/pkg/version/kubermatic"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/pointer"
)

const (
	ServiceAccountName    = "nodeport-proxy"
	EnvoyDeploymentName   = "nodeport-proxy-envoy"
	UpdaterDeploymentName = "nodeport-proxy-updater"
	EnvoyPort             = 8002
	EnvoySNIPort          = 6443
	EnvoyTunnelingPort    = 8088
)

func EnvoyDeploymentCreator(cfg *kubermaticv1.KubermaticConfiguration, seed *kubermaticv1.Seed, supportsFailureDomainZoneAntiAffinity bool, versions kubermatic.Versions) reconciling.NamedDeploymentCreatorGetter {
	return func() (string, reconciling.DeploymentCreator) {
		return EnvoyDeploymentName, func(d *appsv1.Deployment) (*appsv1.Deployment, error) {
			d.Spec.Replicas = pointer.Int32Ptr(3)
			d.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: map[string]string{
					common.NameLabel: EnvoyDeploymentName,
				},
			}

			maxSurge := intstr.FromString("25%")
			d.Spec.Strategy = appsv1.DeploymentStrategy{
				Type: appsv1.RollingUpdateDeploymentStrategyType,
				RollingUpdate: &appsv1.RollingUpdateDeployment{
					MaxSurge:       &maxSurge,
					MaxUnavailable: &maxSurge,
				},
			}

			d.Spec.Template.Labels = d.Spec.Selector.MatchLabels
			d.Spec.Template.Annotations = map[string]string{
				"prometheus.io/scrape":       "true",
				"prometheus.io/port":         strconv.Itoa(EnvoyPort),
				"prometheus.io/metrics_path": "/stats/prometheus",
				"fluentbit.io/parser":        "json_iso",
			}

			d.Spec.Template.Spec.RestartPolicy = corev1.RestartPolicyAlways
			d.Spec.Template.Spec.ServiceAccountName = ServiceAccountName

			d.Spec.Template.Spec.InitContainers = []corev1.Container{
				{
					Name:    "copy-envoy-config",
					Image:   seed.Spec.NodeportProxy.EnvoyManager.DockerRepository + ":" + versions.Kubermatic,
					Command: []string{"/bin/cp"},
					Args:    []string{"/envoy.yaml", "/etc/envoy/envoy.yaml"},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "envoy-config",
							MountPath: "/etc/envoy",
						},
					},
				},
			}

			args := []string{
				"-listen-address=:8001",
				"-envoy-node-name=kube",
				"-envoy-admin-port=9001",
				fmt.Sprintf("-envoy-stats-port=%d", EnvoyPort),
			}
			if cfg.Spec.FeatureGates[features.TunnelingExposeStrategy] {
				args = append(args,
					fmt.Sprintf("-envoy-sni-port=%d", EnvoySNIPort),
					fmt.Sprintf("-envoy-tunneling-port=%d", EnvoyTunnelingPort))
			}
			d.Spec.Template.Spec.Containers = []corev1.Container{
				{
					Name:    "envoy-manager",
					Image:   seed.Spec.NodeportProxy.EnvoyManager.DockerRepository + ":" + versions.Kubermatic,
					Command: []string{"/envoy-manager"},
					Args:    args,
					Ports: []corev1.ContainerPort{
						{
							Name:          "grpc",
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: 8001,
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "envoy-config",
							MountPath: "/etc/envoy",
						},
					},
					Resources: seed.Spec.NodeportProxy.EnvoyManager.Resources,
				},

				{
					Name:    "envoy",
					Image:   seed.Spec.NodeportProxy.Envoy.DockerRepository + ":" + versions.Envoy,
					Command: []string{"/usr/local/bin/envoy"},
					Args: []string{
						"-c",
						"/etc/envoy/envoy.yaml",
						"--service-cluster",
						"cluster0",
						"--service-node",
						"kube",
					},
					Ports: []corev1.ContainerPort{
						{
							Name:          "stats",
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: EnvoyPort,
						},
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "envoy-config",
							MountPath: "/etc/envoy",
						},
					},
					LivenessProbe: &corev1.Probe{
						FailureThreshold: 3,
						SuccessThreshold: 1,
						TimeoutSeconds:   1,
						PeriodSeconds:    3,
						ProbeHandler: corev1.ProbeHandler{
							HTTPGet: &corev1.HTTPGetAction{
								Port:   intstr.FromInt(EnvoyPort),
								Scheme: corev1.URISchemeHTTP,
								Path:   "/healthz",
							},
						},
					},
					Lifecycle: &corev1.Lifecycle{
						PreStop: &corev1.LifecycleHandler{
							Exec: &corev1.ExecAction{
								Command: []string{
									"wget",
									"-qO-",
									"http://127.0.0.1:9001/healthcheck/fail",
								},
							},
						},
					},
					Resources: seed.Spec.NodeportProxy.Envoy.Resources,
				},
			}
			d.Spec.Template.Spec.Affinity = HostnameAntiAffinity(EnvoyDeploymentName)
			if supportsFailureDomainZoneAntiAffinity {
				antiAffinities := d.Spec.Template.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution
				antiAffinities = append(antiAffinities, resources.FailureDomainZoneAntiAffinity(EnvoyDeploymentName))
				d.Spec.Template.Spec.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution = antiAffinities
			}

			d.Spec.Template.Spec.Volumes = []corev1.Volume{
				{
					Name: "envoy-config",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				},
			}

			return d, nil
		}
	}
}

func HostnameAntiAffinity(app string) *corev1.Affinity {
	return &corev1.Affinity{
		PodAntiAffinity: &corev1.PodAntiAffinity{
			PreferredDuringSchedulingIgnoredDuringExecution: hostnameAntiAffinity(app),
		},
	}
}

func hostnameAntiAffinity(app string) []corev1.WeightedPodAffinityTerm {
	return []corev1.WeightedPodAffinityTerm{
		// Avoid that we schedule multiple same-kind pods of a cluster on a single node
		{
			Weight: 10,
			PodAffinityTerm: corev1.PodAffinityTerm{
				LabelSelector: &metav1.LabelSelector{
					MatchLabels: map[string]string{
						resources.AppLabelKey: app,
					},
				},
				TopologyKey: resources.TopologyKeyHostname,
			},
		},
	}
}

func EnvoyPDBCreator() reconciling.NamedPodDisruptionBudgetCreatorGetter {
	maxUnavailable := intstr.FromInt(1)
	return func() (string, reconciling.PodDisruptionBudgetCreator) {
		return EnvoyDeploymentName, func(pdb *policyv1.PodDisruptionBudget) (*policyv1.PodDisruptionBudget, error) {
			pdb.Spec.MaxUnavailable = &maxUnavailable
			pdb.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: map[string]string{
					common.NameLabel: EnvoyDeploymentName,
				},
			}
			return pdb, nil
		}
	}
}

func UpdaterDeploymentCreator(cfg *kubermaticv1.KubermaticConfiguration, seed *kubermaticv1.Seed, versions kubermatic.Versions) reconciling.NamedDeploymentCreatorGetter {
	return func() (string, reconciling.DeploymentCreator) {
		return UpdaterDeploymentName, func(d *appsv1.Deployment) (*appsv1.Deployment, error) {
			d.Spec.Replicas = pointer.Int32Ptr(1)
			d.Spec.Selector = &metav1.LabelSelector{
				MatchLabels: map[string]string{
					common.NameLabel: UpdaterDeploymentName,
				},
			}

			d.Spec.Template.Labels = d.Spec.Selector.MatchLabels
			d.Spec.Template.Annotations = map[string]string{
				"fluentbit.io/parser": "json_iso",
			}

			d.Spec.Template.Spec.ServiceAccountName = ServiceAccountName

			args := []string{
				"-lb-namespace=$(NAMESPACE)",
				fmt.Sprintf("-lb-name=%s", ServiceName),
			}
			if cfg.Spec.FeatureGates[features.TunnelingExposeStrategy] {
				args = append(args,
					fmt.Sprintf("-envoy-sni-port=%d", EnvoySNIPort),
					fmt.Sprintf("-envoy-tunneling-port=%d", EnvoyTunnelingPort))
			}
			d.Spec.Template.Spec.Containers = []corev1.Container{
				{
					Name:    "lb-updater",
					Image:   seed.Spec.NodeportProxy.Updater.DockerRepository + ":" + versions.Kubermatic,
					Command: []string{"/lb-updater"},
					Args:    args,
					Env: []corev1.EnvVar{
						{
							Name: "NAMESPACE",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.namespace",
								},
							},
						},
					},
					Resources: seed.Spec.NodeportProxy.Updater.Resources,
				},
			}

			return d, nil
		}
	}
}
