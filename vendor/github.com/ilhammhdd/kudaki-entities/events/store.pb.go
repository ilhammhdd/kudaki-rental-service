// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/store.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	store "github.com/ilhammhdd/kudaki-entities/store"
	user "github.com/ilhammhdd/kudaki-entities/user"
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

type AddStorefrontItemRequested struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AddStorefrontItemRequested) Reset()         { *m = AddStorefrontItemRequested{} }
func (m *AddStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*AddStorefrontItemRequested) ProtoMessage()    {}
func (*AddStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{0}
}

func (m *AddStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddStorefrontItemRequested.Unmarshal(m, b)
}
func (m *AddStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *AddStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddStorefrontItemRequested.Merge(m, src)
}
func (m *AddStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_AddStorefrontItemRequested.Size(m)
}
func (m *AddStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_AddStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_AddStorefrontItemRequested proto.InternalMessageInfo

func (m *AddStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AddStorefrontItemRequested) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type StorefrontItemAdded struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorefrontItemAdded) Reset()         { *m = StorefrontItemAdded{} }
func (m *StorefrontItemAdded) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemAdded) ProtoMessage()    {}
func (*StorefrontItemAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{1}
}

func (m *StorefrontItemAdded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemAdded.Unmarshal(m, b)
}
func (m *StorefrontItemAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemAdded.Marshal(b, m, deterministic)
}
func (m *StorefrontItemAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemAdded.Merge(m, src)
}
func (m *StorefrontItemAdded) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemAdded.Size(m)
}
func (m *StorefrontItemAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemAdded.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemAdded proto.InternalMessageInfo

func (m *StorefrontItemAdded) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemAdded) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemAdded) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

type DeleteStorefrontItemRequested struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeleteStorefrontItemRequested) Reset()         { *m = DeleteStorefrontItemRequested{} }
func (m *DeleteStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*DeleteStorefrontItemRequested) ProtoMessage()    {}
func (*DeleteStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{2}
}

func (m *DeleteStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Unmarshal(m, b)
}
func (m *DeleteStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *DeleteStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteStorefrontItemRequested.Merge(m, src)
}
func (m *DeleteStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_DeleteStorefrontItemRequested.Size(m)
}
func (m *DeleteStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteStorefrontItemRequested proto.InternalMessageInfo

func (m *DeleteStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *DeleteStorefrontItemRequested) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

type StorefrontItemDeleted struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorefrontItemDeleted) Reset()         { *m = StorefrontItemDeleted{} }
func (m *StorefrontItemDeleted) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemDeleted) ProtoMessage()    {}
func (*StorefrontItemDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{3}
}

func (m *StorefrontItemDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemDeleted.Unmarshal(m, b)
}
func (m *StorefrontItemDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemDeleted.Marshal(b, m, deterministic)
}
func (m *StorefrontItemDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemDeleted.Merge(m, src)
}
func (m *StorefrontItemDeleted) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemDeleted.Size(m)
}
func (m *StorefrontItemDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemDeleted proto.InternalMessageInfo

func (m *StorefrontItemDeleted) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemDeleted) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemDeleted) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

type UpdateStorefrontItemRequested struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	User                 *user.User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateStorefrontItemRequested) Reset()         { *m = UpdateStorefrontItemRequested{} }
func (m *UpdateStorefrontItemRequested) String() string { return proto.CompactTextString(m) }
func (*UpdateStorefrontItemRequested) ProtoMessage()    {}
func (*UpdateStorefrontItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{4}
}

func (m *UpdateStorefrontItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Unmarshal(m, b)
}
func (m *UpdateStorefrontItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Marshal(b, m, deterministic)
}
func (m *UpdateStorefrontItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStorefrontItemRequested.Merge(m, src)
}
func (m *UpdateStorefrontItemRequested) XXX_Size() int {
	return xxx_messageInfo_UpdateStorefrontItemRequested.Size(m)
}
func (m *UpdateStorefrontItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStorefrontItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStorefrontItemRequested proto.InternalMessageInfo

func (m *UpdateStorefrontItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *UpdateStorefrontItemRequested) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *UpdateStorefrontItemRequested) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type StorefrontItemUpdated struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	User                 *user.User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *StorefrontItemUpdated) Reset()         { *m = StorefrontItemUpdated{} }
func (m *StorefrontItemUpdated) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemUpdated) ProtoMessage()    {}
func (*StorefrontItemUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{5}
}

func (m *StorefrontItemUpdated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemUpdated.Unmarshal(m, b)
}
func (m *StorefrontItemUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemUpdated.Marshal(b, m, deterministic)
}
func (m *StorefrontItemUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemUpdated.Merge(m, src)
}
func (m *StorefrontItemUpdated) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemUpdated.Size(m)
}
func (m *StorefrontItemUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemUpdated proto.InternalMessageInfo

func (m *StorefrontItemUpdated) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemUpdated) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *StorefrontItemUpdated) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *StorefrontItemUpdated) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

type RetrieveStorefrontItemsRequested struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	From                 int32      `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	Limit                int32      `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RetrieveStorefrontItemsRequested) Reset()         { *m = RetrieveStorefrontItemsRequested{} }
func (m *RetrieveStorefrontItemsRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveStorefrontItemsRequested) ProtoMessage()    {}
func (*RetrieveStorefrontItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{6}
}

func (m *RetrieveStorefrontItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Unmarshal(m, b)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveStorefrontItemsRequested.Merge(m, src)
}
func (m *RetrieveStorefrontItemsRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveStorefrontItemsRequested.Size(m)
}
func (m *RetrieveStorefrontItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveStorefrontItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveStorefrontItemsRequested proto.InternalMessageInfo

func (m *RetrieveStorefrontItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveStorefrontItemsRequested) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RetrieveStorefrontItemsRequested) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *RetrieveStorefrontItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type StorefrontItemsRetrieved struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Items                *store.Items `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
	EventStatus          *Status      `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	First                int32        `protobuf:"varint,4,opt,name=first,proto3" json:"first,omitempty"`
	Limit                int32        `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Last                 int32        `protobuf:"varint,6,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StorefrontItemsRetrieved) Reset()         { *m = StorefrontItemsRetrieved{} }
func (m *StorefrontItemsRetrieved) String() string { return proto.CompactTextString(m) }
func (*StorefrontItemsRetrieved) ProtoMessage()    {}
func (*StorefrontItemsRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{7}
}

func (m *StorefrontItemsRetrieved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorefrontItemsRetrieved.Unmarshal(m, b)
}
func (m *StorefrontItemsRetrieved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorefrontItemsRetrieved.Marshal(b, m, deterministic)
}
func (m *StorefrontItemsRetrieved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorefrontItemsRetrieved.Merge(m, src)
}
func (m *StorefrontItemsRetrieved) XXX_Size() int {
	return xxx_messageInfo_StorefrontItemsRetrieved.Size(m)
}
func (m *StorefrontItemsRetrieved) XXX_DiscardUnknown() {
	xxx_messageInfo_StorefrontItemsRetrieved.DiscardUnknown(m)
}

var xxx_messageInfo_StorefrontItemsRetrieved proto.InternalMessageInfo

func (m *StorefrontItemsRetrieved) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *StorefrontItemsRetrieved) GetItems() *store.Items {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *StorefrontItemsRetrieved) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *StorefrontItemsRetrieved) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *StorefrontItemsRetrieved) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *StorefrontItemsRetrieved) GetLast() int32 {
	if m != nil {
		return m.Last
	}
	return 0
}

type RetrieveItemsRequested struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	From                 int32    `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Limit                int32    `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetrieveItemsRequested) Reset()         { *m = RetrieveItemsRequested{} }
func (m *RetrieveItemsRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveItemsRequested) ProtoMessage()    {}
func (*RetrieveItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{8}
}

func (m *RetrieveItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveItemsRequested.Unmarshal(m, b)
}
func (m *RetrieveItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveItemsRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveItemsRequested.Merge(m, src)
}
func (m *RetrieveItemsRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveItemsRequested.Size(m)
}
func (m *RetrieveItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveItemsRequested proto.InternalMessageInfo

func (m *RetrieveItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveItemsRequested) GetFrom() int32 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *RetrieveItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ItemsRetrieved struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Items                *store.Items `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
	EventStatus          *Status      `protobuf:"bytes,3,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	First                int32        `protobuf:"varint,4,opt,name=first,proto3" json:"first,omitempty"`
	Limit                int32        `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Last                 int32        `protobuf:"varint,6,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemsRetrieved) Reset()         { *m = ItemsRetrieved{} }
func (m *ItemsRetrieved) String() string { return proto.CompactTextString(m) }
func (*ItemsRetrieved) ProtoMessage()    {}
func (*ItemsRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{9}
}

func (m *ItemsRetrieved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemsRetrieved.Unmarshal(m, b)
}
func (m *ItemsRetrieved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemsRetrieved.Marshal(b, m, deterministic)
}
func (m *ItemsRetrieved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemsRetrieved.Merge(m, src)
}
func (m *ItemsRetrieved) XXX_Size() int {
	return xxx_messageInfo_ItemsRetrieved.Size(m)
}
func (m *ItemsRetrieved) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemsRetrieved.DiscardUnknown(m)
}

var xxx_messageInfo_ItemsRetrieved proto.InternalMessageInfo

func (m *ItemsRetrieved) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ItemsRetrieved) GetItems() *store.Items {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemsRetrieved) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *ItemsRetrieved) GetFirst() int32 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *ItemsRetrieved) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ItemsRetrieved) GetLast() int32 {
	if m != nil {
		return m.Last
	}
	return 0
}

type RetrieveItemRequested struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	ItemUuid             string     `protobuf:"bytes,2,opt,name=item_uuid,json=itemUuid,proto3" json:"item_uuid,omitempty"`
	User                 *user.User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RetrieveItemRequested) Reset()         { *m = RetrieveItemRequested{} }
func (m *RetrieveItemRequested) String() string { return proto.CompactTextString(m) }
func (*RetrieveItemRequested) ProtoMessage()    {}
func (*RetrieveItemRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{10}
}

func (m *RetrieveItemRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetrieveItemRequested.Unmarshal(m, b)
}
func (m *RetrieveItemRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetrieveItemRequested.Marshal(b, m, deterministic)
}
func (m *RetrieveItemRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetrieveItemRequested.Merge(m, src)
}
func (m *RetrieveItemRequested) XXX_Size() int {
	return xxx_messageInfo_RetrieveItemRequested.Size(m)
}
func (m *RetrieveItemRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_RetrieveItemRequested.DiscardUnknown(m)
}

var xxx_messageInfo_RetrieveItemRequested proto.InternalMessageInfo

func (m *RetrieveItemRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *RetrieveItemRequested) GetItemUuid() string {
	if m != nil {
		return m.ItemUuid
	}
	return ""
}

func (m *RetrieveItemRequested) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ItemRetrieved struct {
	Uid                  string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Item                 *store.Item `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	User                 *user.User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	EventStatus          *Status     `protobuf:"bytes,4,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ItemRetrieved) Reset()         { *m = ItemRetrieved{} }
func (m *ItemRetrieved) String() string { return proto.CompactTextString(m) }
func (*ItemRetrieved) ProtoMessage()    {}
func (*ItemRetrieved) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{11}
}

func (m *ItemRetrieved) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRetrieved.Unmarshal(m, b)
}
func (m *ItemRetrieved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRetrieved.Marshal(b, m, deterministic)
}
func (m *ItemRetrieved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRetrieved.Merge(m, src)
}
func (m *ItemRetrieved) XXX_Size() int {
	return xxx_messageInfo_ItemRetrieved.Size(m)
}
func (m *ItemRetrieved) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRetrieved.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRetrieved proto.InternalMessageInfo

func (m *ItemRetrieved) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ItemRetrieved) GetItem() *store.Item {
	if m != nil {
		return m.Item
	}
	return nil
}

func (m *ItemRetrieved) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ItemRetrieved) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

type SearchItemsRequested struct {
	Uid                  string     `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Keyword              string     `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	From                 uint64     `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	Limit                int32      `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SearchItemsRequested) Reset()         { *m = SearchItemsRequested{} }
func (m *SearchItemsRequested) String() string { return proto.CompactTextString(m) }
func (*SearchItemsRequested) ProtoMessage()    {}
func (*SearchItemsRequested) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{12}
}

func (m *SearchItemsRequested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchItemsRequested.Unmarshal(m, b)
}
func (m *SearchItemsRequested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchItemsRequested.Marshal(b, m, deterministic)
}
func (m *SearchItemsRequested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchItemsRequested.Merge(m, src)
}
func (m *SearchItemsRequested) XXX_Size() int {
	return xxx_messageInfo_SearchItemsRequested.Size(m)
}
func (m *SearchItemsRequested) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchItemsRequested.DiscardUnknown(m)
}

var xxx_messageInfo_SearchItemsRequested proto.InternalMessageInfo

func (m *SearchItemsRequested) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SearchItemsRequested) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *SearchItemsRequested) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *SearchItemsRequested) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *SearchItemsRequested) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type ItemsSearched struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	User                 *user.User   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Keyword              string       `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Items                *store.Items `protobuf:"bytes,4,opt,name=items,proto3" json:"items,omitempty"`
	EventStatus          *Status      `protobuf:"bytes,5,opt,name=event_status,json=eventStatus,proto3" json:"event_status,omitempty"`
	First                uint64       `protobuf:"varint,6,opt,name=first,proto3" json:"first,omitempty"`
	Limit                int32        `protobuf:"varint,8,opt,name=limit,proto3" json:"limit,omitempty"`
	Last                 uint64       `protobuf:"varint,7,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ItemsSearched) Reset()         { *m = ItemsSearched{} }
func (m *ItemsSearched) String() string { return proto.CompactTextString(m) }
func (*ItemsSearched) ProtoMessage()    {}
func (*ItemsSearched) Descriptor() ([]byte, []int) {
	return fileDescriptor_4f52bba9433e5948, []int{13}
}

func (m *ItemsSearched) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemsSearched.Unmarshal(m, b)
}
func (m *ItemsSearched) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemsSearched.Marshal(b, m, deterministic)
}
func (m *ItemsSearched) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemsSearched.Merge(m, src)
}
func (m *ItemsSearched) XXX_Size() int {
	return xxx_messageInfo_ItemsSearched.Size(m)
}
func (m *ItemsSearched) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemsSearched.DiscardUnknown(m)
}

var xxx_messageInfo_ItemsSearched proto.InternalMessageInfo

func (m *ItemsSearched) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ItemsSearched) GetUser() *user.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ItemsSearched) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func (m *ItemsSearched) GetItems() *store.Items {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *ItemsSearched) GetEventStatus() *Status {
	if m != nil {
		return m.EventStatus
	}
	return nil
}

func (m *ItemsSearched) GetFirst() uint64 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *ItemsSearched) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ItemsSearched) GetLast() uint64 {
	if m != nil {
		return m.Last
	}
	return 0
}

func init() {
	proto.RegisterType((*AddStorefrontItemRequested)(nil), "event.AddStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemAdded)(nil), "event.StorefrontItemAdded")
	proto.RegisterType((*DeleteStorefrontItemRequested)(nil), "event.DeleteStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemDeleted)(nil), "event.StorefrontItemDeleted")
	proto.RegisterType((*UpdateStorefrontItemRequested)(nil), "event.UpdateStorefrontItemRequested")
	proto.RegisterType((*StorefrontItemUpdated)(nil), "event.StorefrontItemUpdated")
	proto.RegisterType((*RetrieveStorefrontItemsRequested)(nil), "event.RetrieveStorefrontItemsRequested")
	proto.RegisterType((*StorefrontItemsRetrieved)(nil), "event.StorefrontItemsRetrieved")
	proto.RegisterType((*RetrieveItemsRequested)(nil), "event.RetrieveItemsRequested")
	proto.RegisterType((*ItemsRetrieved)(nil), "event.ItemsRetrieved")
	proto.RegisterType((*RetrieveItemRequested)(nil), "event.RetrieveItemRequested")
	proto.RegisterType((*ItemRetrieved)(nil), "event.ItemRetrieved")
	proto.RegisterType((*SearchItemsRequested)(nil), "event.SearchItemsRequested")
	proto.RegisterType((*ItemsSearched)(nil), "event.ItemsSearched")
}

func init() { proto.RegisterFile("events/store.proto", fileDescriptor_4f52bba9433e5948) }

var fileDescriptor_4f52bba9433e5948 = []byte{
	// 539 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x96, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb5, 0x89, 0x9d, 0x36, 0x53, 0x0a, 0x95, 0x93, 0x20, 0x2b, 0xa8, 0x52, 0xe4, 0x0b,
	0x91, 0xaa, 0xda, 0x08, 0x9e, 0xa0, 0x88, 0x0b, 0x57, 0x87, 0x48, 0x08, 0x0e, 0x95, 0xdb, 0x9d,
	0x90, 0x55, 0xe2, 0xb8, 0xec, 0xae, 0x8b, 0x38, 0x17, 0xc4, 0x13, 0xf0, 0x0a, 0x1c, 0x79, 0x07,
	0xce, 0xbc, 0x14, 0xda, 0xd9, 0x38, 0x75, 0x93, 0x50, 0xd3, 0x96, 0x22, 0x7a, 0xa9, 0x76, 0xa6,
	0xb3, 0xbb, 0xdf, 0xff, 0xe7, 0x5f, 0xc9, 0xe0, 0xe1, 0x29, 0xce, 0xb4, 0x8a, 0x94, 0xce, 0x24,
	0x86, 0x27, 0x32, 0xd3, 0x99, 0xe7, 0x52, 0xaf, 0xbb, 0x43, 0xbd, 0x48, 0x68, 0x4c, 0xed, 0x3f,
	0xba, 0xad, 0xc5, 0x70, 0xa2, 0x73, 0x35, 0x6f, 0x3e, 0xc8, 0x15, 0xca, 0xc8, 0xfc, 0xb1, 0x8d,
	0xe0, 0x35, 0x74, 0x0f, 0x38, 0x1f, 0x98, 0xcd, 0x23, 0x99, 0xcd, 0xf4, 0x4b, 0x8d, 0x69, 0x8c,
	0xef, 0x73, 0x54, 0x1a, 0xb9, 0xb7, 0x03, 0xf5, 0x5c, 0x70, 0x9f, 0xf5, 0x58, 0xbf, 0x19, 0x9b,
	0xa5, 0xd7, 0x07, 0xc7, 0xdc, 0xe1, 0xd7, 0x7b, 0xac, 0xbf, 0xf5, 0xb4, 0x1d, 0xe2, 0x4c, 0x0b,
	0x2d, 0x50, 0x85, 0x96, 0x89, 0xb6, 0xd3, 0x44, 0x70, 0xc6, 0xa0, 0x75, 0xf1, 0xdc, 0x03, 0xce,
	0x6f, 0x76, 0xa6, 0xf7, 0x04, 0xee, 0x91, 0xaa, 0x43, 0x2b, 0xca, 0x77, 0x68, 0xc7, 0x76, 0x48,
	0xcd, 0x70, 0x40, 0xcd, 0x78, 0x8b, 0x2a, 0x5b, 0x04, 0x6f, 0x61, 0xf7, 0x05, 0x4e, 0x51, 0xe3,
	0xd5, 0x25, 0xd6, 0x2a, 0x25, 0x7e, 0x66, 0xd0, 0xb9, 0x78, 0xae, 0xbd, 0xeb, 0x46, 0xa7, 0xae,
	0x88, 0xac, 0x57, 0x8a, 0xfc, 0xc4, 0x60, 0x77, 0x78, 0xc2, 0x93, 0x5b, 0x51, 0xe9, 0x3d, 0x06,
	0xc7, 0x04, 0x66, 0xce, 0xd1, 0x3a, 0x9f, 0xa4, 0x18, 0x0d, 0x15, 0xca, 0x98, 0x06, 0x82, 0xef,
	0x2b, 0x76, 0x58, 0xa8, 0x7f, 0x73, 0xfd, 0x35, 0xc2, 0xf1, 0x85, 0x41, 0x2f, 0x46, 0x2d, 0x05,
	0x9e, 0x2e, 0x39, 0xa7, 0x2e, 0xb3, 0xae, 0x20, 0xaa, 0x55, 0x11, 0x79, 0xe0, 0x8c, 0x64, 0x66,
	0x83, 0xed, 0xc6, 0xb4, 0xf6, 0xda, 0xe0, 0x4e, 0x45, 0x2a, 0x34, 0xe1, 0xb9, 0xb1, 0x2d, 0x82,
	0x9f, 0x0c, 0xfc, 0x15, 0x02, 0x0b, 0xb6, 0x8e, 0x60, 0x0f, 0x5c, 0xe3, 0x8d, 0x9a, 0x23, 0x74,
	0xd6, 0xd9, 0xa7, 0x62, 0x3b, 0x73, 0xf5, 0x3c, 0x19, 0xc6, 0x91, 0x90, 0x6a, 0xc1, 0x48, 0xc5,
	0x39, 0xb9, 0x5b, 0x22, 0x37, 0x1a, 0xa7, 0x89, 0xd2, 0x7e, 0xc3, 0x6a, 0x34, 0xeb, 0xe0, 0x15,
	0x3c, 0x2c, 0xe8, 0x2b, 0xcd, 0x2c, 0x3c, 0xaa, 0xad, 0xf3, 0xa8, 0x5e, 0xf6, 0xe8, 0x07, 0x83,
	0xfb, 0x77, 0xdd, 0x99, 0x0c, 0x3a, 0x65, 0x67, 0x2e, 0x33, 0xe6, 0x11, 0x34, 0x0d, 0xe5, 0x61,
	0x6e, 0xfa, 0x35, 0xea, 0x6f, 0x9a, 0xc6, 0xb0, 0x1c, 0xc1, 0xca, 0x37, 0xf9, 0x8d, 0xc1, 0xb6,
	0xbd, 0xe9, 0xf7, 0x9e, 0xfd, 0x17, 0x6f, 0xf1, 0x2b, 0x83, 0xf6, 0x00, 0x13, 0x79, 0x3c, 0xfe,
	0x7b, 0xef, 0xcf, 0x87, 0x8d, 0x09, 0x7e, 0xfc, 0x90, 0x49, 0x4e, 0xc4, 0xcd, 0xb8, 0x28, 0x17,
	0xa9, 0x33, 0x5c, 0xce, 0x72, 0xea, 0xca, 0xbf, 0x62, 0x70, 0x56, 0xb3, 0x06, 0x2a, 0x0b, 0x77,
	0x5b, 0x40, 0x8b, 0xdc, 0x3a, 0xd7, 0xc8, 0xad, 0xfb, 0xe7, 0xb9, 0x6d, 0x90, 0xe0, 0xe5, 0xdc,
	0x6e, 0xae, 0xcb, 0xed, 0x86, 0xf5, 0xc6, 0xac, 0x9f, 0xef, 0xbf, 0xd9, 0x7b, 0x27, 0xf4, 0x38,
	0x3f, 0x0a, 0x8f, 0xb3, 0x34, 0x12, 0xd3, 0x71, 0x92, 0xa6, 0x63, 0xce, 0xa3, 0x49, 0xce, 0x93,
	0x89, 0xd8, 0x2f, 0x60, 0x23, 0xfb, 0xcd, 0x71, 0xd4, 0xa0, 0x8f, 0x8b, 0x67, 0xbf, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xd8, 0x2f, 0x8a, 0xb1, 0x08, 0x00, 0x00,
}