// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/coin/v1/coin.proto

package types

import (
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

type Coin struct {
	Title                string                                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ConstantReserveRatio uint64                                 `protobuf:"varint,2,opt,name=constant_reserve_ratio,json=constantReserveRatio,proto3" json:"constant_reserve_ratio,omitempty"`
	Symbol               string                                 `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Reserve              github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=reserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserve"`
	LimitVolume          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=limit_volume,json=limitVolume,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"limit_volume"`
	Volume               github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=volume,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"volume"`
	Creator              string                                 `protobuf:"bytes,7,opt,name=creator,proto3" json:"creator,omitempty"`
	Identity             string                                 `protobuf:"bytes,8,opt,name=identity,proto3" json:"identity,omitempty"`
}

func (m *Coin) Reset()      { *m = Coin{} }
func (*Coin) ProtoMessage() {}
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
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x3f, 0x0b, 0x1a, 0x31,
	0x18, 0x87, 0x2f, 0x55, 0x4f, 0x9b, 0x0a, 0x85, 0x20, 0x12, 0x1c, 0xee, 0xa4, 0x43, 0x71, 0xf1,
	0x82, 0xd4, 0xa9, 0xa3, 0x85, 0xd2, 0x2e, 0x85, 0xde, 0xd0, 0xa1, 0x8b, 0xdc, 0x9f, 0x70, 0x0d,
	0x5e, 0xee, 0x95, 0xe4, 0xf5, 0xa8, 0xdf, 0xa0, 0x63, 0xa1, 0x4b, 0x47, 0x3f, 0x8e, 0xa3, 0x63,
	0xe9, 0x20, 0x45, 0x97, 0x7e, 0x8c, 0x62, 0xee, 0xec, 0x07, 0x70, 0x7a, 0xf3, 0xcb, 0x93, 0xf7,
	0x21, 0xe1, 0x0d, 0x9d, 0xe4, 0x32, 0x53, 0x3a, 0x29, 0x45, 0x06, 0xaa, 0x12, 0xf5, 0xc2, 0xd5,
	0x68, 0x6b, 0x00, 0x81, 0x3d, 0x6f, 0x59, 0xe4, 0xf6, 0xea, 0xc5, 0x64, 0x54, 0x40, 0x01, 0x8e,
	0x89, 0xdb, 0xaa, 0x39, 0xf6, 0xe2, 0x47, 0x87, 0x76, 0xdf, 0x80, 0xaa, 0xd8, 0x88, 0xf6, 0x50,
	0x61, 0x29, 0x39, 0x99, 0x92, 0xd9, 0xd3, 0xb8, 0x09, 0x6c, 0x49, 0xc7, 0x19, 0x54, 0x16, 0x93,
	0x0a, 0xd7, 0x46, 0x5a, 0x69, 0x6a, 0xb9, 0x36, 0x09, 0x2a, 0xe0, 0x4f, 0xa6, 0x64, 0xd6, 0x8d,
	0x47, 0x77, 0x1a, 0x37, 0x30, 0xbe, 0x31, 0x36, 0xa6, 0xbe, 0xdd, 0xeb, 0x14, 0x4a, 0xde, 0x71,
	0xb2, 0x36, 0xb1, 0x77, 0xb4, 0xdf, 0x4a, 0x78, 0x77, 0x4a, 0x66, 0xc3, 0x55, 0x74, 0x3c, 0x87,
	0xde, 0xef, 0x73, 0xf8, 0xb2, 0x50, 0xf8, 0x65, 0x97, 0x46, 0x19, 0x68, 0x91, 0x81, 0xd5, 0x60,
	0xdb, 0x32, 0xb7, 0xf9, 0x46, 0xe0, 0x7e, 0x2b, 0x6d, 0xf4, 0xbe, 0xc2, 0xf8, 0xde, 0xce, 0x3e,
	0xd2, 0x61, 0xa9, 0xb4, 0xc2, 0x75, 0x0d, 0xe5, 0x4e, 0x4b, 0xde, 0x7b, 0x48, 0xf7, 0xcc, 0x39,
	0x3e, 0x39, 0x05, 0x7b, 0x4b, 0xfd, 0x56, 0xe6, 0x3f, 0x24, 0x6b, 0xbb, 0x19, 0xa7, 0xfd, 0xcc,
	0xc8, 0x04, 0xc1, 0xf0, 0xbe, 0x7b, 0xfd, 0x3d, 0xb2, 0x09, 0x1d, 0xa8, 0x5c, 0x56, 0xa8, 0x70,
	0xcf, 0x07, 0x0e, 0xfd, 0xcf, 0xaf, 0x87, 0xdf, 0x0e, 0xa1, 0xf7, 0xf3, 0x10, 0x7a, 0x7f, 0x0f,
	0x21, 0x59, 0x7d, 0x38, 0x5e, 0x02, 0x72, 0xba, 0x04, 0xe4, 0xcf, 0x25, 0x20, 0xdf, 0xaf, 0x81,
	0x77, 0xba, 0x06, 0xde, 0xaf, 0x6b, 0xe0, 0x7d, 0x5e, 0xa6, 0x0a, 0xd3, 0x5d, 0xb6, 0x91, 0x18,
	0x81, 0x29, 0x44, 0x3b, 0x64, 0x94, 0x89, 0x16, 0x05, 0xcc, 0xad, 0x4e, 0x0c, 0xce, 0x2b, 0xc8,
	0xa5, 0xf8, 0xda, 0x7c, 0x0a, 0x77, 0xbf, 0xd4, 0x77, 0xc3, 0x7e, 0xf5, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0xf5, 0x9f, 0x18, 0xdb, 0x31, 0x02, 0x00, 0x00,
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
	if this.ConstantReserveRatio != that1.ConstantReserveRatio {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if !this.Reserve.Equal(that1.Reserve) {
		return false
	}
	if !this.LimitVolume.Equal(that1.LimitVolume) {
		return false
	}
	if !this.Volume.Equal(that1.Volume) {
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
		size := m.Volume.Size()
		i -= size
		if _, err := m.Volume.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintCoin(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.LimitVolume.Size()
		i -= size
		if _, err := m.LimitVolume.MarshalTo(dAtA[i:]); err != nil {
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
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintCoin(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ConstantReserveRatio != 0 {
		i = encodeVarintCoin(dAtA, i, uint64(m.ConstantReserveRatio))
		i--
		dAtA[i] = 0x10
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
	if m.ConstantReserveRatio != 0 {
		n += 1 + sovCoin(uint64(m.ConstantReserveRatio))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovCoin(uint64(l))
	}
	l = m.Reserve.Size()
	n += 1 + l + sovCoin(uint64(l))
	l = m.LimitVolume.Size()
	n += 1 + l + sovCoin(uint64(l))
	l = m.Volume.Size()
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConstantReserveRatio", wireType)
			}
			m.ConstantReserveRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConstantReserveRatio |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
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
				return fmt.Errorf("proto: wrong wireType = %d for field LimitVolume", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Volume", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoin
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
				return ErrInvalidLengthCoin
			}
			postIndex := iNdEx + byteLen
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
