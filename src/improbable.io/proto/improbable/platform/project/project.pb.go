// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/platform/project/project.proto

package improbable_platform_project

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	platform "improbable.io/proto/improbable/platform"
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

type Project struct {
	ProjectId            *platform.ProjectId  `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	CreationTime         *timestamp.Timestamp `protobuf:"bytes,2,opt,name=creation_time,json=creationTime,proto3" json:"creation_time,omitempty"`
	Title                string               `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	GraphicsURL          string               `protobuf:"bytes,5,opt,name=graphicsURL,proto3" json:"graphicsURL,omitempty"`
	StorageConfig        string               `protobuf:"bytes,6,opt,name=storage_config,json=storageConfig,proto3" json:"storage_config,omitempty"`
	OrganisationId       int64                `protobuf:"varint,7,opt,name=organisation_id,json=organisationId,proto3" json:"organisation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d172c1cebd324ce, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetProjectId() *platform.ProjectId {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *Project) GetCreationTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Project) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetGraphicsURL() string {
	if m != nil {
		return m.GraphicsURL
	}
	return ""
}

func (m *Project) GetStorageConfig() string {
	if m != nil {
		return m.StorageConfig
	}
	return ""
}

func (m *Project) GetOrganisationId() int64 {
	if m != nil {
		return m.OrganisationId
	}
	return 0
}

type ListRequest struct {
	IncludeDeleted       bool     `protobuf:"varint,1,opt,name=include_deleted,json=includeDeleted,proto3" json:"include_deleted,omitempty"`
	OrganisationId       int64    `protobuf:"varint,2,opt,name=organisation_id,json=organisationId,proto3" json:"organisation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d172c1cebd324ce, []int{1}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetIncludeDeleted() bool {
	if m != nil {
		return m.IncludeDeleted
	}
	return false
}

func (m *ListRequest) GetOrganisationId() int64 {
	if m != nil {
		return m.OrganisationId
	}
	return 0
}

type DeleteProjectRequest struct {
	ProjectId            *platform.ProjectId `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Hard                 bool                `protobuf:"varint,2,opt,name=hard,proto3" json:"hard,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeleteProjectRequest) Reset()         { *m = DeleteProjectRequest{} }
func (m *DeleteProjectRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectRequest) ProtoMessage()    {}
func (*DeleteProjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d172c1cebd324ce, []int{2}
}

func (m *DeleteProjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectRequest.Unmarshal(m, b)
}
func (m *DeleteProjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectRequest.Marshal(b, m, deterministic)
}
func (m *DeleteProjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectRequest.Merge(m, src)
}
func (m *DeleteProjectRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectRequest.Size(m)
}
func (m *DeleteProjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectRequest proto.InternalMessageInfo

func (m *DeleteProjectRequest) GetProjectId() *platform.ProjectId {
	if m != nil {
		return m.ProjectId
	}
	return nil
}

func (m *DeleteProjectRequest) GetHard() bool {
	if m != nil {
		return m.Hard
	}
	return false
}

type DeleteProjectResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteProjectResponse) Reset()         { *m = DeleteProjectResponse{} }
func (m *DeleteProjectResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteProjectResponse) ProtoMessage()    {}
func (*DeleteProjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d172c1cebd324ce, []int{3}
}

func (m *DeleteProjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteProjectResponse.Unmarshal(m, b)
}
func (m *DeleteProjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteProjectResponse.Marshal(b, m, deterministic)
}
func (m *DeleteProjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteProjectResponse.Merge(m, src)
}
func (m *DeleteProjectResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteProjectResponse.Size(m)
}
func (m *DeleteProjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteProjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteProjectResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Project)(nil), "improbable.platform.project.Project")
	proto.RegisterType((*ListRequest)(nil), "improbable.platform.project.ListRequest")
	proto.RegisterType((*DeleteProjectRequest)(nil), "improbable.platform.project.DeleteProjectRequest")
	proto.RegisterType((*DeleteProjectResponse)(nil), "improbable.platform.project.DeleteProjectResponse")
}

func init() {
	proto.RegisterFile("improbable/platform/project/project.proto", fileDescriptor_8d172c1cebd324ce)
}

var fileDescriptor_8d172c1cebd324ce = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x4f, 0xd4, 0x40,
	0x14, 0x4e, 0x17, 0x58, 0x60, 0x96, 0x5d, 0xcc, 0x88, 0xda, 0x14, 0x75, 0x37, 0x0d, 0xc4, 0x95,
	0x48, 0x8b, 0x05, 0x39, 0x78, 0x31, 0x01, 0x89, 0x21, 0xc1, 0xc4, 0x54, 0xbd, 0x19, 0x37, 0xb3,
	0xed, 0x6c, 0x19, 0x6d, 0x3b, 0xc3, 0x74, 0x16, 0x42, 0x08, 0x17, 0xff, 0x82, 0x67, 0xff, 0x90,
	0x57, 0xef, 0x92, 0x10, 0x7f, 0x88, 0xe9, 0xcc, 0x14, 0x0a, 0x54, 0x16, 0x13, 0x4e, 0xdb, 0xfd,
	0xde, 0xfb, 0xde, 0xf7, 0xbe, 0xf7, 0xde, 0x80, 0xa7, 0x24, 0x61, 0x9c, 0xf6, 0x51, 0x3f, 0xc6,
	0x2e, 0x8b, 0x91, 0x18, 0x50, 0x9e, 0xb8, 0x8c, 0xd3, 0x2f, 0x38, 0x10, 0xc5, 0xaf, 0xc3, 0x38,
	0x15, 0x14, 0xce, 0x9f, 0xa7, 0x3a, 0x45, 0xaa, 0xa3, 0x53, 0xac, 0x4e, 0x55, 0x9d, 0x80, 0x26,
	0x09, 0x4d, 0x15, 0xdd, 0x7a, 0x18, 0x51, 0x1a, 0xc5, 0xd8, 0x45, 0x8c, 0xb8, 0x28, 0x4d, 0xa9,
	0x40, 0x82, 0xd0, 0x34, 0xd3, 0xd1, 0x79, 0x1d, 0x95, 0xff, 0xfa, 0xc3, 0x81, 0x8b, 0x13, 0x26,
	0x0e, 0x75, 0xb0, 0x7d, 0x39, 0x28, 0x48, 0x82, 0x33, 0x81, 0x12, 0xa6, 0x13, 0xd6, 0x23, 0x22,
	0x76, 0x87, 0x7d, 0x27, 0xa0, 0x89, 0x9b, 0x1c, 0x10, 0xf1, 0x95, 0x1e, 0xb8, 0x11, 0x5d, 0x96,
	0xc1, 0xe5, 0x7d, 0x14, 0x93, 0x10, 0x09, 0xca, 0x33, 0xf7, 0xec, 0x53, 0xf1, 0xec, 0xdf, 0x35,
	0x30, 0xf9, 0x4e, 0x39, 0x80, 0x5b, 0x00, 0x68, 0x33, 0x3d, 0x12, 0x9a, 0x46, 0xc7, 0xe8, 0x36,
	0xbc, 0xc7, 0x4e, 0x95, 0x67, 0xcd, 0xd8, 0x0e, 0x37, 0xea, 0xa7, 0x27, 0xed, 0x5a, 0xc7, 0xf0,
	0xa7, 0x59, 0x01, 0xc1, 0x57, 0xa0, 0x19, 0x70, 0x2c, 0xbd, 0xf5, 0xf2, 0x36, 0xcd, 0x9a, 0xac,
	0x64, 0x39, 0xca, 0x83, 0x53, 0x78, 0x70, 0x3e, 0x14, 0x1e, 0xfc, 0x99, 0x82, 0x90, 0x43, 0x70,
	0x11, 0x4c, 0x08, 0x22, 0x62, 0x6c, 0x8e, 0x75, 0x8c, 0xee, 0xf4, 0xc6, 0xec, 0xe9, 0x49, 0xbb,
	0x01, 0xa6, 0x3f, 0x3b, 0x47, 0x2b, 0xcf, 0xd6, 0xd7, 0x8e, 0x17, 0x7c, 0x15, 0x85, 0x1e, 0x68,
	0x84, 0x38, 0x0b, 0x38, 0x61, 0x39, 0xd3, 0x1c, 0x97, 0xc9, 0x77, 0x4e, 0x4f, 0xda, 0x33, 0x00,
	0xc8, 0x64, 0xef, 0xc5, 0xfa, 0xf1, 0x82, 0x5f, 0x4e, 0x82, 0x1d, 0xd0, 0x88, 0x38, 0x62, 0xbb,
	0x24, 0xc8, 0x3e, 0xfa, 0x3b, 0xe6, 0x44, 0xce, 0xf1, 0xcb, 0x10, 0x5c, 0x04, 0xad, 0x4c, 0x50,
	0x8e, 0x22, 0xdc, 0x0b, 0x68, 0x3a, 0x20, 0x91, 0x59, 0x97, 0x49, 0x4d, 0x8d, 0x6e, 0x4a, 0x10,
	0x3e, 0x01, 0xb3, 0x94, 0x47, 0x28, 0x25, 0x99, 0x32, 0x4a, 0x42, 0x73, 0xb2, 0x63, 0x74, 0xc7,
	0xfc, 0x56, 0x19, 0xde, 0x0e, 0xed, 0x1e, 0x68, 0xec, 0x90, 0x4c, 0xf8, 0x78, 0x6f, 0x88, 0x33,
	0x91, 0xf3, 0x48, 0x1a, 0xc4, 0xc3, 0x10, 0xf7, 0x42, 0x1c, 0x63, 0x81, 0xd5, 0xa0, 0xa7, 0xfc,
	0x96, 0x86, 0x5f, 0x2b, 0xb4, 0x4a, 0xa0, 0x56, 0x29, 0xb0, 0x07, 0xe6, 0x14, 0x47, 0x2f, 0xa5,
	0x50, 0xba, 0xa5, 0x6d, 0x42, 0x30, 0xbe, 0x8b, 0xb8, 0x12, 0x9f, 0xf2, 0xe5, 0xb7, 0xfd, 0x00,
	0xdc, 0xbb, 0x24, 0x99, 0x31, 0x9a, 0x66, 0xd8, 0xfb, 0x39, 0x01, 0x5a, 0x1a, 0x7b, 0x8f, 0xf9,
	0x3e, 0x09, 0x30, 0x64, 0xa0, 0xbe, 0x99, 0x2f, 0x17, 0xc3, 0x05, 0xe7, 0x9a, 0xe7, 0x53, 0x34,
	0x61, 0xdd, 0xbf, 0x72, 0x26, 0x5b, 0xf9, 0x3b, 0xb0, 0xbb, 0xdf, 0x7e, 0xfd, 0xf9, 0x5e, 0xb3,
	0xad, 0x47, 0x67, 0x6f, 0x73, 0x7f, 0xd5, 0x3d, 0x3a, 0x77, 0xe8, 0xa4, 0x28, 0xc1, 0xc7, 0x2f,
	0x8d, 0xa5, 0x5c, 0xf1, 0x2d, 0x0d, 0xc9, 0xe0, 0xf0, 0x76, 0x14, 0xbd, 0xd1, 0x8a, 0x11, 0x18,
	0x7b, 0x83, 0x05, 0x1c, 0x31, 0x5d, 0xeb, 0x46, 0xed, 0xd8, 0x96, 0x94, 0x9d, 0x83, 0xf0, 0x82,
	0xac, 0xd4, 0x82, 0x9f, 0x40, 0x5d, 0x0d, 0x7e, 0xa4, 0xd6, 0xbf, 0x4c, 0xe9, 0xea, 0x4b, 0x55,
	0xd5, 0x13, 0x30, 0x9e, 0x9f, 0x2a, 0xec, 0x5e, 0xdb, 0x67, 0xe9, 0x9a, 0x6f, 0xe8, 0xe8, 0xae,
	0xd4, 0x6c, 0xc2, 0x46, 0x49, 0x73, 0xc5, 0x80, 0x3f, 0x0c, 0xd0, 0xbc, 0x70, 0x46, 0xf0, 0xf9,
	0xb5, 0xe5, 0xaa, 0xae, 0xdc, 0xf2, 0xfe, 0x87, 0xa2, 0xae, 0xd4, 0x5e, 0x94, 0xfd, 0xb4, 0x97,
	0x4a, 0x8b, 0x5d, 0xbb, 0xba, 0xd8, 0x7e, 0x5d, 0x8e, 0x6e, 0xf5, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x55, 0x1c, 0x80, 0xe1, 0x21, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProjectServiceClient interface {
	Create(ctx context.Context, in *Project, opts ...grpc.CallOption) (*empty.Empty, error)
	Modify(ctx context.Context, in *Project, opts ...grpc.CallOption) (*empty.Empty, error)
	Get(ctx context.Context, in *platform.ProjectId, opts ...grpc.CallOption) (*Project, error)
	Delete(ctx context.Context, in *platform.ProjectId, opts ...grpc.CallOption) (*empty.Empty, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (ProjectService_ListClient, error)
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Create(ctx context.Context, in *Project, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Modify(ctx context.Context, in *Project, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectService/Modify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Get(ctx context.Context, in *platform.ProjectId, opts ...grpc.CallOption) (*Project, error) {
	out := new(Project)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Delete(ctx context.Context, in *platform.ProjectId, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (ProjectService_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProjectService_serviceDesc.Streams[0], "/improbable.platform.project.ProjectService/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_ListClient interface {
	Recv() (*Project, error)
	grpc.ClientStream
}

type projectServiceListClient struct {
	grpc.ClientStream
}

func (x *projectServiceListClient) Recv() (*Project, error) {
	m := new(Project)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, "/improbable.platform.project.ProjectService/DeleteProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
type ProjectServiceServer interface {
	Create(context.Context, *Project) (*empty.Empty, error)
	Modify(context.Context, *Project) (*empty.Empty, error)
	Get(context.Context, *platform.ProjectId) (*Project, error)
	Delete(context.Context, *platform.ProjectId) (*empty.Empty, error)
	List(*ListRequest, ProjectService_ListServer) error
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
}

// UnimplementedProjectServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (*UnimplementedProjectServiceServer) Create(ctx context.Context, req *Project) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedProjectServiceServer) Modify(ctx context.Context, req *Project) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modify not implemented")
}
func (*UnimplementedProjectServiceServer) Get(ctx context.Context, req *platform.ProjectId) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedProjectServiceServer) Delete(ctx context.Context, req *platform.ProjectId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedProjectServiceServer) List(req *ListRequest, srv ProjectService_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProjectServiceServer) DeleteProject(ctx context.Context, req *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}

func RegisterProjectServiceServer(s *grpc.Server, srv ProjectServiceServer) {
	s.RegisterService(&_ProjectService_serviceDesc, srv)
}

func _ProjectService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Create(ctx, req.(*Project))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Project)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectService/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Modify(ctx, req.(*Project))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platform.ProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Get(ctx, req.(*platform.ProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(platform.ProjectId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Delete(ctx, req.(*platform.ProjectId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).List(m, &projectServiceListServer{stream})
}

type ProjectService_ListServer interface {
	Send(*Project) error
	grpc.ServerStream
}

type projectServiceListServer struct {
	grpc.ServerStream
}

func (x *projectServiceListServer) Send(m *Project) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.platform.project.ProjectService/DeleteProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.platform.project.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProjectService_Create_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _ProjectService_Modify_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProjectService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProjectService_Delete_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _ProjectService_DeleteProject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _ProjectService_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "improbable/platform/project/project.proto",
}