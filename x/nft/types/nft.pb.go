// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/nft/v1/nft.proto

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

type BaseNFT struct {
	ID        string                                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"id" yaml:"id"`
	Owners    TokenOwners                            `protobuf:"bytes,2,rep,name=owners,proto3,castrepeated=TokenOwners" json:"owners" yaml:"owners"`
	Creator   string                                 `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	TokenURI  string                                 `protobuf:"bytes,4,opt,name=tokenURI,proto3" json:"token_uri" yaml:"token_uri"`
	Reserve   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=reserve,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"reserve" yaml:"reserve"`
	AllowMint bool                                   `protobuf:"varint,6,opt,name=allowMint,proto3" json:"allowMint,omitempty" yaml:"allow_mint"`
}

func (m *BaseNFT) Reset()      { *m = BaseNFT{} }
func (*BaseNFT) ProtoMessage() {}
func (*BaseNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_48e340f93e0c6acc, []int{0}
}
func (m *BaseNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseNFT.Merge(m, src)
}
func (m *BaseNFT) XXX_Size() int {
	return m.Size()
}
func (m *BaseNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseNFT.DiscardUnknown(m)
}

var xxx_messageInfo_BaseNFT proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseNFT)(nil), "decimal.nft.v1.BaseNFT")
}

func init() { proto.RegisterFile("decimal/nft/v1/nft.proto", fileDescriptor_48e340f93e0c6acc) }

var fileDescriptor_48e340f93e0c6acc = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x41, 0x6b, 0x13, 0x41,
	0x1c, 0xc5, 0x67, 0x13, 0x4d, 0x9a, 0x29, 0x96, 0x3a, 0x2a, 0x2c, 0x39, 0xec, 0xc4, 0x51, 0x24,
	0x07, 0xb3, 0x43, 0xcd, 0xad, 0x20, 0xc8, 0x52, 0x84, 0x1c, 0xaa, 0xb0, 0x54, 0x90, 0x5e, 0xca,
	0x66, 0x77, 0x5c, 0x87, 0xec, 0xee, 0x94, 0x99, 0x49, 0x6a, 0x6f, 0x1e, 0x3d, 0x7a, 0xec, 0x31,
	0x67, 0x3f, 0x49, 0x8f, 0x3d, 0x8a, 0xc8, 0x28, 0xc9, 0x45, 0x7a, 0xdc, 0x4f, 0x20, 0x3b, 0xbb,
	0x6d, 0xd0, 0xd3, 0x0c, 0xff, 0xf7, 0x7e, 0x8f, 0x07, 0x0f, 0xba, 0x09, 0x8b, 0x79, 0x1e, 0x65,
	0xb4, 0xf8, 0xa0, 0xe9, 0x62, 0xaf, 0x7a, 0xfc, 0x53, 0x29, 0xb4, 0x40, 0x3b, 0x8d, 0xe2, 0x57,
	0xa7, 0xc5, 0x5e, 0xff, 0x61, 0x2a, 0x52, 0x61, 0x25, 0x5a, 0xfd, 0x6a, 0x57, 0xbf, 0xff, 0x1f,
	0x2f, 0xce, 0x0a, 0x26, 0x6b, 0x8d, 0x5c, 0xb4, 0x61, 0x37, 0x88, 0x14, 0x7b, 0xf3, 0xfa, 0x08,
	0x3d, 0x81, 0xad, 0xc9, 0x81, 0xeb, 0x0c, 0x9c, 0x61, 0x2f, 0x78, 0x70, 0x6d, 0x70, 0x8b, 0x27,
	0xa5, 0xc1, 0xbd, 0xf3, 0x28, 0xcf, 0xf6, 0x09, 0x4f, 0x48, 0xd8, 0x9a, 0x1c, 0xa0, 0xf7, 0xb0,
	0x63, 0x79, 0xe5, 0xb6, 0x06, 0xed, 0xe1, 0xf6, 0x8b, 0xbe, 0xff, 0x6f, 0x07, 0xff, 0x48, 0xcc,
	0x58, 0xf1, 0xb6, 0xb2, 0x04, 0x4f, 0x2f, 0x0d, 0x06, 0xa5, 0xc1, 0xf7, 0xea, 0x88, 0x9a, 0x23,
	0xdf, 0x7e, 0xe1, 0xed, 0x8d, 0x49, 0x85, 0x4d, 0x1e, 0x7a, 0x0e, 0xbb, 0xb1, 0x64, 0x91, 0x16,
	0xd2, 0x6d, 0xdb, 0x0e, 0xa8, 0x34, 0x78, 0xa7, 0x46, 0x1b, 0x81, 0x84, 0x37, 0x16, 0xf4, 0x12,
	0x6e, 0xe9, 0x2a, 0xe4, 0x5d, 0x38, 0x71, 0xef, 0x58, 0xfb, 0xe3, 0x6b, 0x83, 0x7b, 0xf6, 0x76,
	0x32, 0x97, 0xbc, 0x34, 0x78, 0xb7, 0x66, 0x6f, 0x4f, 0x24, 0xbc, 0x45, 0xd0, 0x31, 0xec, 0x4a,
	0xa6, 0x98, 0x5c, 0x30, 0xf7, 0xae, 0xa5, 0x5f, 0x55, 0x5d, 0x7f, 0x18, 0xfc, 0x2c, 0xe5, 0xfa,
	0xe3, 0x7c, 0xea, 0xc7, 0x22, 0xa7, 0xb1, 0x50, 0xb9, 0x50, 0xcd, 0x33, 0x52, 0xc9, 0x8c, 0xea,
	0xf3, 0x53, 0xa6, 0xfc, 0x49, 0xa1, 0x37, 0xd5, 0x9a, 0x18, 0x12, 0xde, 0x04, 0xa2, 0x31, 0xec,
	0x45, 0x59, 0x26, 0xce, 0x0e, 0x79, 0xa1, 0xdd, 0xce, 0xc0, 0x19, 0x6e, 0x05, 0x8f, 0x4a, 0x83,
	0xef, 0xd7, 0x7e, 0x2b, 0x9d, 0xe4, 0xbc, 0xd0, 0x24, 0xdc, 0xf8, 0xf6, 0x77, 0xbf, 0x2c, 0x31,
	0xb8, 0x58, 0x62, 0xf0, 0x67, 0x89, 0xc1, 0xe7, 0x9f, 0x03, 0x10, 0x1c, 0x5e, 0xae, 0x3c, 0xe7,
	0x6a, 0xe5, 0x39, 0xbf, 0x57, 0x9e, 0xf3, 0x75, 0xed, 0x81, 0xab, 0xb5, 0x07, 0xbe, 0xaf, 0x3d,
	0x70, 0x3c, 0x9e, 0x72, 0x3d, 0x9d, 0xc7, 0x33, 0xa6, 0x7d, 0x21, 0x53, 0xda, 0x0c, 0xa0, 0x59,
	0x94, 0xd3, 0x54, 0x8c, 0x54, 0x1e, 0x49, 0x3d, 0x2a, 0x44, 0xc2, 0xe8, 0x27, 0x3b, 0xb9, 0x2d,
	0x3d, 0xed, 0xd8, 0xc1, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xaa, 0x8e, 0x75, 0xd9, 0x4e,
	0x02, 0x00, 0x00,
}

func (m *BaseNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowMint {
		i--
		if m.AllowMint {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Reserve.Size()
		i -= size
		if _, err := m.Reserve.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.TokenURI) > 0 {
		i -= len(m.TokenURI)
		copy(dAtA[i:], m.TokenURI)
		i = encodeVarintNft(dAtA, i, uint64(len(m.TokenURI)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Owners[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNft(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintNft(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, e := range m.Owners {
			l = e.Size()
			n += 1 + l + sovNft(uint64(l))
		}
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.TokenURI)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = m.Reserve.Size()
	n += 1 + l + sovNft(uint64(l))
	if m.AllowMint {
		n += 2
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
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
			return fmt.Errorf("proto: BaseNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owners = append(m.Owners, TokenOwner{})
			if err := m.Owners[len(m.Owners)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenURI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenURI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserve", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserve.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowMint", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
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
			m.AllowMint = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
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
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
					return 0, ErrIntOverflowNft
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
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)
