// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/history/history.proto

package improbable_platform_history

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	_ "improbable.io/proto/improbable/platform"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *HistoryId) LogFields() map[string]string {
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
		return map[string]string{
			"history_name": this.Name,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range projectIdFields {
		res[k] = v
	}
	res["history_name"] = this.Name
	return res
}

func (this *HistoryId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectId, dst)
	dst["history_name"] = this.Name
}

func (this *History) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	historyIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.HistoryId)
	hasInner = hasInner || len(historyIdFields) > 0
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range historyIdFields {
		res[k] = v
	}
	for k, v := range creationTimeFields {
		res[k] = v
	}
	return res
}

func (this *History) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.HistoryId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
}