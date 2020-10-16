// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/runtime/receptionist.proto

package improbable_platform_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RakNetLinkSettings) Validate() error {
	return nil
}
func (this *TcpLinkSettings) Validate() error {
	return nil
}
func (this *ErasureCodecParameters) Validate() error {
	return nil
}
func (this *HeartbeatParameters) Validate() error {
	return nil
}
func (this *KcpLinkSettings) Validate() error {
	if this.ErasureCodec != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ErasureCodec); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ErasureCodec", err)
		}
	}
	if this.Heartbeat != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Heartbeat); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Heartbeat", err)
		}
	}
	return nil
}
func (this *LinkSettings) Validate() error {
	if oneOfNester, ok := this.GetLinkSettings().(*LinkSettings_RaknetSettings); ok {
		if oneOfNester.RaknetSettings != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RaknetSettings); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RaknetSettings", err)
			}
		}
	}
	if oneOfNester, ok := this.GetLinkSettings().(*LinkSettings_TcpSettings); ok {
		if oneOfNester.TcpSettings != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.TcpSettings); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TcpSettings", err)
			}
		}
	}
	if oneOfNester, ok := this.GetLinkSettings().(*LinkSettings_KcpSettings); ok {
		if oneOfNester.KcpSettings != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.KcpSettings); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("KcpSettings", err)
			}
		}
	}
	return nil
}
func (this *ProtocolCapabilitiesV1) Validate() error {
	if this.LinkSettings != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LinkSettings); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LinkSettings", err)
		}
	}
	return nil
}
func (this *ProtocolCapabilitiesV2) Validate() error {
	return nil
}
func (this *ProtocolCapabilities) Validate() error {
	if this.CapabilitiesV1 != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CapabilitiesV1); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CapabilitiesV1", err)
		}
	}
	if this.CapabilitiesV2 != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CapabilitiesV2); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CapabilitiesV2", err)
		}
	}
	return nil
}
func (this *ReceptionistLoginPayload) Validate() error {
	if this.LinkSettings != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LinkSettings); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LinkSettings", err)
		}
	}
	return nil
}
func (this *LoginRequest) Validate() error {
	return nil
}
func (this *BridgeSession) Validate() error {
	return nil
}
func (this *BridgeConnectionDetails) Validate() error {
	if this.Session != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Session); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Session", err)
		}
	}
	return nil
}
func (this *ProtocolSettingsV1) Validate() error {
	if this.BridgeDetails != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BridgeDetails); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BridgeDetails", err)
		}
	}
	return nil
}
func (this *ProtocolSettingsV2) Validate() error {
	if this.BridgeDetails != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BridgeDetails); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BridgeDetails", err)
		}
	}
	return nil
}
func (this *ProtocolSettings) Validate() error {
	if oneOfNester, ok := this.GetSettings().(*ProtocolSettings_ProtocolSettingsV1); ok {
		if oneOfNester.ProtocolSettingsV1 != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ProtocolSettingsV1); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ProtocolSettingsV1", err)
			}
		}
	}
	if oneOfNester, ok := this.GetSettings().(*ProtocolSettings_ProtocolSettingsV2); ok {
		if oneOfNester.ProtocolSettingsV2 != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.ProtocolSettingsV2); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ProtocolSettingsV2", err)
			}
		}
	}
	return nil
}
func (this *LoginResponse) Validate() error {
	if this.BridgeDetails != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BridgeDetails); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BridgeDetails", err)
		}
	}
	return nil
}
func (this *RemainingCapacityRequest) Validate() error {
	return nil
}
func (this *RemainingCapacityResponse) Validate() error {
	return nil
}
func (this *RemainingCapacity) Validate() error {
	return nil
}
func (this *RateLimit) Validate() error {
	if this.Period != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Period); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Period", err)
		}
	}
	return nil
}
func (this *WorkerConfiguration) Validate() error {
	if this.LoginRate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LoginRate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LoginRate", err)
		}
	}
	return nil
}
func (this *LoginV2Configuration) Validate() error {
	for _, item := range this.WorkerConfigs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkerConfigs", err)
			}
		}
	}
	return nil
}
func (this *SetLoginRateRequest) Validate() error {
	if this.RateLimit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RateLimit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RateLimit", err)
		}
	}
	return nil
}
func (this *SetLoginRateResponse) Validate() error {
	return nil
}
func (this *SetMaximumCapacityRequest) Validate() error {
	return nil
}
func (this *SetMaximumCapacityResponse) Validate() error {
	return nil
}
func (this *ListWorkerConfigurationRequest) Validate() error {
	return nil
}
func (this *ListWorkerConfigurationResponse) Validate() error {
	for _, item := range this.WorkerConfigs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkerConfigs", err)
			}
		}
	}
	return nil
}
func (this *RemainingCapacityV2Request) Validate() error {
	return nil
}
func (this *RemainingCapacityV2Response) Validate() error {
	for _, item := range this.RemainingCapacity {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RemainingCapacity", err)
			}
		}
	}
	for _, item := range this.WorkerConfigs {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WorkerConfigs", err)
			}
		}
	}
	return nil
}
func (this *PlayerIdentity) Validate() error {
	return nil
}
func (this *LoginV2Request) Validate() error {
	if this.Identity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Identity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Identity", err)
		}
	}
	return nil
}
func (this *LoginV2Response) Validate() error {
	if this.RemainingCapacity != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RemainingCapacity); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RemainingCapacity", err)
		}
	}
	return nil
}
