// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/auth/testing/test_accounts.proto

package improbable_int_auth_testing

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "improbable.io/proto/improbable/platform"
	_ "improbable.io/proto/infra/accounts"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *TestAccountId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *TestAccountId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Account) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	testAccountIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.TestAccountId)
	hasInner = hasInner || len(testAccountIdFields) > 0
	accountIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountId)
	hasInner = hasInner || len(accountIdFields) > 0
	expirationTimeFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ExpirationTime)
	hasInner = hasInner || len(expirationTimeFields) > 0
	creatorIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.CreatorId)
	hasInner = hasInner || len(creatorIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range testAccountIdFields {
		res[k] = v
	}
	for k, v := range accountIdFields {
		res[k] = v
	}
	for k, v := range expirationTimeFields {
		res[k] = v
	}
	for k, v := range creatorIdFields {
		res[k] = v
	}
	return res
}

func (this *Account) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.TestAccountId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ExpirationTime, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.CreatorId, dst)
}

func (this *CreateRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CreateResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	testAccountIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.TestAccountId)
	hasInner = hasInner || len(testAccountIdFields) > 0
	accountIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountId)
	hasInner = hasInner || len(accountIdFields) > 0
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	projectNameFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.ProjectName)
	hasInner = hasInner || len(projectNameFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range testAccountIdFields {
		res[k] = v
	}
	for k, v := range accountIdFields {
		res[k] = v
	}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	for k, v := range projectNameFields {
		res[k] = v
	}
	return res
}

func (this *CreateResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.TestAccountId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.ProjectName, dst)
}

func (this *ListRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *DeleteRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	testAccountIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.TestAccountId)
	hasInner = hasInner || len(testAccountIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range testAccountIdFields {
		res[k] = v
	}
	return res
}

func (this *DeleteRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.TestAccountId, dst)
}

func (this *DeleteResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteResponse) ExtractRequestFields(dst map[string]interface{}) {
}
