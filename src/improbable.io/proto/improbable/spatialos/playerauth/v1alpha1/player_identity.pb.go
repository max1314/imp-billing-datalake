// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/spatialos/playerauth/v1alpha1/player_identity.proto

package playerauth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
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

type PlayerToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerToken) Reset()         { *m = PlayerToken{} }
func (m *PlayerToken) String() string { return proto.CompactTextString(m) }
func (*PlayerToken) ProtoMessage()    {}
func (*PlayerToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_08aa3a30eb0b7ff6, []int{0}
}

func (m *PlayerToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerToken.Unmarshal(m, b)
}
func (m *PlayerToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerToken.Marshal(b, m, deterministic)
}
func (m *PlayerToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerToken.Merge(m, src)
}
func (m *PlayerToken) XXX_Size() int {
	return xxx_messageInfo_PlayerToken.Size(m)
}
func (m *PlayerToken) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerToken.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerToken proto.InternalMessageInfo

func (m *PlayerToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CreatePlayerTokenRequest struct {
	PlayerIdentifier     string   `protobuf:"bytes,1,opt,name=player_identifier,json=playerIdentifier,proto3" json:"player_identifier,omitempty"`
	ProjectName          string   `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePlayerTokenRequest) Reset()         { *m = CreatePlayerTokenRequest{} }
func (m *CreatePlayerTokenRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerTokenRequest) ProtoMessage()    {}
func (*CreatePlayerTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08aa3a30eb0b7ff6, []int{1}
}

func (m *CreatePlayerTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerTokenRequest.Unmarshal(m, b)
}
func (m *CreatePlayerTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreatePlayerTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerTokenRequest.Merge(m, src)
}
func (m *CreatePlayerTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerTokenRequest.Size(m)
}
func (m *CreatePlayerTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerTokenRequest proto.InternalMessageInfo

func (m *CreatePlayerTokenRequest) GetPlayerIdentifier() string {
	if m != nil {
		return m.PlayerIdentifier
	}
	return ""
}

func (m *CreatePlayerTokenRequest) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

type CreatePlayerTokenResponse struct {
	Token                *PlayerToken `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreatePlayerTokenResponse) Reset()         { *m = CreatePlayerTokenResponse{} }
func (m *CreatePlayerTokenResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePlayerTokenResponse) ProtoMessage()    {}
func (*CreatePlayerTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08aa3a30eb0b7ff6, []int{2}
}

func (m *CreatePlayerTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePlayerTokenResponse.Unmarshal(m, b)
}
func (m *CreatePlayerTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePlayerTokenResponse.Marshal(b, m, deterministic)
}
func (m *CreatePlayerTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePlayerTokenResponse.Merge(m, src)
}
func (m *CreatePlayerTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePlayerTokenResponse.Size(m)
}
func (m *CreatePlayerTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePlayerTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePlayerTokenResponse proto.InternalMessageInfo

func (m *CreatePlayerTokenResponse) GetToken() *PlayerToken {
	if m != nil {
		return m.Token
	}
	return nil
}

func init() {
	proto.RegisterType((*PlayerToken)(nil), "improbable.spatialos.playerauth.v1alpha1.PlayerToken")
	proto.RegisterType((*CreatePlayerTokenRequest)(nil), "improbable.spatialos.playerauth.v1alpha1.CreatePlayerTokenRequest")
	proto.RegisterType((*CreatePlayerTokenResponse)(nil), "improbable.spatialos.playerauth.v1alpha1.CreatePlayerTokenResponse")
}

func init() {
	proto.RegisterFile("improbable/spatialos/playerauth/v1alpha1/player_identity.proto", fileDescriptor_08aa3a30eb0b7ff6)
}

var fileDescriptor_08aa3a30eb0b7ff6 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x51, 0x8b, 0x13, 0x31,
	0x10, 0xc7, 0xe9, 0x8a, 0x82, 0xa9, 0x88, 0x17, 0x14, 0x6b, 0x1f, 0x4e, 0x59, 0x41, 0x8b, 0xd8,
	0xc4, 0xde, 0x21, 0x88, 0x07, 0xc2, 0xf6, 0xf0, 0xe1, 0x10, 0xf4, 0x68, 0x8f, 0x7b, 0x50, 0xa4,
	0x66, 0xb7, 0x73, 0xbb, 0xa9, 0xd9, 0x4d, 0xcc, 0x66, 0xab, 0x87, 0xf8, 0x2a, 0x7e, 0x06, 0x5f,
	0x0e, 0x7c, 0xd4, 0x6f, 0xe4, 0xd3, 0x7d, 0x14, 0x69, 0xb2, 0x76, 0x73, 0x54, 0xa1, 0x70, 0x6f,
	0xc9, 0xcc, 0xfc, 0x66, 0xfe, 0x33, 0x93, 0xa0, 0x67, 0x3c, 0x57, 0x5a, 0xc6, 0x2c, 0x16, 0x40,
	0x4b, 0xc5, 0x0c, 0x67, 0x42, 0x96, 0x54, 0x09, 0x76, 0x0c, 0x9a, 0x55, 0x26, 0xa3, 0xf3, 0x01,
	0x13, 0x2a, 0x63, 0x83, 0xda, 0x36, 0xe1, 0x53, 0x28, 0x0c, 0x37, 0xc7, 0x44, 0x69, 0x69, 0x24,
	0xee, 0x35, 0x3c, 0x59, 0xf2, 0xa4, 0xe1, 0xc9, 0x5f, 0xbe, 0xbb, 0x93, 0x72, 0x93, 0x55, 0x31,
	0x49, 0x64, 0x4e, 0x1b, 0xa8, 0xcf, 0x25, 0x4d, 0x65, 0xdf, 0x26, 0xeb, 0x0b, 0x99, 0x1e, 0x71,
	0x10, 0xd3, 0x92, 0x2e, 0x4f, 0xae, 0x4c, 0xf7, 0x66, 0x2a, 0x65, 0x2a, 0x80, 0x32, 0xc5, 0x69,
	0x22, 0x38, 0x14, 0xa6, 0x76, 0xdc, 0xf6, 0x1c, 0x96, 0x98, 0xc4, 0x90, 0xb1, 0x39, 0x97, 0xba,
	0x0e, 0xb8, 0xe7, 0x35, 0x08, 0x9f, 0x0c, 0x55, 0xa2, 0x4a, 0x79, 0x41, 0x6d, 0x6f, 0x0a, 0x74,
	0x5e, 0x57, 0x08, 0xef, 0xa2, 0xf6, 0xbe, 0x55, 0x7d, 0x20, 0xdf, 0x43, 0x81, 0xaf, 0xa3, 0x8b,
	0x66, 0x71, 0xe8, 0xb4, 0xee, 0xb4, 0x7a, 0x97, 0x47, 0xee, 0x12, 0x7e, 0x6d, 0xa1, 0xce, 0xae,
	0x06, 0x66, 0xc0, 0x8b, 0x1d, 0xc1, 0x87, 0x0a, 0x4a, 0x83, 0x1f, 0xa1, 0x8d, 0x33, 0x33, 0x3a,
	0xe2, 0xa0, 0x1d, 0x3e, 0xbc, 0x70, 0x1a, 0x05, 0xa3, 0x6b, 0xce, 0xbb, 0xb7, 0x74, 0xe2, 0x27,
	0xe8, 0x8a, 0xd2, 0x72, 0x06, 0x89, 0x99, 0x14, 0x2c, 0x87, 0x4e, 0x60, 0x83, 0x6f, 0x9c, 0x46,
	0xc1, 0x8f, 0x5f, 0x9b, 0x57, 0xcf, 0x3a, 0x47, 0xed, 0xfa, 0xf6, 0x92, 0xe5, 0x10, 0x66, 0xe8,
	0xd6, 0x3f, 0x74, 0x94, 0x4a, 0x16, 0x25, 0xe0, 0x17, 0xbe, 0xf6, 0xf6, 0xd6, 0x63, 0xb2, 0xee,
	0x8e, 0x88, 0x9f, 0xcd, 0xe5, 0xd8, 0x3a, 0x09, 0x10, 0xf6, 0xcc, 0x63, 0xd0, 0x73, 0x9e, 0x00,
	0xfe, 0x16, 0xa0, 0x8d, 0x15, 0x05, 0x78, 0xb8, 0x7e, 0xa9, 0xff, 0x8d, 0xb1, 0xbb, 0x7b, 0xae,
	0x1c, 0x6e, 0x04, 0xe1, 0xec, 0x77, 0xb4, 0xb9, 0xb2, 0x8c, 0x87, 0xfe, 0x3c, 0xbf, 0x9f, 0x84,
	0xcf, 0x51, 0xd7, 0xc5, 0x44, 0x95, 0xc9, 0x0e, 0x07, 0xfb, 0x55, 0x2c, 0x78, 0x52, 0xe7, 0x3c,
	0xc0, 0xf7, 0xdf, 0x7c, 0x7c, 0xfb, 0x54, 0xe9, 0x19, 0xfd, 0xcc, 0x74, 0x6a, 0x5f, 0xc9, 0x82,
	0xfd, 0x42, 0xa7, 0x4a, 0xd0, 0x07, 0xde, 0x0f, 0x19, 0xbe, 0x7b, 0xbd, 0xbd, 0xee, 0x27, 0xda,
	0x69, 0x6c, 0x3f, 0x83, 0xde, 0x5e, 0xd3, 0xe7, 0xd8, 0x51, 0xaf, 0xc6, 0xf5, 0x0a, 0x16, 0x8a,
	0xc8, 0xe1, 0x20, 0xb2, 0x54, 0x7c, 0xc9, 0x3e, 0xd1, 0xed, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0x48, 0x64, 0x38, 0xad, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PlayerTokenServiceClient is the client API for PlayerTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PlayerTokenServiceClient interface {
	CreatePlayerToken(ctx context.Context, in *CreatePlayerTokenRequest, opts ...grpc.CallOption) (*CreatePlayerTokenResponse, error)
}

type playerTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayerTokenServiceClient(cc grpc.ClientConnInterface) PlayerTokenServiceClient {
	return &playerTokenServiceClient{cc}
}

func (c *playerTokenServiceClient) CreatePlayerToken(ctx context.Context, in *CreatePlayerTokenRequest, opts ...grpc.CallOption) (*CreatePlayerTokenResponse, error) {
	out := new(CreatePlayerTokenResponse)
	err := c.cc.Invoke(ctx, "/improbable.spatialos.playerauth.v1alpha1.PlayerTokenService/CreatePlayerToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayerTokenServiceServer is the server API for PlayerTokenService service.
type PlayerTokenServiceServer interface {
	CreatePlayerToken(context.Context, *CreatePlayerTokenRequest) (*CreatePlayerTokenResponse, error)
}

// UnimplementedPlayerTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPlayerTokenServiceServer struct {
}

func (*UnimplementedPlayerTokenServiceServer) CreatePlayerToken(ctx context.Context, req *CreatePlayerTokenRequest) (*CreatePlayerTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayerToken not implemented")
}

func RegisterPlayerTokenServiceServer(s *grpc.Server, srv PlayerTokenServiceServer) {
	s.RegisterService(&_PlayerTokenService_serviceDesc, srv)
}

func _PlayerTokenService_CreatePlayerToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayerTokenServiceServer).CreatePlayerToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.spatialos.playerauth.v1alpha1.PlayerTokenService/CreatePlayerToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayerTokenServiceServer).CreatePlayerToken(ctx, req.(*CreatePlayerTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PlayerTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.spatialos.playerauth.v1alpha1.PlayerTokenService",
	HandlerType: (*PlayerTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlayerToken",
			Handler:    _PlayerTokenService_CreatePlayerToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/spatialos/playerauth/v1alpha1/player_identity.proto",
}
