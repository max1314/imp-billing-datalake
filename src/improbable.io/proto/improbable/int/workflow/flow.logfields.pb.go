// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/workflow/flow.proto

package improbable_int_workflow

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "improbable.io/proto/infra/accounts"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *FlowConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	serviceAccountFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ServiceAccount)
	hasInner = hasInner || len(serviceAccountFields) > 0
	metadataFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Metadata)
	hasInner = hasInner || len(metadataFields) > 0
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	expiryTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpiryTime)
	hasInner = hasInner || len(expiryTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range serviceAccountFields {
		res[k] = v
	}
	for k, v := range metadataFields {
		res[k] = v
	}
	for k, v := range creationTimeFields {
		res[k] = v
	}
	for k, v := range expiryTimeFields {
		res[k] = v
	}
	return res
}

func (this *FlowConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ServiceAccount, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Metadata, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpiryTime, dst)
}

func (this *FlowStatus) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	stateFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.State)
	hasInner = hasInner || len(stateFields) > 0
	errorFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Error)
	hasInner = hasInner || len(errorFields) > 0
	lastStepErrorFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.LastStepError)
	hasInner = hasInner || len(lastStepErrorFields) > 0
	rescheduleTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.RescheduleTime)
	hasInner = hasInner || len(rescheduleTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range stateFields {
		res[k] = v
	}
	for k, v := range errorFields {
		res[k] = v
	}
	for k, v := range lastStepErrorFields {
		res[k] = v
	}
	for k, v := range rescheduleTimeFields {
		res[k] = v
	}
	return res
}

func (this *FlowStatus) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.State, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Error, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LastStepError, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.RescheduleTime, dst)
}

func (this *Error) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Error) ExtractRequestFields(dst map[string]interface{}) {
}
