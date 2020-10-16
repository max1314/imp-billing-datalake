// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/limitz/config.proto

package improbable_infra_limits

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *LimitzConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LimitzConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Rule) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	defaultLimitsFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DefaultLimits)
	hasInner = hasInner || len(defaultLimitsFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range defaultLimitsFields {
		res[k] = v
	}
	return res
}

func (this *Rule) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DefaultLimits, dst)
}

func (this *Limits) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Limits) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Override) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	limitsFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Limits)
	hasInner = hasInner || len(limitsFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range limitsFields {
		res[k] = v
	}
	return res
}

func (this *Override) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Limits, dst)
}