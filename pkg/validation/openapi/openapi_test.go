package openapi

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8s.io/apiextensions-apiserver/pkg/apiserver/validation"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestValidatorFromCRD(t *testing.T) {
	const v1crd = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: samples.kubermatic.k8c.io
spec:
  group: kubermatic.k8c.io
  names:
    kind: Sample
    listKind: SampleList
    plural: Samples
    singular: Sample
  scope: Cluster
  versions:
    - name: v1
      schema:
        openAPIV3Schema:
          description: ""
          properties:
            apiVersion:
              description: ""
              type: string
            kind:
              description: ""
              type: string
            metadata:
              type: object
            spec:
              description: spec
              properties:
                val1:
                  description: val1
                  type: string
                  enum:
                    - valid1
          type: object
      served: true
      storage: true
    - name: v2
      schema:
        openAPIV3Schema:
          description: ""
          properties:
            apiVersion:
              description: ""
              type: string
            kind:
              description: ""
              type: string
            metadata:
              type: object
            spec:
              description: spec
              properties:
                val1:
                  description: val1
                  type: string
                  enum:
                    - valid1
                    - valid2
          type: object
      served: true
      storage: true
`

	const v1beta1crd = `
---
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: samples.kubermatic.k8c.io
spec:
  group: kubermatic.k8c.io
  names:
    kind: Sample
    listKind: SampleList
    plural: Samples
    singular: Sample
  scope: Cluster
  validation:
    openAPIV3Schema:
      description: ""
      properties:
        apiVersion:
          description: ""
          type: string
        kind:
          description: ""
          type: string
        metadata:
          type: object
        spec:
          description: spec
          properties:
            val1:
              description: val1
              type: string
              enum:
                - valid1
    type: object
  versions:
    - name: v1
      served: true
      storage: true
`

	type sampleSpec struct {
		Val1 string `json:"val1"`
	}
	type sample struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec sampleSpec `json:"spec,omitempty"`
	}

	tt := map[string]struct {
		crd          string
		saKind       string
		saAPIVersion string
		saVal        string
		expValErrs   int
		expErr       bool
	}{
		"valid v1 sample": {
			v1crd,
			"sample",
			"kubermatic.k8c.io/v1",
			"valid1",
			0,
			false,
		},
		"invalid v1 sample": {
			v1crd,
			"sample",
			"kubermatic.k8c.io/v1",
			"valid2",
			1,
			false,
		},
		"valid v2 sample": {
			v1crd,
			"sample",
			"kubermatic.k8c.io/v2",
			"valid2",
			0,
			false,
		},
		"unsupported APIVersion": {
			v1crd,
			"sample",
			"kubermatic.k8c.io/vInvalid",
			"valid1",
			0,
			true,
		},
		"valid global validation sample": {
			v1beta1crd,
			"sample",
			"kubermatic.k8c.io/v1",
			"valid1",
			0,
			false,
		},
		"invalid global validation sample": {
			v1beta1crd,
			"sample",
			"kubermatic.k8c.io/v1",
			"valid2",
			1,
			false,
		},
		"empty desired version": {
			v1crd,
			"sample",
			"",
			"",
			0,
			true,
		},
		"unsupported crd version": {
			"apiVersion: apiextensions.k8s.io/vinvalid\nkind: CustomResourceDefinition",
			"sample",
			"kubermatic.k8c.io/v1",
			"valid1",
			0,
			true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := &sample{
				TypeMeta: metav1.TypeMeta{Kind: tc.saKind, APIVersion: tc.saAPIVersion},
				Spec:     sampleSpec{Val1: tc.saVal},
			}

			v, err := ValidatorFromCRD(strings.NewReader(tc.crd), s.GetObjectKind().GroupVersionKind().Version)
			res := validation.ValidateCustomResource(nil, s, v)

			if tc.expValErrs != len(res) {
				t.Errorf("Exp Errorlist length to be %d, got %d", tc.expValErrs, len(res))
			}

			if err != nil && !tc.expErr {
				t.Errorf("Exp err to be nil, but got %q", err)
			}

			// TODO Delete later
			// for _, e := range res {
			// 	fmt.Println(e.Error())
			// }

		})
	}
}

func TestValidatorForType(t *testing.T) {
	tt := map[string]struct {
		in           *metav1.TypeMeta
		expValidator bool
		expErr       bool
	}{
		"k8c.io crd": {
			&metav1.TypeMeta{Kind: "Cluster", APIVersion: "kubermatic.k8c.io/v1"},
			true,
			false,
		},
		"apps.k8c.io crd": {
			&metav1.TypeMeta{Kind: "ApplicationDefinition", APIVersion: "apps.kubermatic.k8c.io/v1"},
			true,
			false,
		},
		"invalid kind": {
			&metav1.TypeMeta{Kind: "Invalid", APIVersion: "kubermatic.k8c.io/v1"},
			false,
			true,
		},
		"invalid apiversion": {
			&metav1.TypeMeta{Kind: "Cluster", APIVersion: "Invalid"},
			false,
			true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res, err := ValidatorForType(tc.in)

			// TODO delete later
			cl := &kubermaticv1.Cluster{
				Spec: kubermaticv1.ClusterSpec{CNIPlugin: &kubermaticv1.CNIPluginSettings{Type: "invalid"}},
			}
			valid := validation.ValidateCustomResource(nil, cl, res)
			fmt.Println(valid)

			if res != nil {
				if tc.expValidator && len(res.Schema.SchemaProps.Properties) == 0 {
					t.Errorf("Root SchemaProps.Properties are empty, when they should not be")
				}
			}

			if !tc.expErr && err != nil {
				t.Errorf("Exp error to be nil, but got %q", err)
			}

		})
	}
}

// TODO delete later
const clusterCrd = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  annotations:
    controller-gen.kubebuilder.io/version: v0.7.0
  creationTimestamp: null
  name: clusters.kubermatic.k8c.io
spec:
  group: kubermatic.k8c.io
  names:
    kind: Cluster
    listKind: ClusterList
    plural: clusters
    singular: cluster
  scope: Cluster
  versions:
  - additionalPrinterColumns:
    - jsonPath: .spec.humanReadableName
      name: HumanReadableName
      type: string
    - jsonPath: .status.userEmail
      name: Owner
      type: string
    - jsonPath: .spec.version
      name: Version
      type: string
    - jsonPath: .spec.cloud.providerName
      name: Provider
      type: string
    - jsonPath: .spec.cloud.dc
      name: Datacenter
      type: string
    - jsonPath: .spec.pause
      name: Paused
      type: boolean
    - jsonPath: .metadata.creationTimestamp
      name: Age
      type: date
    name: v1
    schema:
      openAPIV3Schema:
        description: Cluster is the object representing a cluster.
        properties:
          address:
            description: ClusterAddress stores access and address information of a
              cluster.
            properties:
              adminToken:
                description: AdminToken is the token for the kubeconfig, the user
                  can download
                type: string
              externalName:
                description: ExternalName is the DNS name for this cluster
                type: string
              internalURL:
                description: InternalName is the seed cluster internal absolute DNS
                  name to the API server
                type: string
              ip:
                description: IP is the external IP under which the apiserver is available
                type: string
              port:
                description: Port is the port the API server listens on
                format: int32
                type: integer
              url:
                description: URL under which the Apiserver is available
                type: string
            required:
            - adminToken
            - externalName
            - internalURL
            - ip
            - port
            - url
            type: object
          apiVersion:
            description: 'APIVersion defines the versioned schema of this representation
              of an object. Servers should convert recognized schemas to the latest
              internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
            type: string
          kind:
            description: 'Kind is a string value representing the REST resource this
              object represents. Servers may infer this from the endpoint the client
              submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          metadata:
            type: object
          spec:
            description: ClusterSpec specifies the data for a new cluster.
            properties:
              admissionPlugins:
                description: AdmissionPlugins provides the ability to pass arbitrary
                  names of admission plugins to kube-apiserver
                items:
                  type: string
                type: array
              auditLogging:
                properties:
                  enabled:
                    type: boolean
                  policyPreset:
                    enum:
                    - ""
                    - metadata
                    - recommended
                    - minimal
                    type: string
                type: object
              cloud:
                description: CloudSpec mutually stores access data to a cloud provider.
                properties:
                  alibaba:
                    description: AlibabaCloudSpec specifies the access data to Alibaba.
                    properties:
                      accessKeyID:
                        type: string
                      accessKeySecret:
                        type: string
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                    type: object
                  anexia:
                    description: AnexiaCloudSpec specifies the access data to Anexia.
                    properties:
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      token:
                        type: string
                    type: object
                  aws:
                    description: AWSCloudSpec specifies access data to Amazon Web
                      Services.
                    properties:
                      accessKeyID:
                        type: string
                      assumeRoleARN:
                        type: string
                      assumeRoleExternalID:
                        type: string
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      instanceProfileName:
                        type: string
                      nodePortsAllowedIPRange:
                        type: string
                      roleARN:
                        description: The IAM role, the control plane will use. The
                          control plane will perform an assume-role
                        type: string
                      routeTableID:
                        type: string
                      secretAccessKey:
                        type: string
                      securityGroupID:
                        type: string
                      vpcID:
                        type: string
                    required:
                    - instanceProfileName
                    - roleARN
                    - routeTableID
                    - securityGroupID
                    - vpcID
                    type: object
                  azure:
                    description: AzureCloudSpec specifies access credentials to Azure
                      cloud.
                    properties:
                      assignAvailabilitySet:
                        type: boolean
                      availabilitySet:
                        type: string
                      clientID:
                        type: string
                      clientSecret:
                        type: string
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      loadBalancerSKU:
                        description: LoadBalancerSKU sets the LB type that will be
                          used for the Azure cluster, possible values are "basic"
                          and "standard", if empty, "basic" will be used
                        enum:
                        - standard
                        - basic
                        type: string
                      nodePortsAllowedIPRange:
                        type: string
                      resourceGroup:
                        type: string
                      routeTable:
                        type: string
                      securityGroup:
                        type: string
                      subnet:
                        type: string
                      subscriptionID:
                        type: string
                      tenantID:
                        type: string
                      vnet:
                        type: string
                      vnetResourceGroup:
                        type: string
                    required:
                    - availabilitySet
                    - loadBalancerSKU
                    - resourceGroup
                    - routeTable
                    - securityGroup
                    - subnet
                    - vnet
                    - vnetResourceGroup
                    type: object
                  bringyourown:
                    description: BringYourOwnCloudSpec specifies access data for a
                      bring your own cluster.
                    type: object
                  dc:
                    description: DatacenterName where the users 'cloud' lives in.
                    type: string
                  digitalocean:
                    description: DigitaloceanCloudSpec specifies access data to DigitalOcean.
                    properties:
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      token:
                        type: string
                    type: object
                  fake:
                    description: FakeCloudSpec specifies access data for a fake cloud.
                    properties:
                      token:
                        type: string
                    type: object
                  gcp:
                    description: GCPCloudSpec specifies access data to GCP.
                    properties:
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      network:
                        type: string
                      nodePortsAllowedIPRange:
                        type: string
                      serviceAccount:
                        type: string
                      subnetwork:
                        type: string
                    required:
                    - network
                    - subnetwork
                    type: object
                  hetzner:
                    description: HetznerCloudSpec specifies access data to hetzner
                      cloud.
                    properties:
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      network:
                        description: Network is the pre-existing Hetzner network in
                          which the machines are running. While machines can be in
                          multiple networks, a single one must be chosen for the HCloud
                          CCM to work. If this is empty, the network configured on
                          the datacenter will be used.
                        type: string
                      token:
                        description: Token is used to authenticate with the Hetzner
                          cloud API.
                        type: string
                    type: object
                  kubevirt:
                    description: KubevirtCloudSpec specifies the access data to Kubevirt.
                    properties:
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      kubeconfig:
                        type: string
                    type: object
                  nutanix:
                    description: NutanixCloudSpec specifies the access data to Nutanix.
                      NUTANIX IMPLEMENTATION IS EXPERIMENTAL AND UNSUPPORTED.
                    properties:
                      clusterName:
                        description: ClusterName is the Nutanix cluster that this
                          user cluster will be deployed to.
                        type: string
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      password:
                        type: string
                      projectName:
                        description: ProjectName is the project that this cluster
                          is deployed into. If none is given, no project will be used.
                        type: string
                      proxyURL:
                        type: string
                      username:
                        type: string
                    required:
                    - clusterName
                    type: object
                  openstack:
                    description: OpenstackCloudSpec specifies access data to an OpenStack
                      cloud.
                    properties:
                      applicationCredentialID:
                        type: string
                      applicationCredentialSecret:
                        type: string
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      domain:
                        type: string
                      floatingIPPool:
                        description: "FloatingIPPool holds the name of the public
                          network The public network is reachable from the outside
                          world and should provide the pool of IP addresses to choose
                          from. \n When specified, all worker nodes will receive a
                          public ip from this floating ip pool \n Note that the network
                          is external if the \"External\" field is set to true"
                        type: string
                      network:
                        description: "Network holds the name of the internal network
                          When specified, all worker nodes will be attached to this
                          network. If not specified, a network, subnet & router will
                          be created \n Note that the network is internal if the \"External\"
                          field is set to false"
                        type: string
                      nodePortsAllowedIPRange:
                        type: string
                      password:
                        type: string
                      project:
                        description: project, formally known as tenant.
                        type: string
                      projectID:
                        description: project id, formally known as tenantID.
                        type: string
                      routerID:
                        type: string
                      securityGroups:
                        type: string
                      subnetID:
                        type: string
                      token:
                        description: Used internally during cluster creation
                        type: string
                      useOctavia:
                        description: "Whether or not to use Octavia for LoadBalancer
                          type of Service implementation instead of using Neutron-LBaaS.
                          Attention:Openstack CCM use Octavia as default load balancer
                          implementation since v1.17.0 \n Takes precedence over the
                          'use_octavia' flag provided at datacenter level if both
                          are specified."
                        type: boolean
                      useToken:
                        type: boolean
                      username:
                        type: string
                    required:
                    - floatingIPPool
                    - network
                    - routerID
                    - securityGroups
                    - subnetID
                    type: object
                  packet:
                    description: PacketCloudSpec specifies access data to a Packet
                      cloud.
                    properties:
                      apiKey:
                        type: string
                      billingCycle:
                        type: string
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      projectID:
                        type: string
                    required:
                    - billingCycle
                    type: object
                  providerName:
                    description: ProviderName is the name of the cloud provider used
                      for this cluster. This must match the given provider spec (e.g.
                      if the providerName is "aws", then the AWSCloudSpec must be
                      set)
                    type: string
                  vsphere:
                    description: VSphereCloudSpec specifies access data to VSphere
                      cloud.
                    properties:
                      credentialsReference:
                        description: GlobalObjectKeySelector is needed as we can not
                          use v1.SecretKeySelector because it is not cross namespace
                        properties:
                          apiVersion:
                            description: API version of the referent.
                            type: string
                          fieldPath:
                            description: 'If referring to a piece of an object instead
                              of an entire object, this string should contain a valid
                              JSON/Go field access statement, such as desiredState.manifest.containers[2].
                              For example, if the object reference is to a container
                              within a pod, this would take on a value like: "spec.containers{name}"
                              (where "name" refers to the name of the container that
                              triggered the event) or if no container name is specified
                              "spec.containers[2]" (container with index 2 in this
                              pod). This syntax is chosen only to have some well-defined
                              way of referencing a part of an object. TODO: this design
                              is not final and this field is subject to change in
                              the future.'
                            type: string
                          key:
                            type: string
                          kind:
                            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                          resourceVersion:
                            description: 'Specific resourceVersion to which this reference
                              is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
                            type: string
                          uid:
                            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
                            type: string
                        type: object
                      datastore:
                        description: Datastore to be used for storing virtual machines
                          and as a default for dynamic volume provisioning, it is
                          mutually exclusive with DatastoreCluster.
                        type: string
                      datastoreCluster:
                        description: DatastoreCluster to be used for storing virtual
                          machines, it is mutually exclusive with Datastore.
                        type: string
                      folder:
                        description: Folder is the folder to be used to group the
                          provisioned virtual machines.
                        type: string
                      infraManagementUser:
                        description: This user will be used for everything except
                          cloud provider functionality
                        properties:
                          password:
                            type: string
                          username:
                            type: string
                        type: object
                      password:
                        description: Password is the vSphere user password.
                        type: string
                      resourcePool:
                        description: ResourcePool is used to manage resources such
                          as cpu and memory for vSphere virtual machines. The resource
                          pool should be defined on vSphere cluster level.
                        type: string
                      storagePolicy:
                        description: StoragePolicy to be used for storage provisioning
                        type: string
                      username:
                        description: Username is the vSphere user name.
                        type: string
                      vmNetName:
                        description: VMNetName is the name of the vSphere network.
                        type: string
                    required:
                    - infraManagementUser
                    - storagePolicy
                    - vmNetName
                    type: object
                required:
                - dc
                - providerName
                type: object
              clusterNetwork:
                description: ClusterNetworkingConfig specifies the different networking
                  parameters for a cluster.
                properties:
                  dnsDomain:
                    description: Domain name for services.
                    type: string
                  ipvs:
                    description: IPVS defines kube-proxy ipvs configuration options
                    properties:
                      strictArp:
                        default: true
                        description: StrictArp configure arp_ignore and arp_announce
                          to avoid answering ARP queries from kube-ipvs0 interface.
                          defaults to true.
                        type: boolean
                    type: object
                  konnectivityEnabled:
                    description: KonnectivityEnabled enables konnectivity for controlplane
                      to node network communication.
                    type: boolean
                  nodeLocalDNSCacheEnabled:
                    default: true
                    description: NodeLocalDNSCacheEnabled controls whether the NodeLocal
                      DNS Cache feature is enabled. Defaults to true.
                    type: boolean
                  pods:
                    description: The network ranges from which POD networks are allocated.
                    properties:
                      cidrBlocks:
                        items:
                          type: string
                        type: array
                    required:
                    - cidrBlocks
                    type: object
                  proxyMode:
                    default: ipvs
                    description: ProxyMode defines the kube-proxy mode ("ipvs" / "iptables"
                      / "ebpf"). Defaults to "ipvs". "ebpf" disables kube-proxy and
                      requires CNI support.
                    enum:
                    - ipvs
                    - iptables
                    - ebpf
                    type: string
                  services:
                    description: The network ranges from which service VIPs are allocated.
                    properties:
                      cidrBlocks:
                        items:
                          type: string
                        type: array
                    required:
                    - cidrBlocks
                    type: object
                required:
                - dnsDomain
                - pods
                - proxyMode
                - services
                type: object
              cniPlugin:
                description: CNIPlugin contains the spec of the CNI plugin to be installed
                  in the cluster.
                properties:
                  type:
                    description: CNIPluginType define the type of CNI plugin installed.
                      e.g. Canal.
                    enum:
                    - canal
                    - cilium
                    - none
                    type: string
                  version:
                    type: string
                required:
                - type
                - version
                type: object
              componentsOverride:
                description: Optional component specific overrides
                properties:
                  apiserver:
                    properties:
                      endpointReconcilingDisabled:
                        type: boolean
                      nodePortRange:
                        type: string
                      replicas:
                        format: int32
                        type: integer
                      resources:
                        description: ResourceRequirements describes the compute resource
                          requirements.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Limits describes the maximum amount of compute
                              resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Requests describes the minimum amount of
                              compute resources required. If Requests is omitted for
                              a container, it defaults to Limits if that is explicitly
                              specified, otherwise to an implementation-defined value.
                              More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                        type: object
                      tolerations:
                        items:
                          description: The pod this Toleration is attached to tolerates
                            any taint that matches the triple <key,value,effect> using
                            the matching operator <operator>.
                          properties:
                            effect:
                              description: Effect indicates the taint effect to match.
                                Empty means match all taint effects. When specified,
                                allowed values are NoSchedule, PreferNoSchedule and
                                NoExecute.
                              type: string
                            key:
                              description: Key is the taint key that the toleration
                                applies to. Empty means match all taint keys. If the
                                key is empty, operator must be Exists; this combination
                                means to match all values and all keys.
                              type: string
                            operator:
                              description: Operator represents a key's relationship
                                to the value. Valid operators are Exists and Equal.
                                Defaults to Equal. Exists is equivalent to wildcard
                                for value, so that a pod can tolerate all taints of
                                a particular category.
                              type: string
                            tolerationSeconds:
                              description: TolerationSeconds represents the period
                                of time the toleration (which must be of effect NoExecute,
                                otherwise this field is ignored) tolerates the taint.
                                By default, it is not set, which means tolerate the
                                taint forever (do not evict). Zero and negative values
                                will be treated as 0 (evict immediately) by the system.
                              format: int64
                              type: integer
                            value:
                              description: Value is the taint value the toleration
                                matches to. If the operator is Exists, the value should
                                be empty, otherwise just a regular string.
                              type: string
                          type: object
                        type: array
                    type: object
                  controllerManager:
                    properties:
                      leaderElection:
                        properties:
                          leaseDurationSeconds:
                            description: LeaseDurationSeconds is the duration in seconds
                              that non-leader candidates will wait to force acquire
                              leadership. This is measured against time of last observed
                              ack.
                            format: int32
                            type: integer
                          renewDeadlineSeconds:
                            description: RenewDeadlineSeconds is the duration in seconds
                              that the acting controlplane will retry refreshing leadership
                              before giving up.
                            format: int32
                            type: integer
                          retryPeriodSeconds:
                            description: RetryPeriodSeconds is the duration in seconds
                              the LeaderElector clients should wait between tries
                              of actions.
                            format: int32
                            type: integer
                        type: object
                      replicas:
                        format: int32
                        type: integer
                      resources:
                        description: ResourceRequirements describes the compute resource
                          requirements.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Limits describes the maximum amount of compute
                              resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Requests describes the minimum amount of
                              compute resources required. If Requests is omitted for
                              a container, it defaults to Limits if that is explicitly
                              specified, otherwise to an implementation-defined value.
                              More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                        type: object
                      tolerations:
                        items:
                          description: The pod this Toleration is attached to tolerates
                            any taint that matches the triple <key,value,effect> using
                            the matching operator <operator>.
                          properties:
                            effect:
                              description: Effect indicates the taint effect to match.
                                Empty means match all taint effects. When specified,
                                allowed values are NoSchedule, PreferNoSchedule and
                                NoExecute.
                              type: string
                            key:
                              description: Key is the taint key that the toleration
                                applies to. Empty means match all taint keys. If the
                                key is empty, operator must be Exists; this combination
                                means to match all values and all keys.
                              type: string
                            operator:
                              description: Operator represents a key's relationship
                                to the value. Valid operators are Exists and Equal.
                                Defaults to Equal. Exists is equivalent to wildcard
                                for value, so that a pod can tolerate all taints of
                                a particular category.
                              type: string
                            tolerationSeconds:
                              description: TolerationSeconds represents the period
                                of time the toleration (which must be of effect NoExecute,
                                otherwise this field is ignored) tolerates the taint.
                                By default, it is not set, which means tolerate the
                                taint forever (do not evict). Zero and negative values
                                will be treated as 0 (evict immediately) by the system.
                              format: int64
                              type: integer
                            value:
                              description: Value is the taint value the toleration
                                matches to. If the operator is Exists, the value should
                                be empty, otherwise just a regular string.
                              type: string
                          type: object
                        type: array
                    type: object
                  etcd:
                    properties:
                      clusterSize:
                        format: int32
                        type: integer
                      diskSize:
                        anyOf:
                        - type: integer
                        - type: string
                        pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                        x-kubernetes-int-or-string: true
                      resources:
                        description: ResourceRequirements describes the compute resource
                          requirements.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Limits describes the maximum amount of compute
                              resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Requests describes the minimum amount of
                              compute resources required. If Requests is omitted for
                              a container, it defaults to Limits if that is explicitly
                              specified, otherwise to an implementation-defined value.
                              More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                        type: object
                      storageClass:
                        type: string
                      tolerations:
                        items:
                          description: The pod this Toleration is attached to tolerates
                            any taint that matches the triple <key,value,effect> using
                            the matching operator <operator>.
                          properties:
                            effect:
                              description: Effect indicates the taint effect to match.
                                Empty means match all taint effects. When specified,
                                allowed values are NoSchedule, PreferNoSchedule and
                                NoExecute.
                              type: string
                            key:
                              description: Key is the taint key that the toleration
                                applies to. Empty means match all taint keys. If the
                                key is empty, operator must be Exists; this combination
                                means to match all values and all keys.
                              type: string
                            operator:
                              description: Operator represents a key's relationship
                                to the value. Valid operators are Exists and Equal.
                                Defaults to Equal. Exists is equivalent to wildcard
                                for value, so that a pod can tolerate all taints of
                                a particular category.
                              type: string
                            tolerationSeconds:
                              description: TolerationSeconds represents the period
                                of time the toleration (which must be of effect NoExecute,
                                otherwise this field is ignored) tolerates the taint.
                                By default, it is not set, which means tolerate the
                                taint forever (do not evict). Zero and negative values
                                will be treated as 0 (evict immediately) by the system.
                              format: int64
                              type: integer
                            value:
                              description: Value is the taint value the toleration
                                matches to. If the operator is Exists, the value should
                                be empty, otherwise just a regular string.
                              type: string
                          type: object
                        type: array
                    type: object
                  nodePortProxyEnvoy:
                    properties:
                      dockerRepository:
                        description: DockerRepository is the repository containing
                          the component's image.
                        type: string
                      resources:
                        description: Resources describes the requested and maximum
                          allowed CPU/memory usage.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Limits describes the maximum amount of compute
                              resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Requests describes the minimum amount of
                              compute resources required. If Requests is omitted for
                              a container, it defaults to Limits if that is explicitly
                              specified, otherwise to an implementation-defined value.
                              More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                        type: object
                    type: object
                  prometheus:
                    properties:
                      resources:
                        description: ResourceRequirements describes the compute resource
                          requirements.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Limits describes the maximum amount of compute
                              resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Requests describes the minimum amount of
                              compute resources required. If Requests is omitted for
                              a container, it defaults to Limits if that is explicitly
                              specified, otherwise to an implementation-defined value.
                              More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                        type: object
                    type: object
                  scheduler:
                    properties:
                      leaderElection:
                        properties:
                          leaseDurationSeconds:
                            description: LeaseDurationSeconds is the duration in seconds
                              that non-leader candidates will wait to force acquire
                              leadership. This is measured against time of last observed
                              ack.
                            format: int32
                            type: integer
                          renewDeadlineSeconds:
                            description: RenewDeadlineSeconds is the duration in seconds
                              that the acting controlplane will retry refreshing leadership
                              before giving up.
                            format: int32
                            type: integer
                          retryPeriodSeconds:
                            description: RetryPeriodSeconds is the duration in seconds
                              the LeaderElector clients should wait between tries
                              of actions.
                            format: int32
                            type: integer
                        type: object
                      replicas:
                        format: int32
                        type: integer
                      resources:
                        description: ResourceRequirements describes the compute resource
                          requirements.
                        properties:
                          limits:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Limits describes the maximum amount of compute
                              resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                          requests:
                            additionalProperties:
                              anyOf:
                              - type: integer
                              - type: string
                              pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                              x-kubernetes-int-or-string: true
                            description: 'Requests describes the minimum amount of
                              compute resources required. If Requests is omitted for
                              a container, it defaults to Limits if that is explicitly
                              specified, otherwise to an implementation-defined value.
                              More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                            type: object
                        type: object
                      tolerations:
                        items:
                          description: The pod this Toleration is attached to tolerates
                            any taint that matches the triple <key,value,effect> using
                            the matching operator <operator>.
                          properties:
                            effect:
                              description: Effect indicates the taint effect to match.
                                Empty means match all taint effects. When specified,
                                allowed values are NoSchedule, PreferNoSchedule and
                                NoExecute.
                              type: string
                            key:
                              description: Key is the taint key that the toleration
                                applies to. Empty means match all taint keys. If the
                                key is empty, operator must be Exists; this combination
                                means to match all values and all keys.
                              type: string
                            operator:
                              description: Operator represents a key's relationship
                                to the value. Valid operators are Exists and Equal.
                                Defaults to Equal. Exists is equivalent to wildcard
                                for value, so that a pod can tolerate all taints of
                                a particular category.
                              type: string
                            tolerationSeconds:
                              description: TolerationSeconds represents the period
                                of time the toleration (which must be of effect NoExecute,
                                otherwise this field is ignored) tolerates the taint.
                                By default, it is not set, which means tolerate the
                                taint forever (do not evict). Zero and negative values
                                will be treated as 0 (evict immediately) by the system.
                              format: int64
                              type: integer
                            value:
                              description: Value is the taint value the toleration
                                matches to. If the operator is Exists, the value should
                                be empty, otherwise just a regular string.
                              type: string
                          type: object
                        type: array
                    type: object
                required:
                - apiserver
                - controllerManager
                - etcd
                - nodePortProxyEnvoy
                - prometheus
                - scheduler
                type: object
              containerRuntime:
                default: containerd
                description: ContainerRuntime to use, i.e. Docker or containerd. By
                  default containerd will be used.
                enum:
                - docker
                - containerd
                type: string
              debugLog:
                description: DebugLog enables more verbose logging by KKP's usercluster-controller-manager.
                type: boolean
              enableOperatingSystemManager:
                description: EnableOperatingSystemManager enables OSM which in-turn
                  is responsible for creating and managing worker node configuration
                type: boolean
              enableUserSSHKeyAgent:
                description: EnableUserSSHKeyAgent control whether the UserSSHKeyAgent
                  will be deployed in the user cluster or not. If it was enabled,
                  the agent will be deployed and used to sync the user ssh keys, that
                  the user attach to the created cluster. If the agent was disabled,
                  it won't be deployed in the user cluster, thus after the cluster
                  creation any attached ssh keys won't be synced to the worker nodes.
                  Once the agent is enabled/disabled it cannot be changed after the
                  cluster is being created.
                type: boolean
              eventRateLimitConfig:
                description: EventRateLimitConfig allows configuring the EventRateLimit
                  admission plugin (if enabled via useEventRateLimitAdmissionPlugin)
                  to create limits on Kubernetes event generation. The EventRateLimit
                  plugin is capable of comparing incoming Events to several configured
                  buckets based on the type of event rate limit.
                properties:
                  namespace:
                    properties:
                      burst:
                        format: int32
                        type: integer
                      cacheSize:
                        format: int32
                        type: integer
                      qps:
                        format: int32
                        type: integer
                    required:
                    - burst
                    - qps
                    type: object
                  server:
                    properties:
                      burst:
                        format: int32
                        type: integer
                      cacheSize:
                        format: int32
                        type: integer
                      qps:
                        format: int32
                        type: integer
                    required:
                    - burst
                    - qps
                    type: object
                  sourceAndObject:
                    properties:
                      burst:
                        format: int32
                        type: integer
                      cacheSize:
                        format: int32
                        type: integer
                      qps:
                        format: int32
                        type: integer
                    required:
                    - burst
                    - qps
                    type: object
                  user:
                    properties:
                      burst:
                        format: int32
                        type: integer
                      cacheSize:
                        format: int32
                        type: integer
                      qps:
                        format: int32
                        type: integer
                    required:
                    - burst
                    - qps
                    type: object
                type: object
              exposeStrategy:
                description: ExposeStrategy is the approach we use to expose this
                  cluster, either via NodePort or via a dedicated LoadBalancer
                enum:
                - NodePort
                - LoadBalancer
                - Tunneling
                type: string
              features:
                additionalProperties:
                  type: boolean
                description: Feature flags This unfortunately has to be a string map,
                  because we use it in templating and that can not cope with string
                  types
                type: object
              humanReadableName:
                description: HumanReadableName is the cluster name provided by the
                  user
                type: string
              machineNetworks:
                items:
                  description: MachineNetworkingConfig specifies the networking parameters
                    used for IPAM.
                  properties:
                    cidr:
                      type: string
                    dnsServers:
                      items:
                        type: string
                      type: array
                    gateway:
                      type: string
                  required:
                  - cidr
                  - dnsServers
                  - gateway
                  type: object
                type: array
              mla:
                description: MLA contains monitoring, logging and alerting related
                  settings for the user cluster.
                properties:
                  loggingEnabled:
                    description: LoggingEnabled is the flag for enabling logging in
                      user cluster.
                    type: boolean
                  loggingResources:
                    description: LoggingResources is the resource requirements for
                      user cluster promtail.
                    properties:
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  monitoringEnabled:
                    description: MonitoringEnabled is the flag for enabling monitoring
                      in user cluster.
                    type: boolean
                  monitoringReplicas:
                    description: MonitoringReplicas is the number of desired pods
                      of user cluster prometheus deployment.
                    format: int32
                    type: integer
                  monitoringResources:
                    description: MonitoringResources is the resource requirements
                      for user cluster prometheus.
                    properties:
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                type: object
              oidc:
                properties:
                  clientID:
                    type: string
                  clientSecret:
                    type: string
                  extraScopes:
                    type: string
                  groupsClaim:
                    type: string
                  issuerURL:
                    type: string
                  requiredClaim:
                    type: string
                  usernameClaim:
                    type: string
                type: object
              opaIntegration:
                description: OPAIntegration is a preview feature that enables OPA
                  integration with Kubermatic for the cluster. Enabling it causes
                  gatekeeper and its resources to be deployed on the user cluster.
                  By default it is disabled.
                properties:
                  auditResources:
                    description: AuditResources is the resource requirements for user
                      cluster gatekeeper audit.
                    properties:
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  controllerResources:
                    description: ControllerResources is the resource requirements
                      for user cluster gatekeeper controller.
                    properties:
                      limits:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Limits describes the maximum amount of compute
                          resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                      requests:
                        additionalProperties:
                          anyOf:
                          - type: integer
                          - type: string
                          pattern: ^(\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?$
                          x-kubernetes-int-or-string: true
                        description: 'Requests describes the minimum amount of compute
                          resources required. If Requests is omitted for a container,
                          it defaults to Limits if that is explicitly specified, otherwise
                          to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/'
                        type: object
                    type: object
                  enabled:
                    description: Enabled is the flag for enabling OPA integration
                    type: boolean
                  experimentalEnableMutation:
                    description: Enable mutation
                    type: boolean
                  webhookTimeoutSeconds:
                    default: 10
                    description: WebhookTimeout is the timeout that is set for the
                      gatekeeper validating webhook admission review calls. By default
                      10 seconds.
                    format: int32
                    type: integer
                type: object
              pause:
                description: Pause tells that this cluster is currently not managed
                  by the controller. It indicates that the user needs to do some action
                  to resolve the pause.
                type: boolean
              pauseReason:
                description: PauseReason is the reason why the cluster is no being
                  managed.
                type: string
              podNodeSelectorAdmissionPluginConfig:
                additionalProperties:
                  type: string
                description: 'PodNodeSelectorAdmissionPluginConfig provides the configuration
                  for the PodNodeSelector. It''s used by the backend to create a configuration
                  file for this plugin. The key:value from the map is converted to
                  the namespace:<node-selectors-labels> in the file. The format in
                  a file: podNodeSelectorPluginConfig:  clusterDefaultNodeSelector:
                  <node-selectors-labels>  namespace1: <node-selectors-labels>  namespace2:
                  <node-selectors-labels>'
                type: object
              serviceAccount:
                description: ServiceAccount contains service account related settings
                  for the kube-apiserver of user cluster.
                properties:
                  apiAudiences:
                    description: APIAudiences are the Identifiers of the API If this
                      is not specified, it will be set to a single element list containing
                      the issuer URL
                    items:
                      type: string
                    type: array
                  issuer:
                    description: Issuer is the identifier of the service account token
                      issuer If this is not specified, it will be set to the URL of
                      apiserver by default
                    type: string
                  tokenVolumeProjectionEnabled:
                    type: boolean
                type: object
              updateWindow:
                properties:
                  length:
                    type: string
                  start:
                    type: string
                type: object
              useEventRateLimitAdmissionPlugin:
                type: boolean
              usePodNodeSelectorAdmissionPlugin:
                type: boolean
              usePodSecurityPolicyAdmissionPlugin:
                type: boolean
              version:
                description: Version defines the wanted version of the control plane
                type: string
            required:
            - cloud
            - clusterNetwork
            - exposeStrategy
            - humanReadableName
            - pause
            - version
            type: object
          status:
            description: ClusterStatus stores status information about a cluster.
            properties:
              cloudMigrationRevision:
                description: CloudMigrationRevision describes the latest version of
                  the migration that has been done It is used to avoid redundant and
                  potentially costly migrations
                type: integer
              conditions:
                additionalProperties:
                  properties:
                    kubermaticVersion:
                      description: KubermaticVersion current kubermatic version.
                      type: string
                    lastHeartbeatTime:
                      description: Last time we got an update on a given condition.
                      format: date-time
                      type: string
                    lastTransitionTime:
                      description: Last time the condition transit from one status
                        to another.
                      format: date-time
                      type: string
                    message:
                      description: Human readable message indicating details about
                        last transition.
                      type: string
                    reason:
                      description: (brief) reason for the condition's last transition.
                      type: string
                    status:
                      description: Status of the condition, one of True, False, Unknown.
                      type: string
                  required:
                  - kubermaticVersion
                  - lastHeartbeatTime
                  - status
                  type: object
                description: Conditions contains conditions the cluster is in, its
                  primary use case is status signaling between controllers or between
                  controllers and the API
                type: object
              errorMessage:
                description: ErrorMessage contains a default error message in case
                  the controller encountered an error. Will be reset if the error
                  was resolved
                type: string
              errorReason:
                description: ErrorReason contains a error reason in case the controller
                  encountered an error. Will be reset if the error was resolved
                enum:
                - InvalidConfiguration
                - UnsupportedChange
                - ReconcileError
                type: string
              extendedHealth:
                description: ExtendedHealth exposes information about the current
                  health state. Extends standard health status for new states.
                properties:
                  alertmanagerConfig:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  apiserver:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  cloudProviderInfrastructure:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  controller:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  etcd:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  gatekeeperAudit:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  gatekeeperController:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  logging:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  machineController:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  mlaGateway:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  monitoring:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  openvpn:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  scheduler:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                  userClusterControllerManager:
                    enum:
                    - HealthStatusDown
                    - HealthStatusUp
                    - HealthStatusProvisioning
                    type: string
                type: object
              inheritedLabels:
                additionalProperties:
                  type: string
                description: InheritedLabels are labels the cluster inherited from
                  the project. They are read-only for users.
                type: object
              lastProviderReconciliation:
                description: LastProviderReconciliation is the time when the cloud
                  provider resources were last fully reconciled (during normal cluster
                  reconciliation, KKP does not re-check things like security groups,
                  networks etc.).
                format: date-time
                type: string
              lastUpdated:
                format: date-time
                type: string
              namespaceName:
                description: NamespaceName defines the namespace the control plane
                  of this cluster is deployed in
                type: string
              userEmail:
                description: UserEmail contains the email of the owner of this cluster
                type: string
              userName:
                description: UserName contains the name of the owner of this cluster
                type: string
            type: object
        type: object
    served: true
    storage: true
    subresources:
      status: {}
status:
  acceptedNames:
    kind: ""
    plural: ""
  conditions: []
  storedVersions: []
`

const sampleCrd = `
---
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: samples.kubermatic.k8c.io
spec:
  group: kubermatic.k8c.io
  names:
    kind: Sample
    listKind: SampleList
    plural: Samples
    singular: Sample
  scope: Cluster
  versions:
    - name: v1
      schema:
        openAPIV3Schema:
          description: ""
          properties:
            apiVersion:
              description: ""
              type: string
            kind:
              description: ""
              type: string
            metadata:
              type: object
            spec:
              description: spec
              properties:
                val1:
                  description: val1
                  type: string
                  enum:
                    - valid1
          type: object
      served: true
      storage: true
    - name: v2
      schema:
        openAPIV3Schema:
          description: ""
          properties:
            apiVersion:
              description: ""
              type: string
            kind:
              description: ""
              type: string
            metadata:
              type: object
            spec:
              description: spec
              properties:
                val1:
                  description: val1
                  type: string
                  enum:
                    - valid1
                    - valid2
          type: object
      served: true
      storage: true
`

func TestTest(t *testing.T) {
	f, err := os.Open("crd/k8c.io/kubermatic.k8c.io_clusters.yaml")
	if err != nil {
		t.Fatal(err)
	}

	tt := map[string]struct {
		in io.Reader
	}{
		"cluster from file": {
			f,
		},
		// "cluster from string": {
		// 	strings.NewReader(clusterCrd),
		// },
		"sample from string": {
			strings.NewReader(sampleCrd),
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			v, err := ValidatorFromCRD(tc.in, "v1")
			if err != nil {
				t.Fatal(err)
			}

			if len(v.Schema.SchemaProps.Properties) == 0 {
				t.Error("SchemaProps Empty")
			}
		})
	}

}
