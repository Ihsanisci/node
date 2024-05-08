// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zetachain/zetacore/lightclient/block_header_verification.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// HeaderSupportedChain is a structure containing information of weather a chain
// is enabled or not for block header verification
type HeaderSupportedChain struct {
	ChainId int64 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Enabled bool  `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (m *HeaderSupportedChain) Reset()         { *m = HeaderSupportedChain{} }
func (m *HeaderSupportedChain) String() string { return proto.CompactTextString(m) }
func (*HeaderSupportedChain) ProtoMessage()    {}
func (*HeaderSupportedChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_deea61e47e024601, []int{0}
}
func (m *HeaderSupportedChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HeaderSupportedChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HeaderSupportedChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HeaderSupportedChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderSupportedChain.Merge(m, src)
}
func (m *HeaderSupportedChain) XXX_Size() int {
	return m.Size()
}
func (m *HeaderSupportedChain) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderSupportedChain.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderSupportedChain proto.InternalMessageInfo

func (m *HeaderSupportedChain) GetChainId() int64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *HeaderSupportedChain) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type BlockHeaderVerification struct {
	HeaderSupportedChains []HeaderSupportedChain `protobuf:"bytes,1,rep,name=header_supported_chains,json=headerSupportedChains,proto3" json:"header_supported_chains"`
}

func (m *BlockHeaderVerification) Reset()         { *m = BlockHeaderVerification{} }
func (m *BlockHeaderVerification) String() string { return proto.CompactTextString(m) }
func (*BlockHeaderVerification) ProtoMessage()    {}
func (*BlockHeaderVerification) Descriptor() ([]byte, []int) {
	return fileDescriptor_deea61e47e024601, []int{1}
}
func (m *BlockHeaderVerification) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeaderVerification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeaderVerification.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeaderVerification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeaderVerification.Merge(m, src)
}
func (m *BlockHeaderVerification) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeaderVerification) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeaderVerification.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeaderVerification proto.InternalMessageInfo

func (m *BlockHeaderVerification) GetHeaderSupportedChains() []HeaderSupportedChain {
	if m != nil {
		return m.HeaderSupportedChains
	}
	return nil
}

func init() {
	proto.RegisterType((*HeaderSupportedChain)(nil), "zetachain.zetacore.lightclient.HeaderSupportedChain")
	proto.RegisterType((*BlockHeaderVerification)(nil), "zetachain.zetacore.lightclient.BlockHeaderVerification")
}

func init() {
	proto.RegisterFile("zetachain/zetacore/lightclient/block_header_verification.proto", fileDescriptor_deea61e47e024601)
}

var fileDescriptor_deea61e47e024601 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xab, 0x4a, 0x2d, 0x49,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0xb3, 0xf2, 0x8b, 0x52, 0xf5, 0x73, 0x32, 0xd3, 0x33,
	0x4a, 0x92, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xf4, 0x93, 0x72, 0xf2, 0x93, 0xb3, 0xe3, 0x33, 0x52,
	0x13, 0x53, 0x52, 0x8b, 0xe2, 0xcb, 0x52, 0x8b, 0x32, 0xd3, 0x32, 0x93, 0x13, 0x4b, 0x32, 0xf3,
	0xf3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xe4, 0xe0, 0xfa, 0xf5, 0x60, 0xfa, 0xf5, 0x90,
	0xf4, 0x4b, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x95, 0xea, 0x83, 0x58, 0x10, 0x5d, 0x4a, 0xde,
	0x5c, 0x22, 0x1e, 0x60, 0x23, 0x83, 0x4b, 0x0b, 0x0a, 0xf2, 0x8b, 0x4a, 0x52, 0x53, 0x9c, 0x41,
	0x46, 0x08, 0x49, 0x72, 0x71, 0x80, 0xcd, 0x8a, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60,
	0x0e, 0x62, 0x07, 0xf3, 0x3d, 0x53, 0x84, 0x24, 0xb8, 0xd8, 0x53, 0xf3, 0x12, 0x93, 0x72, 0x52,
	0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x82, 0x60, 0x5c, 0xa5, 0x5e, 0x46, 0x2e, 0x71, 0x27,
	0x90, 0x33, 0x21, 0x46, 0x86, 0x21, 0x39, 0x52, 0xa8, 0x88, 0x4b, 0x1c, 0xea, 0xf6, 0x62, 0x98,
	0x4d, 0xf1, 0x60, 0x13, 0x8b, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x4c, 0xf4, 0xf0, 0x7b,
	0x40, 0x0f, 0x9b, 0x3b, 0x9d, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x12, 0xcd, 0xc0, 0x22, 0x57,
	0xec, 0xe4, 0x73, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x46, 0xe9, 0x99,
	0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xe0, 0xd0, 0xd6, 0x45, 0x0b, 0xf8, 0x0a, 0x94,
	0xa0, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x87, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xf3, 0x4c, 0xcf, 0xdf, 0xa9, 0x01, 0x00, 0x00,
}

func (m *HeaderSupportedChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HeaderSupportedChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HeaderSupportedChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Enabled {
		i--
		if m.Enabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.ChainId != 0 {
		i = encodeVarintBlockHeaderVerification(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockHeaderVerification) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeaderVerification) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeaderVerification) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HeaderSupportedChains) > 0 {
		for iNdEx := len(m.HeaderSupportedChains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HeaderSupportedChains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockHeaderVerification(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockHeaderVerification(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockHeaderVerification(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HeaderSupportedChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovBlockHeaderVerification(uint64(m.ChainId))
	}
	if m.Enabled {
		n += 2
	}
	return n
}

func (m *BlockHeaderVerification) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.HeaderSupportedChains) > 0 {
		for _, e := range m.HeaderSupportedChains {
			l = e.Size()
			n += 1 + l + sovBlockHeaderVerification(uint64(l))
		}
	}
	return n
}

func sovBlockHeaderVerification(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockHeaderVerification(x uint64) (n int) {
	return sovBlockHeaderVerification(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HeaderSupportedChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockHeaderVerification
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
			return fmt.Errorf("proto: HeaderSupportedChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HeaderSupportedChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeaderVerification
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Enabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeaderVerification
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
			m.Enabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBlockHeaderVerification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockHeaderVerification
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
func (m *BlockHeaderVerification) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockHeaderVerification
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
			return fmt.Errorf("proto: BlockHeaderVerification: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeaderVerification: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderSupportedChains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockHeaderVerification
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
				return ErrInvalidLengthBlockHeaderVerification
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockHeaderVerification
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderSupportedChains = append(m.HeaderSupportedChains, HeaderSupportedChain{})
			if err := m.HeaderSupportedChains[len(m.HeaderSupportedChains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockHeaderVerification(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockHeaderVerification
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
func skipBlockHeaderVerification(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockHeaderVerification
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
					return 0, ErrIntOverflowBlockHeaderVerification
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
					return 0, ErrIntOverflowBlockHeaderVerification
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
				return 0, ErrInvalidLengthBlockHeaderVerification
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockHeaderVerification
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockHeaderVerification
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockHeaderVerification        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockHeaderVerification          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockHeaderVerification = fmt.Errorf("proto: unexpected end of group")
)
