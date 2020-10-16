// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/playtest/playtest_share.proto

package improbable_platform_playtest

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PlaytestShareToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	expiryFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Expiry)
	hasInner = hasInner || len(expiryFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range expiryFields {
		res[k] = v
	}
	return res
}

func (this *PlaytestShareToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Expiry, dst)
}

func (this *CreatePlaytestShareTokenRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	lifetimeDurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.LifetimeDuration)
	hasInner = hasInner || len(lifetimeDurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range lifetimeDurationFields {
		res[k] = v
	}
	return res
}

func (this *CreatePlaytestShareTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LifetimeDuration, dst)
}

func (this *CreatePlaytestShareTokenResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreatePlaytestShareTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DecodePlaytestShareTokenRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DecodePlaytestShareTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DecodePlaytestShareTokenResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	tokenFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Token)
	hasInner = hasInner || len(tokenFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range tokenFields {
		res[k] = v
	}
	return res
}

func (this *DecodePlaytestShareTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Token, dst)
}
