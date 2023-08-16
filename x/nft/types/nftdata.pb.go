// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gchain/nft/nftdata.proto

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

type Nftdata struct {
	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Meta        string `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	Uri         string `protobuf:"bytes,5,opt,name=uri,proto3" json:"uri,omitempty"`
	Urihash     string `protobuf:"bytes,6,opt,name=urihash,proto3" json:"urihash,omitempty"`
	Owner       string `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *Nftdata) Reset()         { *m = Nftdata{} }
func (m *Nftdata) String() string { return proto.CompactTextString(m) }
func (*Nftdata) ProtoMessage()    {}
func (*Nftdata) Descriptor() ([]byte, []int) {
	return fileDescriptor_20414795974c5eca, []int{0}
}
func (m *Nftdata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Nftdata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Nftdata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Nftdata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nftdata.Merge(m, src)
}
func (m *Nftdata) XXX_Size() int {
	return m.Size()
}
func (m *Nftdata) XXX_DiscardUnknown() {
	xxx_messageInfo_Nftdata.DiscardUnknown(m)
}

var xxx_messageInfo_Nftdata proto.InternalMessageInfo

func (m *Nftdata) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Nftdata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Nftdata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Nftdata) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Nftdata) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *Nftdata) GetUrihash() string {
	if m != nil {
		return m.Urihash
	}
	return ""
}

func (m *Nftdata) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*Nftdata)(nil), "gchain.nft.Nftdata")
}

func init() { proto.RegisterFile("gchain/nft/nftdata.proto", fileDescriptor_20414795974c5eca) }

var fileDescriptor_20414795974c5eca = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x4f, 0xce, 0x48,
	0xcc, 0xcc, 0xd3, 0xcf, 0x4b, 0x2b, 0x01, 0xe1, 0x94, 0xc4, 0x92, 0x44, 0xbd, 0x82, 0xa2, 0xfc,
	0x92, 0x7c, 0x21, 0x2e, 0x88, 0x8c, 0x5e, 0x5e, 0x5a, 0x89, 0xd2, 0x52, 0x46, 0x2e, 0x76, 0x3f,
	0x88, 0xac, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x53,
	0x66, 0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0x2d, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99,
	0x9f, 0x27, 0xc1, 0x0c, 0x96, 0x42, 0x16, 0x02, 0xe9, 0xca, 0x4d, 0x2d, 0x49, 0x94, 0x60, 0x81,
	0xe8, 0x02, 0xb1, 0x85, 0x04, 0xb8, 0x98, 0x4b, 0x8b, 0x32, 0x25, 0x58, 0xc1, 0x42, 0x20, 0xa6,
	0x90, 0x04, 0x17, 0x7b, 0x69, 0x51, 0x66, 0x46, 0x62, 0x71, 0x86, 0x04, 0x1b, 0x58, 0x14, 0xc6,
	0x15, 0x12, 0xe1, 0x62, 0xcd, 0x2f, 0xcf, 0x4b, 0x2d, 0x92, 0x60, 0x07, 0x8b, 0x43, 0x38, 0x4e,
	0x3a, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x04, 0xf5, 0x67, 0x05,
	0xd8, 0xa7, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x8f, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x15, 0x11, 0x95, 0x79, 0x04, 0x01, 0x00, 0x00,
}

func (m *Nftdata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Nftdata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Nftdata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNftdata(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Urihash) > 0 {
		i -= len(m.Urihash)
		copy(dAtA[i:], m.Urihash)
		i = encodeVarintNftdata(dAtA, i, uint64(len(m.Urihash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintNftdata(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Meta) > 0 {
		i -= len(m.Meta)
		copy(dAtA[i:], m.Meta)
		i = encodeVarintNftdata(dAtA, i, uint64(len(m.Meta)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintNftdata(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNftdata(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintNftdata(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNftdata(dAtA []byte, offset int, v uint64) int {
	offset -= sovNftdata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Nftdata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovNftdata(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNftdata(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovNftdata(uint64(l))
	}
	l = len(m.Meta)
	if l > 0 {
		n += 1 + l + sovNftdata(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovNftdata(uint64(l))
	}
	l = len(m.Urihash)
	if l > 0 {
		n += 1 + l + sovNftdata(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNftdata(uint64(l))
	}
	return n
}

func sovNftdata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNftdata(x uint64) (n int) {
	return sovNftdata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Nftdata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNftdata
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
			return fmt.Errorf("proto: Nftdata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Nftdata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
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
				return ErrInvalidLengthNftdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
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
				return ErrInvalidLengthNftdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
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
				return ErrInvalidLengthNftdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Meta = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
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
				return ErrInvalidLengthNftdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Urihash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
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
				return ErrInvalidLengthNftdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Urihash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNftdata
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
				return ErrInvalidLengthNftdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNftdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNftdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNftdata
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
func skipNftdata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNftdata
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
					return 0, ErrIntOverflowNftdata
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
					return 0, ErrIntOverflowNftdata
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
				return 0, ErrInvalidLengthNftdata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNftdata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNftdata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNftdata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNftdata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNftdata = fmt.Errorf("proto: unexpected end of group")
)
