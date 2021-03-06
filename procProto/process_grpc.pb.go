// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package procProto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcServiceClient is the client API for ProcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProcServiceClient interface {
	CreateProc(ctx context.Context, in *RequestForCreateProc, opts ...grpc.CallOption) (*ProcID, error)
	ReadProcByID(ctx context.Context, in *RequestForReadByID, opts ...grpc.CallOption) (*Proc, error)
	UpdateProcByID(ctx context.Context, in *RequestForUpdateProcByID, opts ...grpc.CallOption) (*ProcID, error)
	DeleteProcByID(ctx context.Context, in *RequestForDeleteByID, opts ...grpc.CallOption) (*ProcID, error)
	ReadAllProcs(ctx context.Context, in *RequestForReadAllProcs, opts ...grpc.CallOption) (*ProcList, error)
}

type procServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProcServiceClient(cc grpc.ClientConnInterface) ProcServiceClient {
	return &procServiceClient{cc}
}

func (c *procServiceClient) CreateProc(ctx context.Context, in *RequestForCreateProc, opts ...grpc.CallOption) (*ProcID, error) {
	out := new(ProcID)
	err := c.cc.Invoke(ctx, "/ProcService/CreateProc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procServiceClient) ReadProcByID(ctx context.Context, in *RequestForReadByID, opts ...grpc.CallOption) (*Proc, error) {
	out := new(Proc)
	err := c.cc.Invoke(ctx, "/ProcService/ReadProcByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procServiceClient) UpdateProcByID(ctx context.Context, in *RequestForUpdateProcByID, opts ...grpc.CallOption) (*ProcID, error) {
	out := new(ProcID)
	err := c.cc.Invoke(ctx, "/ProcService/UpdateProcByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procServiceClient) DeleteProcByID(ctx context.Context, in *RequestForDeleteByID, opts ...grpc.CallOption) (*ProcID, error) {
	out := new(ProcID)
	err := c.cc.Invoke(ctx, "/ProcService/DeleteProcByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *procServiceClient) ReadAllProcs(ctx context.Context, in *RequestForReadAllProcs, opts ...grpc.CallOption) (*ProcList, error) {
	out := new(ProcList)
	err := c.cc.Invoke(ctx, "/ProcService/ReadAllProcs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcServiceServer is the server API for ProcService service.
// All implementations should embed UnimplementedProcServiceServer
// for forward compatibility
type ProcServiceServer interface {
	CreateProc(context.Context, *RequestForCreateProc) (*ProcID, error)
	ReadProcByID(context.Context, *RequestForReadByID) (*Proc, error)
	UpdateProcByID(context.Context, *RequestForUpdateProcByID) (*ProcID, error)
	DeleteProcByID(context.Context, *RequestForDeleteByID) (*ProcID, error)
	ReadAllProcs(context.Context, *RequestForReadAllProcs) (*ProcList, error)
}

// UnimplementedProcServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProcServiceServer struct {
}

func (*UnimplementedProcServiceServer) CreateProc(context.Context, *RequestForCreateProc) (*ProcID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProc not implemented")
}
func (*UnimplementedProcServiceServer) ReadProcByID(context.Context, *RequestForReadByID) (*Proc, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProcByID not implemented")
}
func (*UnimplementedProcServiceServer) UpdateProcByID(context.Context, *RequestForUpdateProcByID) (*ProcID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProcByID not implemented")
}
func (*UnimplementedProcServiceServer) DeleteProcByID(context.Context, *RequestForDeleteByID) (*ProcID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProcByID not implemented")
}
func (*UnimplementedProcServiceServer) ReadAllProcs(context.Context, *RequestForReadAllProcs) (*ProcList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllProcs not implemented")
}

func RegisterProcServiceServer(s *grpc.Server, srv ProcServiceServer) {
	s.RegisterService(&_ProcService_serviceDesc, srv)
}

func _ProcService_CreateProc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForCreateProc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcServiceServer).CreateProc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcService/CreateProc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcServiceServer).CreateProc(ctx, req.(*RequestForCreateProc))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcService_ReadProcByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForReadByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcServiceServer).ReadProcByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcService/ReadProcByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcServiceServer).ReadProcByID(ctx, req.(*RequestForReadByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcService_UpdateProcByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForUpdateProcByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcServiceServer).UpdateProcByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcService/UpdateProcByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcServiceServer).UpdateProcByID(ctx, req.(*RequestForUpdateProcByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcService_DeleteProcByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForDeleteByID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcServiceServer).DeleteProcByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcService/DeleteProcByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcServiceServer).DeleteProcByID(ctx, req.(*RequestForDeleteByID))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcService_ReadAllProcs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestForReadAllProcs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcServiceServer).ReadAllProcs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ProcService/ReadAllProcs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcServiceServer).ReadAllProcs(ctx, req.(*RequestForReadAllProcs))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ProcService",
	HandlerType: (*ProcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProc",
			Handler:    _ProcService_CreateProc_Handler,
		},
		{
			MethodName: "ReadProcByID",
			Handler:    _ProcService_ReadProcByID_Handler,
		},
		{
			MethodName: "UpdateProcByID",
			Handler:    _ProcService_UpdateProcByID_Handler,
		},
		{
			MethodName: "DeleteProcByID",
			Handler:    _ProcService_DeleteProcByID_Handler,
		},
		{
			MethodName: "ReadAllProcs",
			Handler:    _ProcService_ReadAllProcs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "process.proto",
}
