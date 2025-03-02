/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

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

package scenarios

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	clusterv1alpha1 "github.com/kubermatic/machine-controller/pkg/apis/cluster/v1alpha1"
	providerconfig "github.com/kubermatic/machine-controller/pkg/providerconfig/types"
	"github.com/kubermatic/machine-controller/pkg/userdata/flatcar"
	"k8c.io/kubermatic/v2/cmd/conformance-tester/pkg/types"
	apiv1 "k8c.io/kubermatic/v2/pkg/api/v1"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources/machine"
	apimodels "k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"

	"k8s.io/utils/pointer"
)

const (
	nodeCpu      = 2
	nodeDiskSize = 60
	nodeMemory   = 2048
)

type anexiaScenario struct {
	baseScenario
}

func (s *anexiaScenario) IsValid(opts *types.Options, log *zap.SugaredLogger) bool {
	if !s.baseScenario.IsValid(opts, log) {
		return false
	}

	if s.operatingSystem != providerconfig.OperatingSystemFlatcar {
		s.Log(log).Debug("Skipping because provider only supports Flatcar.")
		return false
	}

	return true
}

func (s *anexiaScenario) APICluster(secrets types.Secrets) *apimodels.CreateClusterSpec {
	return &apimodels.CreateClusterSpec{
		Cluster: &apimodels.Cluster{
			Spec: &apimodels.ClusterSpec{
				ContainerRuntime: s.containerRuntime,
				Cloud: &apimodels.CloudSpec{
					DatacenterName: secrets.Anexia.KKPDatacenter,
					Anexia: &apimodels.AnexiaCloudSpec{
						Token: secrets.Anexia.Token,
					},
				},
				Version: apimodels.Semver(s.version.String()),
			},
		},
	}
}

func (s *anexiaScenario) Cluster(secrets types.Secrets) *kubermaticv1.ClusterSpec {
	return &kubermaticv1.ClusterSpec{
		ContainerRuntime: s.containerRuntime,
		Cloud: kubermaticv1.CloudSpec{
			DatacenterName: secrets.Anexia.KKPDatacenter,
			Anexia: &kubermaticv1.AnexiaCloudSpec{
				Token: secrets.Anexia.Token,
			},
		},
		Version: s.version,
	}
}

func (s *anexiaScenario) NodeDeployments(_ context.Context, num int, secrets types.Secrets) ([]apimodels.NodeDeployment, error) {
	replicas := int32(num)

	osSpec, err := s.APIOperatingSystemSpec()
	if err != nil {
		return nil, fmt.Errorf("failed to build OS spec: %w", err)
	}

	return []apimodels.NodeDeployment{
		{
			Spec: &apimodels.NodeDeploymentSpec{
				Replicas: &replicas,
				Template: &apimodels.NodeSpec{
					Cloud: &apimodels.NodeCloudSpec{
						Anexia: &apimodels.AnexiaNodeSpec{
							DiskSize:   nodeDiskSize,
							CPUs:       pointer.Int64(nodeCpu),
							Memory:     pointer.Int64(nodeMemory),
							TemplateID: &secrets.Anexia.TemplateID,
							VlanID:     &secrets.Anexia.VlanID,
						},
					},
					Versions: &apimodels.NodeVersionInfo{
						Kubelet: s.version.String(),
					},
					OperatingSystem: osSpec,
				},
			},
		},
		{
			Spec: &apimodels.NodeDeploymentSpec{
				Replicas: &replicas,
				Template: &apimodels.NodeSpec{
					Cloud: &apimodels.NodeCloudSpec{
						Anexia: &apimodels.AnexiaNodeSpec{
							Disks: []*apimodels.AnexiaDiskConfig{
								{
									Size: pointer.Int64(nodeDiskSize),
								},
							},
							CPUs:       pointer.Int64(nodeCpu),
							Memory:     pointer.Int64(nodeMemory),
							TemplateID: &secrets.Anexia.TemplateID,
							VlanID:     &secrets.Anexia.VlanID,
						},
					},
					Versions: &apimodels.NodeVersionInfo{
						Kubelet: s.version.String(),
					},
					OperatingSystem: osSpec,
				},
			},
		},
	}, nil
}

func (s *anexiaScenario) MachineDeployments(_ context.Context, num int, secrets types.Secrets, cluster *kubermaticv1.Cluster) ([]clusterv1alpha1.MachineDeployment, error) {
	osSpec, err := s.OperatingSystemSpec()
	if err != nil {
		return nil, fmt.Errorf("failed to build OS spec: %w", err)
	}

	osSpec.Flatcar.ProvisioningUtility = flatcar.CloudInit

	nodeSpec := apiv1.NodeSpec{
		OperatingSystem: *osSpec,
		Cloud: apiv1.NodeCloudSpec{
			Anexia: &apiv1.AnexiaNodeSpec{
				CPUs:       nodeCpu,
				Memory:     nodeMemory,
				DiskSize:   pointer.Int64(nodeDiskSize),
				TemplateID: secrets.Anexia.TemplateID,
				VlanID:     secrets.Anexia.VlanID,
			},
		},
	}

	config, err := machine.GetAnexiaProviderConfig(cluster, nodeSpec, s.datacenter)
	if err != nil {
		return nil, err
	}

	md, err := createMachineDeployment(num, &s.version, s.operatingSystem, osSpec, s.cloudProvider, config, s.dualstackEnabled)
	if err != nil {
		return nil, err
	}

	return []clusterv1alpha1.MachineDeployment{md}, nil
}
