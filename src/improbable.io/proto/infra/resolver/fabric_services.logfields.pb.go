// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/resolver/fabric_services.proto

package improbable_infra

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ResolveServiceRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Deployment)
	hasInner = hasInner || len(deploymentFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentFields {
		res[k] = v
	}
	return res
}

func (this *ResolveServiceRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Deployment, dst)
}

func (this *ResolveModuleRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Deployment)
	hasInner = hasInner || len(deploymentFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentFields {
		res[k] = v
	}
	return res
}

func (this *ResolveModuleRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Deployment, dst)
}

func (this *ResolveServiceResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ResolveServiceResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ResolveModuleResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ResolveModuleResponse) ExtractRequestFields(dst map[string]interface{}) {
}