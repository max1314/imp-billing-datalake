// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/resolver/lookup_proxy.proto

package improbable_infra

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *LookupSrvRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LookupSrvRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SrvRecord) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SrvRecord) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *LookupSrvResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LookupSrvResponse) ExtractRequestFields(dst map[string]interface{}) {
}
