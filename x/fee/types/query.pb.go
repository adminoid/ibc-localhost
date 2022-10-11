// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/fee/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryCoinPricesRequest is request type for the Query/CoinPrices RPC method.
type QueryCoinPricesRequest struct {
}

func (m *QueryCoinPricesRequest) Reset()         { *m = QueryCoinPricesRequest{} }
func (m *QueryCoinPricesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCoinPricesRequest) ProtoMessage()    {}
func (*QueryCoinPricesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_451d183aec7ef1fe, []int{0}
}
func (m *QueryCoinPricesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoinPricesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoinPricesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoinPricesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoinPricesRequest.Merge(m, src)
}
func (m *QueryCoinPricesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoinPricesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoinPricesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoinPricesRequest proto.InternalMessageInfo

// QueryCoinPricesResponse is response type for the Query/CoinPrices RPC method.
type QueryCoinPricesResponse struct {
	Prices []CoinPrice `protobuf:"bytes,1,rep,name=prices,proto3" json:"prices"`
}

func (m *QueryCoinPricesResponse) Reset()         { *m = QueryCoinPricesResponse{} }
func (m *QueryCoinPricesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCoinPricesResponse) ProtoMessage()    {}
func (*QueryCoinPricesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_451d183aec7ef1fe, []int{1}
}
func (m *QueryCoinPricesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoinPricesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoinPricesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoinPricesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoinPricesResponse.Merge(m, src)
}
func (m *QueryCoinPricesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoinPricesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoinPricesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoinPricesResponse proto.InternalMessageInfo

func (m *QueryCoinPricesResponse) GetPrices() []CoinPrice {
	if m != nil {
		return m.Prices
	}
	return nil
}

// QueryCoinPriceRequest is request type for the Query/CoinPrice RPC method.
type QueryCoinPriceRequest struct {
	// denom defines the base currency (coin) denomination which is priced.
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// quote defines the quote currency denomination in the pair (USD as the first example).
	Quote string `protobuf:"bytes,2,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (m *QueryCoinPriceRequest) Reset()         { *m = QueryCoinPriceRequest{} }
func (m *QueryCoinPriceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCoinPriceRequest) ProtoMessage()    {}
func (*QueryCoinPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_451d183aec7ef1fe, []int{2}
}
func (m *QueryCoinPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoinPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoinPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoinPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoinPriceRequest.Merge(m, src)
}
func (m *QueryCoinPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoinPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoinPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoinPriceRequest proto.InternalMessageInfo

func (m *QueryCoinPriceRequest) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *QueryCoinPriceRequest) GetQuote() string {
	if m != nil {
		return m.Quote
	}
	return ""
}

// QueryCoinPriceResponse is response type for the Query/CoinPrice RPC method.
type QueryCoinPriceResponse struct {
	Price *CoinPrice `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
}

func (m *QueryCoinPriceResponse) Reset()         { *m = QueryCoinPriceResponse{} }
func (m *QueryCoinPriceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCoinPriceResponse) ProtoMessage()    {}
func (*QueryCoinPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_451d183aec7ef1fe, []int{3}
}
func (m *QueryCoinPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoinPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoinPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoinPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoinPriceResponse.Merge(m, src)
}
func (m *QueryCoinPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoinPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoinPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoinPriceResponse proto.InternalMessageInfo

func (m *QueryCoinPriceResponse) GetPrice() *CoinPrice {
	if m != nil {
		return m.Price
	}
	return nil
}

// QueryModuleParamsRequest is request type for the Query/ModuleParams RPC method.
type QueryModuleParamsRequest struct {
}

func (m *QueryModuleParamsRequest) Reset()         { *m = QueryModuleParamsRequest{} }
func (m *QueryModuleParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryModuleParamsRequest) ProtoMessage()    {}
func (*QueryModuleParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_451d183aec7ef1fe, []int{4}
}
func (m *QueryModuleParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryModuleParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryModuleParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryModuleParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryModuleParamsRequest.Merge(m, src)
}
func (m *QueryModuleParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryModuleParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryModuleParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryModuleParamsRequest proto.InternalMessageInfo

// QueryModuleParamsResponse is response type for the Query/ModuleParams RPC method.
type QueryModuleParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryModuleParamsResponse) Reset()         { *m = QueryModuleParamsResponse{} }
func (m *QueryModuleParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryModuleParamsResponse) ProtoMessage()    {}
func (*QueryModuleParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_451d183aec7ef1fe, []int{5}
}
func (m *QueryModuleParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryModuleParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryModuleParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryModuleParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryModuleParamsResponse.Merge(m, src)
}
func (m *QueryModuleParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryModuleParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryModuleParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryModuleParamsResponse proto.InternalMessageInfo

func (m *QueryModuleParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryCoinPricesRequest)(nil), "decimal.fee.v1.QueryCoinPricesRequest")
	proto.RegisterType((*QueryCoinPricesResponse)(nil), "decimal.fee.v1.QueryCoinPricesResponse")
	proto.RegisterType((*QueryCoinPriceRequest)(nil), "decimal.fee.v1.QueryCoinPriceRequest")
	proto.RegisterType((*QueryCoinPriceResponse)(nil), "decimal.fee.v1.QueryCoinPriceResponse")
	proto.RegisterType((*QueryModuleParamsRequest)(nil), "decimal.fee.v1.QueryModuleParamsRequest")
	proto.RegisterType((*QueryModuleParamsResponse)(nil), "decimal.fee.v1.QueryModuleParamsResponse")
}

func init() { proto.RegisterFile("decimal/fee/v1/query.proto", fileDescriptor_451d183aec7ef1fe) }

var fileDescriptor_451d183aec7ef1fe = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xb5, 0x24, 0x52, 0x0f, 0xc4, 0x70, 0x2a, 0xc1, 0x35, 0xc8, 0x54, 0x96, 0xda, 0x86,
	0x21, 0x3e, 0xb5, 0x45, 0x62, 0xa7, 0x13, 0x43, 0xa5, 0xd6, 0x23, 0xdb, 0x25, 0x7e, 0xb5, 0x2c,
	0x62, 0x3f, 0xc7, 0x3e, 0x57, 0x54, 0x55, 0x16, 0x98, 0xd8, 0x90, 0x58, 0xf8, 0x49, 0x1d, 0x23,
	0xb1, 0x30, 0x21, 0x94, 0xf0, 0x43, 0x90, 0xef, 0x2e, 0x86, 0x18, 0x2b, 0xe9, 0xe6, 0x7b, 0xdf,
	0xf7, 0xbe, 0xef, 0xbd, 0xf7, 0xc9, 0xd4, 0x09, 0x61, 0x14, 0x27, 0x62, 0xcc, 0xaf, 0x00, 0xf8,
	0xf5, 0x31, 0x9f, 0x94, 0x90, 0xdf, 0xf8, 0x59, 0x8e, 0x12, 0xd9, 0x63, 0x83, 0xf9, 0x57, 0x00,
	0xfe, 0xf5, 0xb1, 0xb3, 0x1b, 0x61, 0x84, 0x0a, 0xe2, 0xd5, 0x97, 0x66, 0x39, 0xcf, 0x23, 0xc4,
	0x68, 0x0c, 0x5c, 0x64, 0x31, 0x17, 0x69, 0x8a, 0x52, 0xc8, 0x18, 0xd3, 0xc2, 0xa0, 0x76, 0x43,
	0xbf, 0x92, 0xd2, 0xc8, 0xb3, 0x06, 0x92, 0x89, 0x5c, 0x24, 0xa6, 0xcd, 0xb3, 0x69, 0xef, 0xb2,
	0x9a, 0xe4, 0x0c, 0xe3, 0xf4, 0x22, 0x8f, 0x47, 0x50, 0x04, 0x30, 0x29, 0xa1, 0x90, 0x5e, 0x40,
	0x9f, 0xfe, 0x87, 0x14, 0x19, 0xa6, 0x05, 0xb0, 0xd7, 0xb4, 0x9b, 0xa9, 0x8a, 0x4d, 0xf6, 0xb7,
	0xfb, 0x0f, 0x4f, 0xf6, 0xfc, 0xd5, 0x05, 0xfc, 0xba, 0xe7, 0xcd, 0x83, 0xbb, 0x9f, 0x2f, 0xac,
	0xc0, 0xd0, 0xbd, 0x33, 0xfa, 0x64, 0x55, 0xd3, 0x98, 0xb1, 0x5d, 0xda, 0x09, 0x21, 0xc5, 0xc4,
	0x26, 0xfb, 0xa4, 0xbf, 0x13, 0xe8, 0x47, 0x55, 0x9d, 0x94, 0x28, 0xc1, 0xde, 0xd2, 0x55, 0xf5,
	0xf0, 0xde, 0x36, 0x47, 0xae, 0xe7, 0xe2, 0xb4, 0xa3, 0x8c, 0x94, 0xca, 0xba, 0xb1, 0x02, 0xcd,
	0xf3, 0x1c, 0x6a, 0x2b, 0xa9, 0x73, 0x0c, 0xcb, 0x31, 0x5c, 0xa8, 0xc3, 0x2c, 0xf7, 0xbf, 0xa4,
	0x7b, 0x2d, 0x98, 0x71, 0x7a, 0x45, 0xbb, 0xfa, 0x8c, 0xc6, 0xaa, 0xd7, 0xb4, 0xd2, 0xfc, 0x7a,
	0x7d, 0xf5, 0x3a, 0xf9, 0xb6, 0x4d, 0x3b, 0x4a, 0x93, 0x4d, 0x29, 0xfd, 0x7b, 0x57, 0x76, 0xd8,
	0xec, 0x6e, 0x8f, 0xc4, 0x39, 0xda, 0xc8, 0xd3, 0xe3, 0x79, 0xee, 0xc7, 0xef, 0xbf, 0xbf, 0x6e,
	0xd9, 0xac, 0xc7, 0x9b, 0xd9, 0x6b, 0xc3, 0xcf, 0x84, 0xee, 0xd4, 0x6d, 0xec, 0x60, 0xbd, 0xec,
	0xd2, 0xfd, 0x70, 0x13, 0xcd, 0x98, 0x0f, 0x94, 0xf9, 0x11, 0x3b, 0x68, 0x35, 0xe7, 0xb7, 0x2a,
	0xdb, 0x29, 0xbf, 0x55, 0x69, 0x4e, 0xd9, 0x27, 0x42, 0x1f, 0xfd, 0x7b, 0x63, 0xd6, 0x6f, 0xf5,
	0x69, 0x89, 0xc8, 0x79, 0x79, 0x0f, 0xe6, 0xc6, 0x8b, 0xe8, 0xa0, 0xce, 0xef, 0xe6, 0x2e, 0x99,
	0xcd, 0x5d, 0xf2, 0x6b, 0xee, 0x92, 0x2f, 0x0b, 0xd7, 0x9a, 0x2d, 0x5c, 0xeb, 0xc7, 0xc2, 0xb5,
	0xde, 0x9d, 0x0e, 0x63, 0x39, 0x2c, 0x47, 0xef, 0x41, 0xfa, 0x98, 0x47, 0xcb, 0x76, 0x09, 0x22,
	0xe1, 0x11, 0x0e, 0x8a, 0x44, 0xe4, 0x72, 0x90, 0x62, 0x08, 0xfc, 0x83, 0x92, 0x94, 0x37, 0x19,
	0x14, 0xc3, 0xae, 0xfa, 0xbb, 0x4e, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x79, 0x1c, 0x5e, 0xdb,
	0xf6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// CoinPrices queries all available coin prices.
	CoinPrices(ctx context.Context, in *QueryCoinPricesRequest, opts ...grpc.CallOption) (*QueryCoinPricesResponse, error)
	// CoinPrice queries the specified coin price.
	CoinPrice(ctx context.Context, in *QueryCoinPriceRequest, opts ...grpc.CallOption) (*QueryCoinPriceResponse, error)
	// ModuleParams queries the module params.
	ModuleParams(ctx context.Context, in *QueryModuleParamsRequest, opts ...grpc.CallOption) (*QueryModuleParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CoinPrices(ctx context.Context, in *QueryCoinPricesRequest, opts ...grpc.CallOption) (*QueryCoinPricesResponse, error) {
	out := new(QueryCoinPricesResponse)
	err := c.cc.Invoke(ctx, "/decimal.fee.v1.Query/CoinPrices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CoinPrice(ctx context.Context, in *QueryCoinPriceRequest, opts ...grpc.CallOption) (*QueryCoinPriceResponse, error) {
	out := new(QueryCoinPriceResponse)
	err := c.cc.Invoke(ctx, "/decimal.fee.v1.Query/CoinPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ModuleParams(ctx context.Context, in *QueryModuleParamsRequest, opts ...grpc.CallOption) (*QueryModuleParamsResponse, error) {
	out := new(QueryModuleParamsResponse)
	err := c.cc.Invoke(ctx, "/decimal.fee.v1.Query/ModuleParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CoinPrices queries all available coin prices.
	CoinPrices(context.Context, *QueryCoinPricesRequest) (*QueryCoinPricesResponse, error)
	// CoinPrice queries the specified coin price.
	CoinPrice(context.Context, *QueryCoinPriceRequest) (*QueryCoinPriceResponse, error)
	// ModuleParams queries the module params.
	ModuleParams(context.Context, *QueryModuleParamsRequest) (*QueryModuleParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CoinPrices(ctx context.Context, req *QueryCoinPricesRequest) (*QueryCoinPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinPrices not implemented")
}
func (*UnimplementedQueryServer) CoinPrice(ctx context.Context, req *QueryCoinPriceRequest) (*QueryCoinPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinPrice not implemented")
}
func (*UnimplementedQueryServer) ModuleParams(ctx context.Context, req *QueryModuleParamsRequest) (*QueryModuleParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModuleParams not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CoinPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCoinPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.fee.v1.Query/CoinPrices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinPrices(ctx, req.(*QueryCoinPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CoinPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCoinPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.fee.v1.Query/CoinPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinPrice(ctx, req.(*QueryCoinPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ModuleParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryModuleParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ModuleParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.fee.v1.Query/ModuleParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ModuleParams(ctx, req.(*QueryModuleParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decimal.fee.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoinPrices",
			Handler:    _Query_CoinPrices_Handler,
		},
		{
			MethodName: "CoinPrice",
			Handler:    _Query_CoinPrice_Handler,
		},
		{
			MethodName: "ModuleParams",
			Handler:    _Query_ModuleParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decimal/fee/v1/query.proto",
}

func (m *QueryCoinPricesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoinPricesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoinPricesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryCoinPricesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoinPricesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoinPricesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryCoinPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoinPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoinPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Quote) > 0 {
		i -= len(m.Quote)
		copy(dAtA[i:], m.Quote)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Quote)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCoinPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoinPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoinPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Price != nil {
		{
			size, err := m.Price.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryModuleParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryModuleParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryModuleParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryModuleParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryModuleParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryModuleParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCoinPricesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryCoinPricesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryCoinPriceRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Quote)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCoinPriceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Price != nil {
		l = m.Price.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryModuleParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryModuleParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCoinPricesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCoinPricesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoinPricesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCoinPricesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCoinPricesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoinPricesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prices = append(m.Prices, CoinPrice{})
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCoinPriceRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCoinPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoinPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryCoinPriceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryCoinPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoinPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Price == nil {
				m.Price = &CoinPrice{}
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryModuleParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryModuleParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryModuleParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryModuleParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryModuleParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryModuleParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
