// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/toolshare/toolshare.proto

package improbable_int_toolshare

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Tool) Validate() error {
	if this.Metadata != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Metadata); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Metadata", err)
		}
	}
	return nil
}

var _regex_Tool_Metadata_Owner = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)

func (this *Tool_Metadata) Validate() error {
	if !_regex_Tool_Metadata_Owner.MatchString(this.Owner) {
		return github_com_mwitkow_go_proto_validators.FieldError("Owner", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"`, this.Owner))
	}
	return nil
}
func (this *ToolshareConfiguration) Validate() error {
	for _, item := range this.PinnedTool {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PinnedTool", err)
			}
		}
	}
	return nil
}
func (this *ToolshareConfiguration_ToolVersion) Validate() error {
	return nil
}
func (this *LogFile) Validate() error {
	if this.Log != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Log); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Log", err)
		}
	}
	return nil
}
func (this *LogFile_Log) Validate() error {
	if this.LastUse != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastUse); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastUse", err)
		}
	}
	return nil
}
