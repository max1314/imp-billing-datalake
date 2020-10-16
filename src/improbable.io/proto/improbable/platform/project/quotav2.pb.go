// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/platform/project/quotav2.proto

package improbable_platform_project

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	platform "improbable.io/proto/improbable/platform"
	deployment "improbable.io/proto/improbable/platform/deployment"
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

type SetQuotaGrantRequest struct {
	ProjectId            *platform.ProjectId `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	QuotaGrant           uint64              `protobuf:"varint,2,opt,name=quota_grant,json=quotaGrant,proto3" json:"quota_grant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SetQuotaGrantRequest) Reset()         { *m = SetQuotaGrantRequest{} }
func (m *SetQuotaGrantRequest) String() string { return proto.CompactTextString(m) }
func (*SetQuotaGrantRequest) ProtoMessage()    {}
func (*SetQuotaGrantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f672c276fbaa7afd, []int{0}
}

func (m *SetQuotaGrantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetQuotaGrantRequest.Unmarshal(m, b)
}
func (m *SetQuotaGrantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetQuotaGrantRequest.Marshal(b, m, deterministic)
}
func (m *SetQuotaGrantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetQuotaGrantRequest.Merge(m, src)
}
func (m *SetQuotaGrantRequest) XXX_Size() int {
	return xxx_messageInfo_SetQuotaGrantRequest.Size(m)
}
func (m *SetQuotaGrantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetQuotaGrantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetQuotaGrantRequest proto.InternalMessageInfo

func (m *SetQuotaGrantRequest) GetProjectId() *platform.ProjectId {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *SetQuotaGrantRequest) GetQuotaGrant() uint64 {
	if m != nil {
		return m.QuotaGrant
	}
	return 0
}

type GetQuotaGrantAndUsageRequest struct {
	ProjectId            *platform.ProjectId `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetQuotaGrantAndUsageRequest) Reset()         { *m = GetQuotaGrantAndUsageRequest{} }
func (m *GetQuotaGrantAndUsageRequest) String() string { return proto.CompactTextString(m) }
func (*GetQuotaGrantAndUsageRequest) ProtoMessage()    {}
func (*GetQuotaGrantAndUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f672c276fbaa7afd, []int{1}
}

func (m *GetQuotaGrantAndUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuotaGrantAndUsageRequest.Unmarshal(m, b)
}
func (m *GetQuotaGrantAndUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuotaGrantAndUsageRequest.Marshal(b, m, deterministic)
}
func (m *GetQuotaGrantAndUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuotaGrantAndUsageRequest.Merge(m, src)
}
func (m *GetQuotaGrantAndUsageRequest) XXX_Size() int {
	return xxx_messageInfo_GetQuotaGrantAndUsageRequest.Size(m)
}
func (m *GetQuotaGrantAndUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuotaGrantAndUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuotaGrantAndUsageRequest proto.InternalMessageInfo

func (m *GetQuotaGrantAndUsageRequest) GetProjectId() *platform.ProjectId {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

type GetQuotaGrantAndUsageResponse struct {
	QuotaGrant           uint64        `protobuf:"varint,1,opt,name=quota_grant,json=quotaGrant,proto3" json:"quota_grant,omitempty"`
	QuotaUsages          []*QuotaUsage `protobuf:"bytes,2,rep,name=quota_usages,json=quotaUsages,proto3" json:"quota_usages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetQuotaGrantAndUsageResponse) Reset()         { *m = GetQuotaGrantAndUsageResponse{} }
func (m *GetQuotaGrantAndUsageResponse) String() string { return proto.CompactTextString(m) }
func (*GetQuotaGrantAndUsageResponse) ProtoMessage()    {}
func (*GetQuotaGrantAndUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f672c276fbaa7afd, []int{2}
}

func (m *GetQuotaGrantAndUsageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetQuotaGrantAndUsageResponse.Unmarshal(m, b)
}
func (m *GetQuotaGrantAndUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetQuotaGrantAndUsageResponse.Marshal(b, m, deterministic)
}
func (m *GetQuotaGrantAndUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetQuotaGrantAndUsageResponse.Merge(m, src)
}
func (m *GetQuotaGrantAndUsageResponse) XXX_Size() int {
	return xxx_messageInfo_GetQuotaGrantAndUsageResponse.Size(m)
}
func (m *GetQuotaGrantAndUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetQuotaGrantAndUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetQuotaGrantAndUsageResponse proto.InternalMessageInfo

func (m *GetQuotaGrantAndUsageResponse) GetQuotaGrant() uint64 {
	if m != nil {
		return m.QuotaGrant
	}
	return 0
}

func (m *GetQuotaGrantAndUsageResponse) GetQuotaUsages() []*QuotaUsage {
	if m != nil {
		return m.QuotaUsages
	}
	return nil
}

type QuotaUsage struct {
	ClusterId            *deployment.ClusterId `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	QuotaUsage           uint64                `protobuf:"varint,1,opt,name=quota_usage,json=quotaUsage,proto3" json:"quota_usage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *QuotaUsage) Reset()         { *m = QuotaUsage{} }
func (m *QuotaUsage) String() string { return proto.CompactTextString(m) }
func (*QuotaUsage) ProtoMessage()    {}
func (*QuotaUsage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f672c276fbaa7afd, []int{3}
}

func (m *QuotaUsage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuotaUsage.Unmarshal(m, b)
}
func (m *QuotaUsage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuotaUsage.Marshal(b, m, deterministic)
}
func (m *QuotaUsage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuotaUsage.Merge(m, src)
}
func (m *QuotaUsage) XXX_Size() int {
	return xxx_messageInfo_QuotaUsage.Size(m)
}
func (m *QuotaUsage) XXX_DiscardUnknown() {
	xxx_messageInfo_QuotaUsage.DiscardUnknown(m)
}

var xxx_messageInfo_QuotaUsage proto.InternalMessageInfo

func (m *QuotaUsage) GetClusterId() *deployment.ClusterId {
	if m != nil {
		return m.ClusterId
	}
	return nil
}

func (m *QuotaUsage) GetQuotaUsage() uint64 {
	if m != nil {
		return m.QuotaUsage
	}
	return 0
}

type DeleteQuotaGrantRequest struct {
	ProjectId            *platform.ProjectId `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeleteQuotaGrantRequest) Reset()         { *m = DeleteQuotaGrantRequest{} }
func (m *DeleteQuotaGrantRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteQuotaGrantRequest) ProtoMessage()    {}
func (*DeleteQuotaGrantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f672c276fbaa7afd, []int{4}
}

func (m *DeleteQuotaGrantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteQuotaGrantRequest.Unmarshal(m, b)
}
func (m *DeleteQuotaGrantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteQuotaGrantRequest.Marshal(b, m, deterministic)
}
func (m *DeleteQuotaGrantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteQuotaGrantRequest.Merge(m, src)
}
func (m *DeleteQuotaGrantRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteQuotaGrantRequest.Size(m)
}
func (m *DeleteQuotaGrantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteQuotaGrantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteQuotaGrantRequest proto.InternalMessageInfo

func (m *DeleteQuotaGrantRequest) GetProjectId() *platform.ProjectId {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

type DeleteQuotaGrantResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteQuotaGrantResponse) Reset()         { *m = DeleteQuotaGrantResponse{} }
func (m *DeleteQuotaGrantResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteQuotaGrantResponse) ProtoMessage()    {}
func (*DeleteQuotaGrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f672c276fbaa7afd, []int{5}
}

func (m *DeleteQuotaGrantResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteQuotaGrantResponse.Unmarshal(m, b)
}
func (m *DeleteQuotaGrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteQuotaGrantResponse.Marshal(b, m, deterministic)
}
func (m *DeleteQuotaGrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteQuotaGrantResponse.Merge(m, src)
}
func (m *DeleteQuotaGrantResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteQuotaGrantResponse.Size(m)
}
func (m *DeleteQuotaGrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteQuotaGrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteQuotaGrantResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetQuotaGrantRequest)(nil), "improbable.platform.project.SetQuotaGrantRequest")
	proto.RegisterType((*GetQuotaGrantAndUsageRequest)(nil), "improbable.platform.project.GetQuotaGrantAndUsageRequest")
	proto.RegisterType((*GetQuotaGrantAndUsageResponse)(nil), "improbable.platform.project.GetQuotaGrantAndUsageResponse")
	proto.RegisterType((*QuotaUsage)(nil), "improbable.platform.project.QuotaUsage")
	proto.RegisterType((*DeleteQuotaGrantRequest)(nil), "improbable.platform.project.DeleteQuotaGrantRequest")
	proto.RegisterType((*DeleteQuotaGrantResponse)(nil), "improbable.platform.project.DeleteQuotaGrantResponse")
}

func init() {
	proto.RegisterFile("improbable/platform/project/quotav2.proto", fileDescriptor_f672c276fbaa7afd)
}

var fileDescriptor_f672c276fbaa7afd = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xdf, 0x6a, 0xd4, 0x4e,
	0x14, 0x66, 0xf2, 0x2b, 0x85, 0xdf, 0xac, 0x82, 0x0c, 0x56, 0x97, 0xb4, 0xda, 0x10, 0x44, 0x77,
	0x4b, 0x3b, 0x83, 0xf1, 0x0f, 0xd8, 0x3b, 0xff, 0x94, 0xa2, 0x57, 0x75, 0x8b, 0xde, 0xd6, 0x49,
	0x72, 0x1a, 0xa3, 0xc9, 0x4c, 0x36, 0x99, 0x6c, 0x29, 0xa2, 0x82, 0xb7, 0x7a, 0xa7, 0xaf, 0xe1,
	0x1b, 0xf8, 0x16, 0x3e, 0x80, 0x20, 0x3e, 0x88, 0x64, 0x26, 0xdb, 0x74, 0xd7, 0x34, 0x54, 0xd1,
	0xbb, 0x64, 0xce, 0xf9, 0xe6, 0xfb, 0xbe, 0x33, 0xdf, 0xc1, 0xc3, 0x38, 0xcd, 0x72, 0xe9, 0x73,
	0x3f, 0x01, 0x96, 0x25, 0x5c, 0xed, 0xcb, 0x3c, 0x65, 0x59, 0x2e, 0x5f, 0x40, 0xa0, 0xd8, 0xb8,
	0x94, 0x8a, 0x4f, 0x3c, 0x9a, 0xe5, 0x52, 0x49, 0xb2, 0xdc, 0xb4, 0xd2, 0x69, 0x2b, 0xad, 0x5b,
	0x6d, 0xa7, 0xed, 0x9e, 0x40, 0xa6, 0xa9, 0x14, 0x06, 0x6e, 0xaf, 0xb7, 0x75, 0x84, 0x90, 0x25,
	0xf2, 0x30, 0x05, 0xa1, 0x58, 0x90, 0x94, 0x85, 0x82, 0xbc, 0xee, 0x5e, 0x89, 0xa4, 0x8c, 0x12,
	0x60, 0x3c, 0x8b, 0x19, 0x17, 0x42, 0x2a, 0xae, 0x62, 0x29, 0x8a, 0xba, 0xba, 0x5c, 0x57, 0xf5,
	0x9f, 0x5f, 0xee, 0x33, 0x48, 0x33, 0x75, 0x58, 0x17, 0x6f, 0x47, 0xb1, 0x7a, 0x5e, 0xfa, 0x34,
	0x90, 0x29, 0x4b, 0x0f, 0x62, 0xf5, 0x52, 0x1e, 0xb0, 0x48, 0x6e, 0xe8, 0xe2, 0xc6, 0x84, 0x27,
	0x71, 0xc8, 0x95, 0xcc, 0x0b, 0x76, 0xf4, 0x69, 0x70, 0xee, 0x1b, 0x7c, 0x7e, 0x17, 0xd4, 0xe3,
	0xca, 0xf3, 0x76, 0xce, 0x85, 0x1a, 0xc1, 0xb8, 0x84, 0x42, 0x91, 0x2d, 0x8c, 0x6b, 0x97, 0x7b,
	0x71, 0xd8, 0x47, 0x0e, 0x1a, 0xf4, 0xbc, 0xcb, 0xb4, 0x6d, 0x18, 0x3b, 0xa6, 0xed, 0x61, 0x78,
	0x6f, 0xf1, 0xfb, 0xb7, 0x55, 0xcb, 0x41, 0xa3, 0xff, 0xb3, 0xe9, 0x11, 0x59, 0xc5, 0x3d, 0x3d,
	0xcf, 0xbd, 0xa8, 0xba, 0xbc, 0x6f, 0x39, 0x68, 0xb0, 0x30, 0xc2, 0xe3, 0x23, 0x3a, 0x17, 0xf0,
	0xca, 0xf6, 0x71, 0xfe, 0xbb, 0x22, 0x7c, 0x52, 0xf0, 0x08, 0xfe, 0xae, 0x0e, 0xf7, 0x03, 0xc2,
	0x97, 0x4e, 0xe0, 0x29, 0x32, 0x29, 0x0a, 0x98, 0x57, 0x8a, 0xe6, 0x95, 0x92, 0x47, 0xf8, 0x8c,
	0x69, 0x28, 0x2b, 0x5c, 0xd1, 0xb7, 0x9c, 0xff, 0x06, 0x3d, 0xef, 0x1a, 0xed, 0x08, 0x08, 0xd5,
	0x7c, 0x86, 0xc7, 0xdc, 0xae, 0xbf, 0x0b, 0xf7, 0x2d, 0xc6, 0x4d, 0x89, 0xec, 0x60, 0x5c, 0xe7,
	0xa0, 0xf2, 0x68, 0x69, 0x8f, 0xc3, 0xd6, 0x7b, 0x9b, 0xe4, 0xd0, 0xfb, 0x06, 0x71, 0xdc, 0x6e,
	0x30, 0x3d, 0x6a, 0xcc, 0x68, 0xad, 0x33, 0x66, 0x34, 0xa5, 0xfb, 0x0c, 0x5f, 0x7c, 0x00, 0x09,
	0x28, 0xf8, 0x57, 0x2f, 0xef, 0xda, 0xb8, 0xff, 0x2b, 0x83, 0x99, 0xb5, 0xf7, 0x69, 0x01, 0x2f,
	0xd5, 0x60, 0x5d, 0x7d, 0xea, 0xed, 0x42, 0x3e, 0x89, 0x03, 0x20, 0xef, 0x11, 0x3e, 0x3b, 0x93,
	0x47, 0x72, 0xbd, 0x73, 0xc0, 0x6d, 0xd9, 0xb5, 0x2f, 0x50, 0xb3, 0x29, 0x74, 0xba, 0x29, 0x74,
	0xab, 0xda, 0x14, 0x97, 0xbd, 0xfb, 0xfa, 0xe3, 0xa3, 0x35, 0xb4, 0xaf, 0xb0, 0x58, 0x28, 0xc8,
	0x05, 0x4f, 0xcc, 0xb2, 0xb3, 0x89, 0xc7, 0x5e, 0x35, 0x9e, 0xa9, 0xe0, 0x29, 0xbc, 0xde, 0x44,
	0x6b, 0xe4, 0x0b, 0xc2, 0x4b, 0xad, 0xa9, 0x21, 0x77, 0x3a, 0x55, 0x75, 0x25, 0xda, 0xde, 0xfc,
	0x13, 0xa8, 0x19, 0x9c, 0xbb, 0xae, 0x1d, 0x5c, 0x25, 0xa7, 0x72, 0x40, 0x3e, 0x23, 0x7c, 0x6e,
	0xfe, 0x0d, 0xc8, 0xcd, 0x4e, 0xfa, 0x13, 0x42, 0x61, 0xdf, 0xfa, 0x4d, 0xd4, 0xac, 0xde, 0xb5,
	0x53, 0xe9, 0xf5, 0x17, 0xf5, 0x7b, 0xdd, 0xf8, 0x19, 0x00, 0x00, 0xff, 0xff, 0x32, 0x7b, 0xbc,
	0x2e, 0x9f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProjectQuotaV2ServiceClient is the client API for ProjectQuotaV2Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectQuotaV2ServiceClient interface {
	SetQuotaGrant(ctx context.Context, in *SetQuotaGrantRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetQuotaGrantAndUsage(ctx context.Context, in *GetQuotaGrantAndUsageRequest, opts ...grpc.CallOption) (*GetQuotaGrantAndUsageResponse, error)
	DeleteQuotaGrant(ctx context.Context, in *DeleteQuotaGrantRequest, opts ...grpc.CallOption) (*DeleteQuotaGrantResponse, error)
}

type projectQuotaV2ServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectQuotaV2ServiceClient(cc grpc.ClientConnInterface) ProjectQuotaV2ServiceClient {
	return &projectQuotaV2ServiceClient{cc}
}

func (c *projectQuotaV2ServiceClient) SetQuotaGrant(ctx context.Context, in *SetQuotaGrantRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectQuotaV2Service/SetQuotaGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectQuotaV2ServiceClient) GetQuotaGrantAndUsage(ctx context.Context, in *GetQuotaGrantAndUsageRequest, opts ...grpc.CallOption) (*GetQuotaGrantAndUsageResponse, error) {
	out := new(GetQuotaGrantAndUsageResponse)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectQuotaV2Service/GetQuotaGrantAndUsage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectQuotaV2ServiceClient) DeleteQuotaGrant(ctx context.Context, in *DeleteQuotaGrantRequest, opts ...grpc.CallOption) (*DeleteQuotaGrantResponse, error) {
	out := new(DeleteQuotaGrantResponse)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectQuotaV2Service/DeleteQuotaGrant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectQuotaV2ServiceServer is the server API for ProjectQuotaV2Service service.
type ProjectQuotaV2ServiceServer interface {
	SetQuotaGrant(context.Context, *SetQuotaGrantRequest) (*empty.Empty, error)
	GetQuotaGrantAndUsage(context.Context, *GetQuotaGrantAndUsageRequest) (*GetQuotaGrantAndUsageResponse, error)
	DeleteQuotaGrant(context.Context, *DeleteQuotaGrantRequest) (*DeleteQuotaGrantResponse, error)
}

// UnimplementedProjectQuotaV2ServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProjectQuotaV2ServiceServer struct {
}

func (*UnimplementedProjectQuotaV2ServiceServer) SetQuotaGrant(ctx context.Context, req *SetQuotaGrantRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetQuotaGrant not implemented")
}
func (*UnimplementedProjectQuotaV2ServiceServer) GetQuotaGrantAndUsage(ctx context.Context, req *GetQuotaGrantAndUsageRequest) (*GetQuotaGrantAndUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotaGrantAndUsage not implemented")
}
func (*UnimplementedProjectQuotaV2ServiceServer) DeleteQuotaGrant(ctx context.Context, req *DeleteQuotaGrantRequest) (*DeleteQuotaGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuotaGrant not implemented")
}

func RegisterProjectQuotaV2ServiceServer(s *grpc.Server, srv ProjectQuotaV2ServiceServer) {
	s.RegisterService(&_ProjectQuotaV2Service_serviceDesc, srv)
}

func _ProjectQuotaV2Service_SetQuotaGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetQuotaGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectQuotaV2ServiceServer).SetQuotaGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectQuotaV2Service/SetQuotaGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectQuotaV2ServiceServer).SetQuotaGrant(ctx, req.(*SetQuotaGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectQuotaV2Service_GetQuotaGrantAndUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuotaGrantAndUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectQuotaV2ServiceServer).GetQuotaGrantAndUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectQuotaV2Service/GetQuotaGrantAndUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectQuotaV2ServiceServer).GetQuotaGrantAndUsage(ctx, req.(*GetQuotaGrantAndUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectQuotaV2Service_DeleteQuotaGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuotaGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectQuotaV2ServiceServer).DeleteQuotaGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectQuotaV2Service/DeleteQuotaGrant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectQuotaV2ServiceServer).DeleteQuotaGrant(ctx, req.(*DeleteQuotaGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectQuotaV2Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.platform.project.ProjectQuotaV2Service",
	HandlerType: (*ProjectQuotaV2ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetQuotaGrant",
			Handler:    _ProjectQuotaV2Service_SetQuotaGrant_Handler,
		},
		{
			MethodName: "GetQuotaGrantAndUsage",
			Handler:    _ProjectQuotaV2Service_GetQuotaGrantAndUsage_Handler,
		},
		{
			MethodName: "DeleteQuotaGrant",
			Handler:    _ProjectQuotaV2Service_DeleteQuotaGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/platform/project/quotav2.proto",
}