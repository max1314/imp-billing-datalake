// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base/serverstatus.proto

package base

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *HealthCheckResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *HealthCheckResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *FlagzState) LogFields() map[string]string {
	return map[string]string{}
}

func (this *FlagzState) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *VersionResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *VersionResponse) ExtractRequestFields(dst map[string]interface{}) {
}