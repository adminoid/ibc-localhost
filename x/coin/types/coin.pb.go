// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/coin/v1/coin.proto

package types

import (
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

type Coin struct {
	Title       string                                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title" yaml:"title"`
	Symbol      string                                 `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol" yaml:"symbol"`
	CRR         uint64                                 `protobuf:"varint,3,opt,name=crr,proto3" json:"constant_reserve_ratio" yaml:"constant_reserve_ratio"`
	Reserve     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=reserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserve" yaml:"reserve"`
	Volume      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=volume,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"volume" yaml:"volume"`
	LimitVolume github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=limit_volume,json=limitVolume,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"limit_volume" yaml:"limit_volume"`
	Creator     string                                 `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator" yaml:"creator"`
	Identity    string                                 `protobuf:"bytes,8,opt,name=identity,proto3" json:"identity" yaml:"identity"`
}

func (m *Coin) Reset()         { *m = Coin{} }
func (m *Coin) String() string { return proto.CompactTextString(m) }
func (*Coin) ProtoMessage()    {}
func (*Coin) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0da51b7c52dac0b, []int{0}
}
func (m *Coin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Coin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coin.Merge(m, src)
}
func (m *Coin) XXX_Size() int {
	return m.Size()
}
func (m *Coin) XXX_DiscardUnknown() {
	xxx_messageInfo_Coin.DiscardUnknown(m)
}

var xxx_messageInfo_Coin proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Coin)(nil), "decimal.coin.v1.Coin")
}

func init() { proto.RegisterFile("decimal/coin/v1/coin.proto", fileDescriptor_b0da51b7c52dac0b) }

var fileDescriptor_b0da51b7c52dac0b = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x13, 0xf7, 0x57, 0x1d, 0x57, 0x0b, 0x51, 0x24, 0x56, 0x9a, 0x57, 0xe6, 0x20, 0xbd,
	0x6c, 0x42, 0xa9, 0x20, 0xd4, 0xdb, 0xf6, 0x54, 0x04, 0x0f, 0x83, 0x7a, 0xf0, 0xb2, 0x24, 0xd9,
	0x21, 0x0e, 0xcd, 0x64, 0xca, 0x64, 0x76, 0xe9, 0xfe, 0x07, 0xe2, 0xc9, 0xa3, 0xc7, 0xfd, 0x73,
	0x7a, 0xec, 0x51, 0x3c, 0x0c, 0xb2, 0x7b, 0x91, 0x1c, 0xf3, 0x17, 0xc8, 0xce, 0xcc, 0xb6, 0x0a,
	0xbd, 0xf4, 0x34, 0x79, 0x9f, 0x6f, 0xde, 0xfb, 0x84, 0xc7, 0x04, 0xed, 0x4d, 0x69, 0xce, 0x78,
	0x5a, 0x26, 0xb9, 0x60, 0x55, 0x32, 0x3f, 0x32, 0x67, 0x7c, 0x21, 0x85, 0x12, 0xc1, 0xae, 0xcb,
	0x62, 0xc3, 0xe6, 0x47, 0x7b, 0xcf, 0x0a, 0x51, 0x08, 0x93, 0x25, 0x9b, 0x27, 0xfb, 0x1a, 0xfe,
	0xd6, 0x43, 0xdd, 0x53, 0xc1, 0xaa, 0x20, 0x41, 0x3d, 0xc5, 0x54, 0x49, 0x43, 0xff, 0xc0, 0x3f,
	0x7c, 0x38, 0x7e, 0xd1, 0x68, 0xb0, 0xa0, 0xd5, 0x30, 0x5c, 0xa4, 0xbc, 0x3c, 0xc1, 0xa6, 0xc4,
	0xc4, 0xe2, 0xe0, 0x18, 0xf5, 0xeb, 0x05, 0xcf, 0x44, 0x19, 0x3e, 0x30, 0x1d, 0x2f, 0x1b, 0x0d,
	0x8e, 0xb4, 0x1a, 0x1e, 0xdb, 0x16, 0x5b, 0x63, 0xe2, 0x82, 0xe0, 0x03, 0xea, 0xe4, 0x52, 0x86,
	0x9d, 0x03, 0xff, 0xb0, 0x3b, 0x1e, 0xaf, 0x34, 0x74, 0x4e, 0x09, 0x69, 0x34, 0x3c, 0xcf, 0x45,
	0x55, 0xab, 0xb4, 0x52, 0x13, 0x49, 0x6b, 0x2a, 0xe7, 0x74, 0x22, 0x53, 0xc5, 0x44, 0xab, 0x61,
	0xdf, 0x0e, 0xba, 0x3b, 0xc7, 0x64, 0x33, 0x2e, 0xa0, 0x68, 0xe0, 0x70, 0xd8, 0x35, 0xdf, 0xf2,
	0xee, 0x4a, 0x83, 0xf7, 0x4b, 0xc3, 0xab, 0x82, 0xa9, 0x2f, 0xb3, 0x2c, 0xce, 0x05, 0x4f, 0x72,
	0x51, 0x73, 0x51, 0xbb, 0x63, 0x54, 0x4f, 0xcf, 0x13, 0xb5, 0xb8, 0xa0, 0x75, 0x7c, 0x56, 0xa9,
	0x46, 0xc3, 0x76, 0x40, 0xab, 0xe1, 0x89, 0x35, 0x3a, 0x80, 0xc9, 0x36, 0x0a, 0x52, 0xd4, 0x9f,
	0x8b, 0x72, 0xc6, 0x69, 0xd8, 0x33, 0x96, 0xb3, 0x7b, 0x5b, 0x5c, 0xff, 0xed, 0x7e, 0x6c, 0x8d,
	0x89, 0x0b, 0x82, 0x4b, 0x34, 0x2c, 0x19, 0x67, 0x6a, 0xe2, 0x44, 0x7d, 0x23, 0xfa, 0x78, 0x6f,
	0xd1, 0x7f, 0x53, 0x5a, 0x0d, 0x4f, 0xad, 0xee, 0x5f, 0x8a, 0xc9, 0x23, 0x53, 0x7e, 0xb2, 0xe6,
	0x37, 0x68, 0x90, 0x4b, 0x9a, 0x2a, 0x21, 0xc3, 0x81, 0x91, 0xee, 0x6f, 0xb6, 0xe2, 0xd0, 0xed,
	0x56, 0x1c, 0xc0, 0x64, 0x1b, 0x05, 0x6f, 0xd1, 0x0e, 0x9b, 0xd2, 0x4a, 0x31, 0xb5, 0x08, 0x77,
	0x4c, 0x27, 0x34, 0x1a, 0x6e, 0x58, 0xab, 0x61, 0xd7, 0xb6, 0x6e, 0x09, 0x26, 0x37, 0xe1, 0xc9,
	0xf0, 0xeb, 0x12, 0xbc, 0x1f, 0x4b, 0xf0, 0xff, 0x2c, 0xc1, 0x1f, 0xbf, 0xbf, 0x5a, 0x45, 0xfe,
	0xf5, 0x2a, 0xf2, 0x7f, 0xaf, 0x22, 0xff, 0xfb, 0x3a, 0xf2, 0xae, 0xd7, 0x91, 0xf7, 0x73, 0x1d,
	0x79, 0x9f, 0x5f, 0x67, 0x4c, 0x65, 0xb3, 0xfc, 0x9c, 0xaa, 0x58, 0xc8, 0x22, 0x71, 0x77, 0x5b,
	0xd1, 0x94, 0x27, 0x85, 0x18, 0xd5, 0x3c, 0x95, 0x6a, 0x54, 0x89, 0x29, 0x4d, 0x2e, 0xed, 0xbf,
	0x60, 0x76, 0x91, 0xf5, 0xcd, 0x1d, 0x3f, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x33, 0x46, 0xb3,
	0xc1, 0x28, 0x03, 0x00, 0x00,
}

func (this *Coin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Coin)
	if !ok {
		that2, ok := that.(Coin)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.CRR != that1.CRR {
		return false
	}
	if !this.Reserve.Equal(that1.Reserve) {
		return false
	}
	if !this.Volume.Equal(that1.Volume) {
		return false
	}
	if !this.LimitVolume.Equal(that1.LimitVolume) {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if this.Identity != that1.Identity {
		return false
	}
	return true
}
func (m *Coin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Coin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintCoin(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintCoin(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.LimitVolume.Size()
		i -= size
		if _, err := m.LimitVolume.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.Volume.Size()
		i -= size
		if _, err := m.Volume.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.Reserve.Size()
		i -= size
		if _, err := m.Reserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.CRR != 0 {
		i = encodeVarintCoin(dAtA, i, uint64(m.CRR))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintCoin(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintCoin(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoin(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Coin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	if m.CRR != 0 {
		n += 1 + sovCoin(uint64(m.CRR))
	}
	l = m.Reserve.Size()
	n += 1 + l + sovCoin(uint64(l))
	l = m.Volume.Size()
	n += 1 + l + sovCoin(uint64(l))
	l = m.LimitVolume.Size()
	n += 1 + l + sovCoin(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	return n
}

func sovCoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoin(x uint64) (n int) {
	return sovCoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Coin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoin
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
			return fmt.Errorf("proto: Coin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CRR", wireType)
			}
			m.CRR = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CRR |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Volume.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitVolume", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LimitVolume.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoin
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
func skipCoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoin
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
					return 0, ErrIntOverflowCoin
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
					return 0, ErrIntOverflowCoin
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
				return 0, ErrInvalidLengthCoin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoin = fmt.Errorf("proto: unexpected end of group")
)
