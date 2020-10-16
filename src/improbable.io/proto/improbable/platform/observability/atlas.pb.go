// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/platform/observability/atlas.proto

package improbable_platform_observability

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type LogEntry struct {
	Id                   string                     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp            *timestamp.Timestamp       `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message              string                     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Tenant               string                     `protobuf:"bytes,4,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Resource             string                     `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource,omitempty"`
	Group                string                     `protobuf:"bytes,6,opt,name=group,proto3" json:"group,omitempty"`
	Fields               map[string]*LogEntry_Value `protobuf:"bytes,7,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_06094893b65a9eed, []int{0}
}

func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (m *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(m, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogEntry) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *LogEntry) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogEntry) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *LogEntry) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *LogEntry) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *LogEntry) GetFields() map[string]*LogEntry_Value {
	if m != nil {
		return m.Fields
	}
	return nil
}

type LogEntry_Value struct {
	// Types that are valid to be assigned to Value:
	//	*LogEntry_Value_StringValue
	//	*LogEntry_Value_Int64Value
	//	*LogEntry_Value_Float64Value
	//	*LogEntry_Value_BoolValue
	Value                isLogEntry_Value_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LogEntry_Value) Reset()         { *m = LogEntry_Value{} }
func (m *LogEntry_Value) String() string { return proto.CompactTextString(m) }
func (*LogEntry_Value) ProtoMessage()    {}
func (*LogEntry_Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_06094893b65a9eed, []int{0, 0}
}

func (m *LogEntry_Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry_Value.Unmarshal(m, b)
}
func (m *LogEntry_Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry_Value.Marshal(b, m, deterministic)
}
func (m *LogEntry_Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry_Value.Merge(m, src)
}
func (m *LogEntry_Value) XXX_Size() int {
	return xxx_messageInfo_LogEntry_Value.Size(m)
}
func (m *LogEntry_Value) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry_Value.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry_Value proto.InternalMessageInfo

type isLogEntry_Value_Value interface {
	isLogEntry_Value_Value()
}

type LogEntry_Value_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type LogEntry_Value_Int64Value struct {
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type LogEntry_Value_Float64Value struct {
	Float64Value float64 `protobuf:"fixed64,3,opt,name=float64_value,json=float64Value,proto3,oneof"`
}

type LogEntry_Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,4,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

func (*LogEntry_Value_StringValue) isLogEntry_Value_Value() {}

func (*LogEntry_Value_Int64Value) isLogEntry_Value_Value() {}

func (*LogEntry_Value_Float64Value) isLogEntry_Value_Value() {}

func (*LogEntry_Value_BoolValue) isLogEntry_Value_Value() {}

func (m *LogEntry_Value) GetValue() isLogEntry_Value_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *LogEntry_Value) GetStringValue() string {
	if x, ok := m.GetValue().(*LogEntry_Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *LogEntry_Value) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*LogEntry_Value_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *LogEntry_Value) GetFloat64Value() float64 {
	if x, ok := m.GetValue().(*LogEntry_Value_Float64Value); ok {
		return x.Float64Value
	}
	return 0
}

func (m *LogEntry_Value) GetBoolValue() bool {
	if x, ok := m.GetValue().(*LogEntry_Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LogEntry_Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LogEntry_Value_StringValue)(nil),
		(*LogEntry_Value_Int64Value)(nil),
		(*LogEntry_Value_Float64Value)(nil),
		(*LogEntry_Value_BoolValue)(nil),
	}
}

type Index struct {
	LastCompactedMinute  *timestamp.Timestamp `protobuf:"bytes,1,opt,name=last_compacted_minute,json=lastCompactedMinute,proto3" json:"last_compacted_minute,omitempty"`
	Records              []*Index_Record      `protobuf:"bytes,2,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_06094893b65a9eed, []int{1}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetLastCompactedMinute() *timestamp.Timestamp {
	if m != nil {
		return m.LastCompactedMinute
	}
	return nil
}

func (m *Index) GetRecords() []*Index_Record {
	if m != nil {
		return m.Records
	}
	return nil
}

type Index_Record struct {
	From                 *timestamp.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *timestamp.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Path                 string               `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Index_Record) Reset()         { *m = Index_Record{} }
func (m *Index_Record) String() string { return proto.CompactTextString(m) }
func (*Index_Record) ProtoMessage()    {}
func (*Index_Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_06094893b65a9eed, []int{1, 0}
}

func (m *Index_Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index_Record.Unmarshal(m, b)
}
func (m *Index_Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index_Record.Marshal(b, m, deterministic)
}
func (m *Index_Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index_Record.Merge(m, src)
}
func (m *Index_Record) XXX_Size() int {
	return xxx_messageInfo_Index_Record.Size(m)
}
func (m *Index_Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Index_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Index_Record proto.InternalMessageInfo

func (m *Index_Record) GetFrom() *timestamp.Timestamp {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Index_Record) GetTo() *timestamp.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Index_Record) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*LogEntry)(nil), "improbable.platform.observability.LogEntry")
	proto.RegisterMapType((map[string]*LogEntry_Value)(nil), "improbable.platform.observability.LogEntry.FieldsEntry")
	proto.RegisterType((*LogEntry_Value)(nil), "improbable.platform.observability.LogEntry.Value")
	proto.RegisterType((*Index)(nil), "improbable.platform.observability.Index")
	proto.RegisterType((*Index_Record)(nil), "improbable.platform.observability.Index.Record")
}

func init() {
	proto.RegisterFile("improbable/platform/observability/atlas.proto", fileDescriptor_06094893b65a9eed)
}

var fileDescriptor_06094893b65a9eed = []byte{
	// 475 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x5f, 0x8b, 0xd4, 0x30,
	0x14, 0xc5, 0xa7, 0x9d, 0x99, 0xce, 0xcc, 0xed, 0x2a, 0x12, 0xff, 0x50, 0xfa, 0xb2, 0xb3, 0x2b,
	0xc2, 0x20, 0x98, 0xe2, 0x2a, 0xba, 0xf8, 0xa8, 0xa8, 0xb3, 0xe0, 0x1f, 0x08, 0xe2, 0xeb, 0x90,
	0x4e, 0xd3, 0x1a, 0x4c, 0x9b, 0x92, 0xa4, 0xcb, 0xce, 0x57, 0x11, 0x1f, 0xfc, 0xa8, 0xd2, 0x24,
	0xdd, 0xd1, 0x17, 0x9d, 0x7d, 0xcb, 0xbd, 0xf7, 0x77, 0x4e, 0x6f, 0x4f, 0x02, 0x4f, 0x78, 0xdd,
	0x2a, 0x99, 0xd3, 0x5c, 0xb0, 0xac, 0x15, 0xd4, 0x94, 0x52, 0xd5, 0x99, 0xcc, 0x35, 0x53, 0x97,
	0x34, 0xe7, 0x82, 0x9b, 0x5d, 0x46, 0x8d, 0xa0, 0x1a, 0xb7, 0x4a, 0x1a, 0x89, 0x4e, 0xf6, 0x38,
	0x1e, 0x70, 0xfc, 0x17, 0x9e, 0x1e, 0x57, 0x52, 0x56, 0xbd, 0x5b, 0x2f, 0xc8, 0xbb, 0x32, 0x33,
	0xbc, 0x66, 0xda, 0xd0, 0xba, 0x75, 0x1e, 0xa7, 0x3f, 0x26, 0x30, 0xff, 0x20, 0xab, 0xb7, 0x8d,
	0x51, 0x3b, 0x74, 0x1b, 0x42, 0x5e, 0x24, 0xc1, 0x32, 0x58, 0x2d, 0x48, 0xc8, 0x0b, 0x74, 0x0e,
	0x8b, 0x6b, 0x3e, 0x09, 0x97, 0xc1, 0x2a, 0x3e, 0x4b, 0xb1, 0x73, 0xc4, 0x83, 0x23, 0xfe, 0x32,
	0x10, 0x64, 0x0f, 0xa3, 0x04, 0x66, 0x35, 0xd3, 0x9a, 0x56, 0x2c, 0x19, 0x5b, 0xbb, 0xa1, 0x44,
	0x0f, 0x20, 0x32, 0xac, 0xa1, 0x8d, 0x49, 0x26, 0x76, 0xe0, 0x2b, 0x94, 0xc2, 0x5c, 0x31, 0x2d,
	0x3b, 0xb5, 0x65, 0xc9, 0xd4, 0x4e, 0xae, 0x6b, 0x74, 0x0f, 0xa6, 0x95, 0x92, 0x5d, 0x9b, 0x44,
	0x76, 0xe0, 0x0a, 0xf4, 0x19, 0xa2, 0x92, 0x33, 0x51, 0xe8, 0x64, 0xb6, 0x1c, 0xaf, 0xe2, 0xb3,
	0x97, 0xf8, 0xbf, 0x79, 0xe0, 0xe1, 0x57, 0xf1, 0x3b, 0xab, 0xb4, 0x67, 0xe2, 0x6d, 0xd2, 0x5f,
	0x01, 0x4c, 0xbf, 0x52, 0xd1, 0x31, 0xf4, 0x10, 0x8e, 0xb4, 0x51, 0xbc, 0xa9, 0x36, 0x97, 0x7d,
	0xed, 0x22, 0x59, 0x8f, 0x48, 0xec, 0xba, 0x0e, 0x3a, 0x81, 0x98, 0x37, 0xe6, 0xc5, 0x73, 0xcf,
	0xf4, 0xf9, 0x8c, 0xd7, 0x23, 0x02, 0xb6, 0xe9, 0x90, 0x47, 0x70, 0xab, 0x14, 0x92, 0xee, 0xa1,
	0x3e, 0x8c, 0x60, 0x3d, 0x22, 0x47, 0xbe, 0xed, 0xb0, 0x63, 0x80, 0x5c, 0x4a, 0xe1, 0x99, 0x3e,
	0x97, 0xf9, 0x7a, 0x44, 0x16, 0x7d, 0xcf, 0x02, 0xaf, 0x67, 0x30, 0xb5, 0xb3, 0x54, 0x40, 0xfc,
	0xc7, 0xe6, 0xe8, 0x0e, 0x8c, 0xbf, 0xb3, 0x9d, 0xbf, 0xb1, 0xfe, 0x88, 0xde, 0x7b, 0xd2, 0x5f,
	0xd7, 0xd3, 0x9b, 0x64, 0x62, 0xbf, 0x45, 0x9c, 0xfe, 0x55, 0x78, 0x1e, 0x9c, 0xfe, 0x0c, 0x61,
	0x7a, 0xd1, 0x14, 0xec, 0x0a, 0x7d, 0x82, 0xfb, 0x82, 0x6a, 0xb3, 0xd9, 0xca, 0xba, 0xa5, 0x5b,
	0xc3, 0x8a, 0x4d, 0xcd, 0x9b, 0xce, 0xb8, 0x64, 0xfe, 0xfd, 0x2a, 0xee, 0xf6, 0xc2, 0x37, 0x83,
	0xee, 0xa3, 0x95, 0xa1, 0x0b, 0x98, 0x29, 0xb6, 0x95, 0xaa, 0xd0, 0x49, 0x68, 0x2f, 0x2f, 0x3b,
	0x60, 0x51, 0xbb, 0x0a, 0x26, 0x56, 0x47, 0x06, 0x7d, 0x7a, 0x05, 0x91, 0x6b, 0x21, 0x0c, 0x93,
	0x52, 0xc9, 0xfa, 0x80, 0x9d, 0x2c, 0x87, 0x1e, 0x43, 0x68, 0xe4, 0x01, 0xef, 0x3a, 0x34, 0x12,
	0x21, 0x98, 0xb4, 0xd4, 0x7c, 0xf3, 0xaf, 0xd9, 0x9e, 0xf3, 0xc8, 0xb2, 0xcf, 0x7e, 0x07, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0x56, 0x0a, 0x6b, 0xb7, 0x03, 0x00, 0x00,
}