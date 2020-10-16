// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/deployment/worker_descriptor.proto

package improbable_platform_deployment

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

func (this *WorkerZip) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	zipFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Zip)
	hasInner = hasInner || len(zipFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range zipFields {
		res[k] = v
	}
	return res
}

func (this *WorkerZip) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Zip, dst)
}

func (this *WorkerLocation) LogFields() map[string]string {
	return map[string]string{}
}

func (this *WorkerLocation) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *WorkerStartup) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	var workerStartupTypesFields map[string]string
	switch f := this.WorkerStartupTypes.(type) {
	case *WorkerStartup_WorkerZip:
		workerStartupTypesFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.WorkerZip)
	default:
		workerStartupTypesFields = map[string]string{}
	}
	hasInner = hasInner || len(workerStartupTypesFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range workerStartupTypesFields {
		res[k] = v
	}
	return res
}

func (this *WorkerStartup) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	if f, ok := this.WorkerStartupTypes.(*WorkerStartup_WorkerZip); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.WorkerZip, dst)
	}
}

func (this *WorkerDescriptor) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	launchFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Launch)
	hasInner = hasInner || len(launchFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range launchFields {
		res[k] = v
	}
	return res
}

func (this *WorkerDescriptor) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Launch, dst)
}