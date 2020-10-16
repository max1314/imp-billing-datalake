// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/event/event.proto

package improbable_infra_event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Error) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Error) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SpatialCommand) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SpatialCommand) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Event) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	timestampFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Timestamp)
	hasInner = hasInner || len(timestampFields) > 0
	spatialCommandFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SpatialCommand)
	hasInner = hasInner || len(spatialCommandFields) > 0
	errorFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Error)
	hasInner = hasInner || len(errorFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range timestampFields {
		res[k] = v
	}
	for k, v := range spatialCommandFields {
		res[k] = v
	}
	for k, v := range errorFields {
		res[k] = v
	}
	return res
}

func (this *Event) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Timestamp, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SpatialCommand, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Error, dst)
}

func (this *SendRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SendRequest) ExtractRequestFields(dst map[string]interface{}) {
}