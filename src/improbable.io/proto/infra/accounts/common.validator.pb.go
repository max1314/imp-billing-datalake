// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/accounts/common.proto

package improbable_accounts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *AccountId) Validate() error {
	return nil
}

var _regex_AccountEmail_Email = regexp.MustCompile(`^[^[:cntrl:] ,:;<>@"'()\[\]{}]{0,63}@[[:alnum:].-]{2,63}$`)

func (this *AccountEmail) Validate() error {
	if !_regex_AccountEmail_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex "^[^[:cntrl:] ,:;<>@\"'()\\[\\]{}]{0,63}@[[:alnum:].-]{2,63}$"`, this.Email))
	}
	return nil
}
func (this *ServiceAccountId) Validate() error {
	return nil
}
func (this *OrganisationId) Validate() error {
	return nil
}
