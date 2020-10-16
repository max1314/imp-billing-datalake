// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/scheduler/jobs.proto

package improbable_int_platform_scheduler

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *JobState) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	simulationPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SimulationPlaneId)
	hasInner = hasInner || len(simulationPlaneIdFields) > 0
	assignedMachineFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AssignedMachine)
	hasInner = hasInner || len(assignedMachineFields) > 0
	dockerParametersFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DockerParameters)
	hasInner = hasInner || len(dockerParametersFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range simulationPlaneIdFields {
		res[k] = v
	}
	for k, v := range assignedMachineFields {
		res[k] = v
	}
	for k, v := range dockerParametersFields {
		res[k] = v
	}
	return res
}

func (this *JobState) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SimulationPlaneId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AssignedMachine, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DockerParameters, dst)
}

func (this *GetJobRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	return res
}

func (this *GetJobRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *GetJobResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	jobStateFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.JobState)
	hasInner = hasInner || len(jobStateFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range jobStateFields {
		res[k] = v
	}
	return res
}

func (this *GetJobResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.JobState, dst)
}

func (this *ListDeploymentJobsRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentId)
	hasInner = hasInner || len(deploymentIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentIdFields {
		res[k] = v
	}
	return res
}

func (this *ListDeploymentJobsRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
}

func (this *ListDeploymentJobsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListDeploymentJobsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAllJobsRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAllJobsRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAllJobsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAllJobsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CreateJobRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	dockerParametersFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DockerParameters)
	hasInner = hasInner || len(dockerParametersFields) > 0
	jobParametersFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.JobParameters)
	hasInner = hasInner || len(jobParametersFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range dockerParametersFields {
		res[k] = v
	}
	for k, v := range jobParametersFields {
		res[k] = v
	}
	return res
}

func (this *CreateJobRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DockerParameters, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.JobParameters, dst)
}

func (this *CreateJobResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateJobResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *UnassignJobRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	jobIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.JobId)
	hasInner = hasInner || len(jobIdFields) > 0
	prevMachineIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.PrevMachineId)
	hasInner = hasInner || len(prevMachineIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range jobIdFields {
		res[k] = v
	}
	for k, v := range prevMachineIdFields {
		res[k] = v
	}
	return res
}

func (this *UnassignJobRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.JobId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.PrevMachineId, dst)
}

func (this *UnassignJobResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *UnassignJobResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SetJobAttributesRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	runnerIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.RunnerId)
	hasInner = hasInner || len(runnerIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"current_state": this.CurrentState,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	res["current_state"] = this.CurrentState
	for k, v := range runnerIdFields {
		res[k] = v
	}
	return res
}

func (this *SetJobAttributesRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	dst["current_state"] = this.CurrentState
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.RunnerId, dst)
}

func (this *SetJobAttributesResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SetJobAttributesResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *LockLaunchRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	jobIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.JobId)
	hasInner = hasInner || len(jobIdFields) > 0
	runnerIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.RunnerId)
	hasInner = hasInner || len(runnerIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range jobIdFields {
		res[k] = v
	}
	for k, v := range runnerIdFields {
		res[k] = v
	}
	return res
}

func (this *LockLaunchRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.JobId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.RunnerId, dst)
}

func (this *LockLaunchResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LockLaunchResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SetJobTargetStateRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"target_state": fmt.Sprintf("%v", this.TargetState),
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	res["target_state"] = fmt.Sprintf("%v", this.TargetState)
	return res
}

func (this *SetJobTargetStateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	dst["target_state"] = this.TargetState
}

func (this *SetJobTargetStateResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SetJobTargetStateResponse) ExtractRequestFields(dst map[string]interface{}) {
}