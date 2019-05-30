// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/topics.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UserTopic int32

const (
	UserTopic_SIGNED_UP                           UserTopic = 0
	UserTopic_USER_VERIFICATION_EMAIL_SENT        UserTopic = 1
	UserTopic_SIGN_UP_REQUESTED                   UserTopic = 2
	UserTopic_VERIFY_USER_REQUESTED               UserTopic = 3
	UserTopic_LOGGED_IN                           UserTopic = 4
	UserTopic_LOGIN_REQUESTED                     UserTopic = 5
	UserTopic_USER_AUTHENTICATION_REQUESTED       UserTopic = 6
	UserTopic_USER_AUTHENTICATED                  UserTopic = 7
	UserTopic_RESET_PASSWORD_REQUESTED            UserTopic = 8
	UserTopic_PASSWORD_RESETED                    UserTopic = 9
	UserTopic_USER_AUTHORIZATION_REQUESTED        UserTopic = 10
	UserTopic_USER_AUTHORIZED                     UserTopic = 11
	UserTopic_USER_VERIFIED                       UserTopic = 12
	UserTopic_USER_RETRIEVED                      UserTopic = 13
	UserTopic_CHANGE_PASSWORD_REQUESTED           UserTopic = 14
	UserTopic_PASSWORD_CHANGED                    UserTopic = 15
	UserTopic_RESET_PASSWORD_EMAIL_SENT           UserTopic = 16
	UserTopic_SEND_RESET_PASSWORD_EMAIL_REQUESTED UserTopic = 17
)

var UserTopic_name = map[int32]string{
	0:  "SIGNED_UP",
	1:  "USER_VERIFICATION_EMAIL_SENT",
	2:  "SIGN_UP_REQUESTED",
	3:  "VERIFY_USER_REQUESTED",
	4:  "LOGGED_IN",
	5:  "LOGIN_REQUESTED",
	6:  "USER_AUTHENTICATION_REQUESTED",
	7:  "USER_AUTHENTICATED",
	8:  "RESET_PASSWORD_REQUESTED",
	9:  "PASSWORD_RESETED",
	10: "USER_AUTHORIZATION_REQUESTED",
	11: "USER_AUTHORIZED",
	12: "USER_VERIFIED",
	13: "USER_RETRIEVED",
	14: "CHANGE_PASSWORD_REQUESTED",
	15: "PASSWORD_CHANGED",
	16: "RESET_PASSWORD_EMAIL_SENT",
	17: "SEND_RESET_PASSWORD_EMAIL_REQUESTED",
}

var UserTopic_value = map[string]int32{
	"SIGNED_UP":                           0,
	"USER_VERIFICATION_EMAIL_SENT":        1,
	"SIGN_UP_REQUESTED":                   2,
	"VERIFY_USER_REQUESTED":               3,
	"LOGGED_IN":                           4,
	"LOGIN_REQUESTED":                     5,
	"USER_AUTHENTICATION_REQUESTED":       6,
	"USER_AUTHENTICATED":                  7,
	"RESET_PASSWORD_REQUESTED":            8,
	"PASSWORD_RESETED":                    9,
	"USER_AUTHORIZATION_REQUESTED":        10,
	"USER_AUTHORIZED":                     11,
	"USER_VERIFIED":                       12,
	"USER_RETRIEVED":                      13,
	"CHANGE_PASSWORD_REQUESTED":           14,
	"PASSWORD_CHANGED":                    15,
	"RESET_PASSWORD_EMAIL_SENT":           16,
	"SEND_RESET_PASSWORD_EMAIL_REQUESTED": 17,
}

func (x UserTopic) String() string {
	return proto.EnumName(UserTopic_name, int32(x))
}

func (UserTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{0}
}

type MountainTopic int32

const (
	MountainTopic_CREATE_MOUNTAIN_REQUESTED    MountainTopic = 0
	MountainTopic_MOUNTAIN_CREATED             MountainTopic = 1
	MountainTopic_RETRIEVE_MOUNTAINS_REQUESTED MountainTopic = 2
	MountainTopic_MOUNTAINS_RETRIEVED          MountainTopic = 3
)

var MountainTopic_name = map[int32]string{
	0: "CREATE_MOUNTAIN_REQUESTED",
	1: "MOUNTAIN_CREATED",
	2: "RETRIEVE_MOUNTAINS_REQUESTED",
	3: "MOUNTAINS_RETRIEVED",
}

var MountainTopic_value = map[string]int32{
	"CREATE_MOUNTAIN_REQUESTED":    0,
	"MOUNTAIN_CREATED":             1,
	"RETRIEVE_MOUNTAINS_REQUESTED": 2,
	"MOUNTAINS_RETRIEVED":          3,
}

func (x MountainTopic) String() string {
	return proto.EnumName(MountainTopic_name, int32(x))
}

func (MountainTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{1}
}

type StoreTopic int32

const (
	StoreTopic_ADD_STOREFRONT_ITEM_REQUESTED       StoreTopic = 0
	StoreTopic_STOREFRONT_ITEM_ADDED               StoreTopic = 1
	StoreTopic_DELETE_STOREFRONT_ITEM_REQUESTED    StoreTopic = 2
	StoreTopic_STOREFRONT_ITEM_DELETED             StoreTopic = 3
	StoreTopic_RETRIEVE_STOREFRONT_ITEMS_REQUESTED StoreTopic = 4
	StoreTopic_STOREFRONT_ITEMS_RETRIEVED          StoreTopic = 5
	StoreTopic_UPDATE_STOREFRONT_ITEM_REQUESTED    StoreTopic = 6
	StoreTopic_STOREFRONT_ITEM_UPDATED             StoreTopic = 7
	StoreTopic_RETRIEVE_ITEMS_REQUESTED            StoreTopic = 8
	StoreTopic_ITEMS_RETRIEVED                     StoreTopic = 9
	StoreTopic_RETRIEVE_ITEM_REQUESTED             StoreTopic = 10
	StoreTopic_ITEM_RETRIEVED                      StoreTopic = 11
	StoreTopic_SEARCH_ITEMS_REQUESTED              StoreTopic = 12
	StoreTopic_ITEMS_SEARCHED                      StoreTopic = 13
)

var StoreTopic_name = map[int32]string{
	0:  "ADD_STOREFRONT_ITEM_REQUESTED",
	1:  "STOREFRONT_ITEM_ADDED",
	2:  "DELETE_STOREFRONT_ITEM_REQUESTED",
	3:  "STOREFRONT_ITEM_DELETED",
	4:  "RETRIEVE_STOREFRONT_ITEMS_REQUESTED",
	5:  "STOREFRONT_ITEMS_RETRIEVED",
	6:  "UPDATE_STOREFRONT_ITEM_REQUESTED",
	7:  "STOREFRONT_ITEM_UPDATED",
	8:  "RETRIEVE_ITEMS_REQUESTED",
	9:  "ITEMS_RETRIEVED",
	10: "RETRIEVE_ITEM_REQUESTED",
	11: "ITEM_RETRIEVED",
	12: "SEARCH_ITEMS_REQUESTED",
	13: "ITEMS_SEARCHED",
}

var StoreTopic_value = map[string]int32{
	"ADD_STOREFRONT_ITEM_REQUESTED":       0,
	"STOREFRONT_ITEM_ADDED":               1,
	"DELETE_STOREFRONT_ITEM_REQUESTED":    2,
	"STOREFRONT_ITEM_DELETED":             3,
	"RETRIEVE_STOREFRONT_ITEMS_REQUESTED": 4,
	"STOREFRONT_ITEMS_RETRIEVED":          5,
	"UPDATE_STOREFRONT_ITEM_REQUESTED":    6,
	"STOREFRONT_ITEM_UPDATED":             7,
	"RETRIEVE_ITEMS_REQUESTED":            8,
	"ITEMS_RETRIEVED":                     9,
	"RETRIEVE_ITEM_REQUESTED":             10,
	"ITEM_RETRIEVED":                      11,
	"SEARCH_ITEMS_REQUESTED":              12,
	"ITEMS_SEARCHED":                      13,
}

func (x StoreTopic) String() string {
	return proto.EnumName(StoreTopic_name, int32(x))
}

func (StoreTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{2}
}

type RentalTopic int32

const (
	RentalTopic_CHECKOUT_REQUESTED RentalTopic = 0
	RentalTopic_CHECKEDOUT         RentalTopic = 1
	RentalTopic_CHECKOUT_REVIEWD   RentalTopic = 2
)

var RentalTopic_name = map[int32]string{
	0: "CHECKOUT_REQUESTED",
	1: "CHECKEDOUT",
	2: "CHECKOUT_REVIEWD",
}

var RentalTopic_value = map[string]int32{
	"CHECKOUT_REQUESTED": 0,
	"CHECKEDOUT":         1,
	"CHECKOUT_REVIEWD":   2,
}

func (x RentalTopic) String() string {
	return proto.EnumName(RentalTopic_name, int32(x))
}

func (RentalTopic) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c97f2476c050d00c, []int{3}
}

func init() {
	proto.RegisterEnum("event.topic.UserTopic", UserTopic_name, UserTopic_value)
	proto.RegisterEnum("event.topic.MountainTopic", MountainTopic_name, MountainTopic_value)
	proto.RegisterEnum("event.topic.StoreTopic", StoreTopic_name, StoreTopic_value)
	proto.RegisterEnum("event.topic.RentalTopic", RentalTopic_name, RentalTopic_value)
}

func init() { proto.RegisterFile("events/topics.proto", fileDescriptor_c97f2476c050d00c) }

var fileDescriptor_c97f2476c050d00c = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x61, 0x4f, 0x1a, 0x4d,
	0x10, 0x56, 0x51, 0x5f, 0x19, 0x04, 0xd7, 0x55, 0xf0, 0xc5, 0x4a, 0xa3, 0xb1, 0x89, 0x09, 0x8d,
	0xf0, 0xa1, 0xbf, 0xe0, 0xca, 0x8e, 0x70, 0x11, 0xee, 0xe8, 0xee, 0x1e, 0xa6, 0x7e, 0xd9, 0xa0,
	0x5c, 0xca, 0x45, 0xe1, 0x0c, 0x1c, 0xfd, 0xda, 0x5f, 0xd4, 0x7f, 0xd4, 0x1f, 0xd3, 0xec, 0x1d,
	0x07, 0x77, 0x57, 0xe2, 0xc7, 0x9d, 0x79, 0x9e, 0x99, 0xe7, 0x99, 0x99, 0x2c, 0x9c, 0xb8, 0x3f,
	0xdd, 0x69, 0x30, 0x6f, 0x06, 0xfe, 0x9b, 0xf7, 0x3c, 0x6f, 0xbc, 0xcd, 0xfc, 0xc0, 0xa7, 0x85,
	0x30, 0xd8, 0x08, 0x63, 0xf5, 0x3f, 0x39, 0xc8, 0x3b, 0x73, 0x77, 0x26, 0xf5, 0x8b, 0x16, 0x21,
	0x2f, 0xcc, 0xb6, 0x85, 0x4c, 0x39, 0x7d, 0xb2, 0x45, 0x2f, 0xe1, 0xc2, 0x11, 0xc8, 0xd5, 0x00,
	0xb9, 0x79, 0x67, 0xb6, 0x0c, 0x69, 0xda, 0x96, 0xc2, 0x9e, 0x61, 0x76, 0x95, 0x40, 0x4b, 0x92,
	0x6d, 0x5a, 0x86, 0x63, 0x4d, 0x50, 0x4e, 0x5f, 0x71, 0xfc, 0xe6, 0xa0, 0x90, 0xc8, 0xc8, 0x0e,
	0xad, 0x42, 0x39, 0xe4, 0x7c, 0x57, 0x21, 0x7f, 0x9d, 0xca, 0xe9, 0x16, 0x5d, 0xbb, 0xdd, 0x46,
	0xa6, 0x4c, 0x8b, 0xec, 0xd2, 0x13, 0x38, 0xea, 0xda, 0x6d, 0xd3, 0x4a, 0x60, 0xf6, 0xe8, 0x15,
	0xd4, 0x42, 0x9e, 0xe1, 0xc8, 0x0e, 0x5a, 0x32, 0xee, 0xbc, 0x86, 0xec, 0xd3, 0x0a, 0xd0, 0x2c,
	0x04, 0x19, 0xf9, 0x8f, 0x5e, 0xc0, 0xff, 0x1c, 0x05, 0x4a, 0xd5, 0x37, 0x84, 0x78, 0xb0, 0x39,
	0x4b, 0xb0, 0x0e, 0xe8, 0x29, 0x90, 0x44, 0x5c, 0xa0, 0x8e, 0xe6, 0x57, 0x36, 0x75, 0x2d, 0x9b,
	0x9b, 0x8f, 0xd9, 0x6e, 0xa0, 0x55, 0xa6, 0x10, 0xc8, 0x48, 0x81, 0x1e, 0x43, 0x31, 0x31, 0x1d,
	0x64, 0xe4, 0x90, 0x52, 0x28, 0x2d, 0x0d, 0x4b, 0x6e, 0xe2, 0x00, 0x19, 0x29, 0xd2, 0x1a, 0x54,
	0x5b, 0x1d, 0xc3, 0x6a, 0xe3, 0x26, 0x49, 0xa5, 0x94, 0xa4, 0x08, 0xc7, 0xc8, 0x91, 0x26, 0x65,
	0x6c, 0x24, 0xc6, 0x4e, 0xe8, 0x0d, 0x5c, 0x0b, 0xb4, 0x96, 0x1e, 0xb2, 0x98, 0x75, 0xf5, 0xe3,
	0xfa, 0x2f, 0x28, 0xf6, 0xfc, 0xc5, 0x34, 0x18, 0x7a, 0xd3, 0x68, 0xc3, 0x5a, 0x0d, 0x47, 0x43,
	0xa2, 0xea, 0xd9, 0x8e, 0x25, 0x8d, 0xd4, 0xe4, 0xb7, 0xb4, 0x9a, 0x55, 0x3c, 0xc2, 0x31, 0xb2,
	0xad, 0x07, 0x14, 0x3b, 0x5a, 0xd1, 0x44, 0x6a, 0xe1, 0x67, 0x70, 0x92, 0x4c, 0xc4, 0xee, 0x73,
	0xf5, 0xdf, 0x39, 0x00, 0x11, 0xf8, 0x33, 0x37, 0x6a, 0x7f, 0x05, 0x35, 0x83, 0x31, 0x25, 0xa4,
	0xcd, 0xf1, 0x8e, 0xdb, 0x96, 0x54, 0xa6, 0xc4, 0x5e, 0x4a, 0x42, 0x15, 0xca, 0xd9, 0xb4, 0xc1,
	0x58, 0xa8, 0xe3, 0x13, 0x5c, 0x32, 0xec, 0xa2, 0xc4, 0x77, 0x0a, 0xec, 0xd0, 0x0f, 0x70, 0x96,
	0x4d, 0x47, 0x2c, 0x7d, 0x7e, 0x37, 0x70, 0xbd, 0xb2, 0x92, 0x41, 0x25, 0x1d, 0xed, 0xd2, 0x8f,
	0x70, 0xbe, 0x21, 0x1f, 0x1b, 0xdb, 0xd3, 0x5a, 0x9c, 0x3e, 0x33, 0xde, 0xd5, 0xb2, 0xbf, 0x49,
	0x4b, 0xc4, 0x5a, 0xdd, 0xea, 0x52, 0x4b, 0x56, 0xc0, 0x81, 0xbe, 0xb9, 0x6c, 0xd7, 0xbc, 0xae,
	0x97, 0xa2, 0xa4, 0xae, 0x94, 0x42, 0x69, 0x19, 0x8b, 0x09, 0x05, 0x7a, 0x0e, 0x15, 0x81, 0x06,
	0x6f, 0x75, 0xfe, 0xe9, 0x70, 0x18, 0xe3, 0x85, 0x8a, 0x10, 0xfa, 0x5a, 0xeb, 0xf7, 0x50, 0xe0,
	0xee, 0x34, 0x18, 0xbe, 0x46, 0xfb, 0xaa, 0x00, 0x6d, 0x75, 0xb0, 0x75, 0x6f, 0x3b, 0x32, 0xb5,
	0xa4, 0x12, 0x40, 0x18, 0x47, 0x66, 0x3b, 0xfa, 0x1f, 0x38, 0x05, 0x92, 0xc0, 0x0d, 0x4c, 0x7c,
	0x60, 0x64, 0xe7, 0xeb, 0xed, 0xe3, 0xe7, 0x1f, 0x5e, 0x30, 0x5e, 0x3c, 0x35, 0x9e, 0xfd, 0x49,
	0xd3, 0x7b, 0x1d, 0x0f, 0x27, 0x93, 0xf1, 0x68, 0xd4, 0x7c, 0x59, 0x8c, 0x86, 0x2f, 0xde, 0xad,
	0x3b, 0x0d, 0xbc, 0xc0, 0x73, 0xe7, 0xcd, 0xe8, 0x97, 0x7a, 0xda, 0x0f, 0xff, 0xa7, 0x2f, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x35, 0x6b, 0xf1, 0x8c, 0xb6, 0x04, 0x00, 0x00,
}