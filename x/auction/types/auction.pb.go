// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bitsong/auction/v1beta1/auction.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Auction defines auction properties
type Auction struct {
	Id          uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuctionType uint32     `protobuf:"varint,2,opt,name=auction_type,json=auctionType,proto3" json:"auction_type,omitempty"`
	NftId       string     `protobuf:"bytes,3,opt,name=nft_id,json=nftId,proto3" json:"nft_id,omitempty"`
	StartTime   uint64     `protobuf:"varint,4,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty" yaml:"start_time"`
	Duration    uint64     `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty" yaml:"duration"`
	MinAmount   types.Coin `protobuf:"bytes,6,opt,name=min_amount,json=minAmount,proto3" json:"min_amount" yaml:"min_amount"`
	Owner       string     `protobuf:"bytes,7,opt,name=owner,proto3" json:"owner,omitempty"`
	Limit       uint32     `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *Auction) Reset()         { *m = Auction{} }
func (m *Auction) String() string { return proto.CompactTextString(m) }
func (*Auction) ProtoMessage()    {}
func (*Auction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7fb35b2a64f45fb, []int{0}
}
func (m *Auction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Auction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Auction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Auction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Auction.Merge(m, src)
}
func (m *Auction) XXX_Size() int {
	return m.Size()
}
func (m *Auction) XXX_DiscardUnknown() {
	xxx_messageInfo_Auction.DiscardUnknown(m)
}

var xxx_messageInfo_Auction proto.InternalMessageInfo

// Bid defines auction bidder and its amount
type Bid struct {
	AuctionId uint64     `protobuf:"varint,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty" yaml:"auction_id"`
	Bidder    string     `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	BidAmount types.Coin `protobuf:"bytes,3,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount" yaml:"bid_amount"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7fb35b2a64f45fb, []int{1}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Auction)(nil), "bitsong.auction.v1beta1.Auction")
	proto.RegisterType((*Bid)(nil), "bitsong.auction.v1beta1.Bid")
}

func init() {
	proto.RegisterFile("bitsong/auction/v1beta1/auction.proto", fileDescriptor_d7fb35b2a64f45fb)
}

var fileDescriptor_d7fb35b2a64f45fb = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xbb, 0x8e, 0xd3, 0x40,
	0x14, 0xf5, 0xe4, 0xb5, 0xeb, 0x59, 0x1e, 0x62, 0xd8, 0x05, 0xef, 0x16, 0x4e, 0xb0, 0x84, 0x94,
	0xca, 0xd6, 0x02, 0x12, 0xd2, 0x76, 0x1b, 0xaa, 0x6d, 0x87, 0xa5, 0xa1, 0x89, 0xc6, 0x9e, 0x49,
	0xf6, 0x4a, 0x99, 0x99, 0xc8, 0x1e, 0x03, 0xf9, 0x0b, 0x3e, 0x81, 0x2f, 0xe0, 0x1f, 0xe8, 0x52,
	0x6e, 0x49, 0x15, 0x41, 0xd2, 0x50, 0xe7, 0x0b, 0x90, 0xc7, 0xe3, 0x98, 0x96, 0xee, 0x9e, 0x7b,
	0xef, 0xb9, 0x33, 0xe7, 0xe8, 0xe0, 0x97, 0x29, 0x98, 0x42, 0xab, 0x79, 0xc2, 0xca, 0xcc, 0x80,
	0x56, 0xc9, 0xa7, 0xcb, 0x54, 0x18, 0x76, 0xd9, 0xe0, 0x78, 0x99, 0x6b, 0xa3, 0xc9, 0x73, 0xb7,
	0x16, 0x37, 0x6d, 0xb7, 0x76, 0x11, 0x66, 0xba, 0x90, 0xba, 0x48, 0x52, 0x56, 0x88, 0x03, 0x37,
	0xd3, 0xe0, 0x88, 0x17, 0xa7, 0x73, 0x3d, 0xd7, 0xb6, 0x4c, 0xaa, 0xaa, 0xee, 0x46, 0x3f, 0x3a,
	0xf8, 0xe8, 0xba, 0xbe, 0x44, 0x1e, 0xe1, 0x0e, 0xf0, 0x00, 0x8d, 0xd0, 0xb8, 0x47, 0x3b, 0xc0,
	0xc9, 0x0b, 0xfc, 0xc0, 0x3d, 0x32, 0x35, 0xab, 0xa5, 0x08, 0x3a, 0x23, 0x34, 0x7e, 0x48, 0x4f,
	0x5c, 0xef, 0x76, 0xb5, 0x14, 0xe4, 0x0c, 0x0f, 0xd4, 0xcc, 0x4c, 0x81, 0x07, 0xdd, 0x11, 0x1a,
	0xfb, 0xb4, 0xaf, 0x66, 0xe6, 0x86, 0x93, 0x37, 0x18, 0x17, 0x86, 0xe5, 0x66, 0x6a, 0x40, 0x8a,
	0xa0, 0x57, 0x5d, 0x9c, 0x9c, 0xed, 0x37, 0xc3, 0x27, 0x2b, 0x26, 0x17, 0x57, 0x51, 0x3b, 0x8b,
	0xa8, 0x6f, 0xc1, 0x2d, 0x48, 0x41, 0x12, 0x7c, 0xcc, 0xcb, 0x9c, 0x55, 0xc7, 0x83, 0xbe, 0xe5,
	0x3c, 0xdd, 0x6f, 0x86, 0x8f, 0x6b, 0x4e, 0x33, 0x89, 0xe8, 0x61, 0x89, 0xbc, 0xc7, 0x58, 0x82,
	0x9a, 0x32, 0xa9, 0x4b, 0x65, 0x82, 0xc1, 0x08, 0x8d, 0x4f, 0x5e, 0x9d, 0xc7, 0xb5, 0x0f, 0x71,
	0xe5, 0x43, 0x63, 0x4e, 0xfc, 0x4e, 0x83, 0x9a, 0x9c, 0xaf, 0x37, 0x43, 0xaf, 0xfd, 0x45, 0x4b,
	0x8d, 0xa8, 0x2f, 0x41, 0x5d, 0xdb, 0x9a, 0x9c, 0xe2, 0xbe, 0xfe, 0xac, 0x44, 0x1e, 0x1c, 0xd5,
	0x8a, 0x2c, 0xa8, 0xba, 0x0b, 0x90, 0x60, 0x82, 0x63, 0x6b, 0x42, 0x0d, 0xae, 0x7a, 0x7f, 0xbe,
	0x0d, 0x51, 0xf4, 0x1d, 0xe1, 0xee, 0x04, 0xac, 0xea, 0xc6, 0xaf, 0xc6, 0xc7, 0x7f, 0x55, 0xb7,
	0xb3, 0x88, 0xfa, 0x0e, 0xdc, 0x70, 0xf2, 0x0c, 0x0f, 0x52, 0xe0, 0x5c, 0xe4, 0xd6, 0x5f, 0x9f,
	0x3a, 0x54, 0x89, 0x4b, 0x81, 0x37, 0xe2, 0xba, 0xff, 0x29, 0xae, 0xa5, 0x46, 0xd4, 0x4f, 0x81,
	0xd7, 0xe2, 0xea, 0x0f, 0x4f, 0x3e, 0xac, 0x7f, 0x87, 0xde, 0x7a, 0x1b, 0xa2, 0xfb, 0x6d, 0x88,
	0x7e, 0x6d, 0x43, 0xf4, 0x75, 0x17, 0x7a, 0xf7, 0xbb, 0xd0, 0xfb, 0xb9, 0x0b, 0xbd, 0x8f, 0x6f,
	0xe7, 0x60, 0xee, 0xca, 0x34, 0xce, 0xb4, 0x4c, 0x5c, 0xd8, 0xf4, 0x6c, 0x06, 0x19, 0xb0, 0x45,
	0x92, 0xdd, 0x31, 0x50, 0x52, 0xf3, 0x72, 0x21, 0x8a, 0xe4, 0xcb, 0x21, 0xaa, 0x55, 0x3c, 0x8a,
	0x74, 0x60, 0x23, 0xf5, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x30, 0x9e, 0x26, 0xca,
	0x02, 0x00, 0x00,
}

func (this *Auction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Auction)
	if !ok {
		that2, ok := that.(Auction)
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
	if this.AuctionType != that1.AuctionType {
		return false
	}
	if this.NftId != that1.NftId {
		return false
	}
	if this.StartTime != that1.StartTime {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if !this.MinAmount.Equal(&that1.MinAmount) {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Limit != that1.Limit {
		return false
	}
	return true
}
func (this *Bid) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bid)
	if !ok {
		that2, ok := that.(Bid)
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
	if this.AuctionId != that1.AuctionId {
		return false
	}
	if this.Bidder != that1.Bidder {
		return false
	}
	if !this.BidAmount.Equal(&that1.BidAmount) {
		return false
	}
	return true
}
func (m *Auction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Auction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Auction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Limit != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Limit))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x3a
	}
	{
		size, err := m.MinAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Duration != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x28
	}
	if m.StartTime != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.StartTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NftId) > 0 {
		i -= len(m.NftId)
		copy(dAtA[i:], m.NftId)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.NftId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AuctionType != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionType))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.BidAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAuction(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintAuction(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.AuctionId != 0 {
		i = encodeVarintAuction(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuction(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuction(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Auction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovAuction(uint64(m.Id))
	}
	if m.AuctionType != 0 {
		n += 1 + sovAuction(uint64(m.AuctionType))
	}
	l = len(m.NftId)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.StartTime != 0 {
		n += 1 + sovAuction(uint64(m.StartTime))
	}
	if m.Duration != 0 {
		n += 1 + sovAuction(uint64(m.Duration))
	}
	l = m.MinAmount.Size()
	n += 1 + l + sovAuction(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	if m.Limit != 0 {
		n += 1 + sovAuction(uint64(m.Limit))
	}
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AuctionId != 0 {
		n += 1 + sovAuction(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovAuction(uint64(l))
	}
	l = m.BidAmount.Size()
	n += 1 + l + sovAuction(uint64(l))
	return n
}

func sovAuction(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuction(x uint64) (n int) {
	return sovAuction(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Auction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Auction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Auction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionType", wireType)
			}
			m.AuctionType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionType |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			m.StartTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuction
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
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuction
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuction
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
				return ErrInvalidLengthAuction
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuction
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuction(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthAuction
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthAuction
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
func skipAuction(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
					return 0, ErrIntOverflowAuction
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
				return 0, ErrInvalidLengthAuction
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuction
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuction
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuction        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuction          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuction = fmt.Errorf("proto: unexpected end of group")
)
