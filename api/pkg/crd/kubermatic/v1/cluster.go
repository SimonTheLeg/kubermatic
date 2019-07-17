package v1

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/kubermatic/kubermatic/api/pkg/semver"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ClusterResourceName represents "Resource" defined in Kubernetes
	ClusterResourceName = "clusters"

	// ClusterKindName represents "Kind" defined in Kubernetes
	ClusterKindName = "Cluster"

	// AnnotationNameClusterAutoscalerEnabled is the name of the annotation that is being
	// used to determine if the cluster-autoscaler is enabled for this cluster. It is
	// enabled when this Annotation is set with any value
	AnnotationNameClusterAutoscalerEnabled = "kubermatic.io/cluster-autoscaler-enabled"
)

const (
	WorkerNameLabelKey = "worker-name"
	ProjectIDLabelKey  = "project-id"
)

//+genclient
//+genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cluster is the object representing a cluster.
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec    ClusterSpec    `json:"spec"`
	Address ClusterAddress `json:"address,omitempty"`
	Status  ClusterStatus  `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList specifies a list of clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Cluster `json:"items"`
}

// ClusterSpec specifies the data for a new cluster.
type ClusterSpec struct {
	Cloud           CloudSpec                 `json:"cloud"`
	ClusterNetwork  ClusterNetworkingConfig   `json:"clusterNetwork"`
	MachineNetworks []MachineNetworkingConfig `json:"machineNetworks,omitempty"`

	// Version defines the wanted version of the control plane
	Version semver.Semver `json:"version"`
	// MasterVersion is Deprecated
	MasterVersion string `json:"masterVersion,omitempty"`

	// HumanReadableName is the cluster name provided by the user
	HumanReadableName string `json:"humanReadableName"`

	// ExposeStrategy is the approach we use to expose this cluster, either via NodePort
	// or via a dedicated LoadBalancer
	ExposeStrategy corev1.ServiceType `json:"exposeStrategy"`

	// Pause tells that this cluster is currently not managed by the controller.
	// It indicates that the user needs to do some action to resolve the pause.
	Pause bool `json:"pause"`
	// PauseReason is the reason why the cluster is no being managed.
	PauseReason string `json:"pauseReason,omitempty"`

	// Optional component specific overrides
	ComponentsOverride ComponentSettings `json:"componentsOverride"`

	OIDC OIDCSettings `json:"oidc,omitempty"`
}

type OIDCSettings struct {
	IssuerURL     string `json:"issuerUrl,omitempty"`
	ClientID      string `json:"clientId,omitempty"`
	ClientSecret  string `json:"clientSecret,omitempty"`
	UsernameClaim string `json:"usernameClaim,omitempty"`
	GroupsClaim   string `json:"groupsClaim,omitempty"`
	RequiredClaim string `json:"requiredClaim,omitempty"`
	ExtraScopes   string `json:"extraScopes,omitempty"`
}

type ComponentSettings struct {
	Apiserver         DeploymentSettings  `json:"apiserver"`
	ControllerManager DeploymentSettings  `json:"controllerManager"`
	Scheduler         DeploymentSettings  `json:"scheduler"`
	Etcd              StatefulSetSettings `json:"etcd"`
	Prometheus        StatefulSetSettings `json:"prometheus"`
}

type DeploymentSettings struct {
	Replicas  *int32                       `json:"replicas,omitempty"`
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type StatefulSetSettings struct {
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// ClusterNetworkingConfig specifies the different networking
// parameters for a cluster.
type ClusterNetworkingConfig struct {
	// The network ranges from which service VIPs are allocated.
	Services NetworkRanges `json:"services"`

	// The network ranges from which POD networks are allocated.
	Pods NetworkRanges `json:"pods"`

	// Domain name for services.
	DNSDomain string `json:"dnsDomain"`
}

// MachineNetworkingConfig specifies the networking parameters used for IPAM.
type MachineNetworkingConfig struct {
	CIDR       string   `json:"cidr"`
	Gateway    string   `json:"gateway"`
	DNSServers []string `json:"dnsServers"`
}

// NetworkRanges represents ranges of network addresses.
type NetworkRanges struct {
	CIDRBlocks []string `json:"cidrBlocks"`
}

// ClusterAddress stores access and address information of a cluster.
type ClusterAddress struct {
	// URL under which the Apiserver is available
	URL string `json:"url"`
	// Port is the port the API server listens on
	Port int32 `json:"port"`
	// ExternalName is the DNS name for this cluster
	ExternalName string `json:"externalName"`
	// InternalName is the seed cluster internal absolute DNS name to the API server
	InternalName string `json:"internalURL"`
	// AdminToken is the token for the kubeconfig, the user can download
	AdminToken string `json:"adminToken"`
	// IP is the external IP under which the apiserver is available
	IP string `json:"ip"`
}

type ClusterConditionType string

type ClusterCondition struct {
	// Type of cluster condition.
	Type ClusterConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	// +optional
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime,omitempty"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// ClusterStatus stores status information about a cluster.
type ClusterStatus struct {
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
	// ExtendedHealth exposes information about the current health state.
	// Extends standard health status for new states.
	ExtendedHealth ExtendedClusterHealth `json:"extendedHealth,omitempty"`

	// Deprecated
	RootCA *KeyCert `json:"rootCA,omitempty"`
	// Deprecated
	ApiserverCert *KeyCert `json:"apiserverCert,omitempty"`
	// Deprecated
	KubeletCert *KeyCert `json:"kubeletCert,omitempty"`
	// Deprecated
	ApiserverSSHKey *RSAKeys `json:"apiserverSshKey,omitempty"`
	// Deprecated
	ServiceAccountKey Bytes `json:"serviceAccountKey,omitempty"`
	// NamespaceName defines the namespace the control plane of this cluster is deployed in
	NamespaceName string `json:"namespaceName"`

	// UserName contains the name of the owner of this cluster
	UserName string `json:"userName,omitempty"`
	// UserEmail contains the email of the owner of this cluster
	UserEmail string `json:"userEmail"`

	// ErrorReason contains a error reason in case the controller encountered an error. Will be reset if the error was resolved
	ErrorReason *ClusterStatusError `json:"errorReason,omitempty"`
	// ErrorMessage contains a defauled error message in case the controller encountered an error. Will be reset if the error was resolved
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Conditions contains conditions the cluster is in, its primary use case is status signaling between controllers or between
	// controllers and the API
	Conditions []ClusterCondition `json:"conditions,omitempty"`

	// CloudMigrationRevision describes the latest version of the migration that has been done
	// It is used to avoid redundant and potentially costly migrations
	CloudMigrationRevision int `json:"cloudMigrationRevision"`
}

type ClusterStatusError string

const (
	InvalidConfigurationClusterError ClusterStatusError = "InvalidConfiguration"
	UnsupportedChangeClusterError    ClusterStatusError = "UnsupportedChange"
	ReconcileClusterError            ClusterStatusError = "ReconcileError"
)

// CloudSpec mutually stores access data to a cloud provider.
type CloudSpec struct {
	// DatacenterName where the users 'cloud' lives in.
	DatacenterName string `json:"dc"`

	Fake         *FakeCloudSpec         `json:"fake,omitempty"`
	Digitalocean *DigitaloceanCloudSpec `json:"digitalocean,omitempty"`
	BringYourOwn *BringYourOwnCloudSpec `json:"bringyourown,omitempty"`
	AWS          *AWSCloudSpec          `json:"aws,omitempty"`
	Azure        *AzureCloudSpec        `json:"azure,omitempty"`
	Openstack    *OpenstackCloudSpec    `json:"openstack,omitempty"`
	Packet       *PacketCloudSpec       `json:"packet,omitempty"`
	Hetzner      *HetznerCloudSpec      `json:"hetzner,omitempty"`
	VSphere      *VSphereCloudSpec      `json:"vsphere,omitempty"`
	GCP          *GCPCloudSpec          `json:"gcp,omitempty"`
}

// KeyCert is a pair of key and cert.
type KeyCert struct {
	Key  Bytes `json:"key"`
	Cert Bytes `json:"cert"`
}

// RSAKeys is a pair of private and public key where the key is not published to the API client.
type RSAKeys struct {
	PrivateKey Bytes `json:"privateKey"`
	PublicKey  Bytes `json:"publicKey"`
}

type Bytes []byte

// FakeCloudSpec specifies access data for a fake cloud.
type FakeCloudSpec struct {
	Token string `json:"token,omitempty"`
}

// DigitaloceanCloudSpec specifies access data to DigitalOcean.
type DigitaloceanCloudSpec struct {
	Token string `json:"token"` // Token is used to authenticate with the DigitalOcean API.
}

// HetznerCloudSpec specifies access data to hetzner cloud.
type HetznerCloudSpec struct {
	Token string `json:"token"` // Token is used to authenticate with the Hetzner cloud API.
}

// AzureCloudSpec specifies acceess credentials to Azure cloud.
type AzureCloudSpec struct {
	TenantID       string `json:"tenantID"`
	SubscriptionID string `json:"subscriptionID"`
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`

	ResourceGroup   string `json:"resourceGroup"`
	VNetName        string `json:"vnet"`
	SubnetName      string `json:"subnet"`
	RouteTableName  string `json:"routeTable"`
	SecurityGroup   string `json:"securityGroup"`
	AvailabilitySet string `json:"availabilitySet"`
}

// VSphereCredentials credentials represents a credential for accessing vSphere
type VSphereCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// VSphereCloudSpec specifies access data to VSphere cloud.
type VSphereCloudSpec struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	VMNetName string `json:"vmNetName"`
	Folder    string `json:"folder,omitempty"`

	// This user will be used for everything except cloud provider functionality
	InfraManagementUser VSphereCredentials `json:"infraManagementUser"`
}

// BringYourOwnCloudSpec specifies access data for a bring your own cluster.
type BringYourOwnCloudSpec struct{}

// AWSCloudSpec specifies access data to Amazon Web Services.
type AWSCloudSpec struct {
	AccessKeyID         string `json:"accessKeyId"`
	SecretAccessKey     string `json:"secretAccessKey"`
	VPCID               string `json:"vpcId"`
	SubnetID            string `json:"subnetId"`
	RoleName            string `json:"roleName"`
	RoleARN             string `json:"roleARN"`
	RouteTableID        string `json:"routeTableId"`
	InstanceProfileName string `json:"instanceProfileName"`
	SecurityGroupID     string `json:"securityGroupID"`

	AvailabilityZone string `json:"availabilityZone"`
}

// OpenstackCloudSpec specifies access data to an OpenStack cloud.
type OpenstackCloudSpec struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tenant   string `json:"tenant"`
	TenantID string `json:"tenantID"`
	Domain   string `json:"domain"`
	// Network holds the name of the internal network
	// When specified, all worker nodes will be attached to this network. If not specified, a network, subnet & router will be created
	//
	// Note that the network is internal if the "External" field is set to false
	Network        string `json:"network"`
	SecurityGroups string `json:"securityGroups"`
	// FloatingIPPool holds the name of the public network
	// The public network is reachable from the outside world
	// and should provide the pool of IP addresses to choose from.
	//
	// When specified, all worker nodes will receive a public ip from this floating ip pool
	//
	// Note that the network is external if the "External" field is set to true
	FloatingIPPool string `json:"floatingIpPool"`
	RouterID       string `json:"routerID"`
	SubnetID       string `json:"subnetID"`
}

// PacketCloudSpec specifies access data to a Packet cloud.
type PacketCloudSpec struct {
	APIKey       string `json:"apiKey"`
	ProjectID    string `json:"projectID"`
	BillingCycle string `json:"billingCycle"`
}

// GCPCloudSpec specifies access data to GCP.
type GCPCloudSpec struct {
	ServiceAccount string `json:"serviceAccount"`
	Network        string `json:"network"`
	Subnetwork     string `json:"subnetwork"`
}

type HealthStatus int

const (
	HealthStatusDown         HealthStatus = iota
	HealthStatusUp           HealthStatus = iota
	HealthStatusProvisioning HealthStatus = iota
)

// ExtendedClusterHealth stores health information of a cluster.
type ExtendedClusterHealth struct {
	Apiserver                    HealthStatus `json:"apiserver"`
	Scheduler                    HealthStatus `json:"scheduler"`
	Controller                   HealthStatus `json:"controller"`
	MachineController            HealthStatus `json:"machineController"`
	Etcd                         HealthStatus `json:"etcd"`
	OpenVPN                      HealthStatus `json:"openvpn"`
	CloudProviderInfrastructure  HealthStatus `json:"cloudProviderInfrastructure"`
	UserClusterControllerManager HealthStatus `json:"userClusterControllerManager"`
}

// AllHealthy returns if all components are healthy
func (h *ExtendedClusterHealth) AllHealthy() bool {
	return h.Etcd == HealthStatusUp &&
		h.MachineController == HealthStatusUp &&
		h.Controller == HealthStatusUp &&
		h.Apiserver == HealthStatusUp &&
		h.Scheduler == HealthStatusUp &&
		h.CloudProviderInfrastructure == HealthStatusUp &&
		h.UserClusterControllerManager == HealthStatusUp
}

// MarshalJSON adds base64 json encoding to the Bytes type.
func (bs Bytes) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", base64.StdEncoding.EncodeToString(bs))), nil
}

// UnmarshalJSON adds base64 json decoding to the Bytes type.
func (bs *Bytes) UnmarshalJSON(src []byte) error {
	if len(src) < 2 {
		return errors.New("base64 string expected")
	}
	if src[0] != '"' || src[len(src)-1] != '"' {
		return errors.New("\" quotations expected")
	}
	if len(src) == 2 {
		*bs = nil
		return nil
	}
	var err error
	*bs, err = base64.StdEncoding.DecodeString(string(src[1 : len(src)-1]))
	return err
}

// Base64 converts a Bytes instance to a base64 string.
func (bs Bytes) Base64() string {
	if []byte(bs) == nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(bs))
}

// NewBytes creates a Bytes instance from a base64 string, returning nil for an empty base64 string.
func NewBytes(b64 string) Bytes {
	if b64 == "" {
		return Bytes(nil)
	}
	bs, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		panic(fmt.Sprintf("Invalid base64 string %q", b64))
	}
	return Bytes(bs)
}
