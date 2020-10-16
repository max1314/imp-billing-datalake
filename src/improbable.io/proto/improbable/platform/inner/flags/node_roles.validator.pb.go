// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/inner/flags/node_roles.proto

package improbable_platform_inner_flags

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *NodeRoles) Validate() error {
	return nil
}
func (this *DeploymentNodeRoles) Validate() error {
	for _, item := range this.NodeRoles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeRoles", err)
			}
		}
	}
	return nil
}
