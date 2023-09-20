// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: agent/v1/agent.proto

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

const (
	Agent_CreateAgent_FullMethodName = "/api.agent.v1.Agent/CreateAgent"
	Agent_UpdateAgent_FullMethodName = "/api.agent.v1.Agent/UpdateAgent"
	Agent_DeleteAgent_FullMethodName = "/api.agent.v1.Agent/DeleteAgent"
	Agent_GetAgent_FullMethodName    = "/api.agent.v1.Agent/GetAgent"
	Agent_ListAgent_FullMethodName   = "/api.agent.v1.Agent/ListAgent"
)

// AgentClient is the client API for Agent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentClient interface {
	CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*CreateAgentReply, error)
	UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*UpdateAgentReply, error)
	DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*DeleteAgentReply, error)
	GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentReply, error)
	ListAgent(ctx context.Context, in *ListAgentRequest, opts ...grpc.CallOption) (*ListAgentReply, error)
}

type agentClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentClient(cc grpc.ClientConnInterface) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CreateAgent(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*CreateAgentReply, error) {
	out := new(CreateAgentReply)
	err := c.cc.Invoke(ctx, Agent_CreateAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) UpdateAgent(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*UpdateAgentReply, error) {
	out := new(UpdateAgentReply)
	err := c.cc.Invoke(ctx, Agent_UpdateAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) DeleteAgent(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*DeleteAgentReply, error) {
	out := new(DeleteAgentReply)
	err := c.cc.Invoke(ctx, Agent_DeleteAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) GetAgent(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*GetAgentReply, error) {
	out := new(GetAgentReply)
	err := c.cc.Invoke(ctx, Agent_GetAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) ListAgent(ctx context.Context, in *ListAgentRequest, opts ...grpc.CallOption) (*ListAgentReply, error) {
	out := new(ListAgentReply)
	err := c.cc.Invoke(ctx, Agent_ListAgent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServer is the server API for Agent service.
// All implementations must embed UnimplementedAgentServer
// for forward compatibility
type AgentServer interface {
	CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentReply, error)
	UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentReply, error)
	DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentReply, error)
	GetAgent(context.Context, *GetAgentRequest) (*GetAgentReply, error)
	ListAgent(context.Context, *ListAgentRequest) (*ListAgentReply, error)
	mustEmbedUnimplementedAgentServer()
}

// UnimplementedAgentServer must be embedded to have forward compatible implementations.
type UnimplementedAgentServer struct {
}

func (UnimplementedAgentServer) CreateAgent(context.Context, *CreateAgentRequest) (*CreateAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgent not implemented")
}
func (UnimplementedAgentServer) UpdateAgent(context.Context, *UpdateAgentRequest) (*UpdateAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgent not implemented")
}
func (UnimplementedAgentServer) DeleteAgent(context.Context, *DeleteAgentRequest) (*DeleteAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgent not implemented")
}
func (UnimplementedAgentServer) GetAgent(context.Context, *GetAgentRequest) (*GetAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgent not implemented")
}
func (UnimplementedAgentServer) ListAgent(context.Context, *ListAgentRequest) (*ListAgentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAgent not implemented")
}
func (UnimplementedAgentServer) mustEmbedUnimplementedAgentServer() {}

// UnsafeAgentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServer will
// result in compilation errors.
type UnsafeAgentServer interface {
	mustEmbedUnimplementedAgentServer()
}

func RegisterAgentServer(s grpc.ServiceRegistrar, srv AgentServer) {
	s.RegisterService(&Agent_ServiceDesc, srv)
}

func _Agent_CreateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CreateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_CreateAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CreateAgent(ctx, req.(*CreateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_UpdateAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).UpdateAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_UpdateAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).UpdateAgent(ctx, req.(*UpdateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_DeleteAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).DeleteAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_DeleteAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).DeleteAgent(ctx, req.(*DeleteAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_GetAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).GetAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_GetAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).GetAgent(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_ListAgent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).ListAgent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Agent_ListAgent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).ListAgent(ctx, req.(*ListAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Agent_ServiceDesc is the grpc.ServiceDesc for Agent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Agent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.agent.v1.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAgent",
			Handler:    _Agent_CreateAgent_Handler,
		},
		{
			MethodName: "UpdateAgent",
			Handler:    _Agent_UpdateAgent_Handler,
		},
		{
			MethodName: "DeleteAgent",
			Handler:    _Agent_DeleteAgent_Handler,
		},
		{
			MethodName: "GetAgent",
			Handler:    _Agent_GetAgent_Handler,
		},
		{
			MethodName: "ListAgent",
			Handler:    _Agent_ListAgent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent/v1/agent.proto",
}
