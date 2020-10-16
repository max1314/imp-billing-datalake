// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/playtest/playtest_share.proto

package improbable_platform_playtest

import (
	context "context"
	errors "errors"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
	improbable_io_lib_oauth2_authctx "improbable.io/lib/oauth2/authctx"
	improbable_io_lib_oauth2_perms "improbable.io/lib/oauth2/perms"
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
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.playtest.PlaytestShareService/CreatePlaytestShareToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.playtest.PlaytestShareService/DecodePlaytestShareToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}

// The arguments for this v1 permission check.
type Permits_V1_PlaytestHubCreatePlaytestShareToken_Args struct {
	Deployment string
	Project    string
}

// Performs the delegated permission check described by delegation "PlaytestHubCreatePlaytestShareToken".
func Permits_V1_PlaytestHubCreatePlaytestShareToken(ctx context.Context, args Permits_V1_PlaytestHubCreatePlaytestShareToken_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/dpl/%v/playtest", args.Project, args.Deployment)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}
