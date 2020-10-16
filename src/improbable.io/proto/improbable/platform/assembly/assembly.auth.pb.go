// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/assembly/assembly.proto

package improbable_platform_assembly

import (
	context "context"
	errors "errors"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/improbable-io/go-proto-logfields"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	improbable_io_lib_oauth2_authctx "improbable.io/lib/oauth2/authctx"
	improbable_io_lib_oauth2_perms "improbable.io/lib/oauth2/perms"
	improbable_io_lib_perms_registry "improbable.io/lib/perms/registry"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	improbable_io_proto_improbable_ext_plugin_auth "improbable.io/proto/improbable/ext/plugin/auth"
	_ "improbable.io/proto/improbable/platform"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func init() {
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.assembly.AssemblyService/Create",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.assembly.AssemblyService/Update",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.assembly.AssemblyService/Get",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.assembly.AssemblyService/Delete",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.assembly.AssemblyService/List",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}

// The arguments for this v1 permission check.
type Permits_V1_AssemblyCreate_Args struct {
	Assembly string
	Project  string
}

// Performs the delegated permission check described by delegation "AssemblyCreate".
func Permits_V1_AssemblyCreate(ctx context.Context, args Permits_V1_AssemblyCreate_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/asm/%v", args.Project, args.Assembly)
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
type Permits_V1_AssemblyUpdate_Args struct {
	Assembly string
	Project  string
}

// Performs the delegated permission check described by delegation "AssemblyUpdate".
func Permits_V1_AssemblyUpdate(ctx context.Context, args Permits_V1_AssemblyUpdate_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/asm/%v", args.Project, args.Assembly)
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
type Permits_V1_AssemblyGet_Args struct {
	Assembly string
	Project  string
}

// Performs the delegated permission check described by delegation "AssemblyGet".
func Permits_V1_AssemblyGet(ctx context.Context, args Permits_V1_AssemblyGet_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/asm/%v", args.Project, args.Assembly)
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
type Permits_V1_AssemblyDelete_Args struct {
	Assembly string
	Project  string
}

// Performs the delegated permission check described by delegation "AssemblyDelete".
func Permits_V1_AssemblyDelete(ctx context.Context, args Permits_V1_AssemblyDelete_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/asm/%v", args.Project, args.Assembly)
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
type Permits_V1_AssemblyList_Args struct {
	Assembly string
	Project  string
}

// Performs the delegated permission check described by delegation "AssemblyList".
func Permits_V1_AssemblyList(ctx context.Context, args Permits_V1_AssemblyList_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/asm/%v", args.Project, args.Assembly)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}