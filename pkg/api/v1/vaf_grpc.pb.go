// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// VafServiceClient is the client API for VafService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VafServiceClient interface {
	// StartLocalFunction starts a Service Function on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the vaf/config.yaml
	//   3. all bytes constituting the Function YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalFunction(ctx context.Context, opts ...grpc.CallOption) (VafService_StartLocalFunctionClient, error)
	// StartFromPreviousFunction starts a new Function based on a previous one.
	// If the previous Function does not have the can-replay condition set this call will result in an error.
	StartFromPreviousFunction(ctx context.Context, in *StartFromPreviousFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error)
	// StartFunctionRequest starts a new Function based on its specification.
	StartFunction(ctx context.Context, in *StartFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error)
	// Searches for Function(s) known to this Function
	ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error)
	// Subscribe listens to new Function(s) updates
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (VafService_SubscribeClient, error)
	// GetFunction retrieves details of a single Function
	GetFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*GetFunctionResponse, error)
	// Listen listens to Function updates and log output of a running Function
	Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (VafService_ListenClient, error)
	// StopFunction stops a currently running Function
	StopFunction(ctx context.Context, in *StopFunctionRequest, opts ...grpc.CallOption) (*StopFunctionResponse, error)
}

type vafServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVafServiceClient(cc grpc.ClientConnInterface) VafServiceClient {
	return &vafServiceClient{cc}
}

func (c *vafServiceClient) StartLocalFunction(ctx context.Context, opts ...grpc.CallOption) (VafService_StartLocalFunctionClient, error) {
	stream, err := c.cc.NewStream(ctx, &VafService_ServiceDesc.Streams[0], "/v1.VafService/StartLocalFunction", opts...)
	if err != nil {
		return nil, err
	}
	x := &vafServiceStartLocalFunctionClient{stream}
	return x, nil
}

type VafService_StartLocalFunctionClient interface {
	Send(*StartLocalFunctionRequest) error
	CloseAndRecv() (*StartFunctionResponse, error)
	grpc.ClientStream
}

type vafServiceStartLocalFunctionClient struct {
	grpc.ClientStream
}

func (x *vafServiceStartLocalFunctionClient) Send(m *StartLocalFunctionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *vafServiceStartLocalFunctionClient) CloseAndRecv() (*StartFunctionResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StartFunctionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vafServiceClient) StartFromPreviousFunction(ctx context.Context, in *StartFromPreviousFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error) {
	out := new(StartFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VafService/StartFromPreviousFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vafServiceClient) StartFunction(ctx context.Context, in *StartFunctionRequest, opts ...grpc.CallOption) (*StartFunctionResponse, error) {
	out := new(StartFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VafService/StartFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vafServiceClient) ListFunctions(ctx context.Context, in *ListFunctionsRequest, opts ...grpc.CallOption) (*ListFunctionsResponse, error) {
	out := new(ListFunctionsResponse)
	err := c.cc.Invoke(ctx, "/v1.VafService/ListFunctions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vafServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (VafService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &VafService_ServiceDesc.Streams[1], "/v1.VafService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &vafServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VafService_SubscribeClient interface {
	Recv() (*SubscribeResponse, error)
	grpc.ClientStream
}

type vafServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *vafServiceSubscribeClient) Recv() (*SubscribeResponse, error) {
	m := new(SubscribeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vafServiceClient) GetFunction(ctx context.Context, in *GetFunctionRequest, opts ...grpc.CallOption) (*GetFunctionResponse, error) {
	out := new(GetFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VafService/GetFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vafServiceClient) Listen(ctx context.Context, in *ListenRequest, opts ...grpc.CallOption) (VafService_ListenClient, error) {
	stream, err := c.cc.NewStream(ctx, &VafService_ServiceDesc.Streams[2], "/v1.VafService/Listen", opts...)
	if err != nil {
		return nil, err
	}
	x := &vafServiceListenClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VafService_ListenClient interface {
	Recv() (*ListenResponse, error)
	grpc.ClientStream
}

type vafServiceListenClient struct {
	grpc.ClientStream
}

func (x *vafServiceListenClient) Recv() (*ListenResponse, error) {
	m := new(ListenResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *vafServiceClient) StopFunction(ctx context.Context, in *StopFunctionRequest, opts ...grpc.CallOption) (*StopFunctionResponse, error) {
	out := new(StopFunctionResponse)
	err := c.cc.Invoke(ctx, "/v1.VafService/StopFunction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VafServiceServer is the server API for VafService service.
// All implementations must embed UnimplementedVafServiceServer
// for forward compatibility
type VafServiceServer interface {
	// StartLocalFunction starts a Service Function on the Bhojpur.NET Platform directly.
	// The incoming requests are expected in the following order:
	//   1. metadata
	//   2. all bytes constituting the vaf/config.yaml
	//   3. all bytes constituting the Function YAML that will be executed (that the config.yaml points to)
	//   4. all bytes constituting the gzipped Bhojpur.NET Platform application tar stream
	//   5. the Bhojpur.NET Platform application tar stream done marker
	StartLocalFunction(VafService_StartLocalFunctionServer) error
	// StartFromPreviousFunction starts a new Function based on a previous one.
	// If the previous Function does not have the can-replay condition set this call will result in an error.
	StartFromPreviousFunction(context.Context, *StartFromPreviousFunctionRequest) (*StartFunctionResponse, error)
	// StartFunctionRequest starts a new Function based on its specification.
	StartFunction(context.Context, *StartFunctionRequest) (*StartFunctionResponse, error)
	// Searches for Function(s) known to this Function
	ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error)
	// Subscribe listens to new Function(s) updates
	Subscribe(*SubscribeRequest, VafService_SubscribeServer) error
	// GetFunction retrieves details of a single Function
	GetFunction(context.Context, *GetFunctionRequest) (*GetFunctionResponse, error)
	// Listen listens to Function updates and log output of a running Function
	Listen(*ListenRequest, VafService_ListenServer) error
	// StopFunction stops a currently running Function
	StopFunction(context.Context, *StopFunctionRequest) (*StopFunctionResponse, error)
	mustEmbedUnimplementedVafServiceServer()
}

// UnimplementedVafServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVafServiceServer struct {
}

func (UnimplementedVafServiceServer) StartLocalFunction(VafService_StartLocalFunctionServer) error {
	return status.Errorf(codes.Unimplemented, "method StartLocalFunction not implemented")
}
func (UnimplementedVafServiceServer) StartFromPreviousFunction(context.Context, *StartFromPreviousFunctionRequest) (*StartFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFromPreviousFunction not implemented")
}
func (UnimplementedVafServiceServer) StartFunction(context.Context, *StartFunctionRequest) (*StartFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFunction not implemented")
}
func (UnimplementedVafServiceServer) ListFunctions(context.Context, *ListFunctionsRequest) (*ListFunctionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFunctions not implemented")
}
func (UnimplementedVafServiceServer) Subscribe(*SubscribeRequest, VafService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedVafServiceServer) GetFunction(context.Context, *GetFunctionRequest) (*GetFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFunction not implemented")
}
func (UnimplementedVafServiceServer) Listen(*ListenRequest, VafService_ListenServer) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedVafServiceServer) StopFunction(context.Context, *StopFunctionRequest) (*StopFunctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopFunction not implemented")
}
func (UnimplementedVafServiceServer) mustEmbedUnimplementedVafServiceServer() {}

// UnsafeVafServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VafServiceServer will
// result in compilation errors.
type UnsafeVafServiceServer interface {
	mustEmbedUnimplementedVafServiceServer()
}

func RegisterVafServiceServer(s grpc.ServiceRegistrar, srv VafServiceServer) {
	s.RegisterService(&VafService_ServiceDesc, srv)
}

func _VafService_StartLocalFunction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VafServiceServer).StartLocalFunction(&vafServiceStartLocalFunctionServer{stream})
}

type VafService_StartLocalFunctionServer interface {
	SendAndClose(*StartFunctionResponse) error
	Recv() (*StartLocalFunctionRequest, error)
	grpc.ServerStream
}

type vafServiceStartLocalFunctionServer struct {
	grpc.ServerStream
}

func (x *vafServiceStartLocalFunctionServer) SendAndClose(m *StartFunctionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *vafServiceStartLocalFunctionServer) Recv() (*StartLocalFunctionRequest, error) {
	m := new(StartLocalFunctionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _VafService_StartFromPreviousFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFromPreviousFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VafServiceServer).StartFromPreviousFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VafService/StartFromPreviousFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VafServiceServer).StartFromPreviousFunction(ctx, req.(*StartFromPreviousFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VafService_StartFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VafServiceServer).StartFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VafService/StartFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VafServiceServer).StartFunction(ctx, req.(*StartFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VafService_ListFunctions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFunctionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VafServiceServer).ListFunctions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VafService/ListFunctions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VafServiceServer).ListFunctions(ctx, req.(*ListFunctionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VafService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VafServiceServer).Subscribe(m, &vafServiceSubscribeServer{stream})
}

type VafService_SubscribeServer interface {
	Send(*SubscribeResponse) error
	grpc.ServerStream
}

type vafServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *vafServiceSubscribeServer) Send(m *SubscribeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VafService_GetFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VafServiceServer).GetFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VafService/GetFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VafServiceServer).GetFunction(ctx, req.(*GetFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VafService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VafServiceServer).Listen(m, &vafServiceListenServer{stream})
}

type VafService_ListenServer interface {
	Send(*ListenResponse) error
	grpc.ServerStream
}

type vafServiceListenServer struct {
	grpc.ServerStream
}

func (x *vafServiceListenServer) Send(m *ListenResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _VafService_StopFunction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopFunctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VafServiceServer).StopFunction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.VafService/StopFunction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VafServiceServer).StopFunction(ctx, req.(*StopFunctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VafService_ServiceDesc is the grpc.ServiceDesc for VafService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VafService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.VafService",
	HandlerType: (*VafServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartFromPreviousFunction",
			Handler:    _VafService_StartFromPreviousFunction_Handler,
		},
		{
			MethodName: "StartFunction",
			Handler:    _VafService_StartFunction_Handler,
		},
		{
			MethodName: "ListFunctions",
			Handler:    _VafService_ListFunctions_Handler,
		},
		{
			MethodName: "GetFunction",
			Handler:    _VafService_GetFunction_Handler,
		},
		{
			MethodName: "StopFunction",
			Handler:    _VafService_StopFunction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartLocalFunction",
			Handler:       _VafService_StartLocalFunction_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Subscribe",
			Handler:       _VafService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Listen",
			Handler:       _VafService_Listen_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "vaf.proto",
}
