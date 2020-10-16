// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/accounts/refresh_token.proto

package improbable_accounts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *RefreshToken) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	accountFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Account)
	hasInner = hasInner || len(accountFields) > 0
	serviceAccountFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ServiceAccount)
	hasInner = hasInner || len(serviceAccountFields) > 0
	issueTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.IssueTime)
	hasInner = hasInner || len(issueTimeFields) > 0
	expirationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpirationTime)
	hasInner = hasInner || len(expirationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range accountFields {
		res[k] = v
	}
	for k, v := range serviceAccountFields {
		res[k] = v
	}
	for k, v := range issueTimeFields {
		res[k] = v
	}
	for k, v := range expirationTimeFields {
		res[k] = v
	}
	return res
}

func (this *RefreshToken) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Account, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ServiceAccount, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.IssueTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpirationTime, dst)
}

func (this *RefreshTokenId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *RefreshTokenId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *FetchRefreshTokenRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *FetchRefreshTokenRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *FetchRefreshTokenResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	accountFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Account)
	hasInner = hasInner || len(accountFields) > 0
	serviceAccountFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ServiceAccount)
	hasInner = hasInner || len(serviceAccountFields) > 0
	issueTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.IssueTime)
	hasInner = hasInner || len(issueTimeFields) > 0
	expirationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpirationTime)
	hasInner = hasInner || len(expirationTimeFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range accountFields {
		res[k] = v
	}
	for k, v := range serviceAccountFields {
		res[k] = v
	}
	for k, v := range issueTimeFields {
		res[k] = v
	}
	for k, v := range expirationTimeFields {
		res[k] = v
	}
	return res
}

func (this *FetchRefreshTokenResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Account, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ServiceAccount, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.IssueTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpirationTime, dst)
}

func (this *RemoveExpiredRefreshTokensRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *RemoveExpiredRefreshTokensRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *RemoveExpiredRefreshTokensResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *RemoveExpiredRefreshTokensResponse) ExtractRequestFields(dst map[string]interface{}) {
}
