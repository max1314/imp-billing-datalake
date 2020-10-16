// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/workflow/service.proto

package improbable_int_workflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FlowId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *FlowId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SubmitFlowRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	metadataFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Metadata)
	hasInner = hasInner || len(metadataFields) > 0
	rescheduleTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.RescheduleTime)
	hasInner = hasInner || len(rescheduleTimeFields) > 0
	expiryTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpiryTime)
	hasInner = hasInner || len(expiryTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range metadataFields {
		res[k] = v
	}
	for k, v := range rescheduleTimeFields {
		res[k] = v
	}
	for k, v := range expiryTimeFields {
		res[k] = v
	}
	return res
}

func (this *SubmitFlowRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Metadata, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.RescheduleTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpiryTime, dst)
}

func (this *SubmitFlowResponse) LogFields() map[string]string {
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

func (this *SubmitFlowResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *StartFlowRequest) LogFields() map[string]string {
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

func (this *StartFlowRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *StartFlowResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *StartFlowResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetFlowRequest) LogFields() map[string]string {
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

func (this *GetFlowRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *GetFlowResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	statusFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Status)
	hasInner = hasInner || len(statusFields) > 0
	configFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Config)
	hasInner = hasInner || len(configFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range statusFields {
		res[k] = v
	}
	for k, v := range configFields {
		res[k] = v
	}
	return res
}

func (this *GetFlowResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Status, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Config, dst)
}

func (this *ListFlowsRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	startTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.StartTime)
	hasInner = hasInner || len(startTimeFields) > 0
	endTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.EndTime)
	hasInner = hasInner || len(endTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range startTimeFields {
		res[k] = v
	}
	for k, v := range endTimeFields {
		res[k] = v
	}
	return res
}

func (this *ListFlowsRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.StartTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.EndTime, dst)
}

func (this *ListFlowsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListFlowsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CancelFlowRequest) LogFields() map[string]string {
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

func (this *CancelFlowRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *CancelFlowResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CancelFlowResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CancelRunningFlowByDeploymentIdRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CancelRunningFlowByDeploymentIdRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CancelRunningFlowByDeploymentIdResponse) LogFields() map[string]string {
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

func (this *CancelRunningFlowByDeploymentIdResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *GetRunningFlowByDeploymentIdRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *GetRunningFlowByDeploymentIdRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetRunningFlowByDeploymentIdResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	statusFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Status)
	hasInner = hasInner || len(statusFields) > 0
	configFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Config)
	hasInner = hasInner || len(configFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range statusFields {
		res[k] = v
	}
	for k, v := range configFields {
		res[k] = v
	}
	return res
}

func (this *GetRunningFlowByDeploymentIdResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Status, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Config, dst)
}
