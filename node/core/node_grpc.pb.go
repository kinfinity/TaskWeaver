// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: node.proto

package core

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
	NodeService_GetNodeInfo_FullMethodName               = "/core.NodeService/GetNodeInfo"
	NodeService_AuthenticateNode_FullMethodName          = "/core.NodeService/AuthenticateNode"
	NodeService_UpdateNodeStatus_FullMethodName          = "/core.NodeService/UpdateNodeStatus"
	NodeService_AssignTask_FullMethodName                = "/core.NodeService/AssignTask"
	NodeService_UpdateTaskStatus_FullMethodName          = "/core.NodeService/UpdateTaskStatus"
	NodeService_GetTaskStatus_FullMethodName             = "/core.NodeService/GetTaskStatus"
	NodeService_CancelTask_FullMethodName                = "/core.NodeService/CancelTask"
	NodeService_GetTaskList_FullMethodName               = "/core.NodeService/GetTaskList"
	NodeService_AcknowledgeTaskCompletion_FullMethodName = "/core.NodeService/AcknowledgeTaskCompletion"
	NodeService_HandleTaskError_FullMethodName           = "/core.NodeService/HandleTaskError"
	NodeService_SendTaskProgress_FullMethodName          = "/core.NodeService/SendTaskProgress"
)

// NodeServiceClient is the client API for NodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeServiceClient interface {
	GetNodeInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	AuthenticateNode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	UpdateNodeStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// Task management and communication
	AssignTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_AssignTaskClient, error)
	UpdateTaskStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_UpdateTaskStatusClient, error)
	GetTaskStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	CancelTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetTaskList(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_GetTaskListClient, error)
	AcknowledgeTaskCompletion(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	HandleTaskError(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	SendTaskProgress(ctx context.Context, in *TaskUpdate, opts ...grpc.CallOption) (*Response, error)
}

type nodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeServiceClient(cc grpc.ClientConnInterface) NodeServiceClient {
	return &nodeServiceClient{cc}
}

func (c *nodeServiceClient) GetNodeInfo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_GetNodeInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) AuthenticateNode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_AuthenticateNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) UpdateNodeStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_UpdateNodeStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) AssignTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_AssignTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[0], NodeService_AssignTask_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceAssignTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_AssignTaskClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type nodeServiceAssignTaskClient struct {
	grpc.ClientStream
}

func (x *nodeServiceAssignTaskClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) UpdateTaskStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_UpdateTaskStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[1], NodeService_UpdateTaskStatus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceUpdateTaskStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_UpdateTaskStatusClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type nodeServiceUpdateTaskStatusClient struct {
	grpc.ClientStream
}

func (x *nodeServiceUpdateTaskStatusClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) GetTaskStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_GetTaskStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) CancelTask(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_CancelTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) GetTaskList(ctx context.Context, in *Request, opts ...grpc.CallOption) (NodeService_GetTaskListClient, error) {
	stream, err := c.cc.NewStream(ctx, &NodeService_ServiceDesc.Streams[2], NodeService_GetTaskList_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeServiceGetTaskListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NodeService_GetTaskListClient interface {
	Recv() (*Task, error)
	grpc.ClientStream
}

type nodeServiceGetTaskListClient struct {
	grpc.ClientStream
}

func (x *nodeServiceGetTaskListClient) Recv() (*Task, error) {
	m := new(Task)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeServiceClient) AcknowledgeTaskCompletion(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_AcknowledgeTaskCompletion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) HandleTaskError(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_HandleTaskError_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeServiceClient) SendTaskProgress(ctx context.Context, in *TaskUpdate, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, NodeService_SendTaskProgress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServiceServer is the server API for NodeService service.
// All implementations must embed UnimplementedNodeServiceServer
// for forward compatibility
type NodeServiceServer interface {
	GetNodeInfo(context.Context, *Request) (*Response, error)
	AuthenticateNode(context.Context, *Request) (*Response, error)
	UpdateNodeStatus(context.Context, *Request) (*Response, error)
	// Task management and communication
	AssignTask(*Request, NodeService_AssignTaskServer) error
	UpdateTaskStatus(*Request, NodeService_UpdateTaskStatusServer) error
	GetTaskStatus(context.Context, *Request) (*Response, error)
	CancelTask(context.Context, *Request) (*Response, error)
	GetTaskList(*Request, NodeService_GetTaskListServer) error
	AcknowledgeTaskCompletion(context.Context, *Request) (*Response, error)
	HandleTaskError(context.Context, *Request) (*Response, error)
	SendTaskProgress(context.Context, *TaskUpdate) (*Response, error)
	mustEmbedUnimplementedNodeServiceServer()
}

// UnimplementedNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServiceServer struct {
}

func (UnimplementedNodeServiceServer) GetNodeInfo(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeInfo not implemented")
}
func (UnimplementedNodeServiceServer) AuthenticateNode(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticateNode not implemented")
}
func (UnimplementedNodeServiceServer) UpdateNodeStatus(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNodeStatus not implemented")
}
func (UnimplementedNodeServiceServer) AssignTask(*Request, NodeService_AssignTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method AssignTask not implemented")
}
func (UnimplementedNodeServiceServer) UpdateTaskStatus(*Request, NodeService_UpdateTaskStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateTaskStatus not implemented")
}
func (UnimplementedNodeServiceServer) GetTaskStatus(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskStatus not implemented")
}
func (UnimplementedNodeServiceServer) CancelTask(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTask not implemented")
}
func (UnimplementedNodeServiceServer) GetTaskList(*Request, NodeService_GetTaskListServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTaskList not implemented")
}
func (UnimplementedNodeServiceServer) AcknowledgeTaskCompletion(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcknowledgeTaskCompletion not implemented")
}
func (UnimplementedNodeServiceServer) HandleTaskError(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleTaskError not implemented")
}
func (UnimplementedNodeServiceServer) SendTaskProgress(context.Context, *TaskUpdate) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTaskProgress not implemented")
}
func (UnimplementedNodeServiceServer) mustEmbedUnimplementedNodeServiceServer() {}

// UnsafeNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServiceServer will
// result in compilation errors.
type UnsafeNodeServiceServer interface {
	mustEmbedUnimplementedNodeServiceServer()
}

func RegisterNodeServiceServer(s grpc.ServiceRegistrar, srv NodeServiceServer) {
	s.RegisterService(&NodeService_ServiceDesc, srv)
}

func _NodeService_GetNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetNodeInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetNodeInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetNodeInfo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_AuthenticateNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).AuthenticateNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_AuthenticateNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).AuthenticateNode(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_UpdateNodeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).UpdateNodeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_UpdateNodeStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).UpdateNodeStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_AssignTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).AssignTask(m, &nodeServiceAssignTaskServer{stream})
}

type NodeService_AssignTaskServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type nodeServiceAssignTaskServer struct {
	grpc.ServerStream
}

func (x *nodeServiceAssignTaskServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_UpdateTaskStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).UpdateTaskStatus(m, &nodeServiceUpdateTaskStatusServer{stream})
}

type NodeService_UpdateTaskStatusServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type nodeServiceUpdateTaskStatusServer struct {
	grpc.ServerStream
}

func (x *nodeServiceUpdateTaskStatusServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_GetTaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).GetTaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_GetTaskStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).GetTaskStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_CancelTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).CancelTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_CancelTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).CancelTask(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_GetTaskList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NodeServiceServer).GetTaskList(m, &nodeServiceGetTaskListServer{stream})
}

type NodeService_GetTaskListServer interface {
	Send(*Task) error
	grpc.ServerStream
}

type nodeServiceGetTaskListServer struct {
	grpc.ServerStream
}

func (x *nodeServiceGetTaskListServer) Send(m *Task) error {
	return x.ServerStream.SendMsg(m)
}

func _NodeService_AcknowledgeTaskCompletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).AcknowledgeTaskCompletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_AcknowledgeTaskCompletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).AcknowledgeTaskCompletion(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_HandleTaskError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).HandleTaskError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_HandleTaskError_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).HandleTaskError(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeService_SendTaskProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServiceServer).SendTaskProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeService_SendTaskProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServiceServer).SendTaskProgress(ctx, req.(*TaskUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeService_ServiceDesc is the grpc.ServiceDesc for NodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.NodeService",
	HandlerType: (*NodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeInfo",
			Handler:    _NodeService_GetNodeInfo_Handler,
		},
		{
			MethodName: "AuthenticateNode",
			Handler:    _NodeService_AuthenticateNode_Handler,
		},
		{
			MethodName: "UpdateNodeStatus",
			Handler:    _NodeService_UpdateNodeStatus_Handler,
		},
		{
			MethodName: "GetTaskStatus",
			Handler:    _NodeService_GetTaskStatus_Handler,
		},
		{
			MethodName: "CancelTask",
			Handler:    _NodeService_CancelTask_Handler,
		},
		{
			MethodName: "AcknowledgeTaskCompletion",
			Handler:    _NodeService_AcknowledgeTaskCompletion_Handler,
		},
		{
			MethodName: "HandleTaskError",
			Handler:    _NodeService_HandleTaskError_Handler,
		},
		{
			MethodName: "SendTaskProgress",
			Handler:    _NodeService_SendTaskProgress_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AssignTask",
			Handler:       _NodeService_AssignTask_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateTaskStatus",
			Handler:       _NodeService_UpdateTaskStatus_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTaskList",
			Handler:       _NodeService_GetTaskList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}
