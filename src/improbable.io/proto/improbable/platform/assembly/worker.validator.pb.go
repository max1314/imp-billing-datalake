// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/assembly/worker.proto

package improbable_platform_assembly

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_WorkerName_Name = regexp.MustCompile(`^[a-zA-Z0-9_-]{3,128}$`)

func (this *WorkerName) Validate() error {
	if !_regex_WorkerName_Name.MatchString(this.Name) {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9_-]{3,128}$"`, this.Name))
	}
	return nil
}
func (this *WorkerAssemblyId) Validate() error {
	if nil == this.AssemblyId {
		return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", fmt.Errorf("message must exist"))
	}
	if this.AssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", err)
		}
	}
	if nil == this.Name {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf("message must exist"))
	}
	if this.Name != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Name); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Name", err)
		}
	}
	return nil
}
func (this *WorkerAssembly) Validate() error {
	if nil == this.WorkerAssemblyId {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkerAssemblyId", fmt.Errorf("message must exist"))
	}
	if this.WorkerAssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerAssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerAssemblyId", err)
		}
	}
	for _, item := range this.WorkerLaunchConfiguration {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkerLaunchConfiguration", err)
			}
		}
	}
	if this.BridgeConfiguration != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BridgeConfiguration); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BridgeConfiguration", err)
		}
	}
	if this.WorkerConfigJson != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerConfigJson); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerConfigJson", err)
		}
	}
	return nil
}
func (this *WorkerConfigJson) Validate() error {
	return nil
}
func (this *ListWorkerAssembliesRequest) Validate() error {
	if nil == this.AssemblyId {
		return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", fmt.Errorf("message must exist"))
	}
	if this.AssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AssemblyId", err)
		}
	}
	return nil
}
func (this *ListWorkerAssembliesResponse) Validate() error {
	for _, item := range this.Worker {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Worker", err)
			}
		}
	}
	return nil
}
func (this *GetWorkerAssemblyRequest) Validate() error {
	if nil == this.WorkerAssemblyId {
		return github_com_mwitkow_go_proto_validators.FieldError("WorkerAssemblyId", fmt.Errorf("message must exist"))
	}
	if this.WorkerAssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerAssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerAssemblyId", err)
		}
	}
	return nil
}
func (this *GetWorkerAssemblyResponse) Validate() error {
	if this.Worker != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Worker); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Worker", err)
		}
	}
	return nil
}
func (this *SetWorkerAssemblyRequest) Validate() error {
	if this.Worker != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Worker); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Worker", err)
		}
	}
	for _, item := range this.Workers {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Workers", err)
			}
		}
	}
	return nil
}
func (this *SetWorkerAssemblyResponse) Validate() error {
	return nil
}
func (this *DeleteWorkerAssemblyRequest) Validate() error {
	if this.WorkerAssemblyId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerAssemblyId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerAssemblyId", err)
		}
	}
	return nil
}
func (this *DeleteWorkerAssemblyResponse) Validate() error {
	return nil
}
