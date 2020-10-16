// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/inner/nanny_status.proto

package improbable_platform_inner

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SetStartedRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	nodeIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.NodeId)
	hasInner = hasInner || len(nodeIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range nodeIdFields {
		res[k] = v
	}
	return res
}

func (this *SetStartedRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.NodeId, dst)
}

func (this *IsStartedRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	nodeIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.NodeId)
	hasInner = hasInner || len(nodeIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range nodeIdFields {
		res[k] = v
	}
	return res
}

func (this *IsStartedRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.NodeId, dst)
}

func (this *IsStartedResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *IsStartedResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SetWorkerTerminatedRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	nodeIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.NodeId)
	hasInner = hasInner || len(nodeIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range nodeIdFields {
		res[k] = v
	}
	return res
}

func (this *SetWorkerTerminatedRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.NodeId, dst)
}

func (this *IsWorkerTerminatedRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	nodeIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.NodeId)
	hasInner = hasInner || len(nodeIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range nodeIdFields {
		res[k] = v
	}
	return res
}

func (this *IsWorkerTerminatedRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.NodeId, dst)
}

func (this *IsWorkerTerminatedResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *IsWorkerTerminatedResponse) ExtractRequestFields(dst map[string]interface{}) {
}
