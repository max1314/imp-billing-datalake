// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/int/platform/audit/event.proto

package improbable_platform_audit

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

type Event struct {
	Hostname                 string               `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Environment              string               `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	Cluster                  string               `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	BinaryName               string               `protobuf:"bytes,4,opt,name=binary_name,json=binaryName,proto3" json:"binary_name,omitempty"`
	Subject                  string               `protobuf:"bytes,5,opt,name=subject,proto3" json:"subject,omitempty"`
	RequestId                string               `protobuf:"bytes,6,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	TraceId                  string               `protobuf:"bytes,7,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpatialosProjectId       string               `protobuf:"bytes,8,opt,name=spatialos_project_id,json=spatialosProjectId,proto3" json:"spatialos_project_id,omitempty"`
	SpatialosProjectName     string               `protobuf:"bytes,9,opt,name=spatialos_project_name,json=spatialosProjectName,proto3" json:"spatialos_project_name,omitempty"`
	SpatialosDeploymentId    string               `protobuf:"bytes,10,opt,name=spatialos_deployment_id,json=spatialosDeploymentId,proto3" json:"spatialos_deployment_id,omitempty"`
	SpatialosDeploymentName  string               `protobuf:"bytes,11,opt,name=spatialos_deployment_name,json=spatialosDeploymentName,proto3" json:"spatialos_deployment_name,omitempty"`
	ResponseTimestamp        *timestamp.Timestamp `protobuf:"bytes,12,opt,name=response_timestamp,json=responseTimestamp,proto3" json:"response_timestamp,omitempty"`
	HttpHandlerDurationMs    float32              `protobuf:"fixed32,13,opt,name=http_handler_duration_ms,json=httpHandlerDurationMs,proto3" json:"http_handler_duration_ms,omitempty"`
	HttpRequestHost          string               `protobuf:"bytes,14,opt,name=http_request_host,json=httpRequestHost,proto3" json:"http_request_host,omitempty"`
	HttpRequestPath          string               `protobuf:"bytes,15,opt,name=http_request_path,json=httpRequestPath,proto3" json:"http_request_path,omitempty"`
	HttpRequestProtoMajor    int32                `protobuf:"varint,16,opt,name=http_request_proto_major,json=httpRequestProtoMajor,proto3" json:"http_request_proto_major,omitempty"`
	HttpRequestMethod        string               `protobuf:"bytes,17,opt,name=http_request_method,json=httpRequestMethod,proto3" json:"http_request_method,omitempty"`
	HttpRequestUserAgent     string               `protobuf:"bytes,18,opt,name=http_request_user_agent,json=httpRequestUserAgent,proto3" json:"http_request_user_agent,omitempty"`
	HttpRequestLengthBytes   int32                `protobuf:"varint,19,opt,name=http_request_length_bytes,json=httpRequestLengthBytes,proto3" json:"http_request_length_bytes,omitempty"`
	HttpRequestReferer       string               `protobuf:"bytes,20,opt,name=http_request_referer,json=httpRequestReferer,proto3" json:"http_request_referer,omitempty"`
	HttpResponseStatusCode   int32                `protobuf:"varint,21,opt,name=http_response_status_code,json=httpResponseStatusCode,proto3" json:"http_response_status_code,omitempty"`
	HttpResponseLengthBytes  int32                `protobuf:"varint,22,opt,name=http_response_length_bytes,json=httpResponseLengthBytes,proto3" json:"http_response_length_bytes,omitempty"`
	GrpcHandlerDurationMs    float32              `protobuf:"fixed32,23,opt,name=grpc_handler_duration_ms,json=grpcHandlerDurationMs,proto3" json:"grpc_handler_duration_ms,omitempty"`
	GrpcRequestService       string               `protobuf:"bytes,24,opt,name=grpc_request_service,json=grpcRequestService,proto3" json:"grpc_request_service,omitempty"`
	GrpcRequestMethod        string               `protobuf:"bytes,25,opt,name=grpc_request_method,json=grpcRequestMethod,proto3" json:"grpc_request_method,omitempty"`
	GrpcRequestClientName    string               `protobuf:"bytes,26,opt,name=grpc_request_client_name,json=grpcRequestClientName,proto3" json:"grpc_request_client_name,omitempty"`
	GrpcRequestClientVersion string               `protobuf:"bytes,27,opt,name=grpc_request_client_version,json=grpcRequestClientVersion,proto3" json:"grpc_request_client_version,omitempty"`
	GrpcResponseStatusCode   string               `protobuf:"bytes,28,opt,name=grpc_response_status_code,json=grpcResponseStatusCode,proto3" json:"grpc_response_status_code,omitempty"`
	PeerAddress              string               `protobuf:"bytes,29,opt,name=peer_address,json=peerAddress,proto3" json:"peer_address,omitempty"`
	PeerPort                 string               `protobuf:"bytes,30,opt,name=peer_port,json=peerPort,proto3" json:"peer_port,omitempty"`
	Server                   string               `protobuf:"bytes,31,opt,name=server,proto3" json:"server,omitempty"`
	XXX_NoUnkeyedLiteral     struct{}             `json:"-"`
	XXX_unrecognized         []byte               `json:"-"`
	XXX_sizecache            int32                `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e3ed9811ed35817, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Event) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *Event) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *Event) GetBinaryName() string {
	if m != nil {
		return m.BinaryName
	}
	return ""
}

func (m *Event) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Event) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Event) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *Event) GetSpatialosProjectId() string {
	if m != nil {
		return m.SpatialosProjectId
	}
	return ""
}

func (m *Event) GetSpatialosProjectName() string {
	if m != nil {
		return m.SpatialosProjectName
	}
	return ""
}

func (m *Event) GetSpatialosDeploymentId() string {
	if m != nil {
		return m.SpatialosDeploymentId
	}
	return ""
}

func (m *Event) GetSpatialosDeploymentName() string {
	if m != nil {
		return m.SpatialosDeploymentName
	}
	return ""
}

func (m *Event) GetResponseTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ResponseTimestamp
	}
	return nil
}

func (m *Event) GetHttpHandlerDurationMs() float32 {
	if m != nil {
		return m.HttpHandlerDurationMs
	}
	return 0
}

func (m *Event) GetHttpRequestHost() string {
	if m != nil {
		return m.HttpRequestHost
	}
	return ""
}

func (m *Event) GetHttpRequestPath() string {
	if m != nil {
		return m.HttpRequestPath
	}
	return ""
}

func (m *Event) GetHttpRequestProtoMajor() int32 {
	if m != nil {
		return m.HttpRequestProtoMajor
	}
	return 0
}

func (m *Event) GetHttpRequestMethod() string {
	if m != nil {
		return m.HttpRequestMethod
	}
	return ""
}

func (m *Event) GetHttpRequestUserAgent() string {
	if m != nil {
		return m.HttpRequestUserAgent
	}
	return ""
}

func (m *Event) GetHttpRequestLengthBytes() int32 {
	if m != nil {
		return m.HttpRequestLengthBytes
	}
	return 0
}

func (m *Event) GetHttpRequestReferer() string {
	if m != nil {
		return m.HttpRequestReferer
	}
	return ""
}

func (m *Event) GetHttpResponseStatusCode() int32 {
	if m != nil {
		return m.HttpResponseStatusCode
	}
	return 0
}

func (m *Event) GetHttpResponseLengthBytes() int32 {
	if m != nil {
		return m.HttpResponseLengthBytes
	}
	return 0
}

func (m *Event) GetGrpcHandlerDurationMs() float32 {
	if m != nil {
		return m.GrpcHandlerDurationMs
	}
	return 0
}

func (m *Event) GetGrpcRequestService() string {
	if m != nil {
		return m.GrpcRequestService
	}
	return ""
}

func (m *Event) GetGrpcRequestMethod() string {
	if m != nil {
		return m.GrpcRequestMethod
	}
	return ""
}

func (m *Event) GetGrpcRequestClientName() string {
	if m != nil {
		return m.GrpcRequestClientName
	}
	return ""
}

func (m *Event) GetGrpcRequestClientVersion() string {
	if m != nil {
		return m.GrpcRequestClientVersion
	}
	return ""
}

func (m *Event) GetGrpcResponseStatusCode() string {
	if m != nil {
		return m.GrpcResponseStatusCode
	}
	return ""
}

func (m *Event) GetPeerAddress() string {
	if m != nil {
		return m.PeerAddress
	}
	return ""
}

func (m *Event) GetPeerPort() string {
	if m != nil {
		return m.PeerPort
	}
	return ""
}

func (m *Event) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "improbable.platform.audit.Event")
}

func init() {
	proto.RegisterFile("improbable/int/platform/audit/event.proto", fileDescriptor_7e3ed9811ed35817)
}

var fileDescriptor_7e3ed9811ed35817 = []byte{
	// 705 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0xdf, 0x6f, 0xd3, 0x3e,
	0x10, 0xc0, 0xd5, 0x7d, 0xbf, 0xfb, 0xd1, 0xeb, 0x60, 0xd4, 0xeb, 0x5a, 0xb7, 0x63, 0xac, 0xf0,
	0x54, 0x78, 0x68, 0x27, 0x7e, 0x4d, 0x80, 0x78, 0x18, 0x1b, 0xd2, 0x2a, 0x31, 0x34, 0x75, 0xc0,
	0x6b, 0xe4, 0x36, 0xb7, 0x36, 0x53, 0x12, 0x07, 0xdb, 0xa9, 0xb4, 0xbf, 0x8c, 0x7f, 0x0f, 0xf9,
	0x9c, 0x64, 0x09, 0xed, 0xa3, 0xef, 0x73, 0x1f, 0xfb, 0x9c, 0x3b, 0x07, 0x5e, 0x06, 0x51, 0xa2,
	0xe4, 0x54, 0x4c, 0x43, 0x1c, 0x05, 0xb1, 0x19, 0x25, 0xa1, 0x30, 0xb7, 0x52, 0x45, 0x23, 0x91,
	0xfa, 0x81, 0x19, 0xe1, 0x12, 0x63, 0x33, 0x4c, 0x94, 0x34, 0x92, 0x75, 0x1f, 0x52, 0x87, 0x79,
	0xda, 0x90, 0xd2, 0x7a, 0xc7, 0x73, 0x29, 0xe7, 0x21, 0x8e, 0x28, 0x71, 0x9a, 0xde, 0x8e, 0x4c,
	0x10, 0xa1, 0x36, 0x22, 0x4a, 0x9c, 0xfb, 0xe2, 0x4f, 0x03, 0x36, 0xbf, 0xda, 0xbd, 0x58, 0x0f,
	0x76, 0x16, 0x52, 0x9b, 0x58, 0x44, 0xc8, 0x6b, 0xfd, 0xda, 0xa0, 0x3e, 0x29, 0xd6, 0xac, 0x0f,
	0x0d, 0x8c, 0x97, 0x81, 0x92, 0x71, 0x84, 0xb1, 0xe1, 0x1b, 0x84, 0xcb, 0x21, 0xc6, 0x61, 0x7b,
	0x16, 0xa6, 0xda, 0xa0, 0xe2, 0xff, 0x11, 0xcd, 0x97, 0xec, 0x18, 0x1a, 0xd3, 0x20, 0x16, 0xea,
	0xde, 0xa3, 0xad, 0xff, 0x27, 0x0a, 0x2e, 0xf4, 0xdd, 0x6e, 0xce, 0x61, 0x5b, 0xa7, 0xd3, 0x3b,
	0x9c, 0x19, 0xbe, 0xe9, 0xd4, 0x6c, 0xc9, 0x8e, 0x00, 0x14, 0xfe, 0x4e, 0x51, 0x1b, 0x2f, 0xf0,
	0xf9, 0x16, 0xc1, 0x7a, 0x16, 0x19, 0xfb, 0xac, 0x0b, 0x3b, 0x46, 0x89, 0x19, 0x5a, 0xb8, 0xed,
	0x4c, 0x5a, 0x8f, 0x7d, 0x76, 0x02, 0x2d, 0x9d, 0x08, 0x13, 0x88, 0x50, 0x6a, 0x2f, 0x51, 0xd2,
	0x6e, 0x67, 0xd3, 0x76, 0x28, 0x8d, 0x15, 0xec, 0xda, 0xa1, 0xb1, 0xcf, 0xde, 0x42, 0x7b, 0xd5,
	0xa0, 0x8a, 0xeb, 0xe4, 0xb4, 0xfe, 0x75, 0xa8, 0xf6, 0xf7, 0xd0, 0x79, 0xb0, 0x7c, 0x4c, 0x42,
	0x79, 0x6f, 0x3f, 0x87, 0x3d, 0x0a, 0x48, 0x3b, 0x28, 0xf0, 0x45, 0x41, 0xc7, 0x3e, 0xfb, 0x08,
	0xdd, 0xb5, 0x1e, 0x1d, 0xd8, 0x20, 0xb3, 0xb3, 0xc6, 0xa4, 0x33, 0xc7, 0xc0, 0x14, 0xea, 0x44,
	0xc6, 0x1a, 0xbd, 0xa2, 0x9d, 0x7c, 0xb7, 0x5f, 0x1b, 0x34, 0x5e, 0xf7, 0x86, 0xae, 0xe1, 0xc3,
	0xbc, 0xe1, 0xc3, 0x1f, 0x79, 0xc6, 0xa4, 0x99, 0x5b, 0x45, 0x88, 0x9d, 0x02, 0x5f, 0x18, 0x93,
	0x78, 0x0b, 0x11, 0xfb, 0x21, 0x2a, 0xcf, 0x4f, 0x95, 0x30, 0x81, 0x8c, 0xbd, 0x48, 0xf3, 0x47,
	0xfd, 0xda, 0x60, 0x63, 0x72, 0x60, 0xf9, 0xa5, 0xc3, 0x17, 0x19, 0xbd, 0xd2, 0xec, 0x15, 0x34,
	0x49, 0xcc, 0xdb, 0x63, 0x27, 0x85, 0x3f, 0xa6, 0xba, 0xf7, 0x2c, 0x98, 0xb8, 0xf8, 0xa5, 0xd4,
	0x66, 0x25, 0x37, 0x11, 0x66, 0xc1, 0xf7, 0x56, 0x72, 0xaf, 0x85, 0x59, 0x14, 0x05, 0x15, 0xb9,
	0xf6, 0x1a, 0x5e, 0x24, 0xee, 0xa4, 0xe2, 0x4f, 0xfa, 0xb5, 0xc1, 0xa6, 0x2b, 0x28, 0x57, 0x2c,
	0xbd, 0xb2, 0x90, 0x0d, 0x61, 0xbf, 0x22, 0x46, 0x68, 0x16, 0xd2, 0xe7, 0x4d, 0x3a, 0xa6, 0x59,
	0x72, 0xae, 0x08, 0xb0, 0x77, 0xd0, 0xa9, 0xe4, 0xa7, 0x1a, 0x95, 0x27, 0xe6, 0x76, 0xba, 0x99,
	0xeb, 0x77, 0xc9, 0xf9, 0xa9, 0x51, 0x9d, 0x59, 0xc6, 0x3e, 0x40, 0xb7, 0xa2, 0x85, 0x18, 0xcf,
	0xcd, 0xc2, 0x9b, 0xde, 0x1b, 0xd4, 0x7c, 0x9f, 0x0a, 0x6c, 0x97, 0xc4, 0x6f, 0x84, 0xbf, 0x58,
	0x6a, 0x47, 0xb2, 0xa2, 0x2a, 0xbc, 0x45, 0x85, 0x8a, 0xb7, 0xdc, 0x48, 0x96, 0xac, 0x89, 0x23,
	0xa5, 0xc3, 0xb2, 0x6e, 0x6b, 0x23, 0x4c, 0xaa, 0xbd, 0x99, 0xf4, 0x91, 0x1f, 0x94, 0x0f, 0x73,
	0xfc, 0x86, 0xf0, 0xb9, 0xf4, 0x91, 0x7d, 0x82, 0x5e, 0x55, 0xad, 0x14, 0xda, 0x26, 0xb7, 0x53,
	0x76, 0xcb, 0x95, 0x9e, 0x02, 0x9f, 0xab, 0x64, 0xb6, 0x76, 0x2a, 0x3a, 0x6e, 0x2a, 0x2c, 0x5f,
	0x9d, 0x8a, 0x13, 0x68, 0x91, 0x98, 0x5f, 0x51, 0xa3, 0x5a, 0x06, 0x33, 0xe4, 0xdc, 0x5d, 0xd1,
	0xb2, 0xec, 0x8a, 0x37, 0x8e, 0xd8, 0xb6, 0x55, 0x8c, 0xac, 0x6d, 0x5d, 0xd7, 0xb6, 0x92, 0x90,
	0xb5, 0x2d, 0x2f, 0x2d, 0xcf, 0x9f, 0x85, 0x41, 0xf1, 0x6c, 0x7a, 0xee, 0xc1, 0x95, 0xa4, 0x73,
	0xa2, 0xf4, 0x68, 0x3e, 0xc3, 0xe1, 0x3a, 0x71, 0x89, 0x4a, 0x07, 0x32, 0xe6, 0x87, 0xe4, 0xf2,
	0x15, 0xf7, 0x97, 0xe3, 0xb6, 0x15, 0x99, 0xbe, 0xa6, 0x15, 0x4f, 0x49, 0x6e, 0x3b, 0x79, 0xa5,
	0x15, 0xcf, 0x61, 0x37, 0x41, 0x3b, 0x5c, 0xbe, 0xaf, 0x50, 0x6b, 0x7e, 0xe4, 0x7e, 0x9e, 0x36,
	0x76, 0xe6, 0x42, 0xec, 0x10, 0xea, 0x94, 0x92, 0x48, 0x65, 0xf8, 0x33, 0xf7, 0xef, 0xb5, 0x81,
	0x6b, 0xa9, 0x0c, 0x6b, 0xc3, 0x96, 0xfd, 0x8e, 0xa8, 0xf8, 0x31, 0x91, 0x6c, 0x35, 0xdd, 0xa2,
	0xb7, 0xf1, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xc6, 0xe4, 0x00, 0x29, 0x06, 0x00,
	0x00,
}