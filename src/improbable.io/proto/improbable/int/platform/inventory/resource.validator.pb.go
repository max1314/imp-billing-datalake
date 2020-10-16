// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/int/platform/inventory/resource.proto

package improbable_int_platform_inventory

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
	regexp "regexp"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DateRange) Validate() error {
	if this.From != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.From); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("From", err)
		}
	}
	if this.To != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.To); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("To", err)
		}
	}
	return nil
}
func (this *Resource) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *ResourceEvent) Validate() error {
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *ListResourceRequest) Validate() error {
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	if this.DeletedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.DeletedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("DeletedAt", err)
		}
	}
	return nil
}
func (this *ListResourceResponse) Validate() error {
	for _, item := range this.Resources {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Resources", err)
			}
		}
	}
	return nil
}

var _regex_GetResourceRequest_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *GetResourceRequest) Validate() error {
	if !_regex_GetResourceRequest_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`Please provide a valid resource ID.`))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`Please provide a valid resource ID.`))
	}
	return nil
}
func (this *GetResourceResponse) Validate() error {
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *CreateResourceRequest) Validate() error {
	if nil == this.Resource {
		return github_com_mwitkow_go_proto_validators.FieldError("Resource", fmt.Errorf("message must exist"))
	}
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *CreateResourceResponse) Validate() error {
	return nil
}

var _regex_DeleteResourceRequest_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)

func (this *DeleteResourceRequest) Validate() error {
	if !_regex_DeleteResourceRequest_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`Please provide a valid resource ID.`))
	}
	if this.Id == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`Please provide a valid resource ID.`))
	}
	return nil
}
func (this *DeleteResourceResponse) Validate() error {
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *UpdateResourceRequest) Validate() error {
	if nil == this.Resource {
		return github_com_mwitkow_go_proto_validators.FieldError("Resource", fmt.Errorf("message must exist"))
	}
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *UpdateResourceResponse) Validate() error {
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *ListResourceEventsRequest) Validate() error {
	if this.StartTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.StartTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("StartTimestamp", err)
		}
	}
	if this.EndTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EndTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EndTimestamp", err)
		}
	}
	if this.Resource != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Resource); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Resource", err)
		}
	}
	return nil
}
func (this *ListResourceEventsResponse) Validate() error {
	for _, item := range this.ResourceEvents {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("ResourceEvents", err)
			}
		}
	}
	return nil
}
func (this *ListAttributesRequest) Validate() error {
	return nil
}
func (this *Attribute) Validate() error {
	return nil
}
func (this *ListAttributesResponse) Validate() error {
	for _, item := range this.Attributes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Attributes", err)
			}
		}
	}
	return nil
}
