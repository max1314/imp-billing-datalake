// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/ext/client_launch/client/client.proto

package improbable_ext_client_launch_client

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/ext/client_launch/artifact"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ClientId) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BundleId)
	hasInner = hasInner || len(bundleIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleIdFields {
		res[k] = v
	}
	return res
}

func (this *ClientId) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BundleId, dst)
}

func (this *Client) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ClientId)
	hasInner = hasInner || len(clientIdFields) > 0
	metadataFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Metadata)
	hasInner = hasInner || len(metadataFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientIdFields {
		res[k] = v
	}
	for k, v := range metadataFields {
		res[k] = v
	}
	return res
}

func (this *Client) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ClientId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Metadata, dst)
}

func (this *CreateClientRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Client)
	hasInner = hasInner || len(clientFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientFields {
		res[k] = v
	}
	return res
}

func (this *CreateClientRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Client, dst)
}

func (this *CreateClientResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Client)
	hasInner = hasInner || len(clientFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientFields {
		res[k] = v
	}
	return res
}

func (this *CreateClientResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Client, dst)
}

func (this *GetClientRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ClientId)
	hasInner = hasInner || len(clientIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientIdFields {
		res[k] = v
	}
	return res
}

func (this *GetClientRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ClientId, dst)
}

func (this *GetClientResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Client)
	hasInner = hasInner || len(clientFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientFields {
		res[k] = v
	}
	return res
}

func (this *GetClientResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Client, dst)
}

func (this *DownloadClientRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ClientId)
	hasInner = hasInner || len(clientIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientIdFields {
		res[k] = v
	}
	return res
}

func (this *DownloadClientRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ClientId, dst)
}

func (this *DownloadClientResponse) LogFields() map[string]string {
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

func (this *DownloadClientResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Metadata, dst)
}

func (this *ListClientsRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BundleId)
	hasInner = hasInner || len(bundleIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleIdFields {
		res[k] = v
	}
	return res
}

func (this *ListClientsRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BundleId, dst)
}

func (this *ListClientsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListClientsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeleteClientRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	clientIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ClientId)
	hasInner = hasInner || len(clientIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range clientIdFields {
		res[k] = v
	}
	return res
}

func (this *DeleteClientRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ClientId, dst)
}

func (this *DeleteClientResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteClientResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *BundleId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *BundleId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Bundle) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BundleId)
	hasInner = hasInner || len(bundleIdFields) > 0
	creationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreationTime)
	hasInner = hasInner || len(creationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleIdFields {
		res[k] = v
	}
	for k, v := range creationTimeFields {
		res[k] = v
	}
	return res
}

func (this *Bundle) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BundleId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreationTime, dst)
}

func (this *CreateBundleRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Bundle)
	hasInner = hasInner || len(bundleFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleFields {
		res[k] = v
	}
	return res
}

func (this *CreateBundleRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Bundle, dst)
}

func (this *CreateBundleResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Bundle)
	hasInner = hasInner || len(bundleFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleFields {
		res[k] = v
	}
	return res
}

func (this *CreateBundleResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Bundle, dst)
}

func (this *GetBundleRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BundleId)
	hasInner = hasInner || len(bundleIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleIdFields {
		res[k] = v
	}
	return res
}

func (this *GetBundleRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BundleId, dst)
}

func (this *GetBundleResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Bundle)
	hasInner = hasInner || len(bundleFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleFields {
		res[k] = v
	}
	return res
}

func (this *GetBundleResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Bundle, dst)
}

func (this *ListBundlesRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListBundlesRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListBundlesResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListBundlesResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeleteBundleRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	bundleIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.BundleId)
	hasInner = hasInner || len(bundleIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range bundleIdFields {
		res[k] = v
	}
	return res
}

func (this *DeleteBundleRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.BundleId, dst)
}

func (this *DeleteBundleResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteBundleResponse) ExtractRequestFields(dst map[string]interface{}) {
}
