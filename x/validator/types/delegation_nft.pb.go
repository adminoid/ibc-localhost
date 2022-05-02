// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/validator/v1/delegation_nft.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DelegationNFT struct {
	DelegatorAddress string                                  `protobuf:"bytes,1,opt,name=delegator_address,json=delegatorAddress,proto3" json:"delegator_address,omitempty" yaml:"delegator_address"`
	ValidatorAddress string                                  `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty" yaml:"validator_address"`
	Denom            string                                  `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	TokenId          string                                  `protobuf:"bytes,4,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty" yaml:"token_id"`
	Quantity         github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,5,opt,name=quantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"quantity" yaml:"quantity"`
	Coin             github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=coin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"coin" yaml:"coin"`
}

func (m *DelegationNFT) Reset()         { *m = DelegationNFT{} }
func (m *DelegationNFT) String() string { return proto.CompactTextString(m) }
func (*DelegationNFT) ProtoMessage()    {}
func (*DelegationNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_93035fcf01557794, []int{0}
}
func (m *DelegationNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DelegationNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DelegationNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DelegationNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DelegationNFT.Merge(m, src)
}
func (m *DelegationNFT) XXX_Size() int {
	return m.Size()
}
func (m *DelegationNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_DelegationNFT.DiscardUnknown(m)
}

var xxx_messageInfo_DelegationNFT proto.InternalMessageInfo

type UnbondingDelegationNFTEntry struct {
	CreationHeight int64                                   `protobuf:"varint,1,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty" yaml:"creation_height"`
	CompletionTime time.Time                               `protobuf:"bytes,2,opt,name=completion_time,json=completionTime,proto3,stdtime" json:"completion_time" yaml:"completion_time"`
	Denom          string                                  `protobuf:"bytes,3,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	TokenId        string                                  `protobuf:"bytes,4,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty" yaml:"token_id"`
	Quantity       github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,5,opt,name=quantity,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"quantity" yaml:"quantity"`
	Balance        github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=balance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"balance" yaml:"balance"`
}

func (m *UnbondingDelegationNFTEntry) Reset()         { *m = UnbondingDelegationNFTEntry{} }
func (m *UnbondingDelegationNFTEntry) String() string { return proto.CompactTextString(m) }
func (*UnbondingDelegationNFTEntry) ProtoMessage()    {}
func (*UnbondingDelegationNFTEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_93035fcf01557794, []int{1}
}
func (m *UnbondingDelegationNFTEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnbondingDelegationNFTEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnbondingDelegationNFTEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnbondingDelegationNFTEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbondingDelegationNFTEntry.Merge(m, src)
}
func (m *UnbondingDelegationNFTEntry) XXX_Size() int {
	return m.Size()
}
func (m *UnbondingDelegationNFTEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbondingDelegationNFTEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UnbondingDelegationNFTEntry proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DelegationNFT)(nil), "decimal.validator.v1.DelegationNFT")
	proto.RegisterType((*UnbondingDelegationNFTEntry)(nil), "decimal.validator.v1.UnbondingDelegationNFTEntry")
}

func init() {
	proto.RegisterFile("decimal/validator/v1/delegation_nft.proto", fileDescriptor_93035fcf01557794)
}

var fileDescriptor_93035fcf01557794 = []byte{
	// 545 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x94, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0x6d, 0xd2, 0x37, 0x5c, 0x68, 0x83, 0xa9, 0x90, 0x15, 0x90, 0x8d, 0x6e, 0x28, 0x30,
	0xc4, 0x56, 0x61, 0xa2, 0x03, 0x52, 0x52, 0x40, 0x64, 0x61, 0x30, 0x65, 0x01, 0xa1, 0xe8, 0x6c,
	0x5f, 0x2f, 0xa7, 0xf8, 0xee, 0x09, 0xf6, 0x25, 0x22, 0xdf, 0x80, 0xb1, 0x1f, 0xa1, 0x9f, 0x83,
	0x4f, 0xd0, 0xb1, 0x23, 0x62, 0x30, 0x28, 0x59, 0x60, 0xcd, 0x27, 0x40, 0x3e, 0xbf, 0xa4, 0xa1,
	0x0b, 0x8c, 0x4c, 0xb9, 0x3c, 0xf7, 0xbb, 0xdf, 0x73, 0xd6, 0xff, 0xd1, 0x19, 0x8f, 0x22, 0x12,
	0x32, 0x8e, 0x63, 0x6f, 0x82, 0x63, 0x16, 0x61, 0x09, 0x89, 0x37, 0x39, 0xf0, 0x22, 0x12, 0x13,
	0x8a, 0x25, 0x03, 0xd1, 0x17, 0x27, 0xd2, 0x1d, 0x25, 0x20, 0xc1, 0xdc, 0x2b, 0x51, 0xb7, 0x46,
	0xdd, 0xc9, 0x41, 0x6b, 0x8f, 0x02, 0x05, 0x05, 0x78, 0xf9, 0xaa, 0x60, 0x5b, 0x0e, 0x05, 0xa0,
	0x31, 0xf1, 0xd4, 0xbf, 0x60, 0x7c, 0xe2, 0x49, 0xc6, 0x49, 0x2a, 0x31, 0x1f, 0x15, 0x00, 0xfa,
	0xd2, 0x30, 0x6e, 0x3e, 0xaf, 0xbb, 0xbc, 0x7e, 0x79, 0x6c, 0xf6, 0x8c, 0x5b, 0x65, 0x5b, 0x48,
	0xfa, 0x38, 0x8a, 0x12, 0x92, 0xa6, 0x96, 0x7e, 0x5f, 0x7f, 0x78, 0xbd, 0x7b, 0x6f, 0x91, 0x39,
	0xd6, 0x14, 0xf3, 0xf8, 0x10, 0x5d, 0x41, 0x90, 0xdf, 0xac, 0x6b, 0x9d, 0xa2, 0x94, 0xab, 0xea,
	0x3b, 0xd6, 0xaa, 0x6b, 0x7f, 0xaa, 0xae, 0x20, 0xc8, 0x6f, 0xd6, 0xb5, 0x4a, 0xb5, 0x6f, 0xac,
	0x47, 0x44, 0x00, 0xb7, 0x1a, 0xea, 0x78, 0x73, 0x91, 0x39, 0x37, 0xaa, 0x9b, 0x08, 0xe0, 0xc8,
	0x2f, 0xb6, 0x4d, 0xd7, 0xd8, 0x92, 0x30, 0x24, 0xa2, 0xcf, 0x22, 0x6b, 0x4d, 0xa1, 0xb7, 0x17,
	0x99, 0xb3, 0x5b, 0xa0, 0xd5, 0x0e, 0xf2, 0x37, 0xd5, 0xb2, 0x17, 0x99, 0x1f, 0x8c, 0xad, 0x8f,
	0x63, 0x2c, 0x24, 0x93, 0x53, 0x6b, 0x5d, 0xf1, 0x9d, 0xf3, 0xcc, 0xd1, 0xbe, 0x65, 0xce, 0x3e,
	0x65, 0x72, 0x30, 0x0e, 0xdc, 0x10, 0xb8, 0x17, 0x42, 0xca, 0x21, 0x2d, 0x7f, 0xda, 0x69, 0x34,
	0xf4, 0xe4, 0x74, 0x44, 0x52, 0xb7, 0x27, 0xe4, 0xd2, 0x5e, 0x79, 0x90, 0x5f, 0x2b, 0x4d, 0xdf,
	0x58, 0x0b, 0x81, 0x09, 0x6b, 0x43, 0xa9, 0x9f, 0x95, 0xea, 0x07, 0x7f, 0xa1, 0x3e, 0x02, 0x26,
	0x16, 0x99, 0xb3, 0x5d, 0xb8, 0x73, 0x09, 0xf2, 0x95, 0xeb, 0x70, 0xeb, 0xf3, 0x99, 0xa3, 0xfd,
	0x3c, 0x73, 0x34, 0xf4, 0xab, 0x61, 0xdc, 0x7d, 0x2b, 0x02, 0x10, 0x11, 0x13, 0x74, 0x25, 0xc5,
	0x17, 0x42, 0x26, 0x53, 0xf3, 0xc8, 0xd8, 0x0d, 0x13, 0x52, 0xcc, 0xcf, 0x80, 0x30, 0x3a, 0x90,
	0x2a, 0xc8, 0x46, 0xb7, 0xb5, 0xc8, 0x9c, 0x3b, 0xa5, 0x79, 0x15, 0x40, 0xfe, 0x4e, 0x55, 0x79,
	0xa5, 0x0a, 0x26, 0x35, 0x76, 0x43, 0xe0, 0xa3, 0x98, 0x28, 0x2a, 0x9f, 0x1f, 0x15, 0xe1, 0xf6,
	0xe3, 0x96, 0x5b, 0x0c, 0x97, 0x5b, 0x0d, 0x97, 0x7b, 0x5c, 0x0d, 0x57, 0x17, 0xe5, 0x5f, 0x7a,
	0xa9, 0xc9, 0xaa, 0x00, 0x9d, 0x7e, 0x77, 0x74, 0x7f, 0x67, 0x59, 0xcd, 0x0f, 0xfe, 0xaf, 0x11,
	0xbf, 0x37, 0x36, 0x03, 0x1c, 0x63, 0x11, 0x92, 0x32, 0xe5, 0xce, 0xbf, 0xa7, 0xbc, 0x53, 0xe8,
	0x4b, 0x0f, 0xf2, 0x2b, 0xe3, 0x32, 0xeb, 0xee, 0x9b, 0xf3, 0x99, 0xad, 0x5f, 0xcc, 0x6c, 0xfd,
	0xc7, 0xcc, 0xd6, 0x4f, 0xe7, 0xb6, 0x76, 0x31, 0xb7, 0xb5, 0xaf, 0x73, 0x5b, 0x7b, 0xf7, 0x34,
	0x60, 0x32, 0x18, 0x87, 0x43, 0x22, 0x5d, 0x48, 0xa8, 0x57, 0xbe, 0x0e, 0x92, 0x60, 0xee, 0x51,
	0x68, 0xa7, 0x1c, 0x27, 0xb2, 0x2d, 0x20, 0x22, 0xde, 0xa7, 0x4b, 0x8f, 0x8b, 0x6a, 0x1f, 0x6c,
	0xa8, 0xe8, 0x9e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x19, 0x9f, 0x7e, 0x4c, 0x7e, 0x04, 0x00,
	0x00,
}

func (m *DelegationNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DelegationNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DelegationNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Coin.Size()
		i -= size
		if _, err := m.Coin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegationNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Quantity.Size()
		i -= size
		if _, err := m.Quantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegationNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintDelegationNft(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDelegationNft(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintDelegationNft(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DelegatorAddress) > 0 {
		i -= len(m.DelegatorAddress)
		copy(dAtA[i:], m.DelegatorAddress)
		i = encodeVarintDelegationNft(dAtA, i, uint64(len(m.DelegatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnbondingDelegationNFTEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnbondingDelegationNFTEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnbondingDelegationNFTEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Balance.Size()
		i -= size
		if _, err := m.Balance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegationNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Quantity.Size()
		i -= size
		if _, err := m.Quantity.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDelegationNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintDelegationNft(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintDelegationNft(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x1a
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CompletionTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintDelegationNft(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.CreationHeight != 0 {
		i = encodeVarintDelegationNft(dAtA, i, uint64(m.CreationHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintDelegationNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovDelegationNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DelegationNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DelegatorAddress)
	if l > 0 {
		n += 1 + l + sovDelegationNft(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovDelegationNft(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDelegationNft(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovDelegationNft(uint64(l))
	}
	l = m.Quantity.Size()
	n += 1 + l + sovDelegationNft(uint64(l))
	l = m.Coin.Size()
	n += 1 + l + sovDelegationNft(uint64(l))
	return n
}

func (m *UnbondingDelegationNFTEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CreationHeight != 0 {
		n += 1 + sovDelegationNft(uint64(m.CreationHeight))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CompletionTime)
	n += 1 + l + sovDelegationNft(uint64(l))
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovDelegationNft(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovDelegationNft(uint64(l))
	}
	l = m.Quantity.Size()
	n += 1 + l + sovDelegationNft(uint64(l))
	l = m.Balance.Size()
	n += 1 + l + sovDelegationNft(uint64(l))
	return n
}

func sovDelegationNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDelegationNft(x uint64) (n int) {
	return sovDelegationNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DelegationNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegationNft
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
			return fmt.Errorf("proto: DelegationNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DelegationNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DelegatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegationNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegationNft
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
func (m *UnbondingDelegationNFTEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDelegationNft
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
			return fmt.Errorf("proto: UnbondingDelegationNFTEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnbondingDelegationNFTEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationHeight", wireType)
			}
			m.CreationHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreationHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompletionTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CompletionTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Quantity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDelegationNft
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
				return ErrInvalidLengthDelegationNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDelegationNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDelegationNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDelegationNft
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
func skipDelegationNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDelegationNft
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
					return 0, ErrIntOverflowDelegationNft
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
					return 0, ErrIntOverflowDelegationNft
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
				return 0, ErrInvalidLengthDelegationNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDelegationNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDelegationNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDelegationNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDelegationNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDelegationNft = fmt.Errorf("proto: unexpected end of group")
)
