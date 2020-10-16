// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/resolver/lookup_proxy.proto

package improbable_infra

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

func (this *LookupSrvRequest) Validate() error {
	return nil
}

var _regex_SrvRecord_Ip = regexp.MustCompile(`([0-9]+[.]){3}[0-9]+`)

func (this *SrvRecord) Validate() error {
	if !_regex_SrvRecord_Ip.MatchString(this.Ip) {
		return github_com_mwitkow_go_proto_validators.FieldError("Ip", fmt.Errorf(`value '%v' must be a string conforming to regex "([0-9]+[.]){3}[0-9]+"`, this.Ip))
	}
	if !(this.Port < 32768) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be less than '32768'`, this.Port))
	}
	return nil
}
func (this *LookupSrvResponse) Validate() error {
	for _, item := range this.Records {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Records", err)
			}
		}
	}
	return nil
}
