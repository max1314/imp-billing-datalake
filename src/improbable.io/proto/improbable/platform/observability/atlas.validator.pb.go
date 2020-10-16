// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/observability/atlas.proto

package improbable_platform_observability

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

func (this *LogEntry) Validate() error {
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *LogEntry_Value) Validate() error {
	return nil
}
func (this *Index) Validate() error {
	if this.LastCompactedMinute != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastCompactedMinute); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastCompactedMinute", err)
		}
	}
	for _, item := range this.Records {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Records", err)
			}
		}
	}
	return nil
}
func (this *Index_Record) Validate() error {
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}