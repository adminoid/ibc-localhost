// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/coin/v1/params.proto

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

// Params defines the parameters for the module.
type Params struct {
	// title of the base coin
	BaseTitle string `protobuf:"bytes,1,opt,name=base_title,json=baseTitle,proto3" json:"base_title" yaml:"base_title"`
	// symbol of the base coin (denom)
	BaseSymbol string `protobuf:"bytes,2,opt,name=base_symbol,json=baseSymbol,proto3" json:"base_symbol" yaml:"base_symbol"`
	// initial volume of the base coin
	BaseInitialVolume github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=base_initial_volume,json=baseInitialVolume,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"base_initial_volume" yaml:"base_initial_volume"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fe8125aa91bb9df, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "decimal.coin.v1.Params")
}

func init() { proto.RegisterFile("decimal/coin/v1/params.proto", fileDescriptor_6fe8125aa91bb9df) }

var fileDescriptor_6fe8125aa91bb9df = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x31, 0x4b, 0x33, 0x31,
	0x18, 0xc7, 0x2f, 0x7d, 0xa1, 0xbc, 0x3d, 0x07, 0xe9, 0xe9, 0x50, 0x8a, 0x5c, 0xe4, 0x44, 0x71,
	0xe9, 0x85, 0xa2, 0x53, 0xc7, 0x0e, 0x42, 0x17, 0x91, 0x2a, 0x0e, 0x2e, 0x25, 0x77, 0x0d, 0x67,
	0x68, 0x72, 0x4f, 0xb9, 0xa4, 0xc5, 0x7e, 0x03, 0x47, 0xc1, 0x2f, 0xd0, 0x0f, 0xe1, 0x87, 0xe8,
	0xd8, 0x51, 0x1c, 0x82, 0xb4, 0x8b, 0x74, 0xec, 0x27, 0x90, 0xcb, 0x15, 0x3c, 0xc1, 0x29, 0xcf,
	0xf3, 0xfb, 0x25, 0xff, 0x3c, 0xf0, 0xb8, 0x47, 0x43, 0x16, 0x73, 0x49, 0x05, 0x89, 0x81, 0xa7,
	0x64, 0xda, 0x26, 0x63, 0x9a, 0x51, 0xa9, 0xc2, 0x71, 0x06, 0x1a, 0xbc, 0xfd, 0x9d, 0x0d, 0x73,
	0x1b, 0x4e, 0xdb, 0xcd, 0xc3, 0x04, 0x12, 0xb0, 0x8e, 0xe4, 0x55, 0x71, 0x2d, 0x78, 0xab, 0xb8,
	0xd5, 0x1b, 0xfb, 0xce, 0xeb, 0xba, 0x6e, 0x44, 0x15, 0x1b, 0x68, 0xae, 0x05, 0x6b, 0xa0, 0x63,
	0x74, 0x5e, 0xeb, 0x9e, 0x6c, 0x0c, 0x2e, 0xd1, 0xad, 0xc1, 0xf5, 0x19, 0x95, 0xa2, 0x13, 0xfc,
	0xb0, 0xa0, 0x5f, 0xcb, 0x9b, 0xbb, 0xbc, 0xf6, 0xae, 0xdc, 0x3d, 0x6b, 0xd4, 0x4c, 0x46, 0x20,
	0x1a, 0x15, 0x1b, 0x72, 0xba, 0x31, 0xb8, 0x8c, 0xb7, 0x06, 0x7b, 0xa5, 0x94, 0x02, 0x06, 0x7d,
	0xfb, 0xcf, 0xad, 0x6d, 0xbc, 0x57, 0xe4, 0x1e, 0x58, 0xc9, 0x53, 0xae, 0x39, 0x15, 0x83, 0x29,
	0x88, 0x89, 0x64, 0x8d, 0x7f, 0x36, 0x30, 0x5e, 0x18, 0xec, 0x7c, 0x18, 0x7c, 0x96, 0x70, 0xfd,
	0x38, 0x89, 0xc2, 0x18, 0x24, 0x89, 0x41, 0x49, 0x50, 0xbb, 0xa3, 0xa5, 0x86, 0x23, 0xa2, 0x67,
	0x63, 0xa6, 0xc2, 0x5e, 0xaa, 0x37, 0x06, 0xff, 0x15, 0xb6, 0x35, 0xb8, 0x59, 0x1a, 0xe3, 0xb7,
	0x0c, 0xfa, 0xf5, 0x9c, 0xf6, 0x0a, 0x78, 0x6f, 0x59, 0xe7, 0xff, 0xf3, 0x1c, 0x3b, 0x5f, 0x73,
	0x8c, 0xba, 0xd7, 0x8b, 0x95, 0x8f, 0x96, 0x2b, 0x1f, 0x7d, 0xae, 0x7c, 0xf4, 0xb2, 0xf6, 0x9d,
	0xe5, 0xda, 0x77, 0xde, 0xd7, 0xbe, 0xf3, 0x70, 0x19, 0x71, 0x1d, 0x4d, 0xe2, 0x11, 0xd3, 0x21,
	0x64, 0x09, 0xd9, 0x6d, 0x41, 0x33, 0x2a, 0x49, 0x02, 0x2d, 0x25, 0x69, 0xa6, 0x5b, 0x29, 0x0c,
	0x19, 0x79, 0x2a, 0xf6, 0x66, 0xa7, 0x8c, 0xaa, 0x76, 0x1b, 0x17, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x39, 0x56, 0x79, 0x45, 0xd4, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.BaseTitle != that1.BaseTitle {
		return false
	}
	if this.BaseSymbol != that1.BaseSymbol {
		return false
	}
	if !this.BaseInitialVolume.Equal(that1.BaseInitialVolume) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.BaseInitialVolume.Size()
		i -= size
		if _, err := m.BaseInitialVolume.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.BaseSymbol) > 0 {
		i -= len(m.BaseSymbol)
		copy(dAtA[i:], m.BaseSymbol)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseSymbol)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BaseTitle) > 0 {
		i -= len(m.BaseTitle)
		copy(dAtA[i:], m.BaseTitle)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BaseTitle)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseTitle)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.BaseSymbol)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.BaseInitialVolume.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseTitle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseTitle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseInitialVolume", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BaseInitialVolume.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
