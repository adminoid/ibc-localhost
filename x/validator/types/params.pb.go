// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: decimal/validator/v1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google/protobuf"
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

// Params defines the parameters for the module.
type Params struct {
	// unbonding_time is the time duration of unbonding.
	UnbondingTime time.Duration `protobuf:"bytes,1,opt,name=unbonding_time,json=unbondingTime,proto3,stdduration" json:"unbonding_time"`
	// max_validators is the maximum number of validators.
	MaxValidators uint32 `protobuf:"varint,2,opt,name=max_validators,json=maxValidators,proto3" json:"max_validators,omitempty"`
	// max_entries is the max entries for either unbonding delegation or redelegation (per pair/trio).
	MaxEntries uint32 `protobuf:"varint,3,opt,name=max_entries,json=maxEntries,proto3" json:"max_entries,omitempty"`
	// historical_entries is the number of historical entries to persist.
	HistoricalEntries uint32 `protobuf:"varint,4,opt,name=historical_entries,json=historicalEntries,proto3" json:"historical_entries,omitempty"`
	// max_delegations is the maximum number of delegations per validator.
	MaxDelegations uint32 `protobuf:"varint,5,opt,name=max_delegations,json=maxDelegations,proto3" json:"max_delegations,omitempty"`
	// min_commission_rate is the chain-wide minimum commission rate that a validator can charge their delegators.
	MinCommissionRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=min_commission_rate,json=minCommissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"min_commission_rate"`
	// bond_denom defines the bondable coin denomination.
	BondDenom string `protobuf:"bytes,7,opt,name=bond_denom,json=bondDenom,proto3" json:"bond_denom,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_3180e3eded1f75a1, []int{0}
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

func (m *Params) GetUnbondingTime() time.Duration {
	if m != nil {
		return m.UnbondingTime
	}
	return 0
}

func (m *Params) GetMaxValidators() uint32 {
	if m != nil {
		return m.MaxValidators
	}
	return 0
}

func (m *Params) GetMaxEntries() uint32 {
	if m != nil {
		return m.MaxEntries
	}
	return 0
}

func (m *Params) GetHistoricalEntries() uint32 {
	if m != nil {
		return m.HistoricalEntries
	}
	return 0
}

func (m *Params) GetMaxDelegations() uint32 {
	if m != nil {
		return m.MaxDelegations
	}
	return 0
}

func (m *Params) GetBondDenom() string {
	if m != nil {
		return m.BondDenom
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "decimal.validator.v1.Params")
}

func init() { proto.RegisterFile("decimal/validator/v1/params.proto", fileDescriptor_3180e3eded1f75a1) }

var fileDescriptor_3180e3eded1f75a1 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x92, 0x31, 0x6f, 0xd4, 0x30,
	0x14, 0xc7, 0xcf, 0x14, 0x8e, 0xd6, 0xd5, 0x15, 0x35, 0x74, 0x48, 0x2b, 0x91, 0x1c, 0x48, 0xc0,
	0x2d, 0x89, 0x55, 0x98, 0x40, 0x4c, 0x47, 0x58, 0x98, 0x50, 0x40, 0x0c, 0x2c, 0x91, 0x93, 0x18,
	0xd7, 0x6a, 0xec, 0x77, 0xb2, 0x9d, 0xd3, 0xf1, 0x19, 0x58, 0x18, 0x3b, 0xf6, 0x43, 0xf0, 0x21,
	0x3a, 0x56, 0x4c, 0x88, 0xa1, 0xa0, 0xbb, 0x85, 0x8f, 0x81, 0xec, 0xa4, 0x29, 0x53, 0xe2, 0xff,
	0xfb, 0xf9, 0xff, 0xfe, 0x79, 0x79, 0xf8, 0x61, 0xcd, 0x2a, 0x21, 0x69, 0x43, 0x96, 0xb4, 0x11,
	0x35, 0xb5, 0xa0, 0xc9, 0xf2, 0x98, 0x2c, 0xa8, 0xa6, 0xd2, 0xa4, 0x0b, 0x0d, 0x16, 0x82, 0x83,
	0x1e, 0x49, 0x07, 0x24, 0x5d, 0x1e, 0x1f, 0x1d, 0x70, 0xe0, 0xe0, 0x01, 0xe2, 0xde, 0x3a, 0xf6,
	0x28, 0xe2, 0x00, 0xbc, 0x61, 0xc4, 0x9f, 0xca, 0xf6, 0x33, 0xa9, 0x5b, 0x4d, 0xad, 0x00, 0xd5,
	0xd7, 0x0f, 0x2b, 0x30, 0x12, 0x4c, 0xd1, 0x5d, 0xec, 0x0e, 0x5d, 0xe9, 0xd1, 0xd7, 0x2d, 0x3c,
	0x7e, 0xe7, 0xfb, 0x06, 0x6f, 0xf1, 0x5e, 0xab, 0x4a, 0x50, 0xb5, 0x50, 0xbc, 0xb0, 0x42, 0xb2,
	0x10, 0x4d, 0xd1, 0x6c, 0xf7, 0xd9, 0x61, 0xda, 0xd9, 0xa7, 0xd7, 0xf6, 0x69, 0xd6, 0xdb, 0xcf,
	0xb7, 0x2f, 0xae, 0xe2, 0xd1, 0xd9, 0xef, 0x18, 0xe5, 0x93, 0xe1, 0xea, 0x07, 0x21, 0x59, 0xf0,
	0x18, 0xef, 0x49, 0xba, 0x2a, 0x86, 0xec, 0x26, 0xbc, 0x35, 0x45, 0xb3, 0x49, 0x3e, 0x91, 0x74,
	0xf5, 0x71, 0x10, 0x83, 0x18, 0xef, 0x3a, 0x8c, 0x29, 0xab, 0x05, 0x33, 0xe1, 0x96, 0x67, 0xb0,
	0xa4, 0xab, 0x37, 0x9d, 0x12, 0x24, 0x38, 0x38, 0x11, 0xc6, 0x82, 0x16, 0x15, 0x6d, 0x06, 0xee,
	0xb6, 0xe7, 0xf6, 0x6f, 0x2a, 0xd7, 0xf8, 0x53, 0x7c, 0xcf, 0xf9, 0xd5, 0xac, 0x61, 0xdc, 0x27,
	0x34, 0xe1, 0x1d, 0xcf, 0xba, 0x34, 0xd9, 0x8d, 0x1a, 0x34, 0xf8, 0xbe, 0x14, 0xaa, 0xa8, 0x40,
	0x4a, 0x61, 0x8c, 0x00, 0x55, 0x68, 0x6a, 0x59, 0x38, 0x9e, 0xa2, 0xd9, 0xce, 0xfc, 0x95, 0xfb,
	0xaa, 0x5f, 0x57, 0xf1, 0x13, 0x2e, 0xec, 0x49, 0x5b, 0xa6, 0x15, 0xc8, 0x7e, 0x68, 0xfd, 0x23,
	0x31, 0xf5, 0x29, 0xb1, 0x5f, 0x16, 0xcc, 0xa4, 0x19, 0xab, 0x7e, 0x7c, 0x4f, 0x70, 0x3f, 0xd3,
	0x8c, 0x55, 0xf9, 0xbe, 0x14, 0xea, 0xf5, 0xe0, 0x9b, 0x53, 0xcb, 0x82, 0x07, 0x18, 0xbb, 0xe1,
	0x14, 0x35, 0x53, 0x20, 0xc3, 0xbb, 0xae, 0x49, 0xbe, 0xe3, 0x94, 0xcc, 0x09, 0x2f, 0xb7, 0xcf,
	0xce, 0x63, 0xf4, 0xf7, 0x3c, 0x46, 0xf3, 0xf7, 0x17, 0xeb, 0x08, 0x5d, 0xae, 0x23, 0xf4, 0x67,
	0x1d, 0xa1, 0x6f, 0x9b, 0x68, 0x74, 0xb9, 0x89, 0x46, 0x3f, 0x37, 0xd1, 0xe8, 0xd3, 0x8b, 0x52,
	0xd8, 0xb2, 0xad, 0x4e, 0x99, 0x4d, 0x41, 0x73, 0xd2, 0x2f, 0x87, 0x65, 0x54, 0x12, 0x0e, 0x89,
	0x91, 0x54, 0xdb, 0x44, 0x41, 0xcd, 0xc8, 0xea, 0xbf, 0x9d, 0xf2, 0x11, 0xcb, 0xb1, 0xff, 0x6f,
	0xcf, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x73, 0x89, 0x66, 0x75, 0x02, 0x00, 0x00,
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
	if this.UnbondingTime != that1.UnbondingTime {
		return false
	}
	if this.MaxValidators != that1.MaxValidators {
		return false
	}
	if this.MaxEntries != that1.MaxEntries {
		return false
	}
	if this.HistoricalEntries != that1.HistoricalEntries {
		return false
	}
	if this.MaxDelegations != that1.MaxDelegations {
		return false
	}
	if !this.MinCommissionRate.Equal(that1.MinCommissionRate) {
		return false
	}
	if this.BondDenom != that1.BondDenom {
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
	if len(m.BondDenom) > 0 {
		i -= len(m.BondDenom)
		copy(dAtA[i:], m.BondDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.BondDenom)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.MinCommissionRate.Size()
		i -= size
		if _, err := m.MinCommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.MaxDelegations != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxDelegations))
		i--
		dAtA[i] = 0x28
	}
	if m.HistoricalEntries != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.HistoricalEntries))
		i--
		dAtA[i] = 0x20
	}
	if m.MaxEntries != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxEntries))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxValidators != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxValidators))
		i--
		dAtA[i] = 0x10
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.UnbondingTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnbondingTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
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
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.UnbondingTime)
	n += 1 + l + sovParams(uint64(l))
	if m.MaxValidators != 0 {
		n += 1 + sovParams(uint64(m.MaxValidators))
	}
	if m.MaxEntries != 0 {
		n += 1 + sovParams(uint64(m.MaxEntries))
	}
	if m.HistoricalEntries != 0 {
		n += 1 + sovParams(uint64(m.HistoricalEntries))
	}
	if m.MaxDelegations != 0 {
		n += 1 + sovParams(uint64(m.MaxDelegations))
	}
	l = m.MinCommissionRate.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.BondDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.UnbondingTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxValidators", wireType)
			}
			m.MaxValidators = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxValidators |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxEntries", wireType)
			}
			m.MaxEntries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxEntries |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalEntries", wireType)
			}
			m.HistoricalEntries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HistoricalEntries |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDelegations", wireType)
			}
			m.MaxDelegations = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDelegations |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCommissionRate", wireType)
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
			if err := m.MinCommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BondDenom", wireType)
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
			m.BondDenom = string(dAtA[iNdEx:postIndex])
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
