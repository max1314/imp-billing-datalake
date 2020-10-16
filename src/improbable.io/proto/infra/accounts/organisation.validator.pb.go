// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/accounts/organisation.proto

package improbable_accounts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Organisation) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *CreateOrganisationRequest) Validate() error {
	if this.Organisation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Organisation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Organisation", err)
		}
	}
	return nil
}
func (this *CreateOrganisationResponse) Validate() error {
	if this.Organisation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Organisation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Organisation", err)
		}
	}
	return nil
}
func (this *AddAccountToOrganisationRequest) Validate() error {
	if this.AccountId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.AccountId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("AccountId", err)
		}
	}
	if this.OrganisationId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OrganisationId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OrganisationId", err)
		}
	}
	return nil
}
func (this *AddAccountToOrganisationResponse) Validate() error {
	return nil
}
func (this *GetOrganisationRequest) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	return nil
}
func (this *GetOrganisationResponse) Validate() error {
	if this.Id != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Id); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Id", err)
		}
	}
	if this.Organisation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Organisation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Organisation", err)
		}
	}
	return nil
}
func (this *ListOrganisationsRequest) Validate() error {
	return nil
}
func (this *ListOrganisationsResponse) Validate() error {
	for _, item := range this.Organisations {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Organisations", err)
			}
		}
	}
	return nil
}
func (this *RemoveUserFromOrganisationRequest) Validate() error {
	if this.OrganisationId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OrganisationId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OrganisationId", err)
		}
	}
	return nil
}
func (this *RemoveUserFromOrganisationResponse) Validate() error {
	return nil
}
func (this *SetOrganisationDefaultGroupRequest) Validate() error {
	if this.OrganisationId != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.OrganisationId); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("OrganisationId", err)
		}
	}
	return nil
}
func (this *SetOrganisationDefaultGroupResponse) Validate() error {
	return nil
}
