// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetArea

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

//web->srv
type Request struct {
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

//srv->web
type Response struct {
	//错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//地域信息切片
	Data                 []*Address `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
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

func (m *Response) GetData() []*Address {
	if m != nil {
		return m.Data
	}
	return nil
}

//地域信息消息体
type Address struct {
	Aid                  int32    `protobuf:"varint,1,opt,name=aid,proto3" json:"aid,omitempty"`
	Aname                string   `protobuf:"bytes,2,opt,name=aname,proto3" json:"aname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{2}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetAid() int32 {
	if m != nil {
		return m.Aid
	}
	return 0
}

func (m *Address) GetAname() string {
	if m != nil {
		return m.Aname
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetArea.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetArea.Response")
	proto.RegisterType((*Address)(nil), "go.micro.srv.GetArea.Address")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xc7, 0x30,
	0x0c, 0xc6, 0xfd, 0x3b, 0xff, 0xab, 0x8b, 0x17, 0x29, 0x43, 0x86, 0xa2, 0x8c, 0x9e, 0x76, 0xaa,
	0x6c, 0x3e, 0xc1, 0x0e, 0x22, 0x78, 0x2c, 0xf8, 0x00, 0xd5, 0x86, 0x31, 0xb4, 0xeb, 0x4c, 0xab,
	0xf8, 0xf8, 0xb2, 0xb6, 0xde, 0xdc, 0x29, 0xf9, 0xbe, 0x84, 0xfc, 0xf8, 0x02, 0x37, 0x2b, 0xb9,
	0xe0, 0xee, 0xf1, 0x47, 0xdb, 0xf5, 0x03, 0xff, 0xaa, 0x8c, 0x2e, 0xaf, 0x27, 0x27, 0xed, 0xfc,
	0x46, 0x4e, 0x7a, 0xfa, 0x96, 0x4f, 0x18, 0x46, 0x42, 0x2d, 0x2a, 0x60, 0x0a, 0x3f, 0xbf, 0xd0,
	0x07, 0xf1, 0x0e, 0xe7, 0x0a, 0xfd, 0xea, 0x16, 0x8f, 0xbc, 0x86, 0x23, 0x12, 0x2d, 0xae, 0x39,
	0xb4, 0x87, 0xae, 0x52, 0x49, 0xf0, 0x2b, 0x28, 0x91, 0xc8, 0xfa, 0xa9, 0x39, 0x8d, 0x76, 0x56,
	0xbc, 0x87, 0x33, 0xa3, 0x83, 0x6e, 0x8a, 0xb6, 0xe8, 0x2e, 0x86, 0x5b, 0xf9, 0x1f, 0x49, 0x8e,
	0xc6, 0x10, 0x7a, 0xaf, 0xe2, 0xaa, 0xe8, 0x81, 0x65, 0x83, 0x5f, 0x42, 0xa1, 0x67, 0x13, 0x49,
	0x47, 0xb5, 0xb5, 0x1b, 0x5d, 0x2f, 0xda, 0x62, 0xc6, 0x24, 0x31, 0xbc, 0x00, 0x7b, 0x4c, 0x89,
	0xf8, 0x33, 0xb0, 0x7c, 0x96, 0xef, 0xd0, 0x72, 0xa8, 0xeb, 0xbb, 0xbd, 0x71, 0x0a, 0x2a, 0x4e,
	0x5e, 0xcb, 0xf8, 0x9e, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x20, 0xc7, 0x60, 0x3d,
	0x01, 0x00, 0x00,
}
