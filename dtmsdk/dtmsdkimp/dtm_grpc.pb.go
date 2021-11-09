// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dtmsdkimp

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DtmSvcClient is the client API for DtmSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DtmSvcClient interface {
	// rpc RegisterBranch(DtmBranchRequest)  returns (google.protobuf.Empty) {}
	Submit(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type dtmSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewDtmSvcClient(cc grpc.ClientConnInterface) DtmSvcClient {
	return &dtmSvcClient{cc}
}

func (c *dtmSvcClient) Submit(ctx context.Context, in *DtmRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/dtmsdkimp.DtmSvc/Submit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DtmSvcServer is the server API for DtmSvc service.
// All implementations must embed UnimplementedDtmSvcServer
// for forward compatibility
type DtmSvcServer interface {
	// rpc RegisterBranch(DtmBranchRequest)  returns (google.protobuf.Empty) {}
	Submit(context.Context, *DtmRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDtmSvcServer()
}

// UnimplementedDtmSvcServer must be embedded to have forward compatible implementations.
type UnimplementedDtmSvcServer struct {
}

func (UnimplementedDtmSvcServer) Submit(context.Context, *DtmRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Submit not implemented")
}
func (UnimplementedDtmSvcServer) mustEmbedUnimplementedDtmSvcServer() {}

// UnsafeDtmSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DtmSvcServer will
// result in compilation errors.
type UnsafeDtmSvcServer interface {
	mustEmbedUnimplementedDtmSvcServer()
}

func RegisterDtmSvcServer(s grpc.ServiceRegistrar, srv DtmSvcServer) {
	s.RegisterService(&DtmSvc_ServiceDesc, srv)
}

func _DtmSvc_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DtmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DtmSvcServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dtmsdkimp.DtmSvc/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DtmSvcServer).Submit(ctx, req.(*DtmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DtmSvc_ServiceDesc is the grpc.ServiceDesc for DtmSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DtmSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dtmsdkimp.DtmSvc",
	HandlerType: (*DtmSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _DtmSvc_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dtmsdk/dtmsdkimp/dtm.proto",
}
