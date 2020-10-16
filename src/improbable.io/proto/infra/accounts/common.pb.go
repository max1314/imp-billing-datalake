// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/accounts/common.proto

package improbable_accounts

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

type AccountId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountId) Reset()         { *m = AccountId{} }
func (m *AccountId) String() string { return proto.CompactTextString(m) }
func (*AccountId) ProtoMessage()    {}
func (*AccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e395b37e64aa0d, []int{0}
}

func (m *AccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountId.Unmarshal(m, b)
}
func (m *AccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountId.Marshal(b, m, deterministic)
}
func (m *AccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountId.Merge(m, src)
}
func (m *AccountId) XXX_Size() int {
	return xxx_messageInfo_AccountId.Size(m)
}
func (m *AccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountId.DiscardUnknown(m)
}

var xxx_messageInfo_AccountId proto.InternalMessageInfo

func (m *AccountId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type AccountEmail struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountEmail) Reset()         { *m = AccountEmail{} }
func (m *AccountEmail) String() string { return proto.CompactTextString(m) }
func (*AccountEmail) ProtoMessage()    {}
func (*AccountEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e395b37e64aa0d, []int{1}
}

func (m *AccountEmail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountEmail.Unmarshal(m, b)
}
func (m *AccountEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountEmail.Marshal(b, m, deterministic)
}
func (m *AccountEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountEmail.Merge(m, src)
}
func (m *AccountEmail) XXX_Size() int {
	return xxx_messageInfo_AccountEmail.Size(m)
}
func (m *AccountEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountEmail.DiscardUnknown(m)
}

var xxx_messageInfo_AccountEmail proto.InternalMessageInfo

func (m *AccountEmail) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ServiceAccountId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAccountId) Reset()         { *m = ServiceAccountId{} }
func (m *ServiceAccountId) String() string { return proto.CompactTextString(m) }
func (*ServiceAccountId) ProtoMessage()    {}
func (*ServiceAccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e395b37e64aa0d, []int{2}
}

func (m *ServiceAccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccountId.Unmarshal(m, b)
}
func (m *ServiceAccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccountId.Marshal(b, m, deterministic)
}
func (m *ServiceAccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccountId.Merge(m, src)
}
func (m *ServiceAccountId) XXX_Size() int {
	return xxx_messageInfo_ServiceAccountId.Size(m)
}
func (m *ServiceAccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccountId.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccountId proto.InternalMessageInfo

func (m *ServiceAccountId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type OrganisationId struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrganisationId) Reset()         { *m = OrganisationId{} }
func (m *OrganisationId) String() string { return proto.CompactTextString(m) }
func (*OrganisationId) ProtoMessage()    {}
func (*OrganisationId) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5e395b37e64aa0d, []int{3}
}

func (m *OrganisationId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrganisationId.Unmarshal(m, b)
}
func (m *OrganisationId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrganisationId.Marshal(b, m, deterministic)
}
func (m *OrganisationId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrganisationId.Merge(m, src)
}
func (m *OrganisationId) XXX_Size() int {
	return xxx_messageInfo_OrganisationId.Size(m)
}
func (m *OrganisationId) XXX_DiscardUnknown() {
	xxx_messageInfo_OrganisationId.DiscardUnknown(m)
}

var xxx_messageInfo_OrganisationId proto.InternalMessageInfo

func (m *OrganisationId) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountId)(nil), "improbable.accounts.AccountId")
	proto.RegisterType((*AccountEmail)(nil), "improbable.accounts.AccountEmail")
	proto.RegisterType((*ServiceAccountId)(nil), "improbable.accounts.ServiceAccountId")
	proto.RegisterType((*OrganisationId)(nil), "improbable.accounts.OrganisationId")
}

func init() {
	proto.RegisterFile("infra/accounts/common.proto", fileDescriptor_c5e395b37e64aa0d)
}

var fileDescriptor_c5e395b37e64aa0d = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0xcc, 0x4b, 0x2b,
	0x4a, 0xd4, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x29, 0xd6, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0xcc, 0x2d, 0x28, 0xca, 0x4f, 0x4a,
	0x4c, 0xca, 0x49, 0xd5, 0x83, 0xa9, 0x90, 0x32, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0xcf, 0x2d, 0xcf, 0x2c, 0xc9, 0xce, 0x2f, 0xd7, 0x4f, 0xcf, 0xd7, 0x05, 0xeb,
	0xd0, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x86,
	0x29, 0x49, 0x73, 0x71, 0x3a, 0x42, 0xcc, 0xf0, 0x4c, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x51, 0x4a, 0xe5, 0xe2, 0x81, 0x4a, 0xba,
	0xe6, 0x26, 0x66, 0xe6, 0x08, 0x85, 0x72, 0xb1, 0xa6, 0x82, 0x18, 0x60, 0x25, 0x9c, 0x4e, 0xf6,
	0x8f, 0xee, 0xcb, 0x5b, 0x73, 0x59, 0xc6, 0x45, 0xc7, 0x45, 0x5b, 0x25, 0xe7, 0x95, 0x14, 0xe5,
	0x58, 0xc5, 0x2a, 0xe8, 0x58, 0x59, 0xdb, 0xd8, 0x39, 0x28, 0xa9, 0x6b, 0x68, 0xc6, 0x44, 0xc7,
	0xc4, 0x56, 0xd7, 0xc6, 0x56, 0x1b, 0xe8, 0x98, 0x19, 0xd7, 0x3a, 0x44, 0x47, 0x5b, 0x25, 0xe6,
	0xe4, 0x95, 0xe6, 0x5a, 0xc5, 0xea, 0xe9, 0xc6, 0x56, 0x1b, 0x81, 0xc4, 0x54, 0x82, 0x20, 0xa6,
	0x29, 0x29, 0x71, 0x09, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0xe2, 0x76, 0x8a, 0x02, 0x17,
	0x9f, 0x7f, 0x51, 0x7a, 0x62, 0x5e, 0x66, 0x71, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0xa6, 0x8a, 0x24,
	0x36, 0xb0, 0x87, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x9d, 0x10, 0x86, 0x3c, 0x01,
	0x00, 0x00,
}
