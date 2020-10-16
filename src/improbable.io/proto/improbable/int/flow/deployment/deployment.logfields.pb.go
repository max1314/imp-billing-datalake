// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/flow/deployment/deployment.proto

package improbable_int_flow_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/int/flow"
	_ "improbable.io/proto/improbable/platform/assembly"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DeploymentCreate) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentId)
	hasInner = hasInner || len(deploymentIdFields) > 0
	clusterIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ClusterId)
	hasInner = hasInner || len(clusterIdFields) > 0
	simulationPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SimulationPlaneId)
	hasInner = hasInner || len(simulationPlaneIdFields) > 0
	controlPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ControlPlaneId)
	hasInner = hasInner || len(controlPlaneIdFields) > 0
	deploymentConfigurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentConfiguration)
	hasInner = hasInner || len(deploymentConfigurationFields) > 0
	assemblyIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AssemblyId)
	hasInner = hasInner || len(assemblyIdFields) > 0
	serviceAccountFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ServiceAccount)
	hasInner = hasInner || len(serviceAccountFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"deployment_id": fmt.Sprintf("%v", this.DeploymentUuid),
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentIdFields {
		res[k] = v
	}
	for k, v := range clusterIdFields {
		res[k] = v
	}
	for k, v := range simulationPlaneIdFields {
		res[k] = v
	}
	for k, v := range controlPlaneIdFields {
		res[k] = v
	}
	for k, v := range deploymentConfigurationFields {
		res[k] = v
	}
	for k, v := range assemblyIdFields {
		res[k] = v
	}
	for k, v := range serviceAccountFields {
		res[k] = v
	}
	res["deployment_id"] = fmt.Sprintf("%v", this.DeploymentUuid)
	return res
}

func (this *DeploymentCreate) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ClusterId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SimulationPlaneId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ControlPlaneId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentConfiguration, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AssemblyId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ServiceAccount, dst)
	dst["deployment_id"] = this.DeploymentUuid
}

func (this *DeploymentDelete) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentId)
	hasInner = hasInner || len(deploymentIdFields) > 0
	clusterIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ClusterId)
	hasInner = hasInner || len(clusterIdFields) > 0
	simulationPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SimulationPlaneId)
	hasInner = hasInner || len(simulationPlaneIdFields) > 0
	controlPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ControlPlaneId)
	hasInner = hasInner || len(controlPlaneIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"deployment_id": this.DeploymentUuid,
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentIdFields {
		res[k] = v
	}
	for k, v := range clusterIdFields {
		res[k] = v
	}
	for k, v := range simulationPlaneIdFields {
		res[k] = v
	}
	for k, v := range controlPlaneIdFields {
		res[k] = v
	}
	res["deployment_id"] = this.DeploymentUuid
	return res
}

func (this *DeploymentDelete) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ClusterId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SimulationPlaneId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ControlPlaneId, dst)
	dst["deployment_id"] = this.DeploymentUuid
}

func (this *DeploymentRestart) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	deploymentIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentId)
	hasInner = hasInner || len(deploymentIdFields) > 0
	simulationPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SimulationPlaneId)
	hasInner = hasInner || len(simulationPlaneIdFields) > 0
	controlPlaneIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ControlPlaneId)
	hasInner = hasInner || len(controlPlaneIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{
			"deployment_id": fmt.Sprintf("%v", this.DeploymentUuid),
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentIdFields {
		res[k] = v
	}
	for k, v := range simulationPlaneIdFields {
		res[k] = v
	}
	for k, v := range controlPlaneIdFields {
		res[k] = v
	}
	res["deployment_id"] = fmt.Sprintf("%v", this.DeploymentUuid)
	return res
}

func (this *DeploymentRestart) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SimulationPlaneId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ControlPlaneId, dst)
	dst["deployment_id"] = this.DeploymentUuid
}

func (this *DeploymentDeleteFlowState) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeploymentDeleteFlowState) ExtractRequestFields(dst map[string]interface{}) {
}