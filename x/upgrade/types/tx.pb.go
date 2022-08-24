// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/upgrade/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgSoftwareUpgrade is the Msg/SoftwareUpgrade request type.
type MsgSoftwareUpgrade struct {
	// authority is the address of the governance account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// plan is the upgrade plan.
	Plan Plan `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan"`
}

func (m *MsgSoftwareUpgrade) Reset()         { *m = MsgSoftwareUpgrade{} }
func (m *MsgSoftwareUpgrade) String() string { return proto.CompactTextString(m) }
func (*MsgSoftwareUpgrade) ProtoMessage()    {}
func (*MsgSoftwareUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4671ce387cbe8f4, []int{0}
}
func (m *MsgSoftwareUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSoftwareUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSoftwareUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSoftwareUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSoftwareUpgrade.Merge(m, src)
}
func (m *MsgSoftwareUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *MsgSoftwareUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSoftwareUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSoftwareUpgrade proto.InternalMessageInfo

func (m *MsgSoftwareUpgrade) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgSoftwareUpgrade) GetPlan() Plan {
	if m != nil {
		return m.Plan
	}
	return Plan{}
}

// MsgSoftwareUpgradeResponse is the Msg/SoftwareUpgrade response type.
type MsgSoftwareUpgradeResponse struct {
}

func (m *MsgSoftwareUpgradeResponse) Reset()         { *m = MsgSoftwareUpgradeResponse{} }
func (m *MsgSoftwareUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSoftwareUpgradeResponse) ProtoMessage()    {}
func (*MsgSoftwareUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4671ce387cbe8f4, []int{1}
}
func (m *MsgSoftwareUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSoftwareUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSoftwareUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSoftwareUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSoftwareUpgradeResponse.Merge(m, src)
}
func (m *MsgSoftwareUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSoftwareUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSoftwareUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSoftwareUpgradeResponse proto.InternalMessageInfo

// MsgCancelUpgrade is the Msg/CancelUpgrade request type.
type MsgCancelUpgrade struct {
	// authority is the address of the governance account.
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (m *MsgCancelUpgrade) Reset()         { *m = MsgCancelUpgrade{} }
func (m *MsgCancelUpgrade) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgrade) ProtoMessage()    {}
func (*MsgCancelUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4671ce387cbe8f4, []int{2}
}
func (m *MsgCancelUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgrade.Merge(m, src)
}
func (m *MsgCancelUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgrade proto.InternalMessageInfo

func (m *MsgCancelUpgrade) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

// MsgCancelUpgradeResponse is the Msg/CancelUpgrade response type.
type MsgCancelUpgradeResponse struct {
}

func (m *MsgCancelUpgradeResponse) Reset()         { *m = MsgCancelUpgradeResponse{} }
func (m *MsgCancelUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCancelUpgradeResponse) ProtoMessage()    {}
func (*MsgCancelUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4671ce387cbe8f4, []int{3}
}
func (m *MsgCancelUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCancelUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCancelUpgradeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCancelUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCancelUpgradeResponse.Merge(m, src)
}
func (m *MsgCancelUpgradeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCancelUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCancelUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCancelUpgradeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSoftwareUpgrade)(nil), "decimal.upgrade.v1.MsgSoftwareUpgrade")
	proto.RegisterType((*MsgSoftwareUpgradeResponse)(nil), "decimal.upgrade.v1.MsgSoftwareUpgradeResponse")
	proto.RegisterType((*MsgCancelUpgrade)(nil), "decimal.upgrade.v1.MsgCancelUpgrade")
	proto.RegisterType((*MsgCancelUpgradeResponse)(nil), "decimal.upgrade.v1.MsgCancelUpgradeResponse")
}

func init() { proto.RegisterFile("decimal/upgrade/v1/tx.proto", fileDescriptor_d4671ce387cbe8f4) }

var fileDescriptor_d4671ce387cbe8f4 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0x02, 0x41,
	0x18, 0xc6, 0x77, 0x4a, 0x02, 0x27, 0xfa, 0xc3, 0x22, 0xb4, 0x6d, 0xb1, 0x89, 0x44, 0x48, 0xe4,
	0x0e, 0x1a, 0x14, 0x74, 0xcb, 0xce, 0x42, 0x29, 0x5d, 0xbc, 0xc4, 0xb8, 0x3b, 0x4d, 0x4b, 0xbb,
	0x3b, 0xcb, 0xcc, 0x68, 0x7a, 0xed, 0x13, 0xf8, 0x51, 0x3a, 0xf4, 0x21, 0x3c, 0x4a, 0x10, 0x74,
	0x8a, 0xd0, 0x43, 0x5f, 0x23, 0xdc, 0x3f, 0x89, 0xae, 0x81, 0xd0, 0x6d, 0x5e, 0x9e, 0x87, 0xe7,
	0xf7, 0x3e, 0xc3, 0x0b, 0xf7, 0x6c, 0x62, 0x39, 0x1e, 0x76, 0x51, 0x3b, 0xa0, 0x1c, 0xdb, 0x04,
	0x75, 0xca, 0x48, 0x76, 0xcd, 0x80, 0x33, 0xc9, 0x54, 0x35, 0x16, 0xcd, 0x58, 0x34, 0x3b, 0x65,
	0x3d, 0x47, 0x19, 0x65, 0xa1, 0x8c, 0x26, 0xaf, 0xc8, 0xa9, 0xef, 0x5a, 0x4c, 0x78, 0x4c, 0xdc,
	0x45, 0x42, 0x34, 0xc4, 0xd2, 0x4e, 0x34, 0x21, 0x4f, 0xd0, 0x49, 0xb8, 0x27, 0x68, 0x2c, 0xe4,
	0x17, 0xa0, 0x13, 0x50, 0xe8, 0x28, 0xf4, 0x01, 0x54, 0x6b, 0x82, 0x36, 0xd8, 0xbd, 0x7c, 0xc2,
	0x9c, 0xdc, 0x46, 0xa2, 0x7a, 0x06, 0xb3, 0xb8, 0x2d, 0x1f, 0x18, 0x77, 0x64, 0x4f, 0x03, 0x79,
	0x50, 0xcc, 0x56, 0xb5, 0xb7, 0xd7, 0x52, 0x2e, 0xc6, 0x5e, 0xda, 0x36, 0x27, 0x42, 0x34, 0x24,
	0x77, 0x7c, 0x5a, 0x9f, 0x5a, 0xd5, 0x0a, 0xcc, 0x04, 0x2e, 0xf6, 0xb5, 0x95, 0x3c, 0x28, 0xae,
	0x57, 0x34, 0x33, 0xdd, 0xce, 0xbc, 0x76, 0xb1, 0x5f, 0xcd, 0x0c, 0x3e, 0x0f, 0x94, 0x7a, 0xe8,
	0xbd, 0xd8, 0x7c, 0xfe, 0x7e, 0x39, 0x9e, 0x66, 0x14, 0xf6, 0xa1, 0x9e, 0xde, 0xa8, 0x4e, 0x44,
	0xc0, 0x7c, 0x41, 0x0a, 0x4d, 0xb8, 0x5d, 0x13, 0xf4, 0x0a, 0xfb, 0x16, 0x71, 0xff, 0xb9, 0x6d,
	0x8a, 0xac, 0x43, 0x6d, 0x3e, 0x3b, 0xe1, 0x56, 0xde, 0x01, 0x5c, 0xad, 0x09, 0xaa, 0x3a, 0x70,
	0x6b, 0xfe, 0xb3, 0x8e, 0x16, 0xd5, 0x4c, 0x57, 0xd0, 0xcd, 0xe5, 0x7c, 0x09, 0x52, 0xb5, 0xe0,
	0xc6, 0x6c, 0xcf, 0xc3, 0x3f, 0x02, 0x66, 0x5c, 0xfa, 0xc9, 0x32, 0xae, 0x04, 0x52, 0xbd, 0x19,
	0x8c, 0x0c, 0x30, 0x1c, 0x19, 0xe0, 0x6b, 0x64, 0x80, 0xfe, 0xd8, 0x50, 0x86, 0x63, 0x43, 0xf9,
	0x18, 0x1b, 0x4a, 0xf3, 0xbc, 0xe5, 0xc8, 0x56, 0xdb, 0x7a, 0x24, 0xd2, 0x64, 0x9c, 0xa2, 0x38,
	0x54, 0x12, 0xec, 0x21, 0xca, 0x4a, 0xc2, 0xc3, 0x5c, 0x96, 0x7c, 0x66, 0x13, 0xd4, 0xfd, 0x3d,
	0x2f, 0xd9, 0x0b, 0x88, 0x68, 0xad, 0x85, 0xa7, 0x75, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0x99,
	0x6b, 0xa5, 0x6b, 0xf9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SoftwareUpgrade is a governance operation for initiating a software upgrade.
	SoftwareUpgrade(ctx context.Context, in *MsgSoftwareUpgrade, opts ...grpc.CallOption) (*MsgSoftwareUpgradeResponse, error)
	// CancelUpgrade is a governance operation for cancelling a previously
	// approvid software upgrade.
	CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SoftwareUpgrade(ctx context.Context, in *MsgSoftwareUpgrade, opts ...grpc.CallOption) (*MsgSoftwareUpgradeResponse, error) {
	out := new(MsgSoftwareUpgradeResponse)
	err := c.cc.Invoke(ctx, "/decimal.upgrade.v1.Msg/SoftwareUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelUpgrade(ctx context.Context, in *MsgCancelUpgrade, opts ...grpc.CallOption) (*MsgCancelUpgradeResponse, error) {
	out := new(MsgCancelUpgradeResponse)
	err := c.cc.Invoke(ctx, "/decimal.upgrade.v1.Msg/CancelUpgrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SoftwareUpgrade is a governance operation for initiating a software upgrade.
	SoftwareUpgrade(context.Context, *MsgSoftwareUpgrade) (*MsgSoftwareUpgradeResponse, error)
	// CancelUpgrade is a governance operation for cancelling a previously
	// approvid software upgrade.
	CancelUpgrade(context.Context, *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SoftwareUpgrade(ctx context.Context, req *MsgSoftwareUpgrade) (*MsgSoftwareUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SoftwareUpgrade not implemented")
}
func (*UnimplementedMsgServer) CancelUpgrade(ctx context.Context, req *MsgCancelUpgrade) (*MsgCancelUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelUpgrade not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SoftwareUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSoftwareUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SoftwareUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.upgrade.v1.Msg/SoftwareUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SoftwareUpgrade(ctx, req.(*MsgSoftwareUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.upgrade.v1.Msg/CancelUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelUpgrade(ctx, req.(*MsgCancelUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decimal.upgrade.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SoftwareUpgrade",
			Handler:    _Msg_SoftwareUpgrade_Handler,
		},
		{
			MethodName: "CancelUpgrade",
			Handler:    _Msg_CancelUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decimal/upgrade/v1/tx.proto",
}

func (m *MsgSoftwareUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSoftwareUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSoftwareUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSoftwareUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSoftwareUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSoftwareUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCancelUpgradeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCancelUpgradeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCancelUpgradeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSoftwareUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Plan.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSoftwareUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCancelUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCancelUpgradeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSoftwareUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSoftwareUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSoftwareUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSoftwareUpgradeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSoftwareUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSoftwareUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCancelUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCancelUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCancelUpgradeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCancelUpgradeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)