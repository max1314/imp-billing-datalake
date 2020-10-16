// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/locator/queue.proto

package improbable_platform_locator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QueueSizeForDeploymentRequest) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	return nil
}
func (this *QueueSizeForDeploymentResponse) Validate() error {
	return nil
}