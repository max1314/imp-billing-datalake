// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/deployment/worker_node.proto

package improbable_platform_deployment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "improbable.io/proto/improbable/platform"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *WorkerNodeId) LogFields() map[string]string {
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

func (this *WorkerNodeId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
}

func (this *WorkerNode) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	workerNodeIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.WorkerNodeId)
	hasInner = hasInner || len(workerNodeIdFields) > 0
	lastMachineErrorFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.LastMachineError)
	hasInner = hasInner || len(lastMachineErrorFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range workerNodeIdFields {
		res[k] = v
	}
	for k, v := range lastMachineErrorFields {
		res[k] = v
	}
	return res
}

func (this *WorkerNode) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.WorkerNodeId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LastMachineError, dst)
}

func (this *CreateMultiWorkerNodeRequest) LogFields() map[string]string {
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
		return map[string]string{
			"deployment_id": fmt.Sprintf("%v", this.DeploymentUuid),
		}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range deploymentIdFields {
		res[k] = v
	}
	res["deployment_id"] = fmt.Sprintf("%v", this.DeploymentUuid)
	return res
}

func (this *CreateMultiWorkerNodeRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
	dst["deployment_id"] = this.DeploymentUuid
}

func (this *CreateMultiWorkerNodeResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateMultiWorkerNodeResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListWorkerNodesRequest) LogFields() map[string]string {
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

func (this *ListWorkerNodesRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
}

func (this *ListWorkerNodesResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListWorkerNodesResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SetAllWorkerNodeStateRequest) LogFields() map[string]string {
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

func (this *SetAllWorkerNodeStateRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
}

func (this *SetAllWorkerNodeStateResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SetAllWorkerNodeStateResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetWorkerNodeCredsRequest) LogFields() map[string]string {
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

func (this *GetWorkerNodeCredsRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentId, dst)
}

func (this *GetWorkerNodeCredsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *GetWorkerNodeCredsResponse) ExtractRequestFields(dst map[string]interface{}) {
}
