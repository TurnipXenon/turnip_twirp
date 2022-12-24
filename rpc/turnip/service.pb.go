// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc/turnip/service.proto

package turnip

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// region models
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[0]
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
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Username    string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{1}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Token) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Token) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type AccessDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllowedDomains []string `protobuf:"bytes,1,rep,name=allowed_domains,json=allowedDomains,proto3" json:"allowed_domains,omitempty"`
}

func (x *AccessDetails) Reset() {
	*x = AccessDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessDetails) ProtoMessage() {}

func (x *AccessDetails) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessDetails.ProtoReflect.Descriptor instead.
func (*AccessDetails) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{2}
}

func (x *AccessDetails) GetAllowedDomains() []string {
	if x != nil {
		return x.AllowedDomains
	}
	return nil
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Media         *string                `protobuf:"bytes,4,opt,name=media,proto3,oneof" json:"media,omitempty"`
	TagList       []string               `protobuf:"bytes,5,rep,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	AccessDetails *AccessDetails         `protobuf:"bytes,6,opt,name=access_details,json=accessDetails,proto3" json:"access_details,omitempty"`
	Meta          map[string]string      `protobuf:"bytes,7,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PrimaryId     string                 `protobuf:"bytes,8,opt,name=primary_id,json=primaryId,proto3" json:"primary_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{3}
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Content) GetMedia() string {
	if x != nil && x.Media != nil {
		return *x.Media
	}
	return ""
}

func (x *Content) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *Content) GetAccessDetails() *AccessDetails {
	if x != nil {
		return x.AccessDetails
	}
	return nil
}

func (x *Content) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Content) GetPrimaryId() string {
	if x != nil {
		return x.PrimaryId
	}
	return ""
}

func (x *Content) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// region request and responses
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"` // using msg to follow twirp's format
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{6}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token *Token `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	User  *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{7}
}

func (x *LoginResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *LoginResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type CreateContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string            `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Content       string            `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Media         *string           `protobuf:"bytes,4,opt,name=media,proto3,oneof" json:"media,omitempty"`
	TagList       []string          `protobuf:"bytes,5,rep,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	AccessDetails *AccessDetails    `protobuf:"bytes,6,opt,name=access_details,json=accessDetails,proto3" json:"access_details,omitempty"`
	Meta          map[string]string `protobuf:"bytes,7,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CreateContentRequest) Reset() {
	*x = CreateContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentRequest) ProtoMessage() {}

func (x *CreateContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentRequest.ProtoReflect.Descriptor instead.
func (*CreateContentRequest) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateContentRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateContentRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateContentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateContentRequest) GetMedia() string {
	if x != nil && x.Media != nil {
		return *x.Media
	}
	return ""
}

func (x *CreateContentRequest) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *CreateContentRequest) GetAccessDetails() *AccessDetails {
	if x != nil {
		return x.AccessDetails
	}
	return nil
}

func (x *CreateContentRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type CreateContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Media         *string                `protobuf:"bytes,4,opt,name=media,proto3,oneof" json:"media,omitempty"`
	TagList       []string               `protobuf:"bytes,5,rep,name=tag_list,json=tagList,proto3" json:"tag_list,omitempty"`
	AccessDetails *AccessDetails         `protobuf:"bytes,6,opt,name=access_details,json=accessDetails,proto3" json:"access_details,omitempty"`
	Meta          map[string]string      `protobuf:"bytes,7,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PrimaryId     string                 `protobuf:"bytes,8,opt,name=primary_id,json=primaryId,proto3" json:"primary_id,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateContentResponse) Reset() {
	*x = CreateContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_turnip_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateContentResponse) ProtoMessage() {}

func (x *CreateContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_turnip_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateContentResponse.ProtoReflect.Descriptor instead.
func (*CreateContentResponse) Descriptor() ([]byte, []int) {
	return file_rpc_turnip_service_proto_rawDescGZIP(), []int{9}
}

func (x *CreateContentResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateContentResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateContentResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateContentResponse) GetMedia() string {
	if x != nil && x.Media != nil {
		return *x.Media
	}
	return ""
}

func (x *CreateContentResponse) GetTagList() []string {
	if x != nil {
		return x.TagList
	}
	return nil
}

func (x *CreateContentResponse) GetAccessDetails() *AccessDetails {
	if x != nil {
		return x.AccessDetails
	}
	return nil
}

func (x *CreateContentResponse) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CreateContentResponse) GetPrimaryId() string {
	if x != nil {
		return x.PrimaryId
	}
	return ""
}

func (x *CreateContentResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_rpc_turnip_service_proto protoreflect.FileDescriptor

var file_rpc_turnip_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x2f, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x75, 0x72, 0x6e,
	0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xbc, 0x01, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x38,
	0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65,
	0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x22, 0xab, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0x4b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x46, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x22, 0x66, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x28, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xeb, 0x02, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x42, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x22, 0xc7, 0x03, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x61, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x0e, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x43, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x32, 0x81, 0x02, 0x0a, 0x06, 0x54, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x12, 0x53,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x74,
	0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x22, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1c, 0x2e, 0x74,
	0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x75, 0x72,
	0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x74, 0x75, 0x72,
	0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x74, 0x75, 0x72, 0x6e, 0x69, 0x70, 0x78, 0x65, 0x6e, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x72, 0x70, 0x63, 0x2f, 0x74,
	0x75, 0x72, 0x6e, 0x69, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_turnip_service_proto_rawDescOnce sync.Once
	file_rpc_turnip_service_proto_rawDescData = file_rpc_turnip_service_proto_rawDesc
)

func file_rpc_turnip_service_proto_rawDescGZIP() []byte {
	file_rpc_turnip_service_proto_rawDescOnce.Do(func() {
		file_rpc_turnip_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_turnip_service_proto_rawDescData)
	})
	return file_rpc_turnip_service_proto_rawDescData
}

var file_rpc_turnip_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_rpc_turnip_service_proto_goTypes = []interface{}{
	(*User)(nil),                  // 0: turnipxenon.v1.User
	(*Token)(nil),                 // 1: turnipxenon.v1.Token
	(*AccessDetails)(nil),         // 2: turnipxenon.v1.AccessDetails
	(*Content)(nil),               // 3: turnipxenon.v1.Content
	(*CreateUserRequest)(nil),     // 4: turnipxenon.v1.CreateUserRequest
	(*CreateUserResponse)(nil),    // 5: turnipxenon.v1.CreateUserResponse
	(*LoginRequest)(nil),          // 6: turnipxenon.v1.LoginRequest
	(*LoginResponse)(nil),         // 7: turnipxenon.v1.LoginResponse
	(*CreateContentRequest)(nil),  // 8: turnipxenon.v1.CreateContentRequest
	(*CreateContentResponse)(nil), // 9: turnipxenon.v1.CreateContentResponse
	nil,                           // 10: turnipxenon.v1.Content.MetaEntry
	nil,                           // 11: turnipxenon.v1.CreateContentRequest.MetaEntry
	nil,                           // 12: turnipxenon.v1.CreateContentResponse.MetaEntry
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_rpc_turnip_service_proto_depIdxs = []int32{
	13, // 0: turnipxenon.v1.Token.created_at:type_name -> google.protobuf.Timestamp
	13, // 1: turnipxenon.v1.Token.expires_at:type_name -> google.protobuf.Timestamp
	2,  // 2: turnipxenon.v1.Content.access_details:type_name -> turnipxenon.v1.AccessDetails
	10, // 3: turnipxenon.v1.Content.meta:type_name -> turnipxenon.v1.Content.MetaEntry
	13, // 4: turnipxenon.v1.Content.created_at:type_name -> google.protobuf.Timestamp
	1,  // 5: turnipxenon.v1.LoginResponse.token:type_name -> turnipxenon.v1.Token
	0,  // 6: turnipxenon.v1.LoginResponse.user:type_name -> turnipxenon.v1.User
	2,  // 7: turnipxenon.v1.CreateContentRequest.access_details:type_name -> turnipxenon.v1.AccessDetails
	11, // 8: turnipxenon.v1.CreateContentRequest.meta:type_name -> turnipxenon.v1.CreateContentRequest.MetaEntry
	2,  // 9: turnipxenon.v1.CreateContentResponse.access_details:type_name -> turnipxenon.v1.AccessDetails
	12, // 10: turnipxenon.v1.CreateContentResponse.meta:type_name -> turnipxenon.v1.CreateContentResponse.MetaEntry
	13, // 11: turnipxenon.v1.CreateContentResponse.created_at:type_name -> google.protobuf.Timestamp
	4,  // 12: turnipxenon.v1.Turnip.CreateUser:input_type -> turnipxenon.v1.CreateUserRequest
	6,  // 13: turnipxenon.v1.Turnip.Login:input_type -> turnipxenon.v1.LoginRequest
	8,  // 14: turnipxenon.v1.Turnip.CreateContent:input_type -> turnipxenon.v1.CreateContentRequest
	5,  // 15: turnipxenon.v1.Turnip.CreateUser:output_type -> turnipxenon.v1.CreateUserResponse
	7,  // 16: turnipxenon.v1.Turnip.Login:output_type -> turnipxenon.v1.LoginResponse
	9,  // 17: turnipxenon.v1.Turnip.CreateContent:output_type -> turnipxenon.v1.CreateContentResponse
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_rpc_turnip_service_proto_init() }
func file_rpc_turnip_service_proto_init() {
	if File_rpc_turnip_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_turnip_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpc_turnip_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_rpc_turnip_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessDetails); i {
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
		file_rpc_turnip_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_rpc_turnip_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_rpc_turnip_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_rpc_turnip_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_rpc_turnip_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
		file_rpc_turnip_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentRequest); i {
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
		file_rpc_turnip_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateContentResponse); i {
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
	file_rpc_turnip_service_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_rpc_turnip_service_proto_msgTypes[8].OneofWrappers = []interface{}{}
	file_rpc_turnip_service_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_turnip_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_turnip_service_proto_goTypes,
		DependencyIndexes: file_rpc_turnip_service_proto_depIdxs,
		MessageInfos:      file_rpc_turnip_service_proto_msgTypes,
	}.Build()
	File_rpc_turnip_service_proto = out.File
	file_rpc_turnip_service_proto_rawDesc = nil
	file_rpc_turnip_service_proto_goTypes = nil
	file_rpc_turnip_service_proto_depIdxs = nil
}