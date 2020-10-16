// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/fabric/regisseur/scenario.proto

package improbable_regisseur

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "improbable.io/proto/improbable/inner/fabric_bundle"
	_ "improbable.io/proto/improbable/platform"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Scenario) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	identifierFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Identifier)
	hasInner = hasInner || len(identifierFields) > 0
	projectIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ProjectId)
	hasInner = hasInner || len(projectIdFields) > 0
	assemblyFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Assembly)
	hasInner = hasInner || len(assemblyFields) > 0
	deploymentConfigurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.DeploymentConfiguration)
	hasInner = hasInner || len(deploymentConfigurationFields) > 0
	lifetimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Lifetime)
	hasInner = hasInner || len(lifetimeFields) > 0
	fabricVersionFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.FabricVersion)
	hasInner = hasInner || len(fabricVersionFields) > 0
	initialStateFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.InitialState)
	hasInner = hasInner || len(initialStateFields) > 0
	launchConfigurationFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.LaunchConfiguration)
	hasInner = hasInner || len(launchConfigurationFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range identifierFields {
		res[k] = v
	}
	for k, v := range projectIdFields {
		res[k] = v
	}
	for k, v := range assemblyFields {
		res[k] = v
	}
	for k, v := range deploymentConfigurationFields {
		res[k] = v
	}
	for k, v := range lifetimeFields {
		res[k] = v
	}
	for k, v := range fabricVersionFields {
		res[k] = v
	}
	for k, v := range initialStateFields {
		res[k] = v
	}
	for k, v := range launchConfigurationFields {
		res[k] = v
	}
	return res
}

func (this *Scenario) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Identifier, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Assembly, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.DeploymentConfiguration, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Lifetime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.FabricVersion, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.InitialState, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LaunchConfiguration, dst)
}

func (this *Scenario_Lifetime) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Scenario_Lifetime) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *InitialState) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	snapshotIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.SnapshotId)
	hasInner = hasInner || len(snapshotIdFields) > 0
	taggedSnapshotFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.TaggedSnapshot)
	hasInner = hasInner || len(taggedSnapshotFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range snapshotIdFields {
		res[k] = v
	}
	for k, v := range taggedSnapshotFields {
		res[k] = v
	}
	return res
}

func (this *InitialState) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.SnapshotId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.TaggedSnapshot, dst)
}

func (this *TaggedSnapshot) LogFields() map[string]string {
	return map[string]string{}
}

func (this *TaggedSnapshot) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SnapshotId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	historyIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.HistoryId)
	hasInner = hasInner || len(historyIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range historyIdFields {
		res[k] = v
	}
	return res
}

func (this *SnapshotId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.HistoryId, dst)
}

func (this *HistoryId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	projectIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ProjectId)
	hasInner = hasInner || len(projectIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range projectIdFields {
		res[k] = v
	}
	return res
}

func (this *HistoryId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectId, dst)
}

func (this *ScenarioIdentifier) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ScenarioIdentifier) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ScenarioTag) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ScenarioTag) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Assembly) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Assembly) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Event) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	var eventTypeFields map[string]string
	switch f := this.EventType.(type) {
	case *Event_WorkerFlagUpdate:
		eventTypeFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.WorkerFlagUpdate)
	case *Event_ProfileRuntime:
		eventTypeFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.ProfileRuntime)
	default:
		eventTypeFields = map[string]string{}
	}
	hasInner = hasInner || len(eventTypeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range eventTypeFields {
		res[k] = v
	}
	return res
}

func (this *Event) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	if f, ok := this.EventType.(*Event_WorkerFlagUpdate); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.WorkerFlagUpdate, dst)
	}
	if f, ok := this.EventType.(*Event_ProfileRuntime); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.ProfileRuntime, dst)
	}
}

func (this *UpdateWorkerFlagEvent) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	workerTypeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.WorkerType)
	hasInner = hasInner || len(workerTypeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range workerTypeFields {
		res[k] = v
	}
	return res
}

func (this *UpdateWorkerFlagEvent) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.WorkerType, dst)
}

func (this *UpdateWorkerFlagEvent_WorkerFlagUpdate) LogFields() map[string]string {
	return map[string]string{}
}

func (this *UpdateWorkerFlagEvent_WorkerFlagUpdate) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *AsyncProfilerEvent) LogFields() map[string]string {
	return map[string]string{}
}

func (this *AsyncProfilerEvent) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Metric) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	fallbackFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Fallback)
	hasInner = hasInner || len(fallbackFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range fallbackFields {
		res[k] = v
	}
	return res
}

func (this *Metric) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Fallback, dst)
}

func (this *Metric_Fallback) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	valueFallbackFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ValueFallback)
	hasInner = hasInner || len(valueFallbackFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range valueFallbackFields {
		res[k] = v
	}
	return res
}

func (this *Metric_Fallback) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ValueFallback, dst)
}
