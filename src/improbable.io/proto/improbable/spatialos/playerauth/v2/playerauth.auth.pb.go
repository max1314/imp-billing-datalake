// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/spatialos/playerauth/v2/playerauth.proto

package pb_playerauthv2

import (
	context "context"
	errors "errors"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/CreateLoginToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/CreatePlayerIdentityToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/DecodePlayerIdentityToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/CreateDevelopmentAuthenticationToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/GetDevelopmentAuthenticationToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/ListDevelopmentAuthenticationTokens",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/UpdateDevelopmentAuthenticationToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/ExpireDevelopmentAuthenticationToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.spatialos.playerauth.v2.PlayerAuthService/DeleteDevelopmentAuthenticationToken",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicCreateLT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicCreateLT".
func Permits_V1_PlayerAuthV2PublicCreateLT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicCreateLT_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicCreatePIT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicCreatePIT".
func Permits_V1_PlayerAuthV2PublicCreatePIT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicCreatePIT_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicDecodePIT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicDecodePIT".
func Permits_V1_PlayerAuthV2PublicDecodePIT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicDecodePIT_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicCreateDAT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicCreateDAT".
func Permits_V1_PlayerAuthV2PublicCreateDAT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicCreateDAT_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicGetDAT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicGetDAT".
func Permits_V1_PlayerAuthV2PublicGetDAT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicGetDAT_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicListDATs_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicListDATs".
func Permits_V1_PlayerAuthV2PublicListDATs(ctx context.Context, args Permits_V1_PlayerAuthV2PublicListDATs_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicUpdateDAT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicUpdateDAT".
func Permits_V1_PlayerAuthV2PublicUpdateDAT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicUpdateDAT_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicExpireDAT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicExpireDAT".
func Permits_V1_PlayerAuthV2PublicExpireDAT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicExpireDAT_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}

// The arguments for this v1 permission check.
type Permits_V1_PlayerAuthV2PublicDeleteDAT_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "PlayerAuthV2PublicDeleteDAT".
func Permits_V1_PlayerAuthV2PublicDeleteDAT(ctx context.Context, args Permits_V1_PlayerAuthV2PublicDeleteDAT_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/playerauth", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}
