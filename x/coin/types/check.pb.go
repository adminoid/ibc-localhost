// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/coin/v1/check.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/gogo/protobuf/gogoproto"
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
	Lock     []byte                                  `protobuf:"bytes,6,opt,name=lock,proto3" json:"lock" yaml:"lock"`
	V        *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=v,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"v" yaml:"v"`
	R        *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,8,opt,name=r,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"r" yaml:"r"`
	S        *github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,9,opt,name=s,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"s" yaml:"s"`
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
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbf, 0x8e, 0xd4, 0x30,
	0x10, 0xc6, 0x63, 0xd8, 0xbf, 0x61, 0xd1, 0xa1, 0x08, 0x89, 0x00, 0x52, 0xbc, 0xb8, 0x40, 0x2b,
	0xd0, 0xc6, 0x3a, 0x41, 0x75, 0x48, 0x14, 0x39, 0x28, 0xd2, 0x50, 0xa4, 0xa4, 0x39, 0x25, 0x8e,
	0x95, 0x8d, 0x76, 0x13, 0x9f, 0x62, 0x27, 0xe2, 0xde, 0x80, 0x92, 0x92, 0x72, 0x1f, 0xe7, 0xca,
	0x2b, 0x11, 0x85, 0x85, 0xb2, 0x0d, 0x0a, 0x5d, 0x9e, 0x00, 0xd9, 0x0e, 0xb7, 0x35, 0x57, 0x65,
	0xe6, 0xfb, 0xf2, 0xfd, 0x46, 0x1e, 0x8d, 0xfd, 0x3c, 0xa5, 0x24, 0x2f, 0xe2, 0x1d, 0x26, 0x2c,
	0x2f, 0x71, 0x73, 0x8a, 0xc9, 0x86, 0x92, 0xad, 0x7f, 0x59, 0x31, 0xc1, 0x9c, 0x93, 0xc1, 0xf4,
	0x95, 0xe9, 0x37, 0xa7, 0xcf, 0x1e, 0x67, 0x2c, 0x63, 0xda, 0xc3, 0xaa, 0x32, 0xbf, 0xa1, 0x3f,
	0x23, 0x7b, 0x7c, 0xae, 0x62, 0xce, 0x47, 0x7b, 0x46, 0x36, 0x71, 0x5e, 0x5e, 0xe4, 0xa9, 0x0b,
	0x96, 0x60, 0x35, 0x0f, 0x5e, 0xb5, 0x12, 0x4e, 0xcf, 0x95, 0x16, 0x7e, 0xe8, 0x24, 0xbc, 0xb5,
	0x7b, 0x09, 0x4f, 0xae, 0xe2, 0x62, 0x77, 0x86, 0xfe, 0x29, 0x28, 0x9a, 0xea, 0x32, 0x4c, 0x9d,
	0xd7, 0xf6, 0x48, 0x4d, 0x74, 0xef, 0x69, 0xc4, 0x93, 0x4e, 0x42, 0xdd, 0xf7, 0x12, 0x3e, 0x18,
	0x32, 0x2c, 0x2f, 0x51, 0xa4, 0x45, 0x27, 0xb6, 0x27, 0x71, 0xc1, 0xea, 0x52, 0xb8, 0xf7, 0xf5,
	0xef, 0xe1, 0xb5, 0x84, 0xd6, 0x4f, 0x09, 0x5f, 0x66, 0xb9, 0xd8, 0xd4, 0x89, 0x4f, 0x58, 0x81,
	0x09, 0xe3, 0x05, 0xe3, 0xc3, 0x67, 0xcd, 0xd3, 0x2d, 0x16, 0x57, 0x97, 0x94, 0xfb, 0x61, 0x29,
	0x3a, 0x09, 0x87, 0x7c, 0x2f, 0xe1, 0x43, 0x83, 0x37, 0x3d, 0x8a, 0x06, 0xc3, 0xc1, 0xf6, 0xb8,
	0x64, 0x25, 0xa1, 0xee, 0x68, 0x09, 0x56, 0x8b, 0xe0, 0x69, 0x27, 0xa1, 0x11, 0x7a, 0x09, 0x17,
	0x26, 0xa2, 0x5b, 0x14, 0x19, 0xd9, 0x79, 0x6f, 0xcf, 0xd3, 0x9a, 0x5e, 0x24, 0x3b, 0x46, 0xb6,
	0xee, 0x78, 0x09, 0x56, 0xa3, 0xe0, 0x45, 0x27, 0xe1, 0x51, 0xec, 0x25, 0x7c, 0x64, 0x82, 0xb7,
	0x12, 0x8a, 0x66, 0x69, 0x4d, 0x03, 0x55, 0xaa, 0x05, 0xe8, 0xe8, 0x44, 0xcf, 0xd3, 0x0b, 0x18,
	0x52, 0xc3, 0x02, 0x4c, 0x40, 0x8b, 0x4e, 0x68, 0x83, 0xc6, 0x9d, 0xea, 0xb7, 0xbf, 0xfb, 0xaf,
	0x77, 0x83, 0xa6, 0x97, 0x70, 0x66, 0x80, 0x0d, 0x8a, 0x40, 0xa3, 0x50, 0x95, 0x3b, 0xbb, 0x0b,
	0xaa, 0x3a, 0xa2, 0x2a, 0x14, 0x81, 0x4a, 0xa1, 0xb8, 0x3b, 0xbf, 0x0b, 0x8a, 0x1f, 0x51, 0x1c,
	0x45, 0x80, 0x9f, 0x2d, 0xbe, 0xee, 0xa1, 0xf5, 0x7d, 0x0f, 0xad, 0xdf, 0x7b, 0x08, 0x82, 0x4f,
	0xd7, 0xad, 0x07, 0x6e, 0x5a, 0x0f, 0xfc, 0x6a, 0x3d, 0xf0, 0xed, 0xe0, 0x59, 0x37, 0x07, 0xcf,
	0xfa, 0x71, 0xf0, 0xac, 0xcf, 0x6f, 0x93, 0x5c, 0x24, 0x35, 0xd9, 0x52, 0xe1, 0xb3, 0x2a, 0xc3,
	0xc3, 0xf1, 0x0a, 0x1a, 0x17, 0x38, 0x63, 0x6b, 0x5e, 0xc4, 0x95, 0x58, 0x97, 0x2c, 0xa5, 0xf8,
	0x8b, 0xb9, 0x76, 0x3d, 0x35, 0x99, 0xe8, 0x23, 0x7e, 0xf3, 0x37, 0x00, 0x00, 0xff, 0xff, 0x32,
	0xe2, 0xed, 0x55, 0x0a, 0x03, 0x00, 0x00,
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
	if !bytes.Equal(this.Lock, that1.Lock) {
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
	if len(m.Lock) > 0 {
		i -= len(m.Lock)
		copy(dAtA[i:], m.Lock)
		i = encodeVarintCheck(dAtA, i, uint64(len(m.Lock)))
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
	l = len(m.Lock)
	if l > 0 {
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
			m.Lock = append(m.Lock[:0], dAtA[iNdEx:postIndex]...)
			if m.Lock == nil {
				m.Lock = []byte{}
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
