// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/platform/runtime/component_diff.proto

package improbable_platform_runtime

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ComponentDiff_Format int32

const (
	ComponentDiff_UNKNOWN      ComponentDiff_Format = 0
	ComponentDiff_DIRECT       ComponentDiff_Format = 1
	ComponentDiff_BINARY_V1    ComponentDiff_Format = 2
	ComponentDiff_QUEUE_V1     ComponentDiff_Format = 3
	ComponentDiff_HEURISTIC_V1 ComponentDiff_Format = 4
)

var ComponentDiff_Format_name = map[int32]string{
	0: "UNKNOWN",
	1: "DIRECT",
	2: "BINARY_V1",
	3: "QUEUE_V1",
	4: "HEURISTIC_V1",
}

var ComponentDiff_Format_value = map[string]int32{
	"UNKNOWN":      0,
	"DIRECT":       1,
	"BINARY_V1":    2,
	"QUEUE_V1":     3,
	"HEURISTIC_V1": 4,
}

func (x ComponentDiff_Format) String() string {
	return proto.EnumName(ComponentDiff_Format_name, int32(x))
}

func (ComponentDiff_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{0, 0}
}

type Encoding_Strategy int32

const (
	Encoding_UNKNOWN   Encoding_Strategy = 0
	Encoding_DIRECT    Encoding_Strategy = 1
	Encoding_BINARY    Encoding_Strategy = 2
	Encoding_QUEUE     Encoding_Strategy = 3
	Encoding_HEURISTIC Encoding_Strategy = 4
)

var Encoding_Strategy_name = map[int32]string{
	0: "UNKNOWN",
	1: "DIRECT",
	2: "BINARY",
	3: "QUEUE",
	4: "HEURISTIC",
}

var Encoding_Strategy_value = map[string]int32{
	"UNKNOWN":   0,
	"DIRECT":    1,
	"BINARY":    2,
	"QUEUE":     3,
	"HEURISTIC": 4,
}

func (x Encoding_Strategy) String() string {
	return proto.EnumName(Encoding_Strategy_name, int32(x))
}

func (Encoding_Strategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{1, 0}
}

type ComponentDiff struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComponentDiff) Reset()         { *m = ComponentDiff{} }
func (m *ComponentDiff) String() string { return proto.CompactTextString(m) }
func (*ComponentDiff) ProtoMessage()    {}
func (*ComponentDiff) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{0}
}

func (m *ComponentDiff) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentDiff.Unmarshal(m, b)
}
func (m *ComponentDiff) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentDiff.Marshal(b, m, deterministic)
}
func (m *ComponentDiff) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentDiff.Merge(m, src)
}
func (m *ComponentDiff) XXX_Size() int {
	return xxx_messageInfo_ComponentDiff.Size(m)
}
func (m *ComponentDiff) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentDiff.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentDiff proto.InternalMessageInfo

type Encoding struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Encoding) Reset()         { *m = Encoding{} }
func (m *Encoding) String() string { return proto.CompactTextString(m) }
func (*Encoding) ProtoMessage()    {}
func (*Encoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{1}
}

func (m *Encoding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Encoding.Unmarshal(m, b)
}
func (m *Encoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Encoding.Marshal(b, m, deterministic)
}
func (m *Encoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Encoding.Merge(m, src)
}
func (m *Encoding) XXX_Size() int {
	return xxx_messageInfo_Encoding.Size(m)
}
func (m *Encoding) XXX_DiscardUnknown() {
	xxx_messageInfo_Encoding.DiscardUnknown(m)
}

var xxx_messageInfo_Encoding proto.InternalMessageInfo

type EncodingStrategyToFormatEntry struct {
	Strategy             Encoding_Strategy    `protobuf:"varint,1,opt,name=strategy,proto3,enum=improbable.platform.runtime.Encoding_Strategy" json:"strategy,omitempty"`
	Format               ComponentDiff_Format `protobuf:"varint,2,opt,name=format,proto3,enum=improbable.platform.runtime.ComponentDiff_Format" json:"format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EncodingStrategyToFormatEntry) Reset()         { *m = EncodingStrategyToFormatEntry{} }
func (m *EncodingStrategyToFormatEntry) String() string { return proto.CompactTextString(m) }
func (*EncodingStrategyToFormatEntry) ProtoMessage()    {}
func (*EncodingStrategyToFormatEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{2}
}

func (m *EncodingStrategyToFormatEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodingStrategyToFormatEntry.Unmarshal(m, b)
}
func (m *EncodingStrategyToFormatEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodingStrategyToFormatEntry.Marshal(b, m, deterministic)
}
func (m *EncodingStrategyToFormatEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodingStrategyToFormatEntry.Merge(m, src)
}
func (m *EncodingStrategyToFormatEntry) XXX_Size() int {
	return xxx_messageInfo_EncodingStrategyToFormatEntry.Size(m)
}
func (m *EncodingStrategyToFormatEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodingStrategyToFormatEntry.DiscardUnknown(m)
}

var xxx_messageInfo_EncodingStrategyToFormatEntry proto.InternalMessageInfo

func (m *EncodingStrategyToFormatEntry) GetStrategy() Encoding_Strategy {
	if m != nil {
		return m.Strategy
	}
	return Encoding_UNKNOWN
}

func (m *EncodingStrategyToFormatEntry) GetFormat() ComponentDiff_Format {
	if m != nil {
		return m.Format
	}
	return ComponentDiff_UNKNOWN
}

type EncodingStrategyToFormat struct {
	StrategyToFormat     []*EncodingStrategyToFormatEntry `protobuf:"bytes,1,rep,name=strategy_to_format,json=strategyToFormat,proto3" json:"strategy_to_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *EncodingStrategyToFormat) Reset()         { *m = EncodingStrategyToFormat{} }
func (m *EncodingStrategyToFormat) String() string { return proto.CompactTextString(m) }
func (*EncodingStrategyToFormat) ProtoMessage()    {}
func (*EncodingStrategyToFormat) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{3}
}

func (m *EncodingStrategyToFormat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EncodingStrategyToFormat.Unmarshal(m, b)
}
func (m *EncodingStrategyToFormat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EncodingStrategyToFormat.Marshal(b, m, deterministic)
}
func (m *EncodingStrategyToFormat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EncodingStrategyToFormat.Merge(m, src)
}
func (m *EncodingStrategyToFormat) XXX_Size() int {
	return xxx_messageInfo_EncodingStrategyToFormat.Size(m)
}
func (m *EncodingStrategyToFormat) XXX_DiscardUnknown() {
	xxx_messageInfo_EncodingStrategyToFormat.DiscardUnknown(m)
}

var xxx_messageInfo_EncodingStrategyToFormat proto.InternalMessageInfo

func (m *EncodingStrategyToFormat) GetStrategyToFormat() []*EncodingStrategyToFormatEntry {
	if m != nil {
		return m.StrategyToFormat
	}
	return nil
}

type ComponentDiffFormats struct {
	Formats              []ComponentDiff_Format `protobuf:"varint,1,rep,packed,name=formats,proto3,enum=improbable.platform.runtime.ComponentDiff_Format" json:"formats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ComponentDiffFormats) Reset()         { *m = ComponentDiffFormats{} }
func (m *ComponentDiffFormats) String() string { return proto.CompactTextString(m) }
func (*ComponentDiffFormats) ProtoMessage()    {}
func (*ComponentDiffFormats) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{4}
}

func (m *ComponentDiffFormats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentDiffFormats.Unmarshal(m, b)
}
func (m *ComponentDiffFormats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentDiffFormats.Marshal(b, m, deterministic)
}
func (m *ComponentDiffFormats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentDiffFormats.Merge(m, src)
}
func (m *ComponentDiffFormats) XXX_Size() int {
	return xxx_messageInfo_ComponentDiffFormats.Size(m)
}
func (m *ComponentDiffFormats) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentDiffFormats.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentDiffFormats proto.InternalMessageInfo

func (m *ComponentDiffFormats) GetFormats() []ComponentDiff_Format {
	if m != nil {
		return m.Formats
	}
	return nil
}

type ComponentDiffSettings struct {
	ComponentIdToFormat  map[uint32]ComponentDiff_Format `protobuf:"bytes,1,rep,name=component_id_to_format,json=componentIdToFormat,proto3" json:"component_id_to_format,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=improbable.platform.runtime.ComponentDiff_Format"`
	DefaultFormat        ComponentDiff_Format            `protobuf:"varint,2,opt,name=default_format,json=defaultFormat,proto3,enum=improbable.platform.runtime.ComponentDiff_Format" json:"default_format,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ComponentDiffSettings) Reset()         { *m = ComponentDiffSettings{} }
func (m *ComponentDiffSettings) String() string { return proto.CompactTextString(m) }
func (*ComponentDiffSettings) ProtoMessage()    {}
func (*ComponentDiffSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{5}
}

func (m *ComponentDiffSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComponentDiffSettings.Unmarshal(m, b)
}
func (m *ComponentDiffSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComponentDiffSettings.Marshal(b, m, deterministic)
}
func (m *ComponentDiffSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComponentDiffSettings.Merge(m, src)
}
func (m *ComponentDiffSettings) XXX_Size() int {
	return xxx_messageInfo_ComponentDiffSettings.Size(m)
}
func (m *ComponentDiffSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_ComponentDiffSettings.DiscardUnknown(m)
}

var xxx_messageInfo_ComponentDiffSettings proto.InternalMessageInfo

func (m *ComponentDiffSettings) GetComponentIdToFormat() map[uint32]ComponentDiff_Format {
	if m != nil {
		return m.ComponentIdToFormat
	}
	return nil
}

func (m *ComponentDiffSettings) GetDefaultFormat() ComponentDiff_Format {
	if m != nil {
		return m.DefaultFormat
	}
	return ComponentDiff_UNKNOWN
}

type FieldDiffs struct {
	Diffs                map[uint32][]byte `protobuf:"bytes,1,rep,name=diffs,proto3" json:"diffs,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *FieldDiffs) Reset()         { *m = FieldDiffs{} }
func (m *FieldDiffs) String() string { return proto.CompactTextString(m) }
func (*FieldDiffs) ProtoMessage()    {}
func (*FieldDiffs) Descriptor() ([]byte, []int) {
	return fileDescriptor_a4cc4d9c17908a00, []int{6}
}

func (m *FieldDiffs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldDiffs.Unmarshal(m, b)
}
func (m *FieldDiffs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldDiffs.Marshal(b, m, deterministic)
}
func (m *FieldDiffs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldDiffs.Merge(m, src)
}
func (m *FieldDiffs) XXX_Size() int {
	return xxx_messageInfo_FieldDiffs.Size(m)
}
func (m *FieldDiffs) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldDiffs.DiscardUnknown(m)
}

var xxx_messageInfo_FieldDiffs proto.InternalMessageInfo

func (m *FieldDiffs) GetDiffs() map[uint32][]byte {
	if m != nil {
		return m.Diffs
	}
	return nil
}

func init() {
	proto.RegisterEnum("improbable.platform.runtime.ComponentDiff_Format", ComponentDiff_Format_name, ComponentDiff_Format_value)
	proto.RegisterEnum("improbable.platform.runtime.Encoding_Strategy", Encoding_Strategy_name, Encoding_Strategy_value)
	proto.RegisterType((*ComponentDiff)(nil), "improbable.platform.runtime.ComponentDiff")
	proto.RegisterType((*Encoding)(nil), "improbable.platform.runtime.Encoding")
	proto.RegisterType((*EncodingStrategyToFormatEntry)(nil), "improbable.platform.runtime.EncodingStrategyToFormatEntry")
	proto.RegisterType((*EncodingStrategyToFormat)(nil), "improbable.platform.runtime.EncodingStrategyToFormat")
	proto.RegisterType((*ComponentDiffFormats)(nil), "improbable.platform.runtime.ComponentDiffFormats")
	proto.RegisterType((*ComponentDiffSettings)(nil), "improbable.platform.runtime.ComponentDiffSettings")
	proto.RegisterMapType((map[uint32]ComponentDiff_Format)(nil), "improbable.platform.runtime.ComponentDiffSettings.ComponentIdToFormatEntry")
	proto.RegisterType((*FieldDiffs)(nil), "improbable.platform.runtime.FieldDiffs")
	proto.RegisterMapType((map[uint32][]byte)(nil), "improbable.platform.runtime.FieldDiffs.DiffsEntry")
}

func init() {
	proto.RegisterFile("improbable/platform/runtime/component_diff.proto", fileDescriptor_a4cc4d9c17908a00)
}

var fileDescriptor_a4cc4d9c17908a00 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcf, 0x6e, 0x9b, 0x40,
	0x10, 0xc6, 0x0b, 0x4e, 0x88, 0x33, 0xb1, 0x23, 0x34, 0x75, 0x2b, 0x2b, 0x55, 0xa5, 0x88, 0x53,
	0x4e, 0xeb, 0xc6, 0xbd, 0x44, 0xb9, 0xd5, 0x0e, 0x69, 0xa8, 0x25, 0x37, 0xc5, 0xa6, 0x7f, 0x4e,
	0x08, 0x9b, 0xc5, 0x45, 0x05, 0xd6, 0x82, 0x75, 0x25, 0xdf, 0x7a, 0xe8, 0x03, 0xf4, 0x69, 0xfa,
	0x5a, 0x7d, 0x85, 0x0a, 0x96, 0x0d, 0xa1, 0x89, 0xad, 0x28, 0xbd, 0x20, 0x76, 0x96, 0xef, 0xf7,
	0x7d, 0x33, 0xac, 0x16, 0x5e, 0x85, 0xf1, 0x32, 0x65, 0x33, 0x6f, 0x16, 0xd1, 0xde, 0x32, 0xf2,
	0x78, 0xc0, 0xd2, 0xb8, 0x97, 0xae, 0x12, 0x1e, 0xc6, 0xb4, 0x37, 0x67, 0xf1, 0x92, 0x25, 0x34,
	0xe1, 0xae, 0x1f, 0x06, 0x01, 0x59, 0xa6, 0x8c, 0x33, 0x7c, 0x51, 0x29, 0x88, 0x54, 0x90, 0x52,
	0x61, 0x78, 0xd0, 0x1e, 0x4a, 0xd1, 0x45, 0x18, 0x04, 0xc6, 0x35, 0x68, 0x97, 0x2c, 0x8d, 0x3d,
	0x8e, 0x07, 0xb0, 0xe7, 0x8c, 0x47, 0xe3, 0xf7, 0x9f, 0xc6, 0xfa, 0x13, 0x04, 0xd0, 0x2e, 0x2c,
	0xdb, 0x1c, 0x4e, 0x75, 0x05, 0xdb, 0xb0, 0x3f, 0xb0, 0xc6, 0x6f, 0xec, 0x2f, 0xee, 0xc7, 0x53,
	0x5d, 0xc5, 0x16, 0x34, 0x3f, 0x38, 0xa6, 0x63, 0xe6, 0xab, 0x06, 0xea, 0xd0, 0xba, 0x32, 0x1d,
	0xdb, 0x9a, 0x4c, 0xad, 0x61, 0x5e, 0xd9, 0x31, 0x1c, 0x68, 0x9a, 0xc9, 0x9c, 0xf9, 0x61, 0xb2,
	0x30, 0x2c, 0x68, 0x4e, 0x78, 0xea, 0x71, 0xba, 0x58, 0x6f, 0xe6, 0x03, 0x68, 0x82, 0xaf, 0xab,
	0xb8, 0x0f, 0xbb, 0x05, 0x5c, 0x6f, 0xe4, 0xb6, 0x37, 0x64, 0x7d, 0xc7, 0xf8, 0xad, 0xc0, 0x4b,
	0xc9, 0x95, 0xcc, 0x29, 0x13, 0xd9, 0xcd, 0x84, 0xa7, 0x6b, 0x7c, 0x07, 0xcd, 0xac, 0xdc, 0xe8,
	0x2a, 0xc7, 0xca, 0xc9, 0x61, 0x9f, 0x90, 0x2d, 0xb3, 0x20, 0x92, 0x46, 0x24, 0xce, 0xbe, 0xd1,
	0xa3, 0x05, 0x5a, 0x50, 0xa0, 0xbb, 0x6a, 0x41, 0x3a, 0xdd, 0x4a, 0xaa, 0x8d, 0x94, 0x88, 0x4c,
	0x76, 0x09, 0x30, 0x7e, 0x2a, 0xd0, 0xdd, 0x14, 0x1c, 0xbf, 0x02, 0x4a, 0x4f, 0x97, 0x33, 0xb7,
	0xf4, 0x54, 0x8e, 0x1b, 0x27, 0x07, 0xfd, 0xf3, 0x07, 0xa5, 0xbf, 0x77, 0x16, 0xb6, 0x9e, 0xfd,
	0x53, 0x36, 0xe6, 0xd0, 0xa9, 0xc5, 0x14, 0xe5, 0x0c, 0x47, 0xb0, 0x27, 0x5c, 0xb3, 0xc2, 0xf6,
	0x51, 0xad, 0x4a, 0x82, 0xf1, 0x47, 0x85, 0x67, 0xb5, 0x2f, 0x26, 0x94, 0xf3, 0x30, 0x59, 0x64,
	0xf8, 0x43, 0x81, 0xe7, 0xd5, 0x71, 0x0d, 0xfd, 0x3b, 0xdd, 0x8e, 0x1e, 0x6e, 0x2b, 0xa1, 0x55,
	0xd5, 0xf2, 0xeb, 0xed, 0x3f, 0x9d, 0xdf, 0xdd, 0xc1, 0xcf, 0x70, 0xe8, 0xd3, 0xc0, 0x5b, 0x45,
	0xdc, 0xfd, 0xdf, 0x7f, 0xdb, 0x2e, 0x41, 0x62, 0x79, 0xb4, 0x86, 0xee, 0xa6, 0x28, 0xa8, 0x43,
	0xe3, 0x1b, 0x15, 0x07, 0xb2, 0x6d, 0xe7, 0xaf, 0xf8, 0x16, 0x76, 0xbf, 0x7b, 0xd1, 0x8a, 0x3e,
	0xde, 0x5e, 0xe8, 0xcf, 0xd5, 0x33, 0xc5, 0xf8, 0xa5, 0x00, 0x5c, 0x86, 0x34, 0xf2, 0xf3, 0xfd,
	0x0c, 0xaf, 0x60, 0x37, 0xbf, 0x0a, 0xb2, 0x72, 0xa8, 0xfd, 0xad, 0xec, 0x4a, 0x47, 0x8a, 0xa7,
	0x98, 0x9d, 0x00, 0x1c, 0x9d, 0x01, 0x54, 0xc5, 0x7b, 0xba, 0xe8, 0xdc, 0xee, 0xa2, 0x75, 0x2b,
	0xd2, 0xa0, 0x33, 0xc0, 0x5a, 0xea, 0xeb, 0xfc, 0x5a, 0x9a, 0x69, 0xc5, 0xed, 0xf4, 0xfa, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc4, 0xd4, 0x6b, 0x2a, 0xd1, 0x04, 0x00, 0x00,
}
