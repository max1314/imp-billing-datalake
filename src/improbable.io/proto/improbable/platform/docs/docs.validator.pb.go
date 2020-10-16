// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/docs/docs.proto

package improbable_platform_docs

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *VersionInfo) Validate() error {
	return nil
}
func (this *VersionsResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *YankResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *TagResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CrossLink) Validate() error {
	return nil
}
func (this *VersionsRequest) Validate() error {
	return nil
}
func (this *UploadRequest) Validate() error {
	return nil
}
func (this *UploadResponse) Validate() error {
	return nil
}
func (this *AlgoliaConfigRequest) Validate() error {
	return nil
}
func (this *AlgoliaConfigResponse) Validate() error {
	return nil
}
func (this *YankRequest) Validate() error {
	return nil
}
func (this *TagRequest) Validate() error {
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
func (this *SdkMapRequest) Validate() error {
	return nil
}
func (this *SdkMapResponse) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *CheckCrossLinksRequest) Validate() error {
	for _, item := range this.CrossLinks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CrossLinks", err)
			}
		}
	}
	return nil
}
func (this *CheckCrossLinksResponse) Validate() error {
	return nil
}
func (this *GetCrossLinksRequest) Validate() error {
	return nil
}
func (this *CrossLinksSet) Validate() error {
	for _, item := range this.CrossLinks {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CrossLinks", err)
			}
		}
	}
	return nil
}
func (this *GetCrossLinksResponse) Validate() error {
	for _, item := range this.CrossLinksSet {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("CrossLinksSet", err)
			}
		}
	}
	return nil
}
