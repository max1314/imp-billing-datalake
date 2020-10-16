// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/release_notes/release_notes.proto

package improbable_int_release_notes

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

func (this *ReleaseNotesToolConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	majorFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Major)
	hasInner = hasInner || len(majorFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range majorFields {
		res[k] = v
	}
	return res
}

func (this *ReleaseNotesToolConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Major, dst)
}

func (this *ReleaseNotesToolConfig_Major) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ReleaseNotesToolConfig_Major) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ReleaseNote) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ReleaseNote) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ReleaseNotesMetadata) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ReleaseNotesMetadata) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ReleaseMetadata) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ReleaseMetadata) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Category) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Category) ExtractRequestFields(dst map[string]interface{}) {
}
