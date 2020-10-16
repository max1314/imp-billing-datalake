// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/scheduler/machines.proto

package improbable_int_platform_scheduler

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

func (this *MachineState) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.AssignedJob != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssignedJob); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssignedJob", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *MachineReconcileState) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.LastSeenError != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastSeenError); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastSeenError", err)
		}
	}
	return nil
}
func (this *MachineReconcileError) Validate() error {
	return nil
}
func (this *GetMachineStateRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *GetMachineStateResponse) Validate() error {
	if this.State != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.State); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("State", err)
		}
	}
	return nil
}
func (this *MintMachineForJobRequest) Validate() error {
	if this.JobId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.JobId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("JobId", err)
		}
	}
	return nil
}
func (this *MintMachineForJobResponse) Validate() error {
	if this.MachineId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MachineId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MachineId", err)
		}
	}
	return nil
}
func (this *RevokeMachineRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *RevokeMachineResponse) Validate() error {
	return nil
}
func (this *ClearMachineRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *ClearMachineResponse) Validate() error {
	return nil
}
func (this *WatchAssignedJobRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *WatchAssignedJobResponse) Validate() error {
	if this.JobId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.JobId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("JobId", err)
		}
	}
	if this.SimulationPlaneId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SimulationPlaneId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SimulationPlaneId", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	if this.AssignedMachine != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssignedMachine); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssignedMachine", err)
		}
	}
	if this.DockerParameters != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DockerParameters); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DockerParameters", err)
		}
	}
	return nil
}
func (this *ListMachinesRequest) Validate() error {
	return nil
}
func (this *ListMachinesResponse) Validate() error {
	for _, item := range this.Machines {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Machines", err)
			}
		}
	}
	return nil
}
func (this *FreezeMachineRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *FreezeMachineResponse) Validate() error {
	return nil
}
func (this *UnfreezeMachineRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *UnfreezeMachineResponse) Validate() error {
	return nil
}
func (this *UpdateMachineReconcileStateRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Error != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Error); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Error", err)
		}
	}
	return nil
}
func (this *UpdateMachineReconcileStateResponse) Validate() error {
	if this.State != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.State); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("State", err)
		}
	}
	return nil
}
func (this *GetMachineReconcileStateRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *GetMachineReconcileStateResponse) Validate() error {
	if this.State != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.State); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("State", err)
		}
	}
	return nil
}