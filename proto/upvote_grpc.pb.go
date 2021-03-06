// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// UpvoteClient is the client API for Upvote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UpvoteClient interface {
	// Increases the Upvote
	Upvote(ctx context.Context, in *UpvoteRequest, opts ...grpc.CallOption) (*UpvoteReply, error)
}

type upvoteClient struct {
	cc grpc.ClientConnInterface
}

func NewUpvoteClient(cc grpc.ClientConnInterface) UpvoteClient {
	return &upvoteClient{cc}
}

func (c *upvoteClient) Upvote(ctx context.Context, in *UpvoteRequest, opts ...grpc.CallOption) (*UpvoteReply, error) {
	out := new(UpvoteReply)
	err := c.cc.Invoke(ctx, "/upvote.Upvote/Upvote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpvoteServer is the server API for Upvote service.
// All implementations must embed UnimplementedUpvoteServer
// for forward compatibility
type UpvoteServer interface {
	// Increases the Upvote
	Upvote(context.Context, *UpvoteRequest) (*UpvoteReply, error)
	mustEmbedUnimplementedUpvoteServer()
}

// UnimplementedUpvoteServer must be embedded to have forward compatible implementations.
type UnimplementedUpvoteServer struct {
}

func (UnimplementedUpvoteServer) Upvote(context.Context, *UpvoteRequest) (*UpvoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upvote not implemented")
}
func (UnimplementedUpvoteServer) mustEmbedUnimplementedUpvoteServer() {}

// UnsafeUpvoteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UpvoteServer will
// result in compilation errors.
type UnsafeUpvoteServer interface {
	mustEmbedUnimplementedUpvoteServer()
}

func RegisterUpvoteServer(s grpc.ServiceRegistrar, srv UpvoteServer) {
	s.RegisterService(&Upvote_ServiceDesc, srv)
}

func _Upvote_Upvote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpvoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpvoteServer).Upvote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/upvote.Upvote/Upvote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpvoteServer).Upvote(ctx, req.(*UpvoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Upvote_ServiceDesc is the grpc.ServiceDesc for Upvote service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Upvote_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "upvote.Upvote",
	HandlerType: (*UpvoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upvote",
			Handler:    _Upvote_Upvote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "upvote.proto",
}
