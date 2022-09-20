// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/nft/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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
	// max_collection_size defines maximum allowed count of NFT tokens per a NFT collection.
	MaxCollectionSize uint32 `protobuf:"varint,1,opt,name=max_collection_size,json=maxCollectionSize,proto3" json:"max_collection_size,omitempty"`
	// max_token_quantity defines maximum allowed count of NFT sub-tokens per a NFT token.
	MaxTokenQuantity uint32 `protobuf:"varint,2,opt,name=max_token_quantity,json=maxTokenQuantity,proto3" json:"max_token_quantity,omitempty"`
	// min_reserve_amount defines minimum allowed reserve for a NFT sub-token in the base coin.
	MinReserveAmount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=min_reserve_amount,json=minReserveAmount,proto3,customtype=cosmossdk.io/math.Int" json:"min_reserve_amount"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_eac3d5aa619f2982, []int{0}
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
	proto.RegisterType((*Params)(nil), "decimal.nft.v1.Params")
}

func init() { proto.RegisterFile("decimal/nft/v1/params.proto", fileDescriptor_eac3d5aa619f2982) }

var fileDescriptor_eac3d5aa619f2982 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0x3a, 0x41,
	0x18, 0xc6, 0x77, 0xfe, 0x7f, 0x10, 0x5a, 0x2a, 0x6c, 0x2b, 0x30, 0x83, 0x5d, 0xe9, 0x24, 0x94,
	0x33, 0x88, 0xb7, 0x6e, 0xd9, 0xc9, 0x43, 0x50, 0xd6, 0xa5, 0x2e, 0xcb, 0xb8, 0x8e, 0xdb, 0xa0,
	0x33, 0xaf, 0xed, 0xbc, 0xca, 0xea, 0x27, 0xe8, 0xd8, 0xb1, 0xa3, 0x1f, 0xa2, 0x0f, 0xe1, 0x2d,
	0xe9, 0x14, 0x1d, 0x24, 0xf4, 0xd2, 0xc7, 0x08, 0x77, 0x96, 0x6e, 0x33, 0xef, 0xef, 0xe1, 0xe1,
	0xe1, 0xe7, 0x1e, 0x77, 0x45, 0x24, 0x15, 0x1f, 0x30, 0xdd, 0x43, 0x36, 0xae, 0xb3, 0x21, 0x4f,
	0xb8, 0x32, 0x74, 0x98, 0x00, 0x82, 0xb7, 0x9b, 0x43, 0xaa, 0x7b, 0x48, 0xc7, 0xf5, 0xf2, 0x41,
	0x0c, 0x31, 0x64, 0x88, 0x6d, 0x5e, 0x36, 0x55, 0x3e, 0x8a, 0xc0, 0x28, 0x30, 0xa1, 0x05, 0xf6,
	0x63, 0xd1, 0xc9, 0x3b, 0x71, 0x0b, 0xd7, 0x59, 0xa3, 0x47, 0xdd, 0x7d, 0xc5, 0xd3, 0x30, 0x82,
	0xc1, 0x40, 0x44, 0x28, 0x41, 0x87, 0x46, 0x4e, 0x45, 0x89, 0x54, 0x48, 0x75, 0xa7, 0xbd, 0xa7,
	0x78, 0x7a, 0xf9, 0x47, 0x6e, 0xe5, 0x54, 0x78, 0x67, 0xae, 0xb7, 0xc9, 0x23, 0xf4, 0x85, 0x0e,
	0x9f, 0x46, 0x5c, 0xa3, 0xc4, 0x49, 0xe9, 0x5f, 0x16, 0x2f, 0x2a, 0x9e, 0xde, 0x6d, 0xc0, 0x4d,
	0x7e, 0xf7, 0xee, 0x5d, 0x4f, 0x49, 0x1d, 0x26, 0xc2, 0x88, 0x64, 0x2c, 0x42, 0xae, 0x60, 0xa4,
	0xb1, 0xf4, 0xbf, 0x42, 0xaa, 0x5b, 0xcd, 0xd3, 0xf9, 0x32, 0x70, 0xbe, 0x96, 0xc1, 0xa1, 0x9d,
	0x66, 0xba, 0x7d, 0x2a, 0x81, 0x29, 0x8e, 0x8f, 0xb4, 0xa5, 0xf1, 0xe3, 0xad, 0xe6, 0xe6, 0x9b,
	0x5b, 0x1a, 0xdb, 0x45, 0x25, 0x75, 0xdb, 0xb6, 0x5c, 0x64, 0x25, 0xe7, 0xdb, 0xcf, 0xb3, 0xc0,
	0x79, 0x9d, 0x05, 0xe4, 0x67, 0x16, 0x90, 0xe6, 0xd5, 0x7c, 0xe5, 0x93, 0xc5, 0xca, 0x27, 0xdf,
	0x2b, 0x9f, 0xbc, 0xac, 0x7d, 0x67, 0xb1, 0xf6, 0x9d, 0xcf, 0xb5, 0xef, 0x3c, 0x34, 0x3a, 0x12,
	0x3b, 0xa3, 0xa8, 0x2f, 0x90, 0x42, 0x12, 0xb3, 0x5c, 0x1d, 0x0a, 0xae, 0x58, 0x0c, 0x35, 0xa3,
	0x78, 0x82, 0x35, 0x0d, 0x5d, 0xc1, 0xd2, 0xcc, 0x35, 0x4e, 0x86, 0xc2, 0x74, 0x0a, 0x99, 0xa7,
	0xc6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x79, 0xb6, 0x25, 0x87, 0x01, 0x00, 0x00,
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
	if this.MaxCollectionSize != that1.MaxCollectionSize {
		return false
	}
	if this.MaxTokenQuantity != that1.MaxTokenQuantity {
		return false
	}
	if !this.MinReserveAmount.Equal(that1.MinReserveAmount) {
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
		size := m.MinReserveAmount.Size()
		i -= size
		if _, err := m.MinReserveAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.MaxTokenQuantity != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxTokenQuantity))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxCollectionSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxCollectionSize))
		i--
		dAtA[i] = 0x8
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
	if m.MaxCollectionSize != 0 {
		n += 1 + sovParams(uint64(m.MaxCollectionSize))
	}
	if m.MaxTokenQuantity != 0 {
		n += 1 + sovParams(uint64(m.MaxTokenQuantity))
	}
	l = m.MinReserveAmount.Size()
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxCollectionSize", wireType)
			}
			m.MaxCollectionSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxCollectionSize |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxTokenQuantity", wireType)
			}
			m.MaxTokenQuantity = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxTokenQuantity |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinReserveAmount", wireType)
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
			if err := m.MinReserveAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
