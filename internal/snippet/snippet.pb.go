// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/snippet/snippet.proto

package snippet

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

type Snippet struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snippet) Reset()         { *m = Snippet{} }
func (m *Snippet) String() string { return proto.CompactTextString(m) }
func (*Snippet) ProtoMessage()    {}
func (*Snippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0156653fa8d1c7d6, []int{0}
}

func (m *Snippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snippet.Unmarshal(m, b)
}
func (m *Snippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snippet.Marshal(b, m, deterministic)
}
func (m *Snippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snippet.Merge(m, src)
}
func (m *Snippet) XXX_Size() int {
	return xxx_messageInfo_Snippet.Size(m)
}
func (m *Snippet) XXX_DiscardUnknown() {
	xxx_messageInfo_Snippet.DiscardUnknown(m)
}

var xxx_messageInfo_Snippet proto.InternalMessageInfo

func (m *Snippet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Snippet) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Snippet) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Snippet) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Snippet)(nil), "snippet.Snippet")
}

func init() { proto.RegisterFile("internal/snippet/snippet.proto", fileDescriptor_0156653fa8d1c7d6) }

var fileDescriptor_0156653fa8d1c7d6 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2f, 0xce, 0xcb, 0x2c, 0x28, 0x48, 0x2d, 0x81, 0xd1, 0x7a, 0x05,
	0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x52, 0x21, 0x17, 0x7b, 0x30, 0x84, 0x29, 0xc4,
	0xc7, 0xc5, 0x94, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x94, 0x99, 0x22, 0x24,
	0xc5, 0xc5, 0x91, 0x96, 0x99, 0x93, 0x9a, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x04, 0x16, 0x85, 0xf3,
	0x85, 0x14, 0xb8, 0xb8, 0x53, 0x52, 0x8b, 0x93, 0x8b, 0x32, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x24,
	0x98, 0xc1, 0xd2, 0xc8, 0x42, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0xf9, 0x79, 0x25, 0xa9, 0x79, 0x25,
	0x12, 0x2c, 0x60, 0x59, 0x18, 0x37, 0x89, 0x0d, 0xec, 0x04, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4f, 0xe4, 0xbf, 0x3f, 0xa4, 0x00, 0x00, 0x00,
}