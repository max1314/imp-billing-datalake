// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/inner/flags/gsim_configuration.proto

package improbable_platform_inner_flags

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GsimConfiguration) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	extentFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Extent)
	hasInner = hasInner || len(extentFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range extentFields {
		res[k] = v
	}
	return res
}

func (this *GsimConfiguration) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Extent, dst)
}

func (this *GsimConfiguration_GsimExtent) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bottomLeftFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BottomLeft)
	hasInner = hasInner || len(bottomLeftFields) > 0
	topRightFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.TopRight)
	hasInner = hasInner || len(topRightFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bottomLeftFields {
		res[k] = v
	}
	for k, v := range topRightFields {
		res[k] = v
	}
	return res
}

func (this *GsimConfiguration_GsimExtent) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BottomLeft, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.TopRight, dst)
}

func (this *GsimConfiguration_GsimExtent_Corner) LogFields() map[string]string {
	return map[string]string{}
}

func (this *GsimConfiguration_GsimExtent_Corner) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *NodeGsimConfiguration) LogFields() map[string]string {
	return map[string]string{}
}

func (this *NodeGsimConfiguration) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeploymentGsimConfiguration) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeploymentGsimConfiguration) ExtractRequestFields(dst map[string]interface{}) {
}
