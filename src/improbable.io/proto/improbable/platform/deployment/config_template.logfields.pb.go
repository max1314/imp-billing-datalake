// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/deployment/config_template.proto

package improbable_platform_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/platform"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateTemplateRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	projectIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ProjectId)
	hasInner = hasInner || len(projectIdFields) > 0
	configurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Configuration)
	hasInner = hasInner || len(configurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range projectIdFields {
		res[k] = v
	}
	for k, v := range configurationFields {
		res[k] = v
	}
	return res
}

func (this *CreateTemplateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Configuration, dst)
}

func (this *CreateBaseTemplateRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	configurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Configuration)
	hasInner = hasInner || len(configurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range configurationFields {
		res[k] = v
	}
	return res
}

func (this *CreateBaseTemplateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Configuration, dst)
}

func (this *CreateTemplateResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateTemplateResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetTemplateRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	projectIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ProjectId)
	hasInner = hasInner || len(projectIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range projectIdFields {
		res[k] = v
	}
	return res
}

func (this *GetTemplateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectId, dst)
}

func (this *GetTemplateResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	configurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Configuration)
	hasInner = hasInner || len(configurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range configurationFields {
		res[k] = v
	}
	return res
}

func (this *GetTemplateResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Configuration, dst)
}

func (this *FromLaunchConfigurationRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	projectIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ProjectId)
	hasInner = hasInner || len(projectIdFields) > 0
	launchConfigurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.LaunchConfiguration)
	hasInner = hasInner || len(launchConfigurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range projectIdFields {
		res[k] = v
	}
	for k, v := range launchConfigurationFields {
		res[k] = v
	}
	return res
}

func (this *FromLaunchConfigurationRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LaunchConfiguration, dst)
}

func (this *FromLaunchConfigurationResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	configurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Configuration)
	hasInner = hasInner || len(configurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range configurationFields {
		res[k] = v
	}
	return res
}

func (this *FromLaunchConfigurationResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Configuration, dst)
}