// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/fabric/devtools/service/release_tool/fabric_version.proto

package improbable_fabric_devtools_service_release_tool

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

func (this *BuildStatus) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	runStartTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.RunStartTime)
	hasInner = hasInner || len(runStartTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range runStartTimeFields {
		res[k] = v
	}
	return res
}

func (this *BuildStatus) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.RunStartTime, dst)
}

func (this *FabricVersion) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	buildStatusFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BuildStatus)
	hasInner = hasInner || len(buildStatusFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range buildStatusFields {
		res[k] = v
	}
	return res
}

func (this *FabricVersion) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BuildStatus, dst)
}
