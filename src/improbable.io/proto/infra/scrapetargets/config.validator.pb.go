// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/scrapetargets/config.proto

package improbable_infra_scrapetargets

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

func (this *ScrapeTargets) Validate() error {
	for _, item := range this.Targets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Targets", err)
			}
		}
	}
	return nil
}
func (this *Target) Validate() error {
	return nil
}
