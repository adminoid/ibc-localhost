// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/coin/v1/check.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Check struct {
	ChainID  string                                  `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id" yaml:"chain_id"`
	Coin     string                                  `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin" yaml:"coin"`
	Amount   github_com_cosmos_cosmos_sdk_types.Int  `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount" yaml:"amount"`
	Nonce    []byte                                  `protobuf:"bytes,4,opt,name=nonce,proto3" json:"nonce" yaml:"nonce"`
	DueBlock uint64                                  `protobuf:"varint,5,opt,name=due_block,json=dueBlock,proto3" json:"due_block" yaml:"due_block"`
	Lock     *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=lock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"lock" yaml:"lock"`
	V        *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=v,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"v" yaml:"v"`
	R        *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=r,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"r" yaml:"r"`
	S        *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=s,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"s" yaml:"s"`
	Redeemed bool                                    `protobuf:"varint,10,opt,name=redeemed,proto3" json:"redeemed" yaml:"redeemed"`
}

func (m *Check) Reset()      { *m = Check{} }
func (*Check) ProtoMessage() {}
func (*Check) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9b08460c6f27dc3, []int{0}
}
func (m *Check) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Check) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Check.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Check) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Check.Merge(m, src)
}
func (m *Check) XXX_Size() int {
	return m.Size()
}
func (m *Check) XXX_DiscardUnknown() {
	xxx_messageInfo_Check.DiscardUnknown(m)
}

var xxx_messageInfo_Check proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Check)(nil), "decimal.coin.v1.Check")
}

func init() { proto.RegisterFile("decimal/coin/v1/check.proto", fileDescriptor_c9b08460c6f27dc3) }

var fileDescriptor_c9b08460c6f27dc3 = []byte{
	// 488 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x33, 0x9a, 0xb6, 0xe9, 0x58, 0x59, 0x09, 0x82, 0x51, 0x21, 0x53, 0xe7, 0x20, 0x45,
	0x69, 0x86, 0x45, 0x4f, 0xbb, 0x20, 0x98, 0xd5, 0x43, 0x2e, 0x1e, 0x02, 0x5e, 0xbc, 0x2c, 0xc9,
	0x64, 0x48, 0x43, 0x9b, 0xcc, 0x92, 0x7f, 0xb8, 0xdf, 0xc0, 0xa3, 0x47, 0x8f, 0xfd, 0x38, 0x7b,
	0x5c, 0x3c, 0x89, 0x87, 0x41, 0xda, 0x8b, 0xe4, 0x98, 0x4f, 0xb0, 0x64, 0x26, 0x6d, 0xcf, 0xdb,
	0x53, 0xde, 0xf7, 0x79, 0xe6, 0xf9, 0x25, 0x13, 0xde, 0x17, 0xbe, 0x8c, 0x18, 0x4d, 0xd2, 0x60,
	0x45, 0x28, 0x4f, 0x32, 0x52, 0x9f, 0x12, 0xba, 0x60, 0x74, 0xe9, 0x5c, 0xe5, 0xbc, 0xe4, 0xe6,
	0x49, 0x6f, 0x3a, 0x9d, 0xe9, 0xd4, 0xa7, 0x2f, 0x9e, 0xc6, 0x3c, 0xe6, 0xd2, 0x23, 0x5d, 0xa5,
	0x8e, 0xe1, 0xdf, 0x03, 0x38, 0xb8, 0xe8, 0x62, 0xe6, 0x67, 0x68, 0xd0, 0x45, 0x90, 0x64, 0x97,
	0x49, 0x64, 0x81, 0x29, 0x98, 0x8d, 0xdd, 0x37, 0x1b, 0x81, 0x46, 0x17, 0x9d, 0xe6, 0x7d, 0x6a,
	0x04, 0xda, 0xdb, 0xad, 0x40, 0x27, 0xd7, 0x41, 0xba, 0x3a, 0xc3, 0x3b, 0x05, 0xfb, 0x23, 0x59,
	0x7a, 0x91, 0xf9, 0x16, 0xea, 0xdd, 0x1b, 0xad, 0x07, 0x12, 0xf1, 0xac, 0x11, 0x48, 0xf6, 0xad,
	0x40, 0x8f, 0xfa, 0x0c, 0x4f, 0x32, 0xec, 0x4b, 0xd1, 0x0c, 0xe0, 0x30, 0x48, 0x79, 0x95, 0x95,
	0xd6, 0x43, 0x79, 0xdc, 0xbb, 0x11, 0x48, 0xfb, 0x2b, 0xd0, 0xeb, 0x38, 0x29, 0x17, 0x55, 0xe8,
	0x50, 0x9e, 0x12, 0xca, 0x8b, 0x94, 0x17, 0xfd, 0x63, 0x5e, 0x44, 0x4b, 0x52, 0x5e, 0x5f, 0xb1,
	0xc2, 0xf1, 0xb2, 0xb2, 0x11, 0xa8, 0xcf, 0xb7, 0x02, 0x3d, 0x56, 0x78, 0xd5, 0x63, 0xbf, 0x37,
	0x4c, 0x02, 0x07, 0x19, 0xcf, 0x28, 0xb3, 0xf4, 0x29, 0x98, 0x4d, 0xdc, 0xe7, 0x8d, 0x40, 0x4a,
	0x68, 0x05, 0x9a, 0xa8, 0x88, 0x6c, 0xb1, 0xaf, 0x64, 0xf3, 0x03, 0x1c, 0x47, 0x15, 0xbb, 0x0c,
	0x57, 0x9c, 0x2e, 0xad, 0xc1, 0x14, 0xcc, 0x74, 0xf7, 0x55, 0x23, 0xd0, 0x41, 0x6c, 0x05, 0x7a,
	0xa2, 0x82, 0x7b, 0x09, 0xfb, 0x46, 0x54, 0x31, 0xb7, 0x2b, 0xcd, 0xaf, 0x50, 0x97, 0xd1, 0xa1,
	0xbc, 0xd1, 0xc7, 0x7b, 0xdd, 0x46, 0xef, 0xf9, 0xfd, 0xaf, 0x52, 0x68, 0x29, 0x9a, 0x1e, 0x04,
	0xb5, 0x35, 0x92, 0xcc, 0xf3, 0x7b, 0x31, 0x41, 0xdd, 0x0a, 0x64, 0x28, 0x60, 0x8d, 0x7d, 0x50,
	0x77, 0xa8, 0xdc, 0x32, 0x8e, 0x41, 0xe5, 0x07, 0x54, 0x8e, 0x7d, 0x90, 0x77, 0xa8, 0xc2, 0x1a,
	0x1f, 0x83, 0x2a, 0x0e, 0xa8, 0x02, 0xfb, 0xa0, 0x30, 0xcf, 0xa1, 0x91, 0xb3, 0x88, 0xb1, 0x94,
	0x45, 0x16, 0x9c, 0x82, 0x99, 0xe1, 0xa2, 0x6e, 0xe8, 0x76, 0xda, 0x61, 0xe8, 0x76, 0x0a, 0xf6,
	0xf7, 0xe6, 0xd9, 0xe4, 0xc7, 0x1a, 0x69, 0xbf, 0xd6, 0x48, 0xfb, 0xbf, 0x46, 0xc0, 0xfd, 0x72,
	0xb3, 0xb1, 0xc1, 0xed, 0xc6, 0x06, 0xff, 0x36, 0x36, 0xf8, 0xb9, 0xb5, 0xb5, 0xdb, 0xad, 0xad,
	0xfd, 0xd9, 0xda, 0xda, 0xb7, 0xf7, 0x61, 0x52, 0x86, 0x15, 0x5d, 0xb2, 0xd2, 0xe1, 0x79, 0x4c,
	0xfa, 0x1d, 0x29, 0x59, 0x90, 0x92, 0x98, 0xcf, 0x8b, 0x34, 0xc8, 0xcb, 0x79, 0xc6, 0x23, 0x46,
	0xbe, 0xab, 0xa5, 0x92, 0x9f, 0x1c, 0x0e, 0xe5, 0xae, 0xbc, 0xbb, 0x0b, 0x00, 0x00, 0xff, 0xff,
	0x08, 0xa4, 0x61, 0x0a, 0x71, 0x03, 0x00, 0x00,
}

func (this *Check) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Check)
	if !ok {
		that2, ok := that.(Check)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ChainID != that1.ChainID {
		return false
	}
	if this.Coin != that1.Coin {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if !bytes.Equal(this.Nonce, that1.Nonce) {
		return false
	}
	if this.DueBlock != that1.DueBlock {
		return false
	}
	if that1.Lock == nil {
		if this.Lock != nil {
			return false
		}
	} else if !this.Lock.Equal(*that1.Lock) {
		return false
	}
	if that1.V == nil {
		if this.V != nil {
			return false
		}
	} else if !this.V.Equal(*that1.V) {
		return false
	}
	if that1.R == nil {
		if this.R != nil {
			return false
		}
	} else if !this.R.Equal(*that1.R) {
		return false
	}
	if that1.S == nil {
		if this.S != nil {
			return false
		}
	} else if !this.S.Equal(*that1.S) {
		return false
	}
	if this.Redeemed != that1.Redeemed {
		return false
	}
	return true
}
func (m *Check) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Check) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Check) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Redeemed {
		i--
		if m.Redeemed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.S != nil {
		{
			size := m.S.Size()
			i -= size
			if _, err := m.S.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintCheck(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.R != nil {
		{
			size := m.R.Size()
			i -= size
			if _, err := m.R.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintCheck(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.V != nil {
		{
			size := m.V.Size()
			i -= size
			if _, err := m.V.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintCheck(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Lock != nil {
		{
			size := m.Lock.Size()
			i -= size
			if _, err := m.Lock.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintCheck(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DueBlock != 0 {
		i = encodeVarintCheck(dAtA, i, uint64(m.DueBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Nonce) > 0 {
		i -= len(m.Nonce)
		copy(dAtA[i:], m.Nonce)
		i = encodeVarintCheck(dAtA, i, uint64(len(m.Nonce)))
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCheck(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Coin) > 0 {
		i -= len(m.Coin)
		copy(dAtA[i:], m.Coin)
		i = encodeVarintCheck(dAtA, i, uint64(len(m.Coin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintCheck(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCheck(dAtA []byte, offset int, v uint64) int {
	offset -= sovCheck(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Check) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovCheck(uint64(l))
	}
	l = len(m.Coin)
	if l > 0 {
		n += 1 + l + sovCheck(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovCheck(uint64(l))
	l = len(m.Nonce)
	if l > 0 {
		n += 1 + l + sovCheck(uint64(l))
	}
	if m.DueBlock != 0 {
		n += 1 + sovCheck(uint64(m.DueBlock))
	}
	if m.Lock != nil {
		l = m.Lock.Size()
		n += 1 + l + sovCheck(uint64(l))
	}
	if m.V != nil {
		l = m.V.Size()
		n += 1 + l + sovCheck(uint64(l))
	}
	if m.R != nil {
		l = m.R.Size()
		n += 1 + l + sovCheck(uint64(l))
	}
	if m.S != nil {
		l = m.S.Size()
		n += 1 + l + sovCheck(uint64(l))
	}
	if m.Redeemed {
		n += 2
	}
	return n
}

func sovCheck(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCheck(x uint64) (n int) {
	return sovCheck(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Check) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheck
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
			return fmt.Errorf("proto: Check: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Check: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nonce = append(m.Nonce[:0], dAtA[iNdEx:postIndex]...)
			if m.Nonce == nil {
				m.Nonce = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DueBlock", wireType)
			}
			m.DueBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DueBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.Lock = &v
			if err := m.Lock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field V", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.V = &v
			if err := m.V.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field R", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.R = &v
			if err := m.R.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field S", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
				return ErrInvalidLengthCheck
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheck
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Int
			m.S = &v
			if err := m.S.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Redeemed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheck
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
			m.Redeemed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCheck(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheck
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
func skipCheck(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheck
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
					return 0, ErrIntOverflowCheck
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
					return 0, ErrIntOverflowCheck
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
				return 0, ErrInvalidLengthCheck
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCheck
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCheck
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCheck        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheck          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCheck = fmt.Errorf("proto: unexpected end of group")
)
