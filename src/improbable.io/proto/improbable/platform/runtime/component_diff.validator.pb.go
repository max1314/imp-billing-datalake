// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/runtime/component_diff.proto

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

func (this *ComponentDiff) Validate() error {
	return nil
}
func (this *Encoding) Validate() error {
	return nil
}
func (this *EncodingStrategyToFormatEntry) Validate() error {
	return nil
}
func (this *EncodingStrategyToFormat) Validate() error {
	for _, item := range this.StrategyToFormat {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StrategyToFormat", err)
			}
		}
	}
	return nil
}
func (this *ComponentDiffFormats) Validate() error {
	return nil
}
func (this *ComponentDiffSettings) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *FieldDiffs) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}