// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/ext/version/version.proto

package improbable_ext_version

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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

type Version struct {
	Major                int32    `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor                int32    `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch                int32    `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
	Suffix               string   `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Metadata             string   `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbf8383da455f3a7, []int{0}
}

func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetMajor() int32 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *Version) GetMinor() int32 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *Version) GetPatch() int32 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *Version) GetSuffix() string {
	if m != nil {
		return m.Suffix
	}
	return ""
}

func (m *Version) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type VersionRange struct {
	MinVersion           *Version `protobuf:"bytes,1,opt,name=min_version,json=minVersion,proto3" json:"min_version,omitempty"`
	MaxVersion           *Version `protobuf:"bytes,2,opt,name=max_version,json=maxVersion,proto3" json:"max_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRange) Reset()         { *m = VersionRange{} }
func (m *VersionRange) String() string { return proto.CompactTextString(m) }
func (*VersionRange) ProtoMessage()    {}
func (*VersionRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbf8383da455f3a7, []int{1}
}

func (m *VersionRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRange.Unmarshal(m, b)
}
func (m *VersionRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRange.Marshal(b, m, deterministic)
}
func (m *VersionRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRange.Merge(m, src)
}
func (m *VersionRange) XXX_Size() int {
	return xxx_messageInfo_VersionRange.Size(m)
}
func (m *VersionRange) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRange.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRange proto.InternalMessageInfo

func (m *VersionRange) GetMinVersion() *Version {
	if m != nil {
		return m.MinVersion
	}
	return nil
}

func (m *VersionRange) GetMaxVersion() *Version {
	if m != nil {
		return m.MaxVersion
	}
	return nil
}

type VersionSpecification struct {
	Versions             []*Version      `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	VersionRanges        []*VersionRange `protobuf:"bytes,2,rep,name=version_ranges,json=versionRanges,proto3" json:"version_ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VersionSpecification) Reset()         { *m = VersionSpecification{} }
func (m *VersionSpecification) String() string { return proto.CompactTextString(m) }
func (*VersionSpecification) ProtoMessage()    {}
func (*VersionSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_cbf8383da455f3a7, []int{2}
}

func (m *VersionSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionSpecification.Unmarshal(m, b)
}
func (m *VersionSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionSpecification.Marshal(b, m, deterministic)
}
func (m *VersionSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionSpecification.Merge(m, src)
}
func (m *VersionSpecification) XXX_Size() int {
	return xxx_messageInfo_VersionSpecification.Size(m)
}
func (m *VersionSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_VersionSpecification proto.InternalMessageInfo

func (m *VersionSpecification) GetVersions() []*Version {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *VersionSpecification) GetVersionRanges() []*VersionRange {
	if m != nil {
		return m.VersionRanges
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "improbable.ext.version.Version")
	proto.RegisterType((*VersionRange)(nil), "improbable.ext.version.VersionRange")
	proto.RegisterType((*VersionSpecification)(nil), "improbable.ext.version.VersionSpecification")
}

func init() {
	proto.RegisterFile("improbable/ext/version/version.proto", fileDescriptor_cbf8383da455f3a7)
}

var fileDescriptor_cbf8383da455f3a7 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe5, 0xf4, 0x6f, 0xff, 0xe2, 0x02, 0x43, 0x54, 0x55, 0x56, 0x97, 0x46, 0x55, 0x87,
	0x2e, 0x75, 0x24, 0x90, 0x58, 0xd8, 0x98, 0x90, 0xd8, 0x8c, 0xc4, 0x5a, 0xdd, 0xa4, 0x6e, 0x6b,
	0xa8, 0xe3, 0xc8, 0x76, 0xd3, 0x6c, 0xbc, 0x0a, 0xe2, 0xc5, 0x90, 0x78, 0x12, 0x64, 0xc7, 0x0d,
	0x0c, 0x48, 0xc0, 0x94, 0x7b, 0x8e, 0x7d, 0xbe, 0xdc, 0x23, 0xe3, 0x99, 0x90, 0xa5, 0x56, 0x19,
	0x64, 0x3b, 0x9e, 0xf2, 0xda, 0xa6, 0x15, 0xd7, 0x46, 0xa8, 0xe2, 0xf8, 0xa5, 0xa5, 0x56, 0x56,
	0xc5, 0xa3, 0xcf, 0x5b, 0x94, 0xd7, 0x96, 0x86, 0xd3, 0xf1, 0xd5, 0x46, 0xd8, 0xed, 0x3e, 0xa3,
	0xb9, 0x92, 0xa9, 0x3c, 0x08, 0xfb, 0xa4, 0x0e, 0xe9, 0x46, 0x2d, 0x7c, 0x68, 0x51, 0xc1, 0x4e,
	0xac, 0xc0, 0x2a, 0x6d, 0xd2, 0x76, 0x6c, 0x78, 0xd3, 0x67, 0xfc, 0xff, 0xa1, 0x41, 0xc4, 0x43,
	0xdc, 0x95, 0xf0, 0xa8, 0x34, 0x41, 0x09, 0x9a, 0x77, 0x59, 0x23, 0xbc, 0x2b, 0x0a, 0xa5, 0x49,
	0x14, 0x5c, 0x27, 0x9c, 0x5b, 0x82, 0xcd, 0xb7, 0xa4, 0xd3, 0xb8, 0x5e, 0xc4, 0x23, 0xdc, 0x33,
	0xfb, 0xf5, 0x5a, 0xd4, 0xe4, 0x5f, 0x82, 0xe6, 0x27, 0x2c, 0xa8, 0x78, 0x8c, 0xfb, 0x92, 0x5b,
	0x58, 0x81, 0x05, 0xd2, 0xf5, 0x27, 0xad, 0x9e, 0xbe, 0x22, 0x7c, 0x1a, 0x36, 0x60, 0x50, 0x6c,
	0x78, 0x7c, 0x8b, 0x07, 0x52, 0x14, 0xcb, 0x50, 0xcc, 0x2f, 0x33, 0xb8, 0x98, 0xd0, 0xef, 0x7b,
	0xd3, 0x10, 0xbd, 0xe9, 0xbd, 0xbf, 0x4d, 0xa2, 0x04, 0x31, 0x2c, 0x45, 0x71, 0x2c, 0xe4, 0x48,
	0x50, 0xb7, 0xa4, 0xe8, 0xaf, 0x24, 0xa8, 0x83, 0x37, 0x7d, 0x41, 0x78, 0x18, 0xe6, 0xfb, 0x92,
	0xe7, 0x62, 0x2d, 0x72, 0xb0, 0xee, 0x17, 0xd7, 0xb8, 0x1f, 0xf2, 0x86, 0xa0, 0xa4, 0xf3, 0x0b,
	0x3e, 0x6b, 0x03, 0xf1, 0x1d, 0x3e, 0x0f, 0xf3, 0x52, 0xbb, 0xea, 0x86, 0x44, 0x1e, 0x31, 0xfb,
	0x09, 0xe1, 0x2e, 0xb3, 0xb3, 0xea, 0x8b, 0x32, 0x59, 0xcf, 0xbf, 0xe7, 0xe5, 0x47, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xb5, 0x91, 0xa5, 0xbc, 0x47, 0x02, 0x00, 0x00,
}