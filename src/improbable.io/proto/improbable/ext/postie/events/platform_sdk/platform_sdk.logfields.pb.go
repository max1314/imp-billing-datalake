// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/ext/postie/events/platform_sdk/platform_sdk.proto

package improbable_ext_postie_events_platform_sdk

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "improbable.io/proto/improbable/ext/postie/options"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Header) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	grpcStartTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GrpcStartTime)
	hasInner = hasInner || len(grpcStartTimeFields) > 0
	grpcRequestDeadlineFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GrpcRequestDeadline)
	hasInner = hasInner || len(grpcRequestDeadlineFields) > 0
	atTimestampFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AtTimestamp)
	hasInner = hasInner || len(atTimestampFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range grpcStartTimeFields {
		res[k] = v
	}
	for k, v := range grpcRequestDeadlineFields {
		res[k] = v
	}
	for k, v := range atTimestampFields {
		res[k] = v
	}
	return res
}

func (this *Header) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GrpcStartTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GrpcRequestDeadline, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AtTimestamp, dst)
}

func (this *ListDeployments) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ListDeployments) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *GetDeployment) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *GetDeployment) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateDeployment) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateDeployment) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *UpdateDeployment) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *UpdateDeployment) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *StopDeployment) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *StopDeployment) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ListSnapshots) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ListSnapshots) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *GetSnapshot) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *GetSnapshot) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *TakeSnapshot) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *TakeSnapshot) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *UploadSnapshot) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *UploadSnapshot) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ConfirmSnapshotUpload) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ConfirmSnapshotUpload) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateServiceAccount) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateServiceAccount) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *DeleteServiceAccount) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *DeleteServiceAccount) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ListServiceAccounts) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ListServiceAccounts) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *UpdateServiceAccount) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *UpdateServiceAccount) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateLoginToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateLoginToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreatePlayerIdentityToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreatePlayerIdentityToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *DecodePlayerIdentityToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *DecodePlayerIdentityToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateDevelopmentAuthenticationToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateDevelopmentAuthenticationToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *GetDevelopmentAuthenticationToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *GetDevelopmentAuthenticationToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ListDevelopmentAuthenticationTokens) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ListDevelopmentAuthenticationTokens) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *UpdateDevelopmentAuthenticationToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *UpdateDevelopmentAuthenticationToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ExpireDevelopmentAuthenticationToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ExpireDevelopmentAuthenticationToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *EntityCommand) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *EntityCommand) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ReserveEntityID) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ReserveEntityID) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateEntity) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateEntity) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *DeleteEntity) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *DeleteEntity) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *EntityQuery) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *EntityQuery) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateDevelopmentPlayerIdentityToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateDevelopmentPlayerIdentityToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *CreateDevelopmentLoginTokens) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *CreateDevelopmentLoginTokens) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *ElevateProjectPermissions) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *ElevateProjectPermissions) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *RemoveGroup) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *RemoveGroup) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}

func (this *AddAccountToGroup) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	headerFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Header)
	hasInner = hasInner || len(headerFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range headerFields {
		res[k] = v
	}
	return res
}

func (this *AddAccountToGroup) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Header, dst)
}