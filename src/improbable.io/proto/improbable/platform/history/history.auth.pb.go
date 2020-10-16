// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/history/history.proto

package improbable_platform_history

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
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.history.HistoryService/Create",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.history.HistoryService/Get",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.history.HistoryService/Delete",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.history.HistoryService/List",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}

// The arguments for this v1 permission check.
type Permits_V1_SnapshotHistoryPrivateCreate_Args struct {
	History string
	Project string
}

// Performs the delegated permission check described by delegation "SnapshotHistoryPrivateCreate".
func Permits_V1_SnapshotHistoryPrivateCreate(ctx context.Context, args Permits_V1_SnapshotHistoryPrivateCreate_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/hist/%v", args.Project, args.History)
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
type Permits_V1_SnapshotHistoryPrivateGet_Args struct {
	History string
	Project string
}

// Performs the delegated permission check described by delegation "SnapshotHistoryPrivateGet".
func Permits_V1_SnapshotHistoryPrivateGet(ctx context.Context, args Permits_V1_SnapshotHistoryPrivateGet_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/hist/%v", args.Project, args.History)
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
type Permits_V1_SnapshotHistoryPrivateDelete_Args struct {
	History string
	Project string
}

// Performs the delegated permission check described by delegation "SnapshotHistoryPrivateDelete".
func Permits_V1_SnapshotHistoryPrivateDelete(ctx context.Context, args Permits_V1_SnapshotHistoryPrivateDelete_Args) error {
	perm := fmt.Sprintf("[w]:prj/%v/hist/%v", args.Project, args.History)
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
type Permits_V1_SnapshotHistoryPrivateList_Args struct {
	Project string
}

// Performs the delegated permission check described by delegation "SnapshotHistoryPrivateList".
func Permits_V1_SnapshotHistoryPrivateList(ctx context.Context, args Permits_V1_SnapshotHistoryPrivateList_Args) error {
	perm := fmt.Sprintf("[r]:prj/%v/hist/*", args.Project)
	grant, err := improbable_io_lib_oauth2_perms.ParseGrant(perm)
	if err != nil {
		return err
	}
	if !improbable_io_lib_oauth2_authctx.Perms(ctx).Permits(grant) {
		return errors.New(fmt.Sprintf("perm %v denied for grants %v", perm, improbable_io_lib_oauth2_authctx.Perms(ctx)))
	}
	return nil
}