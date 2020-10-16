// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/spatialos/playerauth/v2/devauth.proto

package pb_playerauthv2

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
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

type CreateLoginTokensRequest struct {
	PlayerIdentityToken  string             `protobuf:"bytes,1,opt,name=player_identity_token,json=playerIdentityToken,proto3" json:"player_identity_token,omitempty"`
	LifetimeDuration     *duration.Duration `protobuf:"bytes,2,opt,name=lifetime_duration,json=lifetimeDuration,proto3" json:"lifetime_duration,omitempty"`
	WorkerType           string             `protobuf:"bytes,3,opt,name=worker_type,json=workerType,proto3" json:"worker_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateLoginTokensRequest) Reset()         { *m = CreateLoginTokensRequest{} }
func (m *CreateLoginTokensRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLoginTokensRequest) ProtoMessage()    {}
func (*CreateLoginTokensRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69bfbae40e7453c, []int{0}
}

func (m *CreateLoginTokensRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLoginTokensRequest.Unmarshal(m, b)
}
func (m *CreateLoginTokensRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLoginTokensRequest.Marshal(b, m, deterministic)
}
func (m *CreateLoginTokensRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLoginTokensRequest.Merge(m, src)
}
func (m *CreateLoginTokensRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLoginTokensRequest.Size(m)
}
func (m *CreateLoginTokensRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLoginTokensRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLoginTokensRequest proto.InternalMessageInfo

func (m *CreateLoginTokensRequest) GetPlayerIdentityToken() string {
	if m != nil {
		return m.PlayerIdentityToken
	}
	return ""
}

func (m *CreateLoginTokensRequest) GetLifetimeDuration() *duration.Duration {
	if m != nil {
		return m.LifetimeDuration
	}
	return nil
}

func (m *CreateLoginTokensRequest) GetWorkerType() string {
	if m != nil {
		return m.WorkerType
	}
	return ""
}

type CreateLoginTokensResponse struct {
	LoginTokenDetails    []*LoginTokenDetails `protobuf:"bytes,1,rep,name=login_token_details,json=loginTokenDetails,proto3" json:"login_token_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateLoginTokensResponse) Reset()         { *m = CreateLoginTokensResponse{} }
func (m *CreateLoginTokensResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLoginTokensResponse) ProtoMessage()    {}
func (*CreateLoginTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69bfbae40e7453c, []int{1}
}

func (m *CreateLoginTokensResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLoginTokensResponse.Unmarshal(m, b)
}
func (m *CreateLoginTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLoginTokensResponse.Marshal(b, m, deterministic)
}
func (m *CreateLoginTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLoginTokensResponse.Merge(m, src)
}
func (m *CreateLoginTokensResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLoginTokensResponse.Size(m)
}
func (m *CreateLoginTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLoginTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLoginTokensResponse proto.InternalMessageInfo

func (m *CreateLoginTokensResponse) GetLoginTokenDetails() []*LoginTokenDetails {
	if m != nil {
		return m.LoginTokenDetails
	}
	return nil
}

type LoginTokenDetails struct {
	DeploymentId         string   `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	DeploymentName       string   `protobuf:"bytes,2,opt,name=deployment_name,json=deploymentName,proto3" json:"deployment_name,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	LoginToken           string   `protobuf:"bytes,4,opt,name=login_token,json=loginToken,proto3" json:"login_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginTokenDetails) Reset()         { *m = LoginTokenDetails{} }
func (m *LoginTokenDetails) String() string { return proto.CompactTextString(m) }
func (*LoginTokenDetails) ProtoMessage()    {}
func (*LoginTokenDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69bfbae40e7453c, []int{2}
}

func (m *LoginTokenDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginTokenDetails.Unmarshal(m, b)
}
func (m *LoginTokenDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginTokenDetails.Marshal(b, m, deterministic)
}
func (m *LoginTokenDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginTokenDetails.Merge(m, src)
}
func (m *LoginTokenDetails) XXX_Size() int {
	return xxx_messageInfo_LoginTokenDetails.Size(m)
}
func (m *LoginTokenDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginTokenDetails.DiscardUnknown(m)
}

var xxx_messageInfo_LoginTokenDetails proto.InternalMessageInfo

func (m *LoginTokenDetails) GetDeploymentId() string {
	if m != nil {
		return m.DeploymentId
	}
	return ""
}

func (m *LoginTokenDetails) GetDeploymentName() string {
	if m != nil {
		return m.DeploymentName
	}
	return ""
}

func (m *LoginTokenDetails) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *LoginTokenDetails) GetLoginToken() string {
	if m != nil {
		return m.LoginToken
	}
	return ""
}

type CreateDevelopmentPlayerIdentityTokenRequest struct {
	DevelopmentAuthenticationTokenSecret string             `protobuf:"bytes,1,opt,name=development_authentication_token_secret,json=developmentAuthenticationTokenSecret,proto3" json:"development_authentication_token_secret,omitempty"`
	PlayerIdentifier                     string             `protobuf:"bytes,2,opt,name=player_identifier,json=playerIdentifier,proto3" json:"player_identifier,omitempty"`
	LifetimeDuration                     *duration.Duration `protobuf:"bytes,3,opt,name=lifetime_duration,json=lifetimeDuration,proto3" json:"lifetime_duration,omitempty"`
	DisplayName                          string             `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Metadata                             []byte             `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral                 struct{}           `json:"-"`
	XXX_unrecognized                     []byte             `json:"-"`
	XXX_sizecache                        int32              `json:"-"`
}

func (m *CreateDevelopmentPlayerIdentityTokenRequest) Reset() {
	*m = CreateDevelopmentPlayerIdentityTokenRequest{}
}
func (m *CreateDevelopmentPlayerIdentityTokenRequest) String() string {
	return proto.CompactTextString(m)
}
func (*CreateDevelopmentPlayerIdentityTokenRequest) ProtoMessage() {}
func (*CreateDevelopmentPlayerIdentityTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69bfbae40e7453c, []int{3}
}

func (m *CreateDevelopmentPlayerIdentityTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenRequest.Unmarshal(m, b)
}
func (m *CreateDevelopmentPlayerIdentityTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenRequest.Marshal(b, m, deterministic)
}
func (m *CreateDevelopmentPlayerIdentityTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenRequest.Merge(m, src)
}
func (m *CreateDevelopmentPlayerIdentityTokenRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenRequest.Size(m)
}
func (m *CreateDevelopmentPlayerIdentityTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenRequest proto.InternalMessageInfo

func (m *CreateDevelopmentPlayerIdentityTokenRequest) GetDevelopmentAuthenticationTokenSecret() string {
	if m != nil {
		return m.DevelopmentAuthenticationTokenSecret
	}
	return ""
}

func (m *CreateDevelopmentPlayerIdentityTokenRequest) GetPlayerIdentifier() string {
	if m != nil {
		return m.PlayerIdentifier
	}
	return ""
}

func (m *CreateDevelopmentPlayerIdentityTokenRequest) GetLifetimeDuration() *duration.Duration {
	if m != nil {
		return m.LifetimeDuration
	}
	return nil
}

func (m *CreateDevelopmentPlayerIdentityTokenRequest) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *CreateDevelopmentPlayerIdentityTokenRequest) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CreateDevelopmentPlayerIdentityTokenResponse struct {
	PlayerIdentityToken  string   `protobuf:"bytes,1,opt,name=player_identity_token,json=playerIdentityToken,proto3" json:"player_identity_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDevelopmentPlayerIdentityTokenResponse) Reset() {
	*m = CreateDevelopmentPlayerIdentityTokenResponse{}
}
func (m *CreateDevelopmentPlayerIdentityTokenResponse) String() string {
	return proto.CompactTextString(m)
}
func (*CreateDevelopmentPlayerIdentityTokenResponse) ProtoMessage() {}
func (*CreateDevelopmentPlayerIdentityTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d69bfbae40e7453c, []int{4}
}

func (m *CreateDevelopmentPlayerIdentityTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenResponse.Unmarshal(m, b)
}
func (m *CreateDevelopmentPlayerIdentityTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenResponse.Marshal(b, m, deterministic)
}
func (m *CreateDevelopmentPlayerIdentityTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenResponse.Merge(m, src)
}
func (m *CreateDevelopmentPlayerIdentityTokenResponse) XXX_Size() int {
	return xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenResponse.Size(m)
}
func (m *CreateDevelopmentPlayerIdentityTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDevelopmentPlayerIdentityTokenResponse proto.InternalMessageInfo

func (m *CreateDevelopmentPlayerIdentityTokenResponse) GetPlayerIdentityToken() string {
	if m != nil {
		return m.PlayerIdentityToken
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateLoginTokensRequest)(nil), "improbable.spatialos.playerauth.v2.CreateLoginTokensRequest")
	proto.RegisterType((*CreateLoginTokensResponse)(nil), "improbable.spatialos.playerauth.v2.CreateLoginTokensResponse")
	proto.RegisterType((*LoginTokenDetails)(nil), "improbable.spatialos.playerauth.v2.LoginTokenDetails")
	proto.RegisterType((*CreateDevelopmentPlayerIdentityTokenRequest)(nil), "improbable.spatialos.playerauth.v2.CreateDevelopmentPlayerIdentityTokenRequest")
	proto.RegisterType((*CreateDevelopmentPlayerIdentityTokenResponse)(nil), "improbable.spatialos.playerauth.v2.CreateDevelopmentPlayerIdentityTokenResponse")
}

func init() {
	proto.RegisterFile("improbable/spatialos/playerauth/v2/devauth.proto", fileDescriptor_d69bfbae40e7453c)
}

var fileDescriptor_d69bfbae40e7453c = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x4f, 0x13, 0x4d,
	0x14, 0xce, 0xb6, 0xbc, 0xe4, 0x65, 0xca, 0x0b, 0x74, 0x08, 0xc9, 0xd2, 0x37, 0xa1, 0x75, 0xd3,
	0x68, 0x13, 0xe9, 0x2e, 0x2e, 0x6a, 0x50, 0xd4, 0x08, 0x36, 0x26, 0x24, 0x46, 0xc9, 0xc2, 0x85,
	0xc1, 0xc0, 0x66, 0xda, 0x3d, 0x5d, 0x26, 0xcc, 0xee, 0xac, 0xbb, 0xd3, 0xc5, 0xfa, 0x71, 0xe3,
	0x8f, 0xf0, 0x56, 0xc3, 0x2d, 0xff, 0xc2, 0x2b, 0xef, 0xfc, 0x09, 0x18, 0xfe, 0x86, 0x37, 0x66,
	0x3f, 0xda, 0x2d, 0x14, 0xa5, 0x1a, 0xef, 0x66, 0xce, 0x39, 0xcf, 0x73, 0xe6, 0x3c, 0xf3, 0xcc,
	0x2e, 0x5a, 0xa2, 0x8e, 0xe7, 0xf3, 0x26, 0x69, 0x32, 0xd0, 0x02, 0x8f, 0x08, 0x4a, 0x18, 0x0f,
	0x34, 0x8f, 0x91, 0x2e, 0xf8, 0xa4, 0x23, 0xf6, 0xb5, 0x50, 0xd7, 0x2c, 0x08, 0xa3, 0xa5, 0xea,
	0xf9, 0x5c, 0x70, 0xac, 0x64, 0x08, 0xb5, 0x8f, 0x50, 0x33, 0x84, 0x1a, 0xea, 0xa5, 0x55, 0x9b,
	0x8a, 0xfd, 0x4e, 0x53, 0x6d, 0x71, 0x47, 0xcb, 0xca, 0xeb, 0x94, 0x6b, 0x36, 0xaf, 0xc7, 0x34,
	0x75, 0xc6, 0xed, 0x36, 0x05, 0x66, 0x05, 0x5a, 0x7f, 0x95, 0x34, 0x28, 0xdd, 0x1e, 0x00, 0x3b,
	0x87, 0x54, 0x1c, 0xf0, 0xc3, 0x0c, 0x16, 0x12, 0x46, 0x2d, 0x22, 0xb8, 0x1f, 0x68, 0xfd, 0x65,
	0x8a, 0x5b, 0xb0, 0x39, 0xb7, 0x19, 0x68, 0xf1, 0xae, 0xd9, 0x69, 0x6b, 0x56, 0xc7, 0x27, 0x82,
	0x72, 0x37, 0xcd, 0x5f, 0x1d, 0x18, 0x15, 0x5e, 0x09, 0xcd, 0x63, 0x1d, 0x9b, 0xba, 0x5a, 0x3c,
	0xa5, 0x07, 0xbe, 0x93, 0xf6, 0x57, 0xbe, 0x49, 0x48, 0x7e, 0xe4, 0x03, 0x11, 0xf0, 0x84, 0xdb,
	0xd4, 0xdd, 0xe6, 0x07, 0xe0, 0x06, 0x06, 0xbc, 0xec, 0x40, 0x20, 0xf0, 0x5d, 0x34, 0x97, 0x8c,
	0x6a, 0x52, 0x0b, 0x5c, 0x41, 0x45, 0xd7, 0x14, 0x51, 0x81, 0x2c, 0x55, 0xa4, 0xda, 0xc4, 0xfa,
	0xf8, 0xe9, 0x49, 0x39, 0xf7, 0x5c, 0x32, 0x66, 0x93, 0xa2, 0x8d, 0xb4, 0x26, 0xe6, 0xc0, 0x8f,
	0x51, 0x91, 0xd1, 0x36, 0x08, 0xea, 0x80, 0xd9, 0x3b, 0x9b, 0x9c, 0xab, 0x48, 0xb5, 0x82, 0x3e,
	0xaf, 0x26, 0x87, 0x57, 0x7b, 0x87, 0x57, 0x1b, 0x69, 0x81, 0x31, 0xd3, 0xc3, 0xf4, 0x22, 0xf8,
	0x21, 0x2a, 0x1c, 0x72, 0xff, 0x00, 0x7c, 0x53, 0x74, 0x3d, 0x90, 0xf3, 0x71, 0xe7, 0xf2, 0xe9,
	0x49, 0xf9, 0x7f, 0x34, 0xbf, 0x57, 0x7d, 0xbb, 0xf7, 0x82, 0xd4, 0x5f, 0xaf, 0xd5, 0x77, 0x96,
	0xea, 0x77, 0xcc, 0xfa, 0xee, 0x9b, 0xe5, 0xc5, 0x1b, 0xfa, 0xca, 0xbb, 0xaa, 0x81, 0x12, 0xcc,
	0x76, 0xd7, 0x03, 0xe5, 0xbd, 0x84, 0xe6, 0x2f, 0x18, 0x31, 0xf0, 0xb8, 0x1b, 0x00, 0x06, 0x34,
	0xcb, 0xa2, 0x70, 0x32, 0x99, 0x69, 0x81, 0x20, 0x94, 0x05, 0xb2, 0x54, 0xc9, 0xd7, 0x0a, 0xfa,
	0x2d, 0xf5, 0xf2, 0xfb, 0x57, 0x33, 0xd6, 0x46, 0x02, 0x36, 0x8a, 0xec, 0x7c, 0x48, 0xf9, 0x2c,
	0xa1, 0xe2, 0x50, 0x21, 0x5e, 0x41, 0xff, 0x59, 0xe0, 0x31, 0xde, 0x75, 0xc0, 0x15, 0x26, 0xb5,
	0x52, 0x61, 0x67, 0x8f, 0x8e, 0x17, 0xa6, 0xcf, 0xa5, 0x8c, 0xc9, 0x6c, 0xbb, 0x61, 0xe1, 0x07,
	0x68, 0x7a, 0x20, 0xed, 0x12, 0x07, 0x62, 0x71, 0x27, 0xd6, 0xe7, 0x8e, 0x8e, 0x17, 0x8a, 0x43,
	0x49, 0x63, 0x2a, 0x0b, 0x3c, 0x25, 0x0e, 0x60, 0x8c, 0xc6, 0x04, 0xb1, 0x03, 0x39, 0x5f, 0xc9,
	0xd7, 0x26, 0x8c, 0x78, 0x8d, 0xcb, 0xa8, 0x30, 0x20, 0x85, 0x3c, 0x16, 0xf1, 0x19, 0x28, 0x9b,
	0x45, 0xf9, 0x92, 0x43, 0xd7, 0x13, 0x25, 0x1b, 0x10, 0x02, 0xe3, 0x5e, 0x44, 0xb7, 0x39, 0x7c,
	0xf9, 0x3d, 0xff, 0xec, 0xa2, 0x6b, 0x56, 0x56, 0x68, 0x46, 0x82, 0x45, 0x55, 0xad, 0xf8, 0x66,
	0x53, 0xc1, 0x03, 0x68, 0xf9, 0x20, 0xce, 0x39, 0xaa, 0x3a, 0x00, 0x5b, 0x3b, 0x83, 0x8a, 0xf9,
	0xb7, 0x62, 0x0c, 0x5e, 0x46, 0xc5, 0x33, 0xf6, 0x6c, 0x53, 0xf0, 0x53, 0x15, 0x7a, 0x44, 0x33,
	0x83, 0xd6, 0x8c, 0xf2, 0x17, 0xfb, 0x32, 0xff, 0xfb, 0xbe, 0xbc, 0x82, 0x26, 0x2d, 0x1a, 0x44,
	0xf4, 0x89, 0xfa, 0x89, 0x5a, 0x85, 0x34, 0x16, 0x6b, 0x5c, 0x42, 0xff, 0x3a, 0x20, 0x88, 0x45,
	0x04, 0x91, 0xff, 0xa9, 0x48, 0xb5, 0x49, 0xa3, 0xbf, 0x57, 0x9a, 0x68, 0x71, 0x34, 0x25, 0x53,
	0x9b, 0xea, 0xbf, 0x7c, 0x8a, 0x17, 0x3e, 0x41, 0xfd, 0x7b, 0x0e, 0x4d, 0x35, 0x20, 0x8c, 0x04,
	0xdc, 0x02, 0x3f, 0xa4, 0x2d, 0xc0, 0x1f, 0x24, 0x54, 0x1c, 0x7a, 0x0b, 0xf8, 0xde, 0x28, 0x36,
	0xff, 0xd9, 0x57, 0xa2, 0x74, 0xff, 0x0f, 0xd1, 0xc9, 0x64, 0xca, 0xd8, 0xa7, 0x8f, 0x8a, 0x84,
	0xbf, 0x4a, 0xa8, 0x3a, 0x8a, 0x20, 0xf8, 0xd9, 0xe8, 0xdd, 0x46, 0x32, 0x69, 0x69, 0xf3, 0xef,
	0x11, 0x0e, 0x4e, 0xb4, 0x7e, 0x73, 0x47, 0xbf, 0xfc, 0x77, 0xb3, 0xea, 0x35, 0xcd, 0x2c, 0x10,
	0xea, 0xcd, 0xf1, 0xd8, 0x7b, 0xcb, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4e, 0xad, 0xe3, 0x81,
	0xab, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DevAuthServiceClient is the client API for DevAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DevAuthServiceClient interface {
	CreateLoginTokens(ctx context.Context, in *CreateLoginTokensRequest, opts ...grpc.CallOption) (*CreateLoginTokensResponse, error)
	CreateDevelopmentPlayerIdentityToken(ctx context.Context, in *CreateDevelopmentPlayerIdentityTokenRequest, opts ...grpc.CallOption) (*CreateDevelopmentPlayerIdentityTokenResponse, error)
}

type devAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevAuthServiceClient(cc grpc.ClientConnInterface) DevAuthServiceClient {
	return &devAuthServiceClient{cc}
}

func (c *devAuthServiceClient) CreateLoginTokens(ctx context.Context, in *CreateLoginTokensRequest, opts ...grpc.CallOption) (*CreateLoginTokensResponse, error) {
	out := new(CreateLoginTokensResponse)
	err := c.cc.Invoke(ctx, "/improbable.spatialos.playerauth.v2.DevAuthService/CreateLoginTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devAuthServiceClient) CreateDevelopmentPlayerIdentityToken(ctx context.Context, in *CreateDevelopmentPlayerIdentityTokenRequest, opts ...grpc.CallOption) (*CreateDevelopmentPlayerIdentityTokenResponse, error) {
	out := new(CreateDevelopmentPlayerIdentityTokenResponse)
	err := c.cc.Invoke(ctx, "/improbable.spatialos.playerauth.v2.DevAuthService/CreateDevelopmentPlayerIdentityToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevAuthServiceServer is the server API for DevAuthService service.
type DevAuthServiceServer interface {
	CreateLoginTokens(context.Context, *CreateLoginTokensRequest) (*CreateLoginTokensResponse, error)
	CreateDevelopmentPlayerIdentityToken(context.Context, *CreateDevelopmentPlayerIdentityTokenRequest) (*CreateDevelopmentPlayerIdentityTokenResponse, error)
}

// UnimplementedDevAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDevAuthServiceServer struct {
}

func (*UnimplementedDevAuthServiceServer) CreateLoginTokens(ctx context.Context, req *CreateLoginTokensRequest) (*CreateLoginTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLoginTokens not implemented")
}
func (*UnimplementedDevAuthServiceServer) CreateDevelopmentPlayerIdentityToken(ctx context.Context, req *CreateDevelopmentPlayerIdentityTokenRequest) (*CreateDevelopmentPlayerIdentityTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevelopmentPlayerIdentityToken not implemented")
}

func RegisterDevAuthServiceServer(s *grpc.Server, srv DevAuthServiceServer) {
	s.RegisterService(&_DevAuthService_serviceDesc, srv)
}

func _DevAuthService_CreateLoginTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLoginTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevAuthServiceServer).CreateLoginTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.spatialos.playerauth.v2.DevAuthService/CreateLoginTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevAuthServiceServer).CreateLoginTokens(ctx, req.(*CreateLoginTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevAuthService_CreateDevelopmentPlayerIdentityToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDevelopmentPlayerIdentityTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevAuthServiceServer).CreateDevelopmentPlayerIdentityToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.spatialos.playerauth.v2.DevAuthService/CreateDevelopmentPlayerIdentityToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevAuthServiceServer).CreateDevelopmentPlayerIdentityToken(ctx, req.(*CreateDevelopmentPlayerIdentityTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevAuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.spatialos.playerauth.v2.DevAuthService",
	HandlerType: (*DevAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLoginTokens",
			Handler:    _DevAuthService_CreateLoginTokens_Handler,
		},
		{
			MethodName: "CreateDevelopmentPlayerIdentityToken",
			Handler:    _DevAuthService_CreateDevelopmentPlayerIdentityToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/spatialos/playerauth/v2/devauth.proto",
}