// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/package/package.proto

package improbable_platform_package

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/improbable-io/go-proto-logfields"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PackageId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Generate fields for this message.
	return map[string]string{
		"package_type":    this.Type,
		"package_name":    this.Name,
		"package_version": this.Version,
	}
}

func (this *PackageId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	dst["package_type"] = this.Type
	dst["package_name"] = this.Name
	dst["package_version"] = this.Version
}

func (this *PackageCreationRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	return res
}

func (this *PackageCreationRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *Package) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range creationTimeFields {
		res[k] = v
	}
	return res
}

func (this *Package) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
}

func (this *PackageUploadUrl) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageUploadUrl) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListTypesResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListTypesResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListNamesByTypeRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListNamesByTypeRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListNamesByTypeResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListNamesByTypeResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListVersionsByNameRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListVersionsByNameRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListVersionsByNameResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListVersionsByNameResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListPackagesRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListPackagesRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListPackagesWithMetadataRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PackageListPackagesWithMetadataRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *PackageListPackagesWithMetadataResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	packageFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Package)
	hasInner = hasInner || len(packageFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range packageFields {
		res[k] = v
	}
	return res
}

func (this *PackageListPackagesWithMetadataResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Package, dst)
}