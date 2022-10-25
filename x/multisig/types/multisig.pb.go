// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/multisig/v1/multisig.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// Wallet defines multisig wallet.
type Wallet struct {
	Address   string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Owners    []string `protobuf:"bytes,2,rep,name=owners,proto3" json:"owners,omitempty"`
	Weights   []uint32 `protobuf:"varint,3,rep,packed,name=weights,proto3" json:"weights,omitempty"`
	Threshold uint32   `protobuf:"varint,4,opt,name=threshold,proto3" json:"threshold,omitempty"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfab0ba1f9829c94, []int{0}
}
func (m *Wallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Wallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Wallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Wallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wallet.Merge(m, src)
}
func (m *Wallet) XXX_Size() int {
	return m.Size()
}
func (m *Wallet) XXX_DiscardUnknown() {
	xxx_messageInfo_Wallet.DiscardUnknown(m)
}

var xxx_messageInfo_Wallet proto.InternalMessageInfo

// Transaction defines multisig transaction.
type Transaction struct {
	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Wallet    string    `protobuf:"bytes,2,opt,name=wallet,proto3" json:"wallet,omitempty"`
	Message   types.Any `protobuf:"bytes,3,opt,name=message,proto3" json:"message"`
	CreatedAt int64     `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_dfab0ba1f9829c94, []int{1}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return m.Size()
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Wallet)(nil), "decimal.multisig.v1.Wallet")
	proto.RegisterType((*Transaction)(nil), "decimal.multisig.v1.Transaction")
}

func init() {
	proto.RegisterFile("decimal/multisig/v1/multisig.proto", fileDescriptor_dfab0ba1f9829c94)
}

var fileDescriptor_dfab0ba1f9829c94 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x8b, 0x13, 0x41,
	0x14, 0xc7, 0x77, 0xb2, 0x92, 0x90, 0x39, 0x4f, 0x61, 0x0c, 0xb2, 0x1e, 0xba, 0x59, 0x52, 0x6d,
	0x93, 0x1d, 0x73, 0x36, 0x62, 0x23, 0x49, 0x6f, 0xb3, 0x0a, 0x82, 0xcd, 0x31, 0xbb, 0xf3, 0x9c,
	0x0c, 0xb7, 0xbb, 0x73, 0xcc, 0xbc, 0x24, 0xe6, 0x3f, 0xb0, 0xb4, 0x14, 0xab, 0xfc, 0x11, 0xd7,
	0xda, 0x1f, 0x56, 0x87, 0x95, 0x95, 0x48, 0xd2, 0xf8, 0x67, 0x88, 0xfb, 0xe3, 0x0e, 0x2b, 0xbb,
	0xf9, 0xfe, 0x78, 0xc3, 0x87, 0xc7, 0xa3, 0x13, 0x09, 0xb9, 0x2e, 0x45, 0xc1, 0xcb, 0x55, 0x81,
	0xda, 0x69, 0xc5, 0xd7, 0xb3, 0x9b, 0x77, 0x72, 0x61, 0x0d, 0x1a, 0xf6, 0xa0, 0xed, 0x24, 0x37,
	0xfe, 0x7a, 0x76, 0x32, 0x52, 0x46, 0x99, 0x3a, 0xe7, 0x7f, 0x5f, 0x4d, 0xf5, 0xe4, 0x91, 0x32,
	0x46, 0x15, 0xc0, 0x6b, 0x95, 0xad, 0xde, 0x73, 0x51, 0x6d, 0xbb, 0x28, 0x37, 0xae, 0x34, 0xee,
	0xac, 0x99, 0x69, 0x44, 0x1b, 0x85, 0x8d, 0xe2, 0x99, 0x70, 0xc0, 0xd7, 0xb3, 0x0c, 0x50, 0xcc,
	0x78, 0x6e, 0x74, 0xd5, 0xe4, 0x93, 0x2f, 0x84, 0xf6, 0xdf, 0x8a, 0xa2, 0x00, 0x64, 0xa7, 0x74,
	0x20, 0xa4, 0xb4, 0xe0, 0x5c, 0x40, 0x22, 0x12, 0x0f, 0x17, 0xc1, 0xf7, 0xcb, 0xe9, 0xa8, 0xfd,
	0x6d, 0xde, 0x24, 0xaf, 0xd1, 0xea, 0x4a, 0xa5, 0x5d, 0x91, 0x3d, 0xa4, 0x7d, 0xb3, 0xa9, 0xc0,
	0xba, 0xa0, 0x17, 0xf9, 0xf1, 0x30, 0x6d, 0x15, 0x0b, 0xe8, 0x60, 0x03, 0x5a, 0x2d, 0xd1, 0x05,
	0x7e, 0xe4, 0xc7, 0xc7, 0x69, 0x27, 0xd9, 0x63, 0x3a, 0xc4, 0xa5, 0x05, 0xb7, 0x34, 0x85, 0x0c,
	0xee, 0x44, 0x24, 0x3e, 0x4e, 0x6f, 0x8d, 0x17, 0x77, 0x3f, 0xee, 0xc6, 0xde, 0xe7, 0xdd, 0x98,
	0xfc, 0xde, 0x8d, 0xc9, 0xe4, 0x2b, 0xa1, 0x47, 0x6f, 0xac, 0xa8, 0x9c, 0xc8, 0x51, 0x9b, 0x8a,
	0xdd, 0xa3, 0x3d, 0x2d, 0x1b, 0xb8, 0xb4, 0xa7, 0x25, 0x7b, 0x4a, 0xfb, 0x9b, 0x9a, 0x3d, 0xe8,
	0xfd, 0x07, 0xb8, 0xed, 0xb1, 0x97, 0x74, 0x50, 0x82, 0x73, 0x42, 0x41, 0xe0, 0x47, 0x24, 0x3e,
	0x3a, 0x1d, 0x25, 0xcd, 0x5a, 0x93, 0x6e, 0xad, 0xc9, 0xbc, 0xda, 0x2e, 0xee, 0x5f, 0xfd, 0x1c,
	0x7b, 0xdf, 0x2e, 0xa7, 0x03, 0x27, 0xcf, 0x93, 0x57, 0x4e, 0xa5, 0xdd, 0x14, 0x7b, 0x42, 0x69,
	0x6e, 0x41, 0x20, 0xc8, 0x33, 0x81, 0x35, 0xbf, 0x9f, 0x0e, 0x5b, 0x67, 0x8e, 0xff, 0xf2, 0x2f,
	0xd2, 0xab, 0x7d, 0x48, 0xae, 0xf7, 0x21, 0xf9, 0xb5, 0x0f, 0xc9, 0xa7, 0x43, 0xe8, 0x5d, 0x1f,
	0x42, 0xef, 0xc7, 0x21, 0xf4, 0xde, 0x3d, 0xcf, 0x34, 0x66, 0xab, 0xfc, 0x1c, 0x30, 0x31, 0x56,
	0xf1, 0xf6, 0x0a, 0x10, 0x44, 0xc9, 0x95, 0x99, 0xba, 0x52, 0x58, 0x9c, 0x56, 0x46, 0x02, 0xff,
	0x70, 0x7b, 0x3d, 0xb8, 0xbd, 0x00, 0x97, 0xf5, 0x6b, 0xd0, 0x67, 0x7f, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x9d, 0x2b, 0xb8, 0x1e, 0x5e, 0x02, 0x00, 0x00,
}

func (this *Wallet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Wallet)
	if !ok {
		that2, ok := that.(Wallet)
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
	if this.Address != that1.Address {
		return false
	}
	if len(this.Owners) != len(that1.Owners) {
		return false
	}
	for i := range this.Owners {
		if this.Owners[i] != that1.Owners[i] {
			return false
		}
	}
	if len(this.Weights) != len(that1.Weights) {
		return false
	}
	for i := range this.Weights {
		if this.Weights[i] != that1.Weights[i] {
			return false
		}
	}
	if this.Threshold != that1.Threshold {
		return false
	}
	return true
}
func (this *Transaction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transaction)
	if !ok {
		that2, ok := that.(Transaction)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Wallet != that1.Wallet {
		return false
	}
	if !this.Message.Equal(&that1.Message) {
		return false
	}
	if this.CreatedAt != that1.CreatedAt {
		return false
	}
	return true
}
func (m *Wallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Wallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Wallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Threshold != 0 {
		i = encodeVarintMultisig(dAtA, i, uint64(m.Threshold))
		i--
		dAtA[i] = 0x20
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
		i = encodeVarintMultisig(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owners) > 0 {
		for iNdEx := len(m.Owners) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Owners[iNdEx])
			copy(dAtA[i:], m.Owners[iNdEx])
			i = encodeVarintMultisig(dAtA, i, uint64(len(m.Owners[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Transaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Transaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Transaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreatedAt != 0 {
		i = encodeVarintMultisig(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Message.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMultisig(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Wallet) > 0 {
		i -= len(m.Wallet)
		copy(dAtA[i:], m.Wallet)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.Wallet)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintMultisig(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMultisig(dAtA []byte, offset int, v uint64) int {
	offset -= sovMultisig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Wallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	if len(m.Owners) > 0 {
		for _, s := range m.Owners {
			l = len(s)
			n += 1 + l + sovMultisig(uint64(l))
		}
	}
	if len(m.Weights) > 0 {
		l = 0
		for _, e := range m.Weights {
			l += sovMultisig(uint64(e))
		}
		n += 1 + sovMultisig(uint64(l)) + l
	}
	if m.Threshold != 0 {
		n += 1 + sovMultisig(uint64(m.Threshold))
	}
	return n
}

func (m *Transaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	l = len(m.Wallet)
	if l > 0 {
		n += 1 + l + sovMultisig(uint64(l))
	}
	l = m.Message.Size()
	n += 1 + l + sovMultisig(uint64(l))
	if m.CreatedAt != 0 {
		n += 1 + sovMultisig(uint64(m.CreatedAt))
	}
	return n
}

func sovMultisig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMultisig(x uint64) (n int) {
	return sovMultisig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Wallet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: Wallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wallet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owners", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owners = append(m.Owners, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowMultisig
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
						return ErrIntOverflowMultisig
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
					return ErrInvalidLengthMultisig
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthMultisig
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
							return ErrIntOverflowMultisig
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			m.Threshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func (m *Transaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMultisig
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
			return fmt.Errorf("proto: Transaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Transaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wallet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Wallet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
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
				return ErrInvalidLengthMultisig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMultisig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMultisig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMultisig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMultisig
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
func skipMultisig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMultisig
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
					return 0, ErrIntOverflowMultisig
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
					return 0, ErrIntOverflowMultisig
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
				return 0, ErrInvalidLengthMultisig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMultisig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMultisig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMultisig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMultisig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMultisig = fmt.Errorf("proto: unexpected end of group")
)
