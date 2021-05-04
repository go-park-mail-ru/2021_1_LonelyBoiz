// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: user.proto

package user_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Email          string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password       string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	SecondPassword string   `protobuf:"bytes,4,opt,name=SecondPassword,proto3" json:"SecondPassword,omitempty"`
	PasswordHash   []byte   `protobuf:"bytes,5,opt,name=PasswordHash,proto3" json:"PasswordHash,omitempty"`
	OldPassword    string   `protobuf:"bytes,6,opt,name=OldPassword,proto3" json:"OldPassword,omitempty"`
	Name           string   `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`
	Birthday       int64    `protobuf:"varint,8,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	Description    string   `protobuf:"bytes,9,opt,name=Description,proto3" json:"Description,omitempty"`
	City           string   `protobuf:"bytes,10,opt,name=City,proto3" json:"City,omitempty"`
	Instagram      string   `protobuf:"bytes,11,opt,name=Instagram,proto3" json:"Instagram,omitempty"`
	Sex            string   `protobuf:"bytes,12,opt,name=Sex,proto3" json:"Sex,omitempty"`
	DatePreference string   `protobuf:"bytes,13,opt,name=DatePreference,proto3" json:"DatePreference,omitempty"`
	IsDeleted      bool     `protobuf:"varint,14,opt,name=IsDeleted,proto3" json:"IsDeleted,omitempty"`
	IsActive       bool     `protobuf:"varint,15,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	Photos         [][]byte `protobuf:"bytes,16,rep,name=Photos,proto3" json:"Photos,omitempty"`
	CaptchaToken   string   `protobuf:"bytes,17,opt,name=CaptchaToken,proto3" json:"CaptchaToken,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetSecondPassword() string {
	if x != nil {
		return x.SecondPassword
	}
	return ""
}

func (x *User) GetPasswordHash() []byte {
	if x != nil {
		return x.PasswordHash
	}
	return nil
}

func (x *User) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetBirthday() int64 {
	if x != nil {
		return x.Birthday
	}
	return 0
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *User) GetInstagram() string {
	if x != nil {
		return x.Instagram
	}
	return ""
}

func (x *User) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *User) GetDatePreference() string {
	if x != nil {
		return x.DatePreference
	}
	return ""
}

func (x *User) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *User) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *User) GetPhotos() [][]byte {
	if x != nil {
		return x.Photos
	}
	return nil
}

func (x *User) GetCaptchaToken() string {
	if x != nil {
		return x.CaptchaToken
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User  *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *UserResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email          string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
	Password       string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	SecondPassword string `protobuf:"bytes,3,opt,name=SecondPassword,proto3" json:"SecondPassword,omitempty"`
}

func (x *UserLogin) Reset() {
	*x = UserLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogin) ProtoMessage() {}

func (x *UserLogin) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogin.ProtoReflect.Descriptor instead.
func (*UserLogin) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserLogin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserLogin) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserLogin) GetSecondPassword() string {
	if x != nil {
		return x.SecondPassword
	}
	return ""
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserId) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserNothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *UserNothing) Reset() {
	*x = UserNothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserNothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserNothing) ProtoMessage() {}

func (x *UserNothing) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserNothing.ProtoReflect.Descriptor instead.
func (*UserNothing) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *UserNothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

type Feed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserId `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *Feed) Reset() {
	*x = Feed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feed) ProtoMessage() {}

func (x *Feed) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feed.ProtoReflect.Descriptor instead.
func (*Feed) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{5}
}

func (x *Feed) GetUsers() []*UserId {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xea, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b,
	0x4f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x20,
	0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72, 0x61,
	0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x67, 0x72,
	0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x65, 0x78, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x53, 0x65, 0x78, 0x12, 0x26, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x65, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x44, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x49, 0x73,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x18, 0x10, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x11,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x47, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x65, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x22, 0x23, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d,
	0x79, 0x22, 0x2d, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x05, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x32, 0xca, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d,
	0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x15, 0x2e,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x46, 0x65, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x22, 0x00, 0x42, 0x66, 0x5a,
	0x64, 0x2f, 0x55, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x2f, 0x32, 0x5f, 0x70, 0x61, 0x72, 0x6b,
	0x4d, 0x61, 0x69, 0x6c, 0x2f, 0x32, 0x30, 0x32, 0x31, 0x5f, 0x31, 0x5f, 0x4c, 0x6f, 0x6e, 0x65,
	0x6c, 0x79, 0x42, 0x6f, 0x69, 0x7a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_proto_goTypes = []interface{}{
	(*User)(nil),         // 0: session.User
	(*UserResponse)(nil), // 1: session.UserResponse
	(*UserLogin)(nil),    // 2: session.UserLogin
	(*UserId)(nil),       // 3: session.UserId
	(*UserNothing)(nil),  // 4: session.UserNothing
	(*Feed)(nil),         // 5: session.Feed
}
var file_user_proto_depIdxs = []int32{
	0, // 0: session.UserResponse.user:type_name -> session.User
	3, // 1: session.Feed.users:type_name -> session.UserId
	0, // 2: session.UserService.CreateUser:input_type -> session.User
	0, // 3: session.UserService.ChangeUser:input_type -> session.User
	2, // 4: session.UserService.CheckUser:input_type -> session.UserLogin
	4, // 5: session.UserService.DeleteUser:input_type -> session.UserNothing
	4, // 6: session.UserService.GetUserById:input_type -> session.UserNothing
	4, // 7: session.UserService.CreateFeed:input_type -> session.UserNothing
	1, // 8: session.UserService.CreateUser:output_type -> session.UserResponse
	0, // 9: session.UserService.ChangeUser:output_type -> session.User
	0, // 10: session.UserService.CheckUser:output_type -> session.User
	4, // 11: session.UserService.DeleteUser:output_type -> session.UserNothing
	0, // 12: session.UserService.GetUserById:output_type -> session.User
	5, // 13: session.UserService.CreateFeed:output_type -> session.Feed
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLogin); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserNothing); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feed); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
