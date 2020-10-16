// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: improbable/platform/locator/deployments.proto

package improbable_platform_locator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	improbable_io_lib_perms_registry "improbable.io/lib/perms/registry"
	_ "improbable.io/proto/improbable/ext/plugin/auth"
	improbable_io_proto_improbable_ext_plugin_auth "improbable.io/proto/improbable/ext/plugin/auth"
	_ "improbable.io/proto/improbable/platform"
	_ "improbable.io/proto/improbable/platform/deployment"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func init() {
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.locator.LocateService/ListDeploymentsForProject",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.locator.LocateService/ListDeployments",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.locator.LocateService/QueueForDeployment",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.locator.LocateService/LoginToDeployment",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
	improbable_io_lib_perms_registry.RegisterPermissions("/improbable.platform.locator.LocateServiceV2/Login",
		[]improbable_io_proto_improbable_ext_plugin_auth.Permission{},
	)
}
