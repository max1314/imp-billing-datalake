// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/runtime/worker_attribute.proto

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

func (this *WorkerAttribute) Validate() error {
	return nil
}
func (this *WorkerAttributeSet) Validate() error {
	for _, item := range this.Attribute {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Attribute", err)
			}
		}
	}
	return nil
}