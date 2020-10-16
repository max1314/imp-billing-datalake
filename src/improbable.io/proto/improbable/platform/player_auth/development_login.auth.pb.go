// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/player_auth/development_login.proto

package improbable_platform_player_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/mwitkow/go-proto-validators"
	improbable_io_lib_perms_registry "improbable.io/lib/perms/registry"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	improbable_io_proto_improbable_ext_plugin_auth "improbable.io/proto/improbable/ext/plugin/auth"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func init() {
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.player_auth.DevelopmentLoginService/CreateLoginTokens",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}
