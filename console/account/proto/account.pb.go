// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account.proto

package proto

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

// 登录请求体
type LoginRequest struct {
	//@inject_tag: json:"username" validate:"required" message:"required:用户名不能为空"
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username" validate:"required" message:"required:用户名不能为空"`
	//@inject_tag: json:"password" validate:"required" message:"required:密码不能为空"
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password" validate:"required" message:"required:密码不能为空"`
	//@inject_tag: json:"source" validate:"required" message:"required:来源不能为空"
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source" validate:"required" message:"required:来源不能为空"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// 登录响应体
type LoginResponse struct {
	//@inject_tag: json:"token"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
	//@inject_tag: json:"expires"
	Expires              string   `protobuf:"bytes,2,opt,name=expires,proto3" json:"expires"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetExpires() string {
	if m != nil {
		return m.Expires
	}
	return ""
}

// 注册请求体
type RegisterRequest struct {
	//@inject_tag: validate:"required" json:"nickname"
	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname" validate:"required"`
	//@inject_tag: validate:"required,gte=11" message:"required:手机号不能为空,gt:手机号不正确" json:"phone"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone" validate:"required,gte=11" message:"required:手机号不能为空,gt:手机号不正确"`
	//@inject_tag: validate:"required" json:"password"
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password" validate:"required"`
	//@inject_tag: validate:"required" json:"code"
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterRequest) Reset()         { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()    {}
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{2}
}

func (m *RegisterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterRequest.Unmarshal(m, b)
}
func (m *RegisterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterRequest.Marshal(b, m, deterministic)
}
func (m *RegisterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterRequest.Merge(m, src)
}
func (m *RegisterRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterRequest.Size(m)
}
func (m *RegisterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterRequest proto.InternalMessageInfo

func (m *RegisterRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *RegisterRequest) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RegisterRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

// 注册响应
type RegisterResponse struct {
	//@inject_tag: json:"id"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResponse) Reset()         { *m = RegisterResponse{} }
func (m *RegisterResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterResponse) ProtoMessage()    {}
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{3}
}

func (m *RegisterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResponse.Unmarshal(m, b)
}
func (m *RegisterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResponse.Marshal(b, m, deterministic)
}
func (m *RegisterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResponse.Merge(m, src)
}
func (m *RegisterResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterResponse.Size(m)
}
func (m *RegisterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResponse proto.InternalMessageInfo

func (m *RegisterResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// 账号信息请求体
type InfoRequest struct {
	//@inject_tag: validate:"required" json:"id"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id" validate:"required"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoRequest) Reset()         { *m = InfoRequest{} }
func (m *InfoRequest) String() string { return proto.CompactTextString(m) }
func (*InfoRequest) ProtoMessage()    {}
func (*InfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{4}
}

func (m *InfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoRequest.Unmarshal(m, b)
}
func (m *InfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoRequest.Marshal(b, m, deterministic)
}
func (m *InfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoRequest.Merge(m, src)
}
func (m *InfoRequest) XXX_Size() int {
	return xxx_messageInfo_InfoRequest.Size(m)
}
func (m *InfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InfoRequest proto.InternalMessageInfo

func (m *InfoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// 账号信息结构体
type InfoResponse struct {
	//@inject_tag: json:"nickname"
	Nickname string `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname"`
	//@inject_tag: json:"phone"
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone"`
	//@inject_tag: json:"email"
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email"`
	//@inject_tag: json:"register_time"
	RegisterTime         string   `protobuf:"bytes,4,opt,name=register_time,json=registerTime,proto3" json:"register_time"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfoResponse) Reset()         { *m = InfoResponse{} }
func (m *InfoResponse) String() string { return proto.CompactTextString(m) }
func (*InfoResponse) ProtoMessage()    {}
func (*InfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{5}
}

func (m *InfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfoResponse.Unmarshal(m, b)
}
func (m *InfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfoResponse.Marshal(b, m, deterministic)
}
func (m *InfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfoResponse.Merge(m, src)
}
func (m *InfoResponse) XXX_Size() int {
	return xxx_messageInfo_InfoResponse.Size(m)
}
func (m *InfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InfoResponse proto.InternalMessageInfo

func (m *InfoResponse) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *InfoResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *InfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *InfoResponse) GetRegisterTime() string {
	if m != nil {
		return m.RegisterTime
	}
	return ""
}

// 修改账号信息请求体
type UpdateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

// 修改账号信息返回体
type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e28828dcb8d24f0, []int{7}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LoginRequest)(nil), "com.dashenwo.srv.account.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "com.dashenwo.srv.account.LoginResponse")
	proto.RegisterType((*RegisterRequest)(nil), "com.dashenwo.srv.account.RegisterRequest")
	proto.RegisterType((*RegisterResponse)(nil), "com.dashenwo.srv.account.RegisterResponse")
	proto.RegisterType((*InfoRequest)(nil), "com.dashenwo.srv.account.InfoRequest")
	proto.RegisterType((*InfoResponse)(nil), "com.dashenwo.srv.account.InfoResponse")
	proto.RegisterType((*UpdateRequest)(nil), "com.dashenwo.srv.account.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "com.dashenwo.srv.account.UpdateResponse")
}

func init() {
	proto.RegisterFile("account.proto", fileDescriptor_8e28828dcb8d24f0)
}

var fileDescriptor_8e28828dcb8d24f0 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0x59, 0xf7, 0xfb, 0xfd, 0xae, 0xdb, 0x08, 0xe3, 0x4b, 0x29, 0x08, 0x12, 0x71, 0x4e,
	0x0f, 0x3d, 0xe8, 0xd1, 0x83, 0xe8, 0x4d, 0xf0, 0x34, 0x1c, 0x88, 0x82, 0x52, 0xdb, 0xd7, 0x2d,
	0xcc, 0x26, 0x35, 0x69, 0x9d, 0xe0, 0x7f, 0xee, 0x49, 0xd6, 0x24, 0xda, 0x0d, 0xe6, 0xf4, 0xb4,
	0x3d, 0x79, 0xde, 0x3c, 0xf9, 0xe4, 0xcd, 0x5b, 0x70, 0xc3, 0x28, 0x12, 0x39, 0xcf, 0x82, 0x54,
	0x8a, 0x4c, 0x10, 0x2f, 0x12, 0x49, 0x10, 0x87, 0x6a, 0x86, 0x7c, 0x21, 0x02, 0x25, 0x5f, 0x03,
	0xe3, 0xd3, 0x7b, 0xe8, 0x5c, 0x89, 0x29, 0xe3, 0x63, 0x7c, 0xc9, 0x51, 0x65, 0xc4, 0x87, 0x56,
	0xae, 0x50, 0xf2, 0x30, 0x41, 0xaf, 0xb2, 0x5b, 0x19, 0xb5, 0xc7, 0x5f, 0x7a, 0xe9, 0xa5, 0xa1,
	0x52, 0x0b, 0x21, 0x63, 0xcf, 0xd1, 0x9e, 0xd5, 0xe4, 0x3f, 0x34, 0x94, 0xc8, 0x65, 0x84, 0x5e,
	0xb5, 0x70, 0x8c, 0xa2, 0x67, 0xe0, 0x9a, 0x7c, 0x95, 0x0a, 0xae, 0x90, 0x0c, 0xa0, 0x9e, 0x89,
	0x39, 0x72, 0x93, 0xae, 0x05, 0xf1, 0xa0, 0x89, 0x6f, 0x29, 0x93, 0xa8, 0x4c, 0xb2, 0x95, 0x54,
	0x41, 0x6f, 0x8c, 0x53, 0xa6, 0x32, 0x94, 0x25, 0x46, 0xce, 0xa2, 0x79, 0x99, 0xd1, 0xea, 0x65,
	0x7c, 0x3a, 0x13, 0x1c, 0x4d, 0x8c, 0x16, 0x2b, 0xe4, 0xd5, 0x35, 0x72, 0x02, 0xb5, 0x48, 0xc4,
	0xe8, 0xd5, 0x8a, 0xf5, 0xe2, 0x3f, 0xa5, 0xd0, 0xff, 0x3e, 0xd4, 0x80, 0x77, 0xc1, 0x61, 0xb1,
	0x39, 0xcf, 0x61, 0x31, 0xdd, 0x81, 0x7f, 0x97, 0xfc, 0x49, 0x58, 0xa8, 0x75, 0xfb, 0x1d, 0x3a,
	0xda, 0x36, 0xdb, 0xff, 0x0e, 0x3d, 0x80, 0x3a, 0x26, 0x21, 0x7b, 0x36, 0xc4, 0x5a, 0x90, 0x3d,
	0x70, 0xa5, 0x41, 0x7b, 0xc8, 0x58, 0x62, 0xb9, 0x3b, 0x76, 0xf1, 0x9a, 0x25, 0x48, 0x7b, 0xe0,
	0x4e, 0xd2, 0x38, 0xcc, 0xd0, 0xd0, 0xd1, 0x3e, 0x74, 0xed, 0x82, 0xe6, 0x39, 0xfe, 0x70, 0xa0,
	0x79, 0xae, 0x87, 0x80, 0xdc, 0x40, 0xbd, 0x78, 0x24, 0x32, 0x0c, 0x36, 0x0d, 0x4a, 0x50, 0x9e,
	0x12, 0xff, 0x60, 0x6b, 0x9d, 0xb9, 0xf5, 0x04, 0x6a, 0xcb, 0x2e, 0x90, 0xfd, 0xcd, 0x1b, 0x4a,
	0x4d, 0xf4, 0x87, 0xdb, 0xca, 0x4c, 0x6c, 0x08, 0x2d, 0xfb, 0x3e, 0xe4, 0x70, 0xf3, 0x9e, 0xb5,
	0xc1, 0xf1, 0x8f, 0x7e, 0x53, 0x6a, 0x8e, 0xb8, 0x83, 0x86, 0xee, 0x18, 0xf9, 0xe1, 0xb2, 0x2b,
	0x4d, 0xf6, 0x47, 0xdb, 0x0b, 0x75, 0xf8, 0x45, 0xfb, 0xb6, 0x19, 0x9c, 0x16, 0x9f, 0xe6, 0x63,
	0xa3, 0xf8, 0x39, 0xf9, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x13, 0xc0, 0x7b, 0xcd, 0xb2, 0x03, 0x00,
	0x00,
}
