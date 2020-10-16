// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/fabric/devtools/service/release_tool/fabric_version.proto

package improbable_fabric_devtools_service_release_tool

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *BuildStatus) Validate() error {
	if this.RunStartTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.RunStartTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("RunStartTime", err)
		}
	}
	return nil
}
func (this *FabricVersion) Validate() error {
	if nil == this.BuildStatus {
		return github_com_mwitkow_go_proto_validators.FieldError("BuildStatus", fmt.Errorf("message must exist"))
	}
	if this.BuildStatus != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BuildStatus); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BuildStatus", err)
		}
	}
	return nil
}
