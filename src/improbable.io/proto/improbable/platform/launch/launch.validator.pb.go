// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/launch/launch.proto

package improbable_platform_launch

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_LoginRequest_ProjectName = regexp.MustCompile(`^[a-z0-9_]{3,32}$`)
var _regex_LoginRequest_DeploymentName = regexp.MustCompile(`^$|^[a-z0-9_]{2,32}$`)
var _regex_LoginRequest_DeploymentTag = regexp.MustCompile(`$|^[a-z0-9_]{3,32}$`)

func (this *LoginRequest) Validate() error {
	if !_regex_LoginRequest_ProjectName.MatchString(this.ProjectName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9_]{3,32}$"`, this.ProjectName))
	}
	if !_regex_LoginRequest_DeploymentName.MatchString(this.DeploymentName) {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentName", fmt.Errorf(`value '%v' must be a string conforming to regex "^$|^[a-z0-9_]{2,32}$"`, this.DeploymentName))
	}
	if !_regex_LoginRequest_DeploymentTag.MatchString(this.DeploymentTag) {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentTag", fmt.Errorf(`value '%v' must be a string conforming to regex "$|^[a-z0-9_]{3,32}$"`, this.DeploymentTag))
	}
	return nil
}

var _regex_GenerateMultipleRequest_ProjectName = regexp.MustCompile(`^[a-z0-9_]{3,32}$`)
var _regex_GenerateMultipleRequest_DeploymentName = regexp.MustCompile(`^$|^[a-z0-9_]{2,32}$`)

func (this *GenerateMultipleRequest) Validate() error {
	if !_regex_GenerateMultipleRequest_ProjectName.MatchString(this.ProjectName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9_]{3,32}$"`, this.ProjectName))
	}
	if !_regex_GenerateMultipleRequest_DeploymentName.MatchString(this.DeploymentName) {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentName", fmt.Errorf(`value '%v' must be a string conforming to regex "^$|^[a-z0-9_]{2,32}$"`, this.DeploymentName))
	}
	return nil
}
func (this *GenerateLoginTokenRequest) Validate() error {
	return nil
}
func (this *GenerateLoginTokenResponse) Validate() error {
	if nil == this.LoginToken {
		return github_com_mwitkow_go_proto_validators.FieldError("LoginToken", fmt.Errorf("message must exist"))
	}
	if this.LoginToken != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LoginToken); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LoginToken", err)
		}
	}
	return nil
}

var _regex_GenerateShareTokenRequest_ProjectName = regexp.MustCompile(`^[a-z0-9_]{3,32}$`)
var _regex_GenerateShareTokenRequest_DeploymentName = regexp.MustCompile(`^$|^[a-z0-9_]{2,32}$`)

func (this *GenerateShareTokenRequest) Validate() error {
	if !_regex_GenerateShareTokenRequest_ProjectName.MatchString(this.ProjectName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9_]{3,32}$"`, this.ProjectName))
	}
	if !_regex_GenerateShareTokenRequest_DeploymentName.MatchString(this.DeploymentName) {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentName", fmt.Errorf(`value '%v' must be a string conforming to regex "^$|^[a-z0-9_]{2,32}$"`, this.DeploymentName))
	}
	return nil
}
func (this *GenerateShareTokenResponse) Validate() error {
	return nil
}
func (this *ValidateShareTokenRequest) Validate() error {
	return nil
}

var _regex_ValidateShareTokenResponse_ProjectName = regexp.MustCompile(`^[a-z0-9_]{3,32}$`)
var _regex_ValidateShareTokenResponse_DeploymentName = regexp.MustCompile(`^$|^[a-z0-9_]{2,32}$`)

func (this *ValidateShareTokenResponse) Validate() error {
	if !_regex_ValidateShareTokenResponse_ProjectName.MatchString(this.ProjectName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9_]{3,32}$"`, this.ProjectName))
	}
	if !_regex_ValidateShareTokenResponse_DeploymentName.MatchString(this.DeploymentName) {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentName", fmt.Errorf(`value '%v' must be a string conforming to regex "^$|^[a-z0-9_]{2,32}$"`, this.DeploymentName))
	}
	return nil
}
func (this *LoginToken) Validate() error {
	return nil
}

var _regex_GenerateTokensRequest_ProjectName = regexp.MustCompile(`^[a-z0-9_]{3,32}$`)

func (this *GenerateTokensRequest) Validate() error {
	if !_regex_GenerateTokensRequest_ProjectName.MatchString(this.ProjectName) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProjectName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z0-9_]{3,32}$"`, this.ProjectName))
	}
	return nil
}
func (this *GenerateTokensResponse) Validate() error {
	return nil
}
func (this *GenerateAnonymousTokensRequest) Validate() error {
	return nil
}
