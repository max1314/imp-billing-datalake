// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/int/platform/node/nanny_management.proto

package improbable_int_platform_node

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type NannyState int32

const (
	NannyState_UNKNOWN_NANNY_STATE NannyState = 0
	NannyState_PENDING             NannyState = 2
	NannyState_STARTED             NannyState = 3
	NannyState_STOPPED             NannyState = 6
)

var NannyState_name = map[int32]string{
	0: "UNKNOWN_NANNY_STATE",
	2: "PENDING",
	3: "STARTED",
	6: "STOPPED",
}

var NannyState_value = map[string]int32{
	"UNKNOWN_NANNY_STATE": 0,
	"PENDING":             2,
	"STARTED":             3,
	"STOPPED":             6,
}

func (x NannyState) String() string {
	return proto.EnumName(NannyState_name, int32(x))
}

func (NannyState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9752702a92329b00, []int{0}
}

type NannyStatusUpdate_StatusUpdateType int32

const (
	NannyStatusUpdate_UNKNOWN_STATUS_UPDATE_TYPE NannyStatusUpdate_StatusUpdateType = 0
	NannyStatusUpdate_INITIAL_STATE              NannyStatusUpdate_StatusUpdateType = 3
	NannyStatusUpdate_STATE_CHANGED              NannyStatusUpdate_StatusUpdateType = 1
	NannyStatusUpdate_ERROR                      NannyStatusUpdate_StatusUpdateType = 2
)

var NannyStatusUpdate_StatusUpdateType_name = map[int32]string{
	0: "UNKNOWN_STATUS_UPDATE_TYPE",
	3: "INITIAL_STATE",
	1: "STATE_CHANGED",
	2: "ERROR",
}

var NannyStatusUpdate_StatusUpdateType_value = map[string]int32{
	"UNKNOWN_STATUS_UPDATE_TYPE": 0,
	"INITIAL_STATE":              3,
	"STATE_CHANGED":              1,
	"ERROR":                      2,
}

func (x NannyStatusUpdate_StatusUpdateType) String() string {
	return proto.EnumName(NannyStatusUpdate_StatusUpdateType_name, int32(x))
}

func (NannyStatusUpdate_StatusUpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9752702a92329b00, []int{0, 0}
}

type NannyCommand_CommandType int32

const (
	NannyCommand_UNKNOWN NannyCommand_CommandType = 0
	NannyCommand_ASSIGN  NannyCommand_CommandType = 1
	NannyCommand_STOP    NannyCommand_CommandType = 2
	NannyCommand_CLEAR   NannyCommand_CommandType = 3
)

var NannyCommand_CommandType_name = map[int32]string{
	0: "UNKNOWN",
	1: "ASSIGN",
	2: "STOP",
	3: "CLEAR",
}

var NannyCommand_CommandType_value = map[string]int32{
	"UNKNOWN": 0,
	"ASSIGN":  1,
	"STOP":    2,
	"CLEAR":   3,
}

func (x NannyCommand_CommandType) String() string {
	return proto.EnumName(NannyCommand_CommandType_name, int32(x))
}

func (NannyCommand_CommandType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9752702a92329b00, []int{1, 0}
}

type NannyStatusUpdate struct {
	StatusUpdateType     NannyStatusUpdate_StatusUpdateType `protobuf:"varint,1,opt,name=status_update_type,json=statusUpdateType,proto3,enum=improbable.int.platform.node.NannyStatusUpdate_StatusUpdateType" json:"status_update_type,omitempty"`
	CurrentNannyState    NannyState                         `protobuf:"varint,2,opt,name=current_nanny_state,json=currentNannyState,proto3,enum=improbable.int.platform.node.NannyState" json:"current_nanny_state,omitempty"`
	PreviousNannyState   NannyState                         `protobuf:"varint,3,opt,name=previous_nanny_state,json=previousNannyState,proto3,enum=improbable.int.platform.node.NannyState" json:"previous_nanny_state,omitempty"`
	ErrorMessage         string                             `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *NannyStatusUpdate) Reset()         { *m = NannyStatusUpdate{} }
func (m *NannyStatusUpdate) String() string { return proto.CompactTextString(m) }
func (*NannyStatusUpdate) ProtoMessage()    {}
func (*NannyStatusUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_9752702a92329b00, []int{0}
}

func (m *NannyStatusUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NannyStatusUpdate.Unmarshal(m, b)
}
func (m *NannyStatusUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NannyStatusUpdate.Marshal(b, m, deterministic)
}
func (m *NannyStatusUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NannyStatusUpdate.Merge(m, src)
}
func (m *NannyStatusUpdate) XXX_Size() int {
	return xxx_messageInfo_NannyStatusUpdate.Size(m)
}
func (m *NannyStatusUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_NannyStatusUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_NannyStatusUpdate proto.InternalMessageInfo

func (m *NannyStatusUpdate) GetStatusUpdateType() NannyStatusUpdate_StatusUpdateType {
	if m != nil {
		return m.StatusUpdateType
	}
	return NannyStatusUpdate_UNKNOWN_STATUS_UPDATE_TYPE
}

func (m *NannyStatusUpdate) GetCurrentNannyState() NannyState {
	if m != nil {
		return m.CurrentNannyState
	}
	return NannyState_UNKNOWN_NANNY_STATE
}

func (m *NannyStatusUpdate) GetPreviousNannyState() NannyState {
	if m != nil {
		return m.PreviousNannyState
	}
	return NannyState_UNKNOWN_NANNY_STATE
}

func (m *NannyStatusUpdate) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type NannyCommand struct {
	CommandType            NannyCommand_CommandType `protobuf:"varint,1,opt,name=command_type,json=commandType,proto3,enum=improbable.int.platform.node.NannyCommand_CommandType" json:"command_type,omitempty"`
	DockerImageName        string                   `protobuf:"bytes,2,opt,name=docker_image_name,json=dockerImageName,proto3" json:"docker_image_name,omitempty"`
	DockerImageCmdlineArgs []string                 `protobuf:"bytes,3,rep,name=docker_image_cmdline_args,json=dockerImageCmdlineArgs,proto3" json:"docker_image_cmdline_args,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                 `json:"-"`
	XXX_unrecognized       []byte                   `json:"-"`
	XXX_sizecache          int32                    `json:"-"`
}

func (m *NannyCommand) Reset()         { *m = NannyCommand{} }
func (m *NannyCommand) String() string { return proto.CompactTextString(m) }
func (*NannyCommand) ProtoMessage()    {}
func (*NannyCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_9752702a92329b00, []int{1}
}

func (m *NannyCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NannyCommand.Unmarshal(m, b)
}
func (m *NannyCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NannyCommand.Marshal(b, m, deterministic)
}
func (m *NannyCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NannyCommand.Merge(m, src)
}
func (m *NannyCommand) XXX_Size() int {
	return xxx_messageInfo_NannyCommand.Size(m)
}
func (m *NannyCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_NannyCommand.DiscardUnknown(m)
}

var xxx_messageInfo_NannyCommand proto.InternalMessageInfo

func (m *NannyCommand) GetCommandType() NannyCommand_CommandType {
	if m != nil {
		return m.CommandType
	}
	return NannyCommand_UNKNOWN
}

func (m *NannyCommand) GetDockerImageName() string {
	if m != nil {
		return m.DockerImageName
	}
	return ""
}

func (m *NannyCommand) GetDockerImageCmdlineArgs() []string {
	if m != nil {
		return m.DockerImageCmdlineArgs
	}
	return nil
}

func init() {
	proto.RegisterEnum("improbable.int.platform.node.NannyState", NannyState_name, NannyState_value)
	proto.RegisterEnum("improbable.int.platform.node.NannyStatusUpdate_StatusUpdateType", NannyStatusUpdate_StatusUpdateType_name, NannyStatusUpdate_StatusUpdateType_value)
	proto.RegisterEnum("improbable.int.platform.node.NannyCommand_CommandType", NannyCommand_CommandType_name, NannyCommand_CommandType_value)
	proto.RegisterType((*NannyStatusUpdate)(nil), "improbable.int.platform.node.NannyStatusUpdate")
	proto.RegisterType((*NannyCommand)(nil), "improbable.int.platform.node.NannyCommand")
}

func init() {
	proto.RegisterFile("improbable/int/platform/node/nanny_management.proto", fileDescriptor_9752702a92329b00)
}

var fileDescriptor_9752702a92329b00 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x8e, 0xd2, 0x40,
	0x14, 0xc6, 0x29, 0x5d, 0x51, 0x0e, 0xac, 0x0e, 0xb3, 0x66, 0xc5, 0x8d, 0x31, 0x04, 0x6f, 0x08,
	0x17, 0xc5, 0xec, 0x26, 0x26, 0xc6, 0x1b, 0x1b, 0x68, 0xb0, 0x91, 0x1d, 0x9a, 0xb6, 0x44, 0xf1,
	0x66, 0x32, 0x94, 0x91, 0x34, 0x32, 0xd3, 0x66, 0x5a, 0x36, 0xe1, 0x05, 0xbc, 0xf3, 0xa9, 0x7c,
	0x31, 0xd3, 0x3f, 0x2b, 0xdd, 0x35, 0x31, 0xec, 0xd5, 0xcc, 0xf9, 0xce, 0x9c, 0xdf, 0xcc, 0x9c,
	0x2f, 0x07, 0xae, 0x42, 0x11, 0xab, 0x68, 0xc5, 0x56, 0x5b, 0x3e, 0x0a, 0x65, 0x3a, 0x8a, 0xb7,
	0x2c, 0xfd, 0x1e, 0x29, 0x31, 0x92, 0xd1, 0x9a, 0x8f, 0x24, 0x93, 0x72, 0x4f, 0x05, 0x93, 0x6c,
	0xc3, 0x05, 0x97, 0xa9, 0x11, 0xab, 0x28, 0x8d, 0xf0, 0xab, 0x43, 0x91, 0x11, 0x66, 0x6a, 0x59,
	0x64, 0x64, 0x45, 0xfd, 0xdf, 0x3a, 0x74, 0x48, 0x56, 0xe8, 0xa5, 0x2c, 0xdd, 0x25, 0x8b, 0x78,
	0xcd, 0x52, 0x8e, 0x25, 0xe0, 0x24, 0x8f, 0xe9, 0x2e, 0x17, 0x68, 0xba, 0x8f, 0x79, 0x57, 0xeb,
	0x69, 0x83, 0xa7, 0x97, 0x1f, 0x8d, 0xff, 0x01, 0x8d, 0x7f, 0x60, 0x46, 0x35, 0xf0, 0xf7, 0x31,
	0x77, 0x51, 0x72, 0x4f, 0xc1, 0x5f, 0xe1, 0x2c, 0xd8, 0x29, 0xc5, 0x65, 0x4a, 0x8b, 0x5f, 0x64,
	0x27, 0x78, 0xb7, 0x9e, 0x5f, 0x38, 0x38, 0xf2, 0x42, 0xee, 0x76, 0x4a, 0xc8, 0x41, 0xc2, 0xdf,
	0xe0, 0x79, 0xac, 0xf8, 0x4d, 0x18, 0xed, 0x92, 0x3b, 0x68, 0xfd, 0x81, 0x68, 0x7c, 0x4b, 0xa9,
	0xb0, 0xdf, 0xc0, 0x29, 0x57, 0x2a, 0x52, 0x54, 0xf0, 0x24, 0x61, 0x1b, 0xde, 0x3d, 0xe9, 0x69,
	0x83, 0xa6, 0xdb, 0xce, 0xc5, 0xeb, 0x42, 0xeb, 0x07, 0x80, 0xee, 0x37, 0x00, 0xbf, 0x86, 0x8b,
	0x05, 0xf9, 0x4c, 0xe6, 0x5f, 0x08, 0xf5, 0x7c, 0xd3, 0x5f, 0x78, 0x74, 0xe1, 0x4c, 0x4c, 0xdf,
	0xa2, 0xfe, 0xd2, 0xb1, 0x50, 0x0d, 0x77, 0xe0, 0xd4, 0x26, 0xb6, 0x6f, 0x9b, 0xb3, 0x3c, 0x6f,
	0x21, 0x3d, 0x93, 0xf2, 0x2d, 0x1d, 0x7f, 0x32, 0xc9, 0xd4, 0x9a, 0x20, 0x0d, 0x37, 0xe1, 0x91,
	0xe5, 0xba, 0x73, 0x17, 0xd5, 0xfb, 0xbf, 0xea, 0xd0, 0xce, 0x1f, 0x36, 0x8e, 0x84, 0x60, 0x72,
	0x8d, 0x97, 0xd0, 0x0e, 0x8a, 0x6d, 0xd5, 0xba, 0x77, 0x47, 0x7c, 0xb7, 0x24, 0x18, 0xe5, 0x9a,
	0x1b, 0xd6, 0x0a, 0x0e, 0x01, 0x1e, 0x42, 0x67, 0x1d, 0x05, 0x3f, 0xb8, 0xa2, 0xa1, 0x60, 0x1b,
	0x4e, 0x25, 0x13, 0x85, 0x53, 0x4d, 0xf7, 0x59, 0x91, 0xb0, 0x33, 0x9d, 0x30, 0xc1, 0xf1, 0x7b,
	0x78, 0x79, 0xe7, 0x6c, 0x20, 0xd6, 0xdb, 0x50, 0x72, 0xca, 0xd4, 0x26, 0xe9, 0xea, 0x3d, 0x7d,
	0xd0, 0x74, 0xcf, 0x2b, 0x35, 0xe3, 0x22, 0x6d, 0xaa, 0x4d, 0xd2, 0xff, 0x00, 0xad, 0xca, 0x13,
	0x70, 0x0b, 0x1e, 0x97, 0x2d, 0x43, 0x35, 0x0c, 0xd0, 0x30, 0x3d, 0xcf, 0x9e, 0x12, 0xa4, 0xe1,
	0x27, 0x70, 0xe2, 0xf9, 0x73, 0x07, 0xd5, 0xb3, 0x7e, 0x8c, 0x67, 0x96, 0xe9, 0x22, 0x7d, 0x38,
	0x03, 0xa8, 0xf8, 0xf4, 0x02, 0xce, 0x6e, 0xdb, 0x4d, 0x4c, 0x42, 0x96, 0x65, 0x53, 0x6b, 0x19,
	0xd4, 0xb1, 0xc8, 0xc4, 0x26, 0x53, 0x54, 0xcf, 0x02, 0xcf, 0x37, 0x5d, 0xdf, 0x9a, 0x20, 0xbd,
	0x08, 0xe6, 0x8e, 0x63, 0x4d, 0x50, 0xe3, 0xf2, 0xa7, 0x06, 0xe7, 0x39, 0xee, 0xfa, 0xef, 0x6c,
	0x79, 0x5c, 0xdd, 0x84, 0x01, 0xc7, 0x5b, 0x68, 0x7a, 0xbb, 0x55, 0x12, 0xa8, 0x70, 0xc5, 0xf1,
	0xe8, 0x81, 0x93, 0x71, 0x31, 0x3c, 0xde, 0x8f, 0x81, 0xf6, 0x56, 0x5b, 0x35, 0xf2, 0x89, 0xbe,
	0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x8c, 0xf2, 0xc8, 0x08, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NannyManagementServiceClient is the client API for NannyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NannyManagementServiceClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (NannyManagementService_SubscribeClient, error)
}

type nannyManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNannyManagementServiceClient(cc grpc.ClientConnInterface) NannyManagementServiceClient {
	return &nannyManagementServiceClient{cc}
}

func (c *nannyManagementServiceClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (NannyManagementService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NannyManagementService_serviceDesc.Streams[0], "/improbable.int.platform.node.NannyManagementService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &nannyManagementServiceSubscribeClient{stream}
	return x, nil
}

type NannyManagementService_SubscribeClient interface {
	Send(*NannyStatusUpdate) error
	Recv() (*NannyCommand, error)
	grpc.ClientStream
}

type nannyManagementServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *nannyManagementServiceSubscribeClient) Send(m *NannyStatusUpdate) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nannyManagementServiceSubscribeClient) Recv() (*NannyCommand, error) {
	m := new(NannyCommand)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NannyManagementServiceServer is the server API for NannyManagementService service.
type NannyManagementServiceServer interface {
	Subscribe(NannyManagementService_SubscribeServer) error
}

// UnimplementedNannyManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNannyManagementServiceServer struct {
}

func (*UnimplementedNannyManagementServiceServer) Subscribe(srv NannyManagementService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterNannyManagementServiceServer(s *grpc.Server, srv NannyManagementServiceServer) {
	s.RegisterService(&_NannyManagementService_serviceDesc, srv)
}

func _NannyManagementService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NannyManagementServiceServer).Subscribe(&nannyManagementServiceSubscribeServer{stream})
}

type NannyManagementService_SubscribeServer interface {
	Send(*NannyCommand) error
	Recv() (*NannyStatusUpdate, error)
	grpc.ServerStream
}

type nannyManagementServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *nannyManagementServiceSubscribeServer) Send(m *NannyCommand) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nannyManagementServiceSubscribeServer) Recv() (*NannyStatusUpdate, error) {
	m := new(NannyStatusUpdate)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _NannyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.int.platform.node.NannyManagementService",
	HandlerType: (*NannyManagementServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _NannyManagementService_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "improbable/int/platform/node/nanny_management.proto",
}