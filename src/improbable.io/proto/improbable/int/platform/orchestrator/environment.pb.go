// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/int/platform/orchestrator/environment.proto

package improbable_int_platform_orchestrator

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Environment struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProjectId            string               `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DisplayName          string               `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Version              string               `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{0}
}

func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (m *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(m, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Environment) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *Environment) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Environment) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Environment) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CreateEnvironmentRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEnvironmentRequest) Reset()         { *m = CreateEnvironmentRequest{} }
func (m *CreateEnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironmentRequest) ProtoMessage()    {}
func (*CreateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{1}
}

func (m *CreateEnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironmentRequest.Unmarshal(m, b)
}
func (m *CreateEnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironmentRequest.Marshal(b, m, deterministic)
}
func (m *CreateEnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironmentRequest.Merge(m, src)
}
func (m *CreateEnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironmentRequest.Size(m)
}
func (m *CreateEnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironmentRequest proto.InternalMessageInfo

func (m *CreateEnvironmentRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CreateEnvironmentRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CreateEnvironmentRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type CreateEnvironmentResponse struct {
	Environment          *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateEnvironmentResponse) Reset()         { *m = CreateEnvironmentResponse{} }
func (m *CreateEnvironmentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironmentResponse) ProtoMessage()    {}
func (*CreateEnvironmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{2}
}

func (m *CreateEnvironmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironmentResponse.Unmarshal(m, b)
}
func (m *CreateEnvironmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironmentResponse.Marshal(b, m, deterministic)
}
func (m *CreateEnvironmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironmentResponse.Merge(m, src)
}
func (m *CreateEnvironmentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironmentResponse.Size(m)
}
func (m *CreateEnvironmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironmentResponse proto.InternalMessageInfo

func (m *CreateEnvironmentResponse) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

type ListEnvironmentRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListEnvironmentRequest) Reset()         { *m = ListEnvironmentRequest{} }
func (m *ListEnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*ListEnvironmentRequest) ProtoMessage()    {}
func (*ListEnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{3}
}

func (m *ListEnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEnvironmentRequest.Unmarshal(m, b)
}
func (m *ListEnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEnvironmentRequest.Marshal(b, m, deterministic)
}
func (m *ListEnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEnvironmentRequest.Merge(m, src)
}
func (m *ListEnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_ListEnvironmentRequest.Size(m)
}
func (m *ListEnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListEnvironmentRequest proto.InternalMessageInfo

func (m *ListEnvironmentRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type ListEnvironmentResponse struct {
	Environments         []*Environment `protobuf:"bytes,1,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListEnvironmentResponse) Reset()         { *m = ListEnvironmentResponse{} }
func (m *ListEnvironmentResponse) String() string { return proto.CompactTextString(m) }
func (*ListEnvironmentResponse) ProtoMessage()    {}
func (*ListEnvironmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{4}
}

func (m *ListEnvironmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListEnvironmentResponse.Unmarshal(m, b)
}
func (m *ListEnvironmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListEnvironmentResponse.Marshal(b, m, deterministic)
}
func (m *ListEnvironmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListEnvironmentResponse.Merge(m, src)
}
func (m *ListEnvironmentResponse) XXX_Size() int {
	return xxx_messageInfo_ListEnvironmentResponse.Size(m)
}
func (m *ListEnvironmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListEnvironmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListEnvironmentResponse proto.InternalMessageInfo

func (m *ListEnvironmentResponse) GetEnvironments() []*Environment {
	if m != nil {
		return m.Environments
	}
	return nil
}

type GetEnvironmentRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	EnvironmentId        string   `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEnvironmentRequest) Reset()         { *m = GetEnvironmentRequest{} }
func (m *GetEnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*GetEnvironmentRequest) ProtoMessage()    {}
func (*GetEnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{5}
}

func (m *GetEnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnvironmentRequest.Unmarshal(m, b)
}
func (m *GetEnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnvironmentRequest.Marshal(b, m, deterministic)
}
func (m *GetEnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnvironmentRequest.Merge(m, src)
}
func (m *GetEnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_GetEnvironmentRequest.Size(m)
}
func (m *GetEnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnvironmentRequest proto.InternalMessageInfo

func (m *GetEnvironmentRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *GetEnvironmentRequest) GetEnvironmentId() string {
	if m != nil {
		return m.EnvironmentId
	}
	return ""
}

type GetEnvironmentResponse struct {
	Environment          *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetEnvironmentResponse) Reset()         { *m = GetEnvironmentResponse{} }
func (m *GetEnvironmentResponse) String() string { return proto.CompactTextString(m) }
func (*GetEnvironmentResponse) ProtoMessage()    {}
func (*GetEnvironmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{6}
}

func (m *GetEnvironmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnvironmentResponse.Unmarshal(m, b)
}
func (m *GetEnvironmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnvironmentResponse.Marshal(b, m, deterministic)
}
func (m *GetEnvironmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnvironmentResponse.Merge(m, src)
}
func (m *GetEnvironmentResponse) XXX_Size() int {
	return xxx_messageInfo_GetEnvironmentResponse.Size(m)
}
func (m *GetEnvironmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnvironmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnvironmentResponse proto.InternalMessageInfo

func (m *GetEnvironmentResponse) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

type DeleteEnvironmentRequest struct {
	ProjectId            string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	EnvironmentId        string   `protobuf:"bytes,2,opt,name=environment_id,json=environmentId,proto3" json:"environment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEnvironmentRequest) Reset()         { *m = DeleteEnvironmentRequest{} }
func (m *DeleteEnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteEnvironmentRequest) ProtoMessage()    {}
func (*DeleteEnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{7}
}

func (m *DeleteEnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEnvironmentRequest.Unmarshal(m, b)
}
func (m *DeleteEnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEnvironmentRequest.Marshal(b, m, deterministic)
}
func (m *DeleteEnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEnvironmentRequest.Merge(m, src)
}
func (m *DeleteEnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteEnvironmentRequest.Size(m)
}
func (m *DeleteEnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEnvironmentRequest proto.InternalMessageInfo

func (m *DeleteEnvironmentRequest) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *DeleteEnvironmentRequest) GetEnvironmentId() string {
	if m != nil {
		return m.EnvironmentId
	}
	return ""
}

type DeleteEnvironmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEnvironmentResponse) Reset()         { *m = DeleteEnvironmentResponse{} }
func (m *DeleteEnvironmentResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEnvironmentResponse) ProtoMessage()    {}
func (*DeleteEnvironmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7f94e7f5e1bbbe1, []int{8}
}

func (m *DeleteEnvironmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEnvironmentResponse.Unmarshal(m, b)
}
func (m *DeleteEnvironmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEnvironmentResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEnvironmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEnvironmentResponse.Merge(m, src)
}
func (m *DeleteEnvironmentResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEnvironmentResponse.Size(m)
}
func (m *DeleteEnvironmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEnvironmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEnvironmentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Environment)(nil), "improbable.int.platform.orchestrator.Environment")
	proto.RegisterType((*CreateEnvironmentRequest)(nil), "improbable.int.platform.orchestrator.CreateEnvironmentRequest")
	proto.RegisterType((*CreateEnvironmentResponse)(nil), "improbable.int.platform.orchestrator.CreateEnvironmentResponse")
	proto.RegisterType((*ListEnvironmentRequest)(nil), "improbable.int.platform.orchestrator.ListEnvironmentRequest")
	proto.RegisterType((*ListEnvironmentResponse)(nil), "improbable.int.platform.orchestrator.ListEnvironmentResponse")
	proto.RegisterType((*GetEnvironmentRequest)(nil), "improbable.int.platform.orchestrator.GetEnvironmentRequest")
	proto.RegisterType((*GetEnvironmentResponse)(nil), "improbable.int.platform.orchestrator.GetEnvironmentResponse")
	proto.RegisterType((*DeleteEnvironmentRequest)(nil), "improbable.int.platform.orchestrator.DeleteEnvironmentRequest")
	proto.RegisterType((*DeleteEnvironmentResponse)(nil), "improbable.int.platform.orchestrator.DeleteEnvironmentResponse")
}

func init() {
	proto.RegisterFile("improbable/int/platform/orchestrator/environment.proto", fileDescriptor_b7f94e7f5e1bbbe1)
}

var fileDescriptor_b7f94e7f5e1bbbe1 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x35, 0x4e, 0x09, 0xea, 0x49, 0xe9, 0x62, 0x24, 0x8a, 0x6b, 0x2e, 0x0d, 0x16, 0x48,
	0x25, 0x22, 0x1e, 0x25, 0x88, 0x4a, 0xd0, 0x96, 0x0a, 0x0a, 0xaa, 0x2a, 0x01, 0x8b, 0x14, 0x10,
	0xbb, 0x68, 0x12, 0x4f, 0xd3, 0x01, 0xdb, 0x63, 0x3c, 0x93, 0x54, 0xa8, 0xea, 0x86, 0x07, 0x60,
	0xc3, 0xab, 0x20, 0xd8, 0xb2, 0x61, 0xcb, 0x86, 0x07, 0x40, 0x42, 0x3c, 0x08, 0xca, 0xd8, 0x49,
	0xec, 0x92, 0x48, 0xae, 0x2b, 0x76, 0x9e, 0x33, 0xe7, 0xf2, 0xfd, 0xce, 0x7f, 0x1c, 0x58, 0xe3,
	0x7e, 0x18, 0x89, 0x0e, 0xed, 0x78, 0x8c, 0xf0, 0x40, 0x91, 0xd0, 0xa3, 0x6a, 0x5f, 0x44, 0x3e,
	0x11, 0x51, 0xf7, 0x80, 0x49, 0x15, 0x51, 0x25, 0x22, 0xc2, 0x82, 0x01, 0x8f, 0x44, 0xe0, 0xb3,
	0x40, 0x39, 0x61, 0x24, 0x94, 0xc0, 0x37, 0x26, 0x75, 0x0e, 0x1f, 0x46, 0x93, 0x3a, 0x27, 0x5d,
	0x67, 0x5d, 0xe9, 0x09, 0xd1, 0xf3, 0x18, 0xa1, 0x21, 0x27, 0x34, 0x08, 0x84, 0xa2, 0x8a, 0x8b,
	0x40, 0xc6, 0x3d, 0xac, 0x95, 0xe4, 0x56, 0x9f, 0x3a, 0xfd, 0x7d, 0xa2, 0xb8, 0xcf, 0xa4, 0xa2,
	0x7e, 0x98, 0x24, 0xac, 0xf5, 0xb8, 0x3a, 0xe8, 0x77, 0x9c, 0xae, 0xf0, 0x89, 0x7f, 0xc8, 0xd5,
	0x5b, 0x71, 0x48, 0x7a, 0xa2, 0xae, 0x2f, 0xeb, 0x03, 0xea, 0x71, 0x77, 0x38, 0x45, 0x92, 0xf1,
	0x63, 0x5c, 0x67, 0x7f, 0x46, 0x50, 0x79, 0x32, 0x41, 0xc6, 0x8b, 0x60, 0x70, 0xd7, 0x44, 0x55,
	0xb4, 0x3a, 0xdf, 0x32, 0xb8, 0x8b, 0xaf, 0x02, 0x84, 0x91, 0x78, 0xc3, 0xba, 0xaa, 0xcd, 0x5d,
	0xd3, 0xd0, 0xf1, 0xf9, 0x24, 0xb2, 0xeb, 0xe2, 0xeb, 0xb0, 0xe0, 0x72, 0x19, 0x7a, 0xf4, 0x7d,
	0x3b, 0xa0, 0x3e, 0x33, 0x4b, 0x3a, 0xa1, 0x92, 0xc4, 0x9e, 0x53, 0x9f, 0x61, 0x13, 0xce, 0x0f,
	0x58, 0x24, 0xb9, 0x08, 0xcc, 0x39, 0x7d, 0x3b, 0x3a, 0xe2, 0x7b, 0x00, 0xdd, 0x88, 0x51, 0xc5,
	0xdc, 0x36, 0x55, 0xe6, 0xb9, 0x2a, 0x5a, 0xad, 0x34, 0x2d, 0x27, 0x56, 0xea, 0x8c, 0x94, 0x3a,
	0x2f, 0x46, 0x4a, 0x5b, 0xf3, 0x49, 0xf6, 0x43, 0x65, 0x7f, 0x44, 0x60, 0x6e, 0xeb, 0x53, 0x0a,
	0xbe, 0xc5, 0xde, 0xf5, 0x99, 0x54, 0xf8, 0x66, 0x86, 0x59, 0x6b, 0x79, 0x54, 0xfe, 0xfd, 0x6b,
	0xc5, 0x78, 0x8d, 0xd2, 0xec, 0xb7, 0x4e, 0xb0, 0x1b, 0x99, 0xc4, 0x8c, 0x86, 0xea, 0x44, 0x43,
	0x29, 0x93, 0x35, 0x0a, 0xdb, 0x21, 0x2c, 0x4f, 0xe1, 0x91, 0xa1, 0x08, 0x24, 0xc3, 0x7b, 0x50,
	0x49, 0xd9, 0x42, 0x13, 0x55, 0x9a, 0x0d, 0x27, 0x8f, 0x2f, 0x9c, 0x74, 0xbf, 0x74, 0x17, 0x7b,
	0x0b, 0x96, 0x9e, 0x72, 0xa9, 0x0a, 0xeb, 0xb7, 0x43, 0xb8, 0xf4, 0x4f, 0x83, 0x04, 0xf8, 0x25,
	0x2c, 0xa4, 0x46, 0x49, 0x13, 0x55, 0x4b, 0xc5, 0x88, 0x33, 0x6d, 0x6c, 0x1f, 0x2e, 0xee, 0xb0,
	0xe2, 0xc4, 0xb8, 0x0e, 0x8b, 0xa9, 0x7e, 0x63, 0x43, 0x8e, 0x53, 0x2f, 0xa4, 0x6e, 0x77, 0x5d,
	0xdb, 0x87, 0xa5, 0x93, 0xe3, 0xfe, 0xe7, 0x0f, 0x12, 0x82, 0xf9, 0x98, 0x79, 0xec, 0x2c, 0x96,
	0x3c, 0xa5, 0xc0, 0xcb, 0xb0, 0x3c, 0x65, 0x62, 0xac, 0xb1, 0xf9, 0xa5, 0x0c, 0xd7, 0x52, 0xf1,
	0x67, 0x34, 0xa0, 0x3d, 0x36, 0x7c, 0xda, 0x63, 0xd1, 0x80, 0x77, 0xd9, 0xab, 0x06, 0xfe, 0x86,
	0xa0, 0x1c, 0xbb, 0x16, 0x3f, 0xc8, 0x27, 0x7e, 0xd6, 0xce, 0x59, 0x5b, 0x85, 0xeb, 0x63, 0x5c,
	0xfb, 0xee, 0x87, 0x9f, 0x7f, 0x3e, 0x19, 0xc4, 0xae, 0xe9, 0x2f, 0xe0, 0xa0, 0x41, 0x92, 0xb7,
	0x22, 0xc9, 0xd1, 0xe4, 0xcd, 0x1d, 0xa7, 0x3f, 0xaf, 0xf2, 0x3e, 0xaa, 0xe1, 0xaf, 0x08, 0xe6,
	0x86, 0x2e, 0xc6, 0x1b, 0xf9, 0x00, 0xa6, 0xaf, 0x8c, 0xb5, 0x59, 0xb0, 0x3a, 0x81, 0x6f, 0x6a,
	0xf8, 0xdb, 0xf8, 0x14, 0xf0, 0xf8, 0x3b, 0x82, 0xd2, 0x0e, 0x53, 0x78, 0x3d, 0xdf, 0xe8, 0xa9,
	0x8b, 0x63, 0x6d, 0x14, 0x2b, 0x4e, 0xb0, 0xb7, 0x35, 0xf6, 0x26, 0x5e, 0xcf, 0x8f, 0x4d, 0x8e,
	0xb2, 0x06, 0x3d, 0xc6, 0x3f, 0x10, 0x94, 0x63, 0x17, 0xe6, 0x35, 0xd1, 0xac, 0x2d, 0xc9, 0x6b,
	0xa2, 0x99, 0x9e, 0x1f, 0x09, 0xaa, 0x9d, 0x45, 0x50, 0xa7, 0xac, 0xff, 0x7a, 0xee, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0x29, 0x98, 0xad, 0xf0, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EnvironmentManagementServiceV1Client is the client API for EnvironmentManagementServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EnvironmentManagementServiceV1Client interface {
	Create(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (*CreateEnvironmentResponse, error)
	List(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error)
	Get(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*GetEnvironmentResponse, error)
	Delete(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (*DeleteEnvironmentResponse, error)
}

type environmentManagementServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentManagementServiceV1Client(cc grpc.ClientConnInterface) EnvironmentManagementServiceV1Client {
	return &environmentManagementServiceV1Client{cc}
}

func (c *environmentManagementServiceV1Client) Create(ctx context.Context, in *CreateEnvironmentRequest, opts ...grpc.CallOption) (*CreateEnvironmentResponse, error) {
	out := new(CreateEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagementServiceV1Client) List(ctx context.Context, in *ListEnvironmentRequest, opts ...grpc.CallOption) (*ListEnvironmentResponse, error) {
	out := new(ListEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagementServiceV1Client) Get(ctx context.Context, in *GetEnvironmentRequest, opts ...grpc.CallOption) (*GetEnvironmentResponse, error) {
	out := new(GetEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentManagementServiceV1Client) Delete(ctx context.Context, in *DeleteEnvironmentRequest, opts ...grpc.CallOption) (*DeleteEnvironmentResponse, error) {
	out := new(DeleteEnvironmentResponse)
	err := c.cc.Invoke(ctx, "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentManagementServiceV1Server is the server API for EnvironmentManagementServiceV1 service.
type EnvironmentManagementServiceV1Server interface {
	Create(context.Context, *CreateEnvironmentRequest) (*CreateEnvironmentResponse, error)
	List(context.Context, *ListEnvironmentRequest) (*ListEnvironmentResponse, error)
	Get(context.Context, *GetEnvironmentRequest) (*GetEnvironmentResponse, error)
	Delete(context.Context, *DeleteEnvironmentRequest) (*DeleteEnvironmentResponse, error)
}

// UnimplementedEnvironmentManagementServiceV1Server can be embedded to have forward compatible implementations.
type UnimplementedEnvironmentManagementServiceV1Server struct {
}

func (*UnimplementedEnvironmentManagementServiceV1Server) Create(ctx context.Context, req *CreateEnvironmentRequest) (*CreateEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEnvironmentManagementServiceV1Server) List(ctx context.Context, req *ListEnvironmentRequest) (*ListEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEnvironmentManagementServiceV1Server) Get(ctx context.Context, req *GetEnvironmentRequest) (*GetEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEnvironmentManagementServiceV1Server) Delete(ctx context.Context, req *DeleteEnvironmentRequest) (*DeleteEnvironmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEnvironmentManagementServiceV1Server(s *grpc.Server, srv EnvironmentManagementServiceV1Server) {
	s.RegisterService(&_EnvironmentManagementServiceV1_serviceDesc, srv)
}

func _EnvironmentManagementServiceV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagementServiceV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagementServiceV1Server).Create(ctx, req.(*CreateEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManagementServiceV1_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagementServiceV1Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagementServiceV1Server).List(ctx, req.(*ListEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManagementServiceV1_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagementServiceV1Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagementServiceV1Server).Get(ctx, req.(*GetEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentManagementServiceV1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEnvironmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentManagementServiceV1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.int.platform.orchestrator.EnvironmentManagementServiceV1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentManagementServiceV1Server).Delete(ctx, req.(*DeleteEnvironmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EnvironmentManagementServiceV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.int.platform.orchestrator.EnvironmentManagementServiceV1",
	HandlerType: (*EnvironmentManagementServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _EnvironmentManagementServiceV1_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _EnvironmentManagementServiceV1_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _EnvironmentManagementServiceV1_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _EnvironmentManagementServiceV1_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/int/platform/orchestrator/environment.proto",
}