// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/ext/platform/deployment/operations.proto

package improbable_ext_platform_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *OperationError) Validate() error {
	return nil
}
func (this *CreateOperationId) Validate() error {
	return nil
}
func (this *CreateDeploymentOperationStatusRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *CreateDeploymentOperationStatusResponse) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if oneOfNester, ok := this.GetResult().(*CreateDeploymentOperationStatusResponse_Error); ok {
		if oneOfNester.Error != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Error); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*CreateDeploymentOperationStatusResponse_Response); ok {
		if oneOfNester.Response != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Response); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
			}
		}
	}
	return nil
}
func (this *CreateDeploymentOperationStatusResponse_OperationResponse) Validate() error {
	return nil
}
func (this *DeleteOperationId) Validate() error {
	return nil
}
func (this *DeleteDeploymentOperationStatusRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *DeleteDeploymentOperationStatusResponse) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if oneOfNester, ok := this.GetResult().(*DeleteDeploymentOperationStatusResponse_Error); ok {
		if oneOfNester.Error != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Error); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*DeleteDeploymentOperationStatusResponse_Response); ok {
		if oneOfNester.Response != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Response); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
			}
		}
	}
	return nil
}
func (this *DeleteDeploymentOperationStatusResponse_OperationResponse) Validate() error {
	return nil
}
func (this *RestartOperationId) Validate() error {
	return nil
}
func (this *RestartDeploymentOperationStatusRequest) Validate() error {
	if nil == this.Id {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf("message must exist"))
	}
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *RestartDeploymentOperationStatusResponse) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if oneOfNester, ok := this.GetResult().(*RestartDeploymentOperationStatusResponse_Error); ok {
		if oneOfNester.Error != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Error); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
			}
		}
	}
	if oneOfNester, ok := this.GetResult().(*RestartDeploymentOperationStatusResponse_Response); ok {
		if oneOfNester.Response != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Response); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
			}
		}
	}
	return nil
}
func (this *RestartDeploymentOperationStatusResponse_OperationResponse) Validate() error {
	return nil
}
