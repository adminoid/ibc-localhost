// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/legacy/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryRecordsRequest is request type for the Query/Records RPC method.
type QueryRecordsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRecordsRequest) Reset()         { *m = QueryRecordsRequest{} }
func (m *QueryRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRecordsRequest) ProtoMessage()    {}
func (*QueryRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cc45feb082d5d7c, []int{0}
}
func (m *QueryRecordsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRecordsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordsRequest.Merge(m, src)
}
func (m *QueryRecordsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordsRequest proto.InternalMessageInfo

func (m *QueryRecordsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryRecordsResponse is response type for the Query/Records RPC method.
type QueryRecordsResponse struct {
	Records    []Record            `protobuf:"bytes,1,rep,name=records,proto3" json:"records"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRecordsResponse) Reset()         { *m = QueryRecordsResponse{} }
func (m *QueryRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRecordsResponse) ProtoMessage()    {}
func (*QueryRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cc45feb082d5d7c, []int{1}
}
func (m *QueryRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordsResponse.Merge(m, src)
}
func (m *QueryRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordsResponse proto.InternalMessageInfo

func (m *QueryRecordsResponse) GetRecords() []Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *QueryRecordsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryRecordRequest is request type for the Query/Record RPC method.
type QueryRecordRequest struct {
	// legacy_address defines legacy account address to found out the legacy record.
	LegacyAddress string `protobuf:"bytes,1,opt,name=legacy_address,json=legacyAddress,proto3" json:"legacy_address,omitempty"`
}

func (m *QueryRecordRequest) Reset()         { *m = QueryRecordRequest{} }
func (m *QueryRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRecordRequest) ProtoMessage()    {}
func (*QueryRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cc45feb082d5d7c, []int{2}
}
func (m *QueryRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordRequest.Merge(m, src)
}
func (m *QueryRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordRequest proto.InternalMessageInfo

func (m *QueryRecordRequest) GetLegacyAddress() string {
	if m != nil {
		return m.LegacyAddress
	}
	return ""
}

// QueryRecordResponse is response type for the Query/Record RPC method.
type QueryRecordResponse struct {
	// record defines legacy record found by the request.
	Record Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record"`
}

func (m *QueryRecordResponse) Reset()         { *m = QueryRecordResponse{} }
func (m *QueryRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRecordResponse) ProtoMessage()    {}
func (*QueryRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cc45feb082d5d7c, []int{3}
}
func (m *QueryRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRecordResponse.Merge(m, src)
}
func (m *QueryRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRecordResponse proto.InternalMessageInfo

func (m *QueryRecordResponse) GetRecord() Record {
	if m != nil {
		return m.Record
	}
	return Record{}
}

// QueryCheckRequest is request type for the Query/Check RPC method.
type QueryCheckRequest struct {
	// pubkey defines account public key as the proof of legacy address authority.
	Pubkey []byte `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
}

func (m *QueryCheckRequest) Reset()         { *m = QueryCheckRequest{} }
func (m *QueryCheckRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCheckRequest) ProtoMessage()    {}
func (*QueryCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cc45feb082d5d7c, []int{4}
}
func (m *QueryCheckRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCheckRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCheckRequest.Merge(m, src)
}
func (m *QueryCheckRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCheckRequest proto.InternalMessageInfo

func (m *QueryCheckRequest) GetPubkey() []byte {
	if m != nil {
		return m.Pubkey
	}
	return nil
}

// QueryCheckResponse is response type for the Query/Check RPC method.
type QueryCheckResponse struct {
	// record defines legacy record found by the request.
	Record Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record"`
}

func (m *QueryCheckResponse) Reset()         { *m = QueryCheckResponse{} }
func (m *QueryCheckResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCheckResponse) ProtoMessage()    {}
func (*QueryCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7cc45feb082d5d7c, []int{5}
}
func (m *QueryCheckResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCheckResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCheckResponse.Merge(m, src)
}
func (m *QueryCheckResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCheckResponse proto.InternalMessageInfo

func (m *QueryCheckResponse) GetRecord() Record {
	if m != nil {
		return m.Record
	}
	return Record{}
}

func init() {
	proto.RegisterType((*QueryRecordsRequest)(nil), "decimal.legacy.v1.QueryRecordsRequest")
	proto.RegisterType((*QueryRecordsResponse)(nil), "decimal.legacy.v1.QueryRecordsResponse")
	proto.RegisterType((*QueryRecordRequest)(nil), "decimal.legacy.v1.QueryRecordRequest")
	proto.RegisterType((*QueryRecordResponse)(nil), "decimal.legacy.v1.QueryRecordResponse")
	proto.RegisterType((*QueryCheckRequest)(nil), "decimal.legacy.v1.QueryCheckRequest")
	proto.RegisterType((*QueryCheckResponse)(nil), "decimal.legacy.v1.QueryCheckResponse")
}

func init() { proto.RegisterFile("decimal/legacy/v1/query.proto", fileDescriptor_7cc45feb082d5d7c) }

var fileDescriptor_7cc45feb082d5d7c = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0xcf, 0x46, 0x9b, 0xe2, 0x54, 0x85, 0x8e, 0x41, 0x92, 0xa8, 0x6b, 0x5d, 0xda, 0xb4, 0x54,
	0xb2, 0x43, 0x22, 0x28, 0x9e, 0xc4, 0x08, 0x7a, 0x52, 0xea, 0x8a, 0x17, 0x41, 0xca, 0xec, 0xee,
	0x30, 0x2e, 0x49, 0x76, 0xb6, 0x3b, 0x93, 0x68, 0x28, 0xbd, 0xe8, 0x17, 0x10, 0xbc, 0xf9, 0x39,
	0xfc, 0x10, 0x3d, 0x16, 0xbd, 0x78, 0x12, 0x49, 0xfc, 0x1c, 0x22, 0x99, 0x79, 0xab, 0xbb, 0xa4,
	0xba, 0x07, 0x6f, 0x3b, 0xf3, 0x7e, 0xef, 0xf7, 0xe7, 0xbd, 0x59, 0x74, 0x2d, 0x64, 0x41, 0x34,
	0xa2, 0x43, 0x32, 0x64, 0x9c, 0x06, 0x53, 0x32, 0xe9, 0x92, 0x83, 0x31, 0x4b, 0xa7, 0x6e, 0x92,
	0x0a, 0x25, 0xf0, 0x3a, 0x94, 0x5d, 0x53, 0x76, 0x27, 0xdd, 0xd6, 0x55, 0x2e, 0x04, 0x1f, 0x32,
	0x42, 0x93, 0x88, 0xd0, 0x38, 0x16, 0x8a, 0xaa, 0x48, 0xc4, 0xd2, 0x34, 0xb4, 0xea, 0x5c, 0x70,
	0xa1, 0x3f, 0xc9, 0xe2, 0x0b, 0x6e, 0x9b, 0x81, 0x90, 0x23, 0x21, 0xf7, 0x4d, 0xc1, 0x1c, 0xa0,
	0xb4, 0x6b, 0x4e, 0xc4, 0xa7, 0x92, 0x19, 0x69, 0x32, 0xe9, 0xfa, 0x4c, 0xd1, 0x2e, 0x49, 0x28,
	0x8f, 0x62, 0xcd, 0x0e, 0x58, 0x7b, 0xd9, 0x2c, 0xf8, 0xd2, 0x75, 0xe7, 0x25, 0xba, 0xf4, 0x74,
	0xc1, 0xe0, 0xb1, 0x40, 0xa4, 0xa1, 0xf4, 0xd8, 0xc1, 0x98, 0x49, 0x85, 0x1f, 0x22, 0xf4, 0x87,
	0xaa, 0x61, 0x6d, 0x58, 0x3b, 0x6b, 0xbd, 0xb6, 0x0b, 0x2e, 0x16, 0xba, 0xae, 0x89, 0x0c, 0xba,
	0xee, 0x1e, 0xe5, 0x0c, 0x7a, 0xbd, 0x5c, 0xa7, 0xf3, 0xd1, 0x42, 0xf5, 0x22, 0xbf, 0x4c, 0x44,
	0x2c, 0x19, 0xbe, 0x8b, 0x56, 0x53, 0x73, 0xd5, 0xb0, 0x36, 0xce, 0xec, 0xac, 0xf5, 0x9a, 0xee,
	0xd2, 0xdc, 0x5c, 0xd3, 0xd4, 0x3f, 0x7b, 0xfc, 0xed, 0x7a, 0xc5, 0xcb, 0xf0, 0xf8, 0x51, 0xc1,
	0x5b, 0x55, 0x7b, 0xdb, 0x2e, 0xf5, 0x66, 0x74, 0x0b, 0xe6, 0x9e, 0x23, 0x9c, 0xf3, 0x96, 0x45,
	0xbf, 0x87, 0x2e, 0x1a, 0x07, 0xfb, 0x34, 0x0c, 0x53, 0x26, 0xa5, 0x8e, 0x7f, 0xae, 0xdf, 0xf8,
	0xfc, 0xa9, 0x53, 0x07, 0x95, 0xfb, 0xa6, 0xf2, 0x4c, 0xa5, 0x51, 0xcc, 0xbd, 0x0b, 0x06, 0x0f,
	0x97, 0xce, 0x93, 0xc2, 0x48, 0x7f, 0x27, 0xbe, 0x83, 0x6a, 0x26, 0x01, 0x8c, 0xb3, 0x34, 0x30,
	0xc0, 0x9d, 0x9b, 0x68, 0x5d, 0xf3, 0x3d, 0x78, 0xc5, 0x82, 0x41, 0xe6, 0xf2, 0x32, 0xaa, 0x25,
	0x63, 0x7f, 0xc0, 0xa6, 0x9a, 0xed, 0xbc, 0x07, 0x27, 0xe7, 0x31, 0x64, 0x02, 0xf0, 0x7f, 0x6a,
	0xf7, 0x7e, 0x56, 0xd1, 0x8a, 0xe6, 0xc3, 0x13, 0xb4, 0x0a, 0x3b, 0xc4, 0xed, 0x53, 0xba, 0x4f,
	0x79, 0x44, 0xad, 0xed, 0x52, 0x9c, 0xb1, 0xe7, 0xb4, 0xde, 0x7e, 0xf9, 0xf1, 0xa1, 0x5a, 0xc7,
	0x38, 0xf7, 0x4a, 0xb3, 0x6d, 0xbf, 0xb3, 0x50, 0xcd, 0xe0, 0xf1, 0xd6, 0xbf, 0xf9, 0x32, 0xd9,
	0x76, 0x19, 0x0c, 0x54, 0x77, 0xb5, 0xea, 0x26, 0x76, 0x96, 0x54, 0xc9, 0x61, 0xf1, 0x05, 0x1c,
	0xe1, 0xd7, 0x68, 0x45, 0x4f, 0x14, 0x6f, 0xfe, 0x8d, 0x3c, 0xbf, 0x9d, 0xd6, 0x56, 0x09, 0x0a,
	0x1c, 0xdc, 0xd0, 0x0e, 0xae, 0xe0, 0x66, 0xce, 0x41, 0xb0, 0x40, 0x90, 0x43, 0xb3, 0xce, 0xa3,
	0xfe, 0xde, 0xf1, 0xcc, 0xb6, 0x4e, 0x66, 0xb6, 0xf5, 0x7d, 0x66, 0x5b, 0xef, 0xe7, 0x76, 0xe5,
	0x64, 0x6e, 0x57, 0xbe, 0xce, 0xed, 0xca, 0x8b, 0xdb, 0x7e, 0xa4, 0xfc, 0x71, 0x30, 0x60, 0xca,
	0x15, 0x29, 0x27, 0x20, 0xa8, 0x18, 0x1d, 0x11, 0x2e, 0x3a, 0x72, 0x44, 0x53, 0xd5, 0x89, 0x45,
	0xc8, 0xc8, 0x9b, 0x8c, 0x5d, 0x4d, 0x13, 0x26, 0xfd, 0x9a, 0xfe, 0xf1, 0x6f, 0xfd, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0xac, 0xf1, 0xf7, 0x58, 0xc7, 0x04, 0x00, 0x00,
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
	// Records queries all legacy records that should be returned to the actual owners.
	Records(ctx context.Context, in *QueryRecordsRequest, opts ...grpc.CallOption) (*QueryRecordsResponse, error)
	// Record queries complete set of different values that should be returned to the actual owner.
	Record(ctx context.Context, in *QueryRecordRequest, opts ...grpc.CallOption) (*QueryRecordResponse, error)
	// Check queries legacy record by specifiec public key.
	Check(ctx context.Context, in *QueryCheckRequest, opts ...grpc.CallOption) (*QueryCheckResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Records(ctx context.Context, in *QueryRecordsRequest, opts ...grpc.CallOption) (*QueryRecordsResponse, error) {
	out := new(QueryRecordsResponse)
	err := c.cc.Invoke(ctx, "/decimal.legacy.v1.Query/Records", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Record(ctx context.Context, in *QueryRecordRequest, opts ...grpc.CallOption) (*QueryRecordResponse, error) {
	out := new(QueryRecordResponse)
	err := c.cc.Invoke(ctx, "/decimal.legacy.v1.Query/Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Check(ctx context.Context, in *QueryCheckRequest, opts ...grpc.CallOption) (*QueryCheckResponse, error) {
	out := new(QueryCheckResponse)
	err := c.cc.Invoke(ctx, "/decimal.legacy.v1.Query/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Records queries all legacy records that should be returned to the actual owners.
	Records(context.Context, *QueryRecordsRequest) (*QueryRecordsResponse, error)
	// Record queries complete set of different values that should be returned to the actual owner.
	Record(context.Context, *QueryRecordRequest) (*QueryRecordResponse, error)
	// Check queries legacy record by specifiec public key.
	Check(context.Context, *QueryCheckRequest) (*QueryCheckResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Records(ctx context.Context, req *QueryRecordsRequest) (*QueryRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Records not implemented")
}
func (*UnimplementedQueryServer) Record(ctx context.Context, req *QueryRecordRequest) (*QueryRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (*UnimplementedQueryServer) Check(ctx context.Context, req *QueryCheckRequest) (*QueryCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Records_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Records(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.legacy.v1.Query/Records",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Records(ctx, req.(*QueryRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.legacy.v1.Query/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Record(ctx, req.(*QueryRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.legacy.v1.Query/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Check(ctx, req.(*QueryCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decimal.legacy.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Records",
			Handler:    _Query_Records_Handler,
		},
		{
			MethodName: "Record",
			Handler:    _Query_Record_Handler,
		},
		{
			MethodName: "Check",
			Handler:    _Query_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decimal/legacy/v1/query.proto",
}

func (m *QueryRecordsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRecordsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRecordsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LegacyAddress) > 0 {
		i -= len(m.LegacyAddress)
		copy(dAtA[i:], m.LegacyAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.LegacyAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryCheckRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCheckRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCheckRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pubkey) > 0 {
		i -= len(m.Pubkey)
		copy(dAtA[i:], m.Pubkey)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Pubkey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCheckResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCheckResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCheckResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryRecordsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.LegacyAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Record.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryCheckRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Pubkey)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryCheckResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Record.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRecordsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRecordsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRecordsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
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
			m.Records = append(m.Records, Record{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryRecordRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LegacyAddress", wireType)
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
			m.LegacyAddress = string(dAtA[iNdEx:postIndex])
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
func (m *QueryRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
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
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryCheckRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCheckRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCheckRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pubkey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pubkey = append(m.Pubkey[:0], dAtA[iNdEx:postIndex]...)
			if m.Pubkey == nil {
				m.Pubkey = []byte{}
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
func (m *QueryCheckResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryCheckResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCheckResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
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
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
