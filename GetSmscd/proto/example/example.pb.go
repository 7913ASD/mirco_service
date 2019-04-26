// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetSmscd

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Text                 string   `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Request) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Response struct {
	Errno                string   `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSmscd.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSmscd.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0xce, 0x82, 0x30,
	0x14, 0x85, 0x7f, 0x7e, 0x15, 0xf0, 0x8e, 0x8d, 0x1a, 0xa2, 0x89, 0x1a, 0x26, 0xa7, 0x9a, 0xe8,
	0xe2, 0x0b, 0x18, 0xe3, 0xe0, 0x82, 0x8b, 0xab, 0xc0, 0x0d, 0x21, 0xa1, 0x14, 0x6f, 0x8b, 0xe1,
	0xf1, 0x0d, 0x6d, 0xd9, 0x74, 0xea, 0x39, 0x5f, 0x93, 0xd3, 0xaf, 0xb0, 0x6a, 0x48, 0x6a, 0xb9,
	0xc7, 0xee, 0x29, 0x9a, 0x0a, 0x87, 0x93, 0x1b, 0xca, 0xe6, 0x85, 0xe4, 0xa2, 0xcc, 0x48, 0x72,
	0x45, 0x6f, 0x7e, 0x41, 0x7d, 0x17, 0x2a, 0xcb, 0xe3, 0x2b, 0x04, 0x09, 0xbe, 0x5a, 0x54, 0x9a,
	0x2d, 0xc0, 0x17, 0x32, 0x2d, 0x2b, 0x8c, 0xbc, 0xad, 0xb7, 0x9b, 0x26, 0xae, 0x31, 0x06, 0xe3,
	0xb6, 0x2d, 0xf3, 0xe8, 0xdf, 0x50, 0x93, 0x7b, 0xa6, 0xb1, 0xd3, 0xd1, 0xc8, 0xb2, 0x3e, 0xc7,
	0x27, 0x08, 0x13, 0x54, 0x8d, 0xac, 0x15, 0xb2, 0x19, 0x4c, 0x90, 0xa8, 0x96, 0x6e, 0xca, 0x96,
	0xfe, 0x05, 0x24, 0x12, 0xaa, 0x70, 0x5b, 0xae, 0x1d, 0x1e, 0x10, 0x9c, 0xad, 0x2c, 0xbb, 0x41,
	0x38, 0xb8, 0xb1, 0x35, 0xff, 0xea, 0xcc, 0x9d, 0xf0, 0x72, 0xf3, 0xf3, 0xde, 0x5a, 0xc4, 0x7f,
	0xa9, 0x6f, 0x3e, 0x7f, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xae, 0x8a, 0x09, 0x32, 0x1b, 0x01,
	0x00, 0x00,
}
