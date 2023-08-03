// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identityrepo/identityrepoproto/protos/identityrepo.proto

package identityrepoproto

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

type Data struct {
	// kind is a string representing the kind of data
	Kind string `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	// data is a byte payload
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// data signature
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff72e8611224dab, []int{0}
}
func (m *Data) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Data.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return m.Size()
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *Data) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Data) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DataPutRequest struct {
	// string representation of identity, must be equal handshake result
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// data to save
	Data []*Data `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *DataPutRequest) Reset()         { *m = DataPutRequest{} }
func (m *DataPutRequest) String() string { return proto.CompactTextString(m) }
func (*DataPutRequest) ProtoMessage()    {}
func (*DataPutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff72e8611224dab, []int{1}
}
func (m *DataPutRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPutRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPutRequest.Merge(m, src)
}
func (m *DataPutRequest) XXX_Size() int {
	return m.Size()
}
func (m *DataPutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataPutRequest proto.InternalMessageInfo

func (m *DataPutRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *DataPutRequest) GetData() []*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type DataDeleteRequest struct {
	// string representation of identity, must be equal handshake result
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// kinds of data to delete, if empty all kinds will be deleted
	Kinds []string `protobuf:"bytes,2,rep,name=kinds,proto3" json:"kinds,omitempty"`
}

func (m *DataDeleteRequest) Reset()         { *m = DataDeleteRequest{} }
func (m *DataDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DataDeleteRequest) ProtoMessage()    {}
func (*DataDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff72e8611224dab, []int{2}
}
func (m *DataDeleteRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataDeleteRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataDeleteRequest.Merge(m, src)
}
func (m *DataDeleteRequest) XXX_Size() int {
	return m.Size()
}
func (m *DataDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataDeleteRequest proto.InternalMessageInfo

func (m *DataDeleteRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *DataDeleteRequest) GetKinds() []string {
	if m != nil {
		return m.Kinds
	}
	return nil
}

type DataPullRequest struct {
	// string representation of identity wanted to request
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// kinds of data wanted to request
	Kinds []string `protobuf:"bytes,2,rep,name=kinds,proto3" json:"kinds,omitempty"`
}

func (m *DataPullRequest) Reset()         { *m = DataPullRequest{} }
func (m *DataPullRequest) String() string { return proto.CompactTextString(m) }
func (*DataPullRequest) ProtoMessage()    {}
func (*DataPullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff72e8611224dab, []int{3}
}
func (m *DataPullRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPullRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPullRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPullRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPullRequest.Merge(m, src)
}
func (m *DataPullRequest) XXX_Size() int {
	return m.Size()
}
func (m *DataPullRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPullRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataPullRequest proto.InternalMessageInfo

func (m *DataPullRequest) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *DataPullRequest) GetKinds() []string {
	if m != nil {
		return m.Kinds
	}
	return nil
}

type DataPullResponse struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// a list of data with requested kinds
	Data []*Data `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *DataPullResponse) Reset()         { *m = DataPullResponse{} }
func (m *DataPullResponse) String() string { return proto.CompactTextString(m) }
func (*DataPullResponse) ProtoMessage()    {}
func (*DataPullResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff72e8611224dab, []int{4}
}
func (m *DataPullResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataPullResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataPullResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataPullResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataPullResponse.Merge(m, src)
}
func (m *DataPullResponse) XXX_Size() int {
	return m.Size()
}
func (m *DataPullResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DataPullResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DataPullResponse proto.InternalMessageInfo

func (m *DataPullResponse) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *DataPullResponse) GetData() []*Data {
	if m != nil {
		return m.Data
	}
	return nil
}

type Ok struct {
}

func (m *Ok) Reset()         { *m = Ok{} }
func (m *Ok) String() string { return proto.CompactTextString(m) }
func (*Ok) ProtoMessage()    {}
func (*Ok) Descriptor() ([]byte, []int) {
	return fileDescriptor_fff72e8611224dab, []int{5}
}
func (m *Ok) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ok) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ok.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ok) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ok.Merge(m, src)
}
func (m *Ok) XXX_Size() int {
	return m.Size()
}
func (m *Ok) XXX_DiscardUnknown() {
	xxx_messageInfo_Ok.DiscardUnknown(m)
}

var xxx_messageInfo_Ok proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Data)(nil), "identityRepo.Data")
	proto.RegisterType((*DataPutRequest)(nil), "identityRepo.DataPutRequest")
	proto.RegisterType((*DataDeleteRequest)(nil), "identityRepo.DataDeleteRequest")
	proto.RegisterType((*DataPullRequest)(nil), "identityRepo.DataPullRequest")
	proto.RegisterType((*DataPullResponse)(nil), "identityRepo.DataPullResponse")
	proto.RegisterType((*Ok)(nil), "identityRepo.Ok")
}

func init() {
	proto.RegisterFile("identityrepo/identityrepoproto/protos/identityrepo.proto", fileDescriptor_fff72e8611224dab)
}

var fileDescriptor_fff72e8611224dab = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xc8, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xa9, 0x2c, 0x4a, 0x2d, 0xc8, 0xd7, 0x47, 0xe6, 0x14, 0x14, 0xe5, 0x97, 0xe4,
	0xeb, 0x83, 0xc9, 0x62, 0x14, 0x09, 0x3d, 0xb0, 0x98, 0x10, 0x0f, 0x4c, 0x2c, 0x28, 0xb5, 0x20,
	0x5f, 0xc9, 0x87, 0x8b, 0xc5, 0x25, 0xb1, 0x24, 0x51, 0x48, 0x88, 0x8b, 0x25, 0x3b, 0x33, 0x2f,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x06, 0x89, 0xa5, 0x24, 0x96, 0x24, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0x99, 0xe9, 0x79,
	0x89, 0x25, 0xa5, 0x45, 0xa9, 0x12, 0xcc, 0x60, 0x09, 0x84, 0x80, 0x52, 0x08, 0x17, 0x1f, 0xc8,
	0xb4, 0x80, 0xd2, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x29, 0x2e, 0x0e, 0x98,
	0x7d, 0x50, 0xb3, 0xe1, 0x7c, 0x21, 0x35, 0xb8, 0xf9, 0xcc, 0x1a, 0xdc, 0x46, 0x42, 0x7a, 0xc8,
	0x0e, 0xd3, 0x03, 0x99, 0x03, 0xb1, 0x53, 0xc9, 0x95, 0x4b, 0x10, 0xc4, 0x73, 0x49, 0xcd, 0x49,
	0x2d, 0x49, 0x25, 0xc6, 0x60, 0x11, 0x2e, 0x56, 0x90, 0x07, 0x8a, 0xc1, 0x26, 0x73, 0x06, 0x41,
	0x38, 0x4a, 0xce, 0x5c, 0xfc, 0x10, 0xc7, 0xe5, 0xe4, 0x90, 0x6f, 0x48, 0x18, 0x97, 0x00, 0xc2,
	0x90, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xaa, 0xf8, 0x91, 0x85, 0x8b, 0xc9, 0x3f, 0xdb, 0xe8,
	0x2a, 0x23, 0x17, 0x8f, 0x27, 0x92, 0x0a, 0x21, 0x4b, 0x2e, 0x76, 0x68, 0x80, 0x0a, 0xc9, 0x60,
	0xea, 0x45, 0x84, 0xb3, 0x94, 0x00, 0xaa, 0xac, 0x7f, 0xb6, 0x90, 0x3d, 0x17, 0x17, 0x22, 0xd4,
	0x84, 0xe4, 0x31, 0x75, 0xa3, 0x84, 0x27, 0x16, 0x03, 0x3c, 0xb9, 0x38, 0x60, 0x5e, 0x15, 0x92,
	0xc5, 0x66, 0x39, 0x3c, 0x1c, 0xa5, 0xe4, 0x70, 0x49, 0x43, 0x42, 0xc8, 0xc9, 0xe2, 0xc4, 0x23,
	0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2,
	0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xe4, 0xf0, 0xa7, 0xe3, 0x24, 0x36, 0x30, 0x65,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xbd, 0x63, 0xa6, 0xf0, 0x02, 0x00, 0x00,
}

func (m *Data) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Data) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Data) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Kind) > 0 {
		i -= len(m.Kind)
		copy(dAtA[i:], m.Kind)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Kind)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPutRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPutRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPutRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIdentityrepo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataDeleteRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataDeleteRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataDeleteRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Kinds) > 0 {
		for iNdEx := len(m.Kinds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Kinds[iNdEx])
			copy(dAtA[i:], m.Kinds[iNdEx])
			i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Kinds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPullRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPullRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPullRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Kinds) > 0 {
		for iNdEx := len(m.Kinds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Kinds[iNdEx])
			copy(dAtA[i:], m.Kinds[iNdEx])
			i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Kinds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataPullResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataPullResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataPullResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintIdentityrepo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintIdentityrepo(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Ok) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ok) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ok) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintIdentityrepo(dAtA []byte, offset int, v uint64) int {
	offset -= sovIdentityrepo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Data) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Kind)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	return n
}

func (m *DataPutRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovIdentityrepo(uint64(l))
		}
	}
	return n
}

func (m *DataDeleteRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	if len(m.Kinds) > 0 {
		for _, s := range m.Kinds {
			l = len(s)
			n += 1 + l + sovIdentityrepo(uint64(l))
		}
	}
	return n
}

func (m *DataPullRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	if len(m.Kinds) > 0 {
		for _, s := range m.Kinds {
			l = len(s)
			n += 1 + l + sovIdentityrepo(uint64(l))
		}
	}
	return n
}

func (m *DataPullResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovIdentityrepo(uint64(l))
	}
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovIdentityrepo(uint64(l))
		}
	}
	return n
}

func (m *Ok) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovIdentityrepo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIdentityrepo(x uint64) (n int) {
	return sovIdentityrepo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Data) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityrepo
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
			return fmt.Errorf("proto: Data: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Data: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kind", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kind = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityrepo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentityrepo
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
func (m *DataPutRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityrepo
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
			return fmt.Errorf("proto: DataPutRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPutRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &Data{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityrepo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentityrepo
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
func (m *DataDeleteRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityrepo
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
			return fmt.Errorf("proto: DataDeleteRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataDeleteRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kinds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kinds = append(m.Kinds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityrepo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentityrepo
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
func (m *DataPullRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityrepo
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
			return fmt.Errorf("proto: DataPullRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPullRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Kinds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Kinds = append(m.Kinds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityrepo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentityrepo
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
func (m *DataPullResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityrepo
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
			return fmt.Errorf("proto: DataPullResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataPullResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIdentityrepo
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
				return ErrInvalidLengthIdentityrepo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIdentityrepo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &Data{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityrepo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentityrepo
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
func (m *Ok) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIdentityrepo
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
			return fmt.Errorf("proto: Ok: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ok: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipIdentityrepo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIdentityrepo
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
func skipIdentityrepo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIdentityrepo
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
					return 0, ErrIntOverflowIdentityrepo
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
					return 0, ErrIntOverflowIdentityrepo
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
				return 0, ErrInvalidLengthIdentityrepo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIdentityrepo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIdentityrepo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIdentityrepo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIdentityrepo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIdentityrepo = fmt.Errorf("proto: unexpected end of group")
)
