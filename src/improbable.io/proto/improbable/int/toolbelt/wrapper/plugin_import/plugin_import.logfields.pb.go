// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/toolbelt/wrapper/plugin_import/plugin_import.proto

package improbable_toolbelt_wrapper_plugin_import

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

func (this *PluginImportRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	rootCommandFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.RootCommand)
	hasInner = hasInner || len(rootCommandFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range rootCommandFields {
		res[k] = v
	}
	return res
}

func (this *PluginImportRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.RootCommand, dst)
}

func (this *PluginImportResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PluginImportResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Command) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Command) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Flag) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Flag) ExtractRequestFields(dst map[string]interface{}) {
}
