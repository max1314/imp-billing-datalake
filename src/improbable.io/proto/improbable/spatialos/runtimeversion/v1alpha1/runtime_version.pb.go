// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/spatialos/runtimeversion/v1alpha1/runtime_version.proto

package runtimeversion

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type ListSupportedRuntimeVersionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListSupportedRuntimeVersionsRequest) Reset()         { *m = ListSupportedRuntimeVersionsRequest{} }
func (m *ListSupportedRuntimeVersionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListSupportedRuntimeVersionsRequest) ProtoMessage()    {}
func (*ListSupportedRuntimeVersionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4224d18f2d1bfb4e, []int{0}
}

func (m *ListSupportedRuntimeVersionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSupportedRuntimeVersionsRequest.Unmarshal(m, b)
}
func (m *ListSupportedRuntimeVersionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSupportedRuntimeVersionsRequest.Marshal(b, m, deterministic)
}
func (m *ListSupportedRuntimeVersionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSupportedRuntimeVersionsRequest.Merge(m, src)
}
func (m *ListSupportedRuntimeVersionsRequest) XXX_Size() int {
	return xxx_messageInfo_ListSupportedRuntimeVersionsRequest.Size(m)
}
func (m *ListSupportedRuntimeVersionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSupportedRuntimeVersionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListSupportedRuntimeVersionsRequest proto.InternalMessageInfo

type ListSupportedRuntimeVersionsResponse struct {
	SupportedVersions    []*ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion `protobuf:"bytes,1,rep,name=supported_versions,json=supportedVersions,proto3" json:"supported_versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                        `json:"-"`
	XXX_unrecognized     []byte                                                          `json:"-"`
	XXX_sizecache        int32                                                           `json:"-"`
}

func (m *ListSupportedRuntimeVersionsResponse) Reset()         { *m = ListSupportedRuntimeVersionsResponse{} }
func (m *ListSupportedRuntimeVersionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListSupportedRuntimeVersionsResponse) ProtoMessage()    {}
func (*ListSupportedRuntimeVersionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4224d18f2d1bfb4e, []int{1}
}

func (m *ListSupportedRuntimeVersionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSupportedRuntimeVersionsResponse.Unmarshal(m, b)
}
func (m *ListSupportedRuntimeVersionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSupportedRuntimeVersionsResponse.Marshal(b, m, deterministic)
}
func (m *ListSupportedRuntimeVersionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSupportedRuntimeVersionsResponse.Merge(m, src)
}
func (m *ListSupportedRuntimeVersionsResponse) XXX_Size() int {
	return xxx_messageInfo_ListSupportedRuntimeVersionsResponse.Size(m)
}
func (m *ListSupportedRuntimeVersionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSupportedRuntimeVersionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListSupportedRuntimeVersionsResponse proto.InternalMessageInfo

func (m *ListSupportedRuntimeVersionsResponse) GetSupportedVersions() []*ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion {
	if m != nil {
		return m.SupportedVersions
	}
	return nil
}

type ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion struct {
	RuntimeVersion       string               `protobuf:"bytes,1,opt,name=runtime_version,json=runtimeVersion,proto3" json:"runtime_version,omitempty"`
	ExpiryTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiry_time,json=expiryTime,proto3" json:"expiry_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) Reset() {
	*m = ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion{}
}
func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) String() string {
	return proto.CompactTextString(m)
}
func (*ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) ProtoMessage() {}
func (*ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_4224d18f2d1bfb4e, []int{1, 0}
}

func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion.Unmarshal(m, b)
}
func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion.Marshal(b, m, deterministic)
}
func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion.Merge(m, src)
}
func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) XXX_Size() int {
	return xxx_messageInfo_ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion.Size(m)
}
func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion.DiscardUnknown(m)
}

var xxx_messageInfo_ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion proto.InternalMessageInfo

func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) GetRuntimeVersion() string {
	if m != nil {
		return m.RuntimeVersion
	}
	return ""
}

func (m *ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion) GetExpiryTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiryTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ListSupportedRuntimeVersionsRequest)(nil), "improbable.spatialos.runtimeversion.v1alpha1.ListSupportedRuntimeVersionsRequest")
	proto.RegisterType((*ListSupportedRuntimeVersionsResponse)(nil), "improbable.spatialos.runtimeversion.v1alpha1.ListSupportedRuntimeVersionsResponse")
	proto.RegisterType((*ListSupportedRuntimeVersionsResponse_SupportedRuntimeVersion)(nil), "improbable.spatialos.runtimeversion.v1alpha1.ListSupportedRuntimeVersionsResponse.SupportedRuntimeVersion")
}

func init() {
	proto.RegisterFile("improbable/spatialos/runtimeversion/v1alpha1/runtime_version.proto", fileDescriptor_4224d18f2d1bfb4e)
}

var fileDescriptor_4224d18f2d1bfb4e = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0xc9, 0x04, 0xc1, 0x0c, 0x14, 0x03, 0xe2, 0x28, 0x82, 0x63, 0x2a, 0xee, 0x62, 0x9c,
	0xb2, 0x79, 0xb9, 0x2b, 0x77, 0x27, 0x08, 0x62, 0x2a, 0xbb, 0xf0, 0x66, 0xa4, 0x1a, 0x67, 0xa4,
	0x6d, 0x62, 0x92, 0x16, 0xbd, 0xf2, 0x39, 0x7c, 0x05, 0x5f, 0x43, 0x7c, 0x0e, 0x5f, 0x45, 0xda,
	0x98, 0x8d, 0x0a, 0x1b, 0x0e, 0x76, 0x7b, 0xfa, 0xff, 0xdf, 0xf9, 0xfb, 0x9f, 0xe0, 0x91, 0x48,
	0x95, 0x96, 0x31, 0x8b, 0x13, 0x1e, 0x1a, 0xc5, 0xac, 0x60, 0x89, 0x34, 0xa1, 0xce, 0x33, 0x2b,
	0x52, 0x5e, 0x70, 0x6d, 0x84, 0xcc, 0xc2, 0xa2, 0xcf, 0x12, 0xf5, 0xc8, 0xfa, 0x7e, 0x3e, 0xf9,
	0xfd, 0x00, 0x4a, 0x4b, 0x2b, 0x49, 0x6f, 0xce, 0x80, 0x19, 0x03, 0xea, 0x0c, 0xf0, 0x8c, 0xe0,
	0x70, 0x2a, 0xe5, 0x34, 0xe1, 0x61, 0xe5, 0x8d, 0xf3, 0x87, 0xb0, 0x54, 0x19, 0xcb, 0x52, 0xe5,
	0x70, 0x9d, 0x13, 0x7c, 0x74, 0x29, 0x8c, 0x8d, 0x72, 0xa5, 0xa4, 0xb6, 0xfc, 0x9e, 0x3a, 0xd0,
	0xd8, 0x81, 0x0c, 0xe5, 0xcf, 0x39, 0x37, 0xb6, 0xf3, 0xd9, 0xc0, 0xc7, 0xcb, 0x75, 0x46, 0xc9,
	0xcc, 0x70, 0xf2, 0x8e, 0x30, 0x31, 0x5e, 0xe4, 0xa3, 0x9b, 0x16, 0x6a, 0x6f, 0x74, 0x9b, 0x83,
	0x27, 0x58, 0x25, 0x3c, 0xfc, 0x67, 0x21, 0x2c, 0x10, 0xd0, 0xdd, 0x59, 0x0a, 0x6f, 0x09, 0xde,
	0xf0, 0xfe, 0x02, 0x35, 0x39, 0xc5, 0x3b, 0x7f, 0xea, 0x6e, 0xa1, 0x36, 0xea, 0x6e, 0xd1, 0x6d,
	0x5d, 0x17, 0x0e, 0x71, 0x93, 0xbf, 0x28, 0xa1, 0x5f, 0x27, 0xe5, 0xb4, 0xd5, 0x68, 0xa3, 0x6e,
	0x73, 0x10, 0x80, 0xab, 0x19, 0x7c, 0xcd, 0x70, 0xe3, 0x6b, 0xa6, 0xd8, 0xc9, 0xcb, 0xc1, 0xe0,
	0x1b, 0xe1, 0xbd, 0xfa, 0xe2, 0x88, 0xeb, 0x42, 0xdc, 0x71, 0xf2, 0x85, 0xf0, 0xc1, 0xb2, 0xdf,
	0x25, 0xd7, 0xeb, 0xac, 0xae, 0xba, 0x69, 0x40, 0xd7, 0x7f, 0x8d, 0x51, 0x72, 0x3b, 0x5c, 0xe5,
	0x8d, 0x0f, 0xeb, 0xf3, 0x8f, 0x46, 0xef, 0x62, 0x1e, 0x29, 0x72, 0xee, 0xab, 0x08, 0xea, 0xdb,
	0x60, 0xdc, 0x3f, 0xaf, 0xdc, 0xf1, 0x66, 0xd5, 0xf7, 0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1e, 0xc6, 0xb1, 0x6a, 0x58, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RuntimeVersionServiceClient is the client API for RuntimeVersionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeVersionServiceClient interface {
	ListSupportedRuntimeVersions(ctx context.Context, in *ListSupportedRuntimeVersionsRequest, opts ...grpc.CallOption) (*ListSupportedRuntimeVersionsResponse, error)
}

type runtimeVersionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuntimeVersionServiceClient(cc grpc.ClientConnInterface) RuntimeVersionServiceClient {
	return &runtimeVersionServiceClient{cc}
}

func (c *runtimeVersionServiceClient) ListSupportedRuntimeVersions(ctx context.Context, in *ListSupportedRuntimeVersionsRequest, opts ...grpc.CallOption) (*ListSupportedRuntimeVersionsResponse, error) {
	out := new(ListSupportedRuntimeVersionsResponse)
	err := c.cc.Invoke(ctx, "/improbable.spatialos.runtimeversion.v1alpha1.RuntimeVersionService/ListSupportedRuntimeVersions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeVersionServiceServer is the server API for RuntimeVersionService service.
type RuntimeVersionServiceServer interface {
	ListSupportedRuntimeVersions(context.Context, *ListSupportedRuntimeVersionsRequest) (*ListSupportedRuntimeVersionsResponse, error)
}

// UnimplementedRuntimeVersionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRuntimeVersionServiceServer struct {
}

func (*UnimplementedRuntimeVersionServiceServer) ListSupportedRuntimeVersions(ctx context.Context, req *ListSupportedRuntimeVersionsRequest) (*ListSupportedRuntimeVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSupportedRuntimeVersions not implemented")
}

func RegisterRuntimeVersionServiceServer(s *grpc.Server, srv RuntimeVersionServiceServer) {
	s.RegisterService(&_RuntimeVersionService_serviceDesc, srv)
}

func _RuntimeVersionService_ListSupportedRuntimeVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSupportedRuntimeVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeVersionServiceServer).ListSupportedRuntimeVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.spatialos.runtimeversion.v1alpha1.RuntimeVersionService/ListSupportedRuntimeVersions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeVersionServiceServer).ListSupportedRuntimeVersions(ctx, req.(*ListSupportedRuntimeVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeVersionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.spatialos.runtimeversion.v1alpha1.RuntimeVersionService",
	HandlerType: (*RuntimeVersionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSupportedRuntimeVersions",
			Handler:    _RuntimeVersionService_ListSupportedRuntimeVersions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/spatialos/runtimeversion/v1alpha1/runtime_version.proto",
}
