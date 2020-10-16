// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/audit/event.proto

package improbable_platform_audit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Event) Validate() error {
	if this.ResponseTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.ResponseTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("ResponseTimestamp", err)
		}
	}
	return nil
}
