// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/spatialos/playerauth/v2/playerauth.proto

package pb_playerauthv2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateLoginTokenRequest) LogFields() map[string]string {
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
		return map[string]string{
			"deployment_id": this.DeploymentId,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	res["deployment_id"] = this.DeploymentId
	for k, v := range lifetimeDurationFields {
		res[k] = v
	}
	return res
}

func (this *CreateLoginTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["deployment_id"] = this.DeploymentId
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LifetimeDuration, dst)
}

func (this *CreateLoginTokenResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateLoginTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CreatePlayerIdentityTokenRequest) LogFields() map[string]string {
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
		return map[string]string{
			"project_name": this.ProjectName,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	res["project_name"] = this.ProjectName
	for k, v := range lifetimeDurationFields {
		res[k] = v
	}
	return res
}

func (this *CreatePlayerIdentityTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["project_name"] = this.ProjectName
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LifetimeDuration, dst)
}

func (this *CreatePlayerIdentityTokenResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreatePlayerIdentityTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DecodePlayerIdentityTokenRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DecodePlayerIdentityTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DecodePlayerIdentityTokenResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	playerIdentityTokenFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.PlayerIdentityToken)
	hasInner = hasInner || len(playerIdentityTokenFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range playerIdentityTokenFields {
		res[k] = v
	}
	return res
}

func (this *DecodePlayerIdentityTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.PlayerIdentityToken, dst)
}

func (this *PlayerIdentityToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	issuedAtTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.IssuedAtTime)
	hasInner = hasInner || len(issuedAtTimeFields) > 0
	expiryTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpiryTime)
	hasInner = hasInner || len(expiryTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range issuedAtTimeFields {
		res[k] = v
	}
	for k, v := range expiryTimeFields {
		res[k] = v
	}
	return res
}

func (this *PlayerIdentityToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.IssuedAtTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpiryTime, dst)
}

func (this *LoginToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	issuedAtTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.IssuedAtTime)
	hasInner = hasInner || len(issuedAtTimeFields) > 0
	expiryTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpiryTime)
	hasInner = hasInner || len(expiryTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range issuedAtTimeFields {
		res[k] = v
	}
	for k, v := range expiryTimeFields {
		res[k] = v
	}
	return res
}

func (this *LoginToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.IssuedAtTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpiryTime, dst)
}

func (this *DevelopmentAuthenticationToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	expirationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpirationTime)
	hasInner = hasInner || len(expirationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"dat_id":       this.Id,
			"project_name": this.ProjectName,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	res["dat_id"] = this.Id
	res["project_name"] = this.ProjectName
	for k, v := range creationTimeFields {
		res[k] = v
	}
	for k, v := range expirationTimeFields {
		res[k] = v
	}
	return res
}

func (this *DevelopmentAuthenticationToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["dat_id"] = this.Id
	dst["project_name"] = this.ProjectName
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpirationTime, dst)
}

func (this *CreateDevelopmentAuthenticationTokenRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	lifetimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Lifetime)
	hasInner = hasInner || len(lifetimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"project_name": this.ProjectName,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	res["project_name"] = this.ProjectName
	for k, v := range lifetimeFields {
		res[k] = v
	}
	return res
}

func (this *CreateDevelopmentAuthenticationTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["project_name"] = this.ProjectName
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Lifetime, dst)
}

func (this *CreateDevelopmentAuthenticationTokenResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	developmentAuthenticationTokenFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DevelopmentAuthenticationToken)
	hasInner = hasInner || len(developmentAuthenticationTokenFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range developmentAuthenticationTokenFields {
		res[k] = v
	}
	return res
}

func (this *CreateDevelopmentAuthenticationTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DevelopmentAuthenticationToken, dst)
}

func (this *GetDevelopmentAuthenticationTokenRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"dat_id": this.Id,
	}
}

func (this *GetDevelopmentAuthenticationTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["dat_id"] = this.Id
}

func (this *GetDevelopmentAuthenticationTokenResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	developmentAuthenticationTokenFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DevelopmentAuthenticationToken)
	hasInner = hasInner || len(developmentAuthenticationTokenFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range developmentAuthenticationTokenFields {
		res[k] = v
	}
	return res
}

func (this *GetDevelopmentAuthenticationTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DevelopmentAuthenticationToken, dst)
}

func (this *ListDevelopmentAuthenticationTokensRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"project_name": this.ProjectName,
	}
}

func (this *ListDevelopmentAuthenticationTokensRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["project_name"] = this.ProjectName
}

func (this *ListDevelopmentAuthenticationTokensResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListDevelopmentAuthenticationTokensResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *UpdateDevelopmentAuthenticationTokenRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	updatedLifetimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.UpdatedLifetime)
	hasInner = hasInner || len(updatedLifetimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"dat_id": this.Id,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	res["dat_id"] = this.Id
	for k, v := range updatedLifetimeFields {
		res[k] = v
	}
	return res
}

func (this *UpdateDevelopmentAuthenticationTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["dat_id"] = this.Id
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.UpdatedLifetime, dst)
}

func (this *UpdateDevelopmentAuthenticationTokenResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	developmentAuthenticationTokenFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DevelopmentAuthenticationToken)
	hasInner = hasInner || len(developmentAuthenticationTokenFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range developmentAuthenticationTokenFields {
		res[k] = v
	}
	return res
}

func (this *UpdateDevelopmentAuthenticationTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DevelopmentAuthenticationToken, dst)
}

func (this *ExpireDevelopmentAuthenticationTokenRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"dat_id": this.Id,
	}
}

func (this *ExpireDevelopmentAuthenticationTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["dat_id"] = this.Id
}

func (this *ExpireDevelopmentAuthenticationTokenResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ExpireDevelopmentAuthenticationTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeleteDevelopmentAuthenticationTokenRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"dat_id": this.Id,
	}
}

func (this *DeleteDevelopmentAuthenticationTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["dat_id"] = this.Id
}

func (this *DeleteDevelopmentAuthenticationTokenResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteDevelopmentAuthenticationTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
}
