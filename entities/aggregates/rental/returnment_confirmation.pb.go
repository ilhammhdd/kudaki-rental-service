// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aggregates/rental/returnment_confirmation.proto

package rental

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

type ReturnmentConfirmation struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderUuid            string               `protobuf:"bytes,2,opt,name=order_uuid,json=orderUuid,proto3" json:"order_uuid,omitempty"`
	TenantUserUuid       string               `protobuf:"bytes,3,opt,name=tenant_user_uuid,json=tenantUserUuid,proto3" json:"tenant_user_uuid,omitempty"`
	OwnerUserUuid        string               `protobuf:"bytes,4,opt,name=owner_user_uuid,json=ownerUserUuid,proto3" json:"owner_user_uuid,omitempty"`
	TenantConfirmed      bool                 `protobuf:"varint,5,opt,name=tenant_confirmed,json=tenantConfirmed,proto3" json:"tenant_confirmed,omitempty"`
	OwnerConfirmed       bool                 `protobuf:"varint,6,opt,name=owner_confirmed,json=ownerConfirmed,proto3" json:"owner_confirmed,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ReturnmentConfirmation) Reset()         { *m = ReturnmentConfirmation{} }
func (m *ReturnmentConfirmation) String() string { return proto.CompactTextString(m) }
func (*ReturnmentConfirmation) ProtoMessage()    {}
func (*ReturnmentConfirmation) Descriptor() ([]byte, []int) {
	return fileDescriptor_24385fc597e09b04, []int{0}
}

func (m *ReturnmentConfirmation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnmentConfirmation.Unmarshal(m, b)
}
func (m *ReturnmentConfirmation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnmentConfirmation.Marshal(b, m, deterministic)
}
func (m *ReturnmentConfirmation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnmentConfirmation.Merge(m, src)
}
func (m *ReturnmentConfirmation) XXX_Size() int {
	return xxx_messageInfo_ReturnmentConfirmation.Size(m)
}
func (m *ReturnmentConfirmation) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnmentConfirmation.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnmentConfirmation proto.InternalMessageInfo

func (m *ReturnmentConfirmation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ReturnmentConfirmation) GetOrderUuid() string {
	if m != nil {
		return m.OrderUuid
	}
	return ""
}

func (m *ReturnmentConfirmation) GetTenantUserUuid() string {
	if m != nil {
		return m.TenantUserUuid
	}
	return ""
}

func (m *ReturnmentConfirmation) GetOwnerUserUuid() string {
	if m != nil {
		return m.OwnerUserUuid
	}
	return ""
}

func (m *ReturnmentConfirmation) GetTenantConfirmed() bool {
	if m != nil {
		return m.TenantConfirmed
	}
	return false
}

func (m *ReturnmentConfirmation) GetOwnerConfirmed() bool {
	if m != nil {
		return m.OwnerConfirmed
	}
	return false
}

func (m *ReturnmentConfirmation) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ReturnmentConfirmation)(nil), "aggregates.rental.ReturnmentConfirmation")
}

func init() {
	proto.RegisterFile("aggregates/rental/returnment_confirmation.proto", fileDescriptor_24385fc597e09b04)
}

var fileDescriptor_24385fc597e09b04 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd1, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0x06, 0x70, 0xda, 0xfd, 0xff, 0xd3, 0x45, 0xdc, 0xb4, 0x07, 0x29, 0x03, 0xb1, 0x78, 0xd0,
	0x7a, 0x58, 0x03, 0x7a, 0xf2, 0xa8, 0x43, 0xbc, 0x17, 0x77, 0xf1, 0x52, 0xb2, 0xe6, 0x5d, 0xf7,
	0xb2, 0x25, 0x19, 0xe9, 0x1b, 0xfd, 0x4a, 0x7e, 0x4c, 0x31, 0x59, 0x3b, 0xc1, 0xeb, 0xc3, 0xef,
	0x7d, 0x4a, 0x9f, 0x30, 0x2e, 0x9a, 0xc6, 0x42, 0x23, 0x08, 0x5a, 0x6e, 0x41, 0x93, 0xd8, 0x72,
	0x0b, 0xe4, 0xac, 0x56, 0xa0, 0xa9, 0xaa, 0x8d, 0x5e, 0xa1, 0x55, 0x82, 0xd0, 0xe8, 0x62, 0x67,
	0x0d, 0x99, 0xe4, 0xfc, 0x70, 0x50, 0x84, 0x83, 0xe9, 0x55, 0x63, 0x4c, 0xb3, 0x05, 0xee, 0xc1,
	0xd2, 0xad, 0x38, 0xa1, 0x82, 0x96, 0x84, 0xda, 0x85, 0x9b, 0xeb, 0xaf, 0x98, 0x5d, 0x94, 0x7d,
	0xeb, 0xfc, 0x57, 0x69, 0x32, 0x66, 0x31, 0xca, 0x34, 0xca, 0xa2, 0x7c, 0x50, 0xc6, 0x28, 0x93,
	0x4b, 0xc6, 0x8c, 0x95, 0x60, 0x2b, 0xe7, 0x50, 0xa6, 0x71, 0x16, 0xe5, 0xa3, 0x72, 0xe4, 0x93,
	0x85, 0x43, 0x99, 0xe4, 0xec, 0x8c, 0x40, 0x0b, 0x4d, 0x95, 0x6b, 0x3b, 0x34, 0xf0, 0x68, 0x1c,
	0xf2, 0x45, 0xbb, 0x97, 0x37, 0x6c, 0x62, 0x3e, 0xf5, 0x8f, 0xe9, 0xe1, 0x3f, 0x0f, 0x4f, 0x7d,
	0xdc, 0xbb, 0xbb, 0xbe, 0x71, 0xff, 0xb3, 0x20, 0xd3, 0xff, 0x59, 0x94, 0x1f, 0x97, 0x93, 0x90,
	0xcf, 0xbb, 0x38, 0xb9, 0xed, 0x2a, 0x0f, 0x72, 0xe8, 0xe5, 0xd8, 0xc7, 0x07, 0xf8, 0xc8, 0x58,
	0x6d, 0x41, 0x10, 0xc8, 0x4a, 0x50, 0x7a, 0x94, 0x45, 0xf9, 0xc9, 0xfd, 0xb4, 0x08, 0x2b, 0x15,
	0xdd, 0x4a, 0xc5, 0x5b, 0xb7, 0x52, 0x39, 0xda, 0xeb, 0x27, 0x7a, 0x7e, 0x7d, 0x7f, 0x69, 0x90,
	0xd6, 0x6e, 0x59, 0xd4, 0x46, 0x71, 0xdc, 0xae, 0x85, 0x52, 0x6b, 0x29, 0xf9, 0xc6, 0x49, 0xb1,
	0xc1, 0x59, 0x58, 0x7c, 0xd6, 0x82, 0xfd, 0xc0, 0x1a, 0x38, 0x68, 0x42, 0x42, 0x68, 0xff, 0xbe,
	0xe2, 0x72, 0xe8, 0xbf, 0xf3, 0xf0, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x00, 0xe8, 0x03, 0xe1,
	0x01, 0x00, 0x00,
}
