// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: v1/msg.proto

package examplev1

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreateContent creates content.
	CreateContent(ctx context.Context, in *MsgCreateContent, opts ...grpc.CallOption) (*MsgCreateContentResponse, error)
	// UpdateContent updates content.
	UpdateContent(ctx context.Context, in *MsgUpdateContent, opts ...grpc.CallOption) (*MsgUpdateContentResponse, error)
	// DeleteContent deletes content.
	DeleteContent(ctx context.Context, in *MsgDeleteContent, opts ...grpc.CallOption) (*MsgDeleteContentResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateContent(ctx context.Context, in *MsgCreateContent, opts ...grpc.CallOption) (*MsgCreateContentResponse, error) {
	out := new(MsgCreateContentResponse)
	err := c.cc.Invoke(ctx, "/chora.example.v1.Msg/CreateContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateContent(ctx context.Context, in *MsgUpdateContent, opts ...grpc.CallOption) (*MsgUpdateContentResponse, error) {
	out := new(MsgUpdateContentResponse)
	err := c.cc.Invoke(ctx, "/chora.example.v1.Msg/UpdateContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteContent(ctx context.Context, in *MsgDeleteContent, opts ...grpc.CallOption) (*MsgDeleteContentResponse, error) {
	out := new(MsgDeleteContentResponse)
	err := c.cc.Invoke(ctx, "/chora.example.v1.Msg/DeleteContent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreateContent creates content.
	CreateContent(context.Context, *MsgCreateContent) (*MsgCreateContentResponse, error)
	// UpdateContent updates content.
	UpdateContent(context.Context, *MsgUpdateContent) (*MsgUpdateContentResponse, error)
	// DeleteContent deletes content.
	DeleteContent(context.Context, *MsgDeleteContent) (*MsgDeleteContentResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateContent(context.Context, *MsgCreateContent) (*MsgCreateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateContent not implemented")
}
func (UnimplementedMsgServer) UpdateContent(context.Context, *MsgUpdateContent) (*MsgUpdateContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContent not implemented")
}
func (UnimplementedMsgServer) DeleteContent(context.Context, *MsgDeleteContent) (*MsgDeleteContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteContent not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.example.v1.Msg/CreateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateContent(ctx, req.(*MsgCreateContent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.example.v1.Msg/UpdateContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateContent(ctx, req.(*MsgUpdateContent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteContent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chora.example.v1.Msg/DeleteContent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteContent(ctx, req.(*MsgDeleteContent))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chora.example.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateContent",
			Handler:    _Msg_CreateContent_Handler,
		},
		{
			MethodName: "UpdateContent",
			Handler:    _Msg_UpdateContent_Handler,
		},
		{
			MethodName: "DeleteContent",
			Handler:    _Msg_DeleteContent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/msg.proto",
}
