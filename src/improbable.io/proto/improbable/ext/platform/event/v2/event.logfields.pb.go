// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/ext/platform/event/v2/event.proto

package improbable_ext_platform_event_v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Event) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	timestampFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Timestamp)
	hasInner = hasInner || len(timestampFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range timestampFields {
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
}

func (this *ProcessRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ProcessRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ProcessResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ProcessResponse) ExtractRequestFields(dst map[string]interface{}) {
}
