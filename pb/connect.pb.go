// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connect.proto

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

type Connect struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_778b7e3040344da6, []int{0}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Connect) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*Connect)(nil), "pb.Connect")
}

func init() { proto.RegisterFile("connect.proto", fileDescriptor_778b7e3040344da6) }

var fileDescriptor_778b7e3040344da6 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcb,
	0x4b, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0x2a,
	0x2d, 0x4e, 0x2d, 0x82, 0xf0, 0x95, 0xec, 0xb9, 0xd8, 0x9d, 0x21, 0x0a, 0x84, 0x64, 0xb8, 0x58,
	0x40, 0x12, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x1c, 0x7a, 0x05, 0x49, 0x7a, 0xa1, 0xc5,
	0xa9, 0x45, 0x41, 0x60, 0x51, 0x21, 0x31, 0x2e, 0xb6, 0xc4, 0xe4, 0x92, 0xcc, 0xb2, 0x54, 0x09,
	0x26, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x28, 0xcf, 0x49, 0x39, 0x4a, 0x31, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x3f, 0x29, 0xb1, 0x38, 0x3f, 0xbf, 0x24, 0x3f, 0x57,
	0x3f, 0x3d, 0x3f, 0x39, 0x23, 0xb1, 0x24, 0xb1, 0xa0, 0x40, 0xbf, 0x20, 0x29, 0x89, 0x0d, 0x6c,
	0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xdb, 0xd3, 0x2e, 0x8d, 0x00, 0x00, 0x00,
}