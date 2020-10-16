// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/platform/logging/logging.proto

package improbable_platform_logging

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

type LogMessage_LogLevel int32

const (
	LogMessage_PANIC LogMessage_LogLevel = 0
	LogMessage_FATAL LogMessage_LogLevel = 1
	LogMessage_ERROR LogMessage_LogLevel = 2
	LogMessage_WARN  LogMessage_LogLevel = 3
	LogMessage_INFO  LogMessage_LogLevel = 4
	LogMessage_DEBUG LogMessage_LogLevel = 5
)

var LogMessage_LogLevel_name = map[int32]string{
	0: "PANIC",
	1: "FATAL",
	2: "ERROR",
	3: "WARN",
	4: "INFO",
	5: "DEBUG",
}

var LogMessage_LogLevel_value = map[string]int32{
	"PANIC": 0,
	"FATAL": 1,
	"ERROR": 2,
	"WARN":  3,
	"INFO":  4,
	"DEBUG": 5,
}

func (x LogMessage_LogLevel) String() string {
	return proto.EnumName(LogMessage_LogLevel_name, int32(x))
}

func (LogMessage_LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4dc0a28404c07b33, []int{0, 0}
}

type LogMessage struct {
	Level                LogMessage_LogLevel `protobuf:"varint,1,opt,name=level,proto3,enum=improbable.platform.logging.LogMessage_LogLevel" json:"level,omitempty"`
	Message              string              `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Fields               map[string]string   `protobuf:"bytes,3,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	StackTrace           string              `protobuf:"bytes,4,opt,name=stack_trace,json=stackTrace,proto3" json:"stack_trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc0a28404c07b33, []int{0}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetLevel() LogMessage_LogLevel {
	if m != nil {
		return m.Level
	}
	return LogMessage_PANIC
}

func (m *LogMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogMessage) GetFields() map[string]string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *LogMessage) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

type LogRequest struct {
	ClientType           string      `protobuf:"bytes,1,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty"`
	Message              *LogMessage `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc0a28404c07b33, []int{1}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *LogRequest) GetMessage() *LogMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func init() {
	proto.RegisterEnum("improbable.platform.logging.LogMessage_LogLevel", LogMessage_LogLevel_name, LogMessage_LogLevel_value)
	proto.RegisterType((*LogMessage)(nil), "improbable.platform.logging.LogMessage")
	proto.RegisterMapType((map[string]string)(nil), "improbable.platform.logging.LogMessage.FieldsEntry")
	proto.RegisterType((*LogRequest)(nil), "improbable.platform.logging.LogRequest")
}

func init() {
	proto.RegisterFile("improbable/platform/logging/logging.proto", fileDescriptor_4dc0a28404c07b33)
}

var fileDescriptor_4dc0a28404c07b33 = []byte{
	// 492 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x6f, 0x12, 0x41,
	0x14, 0x75, 0x59, 0x68, 0xcb, 0xc5, 0xe0, 0x66, 0xd2, 0x18, 0x42, 0x4d, 0x40, 0x5e, 0x44, 0x23,
	0xb3, 0x0a, 0x89, 0xd1, 0xbe, 0x81, 0x82, 0xa9, 0x22, 0x35, 0x23, 0xc6, 0x07, 0x63, 0xc8, 0x2c,
	0x1d, 0xc6, 0x09, 0xb3, 0x3b, 0xeb, 0xee, 0x40, 0x83, 0xc6, 0x17, 0xe3, 0x3f, 0xf0, 0xa7, 0xf9,
	0x03, 0x9a, 0x34, 0x3e, 0xfa, 0x23, 0xcc, 0xec, 0x07, 0xd5, 0x17, 0xdb, 0xa7, 0x39, 0x73, 0xcf,
	0x39, 0x33, 0xf7, 0xde, 0x1c, 0xb8, 0x2b, 0xfc, 0x30, 0x52, 0x1e, 0xf5, 0x24, 0x73, 0x43, 0x49,
	0xf5, 0x42, 0x45, 0xbe, 0x2b, 0x15, 0xe7, 0x22, 0xe0, 0xf9, 0x89, 0xc3, 0x48, 0x69, 0x85, 0x0e,
	0x2e, 0xa4, 0x38, 0x97, 0xe2, 0x4c, 0x52, 0xbf, 0xc5, 0x95, 0xe2, 0x92, 0xb9, 0x34, 0x14, 0x2e,
	0x0d, 0x02, 0xa5, 0xa9, 0x16, 0x2a, 0x88, 0x53, 0x6b, 0xfd, 0x20, 0x63, 0x93, 0x9b, 0xb7, 0x5a,
	0xb8, 0xcc, 0x0f, 0xf5, 0x26, 0x23, 0x1f, 0x71, 0xa1, 0x3f, 0xae, 0x3c, 0x3c, 0x57, 0xbe, 0xeb,
	0x9f, 0x0a, 0xbd, 0x54, 0xa7, 0x2e, 0x57, 0x9d, 0x84, 0xec, 0xac, 0xa9, 0x14, 0x27, 0x54, 0xab,
	0x28, 0x76, 0xb7, 0x30, 0xf5, 0xb5, 0x7e, 0x17, 0x00, 0xc6, 0x8a, 0xbf, 0x62, 0x71, 0x4c, 0x39,
	0x43, 0x23, 0x28, 0x49, 0xb6, 0x66, 0xb2, 0x66, 0x35, 0xad, 0x76, 0xb5, 0xfb, 0x00, 0xff, 0xa7,
	0x5d, 0x7c, 0xe1, 0x33, 0x70, 0x6c, 0x7c, 0x24, 0xb5, 0xa3, 0xdb, 0xb0, 0xeb, 0xa7, 0x54, 0xad,
	0xd0, 0xb4, 0xda, 0xe5, 0xc1, 0xee, 0xf9, 0x59, 0xc3, 0x06, 0x0b, 0x93, 0xbc, 0x8e, 0x5e, 0xc2,
	0xce, 0x42, 0x30, 0x79, 0x12, 0xd7, 0xec, 0xa6, 0xdd, 0xae, 0x74, 0x7b, 0x57, 0xfd, 0x6b, 0x94,
	0xb8, 0x86, 0x81, 0x8e, 0x36, 0x24, 0x7b, 0x02, 0x35, 0xa0, 0x12, 0x6b, 0x3a, 0x5f, 0xce, 0x74,
	0x44, 0xe7, 0xac, 0x56, 0x34, 0x7f, 0x12, 0x48, 0x4a, 0x53, 0x53, 0xa9, 0x3f, 0x81, 0xca, 0x5f,
	0x3e, 0xe4, 0x80, 0xbd, 0x64, 0x9b, 0x64, 0xca, 0x32, 0x31, 0x10, 0xed, 0x43, 0x69, 0x4d, 0xe5,
	0x2a, 0xeb, 0x97, 0xa4, 0x97, 0xc3, 0xc2, 0x63, 0xab, 0xf5, 0x02, 0xf6, 0xf2, 0xf1, 0x50, 0x19,
	0x4a, 0xaf, 0xfb, 0x93, 0xa3, 0xa7, 0xce, 0x35, 0x03, 0x47, 0xfd, 0x69, 0x7f, 0xec, 0x58, 0x06,
	0x0e, 0x09, 0x39, 0x26, 0x4e, 0x01, 0xed, 0x41, 0xf1, 0x5d, 0x9f, 0x4c, 0x1c, 0xdb, 0xa0, 0xa3,
	0xc9, 0xe8, 0xd8, 0x29, 0x1a, 0xfa, 0xd9, 0x70, 0xf0, 0xf6, 0xb9, 0x53, 0x6a, 0x7d, 0xb7, 0x92,
	0x75, 0x13, 0xf6, 0x69, 0xc5, 0x62, 0x8d, 0x7a, 0x50, 0x99, 0x4b, 0xc1, 0x02, 0x3d, 0xd3, 0x9b,
	0x90, 0xa5, 0xed, 0x0c, 0xd0, 0xf9, 0x59, 0xa3, 0x0a, 0xd7, 0xdf, 0xd3, 0xce, 0xe7, 0xd9, 0x87,
	0x2f, 0x0f, 0xef, 0xf7, 0xba, 0x5f, 0x09, 0xa4, 0xb2, 0xe9, 0x26, 0x64, 0xa8, 0xff, 0xef, 0x6e,
	0x2b, 0xdd, 0x3b, 0x57, 0xdc, 0xdc, 0x76, 0xf7, 0x5d, 0x0d, 0xd5, 0x71, 0x4a, 0xbf, 0x61, 0xd1,
	0x5a, 0xcc, 0x19, 0xf2, 0xc0, 0x1e, 0x2b, 0x8e, 0x2e, 0x7d, 0x2a, 0xeb, 0xbc, 0x7e, 0x13, 0xa7,
	0x69, 0xc4, 0x79, 0x1a, 0xf1, 0xd0, 0xa4, 0xb1, 0x55, 0xff, 0xf6, 0xf3, 0xd7, 0x8f, 0xc2, 0x7e,
	0xeb, 0xc6, 0x36, 0xff, 0xeb, 0x9e, 0x81, 0x87, 0xd6, 0x3d, 0x6f, 0x27, 0xd1, 0xf6, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x9b, 0xd8, 0x3c, 0x54, 0x2f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoggingServiceClient is the client API for LoggingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoggingServiceClient interface {
	Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type loggingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoggingServiceClient(cc grpc.ClientConnInterface) LoggingServiceClient {
	return &loggingServiceClient{cc}
}

func (c *loggingServiceClient) Log(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.platform.logging.LoggingService/Log", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoggingServiceServer is the server API for LoggingService service.
type LoggingServiceServer interface {
	Log(context.Context, *LogRequest) (*empty.Empty, error)
}

// UnimplementedLoggingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoggingServiceServer struct {
}

func (*UnimplementedLoggingServiceServer) Log(ctx context.Context, req *LogRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Log not implemented")
}

func RegisterLoggingServiceServer(s *grpc.Server, srv LoggingServiceServer) {
	s.RegisterService(&_LoggingService_serviceDesc, srv)
}

func _LoggingService_Log_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoggingServiceServer).Log(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.logging.LoggingService/Log",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoggingServiceServer).Log(ctx, req.(*LogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoggingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.platform.logging.LoggingService",
	HandlerType: (*LoggingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Log",
			Handler:    _LoggingService_Log_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/platform/logging/logging.proto",
}