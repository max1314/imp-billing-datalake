// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/serverstatus.proto

package base

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type HealthCheckResponse struct {
	IsOk                 bool     `protobuf:"varint,1,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d79a7fe32854758d, []int{0}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

type FlagzState struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Help                 string   `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	CurrentValue         string   `protobuf:"bytes,3,opt,name=current_value,json=currentValue,proto3" json:"current_value,omitempty"`
	DefaultValue         string   `protobuf:"bytes,4,opt,name=default_value,json=defaultValue,proto3" json:"default_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlagzState) Reset()         { *m = FlagzState{} }
func (m *FlagzState) String() string { return proto.CompactTextString(m) }
func (*FlagzState) ProtoMessage()    {}
func (*FlagzState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d79a7fe32854758d, []int{1}
}

func (m *FlagzState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlagzState.Unmarshal(m, b)
}
func (m *FlagzState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlagzState.Marshal(b, m, deterministic)
}
func (m *FlagzState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlagzState.Merge(m, src)
}
func (m *FlagzState) XXX_Size() int {
	return xxx_messageInfo_FlagzState.Size(m)
}
func (m *FlagzState) XXX_DiscardUnknown() {
	xxx_messageInfo_FlagzState.DiscardUnknown(m)
}

var xxx_messageInfo_FlagzState proto.InternalMessageInfo

func (m *FlagzState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FlagzState) GetHelp() string {
	if m != nil {
		return m.Help
	}
	return ""
}

func (m *FlagzState) GetCurrentValue() string {
	if m != nil {
		return m.CurrentValue
	}
	return ""
}

func (m *FlagzState) GetDefaultValue() string {
	if m != nil {
		return m.DefaultValue
	}
	return ""
}

type VersionResponse struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Branchname           string   `protobuf:"bytes,2,opt,name=branchname,proto3" json:"branchname,omitempty"`
	Date                 string   `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Go                   string   `protobuf:"bytes,4,opt,name=go,proto3" json:"go,omitempty"`
	Epoch                string   `protobuf:"bytes,5,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Tag                  string   `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d79a7fe32854758d, []int{2}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *VersionResponse) GetBranchname() string {
	if m != nil {
		return m.Branchname
	}
	return ""
}

func (m *VersionResponse) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *VersionResponse) GetGo() string {
	if m != nil {
		return m.Go
	}
	return ""
}

func (m *VersionResponse) GetEpoch() string {
	if m != nil {
		return m.Epoch
	}
	return ""
}

func (m *VersionResponse) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheckResponse)(nil), "base.HealthCheckResponse")
	proto.RegisterType((*FlagzState)(nil), "base.FlagzState")
	proto.RegisterType((*VersionResponse)(nil), "base.VersionResponse")
}

func init() {
	proto.RegisterFile("base/serverstatus.proto", fileDescriptor_d79a7fe32854758d)
}

var fileDescriptor_d79a7fe32854758d = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x4d, 0x8b, 0xd4, 0x40,
	0x10, 0x25, 0xd9, 0xcc, 0xae, 0xa9, 0x1d, 0xdd, 0xb5, 0xd6, 0x8f, 0x38, 0x8a, 0x48, 0xbc, 0x88,
	0x87, 0x44, 0xf4, 0x27, 0x88, 0xb2, 0x87, 0x05, 0x61, 0x02, 0x7b, 0x0d, 0x9d, 0x4c, 0x4d, 0x12,
	0x26, 0xd3, 0x1d, 0xd2, 0x9d, 0x01, 0xe7, 0xe0, 0xc1, 0x7f, 0x20, 0xfe, 0x34, 0xff, 0x82, 0xff,
	0xc2, 0x8b, 0x74, 0x75, 0xa2, 0x83, 0x38, 0xb7, 0xea, 0xf7, 0x5e, 0xbf, 0x7a, 0xbc, 0x82, 0xc7,
	0x85, 0xd0, 0x94, 0x6a, 0xea, 0x77, 0xd4, 0x6b, 0x23, 0xcc, 0xa0, 0x93, 0xae, 0x57, 0x46, 0x61,
	0x60, 0x89, 0xc5, 0xb3, 0x4a, 0xa9, 0xaa, 0xa5, 0x54, 0x74, 0x4d, 0x2a, 0xa4, 0x54, 0x46, 0x98,
	0x46, 0xc9, 0x51, 0xb3, 0x78, 0x3a, 0xb2, 0xfc, 0x2a, 0x86, 0x75, 0x4a, 0xdb, 0xce, 0x7c, 0x76,
	0x64, 0xfc, 0x1a, 0xae, 0xae, 0x49, 0xb4, 0xa6, 0x7e, 0x5f, 0x53, 0xb9, 0x59, 0x92, 0xee, 0x94,
	0xd4, 0x84, 0x57, 0x30, 0x6b, 0x74, 0xae, 0x36, 0x91, 0xf7, 0xc2, 0x7b, 0x75, 0x67, 0x19, 0x34,
	0xfa, 0xd3, 0x26, 0xfe, 0x02, 0xf0, 0xb1, 0x15, 0xd5, 0x3e, 0x33, 0xc2, 0x10, 0x22, 0x04, 0x52,
	0x6c, 0x89, 0x15, 0xe1, 0x92, 0x67, 0x8b, 0xd5, 0xd4, 0x76, 0x91, 0xef, 0x30, 0x3b, 0xe3, 0x4b,
	0xb8, 0x5b, 0x0e, 0x7d, 0x4f, 0xd2, 0xe4, 0x3b, 0xd1, 0x0e, 0x14, 0x9d, 0x30, 0x39, 0x1f, 0xc1,
	0x5b, 0x8b, 0x59, 0xd1, 0x8a, 0xd6, 0x62, 0x68, 0x27, 0x51, 0xe0, 0x44, 0x23, 0xc8, 0xa2, 0xf8,
	0x9b, 0x07, 0x17, 0xb7, 0xd4, 0xeb, 0x46, 0xc9, 0x3f, 0x41, 0xed, 0x46, 0xa1, 0xeb, 0x29, 0x85,
	0x9d, 0xf1, 0x39, 0x40, 0xd1, 0x0b, 0x59, 0xd6, 0x9c, 0xcf, 0x65, 0x39, 0x40, 0xec, 0x9f, 0x95,
	0x30, 0x53, 0x10, 0x9e, 0xf1, 0x1e, 0xf8, 0x95, 0x1a, 0xb7, 0xfa, 0x95, 0xc2, 0x07, 0x30, 0xa3,
	0x4e, 0x95, 0x75, 0x34, 0x63, 0xc8, 0x3d, 0xf0, 0x12, 0x4e, 0x8c, 0xa8, 0xa2, 0x53, 0xc6, 0xec,
	0xf8, 0xf6, 0x97, 0x07, 0xf3, 0x8c, 0xef, 0x92, 0xf1, 0x5d, 0x30, 0x83, 0xf3, 0x83, 0x42, 0xf1,
	0x51, 0xe2, 0xda, 0x4f, 0xa6, 0xf6, 0x93, 0x0f, 0xb6, 0xfd, 0xc5, 0x93, 0xc4, 0x5e, 0x2e, 0xf9,
	0x4f, 0xf7, 0xf1, 0xfd, 0xaf, 0x3f, 0x7e, 0x7e, 0xf7, 0xcf, 0x31, 0x4c, 0xf3, 0x9a, 0xe9, 0x3d,
	0x5e, 0x43, 0xc8, 0xcd, 0xdf, 0x34, 0xda, 0x1c, 0xb5, 0xbc, 0x74, 0x96, 0x7f, 0x4f, 0x14, 0x5f,
	0xb0, 0x53, 0x88, 0x67, 0x69, 0xbe, 0xb6, 0xe8, 0x1b, 0x0f, 0x6f, 0xe0, 0x6c, 0xac, 0xf0, 0xa8,
	0xcf, 0x43, 0xe7, 0xf3, 0x4f, 0xd3, 0x31, 0xb2, 0xd9, 0x1c, 0x21, 0xcd, 0x77, 0x8e, 0xda, 0x17,
	0xa7, 0xfc, 0xf5, 0xdd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x72, 0xa8, 0x1d, 0xa0, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServerStatusClient is the client API for ServerStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServerStatusClient interface {
	HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	FlagzList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ServerStatus_FlagzListClient, error)
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
}

type serverStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStatusClient(cc grpc.ClientConnInterface) ServerStatusClient {
	return &serverStatusClient{cc}
}

func (c *serverStatusClient) HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/base.ServerStatus/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverStatusClient) FlagzList(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ServerStatus_FlagzListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerStatus_serviceDesc.Streams[0], "/base.ServerStatus/FlagzList", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverStatusFlagzListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStatus_FlagzListClient interface {
	Recv() (*FlagzState, error)
	grpc.ClientStream
}

type serverStatusFlagzListClient struct {
	grpc.ClientStream
}

func (x *serverStatusFlagzListClient) Recv() (*FlagzState, error) {
	m := new(FlagzState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serverStatusClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/base.ServerStatus/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerStatusServer is the server API for ServerStatus service.
type ServerStatusServer interface {
	HealthCheck(context.Context, *empty.Empty) (*HealthCheckResponse, error)
	FlagzList(*empty.Empty, ServerStatus_FlagzListServer) error
	Version(context.Context, *empty.Empty) (*VersionResponse, error)
}

// UnimplementedServerStatusServer can be embedded to have forward compatible implementations.
type UnimplementedServerStatusServer struct {
}

func (*UnimplementedServerStatusServer) HealthCheck(ctx context.Context, req *empty.Empty) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (*UnimplementedServerStatusServer) FlagzList(req *empty.Empty, srv ServerStatus_FlagzListServer) error {
	return status.Errorf(codes.Unimplemented, "method FlagzList not implemented")
}
func (*UnimplementedServerStatusServer) Version(ctx context.Context, req *empty.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}

func RegisterServerStatusServer(s *grpc.Server, srv ServerStatusServer) {
	s.RegisterService(&_ServerStatus_serviceDesc, srv)
}

func _ServerStatus_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatusServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.ServerStatus/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatusServer).HealthCheck(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerStatus_FlagzList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStatusServer).FlagzList(m, &serverStatusFlagzListServer{stream})
}

type ServerStatus_FlagzListServer interface {
	Send(*FlagzState) error
	grpc.ServerStream
}

type serverStatusFlagzListServer struct {
	grpc.ServerStream
}

func (x *serverStatusFlagzListServer) Send(m *FlagzState) error {
	return x.ServerStream.SendMsg(m)
}

func _ServerStatus_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerStatusServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.ServerStatus/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerStatusServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServerStatus_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.ServerStatus",
	HandlerType: (*ServerStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _ServerStatus_HealthCheck_Handler,
		},
		{
			MethodName: "Version",
			Handler:    _ServerStatus_Version_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FlagzList",
			Handler:       _ServerStatus_FlagzList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "base/serverstatus.proto",
}