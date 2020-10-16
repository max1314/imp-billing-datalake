// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/int/platform/imp_billing/record.proto

package improbable_int_platform_imp_billing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Record struct {
	OrganisationId       string   `protobuf:"bytes,1,opt,name=organisation_id,json=organisationId,proto3" json:"organisation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e446db444562c6, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetOrganisationId() string {
	if m != nil {
		return m.OrganisationId
	}
	return ""
}

type ListRecordsRequest struct {
	Currency             string   `protobuf:"bytes,1,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRecordsRequest) Reset()         { *m = ListRecordsRequest{} }
func (m *ListRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*ListRecordsRequest) ProtoMessage()    {}
func (*ListRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e446db444562c6, []int{1}
}

func (m *ListRecordsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecordsRequest.Unmarshal(m, b)
}
func (m *ListRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecordsRequest.Marshal(b, m, deterministic)
}
func (m *ListRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecordsRequest.Merge(m, src)
}
func (m *ListRecordsRequest) XXX_Size() int {
	return xxx_messageInfo_ListRecordsRequest.Size(m)
}
func (m *ListRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecordsRequest proto.InternalMessageInfo

func (m *ListRecordsRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type ListRecordsResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListRecordsResponse) Reset()         { *m = ListRecordsResponse{} }
func (m *ListRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*ListRecordsResponse) ProtoMessage()    {}
func (*ListRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_61e446db444562c6, []int{2}
}

func (m *ListRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRecordsResponse.Unmarshal(m, b)
}
func (m *ListRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRecordsResponse.Marshal(b, m, deterministic)
}
func (m *ListRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRecordsResponse.Merge(m, src)
}
func (m *ListRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_ListRecordsResponse.Size(m)
}
func (m *ListRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListRecordsResponse proto.InternalMessageInfo

func (m *ListRecordsResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "improbable.int.platform.imp_billing.Record")
	proto.RegisterType((*ListRecordsRequest)(nil), "improbable.int.platform.imp_billing.ListRecordsRequest")
	proto.RegisterType((*ListRecordsResponse)(nil), "improbable.int.platform.imp_billing.ListRecordsResponse")
}

func init() {
	proto.RegisterFile("improbable/int/platform/imp_billing/record.proto", fileDescriptor_61e446db444562c6)
}

var fileDescriptor_61e446db444562c6 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xd0, 0x41, 0x4a, 0x03, 0x31,
	0x14, 0x06, 0x60, 0x46, 0xa5, 0x6a, 0x04, 0x0b, 0x71, 0x53, 0x86, 0x22, 0x65, 0x14, 0x2c, 0x48,
	0x93, 0x4e, 0x5d, 0xe8, 0x05, 0x5c, 0x08, 0xae, 0x46, 0x70, 0x25, 0x94, 0xcc, 0x4c, 0x1c, 0x1e,
	0x64, 0xf2, 0x62, 0x92, 0x29, 0xb8, 0xf5, 0x0a, 0x1e, 0xc0, 0x43, 0x78, 0x14, 0xaf, 0xe0, 0x41,
	0xc4, 0xc9, 0x54, 0x2b, 0x6e, 0xda, 0x65, 0x1e, 0xf9, 0xfe, 0xe4, 0x7f, 0x64, 0x0a, 0xb5, 0xb1,
	0x98, 0x8b, 0x5c, 0x49, 0x0e, 0xda, 0x73, 0xa3, 0x84, 0x7f, 0x44, 0x5b, 0x73, 0xa8, 0xcd, 0x3c,
	0x07, 0xa5, 0x40, 0x57, 0xdc, 0xca, 0x02, 0x6d, 0xc9, 0x8c, 0x45, 0x8f, 0xf4, 0xe4, 0x57, 0x30,
	0xd0, 0x9e, 0x2d, 0x05, 0x5b, 0x11, 0xf1, 0xb0, 0x42, 0xac, 0x94, 0xe4, 0xc2, 0x00, 0x17, 0x5a,
	0xa3, 0x17, 0x1e, 0x50, 0xbb, 0x10, 0x91, 0xa4, 0xa4, 0x97, 0xb5, 0x91, 0xf4, 0x8c, 0xf4, 0xd1,
	0x56, 0x42, 0x83, 0x6b, 0x2f, 0xcc, 0xa1, 0x1c, 0x44, 0xa3, 0x68, 0xbc, 0x9f, 0x1d, 0xae, 0x8e,
	0x6f, 0xca, 0x64, 0x4a, 0xe8, 0x2d, 0x38, 0x1f, 0x98, 0xcb, 0xe4, 0x53, 0x23, 0x9d, 0xa7, 0x31,
	0xd9, 0x2b, 0x1a, 0x6b, 0xa5, 0x2e, 0x9e, 0x3b, 0xf7, 0x73, 0x4e, 0x1e, 0xc8, 0xd1, 0x1f, 0xe1,
	0x0c, 0x6a, 0x27, 0xe9, 0x35, 0xd9, 0x0d, 0x75, 0xdc, 0x20, 0x1a, 0x6d, 0x8f, 0x0f, 0x66, 0xe7,
	0x6c, 0x8d, 0x42, 0x2c, 0xc4, 0x64, 0x4b, 0x3b, 0x7b, 0x8f, 0x48, 0x3f, 0xcc, 0xee, 0xa4, 0x5d,
	0x40, 0x21, 0xef, 0x53, 0xfa, 0x16, 0x91, 0x9d, 0xef, 0x27, 0xe9, 0xe5, 0x5a, 0x91, 0xff, 0xfb,
	0xc4, 0x57, 0x9b, 0xc3, 0x50, 0x2b, 0x39, 0x7d, 0xf9, 0xf8, 0x7c, 0xdd, 0x3a, 0xa6, 0x43, 0xbe,
	0x48, 0x79, 0xd1, 0x38, 0x8f, 0xb5, 0xb4, 0x93, 0x4e, 0x4c, 0xba, 0x5f, 0xe7, 0xbd, 0x76, 0xff,
	0x17, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0x4a, 0x12, 0x67, 0xf6, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RecordServiceV1Client is the client API for RecordServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecordServiceV1Client interface {
	List(ctx context.Context, in *ListRecordsRequest, opts ...grpc.CallOption) (*ListRecordsResponse, error)
}

type recordServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewRecordServiceV1Client(cc grpc.ClientConnInterface) RecordServiceV1Client {
	return &recordServiceV1Client{cc}
}

func (c *recordServiceV1Client) List(ctx context.Context, in *ListRecordsRequest, opts ...grpc.CallOption) (*ListRecordsResponse, error) {
	out := new(ListRecordsResponse)
	err := c.cc.Invoke(ctx, "/improbable.int.platform.imp_billing.RecordServiceV1/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecordServiceV1Server is the server API for RecordServiceV1 service.
type RecordServiceV1Server interface {
	List(context.Context, *ListRecordsRequest) (*ListRecordsResponse, error)
}

// UnimplementedRecordServiceV1Server can be embedded to have forward compatible implementations.
type UnimplementedRecordServiceV1Server struct {
}

func (*UnimplementedRecordServiceV1Server) List(ctx context.Context, req *ListRecordsRequest) (*ListRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRecordServiceV1Server(s *grpc.Server, srv RecordServiceV1Server) {
	s.RegisterService(&_RecordServiceV1_serviceDesc, srv)
}

func _RecordServiceV1_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecordServiceV1Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.int.platform.imp_billing.RecordServiceV1/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecordServiceV1Server).List(ctx, req.(*ListRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RecordServiceV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.int.platform.imp_billing.RecordServiceV1",
	HandlerType: (*RecordServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _RecordServiceV1_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/int/platform/imp_billing/record.proto",
}
