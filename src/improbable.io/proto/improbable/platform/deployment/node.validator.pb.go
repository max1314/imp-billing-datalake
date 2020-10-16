// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/deployment/node.proto

package improbable_platform_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/platform"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NodeDebugUrls) Validate() error {
	return nil
}

var _regex_NodeId_Name = regexp.MustCompile(`^[a-zA-Z0-9_]{2,25}$`)

func (this *NodeId) Validate() error {
	if nil == this.DeploymentId {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", fmt.Errorf("message must exist"))
	}
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	if !_regex_NodeId_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9_]{2,25}$"`, this.Name))
	}
	return nil
}
func (this *Node) Validate() error {
	if nil == this.NodeId {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf("message must exist"))
	}
	if this.NodeId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodeId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodeId", err)
		}
	}
	if this.DebugUrls != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DebugUrls); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DebugUrls", err)
		}
	}
	for _, item := range this.Links {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Links", err)
			}
		}
	}
	if this.LastMachineError != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastMachineError); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastMachineError", err)
		}
	}
	return nil
}
func (this *Node_Link) Validate() error {
	return nil
}
func (this *MultiNodeCreateRequest) Validate() error {
	if nil == this.DeploymentId {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", fmt.Errorf("message must exist"))
	}
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	for _, item := range this.Nodes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Nodes", err)
			}
		}
	}
	return nil
}

var _regex_SetStateRequest_State = regexp.MustCompile(`^(configured|provisioned|running)$`)

func (this *SetStateRequest) Validate() error {
	if nil == this.NodeId {
		return github_com_mwitkow_go_proto_validators.FieldError("NodeId", fmt.Errorf("message must exist"))
	}
	if this.NodeId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.NodeId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("NodeId", err)
		}
	}
	if !_regex_SetStateRequest_State.MatchString(this.State) {
		return github_com_mwitkow_go_proto_validators.FieldError("State", fmt.Errorf(`value '%v' must be a string conforming to regex "^(configured|provisioned|running)$"`, this.State))
	}
	return nil
}

var _regex_SetAllStatesRequest_State = regexp.MustCompile(`^(configured|provisioned|running)$`)

func (this *SetAllStatesRequest) Validate() error {
	if nil == this.DeploymentId {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", fmt.Errorf("message must exist"))
	}
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	if !_regex_SetAllStatesRequest_State.MatchString(this.State) {
		return github_com_mwitkow_go_proto_validators.FieldError("State", fmt.Errorf(`value '%v' must be a string conforming to regex "^(configured|provisioned|running)$"`, this.State))
	}
	return nil
}