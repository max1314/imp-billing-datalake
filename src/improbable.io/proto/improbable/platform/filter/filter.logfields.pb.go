// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/filter/filter.proto

package improbable_platform_filter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PropertyFilter) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	valueFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Value)
	hasInner = hasInner || len(valueFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range valueFields {
		res[k] = v
	}
	return res
}

func (this *PropertyFilter) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Value, dst)
}