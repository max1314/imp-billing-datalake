// Code generated by protoc-gen-go. DO NOT EDIT.
// source: improbable/fabric/regisseur/scenario_result.proto

package improbable_regisseur

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	history "improbable.io/proto/improbable/platform/history"
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

type ScenarioResult struct {
	Identifier           *ScenarioIdentifier  `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Metrics              []*CapturedMetric    `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	ExpectationResults   []*ExpectationResult `protobuf:"bytes,3,rep,name=expectation_results,json=expectationResults,proto3" json:"expectation_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ScenarioResult) Reset()         { *m = ScenarioResult{} }
func (m *ScenarioResult) String() string { return proto.CompactTextString(m) }
func (*ScenarioResult) ProtoMessage()    {}
func (*ScenarioResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{0}
}

func (m *ScenarioResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScenarioResult.Unmarshal(m, b)
}
func (m *ScenarioResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScenarioResult.Marshal(b, m, deterministic)
}
func (m *ScenarioResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScenarioResult.Merge(m, src)
}
func (m *ScenarioResult) XXX_Size() int {
	return xxx_messageInfo_ScenarioResult.Size(m)
}
func (m *ScenarioResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ScenarioResult.DiscardUnknown(m)
}

var xxx_messageInfo_ScenarioResult proto.InternalMessageInfo

func (m *ScenarioResult) GetIdentifier() *ScenarioIdentifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *ScenarioResult) GetMetrics() []*CapturedMetric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *ScenarioResult) GetExpectationResults() []*ExpectationResult {
	if m != nil {
		return m.ExpectationResults
	}
	return nil
}

type CapturedMetric struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PrometheusExpression string              `protobuf:"bytes,2,opt,name=prometheus_expression,json=prometheusExpression,proto3" json:"prometheus_expression,omitempty"`
	PrometheusResults    []*PrometheusResult `protobuf:"bytes,3,rep,name=prometheus_results,json=prometheusResults,proto3" json:"prometheus_results,omitempty"`
	Fallback             *Metric_Fallback    `protobuf:"bytes,4,opt,name=fallback,proto3" json:"fallback,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CapturedMetric) Reset()         { *m = CapturedMetric{} }
func (m *CapturedMetric) String() string { return proto.CompactTextString(m) }
func (*CapturedMetric) ProtoMessage()    {}
func (*CapturedMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{1}
}

func (m *CapturedMetric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CapturedMetric.Unmarshal(m, b)
}
func (m *CapturedMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CapturedMetric.Marshal(b, m, deterministic)
}
func (m *CapturedMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapturedMetric.Merge(m, src)
}
func (m *CapturedMetric) XXX_Size() int {
	return xxx_messageInfo_CapturedMetric.Size(m)
}
func (m *CapturedMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_CapturedMetric.DiscardUnknown(m)
}

var xxx_messageInfo_CapturedMetric proto.InternalMessageInfo

func (m *CapturedMetric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CapturedMetric) GetPrometheusExpression() string {
	if m != nil {
		return m.PrometheusExpression
	}
	return ""
}

func (m *CapturedMetric) GetPrometheusResults() []*PrometheusResult {
	if m != nil {
		return m.PrometheusResults
	}
	return nil
}

func (m *CapturedMetric) GetFallback() *Metric_Fallback {
	if m != nil {
		return m.Fallback
	}
	return nil
}

type PrometheusResult struct {
	LabelSet             string        `protobuf:"bytes,1,opt,name=label_set,json=labelSet,proto3" json:"label_set,omitempty"`
	Samples              []*SamplePair `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
	LabelPairs           []*LabelPair  `protobuf:"bytes,3,rep,name=label_pairs,json=labelPairs,proto3" json:"label_pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PrometheusResult) Reset()         { *m = PrometheusResult{} }
func (m *PrometheusResult) String() string { return proto.CompactTextString(m) }
func (*PrometheusResult) ProtoMessage()    {}
func (*PrometheusResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{2}
}

func (m *PrometheusResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrometheusResult.Unmarshal(m, b)
}
func (m *PrometheusResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrometheusResult.Marshal(b, m, deterministic)
}
func (m *PrometheusResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrometheusResult.Merge(m, src)
}
func (m *PrometheusResult) XXX_Size() int {
	return xxx_messageInfo_PrometheusResult.Size(m)
}
func (m *PrometheusResult) XXX_DiscardUnknown() {
	xxx_messageInfo_PrometheusResult.DiscardUnknown(m)
}

var xxx_messageInfo_PrometheusResult proto.InternalMessageInfo

func (m *PrometheusResult) GetLabelSet() string {
	if m != nil {
		return m.LabelSet
	}
	return ""
}

func (m *PrometheusResult) GetSamples() []*SamplePair {
	if m != nil {
		return m.Samples
	}
	return nil
}

func (m *PrometheusResult) GetLabelPairs() []*LabelPair {
	if m != nil {
		return m.LabelPairs
	}
	return nil
}

type LabelPair struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelPair) Reset()         { *m = LabelPair{} }
func (m *LabelPair) String() string { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()    {}
func (*LabelPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{3}
}

func (m *LabelPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelPair.Unmarshal(m, b)
}
func (m *LabelPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelPair.Marshal(b, m, deterministic)
}
func (m *LabelPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelPair.Merge(m, src)
}
func (m *LabelPair) XXX_Size() int {
	return xxx_messageInfo_LabelPair.Size(m)
}
func (m *LabelPair) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelPair.DiscardUnknown(m)
}

var xxx_messageInfo_LabelPair proto.InternalMessageInfo

func (m *LabelPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LabelPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SamplePair struct {
	Timestamp            *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Value                float64              `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SamplePair) Reset()         { *m = SamplePair{} }
func (m *SamplePair) String() string { return proto.CompactTextString(m) }
func (*SamplePair) ProtoMessage()    {}
func (*SamplePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{4}
}

func (m *SamplePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SamplePair.Unmarshal(m, b)
}
func (m *SamplePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SamplePair.Marshal(b, m, deterministic)
}
func (m *SamplePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SamplePair.Merge(m, src)
}
func (m *SamplePair) XXX_Size() int {
	return xxx_messageInfo_SamplePair.Size(m)
}
func (m *SamplePair) XXX_DiscardUnknown() {
	xxx_messageInfo_SamplePair.DiscardUnknown(m)
}

var xxx_messageInfo_SamplePair proto.InternalMessageInfo

func (m *SamplePair) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *SamplePair) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type CapturedSnapshot struct {
	Metadata             *history.Snapshot `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	EntityIds            []int64           `protobuf:"varint,2,rep,packed,name=entity_ids,json=entityIds,proto3" json:"entity_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CapturedSnapshot) Reset()         { *m = CapturedSnapshot{} }
func (m *CapturedSnapshot) String() string { return proto.CompactTextString(m) }
func (*CapturedSnapshot) ProtoMessage()    {}
func (*CapturedSnapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{5}
}

func (m *CapturedSnapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CapturedSnapshot.Unmarshal(m, b)
}
func (m *CapturedSnapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CapturedSnapshot.Marshal(b, m, deterministic)
}
func (m *CapturedSnapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CapturedSnapshot.Merge(m, src)
}
func (m *CapturedSnapshot) XXX_Size() int {
	return xxx_messageInfo_CapturedSnapshot.Size(m)
}
func (m *CapturedSnapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_CapturedSnapshot.DiscardUnknown(m)
}

var xxx_messageInfo_CapturedSnapshot proto.InternalMessageInfo

func (m *CapturedSnapshot) GetMetadata() *history.Snapshot {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CapturedSnapshot) GetEntityIds() []int64 {
	if m != nil {
		return m.EntityIds
	}
	return nil
}

type ExpectationResult struct {
	ExpectationName                        string   `protobuf:"bytes,2,opt,name=expectation_name,json=expectationName,proto3" json:"expectation_name,omitempty"`
	Satisfied                              bool     `protobuf:"varint,1,opt,name=satisfied,proto3" json:"satisfied,omitempty"`
	FailureMessage                         string   `protobuf:"bytes,3,opt,name=failure_message,json=failureMessage,proto3" json:"failure_message,omitempty"`
	FailureIsRecoverableInScenarioLifetime bool     `protobuf:"varint,4,opt,name=failure_is_recoverable_in_scenario_lifetime,json=failureIsRecoverableInScenarioLifetime,proto3" json:"failure_is_recoverable_in_scenario_lifetime,omitempty"`
	XXX_NoUnkeyedLiteral                   struct{} `json:"-"`
	XXX_unrecognized                       []byte   `json:"-"`
	XXX_sizecache                          int32    `json:"-"`
}

func (m *ExpectationResult) Reset()         { *m = ExpectationResult{} }
func (m *ExpectationResult) String() string { return proto.CompactTextString(m) }
func (*ExpectationResult) ProtoMessage()    {}
func (*ExpectationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c8f0e4731b51e26, []int{6}
}

func (m *ExpectationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExpectationResult.Unmarshal(m, b)
}
func (m *ExpectationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExpectationResult.Marshal(b, m, deterministic)
}
func (m *ExpectationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExpectationResult.Merge(m, src)
}
func (m *ExpectationResult) XXX_Size() int {
	return xxx_messageInfo_ExpectationResult.Size(m)
}
func (m *ExpectationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExpectationResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExpectationResult proto.InternalMessageInfo

func (m *ExpectationResult) GetExpectationName() string {
	if m != nil {
		return m.ExpectationName
	}
	return ""
}

func (m *ExpectationResult) GetSatisfied() bool {
	if m != nil {
		return m.Satisfied
	}
	return false
}

func (m *ExpectationResult) GetFailureMessage() string {
	if m != nil {
		return m.FailureMessage
	}
	return ""
}

func (m *ExpectationResult) GetFailureIsRecoverableInScenarioLifetime() bool {
	if m != nil {
		return m.FailureIsRecoverableInScenarioLifetime
	}
	return false
}

func init() {
	proto.RegisterType((*ScenarioResult)(nil), "improbable.regisseur.ScenarioResult")
	proto.RegisterType((*CapturedMetric)(nil), "improbable.regisseur.CapturedMetric")
	proto.RegisterType((*PrometheusResult)(nil), "improbable.regisseur.PrometheusResult")
	proto.RegisterType((*LabelPair)(nil), "improbable.regisseur.LabelPair")
	proto.RegisterType((*SamplePair)(nil), "improbable.regisseur.SamplePair")
	proto.RegisterType((*CapturedSnapshot)(nil), "improbable.regisseur.CapturedSnapshot")
	proto.RegisterType((*ExpectationResult)(nil), "improbable.regisseur.ExpectationResult")
}

func init() {
	proto.RegisterFile("improbable/fabric/regisseur/scenario_result.proto", fileDescriptor_5c8f0e4731b51e26)
}

var fileDescriptor_5c8f0e4731b51e26 = []byte{
	// 615 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xed, 0x6a, 0xd4, 0x4c,
	0x18, 0x25, 0xdd, 0xbe, 0x6f, 0x77, 0x9f, 0x42, 0x3f, 0xc6, 0x0a, 0x4b, 0x55, 0x5a, 0x16, 0x6d,
	0xab, 0x42, 0x16, 0x5b, 0x04, 0xf1, 0x87, 0x28, 0x52, 0x71, 0xa1, 0x95, 0x32, 0x55, 0x10, 0x14,
	0xc2, 0x64, 0xf3, 0x64, 0x3b, 0x98, 0x64, 0x86, 0x99, 0x49, 0x69, 0x2f, 0xca, 0x9b, 0xf2, 0x16,
	0xf4, 0x02, 0x24, 0x93, 0x99, 0x64, 0x77, 0x8d, 0xe0, 0xbf, 0xe4, 0xcc, 0x39, 0x67, 0x9e, 0xcf,
	0x81, 0x67, 0x3c, 0x97, 0x4a, 0xc4, 0x2c, 0xce, 0x70, 0x9c, 0xb2, 0x58, 0xf1, 0xe9, 0x58, 0xe1,
	0x8c, 0x6b, 0x8d, 0xa5, 0x1a, 0xeb, 0x29, 0x16, 0x4c, 0x71, 0x11, 0x29, 0xd4, 0x65, 0x66, 0x42,
	0xa9, 0x84, 0x11, 0x64, 0xa7, 0x95, 0x84, 0x0d, 0x77, 0xf7, 0xc9, 0xbf, 0x18, 0xd5, 0x0e, 0x0b,
	0x5c, 0x99, 0x31, 0x93, 0x0a, 0x95, 0x8f, 0xaf, 0xb8, 0x36, 0x42, 0xdd, 0x8e, 0x75, 0xc1, 0xa4,
	0xbe, 0x12, 0xee, 0xb6, 0xdd, 0xbd, 0x99, 0x10, 0xb3, 0x8a, 0x57, 0xfd, 0xc5, 0x65, 0x3a, 0x36,
	0x3c, 0x47, 0x6d, 0x58, 0x2e, 0x6b, 0xc2, 0xe8, 0x57, 0x00, 0x1b, 0x97, 0xce, 0x9f, 0xda, 0x38,
	0xc9, 0x7b, 0x00, 0x9e, 0x60, 0x61, 0x78, 0xca, 0x51, 0x0d, 0x83, 0xfd, 0xe0, 0x68, 0xfd, 0xf8,
	0x28, 0xec, 0x0a, 0x3b, 0xf4, 0xca, 0x49, 0xc3, 0xa7, 0x73, 0x5a, 0xf2, 0x0a, 0xd6, 0x72, 0x34,
	0x8a, 0x4f, 0xf5, 0x70, 0x65, 0xbf, 0x77, 0xb4, 0x7e, 0xfc, 0xb0, 0xdb, 0xe6, 0x2d, 0x93, 0xa6,
	0x54, 0x98, 0x9c, 0x5b, 0x32, 0xf5, 0x22, 0xf2, 0x19, 0xee, 0xe0, 0x8d, 0xc4, 0xa9, 0x61, 0x86,
	0x8b, 0xc2, 0xd5, 0x51, 0x0f, 0x7b, 0xd6, 0xeb, 0xb0, 0xdb, 0xeb, 0xb4, 0x15, 0xd4, 0xf9, 0x50,
	0x82, 0xcb, 0x90, 0x1e, 0xfd, 0x0c, 0x60, 0x63, 0xf1, 0x56, 0x42, 0x60, 0xb5, 0x60, 0x39, 0xda,
	0x84, 0x07, 0xd4, 0x7e, 0x93, 0x13, 0xb8, 0x2b, 0x95, 0xc8, 0xd1, 0x5c, 0x61, 0xa9, 0x23, 0xbc,
	0x91, 0x0a, 0xb5, 0xe6, 0xa2, 0x18, 0xae, 0x58, 0xd2, 0x4e, 0x7b, 0x78, 0xda, 0x9c, 0x91, 0x4f,
	0x40, 0xe6, 0x44, 0x8b, 0x41, 0x1f, 0x74, 0x07, 0x7d, 0xd1, 0xf0, 0x5d, 0xcc, 0xdb, 0x72, 0x09,
	0xd1, 0xe4, 0x0d, 0xf4, 0x53, 0x96, 0x65, 0x31, 0x9b, 0x7e, 0x1b, 0xae, 0xda, 0xa6, 0x3c, 0xea,
	0x36, 0xab, 0xf3, 0x09, 0xdf, 0x39, 0x32, 0x6d, 0x64, 0xa3, 0xef, 0x01, 0x6c, 0x2d, 0x5f, 0x45,
	0xee, 0xc1, 0x20, 0x63, 0x31, 0x66, 0x91, 0x46, 0xe3, 0x92, 0xef, 0x5b, 0xe0, 0x12, 0x0d, 0x79,
	0x09, 0x6b, 0x9a, 0xe5, 0x32, 0x43, 0xdf, 0xc1, 0xfd, 0xbf, 0x0c, 0x82, 0x25, 0x5d, 0x30, 0xae,
	0xa8, 0x17, 0x90, 0xd7, 0xb0, 0x5e, 0x1b, 0x4b, 0xc6, 0x95, 0x2f, 0xc0, 0x5e, 0xb7, 0xfe, 0xac,
	0x22, 0x5a, 0x39, 0x64, 0xfe, 0x53, 0x8f, 0x9e, 0xc3, 0xa0, 0x39, 0xe8, 0xec, 0xcf, 0x0e, 0xfc,
	0x77, 0xcd, 0xb2, 0x12, 0x5d, 0x3f, 0xea, 0x9f, 0xd1, 0x57, 0x80, 0x36, 0x1e, 0xf2, 0x02, 0x06,
	0xcd, 0xd0, 0xbb, 0x69, 0xde, 0x0d, 0xeb, 0xb5, 0x08, 0xfd, 0x5a, 0x84, 0x1f, 0x3d, 0x83, 0xb6,
	0xe4, 0x45, 0xf7, 0xc0, 0xbb, 0x1b, 0xd8, 0xf2, 0x93, 0x73, 0xe9, 0x96, 0xad, 0xea, 0x4d, 0x8e,
	0x86, 0x25, 0xcc, 0x30, 0x77, 0xc5, 0x42, 0x6f, 0xfc, 0x96, 0x86, 0x6e, 0x4b, 0x43, 0x2f, 0xa4,
	0x8d, 0x8c, 0x3c, 0x00, 0xa8, 0xf6, 0xc6, 0xdc, 0x46, 0x3c, 0xa9, 0x8b, 0xdd, 0xa3, 0x83, 0x1a,
	0x99, 0x24, 0x7a, 0xf4, 0x23, 0x80, 0xed, 0x3f, 0x46, 0x9b, 0x3c, 0x86, 0xad, 0xf9, 0x05, 0xb1,
	0xf5, 0xa9, 0x4b, 0xb1, 0x39, 0x87, 0x7f, 0xa8, 0x4a, 0x75, 0x1f, 0x06, 0x9a, 0x19, 0xae, 0x53,
	0x8e, 0x89, 0x8d, 0xb1, 0x4f, 0x5b, 0x80, 0x1c, 0xc2, 0x66, 0xca, 0x78, 0x56, 0x2a, 0x8c, 0x72,
	0xd4, 0x9a, 0xcd, 0x70, 0xd8, 0xb3, 0x3e, 0x1b, 0x0e, 0x3e, 0xaf, 0x51, 0xf2, 0x05, 0x9e, 0x7a,
	0x22, 0xaf, 0x86, 0x7b, 0x2a, 0xae, 0x51, 0x55, 0x49, 0x46, 0xbc, 0x88, 0x9a, 0x17, 0x2f, 0xe3,
	0x29, 0x56, 0x75, 0xb4, 0x83, 0xda, 0xa7, 0x07, 0x4e, 0x32, 0xd1, 0xb4, 0x15, 0x4c, 0x0a, 0xff,
	0x7a, 0x9c, 0x39, 0x76, 0xfc, 0xbf, 0xed, 0xc7, 0xc9, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x7f, 0x8d, 0x34, 0x57, 0x05, 0x00, 0x00,
}