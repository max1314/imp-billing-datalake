// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/ext/postie/test/test_events/test_events.proto

package improbable_ext_postie_test_test_events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "improbable.io/proto/improbable/ext/postie/options"
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

type Header struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubjectId            string   `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	ClientName           string   `protobuf:"bytes,3,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_e649eb6a6fc2ae8d, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Header) GetSubjectId() string {
	if m != nil {
		return m.SubjectId
	}
	return ""
}

func (m *Header) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

type TestEvent struct {
	Header               *Header              `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Port                 int64                `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Host                 string               `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	GrpcStartTime        *timestamp.Timestamp `protobuf:"bytes,6,opt,name=grpc_start_time,json=grpcStartTime,proto3" json:"grpc_start_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestEvent) Reset()         { *m = TestEvent{} }
func (m *TestEvent) String() string { return proto.CompactTextString(m) }
func (*TestEvent) ProtoMessage()    {}
func (*TestEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e649eb6a6fc2ae8d, []int{1}
}

func (m *TestEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestEvent.Unmarshal(m, b)
}
func (m *TestEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestEvent.Marshal(b, m, deterministic)
}
func (m *TestEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestEvent.Merge(m, src)
}
func (m *TestEvent) XXX_Size() int {
	return xxx_messageInfo_TestEvent.Size(m)
}
func (m *TestEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TestEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TestEvent proto.InternalMessageInfo

func (m *TestEvent) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TestEvent) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *TestEvent) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *TestEvent) GetGrpcStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.GrpcStartTime
	}
	return nil
}

type TestDataTypes struct {
	AString              string               `protobuf:"bytes,1,opt,name=a_string,json=aString,proto3" json:"a_string,omitempty"`
	AInt32               int32                `protobuf:"varint,2,opt,name=a_int32,json=aInt32,proto3" json:"a_int32,omitempty"`
	AInt64               int64                `protobuf:"varint,3,opt,name=a_int64,json=aInt64,proto3" json:"a_int64,omitempty"`
	AUint32              uint32               `protobuf:"varint,4,opt,name=a_uint32,json=aUint32,proto3" json:"a_uint32,omitempty"`
	AUint64              uint64               `protobuf:"varint,5,opt,name=a_uint64,json=aUint64,proto3" json:"a_uint64,omitempty"`
	ADouble              float64              `protobuf:"fixed64,6,opt,name=a_double,json=aDouble,proto3" json:"a_double,omitempty"`
	AFloat               float32              `protobuf:"fixed32,7,opt,name=a_float,json=aFloat,proto3" json:"a_float,omitempty"`
	ABool                bool                 `protobuf:"varint,8,opt,name=a_bool,json=aBool,proto3" json:"a_bool,omitempty"`
	ATimestamp           *timestamp.Timestamp `protobuf:"bytes,9,opt,name=a_timestamp,json=aTimestamp,proto3" json:"a_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TestDataTypes) Reset()         { *m = TestDataTypes{} }
func (m *TestDataTypes) String() string { return proto.CompactTextString(m) }
func (*TestDataTypes) ProtoMessage()    {}
func (*TestDataTypes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e649eb6a6fc2ae8d, []int{2}
}

func (m *TestDataTypes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestDataTypes.Unmarshal(m, b)
}
func (m *TestDataTypes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestDataTypes.Marshal(b, m, deterministic)
}
func (m *TestDataTypes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestDataTypes.Merge(m, src)
}
func (m *TestDataTypes) XXX_Size() int {
	return xxx_messageInfo_TestDataTypes.Size(m)
}
func (m *TestDataTypes) XXX_DiscardUnknown() {
	xxx_messageInfo_TestDataTypes.DiscardUnknown(m)
}

var xxx_messageInfo_TestDataTypes proto.InternalMessageInfo

func (m *TestDataTypes) GetAString() string {
	if m != nil {
		return m.AString
	}
	return ""
}

func (m *TestDataTypes) GetAInt32() int32 {
	if m != nil {
		return m.AInt32
	}
	return 0
}

func (m *TestDataTypes) GetAInt64() int64 {
	if m != nil {
		return m.AInt64
	}
	return 0
}

func (m *TestDataTypes) GetAUint32() uint32 {
	if m != nil {
		return m.AUint32
	}
	return 0
}

func (m *TestDataTypes) GetAUint64() uint64 {
	if m != nil {
		return m.AUint64
	}
	return 0
}

func (m *TestDataTypes) GetADouble() float64 {
	if m != nil {
		return m.ADouble
	}
	return 0
}

func (m *TestDataTypes) GetAFloat() float32 {
	if m != nil {
		return m.AFloat
	}
	return 0
}

func (m *TestDataTypes) GetABool() bool {
	if m != nil {
		return m.ABool
	}
	return false
}

func (m *TestDataTypes) GetATimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ATimestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "improbable.ext.postie.test.test_events.Header")
	proto.RegisterType((*TestEvent)(nil), "improbable.ext.postie.test.test_events.TestEvent")
	proto.RegisterType((*TestDataTypes)(nil), "improbable.ext.postie.test.test_events.TestDataTypes")
}

func init() {
	proto.RegisterFile("improbable/ext/postie/test/test_events/test_events.proto", fileDescriptor_e649eb6a6fc2ae8d)
}

var fileDescriptor_e649eb6a6fc2ae8d = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x6a, 0x1b, 0x3b,
	0x14, 0x65, 0x1c, 0x7b, 0x6c, 0x5f, 0xe3, 0xe4, 0x3d, 0x41, 0xa9, 0x6a, 0x28, 0x35, 0x5e, 0x14,
	0x43, 0x40, 0x86, 0xc4, 0x98, 0x42, 0x57, 0x35, 0x49, 0x48, 0xa0, 0x4d, 0x41, 0x76, 0xd7, 0x42,
	0xe3, 0x91, 0x1d, 0x95, 0x99, 0xd1, 0x74, 0x74, 0xa7, 0xa4, 0x1f, 0xd0, 0x1f, 0xc8, 0xb2, 0xcb,
	0xfe, 0x83, 0xbf, 0xa3, 0xbf, 0x54, 0x24, 0xd9, 0x4e, 0x16, 0x85, 0x76, 0x33, 0x73, 0xcf, 0x3d,
	0xf7, 0x48, 0xf7, 0x1c, 0x04, 0x6f, 0x74, 0x5e, 0x56, 0x26, 0x91, 0x49, 0xa6, 0x26, 0xea, 0x1e,
	0x27, 0xa5, 0xb1, 0xa8, 0xd5, 0x04, 0x95, 0x45, 0xff, 0x11, 0xea, 0xab, 0x2a, 0xd0, 0x3e, 0xad,
	0x59, 0x59, 0x19, 0x34, 0xe4, 0xf5, 0xa3, 0x92, 0xa9, 0x7b, 0x64, 0x41, 0xc9, 0xdc, 0x20, 0x7b,
	0x32, 0x3d, 0x38, 0xfd, 0xf3, 0x0d, 0xa6, 0x44, 0x6d, 0x0a, 0xbb, 0xff, 0x87, 0x43, 0x07, 0xaf,
	0x36, 0xc6, 0x6c, 0x32, 0x35, 0xf1, 0x28, 0xa9, 0xd7, 0x13, 0xd4, 0xb9, 0xb2, 0x28, 0xf3, 0x32,
	0x0c, 0x8c, 0xbe, 0x47, 0x10, 0x5f, 0x2b, 0x99, 0xaa, 0x8a, 0x1c, 0x43, 0x43, 0xa7, 0x34, 0x1a,
	0x46, 0xe3, 0x2e, 0x6f, 0xe8, 0x94, 0xbc, 0x04, 0xb0, 0x75, 0xf2, 0x59, 0xad, 0x50, 0xe8, 0x94,
	0x36, 0x7c, 0xbf, 0xbb, 0xeb, 0xdc, 0xa4, 0xe4, 0x3d, 0xf4, 0x56, 0x99, 0x56, 0x05, 0x8a, 0x42,
	0xe6, 0x8a, 0x1e, 0x39, 0x7e, 0x7e, 0xfa, 0xb0, 0xa5, 0xff, 0xc3, 0x89, 0x35, 0x75, 0xb5, 0x52,
	0x4c, 0x59, 0xb6, 0xd6, 0x2a, 0x4b, 0x7f, 0x6e, 0x29, 0x81, 0xff, 0x50, 0x56, 0x1b, 0x85, 0x22,
	0xf9, 0x22, 0x56, 0x26, 0xab, 0xf3, 0x82, 0x43, 0xd0, 0xdf, 0xca, 0x5c, 0x8d, 0x7e, 0x45, 0xd0,
	0x5d, 0x2a, 0x8b, 0x97, 0xce, 0x24, 0xb9, 0x82, 0xf8, 0xce, 0x2f, 0xe5, 0xd7, 0xe9, 0x9d, 0x31,
	0xf6, 0x6f, 0xe1, 0xb0, 0x60, 0x85, 0xef, 0xd4, 0x84, 0x40, 0xb3, 0x34, 0x15, 0xd2, 0xe6, 0x30,
	0x1a, 0x1f, 0x71, 0x5f, 0x93, 0x11, 0x34, 0xef, 0x8c, 0x45, 0xda, 0xf2, 0x0b, 0x1f, 0x3f, 0x6c,
	0x29, 0x40, 0xe7, 0xfa, 0xe3, 0x62, 0x79, 0xfb, 0xee, 0xc3, 0x25, 0xf7, 0x1c, 0x99, 0xc3, 0xc9,
	0xa6, 0x2a, 0x57, 0xc2, 0xa2, 0xac, 0x50, 0xb8, 0xcc, 0x68, 0xec, 0x17, 0x19, 0xb0, 0x10, 0x28,
	0xdb, 0x07, 0xca, 0x96, 0xfb, 0x40, 0x79, 0xdf, 0x49, 0x16, 0x4e, 0xe1, 0x7a, 0xa3, 0x1f, 0x0d,
	0xe8, 0x3b, 0x47, 0x17, 0x12, 0xe5, 0xf2, 0x5b, 0xa9, 0x2c, 0x79, 0x01, 0x1d, 0x29, 0x2c, 0x56,
	0xba, 0xd8, 0xec, 0x62, 0x6e, 0xcb, 0x85, 0x87, 0xe4, 0x39, 0xb4, 0xa5, 0xd0, 0x05, 0x9e, 0x9f,
	0xf9, 0xa0, 0x5b, 0x3c, 0x96, 0x37, 0x0e, 0x1d, 0x88, 0xd9, 0xd4, 0x27, 0x7c, 0x14, 0x88, 0xd9,
	0x34, 0x1c, 0x56, 0x07, 0x89, 0xb3, 0xd7, 0xe7, 0x6d, 0xf9, 0xc9, 0xc3, 0x47, 0x6a, 0x36, 0xf5,
	0x2e, 0x9b, 0x3b, 0x6a, 0xaf, 0x4a, 0x4d, 0x9d, 0x64, 0xc1, 0x51, 0xc4, 0xdb, 0xf2, 0xc2, 0xc3,
	0x70, 0xd3, 0x3a, 0x33, 0x12, 0x69, 0x7b, 0x18, 0x8d, 0x1b, 0x3c, 0x96, 0x57, 0x0e, 0x91, 0x67,
	0x10, 0x4b, 0x91, 0x18, 0x93, 0xd1, 0xce, 0x30, 0x1a, 0x77, 0x78, 0x4b, 0xce, 0x8d, 0xc9, 0xc8,
	0x5b, 0xe8, 0x49, 0x71, 0x78, 0x4e, 0xb4, 0xfb, 0xd7, 0x7c, 0x40, 0x1e, 0xea, 0x24, 0xf6, 0xfc,
	0xf9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x0a, 0x5d, 0x65, 0x2f, 0x03, 0x00, 0x00,
}
