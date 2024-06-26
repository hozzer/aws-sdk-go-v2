// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ChangeProgressStageStatuses string

// Enum values for ChangeProgressStageStatuses
const (
	ChangeProgressStageStatusesPending    ChangeProgressStageStatuses = "PENDING"
	ChangeProgressStageStatusesInProgress ChangeProgressStageStatuses = "IN_PROGRESS"
	ChangeProgressStageStatusesCompleted  ChangeProgressStageStatuses = "COMPLETED"
	ChangeProgressStageStatusesFailed     ChangeProgressStageStatuses = "FAILED"
)

// Values returns all known values for ChangeProgressStageStatuses. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChangeProgressStageStatuses) Values() []ChangeProgressStageStatuses {
	return []ChangeProgressStageStatuses{
		"PENDING",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
	}
}

type ChangeProgressStatuses string

// Enum values for ChangeProgressStatuses
const (
	ChangeProgressStatusesPending    ChangeProgressStatuses = "PENDING"
	ChangeProgressStatusesInProgress ChangeProgressStatuses = "IN_PROGRESS"
	ChangeProgressStatusesCompleted  ChangeProgressStatuses = "COMPLETED"
	ChangeProgressStatusesFailed     ChangeProgressStatuses = "FAILED"
)

// Values returns all known values for ChangeProgressStatuses. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChangeProgressStatuses) Values() []ChangeProgressStatuses {
	return []ChangeProgressStatuses{
		"PENDING",
		"IN_PROGRESS",
		"COMPLETED",
		"FAILED",
	}
}

type PipelineStatus string

// Enum values for PipelineStatus
const (
	PipelineStatusCreating     PipelineStatus = "CREATING"
	PipelineStatusActive       PipelineStatus = "ACTIVE"
	PipelineStatusUpdating     PipelineStatus = "UPDATING"
	PipelineStatusDeleting     PipelineStatus = "DELETING"
	PipelineStatusCreateFailed PipelineStatus = "CREATE_FAILED"
	PipelineStatusUpdateFailed PipelineStatus = "UPDATE_FAILED"
	PipelineStatusStarting     PipelineStatus = "STARTING"
	PipelineStatusStartFailed  PipelineStatus = "START_FAILED"
	PipelineStatusStopping     PipelineStatus = "STOPPING"
	PipelineStatusStopped      PipelineStatus = "STOPPED"
)

// Values returns all known values for PipelineStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PipelineStatus) Values() []PipelineStatus {
	return []PipelineStatus{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"CREATE_FAILED",
		"UPDATE_FAILED",
		"STARTING",
		"START_FAILED",
		"STOPPING",
		"STOPPED",
	}
}

type VpcEndpointServiceName string

// Enum values for VpcEndpointServiceName
const (
	VpcEndpointServiceNameOpensearchServerless VpcEndpointServiceName = "OPENSEARCH_SERVERLESS"
)

// Values returns all known values for VpcEndpointServiceName. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VpcEndpointServiceName) Values() []VpcEndpointServiceName {
	return []VpcEndpointServiceName{
		"OPENSEARCH_SERVERLESS",
	}
}
