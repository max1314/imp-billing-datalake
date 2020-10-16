// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/deployment/crash_dump.proto

package improbable_platform_deployment

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

func (this *CrashDumpId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CrashDumpId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CrashDumpMetadata) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range creationTimeFields {
		res[k] = v
	}
	return res
}

func (this *CrashDumpMetadata) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
}

func (this *GetUploadUrlRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range creationTimeFields {
		res[k] = v
	}
	return res
}

func (this *GetUploadUrlRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
}

func (this *CrashDumpUrl) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CrashDumpUrl) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetDownloadUrlRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	crashDumpIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CrashDumpId)
	hasInner = hasInner || len(crashDumpIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range crashDumpIdFields {
		res[k] = v
	}
	return res
}

func (this *GetDownloadUrlRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CrashDumpId, dst)
}

func (this *ListCrashDumpsRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListCrashDumpsRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListCrashDumpsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListCrashDumpsResponse) ExtractRequestFields(dst map[string]interface{}) {
}