// Package genesiscloud provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.14.0 DO NOT EDIT.
package genesiscloud

import (
	"time"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for FloatingIPVersion.
const (
	FloatingIPVersionIpv4 FloatingIPVersion = "ipv4"
	FloatingIPVersionIpv6 FloatingIPVersion = "ipv6"
)

var AllFloatingIPVersions = []FloatingIPVersion{
	FloatingIPVersionIpv4,
	FloatingIPVersionIpv6,
}

// Defines values for ImageType.
const (
	ImageTypeBaseOs        ImageType = "base-os"
	ImageTypePreconfigured ImageType = "preconfigured"
	ImageTypeSnapshot      ImageType = "snapshot"
)

var AllImageTypes = []ImageType{
	ImageTypeBaseOs,
	ImageTypePreconfigured,
	ImageTypeSnapshot,
}

// Defines values for InstanceAction.
const (
	InstanceActionReset InstanceAction = "reset"
	InstanceActionStart InstanceAction = "start"
	InstanceActionStop  InstanceAction = "stop"
)

var AllInstanceActions = []InstanceAction{
	InstanceActionReset,
	InstanceActionStart,
	InstanceActionStop,
}

// Defines values for InstancePublicIPType.
const (
	InstancePublicIPTypeDynamic InstancePublicIPType = "dynamic"
	InstancePublicIPTypeStatic  InstancePublicIPType = "static"
)

var AllInstancePublicIPTypes = []InstancePublicIPType{
	InstancePublicIPTypeDynamic,
	InstancePublicIPTypeStatic,
}

// Defines values for InstanceStatus.
const (
	InstanceStatusActive         InstanceStatus = "active"
	InstanceStatusBuild          InstanceStatus = "build"
	InstanceStatusCopying        InstanceStatus = "copying"
	InstanceStatusCreating       InstanceStatus = "creating"
	InstanceStatusDeleting       InstanceStatus = "deleting"
	InstanceStatusEnqueued       InstanceStatus = "enqueued"
	InstanceStatusError          InstanceStatus = "error"
	InstanceStatusPendingPayment InstanceStatus = "pending_payment"
	InstanceStatusResetting      InstanceStatus = "resetting"
	InstanceStatusRestarting     InstanceStatus = "restarting"
	InstanceStatusShutoff        InstanceStatus = "shutoff"
	InstanceStatusStarting       InstanceStatus = "starting"
	InstanceStatusStopped        InstanceStatus = "stopped"
	InstanceStatusStopping       InstanceStatus = "stopping"
	InstanceStatusUnknown        InstanceStatus = "unknown"
	InstanceStatusUpdating       InstanceStatus = "updating"
)

var AllInstanceStatuss = []InstanceStatus{
	InstanceStatusActive,
	InstanceStatusBuild,
	InstanceStatusCopying,
	InstanceStatusCreating,
	InstanceStatusDeleting,
	InstanceStatusEnqueued,
	InstanceStatusError,
	InstanceStatusPendingPayment,
	InstanceStatusResetting,
	InstanceStatusRestarting,
	InstanceStatusShutoff,
	InstanceStatusStarting,
	InstanceStatusStopped,
	InstanceStatusStopping,
	InstanceStatusUnknown,
	InstanceStatusUpdating,
}

// Defines values for Region.
const (
	RegionARCISHAF1  Region = "ARC-IS-HAF-1"
	RegionEUCDEMUC1  Region = "EUC-DE-MUC-1"
	RegionNORDNOKRS1 Region = "NORD-NO-KRS-1"
)

var AllRegions = []Region{
	RegionARCISHAF1,
	RegionEUCDEMUC1,
	RegionNORDNOKRS1,
}

// Defines values for SecurityGroupRuleDirection.
const (
	SecurityGroupRuleDirectionEgress  SecurityGroupRuleDirection = "egress"
	SecurityGroupRuleDirectionIngress SecurityGroupRuleDirection = "ingress"
)

var AllSecurityGroupRuleDirections = []SecurityGroupRuleDirection{
	SecurityGroupRuleDirectionEgress,
	SecurityGroupRuleDirectionIngress,
}

// Defines values for SecurityGroupRuleProtocol.
const (
	SecurityGroupRuleProtocolAll  SecurityGroupRuleProtocol = "all"
	SecurityGroupRuleProtocolIcmp SecurityGroupRuleProtocol = "icmp"
	SecurityGroupRuleProtocolTcp  SecurityGroupRuleProtocol = "tcp"
	SecurityGroupRuleProtocolUdp  SecurityGroupRuleProtocol = "udp"
)

var AllSecurityGroupRuleProtocols = []SecurityGroupRuleProtocol{
	SecurityGroupRuleProtocolAll,
	SecurityGroupRuleProtocolIcmp,
	SecurityGroupRuleProtocolTcp,
	SecurityGroupRuleProtocolUdp,
}

// Defines values for SecurityGroupStatus.
const (
	SecurityGroupStatusCreated  SecurityGroupStatus = "created"
	SecurityGroupStatusCreating SecurityGroupStatus = "creating"
	SecurityGroupStatusDeleting SecurityGroupStatus = "deleting"
	SecurityGroupStatusError    SecurityGroupStatus = "error"
	SecurityGroupStatusUpdating SecurityGroupStatus = "updating"
)

var AllSecurityGroupStatuss = []SecurityGroupStatus{
	SecurityGroupStatusCreated,
	SecurityGroupStatusCreating,
	SecurityGroupStatusDeleting,
	SecurityGroupStatusError,
	SecurityGroupStatusUpdating,
}

// Defines values for SnapshotStatus.
const (
	SnapshotStatusActive        SnapshotStatus = "active"
	SnapshotStatusCreated       SnapshotStatus = "created"
	SnapshotStatusCreating      SnapshotStatus = "creating"
	SnapshotStatusDeleting      SnapshotStatus = "deleting"
	SnapshotStatusError         SnapshotStatus = "error"
	SnapshotStatusPendingDelete SnapshotStatus = "pending_delete"
)

var AllSnapshotStatuss = []SnapshotStatus{
	SnapshotStatusActive,
	SnapshotStatusCreated,
	SnapshotStatusCreating,
	SnapshotStatusDeleting,
	SnapshotStatusError,
	SnapshotStatusPendingDelete,
}

// Defines values for VolumeStatus.
const (
	VolumeStatusAvailable     VolumeStatus = "available"
	VolumeStatusCreated       VolumeStatus = "created"
	VolumeStatusCreating      VolumeStatus = "creating"
	VolumeStatusDeleting      VolumeStatus = "deleting"
	VolumeStatusError         VolumeStatus = "error"
	VolumeStatusInUse         VolumeStatus = "in-use"
	VolumeStatusPendingDelete VolumeStatus = "pending_delete"
)

var AllVolumeStatuss = []VolumeStatus{
	VolumeStatusAvailable,
	VolumeStatusCreated,
	VolumeStatusCreating,
	VolumeStatusDeleting,
	VolumeStatusError,
	VolumeStatusInUse,
	VolumeStatusPendingDelete,
}

// Defines values for VolumeType.
const (
	VolumeTypeHdd VolumeType = "hdd"
	VolumeTypeSsd VolumeType = "ssd"
)

var AllVolumeTypes = []VolumeType{
	VolumeTypeHdd,
	VolumeTypeSsd,
}

// Defines values for CreateFloatingIPJSONBodyVersion.
const (
	CreateFloatingIPJSONBodyVersionIpv4 CreateFloatingIPJSONBodyVersion = "ipv4"
	CreateFloatingIPJSONBodyVersionIpv6 CreateFloatingIPJSONBodyVersion = "ipv6"
)

var AllCreateFloatingIPJSONBodyVersions = []CreateFloatingIPJSONBodyVersion{
	CreateFloatingIPJSONBodyVersionIpv4,
	CreateFloatingIPJSONBodyVersionIpv6,
}

// Error defines model for Error.
type Error struct {
	// Code The Genesis Cloud error code.
	// Check the [developer documentation](https://developers.com/#error-codes) for more information on error codes.
	Code string `json:"code"`

	// Message An explanation of what went wrong.
	Message string `json:"message"`
}

// FloatingIP defines model for FloatingIP.
type FloatingIP struct {
	CreatedAt Timestamp `json:"created_at"`

	// Description The human-readable description for the floating IP.
	Description string `json:"description"`

	// Id A unique identifier for each floating IP. This is automatically generated.
	Id string `json:"id"`

	// IpAddress The IP address of the floating IP.
	IpAddress *string `json:"ip_address,omitempty"`

	// IsPublic A boolean value indicating whether the floating IP is public or private.
	IsPublic bool `json:"is_public"`

	// Name The human-readable name for the floating IP.
	Name string `json:"name"`

	// Region The region identifier.
	Region    Region    `json:"region"`
	UpdatedAt Timestamp `json:"updated_at"`

	// Version The IP version of the floating IP.
	Version *FloatingIPVersion `json:"version,omitempty"`
}

// FloatingIPVersion The IP version of the floating IP.
type FloatingIPVersion string

// Image defines model for Image.
type Image struct {
	CreatedAt Timestamp `json:"created_at"`

	// Id A unique number that can be used to identify and reference a specific image.
	Id ImageId `json:"id"`

	// Name The display name that has been given to an image. This is what is shown in the control panel and is generally a descriptive title for the image in question.
	Name string `json:"name"`

	// Regions The list of regions in which this image can be used in.
	Regions []Region `json:"regions"`

	// Type Describes the kind of image.
	Type ImageType `json:"type"`
}

// ImageType Describes the kind of image.
type ImageType string

// ImageId A unique number that can be used to identify and reference a specific image.
type ImageId = string

// Instance defines model for Instance.
type Instance struct {
	CreatedAt Timestamp `json:"created_at"`

	// Hostname The hostname of your instance.
	Hostname InstanceHostname `json:"hostname"`

	// Id The unique ID of the instance.
	Id string `json:"id"`

	// Image The image of the instance.
	Image struct {
		// Id A unique number that can be used to identify and reference a specific image.
		Id ImageId `json:"id"`

		// Name The image name.
		Name string `json:"name"`
	} `json:"image"`

	// Name The human-readable name set for the instance.
	Name InstanceName `json:"name"`

	// PlacementOption The placement option identifier in which instances are physically located relative to each other within a zone.
	PlacementOption string `json:"placement_option"`

	// PrivateIp The private IPv4 IP-Address (IPv4 address).
	PrivateIp *string `json:"private_ip"`

	// PublicIp The public IPv4 IP-Address (IPv4 address).
	PublicIp *string `json:"public_ip"`

	// PublicIpType When set to `static`, the instance's public IP will not change between start and stop actions.
	PublicIpType InstancePublicIPType `json:"public_ip_type"`

	// Region The region identifier.
	Region Region `json:"region"`

	// SecurityGroups The security groups of the instance.
	SecurityGroups []struct {
		// Id The security group ID.
		Id InstanceSecurityGroupId `json:"id"`

		// Name The name of the security group.
		Name string `json:"name"`
	} `json:"security_groups"`

	// SshKeys The ssh keys of the instance.
	SshKeys []struct {
		// Id The ssh key ID.
		Id InstanceSSHKeyId `json:"id"`

		// Name The name of the ssh key.
		Name string `json:"name"`
	} `json:"ssh_keys"`

	// Status The instance status
	Status InstanceStatus `json:"status"`

	// Type The instance type identifier.
	Type      InstanceType `json:"type"`
	UpdatedAt Timestamp    `json:"updated_at"`

	// Volumes The volumes of the instance
	Volumes []struct {
		// Id A unique identifier for each volume. This is automatically generated.
		Id VolumeId `json:"id"`

		// Name The volume name.
		Name string `json:"name"`
	} `json:"volumes"`
}

// InstanceAction defines model for Instance.Action.
type InstanceAction string

// InstanceHostname The hostname of your instance.
type InstanceHostname = string

// InstanceIsProtected Specifies if the instance is termination protected.
// When set to `true`, it"s not possible to destroy the instance until it"s switched to `false`.
// Set to `true` automatically for long-term billed instances.
type InstanceIsProtected = bool

// InstanceName The human-readable name set for the instance.
type InstanceName = string

// InstancePublicIPType When set to `static`, the instance's public IP will not change between start and stop actions.
type InstancePublicIPType string

// InstanceSSHKeyId The ssh key ID.
type InstanceSSHKeyId = string

// InstanceSecurityGroupId The security group ID.
type InstanceSecurityGroupId = string

// InstanceSecurityGroupIds An array of security group ids.
// **Please Note**: By default the **standard security group** is set if you don"t specify any Security Groups.
// You can override this behavior by providing a different Security Group.
type InstanceSecurityGroupIds = []InstanceSecurityGroupId

// InstanceStatus The instance status
type InstanceStatus string

// InstanceType The instance type identifier.
type InstanceType = string

// InstanceUserData defines model for Instance.UserData.
type InstanceUserData = []struct {
	Content  string  `json:"content"`
	Filename *string `json:"filename,omitempty"`
	Type     string  `json:"type"`
}

// InstancesAvailability defines model for InstancesAvailability.
type InstancesAvailability struct {
	Availability struct {
		InstanceTypes []struct {
			Available bool     `json:"available"`
			Count     *float32 `json:"count,omitempty"`

			// Type The instance type identifier.
			Type InstanceType `json:"type"`
		} `json:"instance_types"`
		Placement *string `json:"placement,omitempty"`

		// Region The region identifier.
		Region Region `json:"region"`
		Type   string `json:"type"`
	} `json:"availability"`
}

// Region The region identifier.
type Region string

// SSHKey defines model for SSHKey.
type SSHKey struct {
	CreatedAt Timestamp `json:"created_at"`

	// Fingerprint The fingerprint of the SSH key.
	Fingerprint string `json:"fingerprint"`

	// Id A unique identifier for each SSH key. This is automatically generated.
	Id string `json:"id"`

	// Name The human-readable name for the SSH key.
	Name string `json:"name"`

	// Size The length of the SSH key.
	Size int `json:"size"`

	// Type The encryption algorithm type of the SSH key.
	Type      string    `json:"type"`
	UpdatedAt Timestamp `json:"updated_at"`

	// Value SSH public key.
	Value string `json:"value"`
}

// SecurityGroup defines model for SecurityGroup.
type SecurityGroup struct {
	CreatedAt Timestamp `json:"created_at"`

	// Description The human-readable description for the security group.
	Description string `json:"description"`

	// Id A unique identifier for each security group. This is automatically generated.
	Id         string `json:"id"`
	IsInternal bool   `json:"is_internal"`

	// Name The human-readable name for the security group.
	Name string `json:"name"`

	// Region The region identifier.
	Region Region              `json:"region"`
	Rules  []SecurityGroupRule `json:"rules"`

	// Status The security group status.
	Status    SecurityGroupStatus `json:"status"`
	UpdatedAt Timestamp           `json:"updated_at"`
}

// SecurityGroupRule defines model for SecurityGroup.Rule.
type SecurityGroupRule struct {
	// Direction The direction of the rule.
	Direction SecurityGroupRuleDirection `json:"direction"`

	// PortRangeMax The maximum port number of the rule.
	PortRangeMax *int `json:"port_range_max"`

	// PortRangeMin The minimum port number of the rule.
	PortRangeMin *int `json:"port_range_min"`

	// Protocol The protocol of the rule.
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
}

// SecurityGroupRuleDirection The direction of the rule.
type SecurityGroupRuleDirection string

// SecurityGroupRuleProtocol The protocol of the rule.
type SecurityGroupRuleProtocol string

// SecurityGroupStatus The security group status.
type SecurityGroupStatus string

// Snapshot defines model for Snapshot.
type Snapshot struct {
	CreatedAt Timestamp `json:"created_at"`

	// Id A unique identifier for each snapshot. This is automatically generated.
	Id string `json:"id"`

	// Name The human-readable name for the snapshot.
	Name string `json:"name"`

	// Region The region identifier.
	Region Region `json:"region"`

	// ResourceId The id of the resource (e.g. instance) that was snapshotted.
	ResourceId string `json:"resource_id"`

	// Size The storage size of this snapshot given in bytes.
	Size *string `json:"size,omitempty"`

	// Status The snapshot status.
	Status    SnapshotStatus `json:"status"`
	UpdatedAt Timestamp      `json:"updated_at"`
}

// SnapshotStatus The snapshot status.
type SnapshotStatus string

// Timestamp defines model for Timestamp.
type Timestamp = time.Time

// Volume defines model for Volume.
type Volume struct {
	CreatedAt Timestamp `json:"created_at"`

	// Description The human-readable description for the volume.
	Description string `json:"description"`

	// Id A unique identifier for each volume. This is automatically generated.
	Id VolumeId `json:"id"`

	// Instances The attached instances.
	Instances []struct {
		// Id The id of the attached instance.
		Id *string `json:"id,omitempty"`

		// Name The name of the attached instance.
		Name *string `json:"name,omitempty"`
	} `json:"instances"`

	// Name The human-readable name for the volume.
	Name string `json:"name"`

	// Region The region identifier.
	Region Region `json:"region"`

	// Size The storage size of this volume given in GiB.
	Size   int          `json:"size"`
	Status VolumeStatus `json:"status"`

	// Type The volume type.
	Type      VolumeType `json:"type"`
	UpdatedAt Timestamp  `json:"updated_at"`
}

// VolumeStatus defines model for Volume.Status.
type VolumeStatus string

// VolumeId A unique identifier for each volume. This is automatically generated.
type VolumeId = string

// VolumeType The volume type.
type VolumeType string

// PageQueryParameter A positive integer to choose the page to return.
type PageQueryParameter = int

// PerPageQueryParameter A positive integer lower or equal to 100 to select the number of items to return.
type PerPageQueryParameter = int

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse = Error

// InstanceListActionsResponse defines model for Instance.ListActionsResponse.
type InstanceListActionsResponse struct {
	Actions []InstanceAction `json:"actions"`
}

// InstancesAvailabilityResponse defines model for InstancesAvailabilityResponse.
type InstancesAvailabilityResponse = InstancesAvailability

// PaginatedFloatingIPsResponse defines model for PaginatedFloatingIPsResponse.
type PaginatedFloatingIPsResponse struct {
	FloatingIps []FloatingIP `json:"floating_ips"`
	Page        int          `json:"page"`
	PerPage     int          `json:"per_page"`
	TotalCount  int          `json:"total_count"`
}

// PaginatedImagesResponse defines model for PaginatedImagesResponse.
type PaginatedImagesResponse struct {
	Images     []Image `json:"images"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	TotalCount int     `json:"total_count"`
}

// PaginatedInstancesResponse defines model for PaginatedInstancesResponse.
type PaginatedInstancesResponse struct {
	Instances  []Instance `json:"instances"`
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	TotalCount int        `json:"total_count"`
}

// PaginatedSSHKeysResponse defines model for PaginatedSSHKeysResponse.
type PaginatedSSHKeysResponse struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	SshKeys    []SSHKey `json:"ssh_keys"`
	TotalCount int      `json:"total_count"`
}

// PaginatedSecurityGroupsResponse defines model for PaginatedSecurityGroupsResponse.
type PaginatedSecurityGroupsResponse struct {
	Page           int             `json:"page"`
	PerPage        int             `json:"per_page"`
	SecurityGroups []SecurityGroup `json:"security_groups"`
	TotalCount     int             `json:"total_count"`
}

// PaginatedSnapshotsResponse defines model for PaginatedSnapshotsResponse.
type PaginatedSnapshotsResponse struct {
	Page       int        `json:"page"`
	PerPage    int        `json:"per_page"`
	Snapshots  []Snapshot `json:"snapshots"`
	TotalCount int        `json:"total_count"`
}

// PaginatedVolumesResponse defines model for PaginatedVolumesResponse.
type PaginatedVolumesResponse struct {
	Page       int      `json:"page"`
	PerPage    int      `json:"per_page"`
	TotalCount int      `json:"total_count"`
	Volumes    []Volume `json:"volumes"`
}

// SingleFloatingIPResponse defines model for SingleFloatingIPResponse.
type SingleFloatingIPResponse struct {
	FloatingIp *FloatingIP `json:"floating_ip,omitempty"`
}

// SingleInstanceResponse defines model for SingleInstanceResponse.
type SingleInstanceResponse struct {
	Instance Instance `json:"instance"`
}

// SingleSSHKeyResponse defines model for SingleSSHKeyResponse.
type SingleSSHKeyResponse = SSHKey

// SingleSecurityGroupResponse defines model for SingleSecurityGroupResponse.
type SingleSecurityGroupResponse struct {
	SecurityGroup SecurityGroup `json:"security_group"`
}

// SingleSnapshotResponse defines model for SingleSnapshotResponse.
type SingleSnapshotResponse struct {
	Snapshot Snapshot `json:"snapshot"`
}

// SingleVolumeResponse defines model for SingleVolumeResponse.
type SingleVolumeResponse struct {
	Volume Volume `json:"volume"`
}

// GetInstancesAvailabilityParams defines parameters for GetInstancesAvailability.
type GetInstancesAvailabilityParams struct {
	Placement *string `form:"placement,omitempty" json:"placement,omitempty"`
}

// ListFloatingIPsParams defines parameters for ListFloatingIPs.
type ListFloatingIPsParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreateFloatingIPJSONBody defines parameters for CreateFloatingIP.
type CreateFloatingIPJSONBody struct {
	// Description The human-readable description set for the volume.
	Description *string `json:"description,omitempty"`

	// IsPublic Whether the floating IP is public or private.
	IsPublic bool `json:"is_public"`

	// Name The human-readable name set for the volume.
	Name string `json:"name"`

	// Region The region identifier.
	Region Region `json:"region"`

	// Version The IP version of the floating IP.
	Version CreateFloatingIPJSONBodyVersion `json:"version"`
}

// CreateFloatingIPJSONBodyVersion defines parameters for CreateFloatingIP.
type CreateFloatingIPJSONBodyVersion string

// ListImagesParams defines parameters for ListImages.
type ListImagesParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
	Type    *ImageType             `form:"type,omitempty" json:"type,omitempty"`
}

// ListInstancesParams defines parameters for ListInstances.
type ListInstancesParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreateInstanceJSONBody defines parameters for CreateInstance.
type CreateInstanceJSONBody struct {
	// Hostname The hostname of your instance.
	Hostname InstanceHostname `json:"hostname"`

	// Image A unique number that can be used to identify and reference a specific image.
	Image ImageId `json:"image"`

	// IsProtected Specifies if the instance is termination protected.
	// When set to `true`, it"s not possible to destroy the instance until it"s switched to `false`.
	// Set to `true` automatically for long-term billed instances.
	IsProtected *InstanceIsProtected `json:"is_protected,omitempty"`

	// Metadata Option to provide metadata.
	Metadata *struct {
		// StartupScript A plain text bash script or "cloud-config" file that will be executed after the first instance boot.
		// It is limited to 64 KiB in size. You can use it to configure your instance, e.g. installing the **NVIDIA GPU driver**.
		// Learn more about [startup scripts and installing the GPU driver](https://support.com/support/solutions/articles/47001122478).
		StartupScript *string           `json:"startup_script,omitempty"`
		UserData      *InstanceUserData `json:"user_data,omitempty"`
	} `json:"metadata,omitempty"`

	// Name The human-readable name set for the instance.
	Name InstanceName `json:"name"`

	// Password The password to access the instance.
	// Your password must have upper and lower chars, digits and length between 8-72.
	// **Please Note**: Only one of `ssh_keys` or `password` can be provided.
	// Password is less secure - we recommend you use an SSH key-pair.
	Password *string `json:"password,omitempty"`

	// PlacementOption The placement option identifier in which instances are physically located relative to each other within a zone.
	PlacementOption *string `json:"placement_option,omitempty"`

	// PublicIpType When set to `static`, the instance's public IP will not change between start and stop actions.
	PublicIpType *InstancePublicIPType `json:"public_ip_type,omitempty"`

	// Region The region identifier.
	Region Region `json:"region"`

	// SecurityGroups An array of security group ids.
	// **Please Note**: By default the **standard security group** is set if you don"t specify any Security Groups.
	// You can override this behavior by providing a different Security Group.
	SecurityGroups *InstanceSecurityGroupIds `json:"security_groups,omitempty"`

	// SshKeys An array of ssh key ids.
	// This should not be provided if `password` authentication method is desired.
	SshKeys *[]InstanceSSHKeyId `json:"ssh_keys,omitempty"`

	// Type The instance type identifier.
	Type InstanceType `json:"type"`

	// Volumes An array of volume ids.
	Volumes *[]VolumeId `json:"volumes,omitempty"`
}

// UpdateInstanceJSONBody defines parameters for UpdateInstance.
type UpdateInstanceJSONBody struct {
	// IsProtected Specifies if the instance is termination protected.
	// When set to `true`, it"s not possible to destroy the instance until it"s switched to `false`.
	// Set to `true` automatically for long-term billed instances.
	IsProtected *InstanceIsProtected `json:"is_protected,omitempty"`

	// Name The human-readable name set for the instance.
	Name *InstanceName `json:"name,omitempty"`

	// SecurityGroups An array of security group ids.
	// **Please Note**: By default the **standard security group** is set if you don"t specify any Security Groups.
	// You can override this behavior by providing a different Security Group.
	SecurityGroups *InstanceSecurityGroupIds `json:"security_groups,omitempty"`

	// Volumes The instance's volumes IDs.
	Volumes *[]VolumeId `json:"volumes,omitempty"`
}

// PerformInstanceActionJSONBody defines parameters for PerformInstanceAction.
type PerformInstanceActionJSONBody struct {
	Action InstanceAction `json:"action"`
}

// ListInstanceSnapshotsParams defines parameters for ListInstanceSnapshots.
type ListInstanceSnapshotsParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreateSnapshotJSONBody defines parameters for CreateSnapshot.
type CreateSnapshotJSONBody struct {
	// Name Name of the snapshot
	Name string `json:"name"`
}

// ListSecurityGroupsParams defines parameters for ListSecurityGroups.
type ListSecurityGroupsParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreateSecurityGroupJSONBody defines parameters for CreateSecurityGroup.
type CreateSecurityGroupJSONBody struct {
	// Description he human-readable description set for the security group.
	Description *string `json:"description,omitempty"`

	// Name The human-readable name set for the security group. **Please Note**: `default` and `standard` are not allowed names (reserved words).
	Name string `json:"name"`

	// Region The region identifier.
	Region Region `json:"region"`

	// Rules The list of rules of the security group.
	Rules []SecurityGroupRule `json:"rules"`
}

// UpdateSecurityGroupJSONBody defines parameters for UpdateSecurityGroup.
type UpdateSecurityGroupJSONBody struct {
	// Description he human-readable description set for the security group.
	Description *string `json:"description,omitempty"`

	// Name The human-readable name set for the security group.
	Name *string `json:"name,omitempty"`

	// Rules The list of rules of the security group.
	Rules *[]SecurityGroupRule `json:"rules,omitempty"`
}

// ListSnapshotsParams defines parameters for ListSnapshots.
type ListSnapshotsParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// UpdateSnapshotJSONBody defines parameters for UpdateSnapshot.
type UpdateSnapshotJSONBody struct {
	// Name The new human-readable name for your snapshot.
	Name *string `json:"name,omitempty"`
}

// ListSSHKeysParams defines parameters for ListSSHKeys.
type ListSSHKeysParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreateSSHKeyJSONBody defines parameters for CreateSSHKey.
type CreateSSHKeyJSONBody struct {
	// Name The human-readable name for your ssh key.
	Name string `json:"name"`

	// Value SSH public key
	Value string `json:"value"`
}

// UpdateSSHKeyJSONBody defines parameters for UpdateSSHKey.
type UpdateSSHKeyJSONBody struct {
	// Name The new human-readable name for your ssh key.
	Name *string `json:"name,omitempty"`
}

// ListVolumesParams defines parameters for ListVolumes.
type ListVolumesParams struct {
	Page    *PageQueryParameter    `form:"page,omitempty" json:"page,omitempty"`
	PerPage *PerPageQueryParameter `form:"per_page,omitempty" json:"per_page,omitempty"`
}

// CreateVolumeJSONBody defines parameters for CreateVolume.
type CreateVolumeJSONBody struct {
	// Description The human-readable description set for the volume.
	Description *string `json:"description,omitempty"`

	// Name The human-readable name set for the volume.
	Name string `json:"name"`

	// Region The region identifier.
	Region Region `json:"region"`

	// Size The storage size of this volume given in GiB (Min: 1GiB).
	Size int `json:"size"`

	// Type The volume type.
	Type *VolumeType `json:"type,omitempty"`
}

// UpdateVolumeJSONBody defines parameters for UpdateVolume.
type UpdateVolumeJSONBody struct {
	// Description The human-readable description set for the volume.
	Description *string `json:"description,omitempty"`

	// Name The human-readable name set for the volume.
	Name *string `json:"name,omitempty"`
}

// CreateFloatingIPJSONRequestBody defines body for CreateFloatingIP for application/json ContentType.
type CreateFloatingIPJSONRequestBody CreateFloatingIPJSONBody

// CreateInstanceJSONRequestBody defines body for CreateInstance for application/json ContentType.
type CreateInstanceJSONRequestBody CreateInstanceJSONBody

// UpdateInstanceJSONRequestBody defines body for UpdateInstance for application/json ContentType.
type UpdateInstanceJSONRequestBody UpdateInstanceJSONBody

// PerformInstanceActionJSONRequestBody defines body for PerformInstanceAction for application/json ContentType.
type PerformInstanceActionJSONRequestBody PerformInstanceActionJSONBody

// CreateSnapshotJSONRequestBody defines body for CreateSnapshot for application/json ContentType.
type CreateSnapshotJSONRequestBody CreateSnapshotJSONBody

// CreateSecurityGroupJSONRequestBody defines body for CreateSecurityGroup for application/json ContentType.
type CreateSecurityGroupJSONRequestBody CreateSecurityGroupJSONBody

// UpdateSecurityGroupJSONRequestBody defines body for UpdateSecurityGroup for application/json ContentType.
type UpdateSecurityGroupJSONRequestBody UpdateSecurityGroupJSONBody

// UpdateSnapshotJSONRequestBody defines body for UpdateSnapshot for application/json ContentType.
type UpdateSnapshotJSONRequestBody UpdateSnapshotJSONBody

// CreateSSHKeyJSONRequestBody defines body for CreateSSHKey for application/json ContentType.
type CreateSSHKeyJSONRequestBody CreateSSHKeyJSONBody

// UpdateSSHKeyJSONRequestBody defines body for UpdateSSHKey for application/json ContentType.
type UpdateSSHKeyJSONRequestBody UpdateSSHKeyJSONBody

// CreateVolumeJSONRequestBody defines body for CreateVolume for application/json ContentType.
type CreateVolumeJSONRequestBody CreateVolumeJSONBody

// UpdateVolumeJSONRequestBody defines body for UpdateVolume for application/json ContentType.
type UpdateVolumeJSONRequestBody UpdateVolumeJSONBody
