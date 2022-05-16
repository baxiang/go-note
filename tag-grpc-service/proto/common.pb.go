// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/common.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Pager struct {
	Page                 int64    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Total                int64    `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pager) Reset()         { *m = Pager{} }
func (m *Pager) String() string { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()    {}
func (*Pager) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{0}
}

func (m *Pager) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pager.Unmarshal(m, b)
}
func (m *Pager) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pager.Marshal(b, m, deterministic)
}
func (m *Pager) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pager.Merge(m, src)
}
func (m *Pager) XXX_Size() int {
	return xxx_messageInfo_Pager.Size(m)
}
func (m *Pager) XXX_DiscardUnknown() {
	xxx_messageInfo_Pager.DiscardUnknown(m)
}

var xxx_messageInfo_Pager proto.InternalMessageInfo

func (m *Pager) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pager) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Pager) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data                 *any.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_1747d3070a2311a0, []int{1}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Error) GetData() *any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Pager)(nil), "proto.Pager")
	proto.RegisterType((*Error)(nil), "proto.Error")
}

func init() {
	proto.RegisterFile("proto/common.proto", fileDescriptor_1747d3070a2311a0)
}

var fileDescriptor_1747d3070a2311a0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0xaa, 0x83, 0x30,
	0x10, 0x45, 0xf1, 0x69, 0xde, 0xe3, 0x4d, 0x77, 0x83, 0x8b, 0xb4, 0xab, 0xe2, 0xca, 0x55, 0x84,
	0xf6, 0x0b, 0xba, 0x70, 0x5f, 0xf2, 0x03, 0x25, 0x6a, 0x1a, 0x0a, 0xea, 0x48, 0x4c, 0x17, 0xf6,
	0xeb, 0x4b, 0x26, 0xb8, 0x9a, 0x7b, 0x0f, 0xc3, 0xe1, 0x02, 0x2e, 0x9e, 0x02, 0x35, 0x3d, 0x4d,
	0x13, 0xcd, 0x8a, 0x0b, 0x0a, 0x3e, 0xa7, 0xa3, 0x23, 0x72, 0xa3, 0x6d, 0xb8, 0x75, 0xef, 0x67,
	0x63, 0xe6, 0x2d, 0x7d, 0x54, 0x2d, 0x88, 0xbb, 0x71, 0xd6, 0x23, 0x42, 0xb1, 0x18, 0x67, 0x65,
	0x76, 0xce, 0xea, 0x5c, 0x73, 0x8e, 0x6c, 0x7d, 0x7d, 0xac, 0xfc, 0x49, 0x2c, 0x66, 0x2c, 0x41,
	0x04, 0x0a, 0x66, 0x94, 0x39, 0xc3, 0x54, 0xaa, 0x07, 0x88, 0xd6, 0x7b, 0x62, 0x4d, 0x4f, 0x43,
	0xd2, 0x08, 0xcd, 0x19, 0x25, 0xfc, 0x4d, 0x76, 0x5d, 0xa3, 0x3d, 0x9a, 0xfe, 0xf5, 0x5e, 0xb1,
	0x86, 0x62, 0x30, 0xc1, 0xb0, 0xeb, 0x70, 0x29, 0x55, 0xda, 0xa9, 0xf6, 0x9d, 0xea, 0x36, 0x6f,
	0x9a, 0x3f, 0xba, 0x5f, 0x66, 0xd7, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0xd7, 0xb6, 0x4b,
	0xe6, 0x00, 0x00, 0x00,
}