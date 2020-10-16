// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/player_auth/development_login.proto

package improbable_platform_player_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_CreateLoginTokensRequest_WorkerType = regexp.MustCompile(`^$|^[a-zA-Z0-9_-]{3,128}$`)

func (this *CreateLoginTokensRequest) Validate() error {
	if this.PlayerIdentityToken == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PlayerIdentityToken", fmt.Errorf(`value '%v' must not be an empty string`, this.PlayerIdentityToken))
	}
	if this.LifetimeDuration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LifetimeDuration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LifetimeDuration", err)
		}
	}
	if !_regex_CreateLoginTokensRequest_WorkerType.MatchString(this.WorkerType) {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkerType", fmt.Errorf(`value '%v' must be a string conforming to regex "^$|^[a-zA-Z0-9_-]{3,128}$"`, this.WorkerType))
	}
	return nil
}
func (this *CreateLoginTokensResponse) Validate() error {
	for _, item := range this.LoginTokenDetails {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LoginTokenDetails", err)
			}
		}
	}
	return nil
}
func (this *LoginTokenDetails) Validate() error {
	return nil
}
