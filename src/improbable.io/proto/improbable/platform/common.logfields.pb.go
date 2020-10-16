// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/common.proto

package improbable_platform

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ProjectId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"project_name": this.Name,
	}
}

func (this *ProjectId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["project_name"] = this.Name
}

func (this *GenericError) LogFields() map[string]string {
	return map[string]string{}
}

func (this *GenericError) ExtractRequestFields(dst map[string]interface{}) {
}
