// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/scheduler/common.proto

package improbable_int_platform_scheduler

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *SimulationPlaneId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"simulation_plane": this.Name,
	}
}

func (this *SimulationPlaneId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["simulation_plane"] = this.Name
}

func (this *DeploymentId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"deployment_name": this.Name,
	}
}

func (this *DeploymentId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["deployment_name"] = this.Name
}

func (this *JobId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentId)
	hasInner = hasInner || len(deploymentIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentIdFields {
		res[k] = v
	}
	return res
}

func (this *JobId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
}

func (this *MachineId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	simulationPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SimulationPlaneId)
	hasInner = hasInner || len(simulationPlaneIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"machine_name": this.Name,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	res["machine_name"] = this.Name
	for k, v := range simulationPlaneIdFields {
		res[k] = v
	}
	return res
}

func (this *MachineId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["machine_name"] = this.Name
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SimulationPlaneId, dst)
}

func (this *DockerParameters) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DockerParameters) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *JobParameters) LogFields() map[string]string {
	return map[string]string{}
}

func (this *JobParameters) ExtractRequestFields(dst map[string]interface{}) {
}