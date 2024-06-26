// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlschema

import (
	"fmt"
	"io"
	"strconv"
)

type ProviderSpecificConfig interface {
	IsProviderSpecificConfig()
}

type AWSProviderConfig struct {
	AwsZones     []*AWSZone `json:"awsZones"`
	VpcCidr      *string    `json:"vpcCidr,omitempty"`
	EnableIMDSv2 *bool      `json:"enableIMDSv2,omitempty"`
}

func (AWSProviderConfig) IsProviderSpecificConfig() {}

type AWSProviderConfigInput struct {
	VpcCidr      string          `json:"vpcCidr"`
	AwsZones     []*AWSZoneInput `json:"awsZones"`
	EnableIMDSv2 *bool           `json:"enableIMDSv2,omitempty"`
}

type AWSZone struct {
	Name         *string `json:"name,omitempty"`
	PublicCidr   *string `json:"publicCidr,omitempty"`
	InternalCidr *string `json:"internalCidr,omitempty"`
	WorkerCidr   *string `json:"workerCidr,omitempty"`
}

type AWSZoneInput struct {
	Name         string `json:"name"`
	PublicCidr   string `json:"publicCidr"`
	InternalCidr string `json:"internalCidr"`
	WorkerCidr   string `json:"workerCidr"`
}

type AzureProviderConfig struct {
	VnetCidr                     *string      `json:"vnetCidr,omitempty"`
	Zones                        []string     `json:"zones,omitempty"`
	AzureZones                   []*AzureZone `json:"azureZones,omitempty"`
	EnableNatGateway             *bool        `json:"enableNatGateway,omitempty"`
	IdleConnectionTimeoutMinutes *int         `json:"idleConnectionTimeoutMinutes,omitempty"`
}

func (AzureProviderConfig) IsProviderSpecificConfig() {}

type AzureProviderConfigInput struct {
	VnetCidr                     string            `json:"vnetCidr"`
	Zones                        []string          `json:"zones,omitempty"`
	AzureZones                   []*AzureZoneInput `json:"azureZones,omitempty"`
	EnableNatGateway             *bool             `json:"enableNatGateway,omitempty"`
	IdleConnectionTimeoutMinutes *int              `json:"idleConnectionTimeoutMinutes,omitempty"`
}

type AzureZone struct {
	Name int    `json:"name"`
	Cidr string `json:"cidr"`
}

type AzureZoneInput struct {
	Name int    `json:"name"`
	Cidr string `json:"cidr"`
}

type ClusterConfigInput struct {
	GardenerConfig *GardenerConfigInput `json:"gardenerConfig"`
	Administrators []string             `json:"administrators,omitempty"`
}

type ComponentConfiguration struct {
	Component     string         `json:"component"`
	Namespace     string         `json:"namespace"`
	Configuration []*ConfigEntry `json:"configuration,omitempty"`
	SourceURL     *string        `json:"sourceURL,omitempty"`
}

type ComponentConfigurationInput struct {
	Component        string              `json:"component"`
	Namespace        string              `json:"namespace"`
	Configuration    []*ConfigEntryInput `json:"configuration,omitempty"`
	SourceURL        *string             `json:"sourceURL,omitempty"`
	ConflictStrategy *ConflictStrategy   `json:"conflictStrategy,omitempty"`
}

type ConfigEntry struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Secret *bool  `json:"secret,omitempty"`
}

type ConfigEntryInput struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	Secret *bool  `json:"secret,omitempty"`
}

type DNSConfig struct {
	Domain    string         `json:"domain"`
	Providers []*DNSProvider `json:"providers,omitempty"`
}

type DNSConfigInput struct {
	Domain    string              `json:"domain"`
	Providers []*DNSProviderInput `json:"providers,omitempty"`
}

type DNSProvider struct {
	DomainsInclude []string `json:"domainsInclude"`
	Primary        bool     `json:"primary"`
	SecretName     string   `json:"secretName"`
	Type           string   `json:"type"`
}

type DNSProviderInput struct {
	DomainsInclude []string `json:"domainsInclude"`
	Primary        bool     `json:"primary"`
	SecretName     string   `json:"secretName"`
	Type           string   `json:"type"`
}

type Error struct {
	Message *string `json:"message,omitempty"`
}

type GCPProviderConfig struct {
	Zones []string `json:"zones"`
}

func (GCPProviderConfig) IsProviderSpecificConfig() {}

type GCPProviderConfigInput struct {
	Zones []string `json:"zones"`
}

type GardenerConfig struct {
	Name                                *string                `json:"name,omitempty"`
	KubernetesVersion                   *string                `json:"kubernetesVersion,omitempty"`
	TargetSecret                        *string                `json:"targetSecret,omitempty"`
	Provider                            *string                `json:"provider,omitempty"`
	Region                              *string                `json:"region,omitempty"`
	Seed                                *string                `json:"seed,omitempty"`
	MachineType                         *string                `json:"machineType,omitempty"`
	MachineImage                        *string                `json:"machineImage,omitempty"`
	MachineImageVersion                 *string                `json:"machineImageVersion,omitempty"`
	DiskType                            *string                `json:"diskType,omitempty"`
	VolumeSizeGb                        *int                   `json:"volumeSizeGB,omitempty"`
	WorkerCidr                          *string                `json:"workerCidr,omitempty"`
	PodsCidr                            *string                `json:"podsCidr,omitempty"`
	ServicesCidr                        *string                `json:"servicesCidr,omitempty"`
	AutoScalerMin                       *int                   `json:"autoScalerMin,omitempty"`
	AutoScalerMax                       *int                   `json:"autoScalerMax,omitempty"`
	MaxSurge                            *int                   `json:"maxSurge,omitempty"`
	MaxUnavailable                      *int                   `json:"maxUnavailable,omitempty"`
	Purpose                             *string                `json:"purpose,omitempty"`
	LicenceType                         *string                `json:"licenceType,omitempty"`
	EnableKubernetesVersionAutoUpdate   *bool                  `json:"enableKubernetesVersionAutoUpdate,omitempty"`
	EnableMachineImageVersionAutoUpdate *bool                  `json:"enableMachineImageVersionAutoUpdate,omitempty"`
	ProviderSpecificConfig              ProviderSpecificConfig `json:"providerSpecificConfig,omitempty"`
	DNSConfig                           *DNSConfig             `json:"dnsConfig,omitempty"`
	OidcConfig                          *OIDCConfig            `json:"oidcConfig,omitempty"`
	ExposureClassName                   *string                `json:"exposureClassName,omitempty"`
	ShootNetworkingFilterDisabled       *bool                  `json:"shootNetworkingFilterDisabled,omitempty"`
	ControlPlaneFailureTolerance        *string                `json:"controlPlaneFailureTolerance,omitempty"`
	EuAccess                            *bool                  `json:"euAccess,omitempty"`
}

type GardenerConfigInput struct {
	Name                                string                 `json:"name"`
	KubernetesVersion                   string                 `json:"kubernetesVersion"`
	Provider                            string                 `json:"provider"`
	TargetSecret                        string                 `json:"targetSecret"`
	Region                              string                 `json:"region"`
	MachineType                         string                 `json:"machineType"`
	MachineImage                        *string                `json:"machineImage,omitempty"`
	MachineImageVersion                 *string                `json:"machineImageVersion,omitempty"`
	DiskType                            *string                `json:"diskType,omitempty"`
	VolumeSizeGb                        *int                   `json:"volumeSizeGB,omitempty"`
	WorkerCidr                          string                 `json:"workerCidr"`
	PodsCidr                            *string                `json:"podsCidr,omitempty"`
	ServicesCidr                        *string                `json:"servicesCidr,omitempty"`
	AutoScalerMin                       int                    `json:"autoScalerMin"`
	AutoScalerMax                       int                    `json:"autoScalerMax"`
	MaxSurge                            int                    `json:"maxSurge"`
	MaxUnavailable                      int                    `json:"maxUnavailable"`
	Purpose                             *string                `json:"purpose,omitempty"`
	LicenceType                         *string                `json:"licenceType,omitempty"`
	EnableKubernetesVersionAutoUpdate   *bool                  `json:"enableKubernetesVersionAutoUpdate,omitempty"`
	EnableMachineImageVersionAutoUpdate *bool                  `json:"enableMachineImageVersionAutoUpdate,omitempty"`
	ProviderSpecificConfig              *ProviderSpecificInput `json:"providerSpecificConfig"`
	DNSConfig                           *DNSConfigInput        `json:"dnsConfig,omitempty"`
	Seed                                *string                `json:"seed,omitempty"`
	OidcConfig                          *OIDCConfigInput       `json:"oidcConfig,omitempty"`
	ExposureClassName                   *string                `json:"exposureClassName,omitempty"`
	ShootNetworkingFilterDisabled       *bool                  `json:"shootNetworkingFilterDisabled,omitempty"`
	ControlPlaneFailureTolerance        *string                `json:"controlPlaneFailureTolerance,omitempty"`
	EuAccess                            *bool                  `json:"euAccess,omitempty"`
}

type GardenerUpgradeInput struct {
	KubernetesVersion                   *string                `json:"kubernetesVersion,omitempty"`
	MachineType                         *string                `json:"machineType,omitempty"`
	DiskType                            *string                `json:"diskType,omitempty"`
	VolumeSizeGb                        *int                   `json:"volumeSizeGB,omitempty"`
	AutoScalerMin                       *int                   `json:"autoScalerMin,omitempty"`
	AutoScalerMax                       *int                   `json:"autoScalerMax,omitempty"`
	MachineImage                        *string                `json:"machineImage,omitempty"`
	MachineImageVersion                 *string                `json:"machineImageVersion,omitempty"`
	MaxSurge                            *int                   `json:"maxSurge,omitempty"`
	MaxUnavailable                      *int                   `json:"maxUnavailable,omitempty"`
	Purpose                             *string                `json:"purpose,omitempty"`
	EnableKubernetesVersionAutoUpdate   *bool                  `json:"enableKubernetesVersionAutoUpdate,omitempty"`
	EnableMachineImageVersionAutoUpdate *bool                  `json:"enableMachineImageVersionAutoUpdate,omitempty"`
	ProviderSpecificConfig              *ProviderSpecificInput `json:"providerSpecificConfig,omitempty"`
	OidcConfig                          *OIDCConfigInput       `json:"oidcConfig,omitempty"`
	ExposureClassName                   *string                `json:"exposureClassName,omitempty"`
	ShootNetworkingFilterDisabled       *bool                  `json:"shootNetworkingFilterDisabled,omitempty"`
}

type HibernationStatus struct {
	Hibernated          *bool `json:"hibernated,omitempty"`
	HibernationPossible *bool `json:"hibernationPossible,omitempty"`
}

type KymaConfig struct {
	Version       *string                   `json:"version,omitempty"`
	Profile       *KymaProfile              `json:"profile,omitempty"`
	Components    []*ComponentConfiguration `json:"components,omitempty"`
	Configuration []*ConfigEntry            `json:"configuration,omitempty"`
}

type KymaConfigInput struct {
	Version          string                         `json:"version"`
	Profile          *KymaProfile                   `json:"profile,omitempty"`
	Components       []*ComponentConfigurationInput `json:"components"`
	Configuration    []*ConfigEntryInput            `json:"configuration,omitempty"`
	ConflictStrategy *ConflictStrategy              `json:"conflictStrategy,omitempty"`
}

type LastError struct {
	ErrMessage string `json:"errMessage"`
	Reason     string `json:"reason"`
	Component  string `json:"component"`
}

type Mutation struct {
}

type OIDCConfig struct {
	ClientID       string   `json:"clientID"`
	GroupsClaim    string   `json:"groupsClaim"`
	IssuerURL      string   `json:"issuerURL"`
	SigningAlgs    []string `json:"signingAlgs"`
	UsernameClaim  string   `json:"usernameClaim"`
	UsernamePrefix string   `json:"usernamePrefix"`
}

type OIDCConfigInput struct {
	ClientID       string   `json:"clientID"`
	GroupsClaim    string   `json:"groupsClaim"`
	IssuerURL      string   `json:"issuerURL"`
	SigningAlgs    []string `json:"signingAlgs"`
	UsernameClaim  string   `json:"usernameClaim"`
	UsernamePrefix string   `json:"usernamePrefix"`
}

type OpenStackProviderConfig struct {
	Zones                []string `json:"zones"`
	FloatingPoolName     string   `json:"floatingPoolName"`
	CloudProfileName     string   `json:"cloudProfileName"`
	LoadBalancerProvider string   `json:"loadBalancerProvider"`
}

func (OpenStackProviderConfig) IsProviderSpecificConfig() {}

type OpenStackProviderConfigInput struct {
	Zones                []string `json:"zones"`
	FloatingPoolName     *string  `json:"floatingPoolName,omitempty"`
	CloudProfileName     *string  `json:"cloudProfileName,omitempty"`
	LoadBalancerProvider string   `json:"loadBalancerProvider"`
}

type OperationStatus struct {
	ID               *string        `json:"id,omitempty"`
	Operation        OperationType  `json:"operation"`
	State            OperationState `json:"state"`
	Message          *string        `json:"message,omitempty"`
	RuntimeID        *string        `json:"runtimeID,omitempty"`
	CompassRuntimeID *string        `json:"compassRuntimeID,omitempty"`
	LastError        *LastError     `json:"lastError,omitempty"`
}

type ProviderSpecificInput struct {
	GcpConfig       *GCPProviderConfigInput       `json:"gcpConfig,omitempty"`
	AzureConfig     *AzureProviderConfigInput     `json:"azureConfig,omitempty"`
	AwsConfig       *AWSProviderConfigInput       `json:"awsConfig,omitempty"`
	OpenStackConfig *OpenStackProviderConfigInput `json:"openStackConfig,omitempty"`
}

type ProvisionRuntimeInput struct {
	RuntimeInput  *RuntimeInput       `json:"runtimeInput"`
	ClusterConfig *ClusterConfigInput `json:"clusterConfig"`
	KymaConfig    *KymaConfigInput    `json:"kymaConfig,omitempty"`
}

type Query struct {
}

type RuntimeConfig struct {
	ClusterConfig *GardenerConfig `json:"clusterConfig,omitempty"`
	KymaConfig    *KymaConfig     `json:"kymaConfig,omitempty"`
	Kubeconfig    *string         `json:"kubeconfig,omitempty"`
}

type RuntimeConnectionStatus struct {
	Status RuntimeAgentConnectionStatus `json:"status"`
	Errors []*Error                     `json:"errors,omitempty"`
}

type RuntimeInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Labels      Labels  `json:"labels,omitempty"`
}

type RuntimeStatus struct {
	LastOperationStatus     *OperationStatus         `json:"lastOperationStatus,omitempty"`
	RuntimeConnectionStatus *RuntimeConnectionStatus `json:"runtimeConnectionStatus,omitempty"`
	RuntimeConfiguration    *RuntimeConfig           `json:"runtimeConfiguration,omitempty"`
	HibernationStatus       *HibernationStatus       `json:"hibernationStatus,omitempty"`
}

type UpgradeRuntimeInput struct {
	KymaConfig *KymaConfigInput `json:"kymaConfig"`
}

type UpgradeShootInput struct {
	GardenerConfig *GardenerUpgradeInput `json:"gardenerConfig"`
	Administrators []string              `json:"administrators,omitempty"`
}

type ConflictStrategy string

const (
	ConflictStrategyMerge   ConflictStrategy = "Merge"
	ConflictStrategyReplace ConflictStrategy = "Replace"
)

var AllConflictStrategy = []ConflictStrategy{
	ConflictStrategyMerge,
	ConflictStrategyReplace,
}

func (e ConflictStrategy) IsValid() bool {
	switch e {
	case ConflictStrategyMerge, ConflictStrategyReplace:
		return true
	}
	return false
}

func (e ConflictStrategy) String() string {
	return string(e)
}

func (e *ConflictStrategy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ConflictStrategy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ConflictStrategy", str)
	}
	return nil
}

func (e ConflictStrategy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type KymaProfile string

const (
	KymaProfileEvaluation KymaProfile = "Evaluation"
	KymaProfileProduction KymaProfile = "Production"
)

var AllKymaProfile = []KymaProfile{
	KymaProfileEvaluation,
	KymaProfileProduction,
}

func (e KymaProfile) IsValid() bool {
	switch e {
	case KymaProfileEvaluation, KymaProfileProduction:
		return true
	}
	return false
}

func (e KymaProfile) String() string {
	return string(e)
}

func (e *KymaProfile) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = KymaProfile(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid KymaProfile", str)
	}
	return nil
}

func (e KymaProfile) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OperationState string

const (
	OperationStatePending    OperationState = "Pending"
	OperationStateInProgress OperationState = "InProgress"
	OperationStateSucceeded  OperationState = "Succeeded"
	OperationStateFailed     OperationState = "Failed"
)

var AllOperationState = []OperationState{
	OperationStatePending,
	OperationStateInProgress,
	OperationStateSucceeded,
	OperationStateFailed,
}

func (e OperationState) IsValid() bool {
	switch e {
	case OperationStatePending, OperationStateInProgress, OperationStateSucceeded, OperationStateFailed:
		return true
	}
	return false
}

func (e OperationState) String() string {
	return string(e)
}

func (e *OperationState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationState", str)
	}
	return nil
}

func (e OperationState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OperationType string

const (
	OperationTypeProvision            OperationType = "Provision"
	OperationTypeProvisionNoInstall   OperationType = "ProvisionNoInstall"
	OperationTypeUpgrade              OperationType = "Upgrade"
	OperationTypeUpgradeShoot         OperationType = "UpgradeShoot"
	OperationTypeDeprovision          OperationType = "Deprovision"
	OperationTypeDeprovisionNoInstall OperationType = "DeprovisionNoInstall"
	OperationTypeReconnectRuntime     OperationType = "ReconnectRuntime"
	OperationTypeHibernate            OperationType = "Hibernate"
)

var AllOperationType = []OperationType{
	OperationTypeProvision,
	OperationTypeProvisionNoInstall,
	OperationTypeUpgrade,
	OperationTypeUpgradeShoot,
	OperationTypeDeprovision,
	OperationTypeDeprovisionNoInstall,
	OperationTypeReconnectRuntime,
	OperationTypeHibernate,
}

func (e OperationType) IsValid() bool {
	switch e {
	case OperationTypeProvision, OperationTypeProvisionNoInstall, OperationTypeUpgrade, OperationTypeUpgradeShoot, OperationTypeDeprovision, OperationTypeDeprovisionNoInstall, OperationTypeReconnectRuntime, OperationTypeHibernate:
		return true
	}
	return false
}

func (e OperationType) String() string {
	return string(e)
}

func (e *OperationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OperationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OperationType", str)
	}
	return nil
}

func (e OperationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RuntimeAgentConnectionStatus string

const (
	RuntimeAgentConnectionStatusPending      RuntimeAgentConnectionStatus = "Pending"
	RuntimeAgentConnectionStatusConnected    RuntimeAgentConnectionStatus = "Connected"
	RuntimeAgentConnectionStatusDisconnected RuntimeAgentConnectionStatus = "Disconnected"
)

var AllRuntimeAgentConnectionStatus = []RuntimeAgentConnectionStatus{
	RuntimeAgentConnectionStatusPending,
	RuntimeAgentConnectionStatusConnected,
	RuntimeAgentConnectionStatusDisconnected,
}

func (e RuntimeAgentConnectionStatus) IsValid() bool {
	switch e {
	case RuntimeAgentConnectionStatusPending, RuntimeAgentConnectionStatusConnected, RuntimeAgentConnectionStatusDisconnected:
		return true
	}
	return false
}

func (e RuntimeAgentConnectionStatus) String() string {
	return string(e)
}

func (e *RuntimeAgentConnectionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RuntimeAgentConnectionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RuntimeAgentConnectionStatus", str)
	}
	return nil
}

func (e RuntimeAgentConnectionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
