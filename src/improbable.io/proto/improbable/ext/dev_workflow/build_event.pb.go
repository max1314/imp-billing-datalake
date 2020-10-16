// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/ext/dev_workflow/build_event.proto

package improbable_ext_dev_workflow

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

type Build struct {
	PipelineName         string   `protobuf:"bytes,1,opt,name=pipeline_name,json=pipelineName,proto3" json:"pipeline_name,omitempty"`
	BranchName           string   `protobuf:"bytes,2,opt,name=branch_name,json=branchName,proto3" json:"branch_name,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_1590f054897dd02d, []int{0}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetPipelineName() string {
	if m != nil {
		return m.PipelineName
	}
	return ""
}

func (m *Build) GetBranchName() string {
	if m != nil {
		return m.BranchName
	}
	return ""
}

func (m *Build) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Build) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type Job struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_1590f054897dd02d, []int{1}
}

func (m *Job) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Job.Unmarshal(m, b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Job.Marshal(b, m, deterministic)
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return xxx_messageInfo_Job.Size(m)
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

type BuildEvent struct {
	Build                *Build               `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	Job                  *Job                 `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ParentName           string               `protobuf:"bytes,6,opt,name=parent_name,json=parentName,proto3" json:"parent_name,omitempty"`
	FiredAt              *timestamp.Timestamp `protobuf:"bytes,7,opt,name=fired_at,json=firedAt,proto3" json:"fired_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BuildEvent) Reset()         { *m = BuildEvent{} }
func (m *BuildEvent) String() string { return proto.CompactTextString(m) }
func (*BuildEvent) ProtoMessage()    {}
func (*BuildEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_1590f054897dd02d, []int{2}
}

func (m *BuildEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEvent.Unmarshal(m, b)
}
func (m *BuildEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEvent.Marshal(b, m, deterministic)
}
func (m *BuildEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEvent.Merge(m, src)
}
func (m *BuildEvent) XXX_Size() int {
	return xxx_messageInfo_BuildEvent.Size(m)
}
func (m *BuildEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEvent.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEvent proto.InternalMessageInfo

func (m *BuildEvent) GetBuild() *Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *BuildEvent) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *BuildEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildEvent) GetParentName() string {
	if m != nil {
		return m.ParentName
	}
	return ""
}

func (m *BuildEvent) GetFiredAt() *timestamp.Timestamp {
	if m != nil {
		return m.FiredAt
	}
	return nil
}

type BuildEventList struct {
	Events               []*BuildEvent `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BuildEventList) Reset()         { *m = BuildEventList{} }
func (m *BuildEventList) String() string { return proto.CompactTextString(m) }
func (*BuildEventList) ProtoMessage()    {}
func (*BuildEventList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1590f054897dd02d, []int{3}
}

func (m *BuildEventList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildEventList.Unmarshal(m, b)
}
func (m *BuildEventList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildEventList.Marshal(b, m, deterministic)
}
func (m *BuildEventList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildEventList.Merge(m, src)
}
func (m *BuildEventList) XXX_Size() int {
	return xxx_messageInfo_BuildEventList.Size(m)
}
func (m *BuildEventList) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildEventList.DiscardUnknown(m)
}

var xxx_messageInfo_BuildEventList proto.InternalMessageInfo

func (m *BuildEventList) GetEvents() []*BuildEvent {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Build)(nil), "improbable.ext.dev_workflow.Build")
	proto.RegisterType((*Job)(nil), "improbable.ext.dev_workflow.Job")
	proto.RegisterType((*BuildEvent)(nil), "improbable.ext.dev_workflow.BuildEvent")
	proto.RegisterType((*BuildEventList)(nil), "improbable.ext.dev_workflow.BuildEventList")
}

func init() {
	proto.RegisterFile("improbable/ext/dev_workflow/build_event.proto", fileDescriptor_1590f054897dd02d)
}

var fileDescriptor_1590f054897dd02d = []byte{
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x49, 0xd2, 0x3f, 0xef, 0xdd, 0xbc, 0xd7, 0xc5, 0xf0, 0x16, 0xa1, 0x6f, 0xd1, 0x12,
	0x17, 0x16, 0xc4, 0x09, 0x44, 0x04, 0x77, 0xa2, 0xe0, 0xa6, 0x88, 0x60, 0x70, 0x1f, 0x66, 0xcc,
	0x6d, 0x1d, 0x9d, 0x64, 0x42, 0x32, 0xfd, 0xf3, 0x89, 0xfd, 0x1c, 0x92, 0x3b, 0x2d, 0x2d, 0x08,
	0x75, 0x97, 0x1c, 0x7e, 0x67, 0x66, 0xce, 0x39, 0x70, 0xa9, 0xca, 0xba, 0x31, 0x52, 0x48, 0x8d,
	0x09, 0x6e, 0x6d, 0x52, 0xe0, 0x3a, 0xdf, 0x98, 0xe6, 0x63, 0xa1, 0xcd, 0x26, 0x91, 0x2b, 0xa5,
	0x8b, 0x1c, 0xd7, 0x58, 0x59, 0x5e, 0x37, 0xc6, 0x1a, 0xf6, 0xff, 0x80, 0x73, 0xdc, 0x5a, 0x7e,
	0x8c, 0x8f, 0x27, 0x4b, 0x63, 0x96, 0x1a, 0x13, 0x42, 0xe5, 0x6a, 0x91, 0x58, 0x55, 0x62, 0x6b,
	0x45, 0x59, 0x3b, 0x77, 0xdc, 0x42, 0xff, 0xbe, 0x3b, 0x92, 0x9d, 0xc1, 0xdf, 0x5a, 0xd5, 0xa8,
	0x55, 0x85, 0x79, 0x25, 0x4a, 0x8c, 0xbc, 0xa9, 0x37, 0xfb, 0x9d, 0xfd, 0xd9, 0x8b, 0x4f, 0xa2,
	0x44, 0x36, 0x81, 0x50, 0x36, 0xa2, 0x7a, 0x7d, 0x73, 0x88, 0x4f, 0x08, 0x38, 0x89, 0x80, 0x11,
	0xf8, 0xaa, 0x88, 0x02, 0xd2, 0x7d, 0x55, 0xb0, 0x7f, 0xd0, 0xd7, 0x42, 0xa2, 0x8e, 0x7a, 0x24,
	0xb9, 0x9f, 0xf8, 0x02, 0x82, 0xb9, 0x91, 0x3b, 0xd8, 0xfb, 0x0e, 0xfb, 0xc7, 0xf0, 0xa7, 0x07,
	0x40, 0x4f, 0x7c, 0xe8, 0x42, 0xb3, 0x1b, 0xe8, 0x53, 0x07, 0xe4, 0x0b, 0xd3, 0x98, 0x9f, 0x88,
	0xcf, 0xc9, 0x97, 0x39, 0x03, 0x4b, 0x21, 0x78, 0x37, 0x92, 0x0e, 0x0f, 0xd3, 0xe9, 0x49, 0xdf,
	0xdc, 0xc8, 0xac, 0x83, 0x19, 0x83, 0x1e, 0x25, 0x75, 0x89, 0xe8, 0xbb, 0x2b, 0xa1, 0x16, 0x0d,
	0x56, 0xd6, 0x95, 0x30, 0x70, 0x25, 0x38, 0x89, 0x4a, 0xb8, 0x86, 0x5f, 0x0b, 0xd5, 0x60, 0x91,
	0x0b, 0x1b, 0x0d, 0xe9, 0xb6, 0x31, 0x77, 0x3b, 0xf0, 0xfd, 0x0e, 0xfc, 0x65, 0xbf, 0x43, 0x36,
	0x24, 0xf6, 0xce, 0xc6, 0xcf, 0x30, 0x3a, 0xe4, 0x7c, 0x54, 0xad, 0x65, 0xb7, 0x30, 0xa0, 0xa5,
	0xdb, 0xc8, 0x9b, 0x06, 0xb3, 0x30, 0x3d, 0xff, 0x39, 0x2c, 0x99, 0xb3, 0x9d, 0x4d, 0x0e, 0xe8,
	0xbe, 0xab, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xb0, 0xc3, 0x5d, 0x53, 0x02, 0x00, 0x00,
}
