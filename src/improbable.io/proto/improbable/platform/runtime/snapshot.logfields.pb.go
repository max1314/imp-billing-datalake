// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/runtime/snapshot.proto

package improbable_platform_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/platform/history"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TakeSnapshotRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *TakeSnapshotRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *TakeSnapshotResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	snapshotIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SnapshotId)
	hasInner = hasInner || len(snapshotIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range snapshotIdFields {
		res[k] = v
	}
	return res
}

func (this *TakeSnapshotResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SnapshotId, dst)
}