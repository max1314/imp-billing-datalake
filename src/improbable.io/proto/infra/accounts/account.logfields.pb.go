// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/accounts/account.proto

package improbable_accounts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	github_com_improbable_io_go_proto_logfields "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GroupId) LogFields() map[string]string {
	return map[string]string{}
}

func (this *GroupId) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Grant) LogFields() map[string]string {
	return map[string]string{}
}

func (this *Grant) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GroupPermPair) LogFields() map[string]string {
	return map[string]string{}
}

func (this *GroupPermPair) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *AccountPermPair) LogFields() map[string]string {
	return map[string]string{}
}

func (this *AccountPermPair) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CreateAccountRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	return res
}

func (this *CreateAccountRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
}

func (this *CreateBetaAccountRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *CreateBetaAccountRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *CreateBetaAccountResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	organisationIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.OrganisationId)
	hasInner = hasInner || len(organisationIdFields) > 0
	accountIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountId)
	hasInner = hasInner || len(accountIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range organisationIdFields {
		res[k] = v
	}
	for k, v := range accountIdFields {
		res[k] = v
	}
	return res
}

func (this *CreateBetaAccountResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.OrganisationId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountId, dst)
}

func (this *CreateGroupRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	return res
}

func (this *CreateGroupRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
}

func (this *RemoveGroupRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	return res
}

func (this *RemoveGroupRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
}

func (this *GroupModifyAccountRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	return res
}

func (this *GroupModifyAccountRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
}

func (this *SetGroupPermissionRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	return res
}

func (this *SetGroupPermissionRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
}

func (this *ListAccountsRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAccountsRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAccountsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAccountsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAccountsByGroupRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAccountsByGroupRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAccountsByGroupResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAccountsByGroupResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *Account) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	emailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Email)
	hasInner = hasInner || len(emailFields) > 0
	organisationIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.OrganisationId)
	hasInner = hasInner || len(organisationIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range emailFields {
		res[k] = v
	}
	for k, v := range organisationIdFields {
		res[k] = v
	}
	return res
}

func (this *Account) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Email, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.OrganisationId, dst)
}

func (this *ListAllAccountsRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	organisationIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.OrganisationId)
	hasInner = hasInner || len(organisationIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range organisationIdFields {
		res[k] = v
	}
	return res
}

func (this *ListAllAccountsRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.OrganisationId, dst)
}

func (this *ListAllAccountsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAllAccountsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GroupMembers) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	return res
}

func (this *GroupMembers) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
}

func (this *AccountMembership) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	return res
}

func (this *AccountMembership) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
}

func (this *DeleteAccountRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *DeleteAccountRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListMembershipRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListMembershipRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListMembershipResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListMembershipResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAccountPermissionRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	return res
}

func (this *ListAccountPermissionRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
}

func (this *ListAccountPermissionResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	return res
}

func (this *ListAccountPermissionResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
}

func (this *ListGroupsRequest) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListGroupsRequest) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListGroupsResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListGroupsResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListGroupPermissionRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	return res
}

func (this *ListGroupPermissionRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
}

func (this *ListGroupPermissionResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	groupIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.GroupId)
	hasInner = hasInner || len(groupIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range groupIdFields {
		res[k] = v
	}
	return res
}

func (this *ListGroupPermissionResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.GroupId, dst)
}

func (this *SetBlacklistStatusRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	accountEmailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.AccountEmail)
	hasInner = hasInner || len(accountEmailFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range accountEmailFields {
		res[k] = v
	}
	return res
}

func (this *SetBlacklistStatusRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.AccountEmail, dst)
}

func (this *SetBlacklistStatusResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *SetBlacklistStatusResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *GetAccountRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	return res
}

func (this *GetAccountRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
}

func (this *GetAccountResponse) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	idFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Id)
	hasInner = hasInner || len(idFields) > 0
	emailFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Email)
	hasInner = hasInner || len(emailFields) > 0
	organisationIdFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.OrganisationId)
	hasInner = hasInner || len(organisationIdFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range idFields {
		res[k] = v
	}
	for k, v := range emailFields {
		res[k] = v
	}
	for k, v := range organisationIdFields {
		res[k] = v
	}
	return res
}

func (this *GetAccountResponse) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Id, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Email, dst)
	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.OrganisationId, dst)
}

func (this *ListGroupsByPermissionRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	grantFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Grant)
	hasInner = hasInner || len(grantFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range grantFields {
		res[k] = v
	}
	return res
}

func (this *ListGroupsByPermissionRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Grant, dst)
}

func (this *ListGroupsByPermissionResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListGroupsByPermissionResponse) ExtractRequestFields(dst map[string]interface{}) {
}

func (this *ListAccountsByPermissionRequest) LogFields() map[string]string {
	// Handle being called on nil message.
	if this == nil {
		return map[string]string{}
	}
	// Gather fields from oneofs and child messages.
	var hasInner bool
	grantFields := github_com_improbable_io_go_proto_logfields.ExtractLogFieldsFromMessage(this.Grant)
	hasInner = hasInner || len(grantFields) > 0
	if !hasInner {
		// If no inner messages added any fields, avoid merging maps.
		return map[string]string{}
	}
	// Merge all the field maps.
	res := map[string]string{}
	for k, v := range grantFields {
		res[k] = v
	}
	return res
}

func (this *ListAccountsByPermissionRequest) ExtractRequestFields(dst map[string]interface{}) {
	// Handle being called on nil message.
	if this == nil {
		return
	}

	github_com_improbable_io_go_proto_logfields.ExtractRequestFieldsFromMessage(this.Grant, dst)
}

func (this *ListAccountsByPermissionResponse) LogFields() map[string]string {
	return map[string]string{}
}

func (this *ListAccountsByPermissionResponse) ExtractRequestFields(dst map[string]interface{}) {
}