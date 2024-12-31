// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proxy/v1/proxy.proto

package proxyv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/deeploop-io/deeploop-protocol/genproto/proxy/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ProxyServiceName is the fully-qualified name of the ProxyService service.
	ProxyServiceName = "deeploop.proxy.v1.ProxyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProxyServiceRPCProcedure is the fully-qualified name of the ProxyService's RPC RPC.
	ProxyServiceRPCProcedure = "/deeploop.proxy.v1.ProxyService/RPC"
	// ProxyServiceAuthenticateProcedure is the fully-qualified name of the ProxyService's Authenticate
	// RPC.
	ProxyServiceAuthenticateProcedure = "/deeploop.proxy.v1.ProxyService/Authenticate"
	// ProxyServiceSubscribeAclProcedure is the fully-qualified name of the ProxyService's SubscribeAcl
	// RPC.
	ProxyServiceSubscribeAclProcedure = "/deeploop.proxy.v1.ProxyService/SubscribeAcl"
	// ProxyServiceOnSubscribeProcedure is the fully-qualified name of the ProxyService's OnSubscribe
	// RPC.
	ProxyServiceOnSubscribeProcedure = "/deeploop.proxy.v1.ProxyService/OnSubscribe"
	// ProxyServiceOnUnsubscribeProcedure is the fully-qualified name of the ProxyService's
	// OnUnsubscribe RPC.
	ProxyServiceOnUnsubscribeProcedure = "/deeploop.proxy.v1.ProxyService/OnUnsubscribe"
	// ProxyServiceOnDisconnectProcedure is the fully-qualified name of the ProxyService's OnDisconnect
	// RPC.
	ProxyServiceOnDisconnectProcedure = "/deeploop.proxy.v1.ProxyService/OnDisconnect"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	proxyServiceServiceDescriptor             = v1.File_proxy_v1_proxy_proto.Services().ByName("ProxyService")
	proxyServiceRPCMethodDescriptor           = proxyServiceServiceDescriptor.Methods().ByName("RPC")
	proxyServiceAuthenticateMethodDescriptor  = proxyServiceServiceDescriptor.Methods().ByName("Authenticate")
	proxyServiceSubscribeAclMethodDescriptor  = proxyServiceServiceDescriptor.Methods().ByName("SubscribeAcl")
	proxyServiceOnSubscribeMethodDescriptor   = proxyServiceServiceDescriptor.Methods().ByName("OnSubscribe")
	proxyServiceOnUnsubscribeMethodDescriptor = proxyServiceServiceDescriptor.Methods().ByName("OnUnsubscribe")
	proxyServiceOnDisconnectMethodDescriptor  = proxyServiceServiceDescriptor.Methods().ByName("OnDisconnect")
)

// ProxyServiceClient is a client for the deeploop.proxy.v1.ProxyService service.
type ProxyServiceClient interface {
	// 远程调用
	RPC(context.Context, *connect.Request[v1.RPCRequest]) (*connect.Response[v1.RPCReply], error)
	// 认证
	Authenticate(context.Context, *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateReply], error)
	// 订阅权限校验
	SubscribeAcl(context.Context, *connect.Request[v1.SubscribeAclRequest]) (*connect.Response[v1.SubscribeAclReply], error)
	OnSubscribe(context.Context, *connect.Request[v1.OnSubscribeRequest]) (*connect.Response[v1.OnSubscribeReply], error)
	OnUnsubscribe(context.Context, *connect.Request[v1.OnUnsubscribeRequest]) (*connect.Response[v1.OnUnsubscribeReply], error)
	OnDisconnect(context.Context, *connect.Request[v1.OnDisconnectRequest]) (*connect.Response[v1.OnDisconnectReply], error)
}

// NewProxyServiceClient constructs a client for the deeploop.proxy.v1.ProxyService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProxyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ProxyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &proxyServiceClient{
		rPC: connect.NewClient[v1.RPCRequest, v1.RPCReply](
			httpClient,
			baseURL+ProxyServiceRPCProcedure,
			connect.WithSchema(proxyServiceRPCMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		authenticate: connect.NewClient[v1.AuthenticateRequest, v1.AuthenticateReply](
			httpClient,
			baseURL+ProxyServiceAuthenticateProcedure,
			connect.WithSchema(proxyServiceAuthenticateMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		subscribeAcl: connect.NewClient[v1.SubscribeAclRequest, v1.SubscribeAclReply](
			httpClient,
			baseURL+ProxyServiceSubscribeAclProcedure,
			connect.WithSchema(proxyServiceSubscribeAclMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		onSubscribe: connect.NewClient[v1.OnSubscribeRequest, v1.OnSubscribeReply](
			httpClient,
			baseURL+ProxyServiceOnSubscribeProcedure,
			connect.WithSchema(proxyServiceOnSubscribeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		onUnsubscribe: connect.NewClient[v1.OnUnsubscribeRequest, v1.OnUnsubscribeReply](
			httpClient,
			baseURL+ProxyServiceOnUnsubscribeProcedure,
			connect.WithSchema(proxyServiceOnUnsubscribeMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		onDisconnect: connect.NewClient[v1.OnDisconnectRequest, v1.OnDisconnectReply](
			httpClient,
			baseURL+ProxyServiceOnDisconnectProcedure,
			connect.WithSchema(proxyServiceOnDisconnectMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// proxyServiceClient implements ProxyServiceClient.
type proxyServiceClient struct {
	rPC           *connect.Client[v1.RPCRequest, v1.RPCReply]
	authenticate  *connect.Client[v1.AuthenticateRequest, v1.AuthenticateReply]
	subscribeAcl  *connect.Client[v1.SubscribeAclRequest, v1.SubscribeAclReply]
	onSubscribe   *connect.Client[v1.OnSubscribeRequest, v1.OnSubscribeReply]
	onUnsubscribe *connect.Client[v1.OnUnsubscribeRequest, v1.OnUnsubscribeReply]
	onDisconnect  *connect.Client[v1.OnDisconnectRequest, v1.OnDisconnectReply]
}

// RPC calls deeploop.proxy.v1.ProxyService.RPC.
func (c *proxyServiceClient) RPC(ctx context.Context, req *connect.Request[v1.RPCRequest]) (*connect.Response[v1.RPCReply], error) {
	return c.rPC.CallUnary(ctx, req)
}

// Authenticate calls deeploop.proxy.v1.ProxyService.Authenticate.
func (c *proxyServiceClient) Authenticate(ctx context.Context, req *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateReply], error) {
	return c.authenticate.CallUnary(ctx, req)
}

// SubscribeAcl calls deeploop.proxy.v1.ProxyService.SubscribeAcl.
func (c *proxyServiceClient) SubscribeAcl(ctx context.Context, req *connect.Request[v1.SubscribeAclRequest]) (*connect.Response[v1.SubscribeAclReply], error) {
	return c.subscribeAcl.CallUnary(ctx, req)
}

// OnSubscribe calls deeploop.proxy.v1.ProxyService.OnSubscribe.
func (c *proxyServiceClient) OnSubscribe(ctx context.Context, req *connect.Request[v1.OnSubscribeRequest]) (*connect.Response[v1.OnSubscribeReply], error) {
	return c.onSubscribe.CallUnary(ctx, req)
}

// OnUnsubscribe calls deeploop.proxy.v1.ProxyService.OnUnsubscribe.
func (c *proxyServiceClient) OnUnsubscribe(ctx context.Context, req *connect.Request[v1.OnUnsubscribeRequest]) (*connect.Response[v1.OnUnsubscribeReply], error) {
	return c.onUnsubscribe.CallUnary(ctx, req)
}

// OnDisconnect calls deeploop.proxy.v1.ProxyService.OnDisconnect.
func (c *proxyServiceClient) OnDisconnect(ctx context.Context, req *connect.Request[v1.OnDisconnectRequest]) (*connect.Response[v1.OnDisconnectReply], error) {
	return c.onDisconnect.CallUnary(ctx, req)
}

// ProxyServiceHandler is an implementation of the deeploop.proxy.v1.ProxyService service.
type ProxyServiceHandler interface {
	// 远程调用
	RPC(context.Context, *connect.Request[v1.RPCRequest]) (*connect.Response[v1.RPCReply], error)
	// 认证
	Authenticate(context.Context, *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateReply], error)
	// 订阅权限校验
	SubscribeAcl(context.Context, *connect.Request[v1.SubscribeAclRequest]) (*connect.Response[v1.SubscribeAclReply], error)
	OnSubscribe(context.Context, *connect.Request[v1.OnSubscribeRequest]) (*connect.Response[v1.OnSubscribeReply], error)
	OnUnsubscribe(context.Context, *connect.Request[v1.OnUnsubscribeRequest]) (*connect.Response[v1.OnUnsubscribeReply], error)
	OnDisconnect(context.Context, *connect.Request[v1.OnDisconnectRequest]) (*connect.Response[v1.OnDisconnectReply], error)
}

// NewProxyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProxyServiceHandler(svc ProxyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	proxyServiceRPCHandler := connect.NewUnaryHandler(
		ProxyServiceRPCProcedure,
		svc.RPC,
		connect.WithSchema(proxyServiceRPCMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	proxyServiceAuthenticateHandler := connect.NewUnaryHandler(
		ProxyServiceAuthenticateProcedure,
		svc.Authenticate,
		connect.WithSchema(proxyServiceAuthenticateMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	proxyServiceSubscribeAclHandler := connect.NewUnaryHandler(
		ProxyServiceSubscribeAclProcedure,
		svc.SubscribeAcl,
		connect.WithSchema(proxyServiceSubscribeAclMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	proxyServiceOnSubscribeHandler := connect.NewUnaryHandler(
		ProxyServiceOnSubscribeProcedure,
		svc.OnSubscribe,
		connect.WithSchema(proxyServiceOnSubscribeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	proxyServiceOnUnsubscribeHandler := connect.NewUnaryHandler(
		ProxyServiceOnUnsubscribeProcedure,
		svc.OnUnsubscribe,
		connect.WithSchema(proxyServiceOnUnsubscribeMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	proxyServiceOnDisconnectHandler := connect.NewUnaryHandler(
		ProxyServiceOnDisconnectProcedure,
		svc.OnDisconnect,
		connect.WithSchema(proxyServiceOnDisconnectMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/deeploop.proxy.v1.ProxyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ProxyServiceRPCProcedure:
			proxyServiceRPCHandler.ServeHTTP(w, r)
		case ProxyServiceAuthenticateProcedure:
			proxyServiceAuthenticateHandler.ServeHTTP(w, r)
		case ProxyServiceSubscribeAclProcedure:
			proxyServiceSubscribeAclHandler.ServeHTTP(w, r)
		case ProxyServiceOnSubscribeProcedure:
			proxyServiceOnSubscribeHandler.ServeHTTP(w, r)
		case ProxyServiceOnUnsubscribeProcedure:
			proxyServiceOnUnsubscribeHandler.ServeHTTP(w, r)
		case ProxyServiceOnDisconnectProcedure:
			proxyServiceOnDisconnectHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedProxyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProxyServiceHandler struct{}

func (UnimplementedProxyServiceHandler) RPC(context.Context, *connect.Request[v1.RPCRequest]) (*connect.Response[v1.RPCReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("deeploop.proxy.v1.ProxyService.RPC is not implemented"))
}

func (UnimplementedProxyServiceHandler) Authenticate(context.Context, *connect.Request[v1.AuthenticateRequest]) (*connect.Response[v1.AuthenticateReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("deeploop.proxy.v1.ProxyService.Authenticate is not implemented"))
}

func (UnimplementedProxyServiceHandler) SubscribeAcl(context.Context, *connect.Request[v1.SubscribeAclRequest]) (*connect.Response[v1.SubscribeAclReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("deeploop.proxy.v1.ProxyService.SubscribeAcl is not implemented"))
}

func (UnimplementedProxyServiceHandler) OnSubscribe(context.Context, *connect.Request[v1.OnSubscribeRequest]) (*connect.Response[v1.OnSubscribeReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("deeploop.proxy.v1.ProxyService.OnSubscribe is not implemented"))
}

func (UnimplementedProxyServiceHandler) OnUnsubscribe(context.Context, *connect.Request[v1.OnUnsubscribeRequest]) (*connect.Response[v1.OnUnsubscribeReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("deeploop.proxy.v1.ProxyService.OnUnsubscribe is not implemented"))
}

func (UnimplementedProxyServiceHandler) OnDisconnect(context.Context, *connect.Request[v1.OnDisconnectRequest]) (*connect.Response[v1.OnDisconnectReply], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("deeploop.proxy.v1.ProxyService.OnDisconnect is not implemented"))
}