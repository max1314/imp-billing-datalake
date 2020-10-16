// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: infra/accounts/email_verification.proto

package improbable_accounts

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	improbable_io_lib_perms_registry "improbable.io/lib/perms/registry"
	improbable_io_proto_improbable_ext_plugin_auth "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func init() {
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.accounts.EmailVerificationService/SendVerificationEmail",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.accounts.EmailVerificationService/VerifyEmail",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}
