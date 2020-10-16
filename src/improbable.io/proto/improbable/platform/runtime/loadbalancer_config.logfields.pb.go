// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/runtime/loadbalancer_config.proto

package improbable_platform_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *WorkerLoadbalancingConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	var workerConfigFields map[string]string
	switch f := this.WorkerConfig.(type) {
	case *WorkerLoadbalancingConfig_StaticHexGrid:
		workerConfigFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.StaticHexGrid)
	case *WorkerLoadbalancingConfig_DynamicLoadbalancer:
		workerConfigFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.DynamicLoadbalancer)
	case *WorkerLoadbalancingConfig_SingletonWorker:
		workerConfigFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.SingletonWorker)
	case *WorkerLoadbalancingConfig_AutoHexGrid:
		workerConfigFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.AutoHexGrid)
	case *WorkerLoadbalancingConfig_PointsOfInterest:
		workerConfigFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.PointsOfInterest)
	default:
		workerConfigFields = map[string]string{}
	}
	hasInner = hasInner || len(workerConfigFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range workerConfigFields {
		res[k] = v
	}
	return res
}

func (this *WorkerLoadbalancingConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	if f, ok := this.WorkerConfig.(*WorkerLoadbalancingConfig_StaticHexGrid); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.StaticHexGrid, dst)
	}
	if f, ok := this.WorkerConfig.(*WorkerLoadbalancingConfig_DynamicLoadbalancer); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.DynamicLoadbalancer, dst)
	}
	if f, ok := this.WorkerConfig.(*WorkerLoadbalancingConfig_SingletonWorker); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.SingletonWorker, dst)
	}
	if f, ok := this.WorkerConfig.(*WorkerLoadbalancingConfig_AutoHexGrid); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.AutoHexGrid, dst)
	}
	if f, ok := this.WorkerConfig.(*WorkerLoadbalancingConfig_PointsOfInterest); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.PointsOfInterest, dst)
	}
}

func (this *LoadbalancingConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LoadbalancingConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SimulationLayerLoadbalancingConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	var strategyFields map[string]string
	switch f := this.Strategy.(type) {
	case *SimulationLayerLoadbalancingConfig_HexGrid:
		strategyFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.HexGrid)
	case *SimulationLayerLoadbalancingConfig_PointsOfInterest:
		strategyFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.PointsOfInterest)
	case *SimulationLayerLoadbalancingConfig_RectangleGrid:
		strategyFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.RectangleGrid)
	case *SimulationLayerLoadbalancingConfig_EntityIdSharding:
		strategyFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.EntityIdSharding)
	default:
		strategyFields = map[string]string{}
	}
	hasInner = hasInner || len(strategyFields) > 0
	optionsFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Options)
	hasInner = hasInner || len(optionsFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range strategyFields {
		res[k] = v
	}
	for k, v := range optionsFields {
		res[k] = v
	}
	return res
}

func (this *SimulationLayerLoadbalancingConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	if f, ok := this.Strategy.(*SimulationLayerLoadbalancingConfig_HexGrid); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.HexGrid, dst)
	}
	if f, ok := this.Strategy.(*SimulationLayerLoadbalancingConfig_PointsOfInterest); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.PointsOfInterest, dst)
	}
	if f, ok := this.Strategy.(*SimulationLayerLoadbalancingConfig_RectangleGrid); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.RectangleGrid, dst)
	}
	if f, ok := this.Strategy.(*SimulationLayerLoadbalancingConfig_EntityIdSharding); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.EntityIdSharding, dst)
	}
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Options, dst)
}

func (this *LoadbalancingOptions) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LoadbalancingOptions) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *StaticHexGrid) LogFields() map[string]string {
	return map[string]string{}
}

func (this *StaticHexGrid) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *AutoHexGrid) LogFields() map[string]string {
	return map[string]string{}
}

func (this *AutoHexGrid) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *StaticPointsOfInterest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *StaticPointsOfInterest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *RectangleGrid) LogFields() map[string]string {
	return map[string]string{}
}

func (this *RectangleGrid) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *EntityIdSharding) LogFields() map[string]string {
	return map[string]string{}
}

func (this *EntityIdSharding) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ConstantWorkerScalerConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ConstantWorkerScalerConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *EntityCountWorkerScalerConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *EntityCountWorkerScalerConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *WorkerScalerConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	var configFields map[string]string
	switch f := this.Config.(type) {
	case *WorkerScalerConfig_ConstantConfig:
		configFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.ConstantConfig)
	case *WorkerScalerConfig_EntityCountConfig:
		configFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.EntityCountConfig)
	default:
		configFields = map[string]string{}
	}
	hasInner = hasInner || len(configFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range configFields {
		res[k] = v
	}
	return res
}

func (this *WorkerScalerConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	if f, ok := this.Config.(*WorkerScalerConfig_ConstantConfig); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.ConstantConfig, dst)
	}
	if f, ok := this.Config.(*WorkerScalerConfig_EntityCountConfig); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.EntityCountConfig, dst)
	}
}

func (this *HexGridWorkerPlacerConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *HexGridWorkerPlacerConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *RandomWorkerPlacerConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *RandomWorkerPlacerConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *AutoHexGridWorkerPlacerConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *AutoHexGridWorkerPlacerConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *WorkerPlacerConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	var configFields map[string]string
	switch f := this.Config.(type) {
	case *WorkerPlacerConfig_HexGridParams:
		configFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.HexGridParams)
	case *WorkerPlacerConfig_RandomParams:
		configFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.RandomParams)
	case *WorkerPlacerConfig_AutoHexGridParams:
		configFields = github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(f.AutoHexGridParams)
	default:
		configFields = map[string]string{}
	}
	hasInner = hasInner || len(configFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range configFields {
		res[k] = v
	}
	return res
}

func (this *WorkerPlacerConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	if f, ok := this.Config.(*WorkerPlacerConfig_HexGridParams); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.HexGridParams, dst)
	}
	if f, ok := this.Config.(*WorkerPlacerConfig_RandomParams); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.RandomParams, dst)
	}
	if f, ok := this.Config.(*WorkerPlacerConfig_AutoHexGridParams); ok {
		github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(f.AutoHexGridParams, dst)
	}
}

func (this *LoadbalancerConfig) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LoadbalancerConfig) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DynamicLoadbalancerConfig) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	workerScalerConfigFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.WorkerScalerConfig)
	hasInner = hasInner || len(workerScalerConfigFields) > 0
	workerPlacerConfigFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.WorkerPlacerConfig)
	hasInner = hasInner || len(workerPlacerConfigFields) > 0
	loadbalancerConfigFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.LoadbalancerConfig)
	hasInner = hasInner || len(loadbalancerConfigFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range workerScalerConfigFields {
		res[k] = v
	}
	for k, v := range workerPlacerConfigFields {
		res[k] = v
	}
	for k, v := range loadbalancerConfigFields {
		res[k] = v
	}
	return res
}

func (this *DynamicLoadbalancerConfig) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.WorkerScalerConfig, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.WorkerPlacerConfig, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.LoadbalancerConfig, dst)
}

func (this *SingletonWorker) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SingletonWorker) ExtractRequestFields(dst map[string]interface{}) {
}