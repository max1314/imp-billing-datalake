// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/platform/local/launch.proto

package improbable_platform_local

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LaunchConfiguration struct {
	Javaparams           []*Parameter   `protobuf:"bytes,1,rep,name=javaparams,proto3" json:"javaparams,omitempty"`
	Flagz                []*FlagValue   `protobuf:"bytes,2,rep,name=flagz,proto3" json:"flagz,omitempty"`
	WorkerFlagz          []*WorkerFlagz `protobuf:"bytes,3,rep,name=worker_flagz,json=workerFlagz,proto3" json:"worker_flagz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LaunchConfiguration) Reset()         { *m = LaunchConfiguration{} }
func (m *LaunchConfiguration) String() string { return proto.CompactTextString(m) }
func (*LaunchConfiguration) ProtoMessage()    {}
func (*LaunchConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_2737822d0df8c77c, []int{0}
}

func (m *LaunchConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LaunchConfiguration.Unmarshal(m, b)
}
func (m *LaunchConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LaunchConfiguration.Marshal(b, m, deterministic)
}
func (m *LaunchConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchConfiguration.Merge(m, src)
}
func (m *LaunchConfiguration) XXX_Size() int {
	return xxx_messageInfo_LaunchConfiguration.Size(m)
}
func (m *LaunchConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchConfiguration proto.InternalMessageInfo

func (m *LaunchConfiguration) GetJavaparams() []*Parameter {
	if m != nil {
		return m.Javaparams
	}
	return nil
}

func (m *LaunchConfiguration) GetFlagz() []*FlagValue {
	if m != nil {
		return m.Flagz
	}
	return nil
}

func (m *LaunchConfiguration) GetWorkerFlagz() []*WorkerFlagz {
	if m != nil {
		return m.WorkerFlagz
	}
	return nil
}

type Parameter struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2737822d0df8c77c, []int{1}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type WorkerFlagz struct {
	WorkerType           string       `protobuf:"bytes,1,opt,name=worker_type,json=workerType,proto3" json:"worker_type,omitempty"`
	Flagz                []*FlagValue `protobuf:"bytes,2,rep,name=flagz,proto3" json:"flagz,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *WorkerFlagz) Reset()         { *m = WorkerFlagz{} }
func (m *WorkerFlagz) String() string { return proto.CompactTextString(m) }
func (*WorkerFlagz) ProtoMessage()    {}
func (*WorkerFlagz) Descriptor() ([]byte, []int) {
	return fileDescriptor_2737822d0df8c77c, []int{2}
}

func (m *WorkerFlagz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerFlagz.Unmarshal(m, b)
}
func (m *WorkerFlagz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerFlagz.Marshal(b, m, deterministic)
}
func (m *WorkerFlagz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerFlagz.Merge(m, src)
}
func (m *WorkerFlagz) XXX_Size() int {
	return xxx_messageInfo_WorkerFlagz.Size(m)
}
func (m *WorkerFlagz) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerFlagz.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerFlagz proto.InternalMessageInfo

func (m *WorkerFlagz) GetWorkerType() string {
	if m != nil {
		return m.WorkerType
	}
	return ""
}

func (m *WorkerFlagz) GetFlagz() []*FlagValue {
	if m != nil {
		return m.Flagz
	}
	return nil
}

type FlagValue struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagValue) Reset()         { *m = FlagValue{} }
func (m *FlagValue) String() string { return proto.CompactTextString(m) }
func (*FlagValue) ProtoMessage()    {}
func (*FlagValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_2737822d0df8c77c, []int{3}
}

func (m *FlagValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagValue.Unmarshal(m, b)
}
func (m *FlagValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagValue.Marshal(b, m, deterministic)
}
func (m *FlagValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagValue.Merge(m, src)
}
func (m *FlagValue) XXX_Size() int {
	return xxx_messageInfo_FlagValue.Size(m)
}
func (m *FlagValue) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagValue.DiscardUnknown(m)
}

var xxx_messageInfo_FlagValue proto.InternalMessageInfo

func (m *FlagValue) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlagValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*LaunchConfiguration)(nil), "improbable.platform.local.LaunchConfiguration")
	proto.RegisterType((*Parameter)(nil), "improbable.platform.local.Parameter")
	proto.RegisterType((*WorkerFlagz)(nil), "improbable.platform.local.WorkerFlagz")
	proto.RegisterType((*FlagValue)(nil), "improbable.platform.local.FlagValue")
}

func init() {
	proto.RegisterFile("improbable/platform/local/launch.proto", fileDescriptor_2737822d0df8c77c)
}

var fileDescriptor_2737822d0df8c77c = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x6b, 0x85, 0x4e, 0x3c, 0xad, 0x1e, 0xe2, 0xc9, 0x12, 0xa4, 0xf4, 0xb4, 0x01,
	0xc5, 0x8b, 0x57, 0x45, 0x10, 0x3c, 0x48, 0x10, 0x3d, 0x96, 0x49, 0xd9, 0xc4, 0xd4, 0x49, 0x66,
	0x19, 0x37, 0x2d, 0xed, 0xcf, 0xf5, 0x97, 0x48, 0x36, 0xf6, 0xe3, 0xe2, 0x07, 0x78, 0xdb, 0x19,
	0x9e, 0xf7, 0xd9, 0x77, 0x59, 0x18, 0x97, 0x95, 0x15, 0xce, 0x30, 0x23, 0x93, 0x58, 0x42, 0x97,
	0xb3, 0x54, 0x09, 0xf1, 0x0c, 0x29, 0x21, 0x6c, 0xea, 0xd9, 0xab, 0xb6, 0xc2, 0x8e, 0xd5, 0xe9,
	0x8e, 0xd3, 0x1b, 0x4e, 0x7b, 0x2e, 0xfe, 0x08, 0xe0, 0xf8, 0xc1, 0xb3, 0x37, 0x5c, 0xe7, 0x65,
	0xd1, 0x08, 0xba, 0x92, 0x6b, 0x75, 0x0b, 0x30, 0xc7, 0x05, 0x5a, 0x14, 0xac, 0xde, 0xa3, 0x60,
	0xd4, 0x9f, 0x84, 0x17, 0xe7, 0xfa, 0x5b, 0x8f, 0x7e, 0x6c, 0x41, 0xe3, 0x8c, 0xa4, 0x7b, 0x39,
	0x75, 0x0d, 0x83, 0x9c, 0xb0, 0x58, 0x47, 0xbd, 0x5f, 0x05, 0x77, 0x84, 0xc5, 0x33, 0x52, 0x63,
	0xd2, 0x2e, 0xa2, 0xee, 0xe1, 0x68, 0xc9, 0xf2, 0x66, 0x64, 0xda, 0x29, 0xfa, 0x5e, 0x31, 0xfe,
	0x41, 0xf1, 0xe2, 0xf1, 0x56, 0xb4, 0x4e, 0xc3, 0xe5, 0x6e, 0x88, 0xaf, 0x60, 0xb8, 0xed, 0xa7,
	0x14, 0x1c, 0xd4, 0x58, 0x99, 0x28, 0x18, 0x05, 0x93, 0x61, 0xea, 0xcf, 0xea, 0x04, 0x06, 0x8b,
	0xf6, 0xee, 0xa8, 0xe7, 0x97, 0xdd, 0x10, 0xcf, 0x21, 0xdc, 0x53, 0xaa, 0x33, 0xf8, 0x92, 0x4e,
	0xdd, 0xca, 0x6e, 0xf2, 0xd0, 0xad, 0x9e, 0x56, 0xd6, 0xfc, 0xe7, 0xb5, 0x6d, 0xc5, 0xed, 0xee,
	0xef, 0x15, 0xb3, 0x43, 0xff, 0xc1, 0x97, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x28, 0x97,
	0x65, 0x0a, 0x02, 0x00, 0x00,
}
