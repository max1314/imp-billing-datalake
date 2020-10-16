// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/assembly/worker_launch_config.proto

package improbable_platform_assembly

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *WorkerZip) Validate() error {
	if this.Zip != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Zip); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Zip", err)
		}
	}
	if this.WorkerLocation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerLocation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerLocation", err)
		}
	}
	if this.ExposedPorts != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ExposedPorts); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ExposedPorts", err)
		}
	}
	if this.ContainerConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ContainerConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ContainerConfig", err)
		}
	}
	return nil
}
func (this *WorkerZip_WorkerLocation) Validate() error {
	return nil
}
func (this *ContainerConfig) Validate() error {
	return nil
}
func (this *ExposedPort) Validate() error {
	return nil
}
func (this *WorkerLaunchConfiguration) Validate() error {
	if oneOfNester, ok := this.GetWorkerLaunch().(*WorkerLaunchConfiguration_WorkerZip); ok {
		if oneOfNester.WorkerZip != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.WorkerZip); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkerZip", err)
			}
		}
	}
	return nil
}
func (this *WorkerLaunchDescriptor) Validate() error {
	for _, item := range this.LaunchConfig {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LaunchConfig", err)
			}
		}
	}
	return nil
}