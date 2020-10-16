// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/scheduler/deployments.proto

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

func (this *CreateDeploymentRequest) LogFields() map[string]string {
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
	return res
}

func (this *CreateDeploymentRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SimulationPlaneId, dst)
}

func (this *CreateDeploymentResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateDeploymentResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeleteDeploymentRequest) LogFields() map[string]string {
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

func (this *DeleteDeploymentRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *DeleteDeploymentResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteDeploymentResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SetDeploymentTargetStateRequest) LogFields() map[string]string {
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

func (this *SetDeploymentTargetStateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	dst["target_state"] = this.TargetState
}

func (this *SetDeploymentTargetStateResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SetDeploymentTargetStateResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *WaitForDeploymentStateRequest) LogFields() map[string]string {
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

func (this *WaitForDeploymentStateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *WaitForDeploymentStateResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *WaitForDeploymentStateResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeploymentState) LogFields() map[string]string {
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

func (this *DeploymentState) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *ListDeploymentsRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListDeploymentsRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListDeploymentsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListDeploymentsResponse) ExtractRequestFields(dst map[string]interface{}) {
}