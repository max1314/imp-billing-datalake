// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/resolver/fabric_services.proto

package improbable_infra

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_ResolveServiceRequest_Service = regexp.MustCompile(`.+`)

func (this *ResolveServiceRequest) Validate() error {
	if nil == this.Deployment {
		return github_com_mwitkow_go_proto_validators.FieldError("Deployment", fmt.Errorf("message must exist"))
	}
	if this.Deployment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deployment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deployment", err)
		}
	}
	if !_regex_ResolveServiceRequest_Service.MatchString(this.Service) {
		return github_com_mwitkow_go_proto_validators.FieldError("Service", fmt.Errorf(`value '%v' must be a string conforming to regex ".+"`, this.Service))
	}
	return nil
}

var _regex_ResolveModuleRequest_Module = regexp.MustCompile(`.+`)

func (this *ResolveModuleRequest) Validate() error {
	if nil == this.Deployment {
		return github_com_mwitkow_go_proto_validators.FieldError("Deployment", fmt.Errorf("message must exist"))
	}
	if this.Deployment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deployment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deployment", err)
		}
	}
	if !_regex_ResolveModuleRequest_Module.MatchString(this.Module) {
		return github_com_mwitkow_go_proto_validators.FieldError("Module", fmt.Errorf(`value '%v' must be a string conforming to regex ".+"`, this.Module))
	}
	return nil
}
func (this *ResolveServiceResponse) Validate() error {
	return nil
}
func (this *ResolveModuleResponse) Validate() error {
	return nil
}