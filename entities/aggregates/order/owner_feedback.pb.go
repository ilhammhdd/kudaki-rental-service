// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/order/owner_feedback.proto

package order

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OwnerFeedback struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Uuid                 string               `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	UserUuid             string               `protobuf:"bytes,3,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Order                *Order               `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty"`
	Rating               float64              `protobuf:"fixed64,5,opt,name=rating,proto3" json:"rating,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OwnerFeedback) Reset()         { *m = OwnerFeedback{} }
func (m *OwnerFeedback) String() string { return proto.CompactTextString(m) }
func (*OwnerFeedback) ProtoMessage()    {}
func (*OwnerFeedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac2c0bc013486200, []int{0}
}

func (m *OwnerFeedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OwnerFeedback.Unmarshal(m, b)
}
func (m *OwnerFeedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OwnerFeedback.Marshal(b, m, deterministic)
}
func (m *OwnerFeedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerFeedback.Merge(m, src)
}
func (m *OwnerFeedback) XXX_Size() int {
	return xxx_messageInfo_OwnerFeedback.Size(m)
}
func (m *OwnerFeedback) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerFeedback.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerFeedback proto.InternalMessageInfo

func (m *OwnerFeedback) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *OwnerFeedback) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *OwnerFeedback) GetUserUuid() string {
	if m != nil {
		return m.UserUuid
	}
	return ""
}

func (m *OwnerFeedback) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

func (m *OwnerFeedback) GetRating() float64 {
	if m != nil {
		return m.Rating
	}
	return 0
}

func (m *OwnerFeedback) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnerFeedback)(nil), "aggregates.order.OwnerFeedback")
}

func init() {
	proto.RegisterFile("aggregates/order/owner_feedback.proto", fileDescriptor_ac2c0bc013486200)
}

var fileDescriptor_ac2c0bc013486200 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4d, 0x4b, 0xf4, 0x30,
	0x14, 0x85, 0xc9, 0x7c, 0xf1, 0x4e, 0x5e, 0x14, 0xc9, 0x42, 0xcb, 0x28, 0x58, 0x04, 0xa1, 0x9b,
	0x26, 0xa0, 0x2b, 0x97, 0x8a, 0xcc, 0x76, 0xa0, 0xe8, 0xc6, 0x4d, 0x49, 0x9b, 0x3b, 0x69, 0xe8,
	0x47, 0x86, 0xf4, 0x46, 0xff, 0xa7, 0xbf, 0x48, 0x9a, 0x76, 0x10, 0xc6, 0xcd, 0x25, 0xc9, 0x73,
	0x72, 0x4e, 0x4e, 0xe8, 0xbd, 0xd4, 0xda, 0x81, 0x96, 0x08, 0xbd, 0xb0, 0x4e, 0x81, 0x13, 0xf6,
	0xab, 0x03, 0x97, 0xef, 0x01, 0x54, 0x21, 0xcb, 0x9a, 0x1f, 0x9c, 0x45, 0xcb, 0x2e, 0x7e, 0x65,
	0x3c, 0xc8, 0x36, 0xb7, 0xda, 0x5a, 0xdd, 0x80, 0x08, 0xbc, 0xf0, 0x7b, 0x81, 0xa6, 0x85, 0x1e,
	0x65, 0x7b, 0x18, 0xaf, 0x6c, 0x6e, 0xfe, 0x3a, 0x0f, 0x73, 0xa4, 0x77, 0xdf, 0x84, 0x9e, 0xed,
	0x86, 0xa4, 0xed, 0x14, 0xc4, 0xce, 0xe9, 0xcc, 0xa8, 0x88, 0xc4, 0x24, 0x99, 0x67, 0x33, 0xa3,
	0x18, 0xa3, 0x0b, 0xef, 0x8d, 0x8a, 0x66, 0x31, 0x49, 0xd6, 0x59, 0x58, 0xb3, 0x6b, 0xba, 0xf6,
	0x3d, 0xb8, 0x3c, 0x80, 0x79, 0x00, 0xff, 0x86, 0x83, 0xf7, 0x01, 0xa6, 0x74, 0x19, 0x12, 0xa2,
	0x45, 0x4c, 0x92, 0xff, 0x0f, 0x57, 0xfc, 0xf4, 0xcd, 0x7c, 0x37, 0xcc, 0x6c, 0x54, 0xb1, 0x4b,
	0xba, 0x72, 0x12, 0x4d, 0xa7, 0xa3, 0x65, 0x4c, 0x12, 0x92, 0x4d, 0x3b, 0xf6, 0x44, 0x69, 0xe9,
	0x40, 0x22, 0xa8, 0x5c, 0x62, 0xb4, 0x0a, 0x5e, 0x1b, 0x3e, 0xb6, 0xe5, 0xc7, 0xb6, 0xfc, 0xed,
	0xd8, 0x36, 0x5b, 0x4f, 0xea, 0x67, 0x7c, 0xd9, 0x7e, 0xbc, 0x6a, 0x83, 0x95, 0x2f, 0x78, 0x69,
	0x5b, 0x61, 0x9a, 0x4a, 0xb6, 0x6d, 0xa5, 0x94, 0xa8, 0xbd, 0x92, 0xb5, 0x49, 0x1d, 0x74, 0x28,
	0x9b, 0xb4, 0x07, 0xf7, 0x69, 0x4a, 0x10, 0xd0, 0xa1, 0x41, 0x03, 0xbd, 0x38, 0xfd, 0xa8, 0x62,
	0x15, 0x62, 0x1e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xbd, 0xfd, 0x32, 0x9d, 0x01, 0x00,
	0x00,
}
