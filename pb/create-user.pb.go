// Code generated by protoc-gen-go. DO NOT EDIT.
// source: create-user.proto

package pb

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

type UserChat struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FullName             string   `protobuf:"bytes,2,opt,name=fullName,proto3" json:"fullName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	PasswordChangedAt    string   `protobuf:"bytes,4,opt,name=password_changed_at,json=passwordChangedAt,proto3" json:"password_changed_at,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserChat) Reset()         { *m = UserChat{} }
func (m *UserChat) String() string { return proto.CompactTextString(m) }
func (*UserChat) ProtoMessage()    {}
func (*UserChat) Descriptor() ([]byte, []int) {
	return fileDescriptor_91d3c652ead4f23b, []int{0}
}

func (m *UserChat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserChat.Unmarshal(m, b)
}
func (m *UserChat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserChat.Marshal(b, m, deterministic)
}
func (m *UserChat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserChat.Merge(m, src)
}
func (m *UserChat) XXX_Size() int {
	return xxx_messageInfo_UserChat.Size(m)
}
func (m *UserChat) XXX_DiscardUnknown() {
	xxx_messageInfo_UserChat.DiscardUnknown(m)
}

var xxx_messageInfo_UserChat proto.InternalMessageInfo

func (m *UserChat) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserChat) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *UserChat) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserChat) GetPasswordChangedAt() string {
	if m != nil {
		return m.PasswordChangedAt
	}
	return ""
}

func (m *UserChat) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func init() {
	proto.RegisterType((*UserChat)(nil), "pb.UserChat")
}

func init() { proto.RegisterFile("create-user.proto", fileDescriptor_91d3c652ead4f23b) }

var fileDescriptor_91d3c652ead4f23b = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcf, 0xbf, 0x6e, 0x83, 0x30,
	0x10, 0xc7, 0x71, 0x41, 0x4b, 0x05, 0xde, 0x70, 0x3b, 0xa0, 0x4a, 0x95, 0xfa, 0x67, 0xe9, 0x52,
	0x18, 0xfa, 0x04, 0x94, 0xbd, 0x43, 0xa5, 0x2e, 0x59, 0xd0, 0xd9, 0x38, 0x18, 0x09, 0x73, 0x96,
	0x7d, 0x28, 0x8f, 0x94, 0xd7, 0x8c, 0xb0, 0x43, 0xc6, 0x9f, 0x3f, 0xdf, 0xc1, 0xc7, 0x4a, 0xe9,
	0x14, 0x90, 0xfa, 0x5a, 0xbd, 0x72, 0xb5, 0x75, 0x48, 0xc8, 0x53, 0x2b, 0xde, 0xcf, 0x09, 0xcb,
	0xff, 0xbd, 0x72, 0x9d, 0x06, 0xe2, 0xcf, 0x2c, 0xdf, 0x78, 0x01, 0xa3, 0xaa, 0xe4, 0x35, 0xf9,
	0x2c, 0xfe, 0x6e, 0x7b, 0xb3, 0xe3, 0x3a, 0xcf, 0xbf, 0x9b, 0xa5, 0xd1, 0xf6, 0xcd, 0x9f, 0x58,
	0xa6, 0x0c, 0x4c, 0x73, 0x75, 0x17, 0x20, 0x0e, 0x5e, 0xb3, 0x47, 0x0b, 0xde, 0x9f, 0xd0, 0x0d,
	0xbd, 0xd4, 0xb0, 0x8c, 0x6a, 0xe8, 0x81, 0xaa, 0xfb, 0xd0, 0x94, 0x3b, 0x75, 0x51, 0x5a, 0xe2,
	0x2f, 0x8c, 0xc5, 0x3f, 0x86, 0x2c, 0x0b, 0x59, 0x71, 0x7d, 0x69, 0xe9, 0xe7, 0xe3, 0xf0, 0x36,
	0x4e, 0xa4, 0x57, 0x51, 0x4b, 0x34, 0x0d, 0x0a, 0xf0, 0x88, 0x84, 0xa6, 0x19, 0x51, 0x6a, 0x20,
	0xb0, 0xb6, 0xb1, 0x42, 0x3c, 0x84, 0xcb, 0xbe, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xee, 0x63,
	0xf4, 0x43, 0xee, 0x00, 0x00, 0x00,
}
