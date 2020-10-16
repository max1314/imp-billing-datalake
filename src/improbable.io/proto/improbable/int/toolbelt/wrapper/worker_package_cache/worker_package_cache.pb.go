// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/int/toolbelt/wrapper/worker_package_cache/worker_package_cache.proto

package improbable_toolbelt_wrapper_worker_package_cache

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_package "improbable.io/proto/improbable/platform/package"
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

type WorkerPackageGetRequest struct {
	PackageId            *_package.PackageId `protobuf:"bytes,1,opt,name=package_id,json=packageId,proto3" json:"package_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *WorkerPackageGetRequest) Reset()         { *m = WorkerPackageGetRequest{} }
func (m *WorkerPackageGetRequest) String() string { return proto.CompactTextString(m) }
func (*WorkerPackageGetRequest) ProtoMessage()    {}
func (*WorkerPackageGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7528fa34839e026b, []int{0}
}

func (m *WorkerPackageGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerPackageGetRequest.Unmarshal(m, b)
}
func (m *WorkerPackageGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerPackageGetRequest.Marshal(b, m, deterministic)
}
func (m *WorkerPackageGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerPackageGetRequest.Merge(m, src)
}
func (m *WorkerPackageGetRequest) XXX_Size() int {
	return xxx_messageInfo_WorkerPackageGetRequest.Size(m)
}
func (m *WorkerPackageGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerPackageGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerPackageGetRequest proto.InternalMessageInfo

func (m *WorkerPackageGetRequest) GetPackageId() *_package.PackageId {
	if m != nil {
		return m.PackageId
	}
	return nil
}

type WorkerPackageGetResponse struct {
	PackagePath          string   `protobuf:"bytes,1,opt,name=package_path,json=packagePath,proto3" json:"package_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WorkerPackageGetResponse) Reset()         { *m = WorkerPackageGetResponse{} }
func (m *WorkerPackageGetResponse) String() string { return proto.CompactTextString(m) }
func (*WorkerPackageGetResponse) ProtoMessage()    {}
func (*WorkerPackageGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7528fa34839e026b, []int{1}
}

func (m *WorkerPackageGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkerPackageGetResponse.Unmarshal(m, b)
}
func (m *WorkerPackageGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkerPackageGetResponse.Marshal(b, m, deterministic)
}
func (m *WorkerPackageGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkerPackageGetResponse.Merge(m, src)
}
func (m *WorkerPackageGetResponse) XXX_Size() int {
	return xxx_messageInfo_WorkerPackageGetResponse.Size(m)
}
func (m *WorkerPackageGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkerPackageGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WorkerPackageGetResponse proto.InternalMessageInfo

func (m *WorkerPackageGetResponse) GetPackagePath() string {
	if m != nil {
		return m.PackagePath
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkerPackageGetRequest)(nil), "improbable.toolbelt.wrapper.worker_package_cache.WorkerPackageGetRequest")
	proto.RegisterType((*WorkerPackageGetResponse)(nil), "improbable.toolbelt.wrapper.worker_package_cache.WorkerPackageGetResponse")
}

func init() {
	proto.RegisterFile("improbable/int/toolbelt/wrapper/worker_package_cache/worker_package_cache.proto", fileDescriptor_7528fa34839e026b)
}

var fileDescriptor_7528fa34839e026b = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0x31, 0x4f, 0xc3, 0x40,
	0x0c, 0x85, 0x95, 0x05, 0xa9, 0x57, 0x06, 0x74, 0x0b, 0xa5, 0x13, 0x74, 0x40, 0xb0, 0xdc, 0xa1,
	0x32, 0x33, 0x21, 0x54, 0x95, 0x85, 0x2a, 0x0c, 0x8c, 0xc5, 0x49, 0x0d, 0x89, 0x9a, 0xd6, 0xc6,
	0x31, 0xf4, 0x37, 0xf1, 0x4b, 0xf8, 0x5b, 0x88, 0xe4, 0x4e, 0xb4, 0x22, 0x0c, 0x48, 0x4c, 0x77,
	0x7a, 0xb2, 0xbf, 0xe7, 0xf7, 0xcc, 0x5d, 0xb9, 0x62, 0xa1, 0x0c, 0xb2, 0x0a, 0x7d, 0xb9, 0x56,
	0xaf, 0x44, 0x55, 0x86, 0x95, 0xfa, 0x8d, 0x00, 0x33, 0x8a, 0xdf, 0x90, 0x2c, 0x51, 0xe6, 0x0c,
	0xf9, 0x12, 0x9e, 0x71, 0x9e, 0x43, 0x5e, 0x60, 0xa7, 0xe8, 0x58, 0x48, 0xc9, 0x5e, 0x7c, 0x03,
	0x5d, 0x84, 0xb9, 0x00, 0x73, 0x5d, 0x7b, 0xc3, 0xf3, 0xad, 0x13, 0xb8, 0x02, 0x7d, 0x22, 0x59,
	0xf9, 0x30, 0x12, 0xdf, 0x16, 0x3e, 0x7a, 0x34, 0x87, 0x0f, 0x0d, 0x62, 0xd6, 0xca, 0x13, 0xd4,
	0x14, 0x5f, 0x5e, 0xb1, 0x56, 0x7b, 0x63, 0x4c, 0xc4, 0x96, 0x8b, 0x41, 0x72, 0x9c, 0x9c, 0xf5,
	0xc7, 0xa7, 0x6e, 0xeb, 0x98, 0x88, 0x76, 0x11, 0x19, 0x18, 0xd3, 0x45, 0xda, 0xe3, 0xf8, 0x1d,
	0x5d, 0x99, 0xc1, 0x4f, 0x87, 0x9a, 0x69, 0x5d, 0xa3, 0x3d, 0x31, 0xfb, 0xd1, 0x82, 0x41, 0x8b,
	0xc6, 0xa4, 0x97, 0xf6, 0x83, 0x36, 0x03, 0x2d, 0xc6, 0x1f, 0x89, 0x39, 0xda, 0xd9, 0xbf, 0xfe,
	0x8a, 0x78, 0x8f, 0xf2, 0x56, 0xe6, 0x68, 0xdf, 0x13, 0x73, 0x30, 0x41, 0xdd, 0x19, 0xb0, 0x53,
	0xf7, 0xd7, 0xc6, 0xdc, 0x2f, 0x1d, 0x0c, 0x6f, 0xff, 0x03, 0xd5, 0x86, 0xcd, 0xf6, 0x9a, 0xc6,
	0x2f, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd6, 0xf9, 0x72, 0x83, 0x21, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerPackageCacheServiceClient is the client API for WorkerPackageCacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkerPackageCacheServiceClient interface {
	GetWorkerPackage(ctx context.Context, in *WorkerPackageGetRequest, opts ...grpc.CallOption) (*WorkerPackageGetResponse, error)
}

type workerPackageCacheServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerPackageCacheServiceClient(cc grpc.ClientConnInterface) WorkerPackageCacheServiceClient {
	return &workerPackageCacheServiceClient{cc}
}

func (c *workerPackageCacheServiceClient) GetWorkerPackage(ctx context.Context, in *WorkerPackageGetRequest, opts ...grpc.CallOption) (*WorkerPackageGetResponse, error) {
	out := new(WorkerPackageGetResponse)
	err := c.cc.Invoke(ctx, "/improbable.toolbelt.wrapper.worker_package_cache.WorkerPackageCacheService/GetWorkerPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerPackageCacheServiceServer is the server API for WorkerPackageCacheService service.
type WorkerPackageCacheServiceServer interface {
	GetWorkerPackage(context.Context, *WorkerPackageGetRequest) (*WorkerPackageGetResponse, error)
}

// UnimplementedWorkerPackageCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkerPackageCacheServiceServer struct {
}

func (*UnimplementedWorkerPackageCacheServiceServer) GetWorkerPackage(ctx context.Context, req *WorkerPackageGetRequest) (*WorkerPackageGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerPackage not implemented")
}

func RegisterWorkerPackageCacheServiceServer(s *grpc.Server, srv WorkerPackageCacheServiceServer) {
	s.RegisterService(&_WorkerPackageCacheService_serviceDesc, srv)
}

func _WorkerPackageCacheService_GetWorkerPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkerPackageGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerPackageCacheServiceServer).GetWorkerPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.toolbelt.wrapper.worker_package_cache.WorkerPackageCacheService/GetWorkerPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerPackageCacheServiceServer).GetWorkerPackage(ctx, req.(*WorkerPackageGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkerPackageCacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.toolbelt.wrapper.worker_package_cache.WorkerPackageCacheService",
	HandlerType: (*WorkerPackageCacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkerPackage",
			Handler:    _WorkerPackageCacheService_GetWorkerPackage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/int/toolbelt/wrapper/worker_package_cache/worker_package_cache.proto",
}