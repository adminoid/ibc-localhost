// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/swap/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type TransferType int32

const (
	TransferType_TransferTypeOut TransferType = 0
	TransferType_TransferTypeIn  TransferType = 1
)

var TransferType_name = map[int32]string{
	0: "TransferTypeOut",
	1: "TransferTypeIn",
}

var TransferType_value = map[string]int32{
	"TransferTypeOut": 0,
	"TransferTypeIn":  1,
}

func (TransferType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{0}
}

type MsgHTLT struct {
	TransferType TransferType                             `protobuf:"varint,1,opt,name=transfer_type,json=transferType,proto3,enum=decimal.swap.v1.TransferType" json:"transfer_type,omitempty" yaml:"transfer_type"`
	From         string                                   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Recipient    string                                   `protobuf:"bytes,3,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
	HashedSecret Hash                                     `protobuf:"bytes,4,opt,name=hashed_secret,json=hashedSecret,proto3,customtype=Hash" json:"hashed_secret" yaml:"hashed_secret"`
	Secret       []byte                                   `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty" yaml:"secret"`
	Amount       github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount" yaml:"amount"`
}

func (m *MsgHTLT) Reset()         { *m = MsgHTLT{} }
func (m *MsgHTLT) String() string { return proto.CompactTextString(m) }
func (*MsgHTLT) ProtoMessage()    {}
func (*MsgHTLT) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{0}
}
func (m *MsgHTLT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgHTLT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgHTLT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgHTLT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHTLT.Merge(m, src)
}
func (m *MsgHTLT) XXX_Size() int {
	return m.Size()
}
func (m *MsgHTLT) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHTLT.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHTLT proto.InternalMessageInfo

func (m *MsgHTLT) GetTransferType() TransferType {
	if m != nil {
		return m.TransferType
	}
	return TransferType_TransferTypeOut
}

func (m *MsgHTLT) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *MsgHTLT) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *MsgHTLT) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *MsgHTLT) GetAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type MsgHTLTResponse struct {
}

func (m *MsgHTLTResponse) Reset()         { *m = MsgHTLTResponse{} }
func (m *MsgHTLTResponse) String() string { return proto.CompactTextString(m) }
func (*MsgHTLTResponse) ProtoMessage()    {}
func (*MsgHTLTResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{1}
}
func (m *MsgHTLTResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgHTLTResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgHTLTResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgHTLTResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgHTLTResponse.Merge(m, src)
}
func (m *MsgHTLTResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgHTLTResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgHTLTResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgHTLTResponse proto.InternalMessageInfo

type MsgRedeem struct {
	From   string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	Secret []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (m *MsgRedeem) Reset()         { *m = MsgRedeem{} }
func (m *MsgRedeem) String() string { return proto.CompactTextString(m) }
func (*MsgRedeem) ProtoMessage()    {}
func (*MsgRedeem) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{2}
}
func (m *MsgRedeem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeem.Merge(m, src)
}
func (m *MsgRedeem) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeem) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeem.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeem proto.InternalMessageInfo

type MsgRedeemResponse struct {
}

func (m *MsgRedeemResponse) Reset()         { *m = MsgRedeemResponse{} }
func (m *MsgRedeemResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRedeemResponse) ProtoMessage()    {}
func (*MsgRedeemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{3}
}
func (m *MsgRedeemResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRedeemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRedeemResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRedeemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRedeemResponse.Merge(m, src)
}
func (m *MsgRedeemResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRedeemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRedeemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRedeemResponse proto.InternalMessageInfo

type MsgRefund struct {
	From         string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty" yaml:"from"`
	HashedSecret Hash   `protobuf:"bytes,2,opt,name=hashed_secret,json=hashedSecret,proto3,customtype=Hash" json:"hashed_secret"`
}

func (m *MsgRefund) Reset()         { *m = MsgRefund{} }
func (m *MsgRefund) String() string { return proto.CompactTextString(m) }
func (*MsgRefund) ProtoMessage()    {}
func (*MsgRefund) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{4}
}
func (m *MsgRefund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRefund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRefund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRefund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRefund.Merge(m, src)
}
func (m *MsgRefund) XXX_Size() int {
	return m.Size()
}
func (m *MsgRefund) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRefund.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRefund proto.InternalMessageInfo

type MsgRefundResponse struct {
}

func (m *MsgRefundResponse) Reset()         { *m = MsgRefundResponse{} }
func (m *MsgRefundResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRefundResponse) ProtoMessage()    {}
func (*MsgRefundResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cd30fd052b718b3, []int{5}
}
func (m *MsgRefundResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRefundResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRefundResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRefundResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRefundResponse.Merge(m, src)
}
func (m *MsgRefundResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRefundResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRefundResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRefundResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("decimal.swap.v1.TransferType", TransferType_name, TransferType_value)
	proto.RegisterType((*MsgHTLT)(nil), "decimal.swap.v1.MsgHTLT")
	proto.RegisterType((*MsgHTLTResponse)(nil), "decimal.swap.v1.MsgHTLTResponse")
	proto.RegisterType((*MsgRedeem)(nil), "decimal.swap.v1.MsgRedeem")
	proto.RegisterType((*MsgRedeemResponse)(nil), "decimal.swap.v1.MsgRedeemResponse")
	proto.RegisterType((*MsgRefund)(nil), "decimal.swap.v1.MsgRefund")
	proto.RegisterType((*MsgRefundResponse)(nil), "decimal.swap.v1.MsgRefundResponse")
}

func init() { proto.RegisterFile("decimal/swap/v1/tx.proto", fileDescriptor_2cd30fd052b718b3) }

var fileDescriptor_2cd30fd052b718b3 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xb5, 0xbf, 0xf8, 0x0b, 0x74, 0xea, 0x36, 0xed, 0xb4, 0x42, 0xc6, 0x12, 0x76, 0x64, 0x58,
	0x04, 0xa4, 0x7a, 0x94, 0xc0, 0xaa, 0x12, 0x0b, 0xcc, 0xa6, 0x95, 0x28, 0x48, 0x26, 0x2b, 0x84,
	0x54, 0xf9, 0x67, 0xea, 0x58, 0xad, 0x3d, 0x96, 0x67, 0x52, 0xda, 0x37, 0x60, 0xc9, 0x23, 0x74,
	0xc1, 0x02, 0xb1, 0xe4, 0x29, 0xba, 0xec, 0x12, 0x21, 0x61, 0x50, 0xb2, 0x61, 0x9d, 0x27, 0x40,
	0x1e, 0x4f, 0x12, 0xa7, 0x34, 0x88, 0x55, 0x32, 0x3e, 0xe7, 0x9e, 0x73, 0xe7, 0xdc, 0xab, 0x01,
	0x5a, 0x88, 0x83, 0x38, 0xf1, 0x4e, 0x10, 0x7d, 0xe7, 0x65, 0xe8, 0xb4, 0x8b, 0xd8, 0x99, 0x9d,
	0xe5, 0x84, 0x11, 0xd8, 0x12, 0x88, 0x5d, 0x22, 0xf6, 0x69, 0x57, 0xdf, 0x8e, 0x48, 0x44, 0x38,
	0x86, 0xca, 0x7f, 0x15, 0x4d, 0x37, 0x02, 0x42, 0x13, 0x42, 0x91, 0xef, 0x51, 0x8c, 0x4e, 0xbb,
	0x3e, 0x66, 0x5e, 0x17, 0x05, 0x24, 0x4e, 0x2b, 0xdc, 0xfa, 0xd2, 0x00, 0xb7, 0x0e, 0x68, 0xb4,
	0xd7, 0x7f, 0xd1, 0x87, 0x6f, 0xc1, 0x1a, 0xcb, 0xbd, 0x94, 0x1e, 0xe1, 0xfc, 0x90, 0x9d, 0x67,
	0x58, 0x93, 0xdb, 0x72, 0x67, 0xbd, 0x77, 0xcf, 0xbe, 0x66, 0x65, 0xf7, 0x05, 0xab, 0x7f, 0x9e,
	0x61, 0x47, 0x9b, 0x14, 0xe6, 0xf6, 0xb9, 0x97, 0x9c, 0xec, 0x5a, 0x0b, 0xd5, 0x96, 0xab, 0xb2,
	0x1a, 0x0f, 0xde, 0x07, 0xca, 0x51, 0x4e, 0x12, 0xed, 0xbf, 0xb6, 0xdc, 0x59, 0x71, 0x5a, 0x93,
	0xc2, 0x5c, 0xad, 0xaa, 0xca, 0xaf, 0x96, 0xcb, 0x41, 0xd8, 0x03, 0x2b, 0x39, 0x0e, 0xe2, 0x2c,
	0xc6, 0x29, 0xd3, 0x1a, 0x9c, 0xb9, 0x3d, 0x29, 0xcc, 0x8d, 0x8a, 0x39, 0x83, 0x2c, 0x77, 0x4e,
	0x83, 0xfb, 0x60, 0x6d, 0xe0, 0xd1, 0x01, 0x0e, 0x0f, 0x29, 0x0e, 0x72, 0xcc, 0x34, 0xa5, 0x2d,
	0x77, 0x54, 0xe7, 0xc1, 0x65, 0x61, 0x4a, 0xdf, 0x0a, 0x53, 0xd9, 0xf3, 0xe8, 0x60, 0xde, 0xe3,
	0x02, 0xd5, 0x72, 0xd5, 0xea, 0xfc, 0x9a, 0x1f, 0xe1, 0x43, 0xd0, 0x14, 0x1a, 0xff, 0x73, 0x8d,
	0xcd, 0x49, 0x61, 0xae, 0x55, 0x75, 0xd3, 0x02, 0x41, 0x80, 0x0c, 0x34, 0xbd, 0x84, 0x0c, 0x53,
	0xa6, 0x35, 0xdb, 0x8d, 0xce, 0x6a, 0xef, 0xae, 0x5d, 0x25, 0x6d, 0x97, 0x49, 0xdb, 0x22, 0x69,
	0xfb, 0x39, 0x89, 0x53, 0xe7, 0x59, 0xd9, 0xc9, 0x5c, 0xa9, 0x2a, 0xb3, 0x3e, 0xff, 0x30, 0x3b,
	0x51, 0xcc, 0x06, 0x43, 0xdf, 0x0e, 0x48, 0x82, 0xc4, 0x9c, 0xaa, 0x9f, 0x1d, 0x1a, 0x1e, 0xa3,
	0x32, 0x44, 0xca, 0x15, 0xa8, 0x2b, 0xbc, 0x76, 0x95, 0x5f, 0x17, 0xa6, 0x64, 0x6d, 0x82, 0x96,
	0x98, 0x99, 0x8b, 0x69, 0x46, 0x52, 0x8a, 0x2d, 0x17, 0xac, 0x1c, 0xd0, 0xc8, 0xc5, 0x21, 0xc6,
	0xc9, 0x2c, 0x6a, 0xf9, 0x6f, 0x51, 0xdf, 0x99, 0xdd, 0xb5, 0x9c, 0x88, 0x3a, 0xbd, 0xd8, 0xee,
	0xed, 0xf7, 0x17, 0xa6, 0xc4, 0x6d, 0xb6, 0xc0, 0xe6, 0x4c, 0x73, 0x66, 0x94, 0x08, 0xa3, 0xa3,
	0x61, 0x1a, 0xfe, 0x9b, 0x51, 0xf7, 0xfa, 0x7c, 0xb8, 0x9f, 0xa3, 0xd6, 0xe7, 0xb3, 0x38, 0x87,
	0x1b, 0x7a, 0x28, 0xed, 0xa6, 0x3d, 0x3c, 0x7a, 0x0a, 0xd4, 0xfa, 0x0a, 0xc2, 0x2d, 0xd0, 0xaa,
	0x9f, 0x5f, 0x0d, 0xd9, 0x86, 0x04, 0x21, 0x58, 0xaf, 0x7f, 0xdc, 0x4f, 0x37, 0x64, 0x5d, 0xf9,
	0xf4, 0xd1, 0x90, 0x7a, 0xdf, 0x65, 0xd0, 0x38, 0xa0, 0x11, 0x74, 0x80, 0xc2, 0xf7, 0x5e, 0xfb,
	0x63, 0xc1, 0x45, 0xba, 0x7a, 0x7b, 0x19, 0x32, 0x6d, 0x05, 0xee, 0x81, 0xa6, 0x08, 0x5d, 0xbf,
	0x89, 0x5b, 0x61, 0xba, 0xb5, 0x1c, 0x5b, 0x54, 0xe2, 0xa9, 0x2e, 0x51, 0x2a, 0xb1, 0x65, 0x4a,
	0xf5, 0x78, 0x9c, 0x97, 0x97, 0x23, 0x43, 0xbe, 0x1a, 0x19, 0xf2, 0xcf, 0x91, 0x21, 0x7f, 0x18,
	0x1b, 0xd2, 0xd5, 0xd8, 0x90, 0xbe, 0x8e, 0x0d, 0xe9, 0xcd, 0x13, 0x3f, 0x66, 0xfe, 0x30, 0x38,
	0xc6, 0xcc, 0x26, 0x79, 0x84, 0x84, 0x14, 0xc3, 0x5e, 0x82, 0x22, 0xb2, 0x43, 0x13, 0x2f, 0x67,
	0x3b, 0x29, 0x09, 0x31, 0x3a, 0xab, 0x1e, 0x1c, 0xbe, 0x82, 0x7e, 0x93, 0x3f, 0x15, 0x8f, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xec, 0xed, 0xc6, 0x16, 0x8d, 0x04, 0x00, 0x00,
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
	HTLT(ctx context.Context, in *MsgHTLT, opts ...grpc.CallOption) (*MsgHTLTResponse, error)
	Redeem(ctx context.Context, in *MsgRedeem, opts ...grpc.CallOption) (*MsgRedeemResponse, error)
	Refund(ctx context.Context, in *MsgRefund, opts ...grpc.CallOption) (*MsgRefundResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) HTLT(ctx context.Context, in *MsgHTLT, opts ...grpc.CallOption) (*MsgHTLTResponse, error) {
	out := new(MsgHTLTResponse)
	err := c.cc.Invoke(ctx, "/decimal.swap.v1.Msg/HTLT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Redeem(ctx context.Context, in *MsgRedeem, opts ...grpc.CallOption) (*MsgRedeemResponse, error) {
	out := new(MsgRedeemResponse)
	err := c.cc.Invoke(ctx, "/decimal.swap.v1.Msg/Redeem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Refund(ctx context.Context, in *MsgRefund, opts ...grpc.CallOption) (*MsgRefundResponse, error) {
	out := new(MsgRefundResponse)
	err := c.cc.Invoke(ctx, "/decimal.swap.v1.Msg/Refund", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	HTLT(context.Context, *MsgHTLT) (*MsgHTLTResponse, error)
	Redeem(context.Context, *MsgRedeem) (*MsgRedeemResponse, error)
	Refund(context.Context, *MsgRefund) (*MsgRefundResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) HTLT(ctx context.Context, req *MsgHTLT) (*MsgHTLTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HTLT not implemented")
}
func (*UnimplementedMsgServer) Redeem(ctx context.Context, req *MsgRedeem) (*MsgRedeemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Redeem not implemented")
}
func (*UnimplementedMsgServer) Refund(ctx context.Context, req *MsgRefund) (*MsgRefundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Refund not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_HTLT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgHTLT)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).HTLT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.swap.v1.Msg/HTLT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).HTLT(ctx, req.(*MsgHTLT))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Redeem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRedeem)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Redeem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.swap.v1.Msg/Redeem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Redeem(ctx, req.(*MsgRedeem))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Refund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRefund)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Refund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/decimal.swap.v1.Msg/Refund",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Refund(ctx, req.(*MsgRefund))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "decimal.swap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HTLT",
			Handler:    _Msg_HTLT_Handler,
		},
		{
			MethodName: "Redeem",
			Handler:    _Msg_Redeem_Handler,
		},
		{
			MethodName: "Refund",
			Handler:    _Msg_Refund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decimal/swap/v1/tx.proto",
}

func (m *MsgHTLT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgHTLT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgHTLT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x2a
	}
	{
		size := m.HashedSecret.Size()
		i -= size
		if _, err := m.HashedSecret.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if m.TransferType != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TransferType))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgHTLTResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgHTLTResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgHTLTResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRedeem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Secret) > 0 {
		i -= len(m.Secret)
		copy(dAtA[i:], m.Secret)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Secret)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRedeemResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRedeemResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRedeemResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRefund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRefund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRefund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.HashedSecret.Size()
		i -= size
		if _, err := m.HashedSecret.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintTx(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRefundResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRefundResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRefundResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgHTLT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TransferType != 0 {
		n += 1 + sovTx(uint64(m.TransferType))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.HashedSecret.Size()
	n += 1 + l + sovTx(uint64(l))
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgHTLTResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRedeem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Secret)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRedeemResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRefund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.HashedSecret.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgRefundResponse) Size() (n int) {
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
func (m *MsgHTLT) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgHTLT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgHTLT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TransferType", wireType)
			}
			m.TransferType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TransferType |= TransferType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
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
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashedSecret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HashedSecret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = append(m.Secret[:0], dAtA[iNdEx:postIndex]...)
			if m.Secret == nil {
				m.Secret = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgHTLTResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgHTLTResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgHTLTResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgRedeem) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRedeem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secret = append(m.Secret[:0], dAtA[iNdEx:postIndex]...)
			if m.Secret == nil {
				m.Secret = []byte{}
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
func (m *MsgRedeemResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRedeemResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRedeemResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgRefund) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRefund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRefund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HashedSecret", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HashedSecret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgRefundResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRefundResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRefundResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
