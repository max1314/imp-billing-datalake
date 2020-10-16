// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/ext/client_launch/artifact/artifact.proto

package improbable_ext_client_launch_artifact

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ArtifactId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ArtifactId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ArtifactMetadata) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ArtifactMetadata) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Artifact) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ArtifactId)
	hasInner = hasInner || len(artifactIdFields) > 0
	metadataFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Metadata)
	hasInner = hasInner || len(metadataFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactIdFields {
		res[k] = v
	}
	for k, v := range metadataFields {
		res[k] = v
	}
	return res
}

func (this *Artifact) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ArtifactId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Metadata, dst)
}

func (this *GetArtifactRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ArtifactId)
	hasInner = hasInner || len(artifactIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactIdFields {
		res[k] = v
	}
	return res
}

func (this *GetArtifactRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ArtifactId, dst)
}

func (this *GetArtifactResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Artifact)
	hasInner = hasInner || len(artifactFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactFields {
		res[k] = v
	}
	return res
}

func (this *GetArtifactResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Artifact, dst)
}

func (this *CreateArtifactRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Artifact)
	hasInner = hasInner || len(artifactFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactFields {
		res[k] = v
	}
	return res
}

func (this *CreateArtifactRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Artifact, dst)
}

func (this *CreateArtifactResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Artifact)
	hasInner = hasInner || len(artifactFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactFields {
		res[k] = v
	}
	return res
}

func (this *CreateArtifactResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Artifact, dst)
}

func (this *DeleteArtifactRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ArtifactId)
	hasInner = hasInner || len(artifactIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactIdFields {
		res[k] = v
	}
	return res
}

func (this *DeleteArtifactRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ArtifactId, dst)
}

func (this *DeleteArtifactResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteArtifactResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetArtifactDownloadUrlRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ArtifactId)
	hasInner = hasInner || len(artifactIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactIdFields {
		res[k] = v
	}
	return res
}

func (this *GetArtifactDownloadUrlRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ArtifactId, dst)
}

func (this *GetArtifactDownloadUrlResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	artifactIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ArtifactId)
	hasInner = hasInner || len(artifactIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range artifactIdFields {
		res[k] = v
	}
	return res
}

func (this *GetArtifactDownloadUrlResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ArtifactId, dst)
}