// Code generated by protoc-gen-go. DO NOT EDIT.
// source: file.proto

package file_proto

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

type File struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Buf                  []byte   `protobuf:"bytes,2,opt,name=buf,proto3" json:"buf,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *File) GetBuf() []byte {
	if m != nil {
		return m.Buf
	}
	return nil
}

func init() {
	proto.RegisterType((*File)(nil), "file.proto.File")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 85 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcb, 0xcc, 0x49,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x42, 0x62, 0x2b, 0xe9, 0x71, 0xb1, 0xb8, 0x65, 0xe6,
	0xa4, 0x0a, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30,
	0x07, 0x41, 0x38, 0x42, 0x02, 0x5c, 0xcc, 0x49, 0xa5, 0x69, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x20, 0x66, 0x12, 0x1b, 0x58, 0x9b, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x01, 0x7e, 0xfc,
	0xa0, 0x50, 0x00, 0x00, 0x00,
}
