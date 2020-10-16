// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/locator/deployments.proto

package improbable_platform_locator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	_ "improbable.io/proto/improbable/platform"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ListDeploymentsForProjectRequest) Validate() error {
	if this.ProjectId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ProjectId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ProjectId", err)
		}
	}
	return nil
}
func (this *Deployment) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	return nil
}
func (this *ListDeploymentsResponse) Validate() error {
	for _, item := range this.Deployments {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deployments", err)
			}
		}
	}
	return nil
}
func (this *QueueForDeploymentRequest) Validate() error {
	if nil == this.DeploymentId {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", fmt.Errorf("message must exist"))
	}
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	return nil
}
func (this *QueueForDeploymentResponse) Validate() error {
	return nil
}
func (this *ConnectionParams) Validate() error {
	return nil
}
func (this *LoginToDeploymentRequest) Validate() error {
	if nil == this.DeploymentId {
		return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", fmt.Errorf("message must exist"))
	}
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	if nil == this.Client {
		return github_com_mwitkow_go_proto_validators.FieldError("Client", fmt.Errorf("message must exist"))
	}
	if this.Client != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Client); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Client", err)
		}
	}
	if this.ConnectionParams != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ConnectionParams); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ConnectionParams", err)
		}
	}
	return nil
}
func (this *LoginToDeploymentRequest_ClientDescription) Validate() error {
	return nil
}
func (this *DeploymentEndpoint) Validate() error {
	return nil
}
func (this *LoginToDeploymentResponse) Validate() error {
	if this.Connection != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Connection); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Connection", err)
		}
	}
	return nil
}
func (this *LoginRequest) Validate() error {
	return nil
}
func (this *LoginResponse) Validate() error {
	if oneOfNester, ok := this.GetOutcome().(*LoginResponse_Failure); ok {
		if oneOfNester.Failure != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Failure); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Failure", err)
			}
		}
	}
	return nil
}
func (this *LoginResponse_LoginFailure) Validate() error {
	return nil
}