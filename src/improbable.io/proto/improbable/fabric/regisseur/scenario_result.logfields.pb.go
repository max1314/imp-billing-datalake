// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/fabric/regisseur/scenario_result.proto

package improbable_regisseur

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "improbable.io/proto/improbable/platform/history"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ScenarioResult) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	identifierFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Identifier)
	hasInner = hasInner || len(identifierFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range identifierFields {
		res[k] = v
	}
	return res
}

func (this *ScenarioResult) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Identifier, dst)
}

func (this *CapturedMetric) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	fallbackFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Fallback)
	hasInner = hasInner || len(fallbackFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range fallbackFields {
		res[k] = v
	}
	return res
}

func (this *CapturedMetric) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Fallback, dst)
}

func (this *PrometheusResult) LogFields() map[string]string {
	return map[string]string{}
}

func (this *PrometheusResult) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *LabelPair) LogFields() map[string]string {
	return map[string]string{}
}

func (this *LabelPair) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *SamplePair) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	timestampFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Timestamp)
	hasInner = hasInner || len(timestampFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range timestampFields {
		res[k] = v
	}
	return res
}

func (this *SamplePair) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Timestamp, dst)
}

func (this *CapturedSnapshot) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	metadataFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Metadata)
	hasInner = hasInner || len(metadataFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range metadataFields {
		res[k] = v
	}
	return res
}

func (this *CapturedSnapshot) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Metadata, dst)
}

func (this *ExpectationResult) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ExpectationResult) ExtractRequestFields(dst map[string]interface{}) {
}