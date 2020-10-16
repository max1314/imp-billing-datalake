// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/fabric/devtools/service/release_tool/fabric_release.proto

package improbable_fabric_devtools_service_release_tool

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type FabricEnvironment int32

const (
	FabricEnvironment_UNKNOWN_FABRIC_ENVIRONMENT FabricEnvironment = 0
	FabricEnvironment_PRODUCTION                 FabricEnvironment = 1
	FabricEnvironment_STAGING                    FabricEnvironment = 2
	FabricEnvironment_TESTING                    FabricEnvironment = 3
)

var FabricEnvironment_name = map[int32]string{
	0: "UNKNOWN_FABRIC_ENVIRONMENT",
	1: "PRODUCTION",
	2: "STAGING",
	3: "TESTING",
}

var FabricEnvironment_value = map[string]int32{
	"UNKNOWN_FABRIC_ENVIRONMENT": 0,
	"PRODUCTION":                 1,
	"STAGING":                    2,
	"TESTING":                    3,
}

func (x FabricEnvironment) String() string {
	return proto.EnumName(FabricEnvironment_name, int32(x))
}

func (FabricEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0066bfa821fc7e49, []int{0}
}

type RolloutType int32

const (
	RolloutType_UNKNOWN_ROLLOUT_TYPE RolloutType = 0
	RolloutType_RELEASE              RolloutType = 1
	RolloutType_ROLLBACK             RolloutType = 2
)

var RolloutType_name = map[int32]string{
	0: "UNKNOWN_ROLLOUT_TYPE",
	1: "RELEASE",
	2: "ROLLBACK",
}

var RolloutType_value = map[string]int32{
	"UNKNOWN_ROLLOUT_TYPE": 0,
	"RELEASE":              1,
	"ROLLBACK":             2,
}

func (x RolloutType) String() string {
	return proto.EnumName(RolloutType_name, int32(x))
}

func (RolloutType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0066bfa821fc7e49, []int{1}
}

type ReleaseRequest struct {
	Environment          FabricEnvironment `protobuf:"varint,1,opt,name=environment,proto3,enum=improbable.fabric.devtools.service.release_tool.FabricEnvironment" json:"environment,omitempty"`
	RolloutType          RolloutType       `protobuf:"varint,2,opt,name=rollout_type,json=rolloutType,proto3,enum=improbable.fabric.devtools.service.release_tool.RolloutType" json:"rollout_type,omitempty"`
	BuildInfo            *FabricVersion    `protobuf:"bytes,3,opt,name=build_info,json=buildInfo,proto3" json:"build_info,omitempty"`
	ReleaseNotes         string            `protobuf:"bytes,4,opt,name=release_notes,json=releaseNotes,proto3" json:"release_notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ReleaseRequest) Reset()         { *m = ReleaseRequest{} }
func (m *ReleaseRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseRequest) ProtoMessage()    {}
func (*ReleaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0066bfa821fc7e49, []int{0}
}

func (m *ReleaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseRequest.Unmarshal(m, b)
}
func (m *ReleaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseRequest.Merge(m, src)
}
func (m *ReleaseRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseRequest.Size(m)
}
func (m *ReleaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseRequest proto.InternalMessageInfo

func (m *ReleaseRequest) GetEnvironment() FabricEnvironment {
	if m != nil {
		return m.Environment
	}
	return FabricEnvironment_UNKNOWN_FABRIC_ENVIRONMENT
}

func (m *ReleaseRequest) GetRolloutType() RolloutType {
	if m != nil {
		return m.RolloutType
	}
	return RolloutType_UNKNOWN_ROLLOUT_TYPE
}

func (m *ReleaseRequest) GetBuildInfo() *FabricVersion {
	if m != nil {
		return m.BuildInfo
	}
	return nil
}

func (m *ReleaseRequest) GetReleaseNotes() string {
	if m != nil {
		return m.ReleaseNotes
	}
	return ""
}

type ReleaseResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseResponse) Reset()         { *m = ReleaseResponse{} }
func (m *ReleaseResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseResponse) ProtoMessage()    {}
func (*ReleaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0066bfa821fc7e49, []int{1}
}

func (m *ReleaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseResponse.Unmarshal(m, b)
}
func (m *ReleaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseResponse.Merge(m, src)
}
func (m *ReleaseResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseResponse.Size(m)
}
func (m *ReleaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("improbable.fabric.devtools.service.release_tool.FabricEnvironment", FabricEnvironment_name, FabricEnvironment_value)
	proto.RegisterEnum("improbable.fabric.devtools.service.release_tool.RolloutType", RolloutType_name, RolloutType_value)
	proto.RegisterType((*ReleaseRequest)(nil), "improbable.fabric.devtools.service.release_tool.ReleaseRequest")
	proto.RegisterType((*ReleaseResponse)(nil), "improbable.fabric.devtools.service.release_tool.ReleaseResponse")
}

func init() {
	proto.RegisterFile("improbable/fabric/devtools/service/release_tool/fabric_release.proto", fileDescriptor_0066bfa821fc7e49)
}

var fileDescriptor_0066bfa821fc7e49 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x6e, 0xd3, 0x40,
	0x18, 0x85, 0x6b, 0x17, 0xb5, 0xf4, 0x4f, 0x08, 0xee, 0xa8, 0x0b, 0x2b, 0x0b, 0x88, 0xca, 0x26,
	0xaa, 0x54, 0x5b, 0x0a, 0x12, 0x2b, 0x04, 0xc4, 0xa9, 0x5b, 0x59, 0x0d, 0x76, 0x35, 0x71, 0x8a,
	0x58, 0x54, 0x96, 0x9d, 0x4c, 0xc2, 0x08, 0xc7, 0x63, 0x66, 0x26, 0xae, 0x7a, 0x07, 0xae, 0xc1,
	0xb9, 0x90, 0xd8, 0x72, 0x09, 0x64, 0x8f, 0xd3, 0x04, 0x75, 0x95, 0xec, 0x3c, 0xcf, 0xcf, 0xdf,
	0xfb, 0xe7, 0xf9, 0x87, 0x0b, 0xba, 0xc8, 0x39, 0x4b, 0xe2, 0x24, 0x25, 0xf6, 0x2c, 0x4e, 0x38,
	0x9d, 0xd8, 0x53, 0x52, 0x48, 0xc6, 0x52, 0x61, 0x0b, 0xc2, 0x0b, 0x3a, 0x21, 0x36, 0x27, 0x29,
	0x89, 0x05, 0x89, 0x4a, 0xb5, 0x36, 0x45, 0xb5, 0x66, 0xe5, 0x9c, 0x49, 0x86, 0xec, 0x35, 0xc5,
	0x52, 0x06, 0x6b, 0x45, 0xb1, 0x6a, 0x8a, 0xb5, 0x49, 0x69, 0xbf, 0x9b, 0x53, 0xf9, 0x6d, 0x99,
	0x58, 0x13, 0xb6, 0xb0, 0x17, 0xf7, 0x54, 0x7e, 0x67, 0xf7, 0xf6, 0x9c, 0x9d, 0x57, 0xb4, 0xf3,
	0x22, 0x4e, 0xe9, 0x34, 0x96, 0x8c, 0x0b, 0xfb, 0xf1, 0x51, 0x05, 0xb5, 0x77, 0x1d, 0xb7, 0x20,
	0x5c, 0x50, 0x96, 0x29, 0xca, 0xe9, 0x5f, 0x1d, 0x5a, 0x58, 0xb9, 0x30, 0xf9, 0xb1, 0x24, 0x42,
	0xa2, 0x14, 0x1a, 0x24, 0x2b, 0x28, 0x67, 0xd9, 0x82, 0x64, 0xd2, 0xd4, 0x3a, 0x5a, 0xb7, 0xd5,
	0x73, 0xac, 0x2d, 0xef, 0x65, 0x5d, 0x56, 0x26, 0x77, 0x4d, 0x72, 0x0e, 0xfe, 0xfc, 0x7e, 0xad,
	0x77, 0x34, 0xbc, 0x89, 0x47, 0x73, 0x68, 0x72, 0x96, 0xa6, 0x6c, 0x29, 0x23, 0xf9, 0x90, 0x13,
	0x53, 0xaf, 0xe2, 0xde, 0x6f, 0x1d, 0x87, 0x15, 0x24, 0x7c, 0xc8, 0xc9, 0x3a, 0x88, 0xaf, 0x45,
	0x74, 0x07, 0x90, 0x2c, 0x69, 0x3a, 0x8d, 0x68, 0x36, 0x63, 0xe6, 0x7e, 0x47, 0xeb, 0x36, 0x7a,
	0x1f, 0x76, 0xbc, 0xd5, 0xad, 0xea, 0x10, 0x1f, 0x55, 0x44, 0x2f, 0x9b, 0x31, 0xf4, 0x06, 0x5e,
	0xac, 0x8c, 0x19, 0x93, 0x44, 0x98, 0xcf, 0x3a, 0x5a, 0xf7, 0x08, 0x37, 0x6b, 0xd1, 0x2f, 0xb5,
	0xd3, 0x63, 0x78, 0xf9, 0x58, 0xb6, 0xc8, 0x59, 0x26, 0xc8, 0xd9, 0x1d, 0x1c, 0x3f, 0x69, 0x0a,
	0xbd, 0x82, 0xf6, 0xd8, 0xbf, 0xf6, 0x83, 0x2f, 0x7e, 0x74, 0xd9, 0x77, 0xb0, 0x37, 0x88, 0x5c,
	0xff, 0xd6, 0xc3, 0x81, 0xff, 0xd9, 0xf5, 0x43, 0x63, 0x0f, 0xb5, 0x00, 0x6e, 0x70, 0x70, 0x31,
	0x1e, 0x84, 0x5e, 0xe0, 0x1b, 0x1a, 0x6a, 0xc0, 0xe1, 0x28, 0xec, 0x5f, 0x79, 0xfe, 0x95, 0xa1,
	0x97, 0x87, 0xd0, 0x1d, 0x85, 0xe5, 0x61, 0xff, 0xcc, 0x81, 0xc6, 0x46, 0x33, 0xc8, 0x84, 0x93,
	0x15, 0x18, 0x07, 0xc3, 0x61, 0x30, 0x0e, 0xa3, 0xf0, 0xeb, 0x8d, 0x6b, 0xec, 0x95, 0x5f, 0x61,
	0x77, 0xe8, 0xf6, 0x47, 0xae, 0xa1, 0xa1, 0x26, 0x3c, 0x2f, 0x5f, 0x3b, 0xfd, 0xc1, 0xb5, 0xa1,
	0xf7, 0x7e, 0x69, 0x70, 0xa2, 0x66, 0xac, 0x87, 0x1f, 0xa9, 0x66, 0xd0, 0x4f, 0x0d, 0x0e, 0x6b,
	0x09, 0x7d, 0xdc, 0xfe, 0x8f, 0xfd, 0xb7, 0x76, 0xed, 0x4f, 0xbb, 0x03, 0x54, 0x95, 0xc9, 0x41,
	0xb5, 0xd2, 0x6f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xb9, 0xe4, 0xa5, 0xc9, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FabricReleaseServiceClient is the client API for FabricReleaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FabricReleaseServiceClient interface {
	Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
}

type fabricReleaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFabricReleaseServiceClient(cc grpc.ClientConnInterface) FabricReleaseServiceClient {
	return &fabricReleaseServiceClient{cc}
}

func (c *fabricReleaseServiceClient) Release(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, "/improbable.fabric.devtools.service.release_tool.FabricReleaseService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FabricReleaseServiceServer is the server API for FabricReleaseService service.
type FabricReleaseServiceServer interface {
	Release(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
}

// UnimplementedFabricReleaseServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFabricReleaseServiceServer struct {
}

func (*UnimplementedFabricReleaseServiceServer) Release(ctx context.Context, req *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}

func RegisterFabricReleaseServiceServer(s *grpc.Server, srv FabricReleaseServiceServer) {
	s.RegisterService(&_FabricReleaseService_serviceDesc, srv)
}

func _FabricReleaseService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FabricReleaseServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.fabric.devtools.service.release_tool.FabricReleaseService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FabricReleaseServiceServer).Release(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FabricReleaseService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.fabric.devtools.service.release_tool.FabricReleaseService",
	HandlerType: (*FabricReleaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Release",
			Handler:    _FabricReleaseService_Release_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/fabric/devtools/service/release_tool/fabric_release.proto",
}
