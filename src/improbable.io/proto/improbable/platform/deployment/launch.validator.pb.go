// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/deployment/launch.proto

package improbable_platform_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_LaunchConfiguration_Template = regexp.MustCompile(`^[a-zA-Z0-9_]{1,}$`)

func (this *LaunchConfiguration) Validate() error {
	if !_regex_LaunchConfiguration_Template.MatchString(this.Template) {
		return github_com_mwitkow_go_proto_validators.FieldError("Template", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9_]{1,}$"`, this.Template))
	}
	if nil == this.World {
		return github_com_mwitkow_go_proto_validators.FieldError("World", fmt.Errorf("message must exist"))
	}
	if this.World != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.World); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("World", err)
		}
	}
	for _, item := range this.Workers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Workers", err)
			}
		}
	}
	if this.LoadBalancing != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LoadBalancing); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LoadBalancing", err)
		}
	}
	return nil
}
func (this *LaunchConfiguration_World) Validate() error {
	if this.Snapshots != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Snapshots); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Snapshots", err)
		}
	}
	for _, item := range this.LegacyFlags {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LegacyFlags", err)
			}
		}
	}
	if nil == this.Dimensions {
		return github_com_mwitkow_go_proto_validators.FieldError("Dimensions", fmt.Errorf("message must exist"))
	}
	if this.Dimensions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Dimensions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Dimensions", err)
		}
	}
	for _, item := range this.LegacyJavaparams {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LegacyJavaparams", err)
			}
		}
	}
	return nil
}
func (this *LaunchConfiguration_World_Snapshots) Validate() error {
	return nil
}
func (this *LaunchConfiguration_World_Dimensions) Validate() error {
	return nil
}
func (this *LaunchConfiguration_WorkerConfig) Validate() error {
	for _, item := range this.Flags {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Flags", err)
			}
		}
	}
	if this.LoadBalancing != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LoadBalancing); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LoadBalancing", err)
		}
	}
	for _, item := range this.Permissions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Permissions", err)
			}
		}
	}
	if this.LoginRateLimit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LoginRateLimit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LoginRateLimit", err)
		}
	}
	if this.ConnectionCapacityLimit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionCapacityLimit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionCapacityLimit", err)
		}
	}
	return nil
}

var _regex_LoginRateLimit_Duration = regexp.MustCompile(`^([0-9]+h)?([0-9]+m)?([0-9]+s)?$`)

func (this *LoginRateLimit) Validate() error {
	if !_regex_LoginRateLimit_Duration.MatchString(this.Duration) {
		return github_com_mwitkow_go_proto_validators.FieldError("Duration", fmt.Errorf(`value '%v' must be a string conforming to regex "^([0-9]+h)?([0-9]+m)?([0-9]+s)?$"`, this.Duration))
	}
	return nil
}
func (this *CapacityLimit) Validate() error {
	if !(this.MaxCapacity > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("MaxCapacity", fmt.Errorf(`value '%v' must be greater than '-1'`, this.MaxCapacity))
	}
	return nil
}
