// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: pb/auth.proto

package pb

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider  string `protobuf:"bytes,1,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Cluster   string `protobuf:"bytes,3,opt,name=Cluster,proto3" json:"Cluster,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=Namespace,proto3" json:"Namespace,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[0]
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
	return file_pb_auth_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *LoginRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type LoginStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthURL     string `protobuf:"bytes,1,opt,name=AuthURL,proto3" json:"AuthURL,omitempty"`
	OneTimeCode string `protobuf:"bytes,2,opt,name=OneTimeCode,proto3" json:"OneTimeCode,omitempty"`
	SecretYAML  string `protobuf:"bytes,3,opt,name=SecretYAML,proto3" json:"SecretYAML,omitempty"`
}

func (x *LoginStatus) Reset() {
	*x = LoginStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginStatus) ProtoMessage() {}

func (x *LoginStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginStatus.ProtoReflect.Descriptor instead.
func (*LoginStatus) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{1}
}

func (x *LoginStatus) GetAuthURL() string {
	if x != nil {
		return x.AuthURL
	}
	return ""
}

func (x *LoginStatus) GetOneTimeCode() string {
	if x != nil {
		return x.OneTimeCode
	}
	return ""
}

func (x *LoginStatus) GetSecretYAML() string {
	if x != nil {
		return x.SecretYAML
	}
	return ""
}

type RefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken      string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken     string `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	JWTSigningSecret string `protobuf:"bytes,3,opt,name=JWTSigningSecret,proto3" json:"JWTSigningSecret,omitempty"`
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *RefreshRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshRequest) GetJWTSigningSecret() string {
	if x != nil {
		return x.JWTSigningSecret
	}
	return ""
}

type RefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *RefreshResponse) Reset() {
	*x = RefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshResponse) ProtoMessage() {}

func (x *RefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshResponse.ProtoReflect.Descriptor instead.
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type GenerateAdminTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminName         string `protobuf:"bytes,1,opt,name=AdminName,proto3" json:"AdminName,omitempty"`
	JWTSigningSecret  string `protobuf:"bytes,2,opt,name=JWTSigningSecret,proto3" json:"JWTSigningSecret,omitempty"`
	RefreshExpiration int64  `protobuf:"varint,3,opt,name=RefreshExpiration,proto3" json:"RefreshExpiration,omitempty"`
	AccessExpiration  int64  `protobuf:"varint,4,opt,name=AccessExpiration,proto3" json:"AccessExpiration,omitempty"`
}

func (x *GenerateAdminTokenRequest) Reset() {
	*x = GenerateAdminTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAdminTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAdminTokenRequest) ProtoMessage() {}

func (x *GenerateAdminTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAdminTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateAdminTokenRequest) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateAdminTokenRequest) GetAdminName() string {
	if x != nil {
		return x.AdminName
	}
	return ""
}

func (x *GenerateAdminTokenRequest) GetJWTSigningSecret() string {
	if x != nil {
		return x.JWTSigningSecret
	}
	return ""
}

func (x *GenerateAdminTokenRequest) GetRefreshExpiration() int64 {
	if x != nil {
		return x.RefreshExpiration
	}
	return 0
}

func (x *GenerateAdminTokenRequest) GetAccessExpiration() int64 {
	if x != nil {
		return x.AccessExpiration
	}
	return 0
}

type GenerateAdminTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *GenerateAdminTokenResponse) Reset() {
	*x = GenerateAdminTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAdminTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAdminTokenResponse) ProtoMessage() {}

func (x *GenerateAdminTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAdminTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateAdminTokenResponse) Descriptor() ([]byte, []int) {
	return file_pb_auth_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateAdminTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_pb_auth_proto protoreflect.FileDescriptor

var file_pb_auth_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x61, 0x74, 0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x22, 0x7e, 0x0a, 0x0c, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a,
	0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x69, 0x0a, 0x0b, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x75,
	0x74, 0x68, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x75, 0x74,
	0x68, 0x55, 0x52, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x6e, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x59, 0x41, 0x4d, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x59, 0x41, 0x4d, 0x4c, 0x22, 0x82, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x4a, 0x57, 0x54, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4a, 0x57, 0x54, 0x53, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x33, 0x0a, 0x0f, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0xbf, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x4a, 0x57, 0x54, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4a, 0x57, 0x54, 0x53, 0x69, 0x67, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x11, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x45, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x10, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x1a, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x93, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x07, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x6c, 0x6c, 0x2f,
	0x6b, 0x61, 0x72, 0x61, 0x76, 0x69, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_auth_proto_rawDescOnce sync.Once
	file_pb_auth_proto_rawDescData = file_pb_auth_proto_rawDesc
)

func file_pb_auth_proto_rawDescGZIP() []byte {
	file_pb_auth_proto_rawDescOnce.Do(func() {
		file_pb_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_auth_proto_rawDescData)
	})
	return file_pb_auth_proto_rawDescData
}

var file_pb_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_auth_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),               // 0: gatekeeper.LoginRequest
	(*LoginStatus)(nil),                // 1: gatekeeper.LoginStatus
	(*RefreshRequest)(nil),             // 2: gatekeeper.RefreshRequest
	(*RefreshResponse)(nil),            // 3: gatekeeper.RefreshResponse
	(*GenerateAdminTokenRequest)(nil),  // 4: gatekeeper.GenerateAdminTokenRequest
	(*GenerateAdminTokenResponse)(nil), // 5: gatekeeper.GenerateAdminTokenResponse
}
var file_pb_auth_proto_depIdxs = []int32{
	0, // 0: gatekeeper.AuthService.Login:input_type -> gatekeeper.LoginRequest
	2, // 1: gatekeeper.AuthService.Refresh:input_type -> gatekeeper.RefreshRequest
	1, // 2: gatekeeper.AuthService.Login:output_type -> gatekeeper.LoginStatus
	3, // 3: gatekeeper.AuthService.Refresh:output_type -> gatekeeper.RefreshResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_auth_proto_init() }
func file_pb_auth_proto_init() {
	if File_pb_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_pb_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginStatus); i {
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
		file_pb_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRequest); i {
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
		file_pb_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshResponse); i {
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
		file_pb_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAdminTokenRequest); i {
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
		file_pb_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateAdminTokenResponse); i {
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
			RawDescriptor: file_pb_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_auth_proto_goTypes,
		DependencyIndexes: file_pb_auth_proto_depIdxs,
		MessageInfos:      file_pb_auth_proto_msgTypes,
	}.Build()
	File_pb_auth_proto = out.File
	file_pb_auth_proto_rawDesc = nil
	file_pb_auth_proto_goTypes = nil
	file_pb_auth_proto_depIdxs = nil
}
