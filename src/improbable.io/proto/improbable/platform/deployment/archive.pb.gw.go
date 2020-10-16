// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: improbable/platform/deployment/archive.proto

/*
Package improbable_platform_deployment is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package improbable_platform_deployment

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/descriptor"
	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = descriptor.ForMessage

var (
	filter_DeploymentArchiveService_LegacyListAll_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_DeploymentArchiveService_LegacyListAll_0(ctx context.Context, marshaler runtime.Marshaler, client DeploymentArchiveServiceClient, req *http.Request, pathParams map[string]string) (DeploymentArchiveService_LegacyListAllClient, runtime.ServerMetadata, error) {
	var protoReq ArchiveListAllRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_DeploymentArchiveService_LegacyListAll_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.LegacyListAll(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

var (
	filter_DeploymentArchiveService_GetConfig_0 = &utilities.DoubleArray{Encoding: map[string]int{"deployment_id": 0, "project_id": 1, "name": 2, "uuid": 3}, Base: []int{1, 4, 1, 1, 4, 0, 3, 0, 0}, Check: []int{0, 1, 2, 3, 1, 4, 2, 7, 5}}
)

func request_DeploymentArchiveService_GetConfig_0(ctx context.Context, marshaler runtime.Marshaler, client DeploymentArchiveServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ArchiveConfigurationGetRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["deployment_id.project_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id.project_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "deployment_id.project_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id.project_id.name", err)
	}

	val, ok = pathParams["deployment_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "deployment_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id.name", err)
	}

	val, ok = pathParams["uuid"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "uuid")
	}

	protoReq.Uuid, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "uuid", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_DeploymentArchiveService_GetConfig_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.GetConfig(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_DeploymentArchiveService_GetConfig_0(ctx context.Context, marshaler runtime.Marshaler, server DeploymentArchiveServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ArchiveConfigurationGetRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["deployment_id.project_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id.project_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "deployment_id.project_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id.project_id.name", err)
	}

	val, ok = pathParams["deployment_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "deployment_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id.name", err)
	}

	val, ok = pathParams["uuid"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "uuid")
	}

	protoReq.Uuid, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "uuid", err)
	}

	if err := runtime.PopulateQueryParameters(&protoReq, req.URL.Query(), filter_DeploymentArchiveService_GetConfig_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.GetConfig(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_DeploymentArchiveService_ListProject_0 = &utilities.DoubleArray{Encoding: map[string]int{"project_id": 0, "name": 1}, Base: []int{1, 1, 1, 0}, Check: []int{0, 1, 2, 3}}
)

func request_DeploymentArchiveService_ListProject_0(ctx context.Context, marshaler runtime.Marshaler, client DeploymentArchiveServiceClient, req *http.Request, pathParams map[string]string) (DeploymentArchiveService_ListProjectClient, runtime.ServerMetadata, error) {
	var protoReq ArchiveListProjectRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["project_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "project_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "project_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "project_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_DeploymentArchiveService_ListProject_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.ListProject(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

var (
	filter_DeploymentArchiveService_ListDeployment_0 = &utilities.DoubleArray{Encoding: map[string]int{"deployment_id": 0, "project_id": 1, "name": 2}, Base: []int{1, 3, 1, 1, 0, 3, 0}, Check: []int{0, 1, 2, 3, 4, 2, 6}}
)

func request_DeploymentArchiveService_ListDeployment_0(ctx context.Context, marshaler runtime.Marshaler, client DeploymentArchiveServiceClient, req *http.Request, pathParams map[string]string) (DeploymentArchiveService_ListDeploymentClient, runtime.ServerMetadata, error) {
	var protoReq ArchiveListDeploymentRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["deployment_id.project_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id.project_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "deployment_id.project_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id.project_id.name", err)
	}

	val, ok = pathParams["deployment_id.name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "deployment_id.name")
	}

	err = runtime.PopulateFieldFromPath(&protoReq, "deployment_id.name", val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "deployment_id.name", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_DeploymentArchiveService_ListDeployment_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.ListDeployment(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterDeploymentArchiveServiceHandlerServer registers the http handlers for service DeploymentArchiveService to "mux".
// UnaryRPC     :call DeploymentArchiveServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterDeploymentArchiveServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server DeploymentArchiveServiceServer) error {

	mux.Handle("GET", pattern_DeploymentArchiveService_LegacyListAll_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_DeploymentArchiveService_GetConfig_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_DeploymentArchiveService_GetConfig_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DeploymentArchiveService_GetConfig_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_DeploymentArchiveService_ListProject_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_DeploymentArchiveService_ListDeployment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterDeploymentArchiveServiceHandlerFromEndpoint is same as RegisterDeploymentArchiveServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterDeploymentArchiveServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterDeploymentArchiveServiceHandler(ctx, mux, conn)
}

// RegisterDeploymentArchiveServiceHandler registers the http handlers for service DeploymentArchiveService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterDeploymentArchiveServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterDeploymentArchiveServiceHandlerClient(ctx, mux, NewDeploymentArchiveServiceClient(conn))
}

// RegisterDeploymentArchiveServiceHandlerClient registers the http handlers for service DeploymentArchiveService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "DeploymentArchiveServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "DeploymentArchiveServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "DeploymentArchiveServiceClient" to call the correct interceptors.
func RegisterDeploymentArchiveServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client DeploymentArchiveServiceClient) error {

	mux.Handle("GET", pattern_DeploymentArchiveService_LegacyListAll_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DeploymentArchiveService_LegacyListAll_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DeploymentArchiveService_LegacyListAll_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_DeploymentArchiveService_GetConfig_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DeploymentArchiveService_GetConfig_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DeploymentArchiveService_GetConfig_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_DeploymentArchiveService_ListProject_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DeploymentArchiveService_ListProject_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DeploymentArchiveService_ListProject_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_DeploymentArchiveService_ListDeployment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_DeploymentArchiveService_ListDeployment_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_DeploymentArchiveService_ListDeployment_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_DeploymentArchiveService_LegacyListAll_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"archived_deployments", "v3"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_DeploymentArchiveService_GetConfig_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3, 1, 0, 4, 1, 5, 4, 2, 5}, []string{"archived_deployments", "v3", "deployment_id.project_id.name", "deployment_id.name", "uuid", "config"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_DeploymentArchiveService_ListProject_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2}, []string{"archived_deployments", "v3", "project_id.name"}, "", runtime.AssumeColonVerbOpt(true)))

	pattern_DeploymentArchiveService_ListDeployment_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 1, 0, 4, 1, 5, 3}, []string{"archived_deployments", "v3", "deployment_id.project_id.name", "deployment_id.name"}, "", runtime.AssumeColonVerbOpt(true)))
)

var (
	forward_DeploymentArchiveService_LegacyListAll_0 = runtime.ForwardResponseStream

	forward_DeploymentArchiveService_GetConfig_0 = runtime.ForwardResponseMessage

	forward_DeploymentArchiveService_ListProject_0 = runtime.ForwardResponseStream

	forward_DeploymentArchiveService_ListDeployment_0 = runtime.ForwardResponseStream
)
