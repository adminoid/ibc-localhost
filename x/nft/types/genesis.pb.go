// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/nft/v1/genesis.proto

package types

import (
	fmt "fmt"
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

type GenesisState struct {
	Collections []Collection         `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections" yaml:"collections"`
	NFTs        []BaseNFT            `protobuf:"bytes,2,rep,name=nfts,proto3" json:"nfts" yaml:"nfts"`
	SubTokens   map[string]SubTokens `protobuf:"bytes,3,rep,name=subTokens,proto3" json:"subTokens" yaml:"sub_tokens" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbce1c13afacb19, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

type SubTokens struct {
	SubTokens []SubToken `protobuf:"bytes,1,rep,name=subTokens,proto3" json:"subTokens" yaml:"sub_tokens"`
}

func (m *SubTokens) Reset()         { *m = SubTokens{} }
func (m *SubTokens) String() string { return proto.CompactTextString(m) }
func (*SubTokens) ProtoMessage()    {}
func (*SubTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_5dbce1c13afacb19, []int{1}
}
func (m *SubTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubTokens.Merge(m, src)
}
func (m *SubTokens) XXX_Size() int {
	return m.Size()
}
func (m *SubTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_SubTokens.DiscardUnknown(m)
}

var xxx_messageInfo_SubTokens proto.InternalMessageInfo

func (m *SubTokens) GetSubTokens() []SubToken {
	if m != nil {
		return m.SubTokens
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "decimal.nft.v1.GenesisState")
	proto.RegisterMapType((map[string]SubTokens)(nil), "decimal.nft.v1.GenesisState.SubTokensEntry")
	proto.RegisterType((*SubTokens)(nil), "decimal.nft.v1.SubTokens")
}

func init() { proto.RegisterFile("decimal/nft/v1/genesis.proto", fileDescriptor_5dbce1c13afacb19) }

var fileDescriptor_5dbce1c13afacb19 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0xaa, 0xd3, 0x40,
	0x18, 0x86, 0x33, 0xe9, 0x51, 0xec, 0x44, 0x0e, 0x3a, 0x08, 0xe6, 0x44, 0x49, 0x4a, 0x56, 0x05,
	0x39, 0x19, 0x4e, 0xbb, 0x91, 0x2e, 0x23, 0x56, 0x37, 0x16, 0x4c, 0x0b, 0x8a, 0x0b, 0x25, 0x49,
	0xa7, 0x21, 0x34, 0xc9, 0x94, 0xcc, 0xa4, 0x98, 0x3b, 0x70, 0xe9, 0x25, 0xf4, 0x72, 0xba, 0xec,
	0xd2, 0x55, 0x90, 0x74, 0xe3, 0xba, 0x5b, 0x37, 0x92, 0x1f, 0xd3, 0x34, 0x20, 0xb8, 0x1b, 0x78,
	0x9f, 0xef, 0x99, 0xef, 0x1d, 0x06, 0x3e, 0x5f, 0x12, 0xd7, 0x0f, 0xed, 0x00, 0x47, 0x2b, 0x8e,
	0xb7, 0x77, 0xd8, 0x23, 0x11, 0x61, 0x3e, 0x33, 0x36, 0x31, 0xe5, 0x14, 0x5d, 0xd7, 0xa9, 0x11,
	0xad, 0xb8, 0xb1, 0xbd, 0x53, 0x9e, 0x78, 0xd4, 0xa3, 0x65, 0x84, 0x8b, 0x53, 0x45, 0x29, 0x72,
	0xc7, 0x51, 0xc0, 0x55, 0xa2, 0x75, 0x12, 0x97, 0x06, 0x01, 0x71, 0xb9, 0x4f, 0xa3, 0x1a, 0x50,
	0x3b, 0x00, 0x4b, 0x9c, 0x2f, 0x9c, 0xae, 0x49, 0x9d, 0xeb, 0xbf, 0x45, 0xf8, 0xf0, 0x4d, 0xb5,
	0xd2, 0x9c, 0xdb, 0x9c, 0xa0, 0x8f, 0x50, 0x3a, 0x4b, 0x98, 0x0c, 0x06, 0xbd, 0xa1, 0x34, 0x52,
	0x8c, 0xcb, 0x3d, 0x8d, 0x57, 0x0d, 0x62, 0x2a, 0xfb, 0x4c, 0x13, 0x4e, 0x99, 0x86, 0x52, 0x3b,
	0x0c, 0x26, 0x7a, 0x6b, 0x58, 0xb7, 0xda, 0x2a, 0xf4, 0x16, 0x5e, 0x45, 0x2b, 0xce, 0x64, 0xb1,
	0x54, 0x3e, 0xed, 0x2a, 0x4d, 0x9b, 0x91, 0xd9, 0x74, 0x61, 0x3e, 0x2b, 0x7c, 0x79, 0xa6, 0x5d,
	0xcd, 0xa6, 0x0b, 0x76, 0xca, 0x34, 0xa9, 0xf2, 0x16, 0xa3, 0xba, 0x55, 0x1a, 0xd0, 0x12, 0xf6,
	0x59, 0xe2, 0x2c, 0x8a, 0x1a, 0x4c, 0xee, 0x95, 0xba, 0x17, 0x5d, 0x5d, 0xbb, 0x94, 0x31, 0xff,
	0x4b, 0xbf, 0x8e, 0x78, 0x9c, 0x9a, 0x37, 0xf5, 0xca, 0x8f, 0x2b, 0x75, 0xf3, 0x26, 0x4c, 0xb7,
	0xce, 0x62, 0xe5, 0x03, 0xbc, 0xbe, 0x9c, 0x43, 0x8f, 0x60, 0x6f, 0x4d, 0x52, 0x19, 0x0c, 0xc0,
	0xb0, 0x6f, 0x15, 0x47, 0x84, 0xe1, 0xbd, 0xad, 0x1d, 0x24, 0x44, 0x16, 0x07, 0x60, 0x28, 0x8d,
	0x6e, 0xba, 0x5b, 0x34, 0x02, 0xab, 0xe2, 0x26, 0xe2, 0x4b, 0x30, 0x79, 0xf0, 0x6d, 0xa7, 0x09,
	0xbf, 0x76, 0x9a, 0xa0, 0x7f, 0x86, 0xfd, 0x86, 0x40, 0xef, 0xdb, 0xad, 0xaa, 0x77, 0x97, 0xff,
	0xe5, 0xfb, 0xbf, 0x0a, 0xe6, 0xbb, 0x7d, 0xae, 0x82, 0x43, 0xae, 0x82, 0x9f, 0xb9, 0x0a, 0xbe,
	0x1f, 0x55, 0xe1, 0x70, 0x54, 0x85, 0x1f, 0x47, 0x55, 0xf8, 0x34, 0x76, 0x7c, 0xee, 0x24, 0xee,
	0x9a, 0x70, 0x83, 0xc6, 0x1e, 0xae, 0xaf, 0xe1, 0xc4, 0x0e, 0xb1, 0x47, 0x6f, 0x59, 0x68, 0xc7,
	0xfc, 0x36, 0xa2, 0x4b, 0x82, 0xbf, 0x96, 0x3f, 0x87, 0xa7, 0x1b, 0xc2, 0x9c, 0xfb, 0xe5, 0x9f,
	0x19, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xe6, 0xdd, 0x09, 0xd4, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubTokens) > 0 {
		for k := range m.SubTokens {
			v := m.SubTokens[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGenesis(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGenesis(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.NFTs) > 0 {
		for iNdEx := len(m.NFTs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NFTs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Collections) > 0 {
		for iNdEx := len(m.Collections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SubTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubTokens) > 0 {
		for iNdEx := len(m.SubTokens) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubTokens[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Collections) > 0 {
		for _, e := range m.Collections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.NFTs) > 0 {
		for _, e := range m.NFTs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SubTokens) > 0 {
		for k, v := range m.SubTokens {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovGenesis(uint64(len(k))) + 1 + l + sovGenesis(uint64(l))
			n += mapEntrySize + 1 + sovGenesis(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *SubTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SubTokens) > 0 {
		for _, e := range m.SubTokens {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collections = append(m.Collections, Collection{})
			if err := m.Collections[len(m.Collections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NFTs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NFTs = append(m.NFTs, BaseNFT{})
			if err := m.NFTs[len(m.NFTs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubTokens == nil {
				m.SubTokens = make(map[string]SubTokens)
			}
			var mapkey string
			mapvalue := &SubTokens{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGenesis
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGenesis
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthGenesis
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &SubTokens{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGenesis(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthGenesis
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.SubTokens[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *SubTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: SubTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubTokens", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubTokens = append(m.SubTokens, SubToken{})
			if err := m.SubTokens[len(m.SubTokens)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
