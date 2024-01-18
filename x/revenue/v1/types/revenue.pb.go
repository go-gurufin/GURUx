// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gurux/revenue/v1/revenue.proto

package types

import (
	fmt "fmt"
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

// Revenue defines an instance that organizes fee distribution conditions for
// the owner of a given smart contract
type Revenue struct {
	// contract_address is the hex address of a registered contract
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// deployer_address is the bech32 address of message sender. It must be the same as the origin EOA
	// sending the transaction which deploys the contract
	DeployerAddress string `protobuf:"bytes,2,opt,name=deployer_address,json=deployerAddress,proto3" json:"deployer_address,omitempty"`
	// withdrawer_address is the bech32 address of account receiving the transaction fees it defaults to
	// deployer_address
	WithdrawerAddress string `protobuf:"bytes,3,opt,name=withdrawer_address,json=withdrawerAddress,proto3" json:"withdrawer_address,omitempty"`
}

func (m *Revenue) Reset()         { *m = Revenue{} }
func (m *Revenue) String() string { return proto.CompactTextString(m) }
func (*Revenue) ProtoMessage()    {}
func (*Revenue) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d087165192e447, []int{0}
}
func (m *Revenue) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Revenue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Revenue.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Revenue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Revenue.Merge(m, src)
}
func (m *Revenue) XXX_Size() int {
	return m.Size()
}
func (m *Revenue) XXX_DiscardUnknown() {
	xxx_messageInfo_Revenue.DiscardUnknown(m)
}

var xxx_messageInfo_Revenue proto.InternalMessageInfo

func (m *Revenue) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *Revenue) GetDeployerAddress() string {
	if m != nil {
		return m.DeployerAddress
	}
	return ""
}

func (m *Revenue) GetWithdrawerAddress() string {
	if m != nil {
		return m.WithdrawerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Revenue)(nil), "gurux.revenue.v1.Revenue")
}

func init() { proto.RegisterFile("gurux/revenue/v1/revenue.proto", fileDescriptor_33d087165192e447) }

var fileDescriptor_33d087165192e447 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2d, 0xcb, 0xcd,
	0x2f, 0xd6, 0x2f, 0x4a, 0x2d, 0x4b, 0xcd, 0x2b, 0x4d, 0xd5, 0x2f, 0x33, 0x84, 0x31, 0xf5, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0xc0, 0xf2, 0x7a, 0x30, 0xc1, 0x32, 0x43, 0xa5, 0x3e, 0x46,
	0x2e, 0xf6, 0x20, 0x08, 0x57, 0x48, 0x93, 0x4b, 0x20, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x31, 0xb9,
	0x24, 0x3e, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x88,
	0x1f, 0x26, 0xee, 0x08, 0x11, 0x06, 0x29, 0x4d, 0x49, 0x2d, 0xc8, 0xc9, 0xaf, 0x4c, 0x2d, 0x82,
	0x2b, 0x65, 0x82, 0x28, 0x85, 0x89, 0xc3, 0x94, 0xea, 0x72, 0x09, 0x95, 0x67, 0x96, 0x64, 0xa4,
	0x14, 0x25, 0x96, 0x23, 0x29, 0x66, 0x06, 0x2b, 0x16, 0x44, 0xc8, 0x40, 0x95, 0x3b, 0xb9, 0x9f,
	0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31,
	0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x6e, 0x7a, 0x66, 0x49, 0x46, 0x69,
	0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x9f, 0x10, 0xb2, 0xcc, 0xd0, 0x50, 0xbf, 0x02, 0xd9,
	0xcf, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x2f, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x2b, 0x34, 0xe5, 0x14, 0x01, 0x00, 0x00,
}

func (m *Revenue) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Revenue) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Revenue) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawerAddress) > 0 {
		i -= len(m.WithdrawerAddress)
		copy(dAtA[i:], m.WithdrawerAddress)
		i = encodeVarintRevenue(dAtA, i, uint64(len(m.WithdrawerAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DeployerAddress) > 0 {
		i -= len(m.DeployerAddress)
		copy(dAtA[i:], m.DeployerAddress)
		i = encodeVarintRevenue(dAtA, i, uint64(len(m.DeployerAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintRevenue(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRevenue(dAtA []byte, offset int, v uint64) int {
	offset -= sovRevenue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Revenue) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovRevenue(uint64(l))
	}
	l = len(m.DeployerAddress)
	if l > 0 {
		n += 1 + l + sovRevenue(uint64(l))
	}
	l = len(m.WithdrawerAddress)
	if l > 0 {
		n += 1 + l + sovRevenue(uint64(l))
	}
	return n
}

func sovRevenue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRevenue(x uint64) (n int) {
	return sovRevenue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Revenue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRevenue
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
			return fmt.Errorf("proto: Revenue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Revenue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevenue
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
				return ErrInvalidLengthRevenue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRevenue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeployerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevenue
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
				return ErrInvalidLengthRevenue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRevenue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeployerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRevenue
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
				return ErrInvalidLengthRevenue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRevenue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRevenue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRevenue
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
func skipRevenue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRevenue
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
					return 0, ErrIntOverflowRevenue
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
					return 0, ErrIntOverflowRevenue
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
				return 0, ErrInvalidLengthRevenue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRevenue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRevenue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRevenue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRevenue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRevenue = fmt.Errorf("proto: unexpected end of group")
)
