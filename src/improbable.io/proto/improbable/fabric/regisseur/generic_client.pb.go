// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/fabric/regisseur/generic_client.proto

package improbable_regisseur

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

type GenericClientNode struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size                 string            `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Config               *TestRunnerConfig `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GenericClientNode) Reset()         { *m = GenericClientNode{} }
func (m *GenericClientNode) String() string { return proto.CompactTextString(m) }
func (*GenericClientNode) ProtoMessage()    {}
func (*GenericClientNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_0157ba334d34ace7, []int{0}
}

func (m *GenericClientNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenericClientNode.Unmarshal(m, b)
}
func (m *GenericClientNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenericClientNode.Marshal(b, m, deterministic)
}
func (m *GenericClientNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericClientNode.Merge(m, src)
}
func (m *GenericClientNode) XXX_Size() int {
	return xxx_messageInfo_GenericClientNode.Size(m)
}
func (m *GenericClientNode) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericClientNode.DiscardUnknown(m)
}

var xxx_messageInfo_GenericClientNode proto.InternalMessageInfo

func (m *GenericClientNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GenericClientNode) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *GenericClientNode) GetConfig() *TestRunnerConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type TestRunnerConfig struct {
	Command              *TestRunnerCommand `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Instances            int32              `protobuf:"varint,2,opt,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *TestRunnerConfig) Reset()         { *m = TestRunnerConfig{} }
func (m *TestRunnerConfig) String() string { return proto.CompactTextString(m) }
func (*TestRunnerConfig) ProtoMessage()    {}
func (*TestRunnerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0157ba334d34ace7, []int{1}
}

func (m *TestRunnerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRunnerConfig.Unmarshal(m, b)
}
func (m *TestRunnerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRunnerConfig.Marshal(b, m, deterministic)
}
func (m *TestRunnerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRunnerConfig.Merge(m, src)
}
func (m *TestRunnerConfig) XXX_Size() int {
	return xxx_messageInfo_TestRunnerConfig.Size(m)
}
func (m *TestRunnerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRunnerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TestRunnerConfig proto.InternalMessageInfo

func (m *TestRunnerConfig) GetCommand() *TestRunnerCommand {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *TestRunnerConfig) GetInstances() int32 {
	if m != nil {
		return m.Instances
	}
	return 0
}

type TestRunnerCommand struct {
	AssemblyZip          string   `protobuf:"bytes,1,opt,name=assembly_zip,json=assemblyZip,proto3" json:"assembly_zip,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Args                 []string `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRunnerCommand) Reset()         { *m = TestRunnerCommand{} }
func (m *TestRunnerCommand) String() string { return proto.CompactTextString(m) }
func (*TestRunnerCommand) ProtoMessage()    {}
func (*TestRunnerCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_0157ba334d34ace7, []int{2}
}

func (m *TestRunnerCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRunnerCommand.Unmarshal(m, b)
}
func (m *TestRunnerCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRunnerCommand.Marshal(b, m, deterministic)
}
func (m *TestRunnerCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRunnerCommand.Merge(m, src)
}
func (m *TestRunnerCommand) XXX_Size() int {
	return xxx_messageInfo_TestRunnerCommand.Size(m)
}
func (m *TestRunnerCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRunnerCommand.DiscardUnknown(m)
}

var xxx_messageInfo_TestRunnerCommand proto.InternalMessageInfo

func (m *TestRunnerCommand) GetAssemblyZip() string {
	if m != nil {
		return m.AssemblyZip
	}
	return ""
}

func (m *TestRunnerCommand) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *TestRunnerCommand) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterType((*GenericClientNode)(nil), "improbable.regisseur.GenericClientNode")
	proto.RegisterType((*TestRunnerConfig)(nil), "improbable.regisseur.TestRunnerConfig")
	proto.RegisterType((*TestRunnerCommand)(nil), "improbable.regisseur.TestRunnerCommand")
}

func init() {
	proto.RegisterFile("improbable/fabric/regisseur/generic_client.proto", fileDescriptor_0157ba334d34ace7)
}

var fileDescriptor_0157ba334d34ace7 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x50, 0x41, 0x4b, 0x3a, 0x41,
	0x1c, 0x65, 0xd5, 0xbf, 0x7f, 0x1c, 0x3b, 0xe4, 0x20, 0xb1, 0x48, 0xe0, 0x26, 0x51, 0x1e, 0xda,
	0x5d, 0xd9, 0x20, 0xe8, 0x12, 0xa4, 0x87, 0xe8, 0xd2, 0x61, 0xe8, 0x64, 0xa4, 0xcc, 0x8e, 0xe3,
	0x3a, 0xb4, 0x33, 0xb3, 0xcc, 0x8c, 0x49, 0x4a, 0x9f, 0xa3, 0x8f, 0x27, 0xf8, 0x49, 0x62, 0xc7,
	0x56, 0xa1, 0x82, 0x6e, 0x8f, 0xf7, 0x7b, 0xef, 0xfd, 0x1e, 0x0f, 0xf4, 0x18, 0xcf, 0x94, 0x8c,
	0x71, 0x9c, 0xd2, 0x70, 0x8a, 0x63, 0xc5, 0x48, 0xa8, 0x68, 0xc2, 0xb4, 0xa6, 0x73, 0x15, 0x26,
	0x54, 0x50, 0xc5, 0xc8, 0x98, 0xa4, 0x8c, 0x0a, 0x13, 0x64, 0x4a, 0x1a, 0x09, 0x9b, 0x7b, 0x47,
	0xb0, 0x93, 0xb6, 0xae, 0x12, 0x66, 0x66, 0xf3, 0x38, 0x20, 0x92, 0x87, 0x7c, 0xc1, 0xcc, 0x8b,
	0x5c, 0x84, 0x89, 0xf4, 0xad, 0xc5, 0x7f, 0xc5, 0x29, 0x9b, 0x60, 0x23, 0x95, 0x0e, 0x77, 0x70,
	0x9b, 0xd6, 0xf9, 0x70, 0x40, 0xe3, 0x6e, 0xfb, 0x66, 0x60, 0xbf, 0x3c, 0xc8, 0x09, 0x85, 0x01,
	0xa8, 0x08, 0xcc, 0xa9, 0xeb, 0x78, 0x4e, 0xb7, 0xd6, 0x6f, 0x6d, 0xd6, 0xed, 0x23, 0xd0, 0x1c,
	0x3d, 0x61, 0x7f, 0x79, 0xeb, 0x0f, 0x7b, 0xfe, 0xf5, 0xf8, 0x79, 0x15, 0x5d, 0x44, 0xbd, 0xf7,
	0x53, 0x64, 0x75, 0x10, 0x82, 0x8a, 0x66, 0x4b, 0xea, 0x96, 0x72, 0x3d, 0xb2, 0x18, 0xde, 0x80,
	0x2a, 0x91, 0x62, 0xca, 0x12, 0xb7, 0xec, 0x39, 0xdd, 0x7a, 0x74, 0x16, 0xfc, 0x56, 0x3c, 0x78,
	0xa4, 0xda, 0xa0, 0xb9, 0x10, 0x54, 0x0d, 0xac, 0x1a, 0x7d, 0xb9, 0x3a, 0x2b, 0x70, 0xf8, 0xfd,
	0x06, 0xef, 0xc1, 0x7f, 0x22, 0x39, 0xc7, 0x62, 0x62, 0xab, 0xd5, 0xa3, 0xf3, 0xbf, 0x43, 0xad,
	0xbc, 0x5f, 0xdd, 0xac, 0xdb, 0x25, 0xcf, 0x41, 0x85, 0x1f, 0x1e, 0x83, 0x1a, 0x13, 0xda, 0x60,
	0x41, 0xa8, 0xb6, 0xbd, 0xff, 0xa1, 0x3d, 0xd1, 0x19, 0x81, 0xc6, 0x8f, 0x0c, 0x78, 0x02, 0x0e,
	0xb0, 0xd6, 0x94, 0xc7, 0xe9, 0xdb, 0x78, 0xc9, 0xb2, 0xed, 0x3a, 0xa8, 0x5e, 0x70, 0x43, 0x96,
	0xe5, 0x43, 0x64, 0xd8, 0xcc, 0x8a, 0x21, 0x72, 0x9c, 0x73, 0x58, 0x25, 0xda, 0x2d, 0x7b, 0xe5,
	0x9c, 0xcb, 0x71, 0x5c, 0xb5, 0xeb, 0x5f, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x80, 0x23, 0x69,
	0x67, 0xff, 0x01, 0x00, 0x00,
}