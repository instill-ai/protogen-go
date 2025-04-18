// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: pipeline/pipeline/v1beta/pipeline_private_service.proto

/*
Package pipelinev1beta is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package pipelinev1beta

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var (
	_ codes.Code
	_ io.Reader
	_ status.Status
	_ = errors.New
	_ = runtime.String
	_ = utilities.NewDoubleArray
	_ = metadata.Join
)

func request_PipelinePrivateService_ListPipelinesAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq ListPipelinesAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.ListPipelinesAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_PipelinePrivateService_ListPipelinesAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq ListPipelinesAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.ListPipelinesAdmin(ctx, &protoReq)
	return msg, metadata, err
}

func request_PipelinePrivateService_LookUpPipelineAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq LookUpPipelineAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.LookUpPipelineAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_PipelinePrivateService_LookUpPipelineAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq LookUpPipelineAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.LookUpPipelineAdmin(ctx, &protoReq)
	return msg, metadata, err
}

func request_PipelinePrivateService_ListPipelineReleasesAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq ListPipelineReleasesAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.ListPipelineReleasesAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_PipelinePrivateService_ListPipelineReleasesAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq ListPipelineReleasesAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.ListPipelineReleasesAdmin(ctx, &protoReq)
	return msg, metadata, err
}

func request_PipelinePrivateService_LookUpConnectionAdmin_0(ctx context.Context, marshaler runtime.Marshaler, client PipelinePrivateServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq LookUpConnectionAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := client.LookUpConnectionAdmin(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err
}

func local_request_PipelinePrivateService_LookUpConnectionAdmin_0(ctx context.Context, marshaler runtime.Marshaler, server PipelinePrivateServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var (
		protoReq LookUpConnectionAdminRequest
		metadata runtime.ServerMetadata
	)
	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && !errors.Is(err, io.EOF) {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	msg, err := server.LookUpConnectionAdmin(ctx, &protoReq)
	return msg, metadata, err
}

// RegisterPipelinePrivateServiceHandlerServer registers the http handlers for service PipelinePrivateService to "mux".
// UnaryRPC     :call PipelinePrivateServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterPipelinePrivateServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterPipelinePrivateServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server PipelinePrivateServiceServer) error {
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_ListPipelinesAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelinesAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelinesAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_LookUpPipelineAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpPipelineAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpPipelineAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_ListPipelineReleasesAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelineReleasesAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelineReleasesAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_ListPipelineReleasesAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_ListPipelineReleasesAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_LookUpConnectionAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpConnectionAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpConnectionAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_PipelinePrivateService_LookUpConnectionAdmin_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_LookUpConnectionAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})

	return nil
}

// RegisterPipelinePrivateServiceHandlerFromEndpoint is same as RegisterPipelinePrivateServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterPipelinePrivateServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()
	return RegisterPipelinePrivateServiceHandler(ctx, mux, conn)
}

// RegisterPipelinePrivateServiceHandler registers the http handlers for service PipelinePrivateService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterPipelinePrivateServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterPipelinePrivateServiceHandlerClient(ctx, mux, NewPipelinePrivateServiceClient(conn))
}

// RegisterPipelinePrivateServiceHandlerClient registers the http handlers for service PipelinePrivateService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "PipelinePrivateServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "PipelinePrivateServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "PipelinePrivateServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterPipelinePrivateServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client PipelinePrivateServiceClient) error {
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_ListPipelinesAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelinesAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelinesAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_ListPipelinesAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_LookUpPipelineAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpPipelineAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpPipelineAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_LookUpPipelineAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_ListPipelineReleasesAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelineReleasesAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/ListPipelineReleasesAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_ListPipelineReleasesAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_ListPipelineReleasesAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	mux.Handle(http.MethodPost, pattern_PipelinePrivateService_LookUpConnectionAdmin_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		annotatedContext, err := runtime.AnnotateContext(ctx, mux, req, "/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpConnectionAdmin", runtime.WithHTTPPathPattern("/pipeline.pipeline.v1beta.PipelinePrivateService/LookUpConnectionAdmin"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_PipelinePrivateService_LookUpConnectionAdmin_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		forward_PipelinePrivateService_LookUpConnectionAdmin_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
	})
	return nil
}

var (
	pattern_PipelinePrivateService_ListPipelinesAdmin_0        = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"pipeline.pipeline.v1beta.PipelinePrivateService", "ListPipelinesAdmin"}, ""))
	pattern_PipelinePrivateService_LookUpPipelineAdmin_0       = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"pipeline.pipeline.v1beta.PipelinePrivateService", "LookUpPipelineAdmin"}, ""))
	pattern_PipelinePrivateService_ListPipelineReleasesAdmin_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"pipeline.pipeline.v1beta.PipelinePrivateService", "ListPipelineReleasesAdmin"}, ""))
	pattern_PipelinePrivateService_LookUpConnectionAdmin_0     = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"pipeline.pipeline.v1beta.PipelinePrivateService", "LookUpConnectionAdmin"}, ""))
)

var (
	forward_PipelinePrivateService_ListPipelinesAdmin_0        = runtime.ForwardResponseMessage
	forward_PipelinePrivateService_LookUpPipelineAdmin_0       = runtime.ForwardResponseMessage
	forward_PipelinePrivateService_ListPipelineReleasesAdmin_0 = runtime.ForwardResponseMessage
	forward_PipelinePrivateService_LookUpConnectionAdmin_0     = runtime.ForwardResponseMessage
)
