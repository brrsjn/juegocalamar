// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// DataNameNodeServiceClient is the client API for DataNameNodeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataNameNodeServiceClient interface {
	RegistraJugada(ctx context.Context, in *JugadaToDataNode, opts ...grpc.CallOption) (*Response, error)
}

type dataNameNodeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataNameNodeServiceClient(cc grpc.ClientConnInterface) DataNameNodeServiceClient {
	return &dataNameNodeServiceClient{cc}
}

func (c *dataNameNodeServiceClient) RegistraJugada(ctx context.Context, in *JugadaToDataNode, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/DataNameNodeService/RegistraJugada", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataNameNodeServiceServer is the server API for DataNameNodeService service.
// All implementations must embed UnimplementedDataNameNodeServiceServer
// for forward compatibility
type DataNameNodeServiceServer interface {
	RegistraJugada(context.Context, *JugadaToDataNode) (*Response, error)
	mustEmbedUnimplementedDataNameNodeServiceServer()
}

// UnimplementedDataNameNodeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataNameNodeServiceServer struct {
}

func (UnimplementedDataNameNodeServiceServer) RegistraJugada(context.Context, *JugadaToDataNode) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistraJugada not implemented")
}
func (UnimplementedDataNameNodeServiceServer) mustEmbedUnimplementedDataNameNodeServiceServer() {}

// UnsafeDataNameNodeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataNameNodeServiceServer will
// result in compilation errors.
type UnsafeDataNameNodeServiceServer interface {
	mustEmbedUnimplementedDataNameNodeServiceServer()
}

func RegisterDataNameNodeServiceServer(s grpc.ServiceRegistrar, srv DataNameNodeServiceServer) {
	s.RegisterService(&DataNameNodeService_ServiceDesc, srv)
}

func _DataNameNodeService_RegistraJugada_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JugadaToDataNode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataNameNodeServiceServer).RegistraJugada(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DataNameNodeService/RegistraJugada",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataNameNodeServiceServer).RegistraJugada(ctx, req.(*JugadaToDataNode))
	}
	return interceptor(ctx, in, info, handler)
}

// DataNameNodeService_ServiceDesc is the grpc.ServiceDesc for DataNameNodeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataNameNodeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DataNameNodeService",
	HandlerType: (*DataNameNodeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistraJugada",
			Handler:    _DataNameNodeService_RegistraJugada_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/data_nameNode.proto",
}