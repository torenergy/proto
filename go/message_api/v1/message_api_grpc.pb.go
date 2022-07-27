// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: message_api/v1/message_api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessageApiClient is the client API for MessageApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageApiClient interface {
	// Publish a message to the network
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	// Subscribe to a stream of envelopers matching a predicate
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MessageApi_SubscribeClient, error)
	// Query the store for messages
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type messageApiClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageApiClient(cc grpc.ClientConnInterface) MessageApiClient {
	return &messageApiClient{cc}
}

func (c *messageApiClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/xmtp.message_api.v1.MessageApi/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageApiClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (MessageApi_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &MessageApi_ServiceDesc.Streams[0], "/xmtp.message_api.v1.MessageApi/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageApiSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MessageApi_SubscribeClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type messageApiSubscribeClient struct {
	grpc.ClientStream
}

func (x *messageApiSubscribeClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageApiClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/xmtp.message_api.v1.MessageApi/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageApiServer is the server API for MessageApi service.
// All implementations must embed UnimplementedMessageApiServer
// for forward compatibility
type MessageApiServer interface {
	// Publish a message to the network
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	// Subscribe to a stream of envelopers matching a predicate
	Subscribe(*SubscribeRequest, MessageApi_SubscribeServer) error
	// Query the store for messages
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
	mustEmbedUnimplementedMessageApiServer()
}

// UnimplementedMessageApiServer must be embedded to have forward compatible implementations.
type UnimplementedMessageApiServer struct {
}

func (UnimplementedMessageApiServer) Publish(context.Context, *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedMessageApiServer) Subscribe(*SubscribeRequest, MessageApi_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedMessageApiServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}
func (UnimplementedMessageApiServer) mustEmbedUnimplementedMessageApiServer() {}

// UnsafeMessageApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageApiServer will
// result in compilation errors.
type UnsafeMessageApiServer interface {
	mustEmbedUnimplementedMessageApiServer()
}

func RegisterMessageApiServer(s grpc.ServiceRegistrar, srv MessageApiServer) {
	s.RegisterService(&MessageApi_ServiceDesc, srv)
}

func _MessageApi_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageApiServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmtp.message_api.v1.MessageApi/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageApiServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageApi_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessageApiServer).Subscribe(m, &messageApiSubscribeServer{stream})
}

type MessageApi_SubscribeServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type messageApiSubscribeServer struct {
	grpc.ServerStream
}

func (x *messageApiSubscribeServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func _MessageApi_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageApiServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/xmtp.message_api.v1.MessageApi/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageApiServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MessageApi_ServiceDesc is the grpc.ServiceDesc for MessageApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xmtp.message_api.v1.MessageApi",
	HandlerType: (*MessageApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _MessageApi_Publish_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _MessageApi_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _MessageApi_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message_api/v1/message_api.proto",
}