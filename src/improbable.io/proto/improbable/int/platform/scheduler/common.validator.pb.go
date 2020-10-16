// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/scheduler/common.proto

package improbable_int_platform_scheduler

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SimulationPlaneId) Validate() error {
	return nil
}
func (this *DeploymentId) Validate() error {
	return nil
}
func (this *JobId) Validate() error {
	if this.DeploymentId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeploymentId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeploymentId", err)
		}
	}
	return nil
}
func (this *MachineId) Validate() error {
	if this.SimulationPlaneId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SimulationPlaneId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SimulationPlaneId", err)
		}
	}
	return nil
}
func (this *DockerParameters) Validate() error {
	return nil
}
func (this *JobParameters) Validate() error {
	return nil
}
