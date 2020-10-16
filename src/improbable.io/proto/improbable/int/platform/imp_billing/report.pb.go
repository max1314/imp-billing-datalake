// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/int/platform/imp_billing/report.proto

package improbable_int_platform_imp_billing

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ItemCharge int32

const (
	ItemCharge_RUNTIME ItemCharge = 0
	ItemCharge_WORKER  ItemCharge = 1
	ItemCharge_EGRESS  ItemCharge = 2
)

var ItemCharge_name = map[int32]string{
	0: "RUNTIME",
	1: "WORKER",
	2: "EGRESS",
}

var ItemCharge_value = map[string]int32{
	"RUNTIME": 0,
	"WORKER":  1,
	"EGRESS":  2,
}

func (x ItemCharge) String() string {
	return proto.EnumName(ItemCharge_name, int32(x))
}

func (ItemCharge) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{0}
}

type Provider int32

const (
	Provider_GOOGLE     Provider = 0
	Provider_AWS        Provider = 1
	Provider_TENCENT    Provider = 2
	Provider_BARE_METAL Provider = 3
)

var Provider_name = map[int32]string{
	0: "GOOGLE",
	1: "AWS",
	2: "TENCENT",
	3: "BARE_METAL",
}

var Provider_value = map[string]int32{
	"GOOGLE":     0,
	"AWS":        1,
	"TENCENT":    2,
	"BARE_METAL": 3,
}

func (x Provider) String() string {
	return proto.EnumName(Provider_name, int32(x))
}

func (Provider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{1}
}

type GroupBy int32

const (
	GroupBy_GROUPBY_UNSPECIFIED GroupBy = 0
	GroupBy_CLUSTER             GroupBy = 1
	GroupBy_REGION              GroupBy = 2
	GroupBy_PROJECT             GroupBy = 3
	GroupBy_ITEM_CHARGE         GroupBy = 4
	GroupBy_STOCK_KEEPING_UNIT  GroupBy = 5
	GroupBy_PROVIDER            GroupBy = 6
	GroupBy_ENVIRONMENT         GroupBy = 7
)

var GroupBy_name = map[int32]string{
	0: "GROUPBY_UNSPECIFIED",
	1: "CLUSTER",
	2: "REGION",
	3: "PROJECT",
	4: "ITEM_CHARGE",
	5: "STOCK_KEEPING_UNIT",
	6: "PROVIDER",
	7: "ENVIRONMENT",
}

var GroupBy_value = map[string]int32{
	"GROUPBY_UNSPECIFIED": 0,
	"CLUSTER":             1,
	"REGION":              2,
	"PROJECT":             3,
	"ITEM_CHARGE":         4,
	"STOCK_KEEPING_UNIT":  5,
	"PROVIDER":            6,
	"ENVIRONMENT":         7,
}

func (x GroupBy) String() string {
	return proto.EnumName(GroupBy_name, int32(x))
}

func (GroupBy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{2}
}

type TimeGranularity int32

const (
	TimeGranularity_MONTHLY TimeGranularity = 0
	TimeGranularity_DAILY   TimeGranularity = 1
)

var TimeGranularity_name = map[int32]string{
	0: "MONTHLY",
	1: "DAILY",
}

var TimeGranularity_value = map[string]int32{
	"MONTHLY": 0,
	"DAILY":   1,
}

func (x TimeGranularity) String() string {
	return proto.EnumName(TimeGranularity_name, int32(x))
}

func (TimeGranularity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{3}
}

type Money struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Money) Reset()         { *m = Money{} }
func (m *Money) String() string { return proto.CompactTextString(m) }
func (*Money) ProtoMessage()    {}
func (*Money) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{0}
}

func (m *Money) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Money.Unmarshal(m, b)
}
func (m *Money) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Money.Marshal(b, m, deterministic)
}
func (m *Money) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Money.Merge(m, src)
}
func (m *Money) XXX_Size() int {
	return xxx_messageInfo_Money.Size(m)
}
func (m *Money) XXX_DiscardUnknown() {
	xxx_messageInfo_Money.DiscardUnknown(m)
}

var xxx_messageInfo_Money proto.InternalMessageInfo

func (m *Money) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Money) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type Group struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	TotalPrice           *Money   `protobuf:"bytes,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Group) Reset()         { *m = Group{} }
func (m *Group) String() string { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()    {}
func (*Group) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{1}
}

func (m *Group) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Group.Unmarshal(m, b)
}
func (m *Group) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Group.Marshal(b, m, deterministic)
}
func (m *Group) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Group.Merge(m, src)
}
func (m *Group) XXX_Size() int {
	return xxx_messageInfo_Group.Size(m)
}
func (m *Group) XXX_DiscardUnknown() {
	xxx_messageInfo_Group.DiscardUnknown(m)
}

var xxx_messageInfo_Group proto.InternalMessageInfo

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetTotalPrice() *Money {
	if m != nil {
		return m.TotalPrice
	}
	return nil
}

type Report struct {
	IsForecast           bool                 `protobuf:"varint,1,opt,name=is_forecast,json=isForecast,proto3" json:"is_forecast,omitempty"`
	BeginTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	TotalPrice           *Money               `protobuf:"bytes,4,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
	Groups               []*Group             `protobuf:"bytes,5,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Report) Reset()         { *m = Report{} }
func (m *Report) String() string { return proto.CompactTextString(m) }
func (*Report) ProtoMessage()    {}
func (*Report) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{2}
}

func (m *Report) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Report.Unmarshal(m, b)
}
func (m *Report) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Report.Marshal(b, m, deterministic)
}
func (m *Report) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Report.Merge(m, src)
}
func (m *Report) XXX_Size() int {
	return xxx_messageInfo_Report.Size(m)
}
func (m *Report) XXX_DiscardUnknown() {
	xxx_messageInfo_Report.DiscardUnknown(m)
}

var xxx_messageInfo_Report proto.InternalMessageInfo

func (m *Report) GetIsForecast() bool {
	if m != nil {
		return m.IsForecast
	}
	return false
}

func (m *Report) GetBeginTime() *timestamp.Timestamp {
	if m != nil {
		return m.BeginTime
	}
	return nil
}

func (m *Report) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Report) GetTotalPrice() *Money {
	if m != nil {
		return m.TotalPrice
	}
	return nil
}

func (m *Report) GetGroups() []*Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

type ListReportsRequest struct {
	ItemCharges          []ItemCharge         `protobuf:"varint,1,rep,packed,name=item_charges,json=itemCharges,proto3,enum=improbable.int.platform.imp_billing.ItemCharge" json:"item_charges,omitempty"`
	Providers            []Provider           `protobuf:"varint,2,rep,packed,name=providers,proto3,enum=improbable.int.platform.imp_billing.Provider" json:"providers,omitempty"`
	StockKeepingUnits    []string             `protobuf:"bytes,3,rep,name=stock_keeping_units,json=stockKeepingUnits,proto3" json:"stock_keeping_units,omitempty"`
	Projects             []string             `protobuf:"bytes,4,rep,name=projects,proto3" json:"projects,omitempty"`
	Environments         []string             `protobuf:"bytes,5,rep,name=environments,proto3" json:"environments,omitempty"`
	Clusters             []string             `protobuf:"bytes,6,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Regions              []string             `protobuf:"bytes,7,rep,name=regions,proto3" json:"regions,omitempty"`
	CustomerTags         []string             `protobuf:"bytes,8,rep,name=customer_tags,json=customerTags,proto3" json:"customer_tags,omitempty"`
	Currency             string               `protobuf:"bytes,9,opt,name=currency,proto3" json:"currency,omitempty"`
	BeginTime            *timestamp.Timestamp `protobuf:"bytes,10,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,11,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	GroupBy              GroupBy              `protobuf:"varint,12,opt,name=group_by,json=groupBy,proto3,enum=improbable.int.platform.imp_billing.GroupBy" json:"group_by,omitempty"`
	TimeGranularity      TimeGranularity      `protobuf:"varint,13,opt,name=time_granularity,json=timeGranularity,proto3,enum=improbable.int.platform.imp_billing.TimeGranularity" json:"time_granularity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListReportsRequest) Reset()         { *m = ListReportsRequest{} }
func (m *ListReportsRequest) String() string { return proto.CompactTextString(m) }
func (*ListReportsRequest) ProtoMessage()    {}
func (*ListReportsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{3}
}

func (m *ListReportsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReportsRequest.Unmarshal(m, b)
}
func (m *ListReportsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReportsRequest.Marshal(b, m, deterministic)
}
func (m *ListReportsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReportsRequest.Merge(m, src)
}
func (m *ListReportsRequest) XXX_Size() int {
	return xxx_messageInfo_ListReportsRequest.Size(m)
}
func (m *ListReportsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReportsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListReportsRequest proto.InternalMessageInfo

func (m *ListReportsRequest) GetItemCharges() []ItemCharge {
	if m != nil {
		return m.ItemCharges
	}
	return nil
}

func (m *ListReportsRequest) GetProviders() []Provider {
	if m != nil {
		return m.Providers
	}
	return nil
}

func (m *ListReportsRequest) GetStockKeepingUnits() []string {
	if m != nil {
		return m.StockKeepingUnits
	}
	return nil
}

func (m *ListReportsRequest) GetProjects() []string {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *ListReportsRequest) GetEnvironments() []string {
	if m != nil {
		return m.Environments
	}
	return nil
}

func (m *ListReportsRequest) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *ListReportsRequest) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *ListReportsRequest) GetCustomerTags() []string {
	if m != nil {
		return m.CustomerTags
	}
	return nil
}

func (m *ListReportsRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *ListReportsRequest) GetBeginTime() *timestamp.Timestamp {
	if m != nil {
		return m.BeginTime
	}
	return nil
}

func (m *ListReportsRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *ListReportsRequest) GetGroupBy() GroupBy {
	if m != nil {
		return m.GroupBy
	}
	return GroupBy_GROUPBY_UNSPECIFIED
}

func (m *ListReportsRequest) GetTimeGranularity() TimeGranularity {
	if m != nil {
		return m.TimeGranularity
	}
	return TimeGranularity_MONTHLY
}

type ListReportsResponse struct {
	Reports              []*Report `protobuf:"bytes,1,rep,name=reports,proto3" json:"reports,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListReportsResponse) Reset()         { *m = ListReportsResponse{} }
func (m *ListReportsResponse) String() string { return proto.CompactTextString(m) }
func (*ListReportsResponse) ProtoMessage()    {}
func (*ListReportsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e184a4face85823, []int{4}
}

func (m *ListReportsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListReportsResponse.Unmarshal(m, b)
}
func (m *ListReportsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListReportsResponse.Marshal(b, m, deterministic)
}
func (m *ListReportsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListReportsResponse.Merge(m, src)
}
func (m *ListReportsResponse) XXX_Size() int {
	return xxx_messageInfo_ListReportsResponse.Size(m)
}
func (m *ListReportsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListReportsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListReportsResponse proto.InternalMessageInfo

func (m *ListReportsResponse) GetReports() []*Report {
	if m != nil {
		return m.Reports
	}
	return nil
}

func init() {
	proto.RegisterEnum("improbable.int.platform.imp_billing.ItemCharge", ItemCharge_name, ItemCharge_value)
	proto.RegisterEnum("improbable.int.platform.imp_billing.Provider", Provider_name, Provider_value)
	proto.RegisterEnum("improbable.int.platform.imp_billing.GroupBy", GroupBy_name, GroupBy_value)
	proto.RegisterEnum("improbable.int.platform.imp_billing.TimeGranularity", TimeGranularity_name, TimeGranularity_value)
	proto.RegisterType((*Money)(nil), "improbable.int.platform.imp_billing.Money")
	proto.RegisterType((*Group)(nil), "improbable.int.platform.imp_billing.Group")
	proto.RegisterType((*Report)(nil), "improbable.int.platform.imp_billing.Report")
	proto.RegisterType((*ListReportsRequest)(nil), "improbable.int.platform.imp_billing.ListReportsRequest")
	proto.RegisterType((*ListReportsResponse)(nil), "improbable.int.platform.imp_billing.ListReportsResponse")
}

func init() {
	proto.RegisterFile("improbable/int/platform/imp_billing/report.proto", fileDescriptor_0e184a4face85823)
}

var fileDescriptor_0e184a4face85823 = []byte{
	// 893 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xc7, 0xeb, 0xa4, 0xf9, 0x3a, 0xe9, 0xb6, 0x66, 0x2a, 0x2d, 0x56, 0xb5, 0x62, 0xa3, 0x2c,
	0x17, 0xa5, 0xd0, 0x84, 0x16, 0x10, 0x20, 0xb8, 0x69, 0x53, 0x6f, 0xd6, 0x24, 0xb1, 0xa3, 0x89,
	0xd3, 0x55, 0x25, 0x24, 0xcb, 0x49, 0xa7, 0xde, 0x61, 0xe3, 0x19, 0x33, 0x33, 0xa9, 0xd4, 0x5b,
	0x9e, 0x00, 0x89, 0x07, 0xe0, 0x05, 0xb8, 0xe3, 0x86, 0xf7, 0xe0, 0x15, 0x78, 0x10, 0x34, 0xe3,
	0xb8, 0x25, 0xe5, 0x82, 0x00, 0x77, 0x3e, 0xe7, 0xcc, 0xef, 0x9c, 0x99, 0x73, 0xfe, 0xc7, 0xf0,
	0x31, 0x4d, 0x33, 0xc1, 0x67, 0xf1, 0x6c, 0x41, 0xba, 0x94, 0xa9, 0x6e, 0xb6, 0x88, 0xd5, 0x0d,
	0x17, 0x69, 0x97, 0xa6, 0x59, 0x34, 0xa3, 0x8b, 0x05, 0x65, 0x49, 0x57, 0x90, 0x8c, 0x0b, 0xd5,
	0xc9, 0x04, 0x57, 0x1c, 0xbd, 0x78, 0x20, 0x3a, 0x94, 0xa9, 0x4e, 0x41, 0x74, 0xfe, 0x42, 0x1c,
	0x3c, 0x4b, 0x38, 0x4f, 0x16, 0xa4, 0x1b, 0x67, 0xb4, 0x1b, 0x33, 0xc6, 0x55, 0xac, 0x28, 0x67,
	0x32, 0x4f, 0x71, 0xf0, 0x7c, 0x15, 0x35, 0xd6, 0x6c, 0x79, 0xd3, 0x55, 0x34, 0x25, 0x52, 0xc5,
	0x69, 0x96, 0x1f, 0x68, 0x7f, 0x05, 0x95, 0x11, 0x67, 0xe4, 0x0e, 0x3d, 0x85, 0x6a, 0x9c, 0xf2,
	0x25, 0x53, 0x8e, 0xd5, 0xb2, 0x0e, 0x2d, 0xbc, 0xb2, 0xd0, 0x01, 0xd4, 0xe7, 0x4b, 0x21, 0x08,
	0x9b, 0xdf, 0x39, 0xa5, 0x96, 0x75, 0xd8, 0xc0, 0xf7, 0x76, 0xfb, 0x0d, 0x54, 0xfa, 0x82, 0x2f,
	0x33, 0x84, 0x60, 0x9b, 0xc5, 0x29, 0x31, 0x68, 0x03, 0x9b, 0x6f, 0x34, 0x80, 0xa6, 0xe2, 0x2a,
	0x5e, 0x44, 0x99, 0xa0, 0x73, 0x62, 0xd8, 0xe6, 0xe9, 0x51, 0x67, 0x83, 0x37, 0x75, 0xcc, 0x8d,
	0x30, 0x18, 0x7c, 0xac, 0xe9, 0xf6, 0x2f, 0x25, 0xa8, 0x62, 0xd3, 0x1b, 0xf4, 0x1c, 0x9a, 0x54,
	0x46, 0x37, 0x5c, 0x90, 0x79, 0x2c, 0xf3, 0xdb, 0xd6, 0x31, 0x50, 0xf9, 0x72, 0xe5, 0x41, 0x5f,
	0x02, 0xcc, 0x48, 0x42, 0x59, 0xa4, 0xdf, 0xba, 0xaa, 0x7b, 0xd0, 0xc9, 0x1b, 0xd1, 0x29, 0x1a,
	0xd1, 0x09, 0x8b, 0x46, 0xe0, 0x86, 0x39, 0xad, 0x6d, 0xf4, 0x19, 0xd4, 0x09, 0xbb, 0xce, 0xc1,
	0xf2, 0x3f, 0x82, 0x35, 0xc2, 0xae, 0x0d, 0xf6, 0xe8, 0xa9, 0xdb, 0xff, 0xe7, 0xa9, 0xe8, 0x1c,
	0xaa, 0x89, 0x6e, 0xaa, 0x74, 0x2a, 0xad, 0xf2, 0xc6, 0x79, 0xcc, 0x1c, 0xf0, 0x8a, 0x6c, 0xff,
	0x56, 0x01, 0x34, 0xa4, 0x52, 0xe5, 0x2d, 0x93, 0x98, 0x7c, 0xbf, 0x24, 0x52, 0x21, 0x0c, 0x3b,
	0x54, 0x91, 0x34, 0x9a, 0xbf, 0x89, 0x45, 0x42, 0xa4, 0x63, 0xb5, 0xca, 0x87, 0xbb, 0xa7, 0xdd,
	0x8d, 0x0a, 0x78, 0x8a, 0xa4, 0x3d, 0xc3, 0xe1, 0x26, 0xbd, 0xff, 0x96, 0x68, 0x00, 0x8d, 0x4c,
	0xf0, 0x5b, 0x7a, 0x4d, 0x84, 0x74, 0x4a, 0x26, 0xe1, 0xf1, 0x46, 0x09, 0xc7, 0x2b, 0x0a, 0x3f,
	0xf0, 0xa8, 0x03, 0xfb, 0x52, 0xf1, 0xf9, 0xdb, 0xe8, 0x2d, 0x21, 0x19, 0x65, 0x49, 0xb4, 0x64,
	0x54, 0x49, 0xa7, 0xdc, 0x2a, 0x1f, 0x36, 0xf0, 0x3b, 0x26, 0x34, 0xc8, 0x23, 0x53, 0x1d, 0xd0,
	0xe2, 0xcc, 0x04, 0xff, 0x8e, 0xcc, 0x95, 0x74, 0xb6, 0xcd, 0xa1, 0x7b, 0x1b, 0xb5, 0x61, 0x87,
	0xb0, 0x5b, 0x2a, 0x38, 0x4b, 0x09, 0x53, 0x79, 0x37, 0x1b, 0x78, 0xcd, 0x67, 0xc4, 0xbd, 0x58,
	0x4a, 0xa5, 0xef, 0x5e, 0xcd, 0xf9, 0xc2, 0x46, 0x0e, 0xd4, 0x04, 0x49, 0xf4, 0x2e, 0x39, 0x35,
	0x13, 0x2a, 0x4c, 0xf4, 0x02, 0x9e, 0xcc, 0x97, 0x52, 0xf1, 0x94, 0x88, 0x48, 0xc5, 0x89, 0x74,
	0xea, 0x79, 0xea, 0xc2, 0x19, 0xc6, 0x89, 0x5c, 0xdb, 0x9b, 0xc6, 0xfa, 0xde, 0x3c, 0x52, 0x28,
	0xfc, 0x57, 0x85, 0x36, 0x37, 0x57, 0x68, 0x1f, 0xea, 0x46, 0x1a, 0xd1, 0xec, 0xce, 0xd9, 0x69,
	0x59, 0x87, 0xbb, 0xa7, 0x1f, 0x6d, 0x2e, 0xab, 0xf3, 0x3b, 0x5c, 0x4b, 0xf2, 0x0f, 0x14, 0x81,
	0xad, 0x6b, 0x47, 0x89, 0x88, 0xd9, 0x72, 0x11, 0x0b, 0xaa, 0xee, 0x9c, 0x27, 0x26, 0xe1, 0xa7,
	0x1b, 0x25, 0xd4, 0xb7, 0xe9, 0x3f, 0xb0, 0x78, 0x4f, 0xad, 0x3b, 0xda, 0xdf, 0xc2, 0xfe, 0x9a,
	0x72, 0x65, 0xc6, 0x99, 0x24, 0xc8, 0xd5, 0xd3, 0x30, 0x2e, 0xa3, 0xda, 0xe6, 0xe9, 0x87, 0x1b,
	0x95, 0xcb, 0xd3, 0xe0, 0x82, 0x3d, 0x3a, 0x01, 0x78, 0x10, 0x32, 0x6a, 0x42, 0x0d, 0x4f, 0xfd,
	0xd0, 0x1b, 0xb9, 0xf6, 0x16, 0x02, 0xa8, 0xbe, 0x0e, 0xf0, 0xc0, 0xc5, 0xb6, 0xa5, 0xbf, 0xdd,
	0x3e, 0x76, 0x27, 0x13, 0xbb, 0x74, 0xf4, 0x35, 0xd4, 0x0b, 0xa9, 0x6a, 0x7f, 0x3f, 0x08, 0xfa,
	0x43, 0x7d, 0xbe, 0x06, 0xe5, 0xb3, 0xd7, 0x13, 0xdb, 0xd2, 0x59, 0x42, 0xd7, 0xef, 0xb9, 0x7e,
	0x68, 0x97, 0xd0, 0x2e, 0xc0, 0xf9, 0x19, 0x76, 0xa3, 0x91, 0x1b, 0x9e, 0x0d, 0xed, 0xf2, 0xd1,
	0x8f, 0x16, 0xd4, 0x56, 0x4d, 0x44, 0xef, 0xc2, 0x7e, 0x1f, 0x07, 0xd3, 0xf1, 0xf9, 0x55, 0x34,
	0xf5, 0x27, 0x63, 0xb7, 0xe7, 0xbd, 0xf4, 0xdc, 0x0b, 0x7b, 0x4b, 0x67, 0xe8, 0x0d, 0xa7, 0x93,
	0xb0, 0xa8, 0x8d, 0xdd, 0xbe, 0x17, 0xf8, 0x76, 0x49, 0x07, 0xc6, 0x38, 0xf8, 0xc6, 0xed, 0x85,
	0x76, 0x19, 0xed, 0x41, 0xd3, 0x0b, 0xdd, 0x51, 0xd4, 0x7b, 0x75, 0x86, 0xfb, 0xae, 0xbd, 0x8d,
	0x9e, 0x02, 0x9a, 0x84, 0x41, 0x6f, 0x10, 0x0d, 0x5c, 0x77, 0xec, 0xf9, 0xfd, 0x68, 0xea, 0x7b,
	0xa1, 0x5d, 0x41, 0x3b, 0x50, 0x1f, 0xe3, 0xe0, 0xd2, 0xbb, 0x70, 0xb1, 0x5d, 0xd5, 0x98, 0xeb,
	0x5f, 0x7a, 0x38, 0xf0, 0x47, 0xfa, 0x8a, 0xb5, 0xa3, 0x0f, 0x60, 0xef, 0xd1, 0x14, 0x74, 0x9d,
	0x51, 0xe0, 0x87, 0xaf, 0x86, 0x57, 0xf6, 0x16, 0x6a, 0x40, 0xe5, 0xe2, 0xcc, 0x1b, 0x5e, 0xd9,
	0xd6, 0xe9, 0xaf, 0x16, 0xec, 0xe5, 0x2d, 0x9c, 0x10, 0x71, 0x4b, 0xe7, 0xe4, 0xf2, 0x04, 0xfd,
	0x6c, 0xc1, 0xb6, 0x9e, 0x10, 0xfa, 0x7c, 0xa3, 0x09, 0xfc, 0xfd, 0x37, 0x74, 0xf0, 0xc5, 0xbf,
	0x07, 0x73, 0x15, 0xb4, 0xdf, 0xff, 0xe1, 0xf7, 0x3f, 0x7e, 0x2a, 0xbd, 0x87, 0x9e, 0x75, 0x6f,
	0x4f, 0xba, 0xc5, 0xba, 0x1d, 0xaf, 0x88, 0xe3, 0xd5, 0x90, 0x67, 0x55, 0xb3, 0x09, 0x9f, 0xfc,
	0x19, 0x00, 0x00, 0xff, 0xff, 0xca, 0xa5, 0x20, 0xc3, 0x72, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReportServiceV1Client is the client API for ReportServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReportServiceV1Client interface {
	List(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error)
}

type reportServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewReportServiceV1Client(cc grpc.ClientConnInterface) ReportServiceV1Client {
	return &reportServiceV1Client{cc}
}

func (c *reportServiceV1Client) List(ctx context.Context, in *ListReportsRequest, opts ...grpc.CallOption) (*ListReportsResponse, error) {
	out := new(ListReportsResponse)
	err := c.cc.Invoke(ctx, "/improbable.int.platform.imp_billing.ReportServiceV1/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportServiceV1Server is the server API for ReportServiceV1 service.
type ReportServiceV1Server interface {
	List(context.Context, *ListReportsRequest) (*ListReportsResponse, error)
}

// UnimplementedReportServiceV1Server can be embedded to have forward compatible implementations.
type UnimplementedReportServiceV1Server struct {
}

func (*UnimplementedReportServiceV1Server) List(ctx context.Context, req *ListReportsRequest) (*ListReportsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterReportServiceV1Server(s *grpc.Server, srv ReportServiceV1Server) {
	s.RegisterService(&_ReportServiceV1_serviceDesc, srv)
}

func _ReportServiceV1_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReportsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportServiceV1Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/improbable.int.platform.imp_billing.ReportServiceV1/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportServiceV1Server).List(ctx, req.(*ListReportsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReportServiceV1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "improbable.int.platform.imp_billing.ReportServiceV1",
	HandlerType: (*ReportServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ReportServiceV1_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "improbable/int/platform/imp_billing/report.proto",
}
