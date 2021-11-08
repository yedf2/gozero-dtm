// Code generated by protoc-gen-go. DO NOT EDIT.
// source: trans.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AdjustInfo struct {
	Amount               int64    `protobuf:"varint,1,opt,name=Amount,proto3" json:"Amount,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdjustInfo) Reset()         { *m = AdjustInfo{} }
func (m *AdjustInfo) String() string { return proto.CompactTextString(m) }
func (*AdjustInfo) ProtoMessage()    {}
func (*AdjustInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbc80efa276a497, []int{0}
}

func (m *AdjustInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdjustInfo.Unmarshal(m, b)
}
func (m *AdjustInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdjustInfo.Marshal(b, m, deterministic)
}
func (m *AdjustInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdjustInfo.Merge(m, src)
}
func (m *AdjustInfo) XXX_Size() int {
	return xxx_messageInfo_AdjustInfo.Size(m)
}
func (m *AdjustInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AdjustInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AdjustInfo proto.InternalMessageInfo

func (m *AdjustInfo) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *AdjustInfo) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_afbc80efa276a497, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdjustInfo)(nil), "trans.AdjustInfo")
	proto.RegisterType((*Response)(nil), "trans.Response")
}

func init() { proto.RegisterFile("trans.proto", fileDescriptor_afbc80efa276a497) }

var fileDescriptor_afbc80efa276a497 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x29, 0x4a, 0xcc,
	0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x6c, 0xb8, 0xb8, 0x1c,
	0x53, 0xb2, 0x4a, 0x8b, 0x4b, 0x3c, 0xf3, 0xd2, 0xf2, 0x85, 0xc4, 0xb8, 0xd8, 0x1c, 0x73, 0xf3,
	0x4b, 0xf3, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98, 0x83, 0xa0, 0x3c, 0x90, 0x78, 0x68, 0x71,
	0x6a, 0x91, 0xa7, 0x8b, 0x04, 0x13, 0x44, 0x1c, 0xc2, 0x53, 0xe2, 0xe2, 0xe2, 0x08, 0x4a, 0x2d,
	0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0xca, 0xe4, 0xe2, 0x08, 0x01, 0x19, 0x19, 0x5c, 0x96, 0x2c,
	0xa4, 0xcb, 0xc5, 0x0e, 0x66, 0x7b, 0xe6, 0x09, 0x09, 0xea, 0x41, 0x6c, 0x45, 0xd8, 0x22, 0xc5,
	0x0f, 0x15, 0x82, 0x69, 0x15, 0xd2, 0x83, 0x6a, 0xf5, 0x2f, 0x2d, 0x21, 0x46, 0xbd, 0x13, 0x4b,
	0x14, 0x53, 0x41, 0x52, 0x12, 0x1b, 0xd8, 0x23, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0xf2, 0xd6, 0xcb, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransSvcClient is the client API for TransSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransSvcClient interface {
	TransIn(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error)
	TransOut(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error)
}

type transSvcClient struct {
	cc *grpc.ClientConn
}

func NewTransSvcClient(cc *grpc.ClientConn) TransSvcClient {
	return &transSvcClient{cc}
}

func (c *transSvcClient) TransIn(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/trans.TransSvc/TransIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transSvcClient) TransOut(ctx context.Context, in *AdjustInfo, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/trans.TransSvc/TransOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransSvcServer is the server API for TransSvc service.
type TransSvcServer interface {
	TransIn(context.Context, *AdjustInfo) (*Response, error)
	TransOut(context.Context, *AdjustInfo) (*Response, error)
}

// UnimplementedTransSvcServer can be embedded to have forward compatible implementations.
type UnimplementedTransSvcServer struct {
}

func (*UnimplementedTransSvcServer) TransIn(ctx context.Context, req *AdjustInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransIn not implemented")
}
func (*UnimplementedTransSvcServer) TransOut(ctx context.Context, req *AdjustInfo) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOut not implemented")
}

func RegisterTransSvcServer(s *grpc.Server, srv TransSvcServer) {
	s.RegisterService(&_TransSvc_serviceDesc, srv)
}

func _TransSvc_TransIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransSvcServer).TransIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trans.TransSvc/TransIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransSvcServer).TransIn(ctx, req.(*AdjustInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransSvc_TransOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdjustInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransSvcServer).TransOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trans.TransSvc/TransOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransSvcServer).TransOut(ctx, req.(*AdjustInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trans.TransSvc",
	HandlerType: (*TransSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransIn",
			Handler:    _TransSvc_TransIn_Handler,
		},
		{
			MethodName: "TransOut",
			Handler:    _TransSvc_TransOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trans.proto",
}