// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/toolbelt/wrapper/worker_package_cache/worker_package_cache.proto

package improbable_toolbelt_wrapper_worker_package_cache

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "improbable.io/proto/improbable/platform/package"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *WorkerPackageGetRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	packageIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.PackageId)
	hasInner = hasInner || len(packageIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range packageIdFields {
		res[k] = v
	}
	return res
}

func (this *WorkerPackageGetRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.PackageId, dst)
}

func (this *WorkerPackageGetResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *WorkerPackageGetResponse) ExtractRequestFields(dst map[string]interface{}) {
}
