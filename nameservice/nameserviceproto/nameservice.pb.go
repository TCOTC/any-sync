// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nameservice/nameserviceproto/protos/nameservice.proto

package nameserviceproto

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

type NameAvailableRequest struct {
	// Name including .any suffix
	FullName string `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
}

func (m *NameAvailableRequest) Reset()         { *m = NameAvailableRequest{} }
func (m *NameAvailableRequest) String() string { return proto.CompactTextString(m) }
func (*NameAvailableRequest) ProtoMessage()    {}
func (*NameAvailableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06bca2ea4304f305, []int{0}
}
func (m *NameAvailableRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameAvailableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameAvailableRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameAvailableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameAvailableRequest.Merge(m, src)
}
func (m *NameAvailableRequest) XXX_Size() int {
	return m.Size()
}
func (m *NameAvailableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameAvailableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameAvailableRequest proto.InternalMessageInfo

func (m *NameAvailableRequest) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

type NameByAddressRequest struct {
	// EOA -> SCW -> name
	// A SCW Ethereum address that owns that name
	OwnerScwEthAddress string `protobuf:"bytes,1,opt,name=ownerScwEthAddress,proto3" json:"ownerScwEthAddress,omitempty"`
}

func (m *NameByAddressRequest) Reset()         { *m = NameByAddressRequest{} }
func (m *NameByAddressRequest) String() string { return proto.CompactTextString(m) }
func (*NameByAddressRequest) ProtoMessage()    {}
func (*NameByAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_06bca2ea4304f305, []int{1}
}
func (m *NameByAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameByAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameByAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameByAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameByAddressRequest.Merge(m, src)
}
func (m *NameByAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *NameByAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NameByAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NameByAddressRequest proto.InternalMessageInfo

func (m *NameByAddressRequest) GetOwnerScwEthAddress() string {
	if m != nil {
		return m.OwnerScwEthAddress
	}
	return ""
}

type NameAvailableResponse struct {
	Available bool `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	// EOA -> SCW -> name
	// This field is non-empty only if name is "already registered"
	OwnerScwEthAddress string `protobuf:"bytes,2,opt,name=ownerScwEthAddress,proto3" json:"ownerScwEthAddress,omitempty"`
	// This field is non-empty only if name is "already registered"
	OwnerEthAddress string `protobuf:"bytes,3,opt,name=ownerEthAddress,proto3" json:"ownerEthAddress,omitempty"`
	// A content hash attached to this name
	// This field is non-empty only if name is "already registered"
	OwnerAnyAddress string `protobuf:"bytes,4,opt,name=ownerAnyAddress,proto3" json:"ownerAnyAddress,omitempty"`
	// A SpaceID attached to this name
	// This field is non-empty only if name is "already registered"
	SpaceId string `protobuf:"bytes,5,opt,name=spaceId,proto3" json:"spaceId,omitempty"`
	// doesn't work with marashalling/unmarshalling
	// google.protobuf.Timestamp nameExpires = 5 [(gogoproto.stdtime) = true];
	NameExpires int64 `protobuf:"varint,6,opt,name=nameExpires,proto3" json:"nameExpires,omitempty"`
}

func (m *NameAvailableResponse) Reset()         { *m = NameAvailableResponse{} }
func (m *NameAvailableResponse) String() string { return proto.CompactTextString(m) }
func (*NameAvailableResponse) ProtoMessage()    {}
func (*NameAvailableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06bca2ea4304f305, []int{2}
}
func (m *NameAvailableResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameAvailableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameAvailableResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameAvailableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameAvailableResponse.Merge(m, src)
}
func (m *NameAvailableResponse) XXX_Size() int {
	return m.Size()
}
func (m *NameAvailableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NameAvailableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NameAvailableResponse proto.InternalMessageInfo

func (m *NameAvailableResponse) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *NameAvailableResponse) GetOwnerScwEthAddress() string {
	if m != nil {
		return m.OwnerScwEthAddress
	}
	return ""
}

func (m *NameAvailableResponse) GetOwnerEthAddress() string {
	if m != nil {
		return m.OwnerEthAddress
	}
	return ""
}

func (m *NameAvailableResponse) GetOwnerAnyAddress() string {
	if m != nil {
		return m.OwnerAnyAddress
	}
	return ""
}

func (m *NameAvailableResponse) GetSpaceId() string {
	if m != nil {
		return m.SpaceId
	}
	return ""
}

func (m *NameAvailableResponse) GetNameExpires() int64 {
	if m != nil {
		return m.NameExpires
	}
	return 0
}

type NameByAddressResponse struct {
	Found bool   `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *NameByAddressResponse) Reset()         { *m = NameByAddressResponse{} }
func (m *NameByAddressResponse) String() string { return proto.CompactTextString(m) }
func (*NameByAddressResponse) ProtoMessage()    {}
func (*NameByAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_06bca2ea4304f305, []int{3}
}
func (m *NameByAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NameByAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NameByAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NameByAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameByAddressResponse.Merge(m, src)
}
func (m *NameByAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *NameByAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NameByAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NameByAddressResponse proto.InternalMessageInfo

func (m *NameByAddressResponse) GetFound() bool {
	if m != nil {
		return m.Found
	}
	return false
}

func (m *NameByAddressResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*NameAvailableRequest)(nil), "NameAvailableRequest")
	proto.RegisterType((*NameByAddressRequest)(nil), "NameByAddressRequest")
	proto.RegisterType((*NameAvailableResponse)(nil), "NameAvailableResponse")
	proto.RegisterType((*NameByAddressResponse)(nil), "NameByAddressResponse")
}

func init() {
	proto.RegisterFile("nameservice/nameserviceproto/protos/nameservice.proto", fileDescriptor_06bca2ea4304f305)
}

var fileDescriptor_06bca2ea4304f305 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xed, 0xca, 0x87, 0x30, 0x1e, 0x30, 0x1b, 0x30, 0x0d, 0x21, 0x0d, 0xe9, 0x89, 0x53, 0x49,
	0x30, 0x7a, 0x2f, 0x06, 0x0d, 0x17, 0x0f, 0xf5, 0xe6, 0x6d, 0xa1, 0x43, 0x24, 0x29, 0xdb, 0xda,
	0x2d, 0x20, 0xff, 0x82, 0x9f, 0xe5, 0x91, 0xa3, 0x47, 0x03, 0xbf, 0xc1, 0xbb, 0x61, 0xd9, 0xc2,
	0x82, 0xc5, 0x4b, 0xbb, 0xf3, 0xe6, 0xcd, 0xcc, 0xbe, 0x37, 0x0b, 0x77, 0x9c, 0x4d, 0x50, 0x60,
	0x3c, 0x1b, 0x0f, 0xb1, 0xad, 0x9d, 0xa3, 0x38, 0x4c, 0xc2, 0xb6, 0xfc, 0x0a, 0x1d, 0x77, 0x24,
	0x64, 0x77, 0xa0, 0xfa, 0xcc, 0x26, 0xe8, 0xce, 0xd8, 0x38, 0x60, 0x83, 0x00, 0x3d, 0x7c, 0x9f,
	0xa2, 0x48, 0x68, 0x1d, 0x4a, 0xa3, 0x69, 0x10, 0x6c, 0x73, 0x26, 0x69, 0x92, 0x56, 0xd9, 0xdb,
	0xc7, 0xf6, 0xe3, 0xae, 0xa6, 0xbb, 0x70, 0x7d, 0x3f, 0x46, 0x21, 0xd2, 0x1a, 0x07, 0x68, 0x38,
	0xe7, 0x18, 0xbf, 0x0c, 0xe7, 0xbd, 0xe4, 0x4d, 0x25, 0x55, 0x75, 0x46, 0xc6, 0xfe, 0x21, 0x50,
	0x3b, 0x19, 0x2e, 0xa2, 0x90, 0x0b, 0xa4, 0x0d, 0x28, 0xb3, 0x14, 0x94, 0x0d, 0x4a, 0xde, 0x01,
	0x38, 0x33, 0xe7, 0xe2, 0xdc, 0x1c, 0xda, 0x82, 0x8a, 0x44, 0x35, 0x72, 0x4e, 0x92, 0x4f, 0xe1,
	0x3d, 0xd3, 0xe5, 0xa9, 0x36, 0x33, 0xaf, 0x31, 0x0f, 0x30, 0x35, 0xe1, 0x52, 0x44, 0x6c, 0x88,
	0x7d, 0xdf, 0x2c, 0x48, 0x46, 0x1a, 0xd2, 0x26, 0x5c, 0x6d, 0x6d, 0xee, 0x7d, 0x44, 0xe3, 0x18,
	0x85, 0x59, 0x6c, 0x92, 0x56, 0xce, 0xd3, 0x21, 0xdb, 0xdd, 0xc9, 0xd6, 0xfc, 0x53, 0xb2, 0xab,
	0x50, 0x18, 0x85, 0x53, 0xee, 0x2b, 0xc9, 0xbb, 0x80, 0x52, 0xc8, 0x6f, 0xab, 0x95, 0x40, 0x79,
	0xee, 0x2c, 0x09, 0x14, 0x5c, 0xbe, 0xe0, 0x82, 0x76, 0xa1, 0xd2, 0x17, 0x47, 0x2e, 0xd2, 0x9a,
	0x93, 0xb5, 0xd2, 0xfa, 0x8d, 0x93, 0x69, 0xb6, 0x6d, 0xd0, 0x07, 0xb8, 0x7e, 0xc2, 0xe4, 0xe8,
	0x4e, 0xaa, 0xc9, 0xe9, 0x8e, 0x55, 0x93, 0x3f, 0x57, 0xb7, 0x8d, 0xee, 0xfd, 0xe7, 0xda, 0x22,
	0xab, 0xb5, 0x45, 0xbe, 0xd7, 0x16, 0x59, 0x6e, 0x2c, 0x63, 0xb5, 0xb1, 0x8c, 0xaf, 0x8d, 0x65,
	0xbc, 0x36, 0xfe, 0x7b, 0x9a, 0x83, 0xa2, 0xfc, 0xdd, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xed,
	0x14, 0xb9, 0xb1, 0xc1, 0x02, 0x00, 0x00,
}

func (m *NameAvailableRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameAvailableRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameAvailableRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FullName) > 0 {
		i -= len(m.FullName)
		copy(dAtA[i:], m.FullName)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.FullName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NameByAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameByAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameByAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OwnerScwEthAddress) > 0 {
		i -= len(m.OwnerScwEthAddress)
		copy(dAtA[i:], m.OwnerScwEthAddress)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.OwnerScwEthAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NameAvailableResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameAvailableResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameAvailableResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NameExpires != 0 {
		i = encodeVarintNameservice(dAtA, i, uint64(m.NameExpires))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SpaceId) > 0 {
		i -= len(m.SpaceId)
		copy(dAtA[i:], m.SpaceId)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.SpaceId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.OwnerAnyAddress) > 0 {
		i -= len(m.OwnerAnyAddress)
		copy(dAtA[i:], m.OwnerAnyAddress)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.OwnerAnyAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OwnerEthAddress) > 0 {
		i -= len(m.OwnerEthAddress)
		copy(dAtA[i:], m.OwnerEthAddress)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.OwnerEthAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OwnerScwEthAddress) > 0 {
		i -= len(m.OwnerScwEthAddress)
		copy(dAtA[i:], m.OwnerScwEthAddress)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.OwnerScwEthAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.Available {
		i--
		if m.Available {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NameByAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NameByAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NameByAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNameservice(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Found {
		i--
		if m.Found {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNameservice(dAtA []byte, offset int, v uint64) int {
	offset -= sovNameservice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NameAvailableRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FullName)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	return n
}

func (m *NameByAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OwnerScwEthAddress)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	return n
}

func (m *NameAvailableResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Available {
		n += 2
	}
	l = len(m.OwnerScwEthAddress)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	l = len(m.OwnerEthAddress)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	l = len(m.OwnerAnyAddress)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	l = len(m.SpaceId)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	if m.NameExpires != 0 {
		n += 1 + sovNameservice(uint64(m.NameExpires))
	}
	return n
}

func (m *NameByAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Found {
		n += 2
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNameservice(uint64(l))
	}
	return n
}

func sovNameservice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNameservice(x uint64) (n int) {
	return sovNameservice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NameAvailableRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNameservice
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
			return fmt.Errorf("proto: NameAvailableRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameAvailableRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FullName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FullName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNameservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNameservice
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
func (m *NameByAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNameservice
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
			return fmt.Errorf("proto: NameByAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameByAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerScwEthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerScwEthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNameservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNameservice
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
func (m *NameAvailableResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNameservice
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
			return fmt.Errorf("proto: NameAvailableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameAvailableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Available", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
			m.Available = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerScwEthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerScwEthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerEthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerEthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerAnyAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OwnerAnyAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NameExpires", wireType)
			}
			m.NameExpires = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NameExpires |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNameservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNameservice
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
func (m *NameByAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNameservice
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
			return fmt.Errorf("proto: NameByAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NameByAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Found", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
			m.Found = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNameservice
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
				return ErrInvalidLengthNameservice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNameservice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNameservice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNameservice
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
func skipNameservice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNameservice
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
					return 0, ErrIntOverflowNameservice
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
					return 0, ErrIntOverflowNameservice
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
				return 0, ErrInvalidLengthNameservice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNameservice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNameservice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNameservice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNameservice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNameservice = fmt.Errorf("proto: unexpected end of group")
)
