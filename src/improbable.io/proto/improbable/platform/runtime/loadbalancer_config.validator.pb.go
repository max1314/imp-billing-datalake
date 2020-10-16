// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/runtime/loadbalancer_config.proto

package improbable_platform_runtime

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

func (this *WorkerLoadbalancingConfig) Validate() error {
	if oneOfNester, ok := this.GetWorkerConfig().(*WorkerLoadbalancingConfig_StaticHexGrid); ok {
		if oneOfNester.StaticHexGrid != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StaticHexGrid); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StaticHexGrid", err)
			}
		}
	}
	if oneOfNester, ok := this.GetWorkerConfig().(*WorkerLoadbalancingConfig_DynamicLoadbalancer); ok {
		if oneOfNester.DynamicLoadbalancer != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.DynamicLoadbalancer); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("DynamicLoadbalancer", err)
			}
		}
	}
	if oneOfNester, ok := this.GetWorkerConfig().(*WorkerLoadbalancingConfig_SingletonWorker); ok {
		if oneOfNester.SingletonWorker != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SingletonWorker); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SingletonWorker", err)
			}
		}
	}
	if oneOfNester, ok := this.GetWorkerConfig().(*WorkerLoadbalancingConfig_AutoHexGrid); ok {
		if oneOfNester.AutoHexGrid != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.AutoHexGrid); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AutoHexGrid", err)
			}
		}
	}
	if oneOfNester, ok := this.GetWorkerConfig().(*WorkerLoadbalancingConfig_PointsOfInterest); ok {
		if oneOfNester.PointsOfInterest != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PointsOfInterest); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PointsOfInterest", err)
			}
		}
	}
	return nil
}
func (this *LoadbalancingConfig) Validate() error {
	for _, item := range this.LayerConfigurations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LayerConfigurations", err)
			}
		}
	}
	return nil
}
func (this *SimulationLayerLoadbalancingConfig) Validate() error {
	if oneOfNester, ok := this.GetStrategy().(*SimulationLayerLoadbalancingConfig_HexGrid); ok {
		if oneOfNester.HexGrid != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.HexGrid); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HexGrid", err)
			}
		}
	}
	if oneOfNester, ok := this.GetStrategy().(*SimulationLayerLoadbalancingConfig_PointsOfInterest); ok {
		if oneOfNester.PointsOfInterest != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PointsOfInterest); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PointsOfInterest", err)
			}
		}
	}
	if oneOfNester, ok := this.GetStrategy().(*SimulationLayerLoadbalancingConfig_RectangleGrid); ok {
		if oneOfNester.RectangleGrid != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RectangleGrid); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RectangleGrid", err)
			}
		}
	}
	if oneOfNester, ok := this.GetStrategy().(*SimulationLayerLoadbalancingConfig_EntityIdSharding); ok {
		if oneOfNester.EntityIdSharding != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.EntityIdSharding); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EntityIdSharding", err)
			}
		}
	}
	if this.Options != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Options); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Options", err)
		}
	}
	return nil
}
func (this *LoadbalancingOptions) Validate() error {
	return nil
}
func (this *StaticHexGrid) Validate() error {
	return nil
}
func (this *AutoHexGrid) Validate() error {
	return nil
}
func (this *StaticPointsOfInterest) Validate() error {
	for _, item := range this.Points {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Points", err)
			}
		}
	}
	return nil
}
func (this *RectangleGrid) Validate() error {
	return nil
}
func (this *EntityIdSharding) Validate() error {
	return nil
}
func (this *ConstantWorkerScalerConfig) Validate() error {
	return nil
}
func (this *EntityCountWorkerScalerConfig) Validate() error {
	return nil
}
func (this *WorkerScalerConfig) Validate() error {
	if oneOfNester, ok := this.GetConfig().(*WorkerScalerConfig_ConstantConfig); ok {
		if oneOfNester.ConstantConfig != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ConstantConfig); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ConstantConfig", err)
			}
		}
	}
	if oneOfNester, ok := this.GetConfig().(*WorkerScalerConfig_EntityCountConfig); ok {
		if oneOfNester.EntityCountConfig != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.EntityCountConfig); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EntityCountConfig", err)
			}
		}
	}
	return nil
}
func (this *HexGridWorkerPlacerConfig) Validate() error {
	return nil
}
func (this *RandomWorkerPlacerConfig) Validate() error {
	return nil
}
func (this *AutoHexGridWorkerPlacerConfig) Validate() error {
	return nil
}
func (this *WorkerPlacerConfig) Validate() error {
	if oneOfNester, ok := this.GetConfig().(*WorkerPlacerConfig_HexGridParams); ok {
		if oneOfNester.HexGridParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.HexGridParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("HexGridParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetConfig().(*WorkerPlacerConfig_RandomParams); ok {
		if oneOfNester.RandomParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RandomParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RandomParams", err)
			}
		}
	}
	if oneOfNester, ok := this.GetConfig().(*WorkerPlacerConfig_AutoHexGridParams); ok {
		if oneOfNester.AutoHexGridParams != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.AutoHexGridParams); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AutoHexGridParams", err)
			}
		}
	}
	return nil
}
func (this *LoadbalancerConfig) Validate() error {
	return nil
}
func (this *DynamicLoadbalancerConfig) Validate() error {
	if this.WorkerScalerConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerScalerConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerScalerConfig", err)
		}
	}
	if this.WorkerPlacerConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WorkerPlacerConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WorkerPlacerConfig", err)
		}
	}
	if this.LoadbalancerConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LoadbalancerConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LoadbalancerConfig", err)
		}
	}
	return nil
}
func (this *SingletonWorker) Validate() error {
	return nil
}
