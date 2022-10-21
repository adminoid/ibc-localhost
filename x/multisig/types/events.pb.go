// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/multisig/v1/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// EventCreateWallet defines event emitted when new multisig wallet is created.
type EventCreateWallet struct {
	Sender    string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Wallet    string   `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Owners    []string `protobuf:"bytes,3,rep,name=owners,proto3" json:"owners,omitempty"`
	Weights   []uint32 `protobuf:"varint,4,rep,packed,name=weights,proto3" json:"weights,omitempty"`
	Threshold uint32   `protobuf:"varint,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (m *EventCreateWallet) Reset()         { *m = EventCreateWallet{} }
func (m *EventCreateWallet) String() string { return proto.CompactTextString(m) }
func (*EventCreateWallet) ProtoMessage()    {}
func (*EventCreateWallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a14a7690049f30, []int{0}
}
func (m *EventCreateWallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateWallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateWallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateWallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateWallet.Merge(m, src)
}
func (m *EventCreateWallet) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateWallet) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateWallet.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateWallet proto.InternalMessageInfo

func (m *EventCreateWallet) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventCreateWallet) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *EventCreateWallet) GetOwners() []string {
	if m != nil {
		return m.Owners
	}
	return nil
}

func (m *EventCreateWallet) GetWeights() []uint32 {
	if m != nil {
		return m.Weights
	}
	return nil
}

func (m *EventCreateWallet) GetThreshold() uint32 {
	if m != nil {
		return m.Threshold
	}
	return 0
}

type EventCreateTransaction struct {
	Sender      string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Wallet      string `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Transaction string `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (m *EventCreateTransaction) Reset()         { *m = EventCreateTransaction{} }
func (m *EventCreateTransaction) String() string { return proto.CompactTextString(m) }
func (*EventCreateTransaction) ProtoMessage()    {}
func (*EventCreateTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a14a7690049f30, []int{1}
}
func (m *EventCreateTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreateTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreateTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreateTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreateTransaction.Merge(m, src)
}
func (m *EventCreateTransaction) XXX_Size() int {
	return m.Size()
}
func (m *EventCreateTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreateTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreateTransaction proto.InternalMessageInfo

func (m *EventCreateTransaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventCreateTransaction) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *EventCreateTransaction) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

type EventSignTransaction struct {
	Sender        string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Wallet        string `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Transaction   string `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction,omitempty"`
	SignerWeight  uint32 `protobuf:"varint,4,opt,name=signer_weight,json=signerWeight,proto3" json:"signer_weight,omitempty"`
	Confirmations uint32 `protobuf:"varint,5,opt,name=confirmations,proto3" json:"confirmations,omitempty"`
	Confirmed     bool   `protobuf:"varint,6,opt,name=confirmed,proto3" json:"confirmed,omitempty"`
}

func (m *EventSignTransaction) Reset()         { *m = EventSignTransaction{} }
func (m *EventSignTransaction) String() string { return proto.CompactTextString(m) }
func (*EventSignTransaction) ProtoMessage()    {}
func (*EventSignTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_52a14a7690049f30, []int{2}
}
func (m *EventSignTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSignTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSignTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSignTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSignTransaction.Merge(m, src)
}
func (m *EventSignTransaction) XXX_Size() int {
	return m.Size()
}
func (m *EventSignTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSignTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_EventSignTransaction proto.InternalMessageInfo

func (m *EventSignTransaction) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *EventSignTransaction) GetWallet() string {
	if m != nil {
		return m.Wallet
	}
	return ""
}

func (m *EventSignTransaction) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *EventSignTransaction) GetSignerWeight() uint32 {
	if m != nil {
		return m.SignerWeight
	}
	return 0
}

func (m *EventSignTransaction) GetConfirmations() uint32 {
	if m != nil {
		return m.Confirmations
	}
	return 0
}

func (m *EventSignTransaction) GetConfirmed() bool {
	if m != nil {
		return m.Confirmed
	}
	return false
}

func init() {
	proto.RegisterType((*EventCreateWallet)(nil), "decimal.multisig.v1.EventCreateWallet")
	proto.RegisterType((*EventCreateTransaction)(nil), "decimal.multisig.v1.EventCreateTransaction")
	proto.RegisterType((*EventSignTransaction)(nil), "decimal.multisig.v1.EventSignTransaction")
}

func init() { proto.RegisterFile("decimal/multisig/v1/events.proto", fileDescriptor_52a14a7690049f30) }

var fileDescriptor_52a14a7690049f30 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0xbd, 0x8e, 0xd3, 0x40,
	0x10, 0xce, 0x5e, 0x8e, 0x40, 0x16, 0x5c, 0x60, 0xa2, 0xd3, 0x72, 0x42, 0x96, 0x15, 0x28, 0xdc,
	0x24, 0x4b, 0x44, 0x43, 0xcb, 0x21, 0x5e, 0xc0, 0x87, 0x74, 0x12, 0xcd, 0x69, 0x6d, 0x0f, 0x9b,
	0x15, 0xf6, 0xee, 0x69, 0x67, 0x2e, 0x81, 0x07, 0xa0, 0xe7, 0x09, 0x78, 0x0a, 0xde, 0x01, 0xca,
	0x13, 0x15, 0x25, 0x4a, 0x5e, 0x04, 0xd9, 0xde, 0x93, 0x0f, 0x51, 0xd0, 0x41, 0xe7, 0xef, 0xcf,
	0x33, 0xe3, 0xf1, 0xf0, 0xb4, 0x82, 0xd2, 0x34, 0xaa, 0x96, 0xcd, 0x65, 0x4d, 0x06, 0x8d, 0x96,
	0x9b, 0x95, 0x84, 0x0d, 0x58, 0xc2, 0xe5, 0x85, 0x77, 0xe4, 0xe2, 0x07, 0xc1, 0xb1, 0xbc, 0x76,
	0x2c, 0x37, 0xab, 0xe3, 0x99, 0x76, 0xda, 0x75, 0xba, 0x6c, 0x9f, 0x7a, 0xeb, 0xf1, 0xc3, 0xd2,
	0x61, 0xe3, 0xf0, 0xbc, 0x17, 0x7a, 0x10, 0xa4, 0xa4, 0x47, 0xb2, 0x50, 0x08, 0x72, 0xb3, 0x2a,
	0x80, 0xd4, 0x4a, 0x96, 0xce, 0xd8, 0x5e, 0x9f, 0x7f, 0x65, 0xfc, 0xfe, 0xab, 0xb6, 0xec, 0x4b,
	0x0f, 0x8a, 0xe0, 0x4c, 0xd5, 0x35, 0x50, 0xfc, 0x94, 0x4f, 0x10, 0x6c, 0x05, 0x5e, 0xb0, 0x94,
	0x65, 0xd3, 0x13, 0xf1, 0xfd, 0xcb, 0x62, 0x16, 0xde, 0xfb, 0xa2, 0xaa, 0x3c, 0x20, 0x9e, 0x92,
	0x37, 0x56, 0xe7, 0xc1, 0xd7, 0x26, 0xb6, 0x5d, 0x56, 0x1c, 0xfc, 0x2d, 0xd1, 0xfb, 0xe2, 0x23,
	0x3e, 0x71, 0x5b, 0x0b, 0x1e, 0xc5, 0x38, 0x1d, 0x67, 0xd3, 0x3c, 0xa0, 0x58, 0xf0, 0xdb, 0x5b,
	0x30, 0x7a, 0x4d, 0x28, 0x0e, 0xd3, 0x71, 0x16, 0xe5, 0xd7, 0x30, 0x7e, 0xc4, 0xa7, 0xb4, 0xf6,
	0x80, 0x6b, 0x57, 0x57, 0xe2, 0x56, 0xca, 0xb2, 0x28, 0x1f, 0x88, 0xf9, 0x67, 0xc6, 0x8f, 0x6e,
	0x4c, 0xf2, 0xda, 0x2b, 0x8b, 0xaa, 0x24, 0xe3, 0xec, 0x3f, 0x19, 0x27, 0xe5, 0x77, 0x69, 0x28,
	0x29, 0x0e, 0xdb, 0x58, 0x7e, 0x93, 0x9a, 0x7f, 0x3c, 0xe0, 0xb3, 0xae, 0xc1, 0x53, 0xa3, 0xed,
	0x7f, 0x6e, 0x6f, 0xfc, 0x47, 0x7b, 0xf1, 0x63, 0x1e, 0xa1, 0xd1, 0x16, 0xfc, 0x79, 0xff, 0xbd,
	0xbb, 0x11, 0xa2, 0xfc, 0x5e, 0x4f, 0x9e, 0x75, 0x5c, 0xfc, 0x84, 0x47, 0xa5, 0xb3, 0x6f, 0x8d,
	0x6f, 0x54, 0x1b, 0xc2, 0xb0, 0x86, 0xdf, 0xc9, 0x76, 0x51, 0x81, 0x80, 0x4a, 0x4c, 0x52, 0x96,
	0xdd, 0xc9, 0x07, 0xe2, 0x24, 0xff, 0xb6, 0x4b, 0xd8, 0xd5, 0x2e, 0x61, 0x3f, 0x77, 0x09, 0xfb,
	0xb4, 0x4f, 0x46, 0x57, 0xfb, 0x64, 0xf4, 0x63, 0x9f, 0x8c, 0xde, 0x3c, 0x2f, 0x0c, 0x15, 0x97,
	0xe5, 0x3b, 0xa0, 0xa5, 0xf3, 0x5a, 0x86, 0x03, 0x20, 0x50, 0x8d, 0xd4, 0x6e, 0x81, 0x8d, 0xf2,
	0xb4, 0xb0, 0xae, 0x02, 0xf9, 0x7e, 0x38, 0x1b, 0xfa, 0x70, 0x01, 0x58, 0x4c, 0xba, 0xbf, 0xf9,
	0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa0, 0xf3, 0xf1, 0x57, 0x03, 0x00, 0x00,
}

func (m *EventCreateWallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateWallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateWallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Threshold != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Weights) > 0 {
		dAtA2 := make([]byte, len(m.Weights)*10)
		var j1 int
		for _, num := range m.Weights {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintEvents(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Owners[iNdEx])
			copy(dAtA[i:], m.Owners[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Owners[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Wallet) > 0 {
		i -= len(m.Wallet)
		copy(dAtA[i:], m.Wallet)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Wallet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventCreateTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreateTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreateTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Transaction) > 0 {
		i -= len(m.Transaction)
		copy(dAtA[i:], m.Transaction)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Transaction)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Wallet) > 0 {
		i -= len(m.Wallet)
		copy(dAtA[i:], m.Wallet)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Wallet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventSignTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSignTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSignTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Confirmed {
		i--
		if m.Confirmed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Confirmations != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Confirmations))
		i--
		dAtA[i] = 0x28
	}
	if m.SignerWeight != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.SignerWeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Transaction) > 0 {
		i -= len(m.Transaction)
		copy(dAtA[i:], m.Transaction)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Transaction)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Wallet) > 0 {
		i -= len(m.Wallet)
		copy(dAtA[i:], m.Wallet)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Wallet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreateWallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Wallet)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, s := range m.Owners {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if len(m.Weights) > 0 {
		l = 0
		for _, e := range m.Weights {
			l += sovEvents(uint64(e))
		}
		n += 1 + sovEvents(uint64(l)) + l
	}
	if m.Threshold != 0 {
		n += 1 + sovEvents(uint64(m.Threshold))
	}
	return n
}

func (m *EventCreateTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Wallet)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Transaction)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *EventSignTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Wallet)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Transaction)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.SignerWeight != 0 {
		n += 1 + sovEvents(uint64(m.SignerWeight))
	}
	if m.Confirmations != 0 {
		n += 1 + sovEvents(uint64(m.Confirmations))
	}
	if m.Confirmed {
		n += 2
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreateWallet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateWallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateWallet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owners = append(m.Owners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Weights = append(m.Weights, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEvents
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEvents
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthEvents
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Weights) == 0 {
					m.Weights = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEvents
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Weights = append(m.Weights, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Weights", wireType)
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Threshold |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventCreateTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventCreateTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreateTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transaction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSignTransaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSignTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSignTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Transaction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Transaction = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerWeight", wireType)
			}
			m.SignerWeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SignerWeight |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmations", wireType)
			}
			m.Confirmations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Confirmations |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Confirmed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Confirmed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
