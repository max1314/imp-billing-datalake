// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/ext/platform/deployment/deployment.proto

package improbable_ext_platform_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/platform"
	_ "improbable.io/proto/improbable/platform/assembly"
	_ "improbable.io/proto/improbable/platform/deployment"
	_ "improbable.io/proto/improbable/platform/filter"
	_ "improbable.io/proto/improbable/platform/history"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DeploymentCreateRequest) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	if this.ClusterId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ClusterId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ClusterId", err)
		}
	}
	if nil == this.AssemblyId {
		return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", fmt.Errorf("message must exist"))
	}
	if this.AssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	if this.LaunchConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LaunchConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LaunchConfig", err)
		}
	}
	if this.StartingSnapshotId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartingSnapshotId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartingSnapshotId", err)
		}
	}
	return nil
}
func (this *DeploymentCreateResponse) Validate() error {
	return nil
}
func (this *DeploymentDeleteRequest) Validate() error {
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
func (this *DeploymentDeleteResponse) Validate() error {
	return nil
}
func (this *DeploymentRestartRequest) Validate() error {
	return nil
}
func (this *DeploymentRestartResponse) Validate() error {
	return nil
}
func (this *CreateDeploymentRequest) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	if this.ClusterId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ClusterId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ClusterId", err)
		}
	}
	if nil == this.AssemblyId {
		return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", fmt.Errorf("message must exist"))
	}
	if this.AssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", err)
		}
	}
	if this.Config != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Config); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Config", err)
		}
	}
	if this.LaunchConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LaunchConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LaunchConfig", err)
		}
	}
	if this.StartingSnapshotId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartingSnapshotId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartingSnapshotId", err)
		}
	}
	return nil
}
func (this *CreateDeploymentResponse) Validate() error {
	if this.CreateOperationId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreateOperationId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreateOperationId", err)
		}
	}
	return nil
}
func (this *DeleteDeploymentRequest) Validate() error {
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
func (this *DeleteDeploymentResponse) Validate() error {
	if this.DeleteOperationId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeleteOperationId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeleteOperationId", err)
		}
	}
	return nil
}
func (this *Deployment) Validate() error {
	if this.LegacyDeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LegacyDeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LegacyDeploymentId", err)
		}
	}
	if this.ClusterId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ClusterId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ClusterId", err)
		}
	}
	if this.CreationTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreationTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreationTime", err)
		}
	}
	if this.StartingSnapshotId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartingSnapshotId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartingSnapshotId", err)
		}
	}
	if this.StartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTime", err)
		}
	}
	if this.StopTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StopTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StopTime", err)
		}
	}
	if this.ExpiryTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExpiryTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExpiryTime", err)
		}
	}
	return nil
}
func (this *GetDeploymentRequest) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	return nil
}
func (this *GetDeploymentResponse) Validate() error {
	if this.Deployment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deployment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deployment", err)
		}
	}
	return nil
}
func (this *ListDeploymentsRequest) Validate() error {
	if this.ProjectId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ProjectId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ProjectId", err)
		}
	}
	for _, item := range this.Filters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Filters", err)
			}
		}
	}
	return nil
}
func (this *ListDeploymentsResponse) Validate() error {
	for _, item := range this.Deployment {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deployment", err)
			}
		}
	}
	return nil
}
func (this *UpdateDeploymentTagsRequest) Validate() error {
	return nil
}
func (this *UpdateDeploymentTagsResponse) Validate() error {
	return nil
}
func (this *UpdateDeploymentRequest) Validate() error {
	if this.ProjectId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ProjectId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ProjectId", err)
		}
	}
	return nil
}
func (this *UpdateDeploymentResponse) Validate() error {
	if this.Deployment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deployment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deployment", err)
		}
	}
	return nil
}
func (this *StopDeploymentRequest) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	return nil
}
func (this *StopDeploymentResponse) Validate() error {
	return nil
}