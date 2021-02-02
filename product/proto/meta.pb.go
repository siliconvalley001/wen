// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/meta.proto

package product

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

type InsertSonMetaRequest struct {
	Id                   int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Authname             string                  `protobuf:"bytes,2,opt,name=authname,proto3" json:"authname,omitempty"`
	Path                 string                  `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Children             []*InsertSonMetaRequest `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InsertSonMetaRequest) Reset()         { *m = InsertSonMetaRequest{} }
func (m *InsertSonMetaRequest) String() string { return proto.CompactTextString(m) }
func (*InsertSonMetaRequest) ProtoMessage()    {}
func (*InsertSonMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{0}
}

func (m *InsertSonMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertSonMetaRequest.Unmarshal(m, b)
}
func (m *InsertSonMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertSonMetaRequest.Marshal(b, m, deterministic)
}
func (m *InsertSonMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertSonMetaRequest.Merge(m, src)
}
func (m *InsertSonMetaRequest) XXX_Size() int {
	return xxx_messageInfo_InsertSonMetaRequest.Size(m)
}
func (m *InsertSonMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertSonMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertSonMetaRequest proto.InternalMessageInfo

func (m *InsertSonMetaRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertSonMetaRequest) GetAuthname() string {
	if m != nil {
		return m.Authname
	}
	return ""
}

func (m *InsertSonMetaRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *InsertSonMetaRequest) GetChildren() []*InsertSonMetaRequest {
	if m != nil {
		return m.Children
	}
	return nil
}

type InsertMetaRequest struct {
	Id                   int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Authname             string                  `protobuf:"bytes,2,opt,name=authname,proto3" json:"authname,omitempty"`
	Path                 string                  `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Children             []*InsertSonMetaRequest `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *InsertMetaRequest) Reset()         { *m = InsertMetaRequest{} }
func (m *InsertMetaRequest) String() string { return proto.CompactTextString(m) }
func (*InsertMetaRequest) ProtoMessage()    {}
func (*InsertMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{1}
}

func (m *InsertMetaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertMetaRequest.Unmarshal(m, b)
}
func (m *InsertMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertMetaRequest.Marshal(b, m, deterministic)
}
func (m *InsertMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertMetaRequest.Merge(m, src)
}
func (m *InsertMetaRequest) XXX_Size() int {
	return xxx_messageInfo_InsertMetaRequest.Size(m)
}
func (m *InsertMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertMetaRequest proto.InternalMessageInfo

func (m *InsertMetaRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *InsertMetaRequest) GetAuthname() string {
	if m != nil {
		return m.Authname
	}
	return ""
}

func (m *InsertMetaRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *InsertMetaRequest) GetChildren() []*InsertSonMetaRequest {
	if m != nil {
		return m.Children
	}
	return nil
}

type MetaInfo struct {
	Id                   int64                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Authname             string                  `protobuf:"bytes,2,opt,name=authname,proto3" json:"authname,omitempty"`
	Path                 string                  `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Children             []*InsertSonMetaRequest `protobuf:"bytes,4,rep,name=children,proto3" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MetaInfo) Reset()         { *m = MetaInfo{} }
func (m *MetaInfo) String() string { return proto.CompactTextString(m) }
func (*MetaInfo) ProtoMessage()    {}
func (*MetaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{2}
}

func (m *MetaInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaInfo.Unmarshal(m, b)
}
func (m *MetaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaInfo.Marshal(b, m, deterministic)
}
func (m *MetaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaInfo.Merge(m, src)
}
func (m *MetaInfo) XXX_Size() int {
	return xxx_messageInfo_MetaInfo.Size(m)
}
func (m *MetaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MetaInfo proto.InternalMessageInfo

func (m *MetaInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MetaInfo) GetAuthname() string {
	if m != nil {
		return m.Authname
	}
	return ""
}

func (m *MetaInfo) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *MetaInfo) GetChildren() []*InsertSonMetaRequest {
	if m != nil {
		return m.Children
	}
	return nil
}

type MetaByIdRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaByIdRequest) Reset()         { *m = MetaByIdRequest{} }
func (m *MetaByIdRequest) String() string { return proto.CompactTextString(m) }
func (*MetaByIdRequest) ProtoMessage()    {}
func (*MetaByIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{3}
}

func (m *MetaByIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaByIdRequest.Unmarshal(m, b)
}
func (m *MetaByIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaByIdRequest.Marshal(b, m, deterministic)
}
func (m *MetaByIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaByIdRequest.Merge(m, src)
}
func (m *MetaByIdRequest) XXX_Size() int {
	return xxx_messageInfo_MetaByIdRequest.Size(m)
}
func (m *MetaByIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaByIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetaByIdRequest proto.InternalMessageInfo

func (m *MetaByIdRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MetaByNameRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaByNameRequest) Reset()         { *m = MetaByNameRequest{} }
func (m *MetaByNameRequest) String() string { return proto.CompactTextString(m) }
func (*MetaByNameRequest) ProtoMessage()    {}
func (*MetaByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{4}
}

func (m *MetaByNameRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaByNameRequest.Unmarshal(m, b)
}
func (m *MetaByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaByNameRequest.Marshal(b, m, deterministic)
}
func (m *MetaByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaByNameRequest.Merge(m, src)
}
func (m *MetaByNameRequest) XXX_Size() int {
	return xxx_messageInfo_MetaByNameRequest.Size(m)
}
func (m *MetaByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetaByNameRequest proto.InternalMessageInfo

func (m *MetaByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MetaResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Status               int64    `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaResponse) Reset()         { *m = MetaResponse{} }
func (m *MetaResponse) String() string { return proto.CompactTextString(m) }
func (*MetaResponse) ProtoMessage()    {}
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{5}
}

func (m *MetaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaResponse.Unmarshal(m, b)
}
func (m *MetaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaResponse.Marshal(b, m, deterministic)
}
func (m *MetaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaResponse.Merge(m, src)
}
func (m *MetaResponse) XXX_Size() int {
	return xxx_messageInfo_MetaResponse.Size(m)
}
func (m *MetaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetaResponse proto.InternalMessageInfo

func (m *MetaResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MetaResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type MetaAllRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetaAllRequest) Reset()         { *m = MetaAllRequest{} }
func (m *MetaAllRequest) String() string { return proto.CompactTextString(m) }
func (*MetaAllRequest) ProtoMessage()    {}
func (*MetaAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{6}
}

func (m *MetaAllRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaAllRequest.Unmarshal(m, b)
}
func (m *MetaAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaAllRequest.Marshal(b, m, deterministic)
}
func (m *MetaAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaAllRequest.Merge(m, src)
}
func (m *MetaAllRequest) XXX_Size() int {
	return xxx_messageInfo_MetaAllRequest.Size(m)
}
func (m *MetaAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetaAllRequest proto.InternalMessageInfo

type MetaAllResponse struct {
	Info                 []*MetaInfo `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	Msg                  string      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Status               int64       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MetaAllResponse) Reset()         { *m = MetaAllResponse{} }
func (m *MetaAllResponse) String() string { return proto.CompactTextString(m) }
func (*MetaAllResponse) ProtoMessage()    {}
func (*MetaAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bcdfb81738caa4f, []int{7}
}

func (m *MetaAllResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetaAllResponse.Unmarshal(m, b)
}
func (m *MetaAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetaAllResponse.Marshal(b, m, deterministic)
}
func (m *MetaAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaAllResponse.Merge(m, src)
}
func (m *MetaAllResponse) XXX_Size() int {
	return xxx_messageInfo_MetaAllResponse.Size(m)
}
func (m *MetaAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetaAllResponse proto.InternalMessageInfo

func (m *MetaAllResponse) GetInfo() []*MetaInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *MetaAllResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *MetaAllResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterType((*InsertSonMetaRequest)(nil), "product.InsertSonMetaRequest")
	proto.RegisterType((*InsertMetaRequest)(nil), "product.InsertMetaRequest")
	proto.RegisterType((*MetaInfo)(nil), "product.MetaInfo")
	proto.RegisterType((*MetaByIdRequest)(nil), "product.MetaByIdRequest")
	proto.RegisterType((*MetaByNameRequest)(nil), "product.MetaByNameRequest")
	proto.RegisterType((*MetaResponse)(nil), "product.MetaResponse")
	proto.RegisterType((*MetaAllRequest)(nil), "product.MetaAllRequest")
	proto.RegisterType((*MetaAllResponse)(nil), "product.MetaAllResponse")
}

func init() { proto.RegisterFile("proto/meta.proto", fileDescriptor_0bcdfb81738caa4f) }

var fileDescriptor_0bcdfb81738caa4f = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0xfd, 0x26, 0x09, 0xfd, 0xea, 0xad, 0xfd, 0xbb, 0xf8, 0x13, 0x02, 0x42, 0x0d, 0x88, 0x5d,
	0x55, 0xa8, 0x1b, 0xc5, 0x8d, 0x2d, 0x82, 0x74, 0xa1, 0x8b, 0x74, 0xe7, 0x2e, 0x6d, 0xa6, 0x36,
	0x90, 0x64, 0x62, 0x66, 0x22, 0xf8, 0x00, 0x6e, 0x05, 0xdf, 0xce, 0xc7, 0x91, 0x99, 0x4c, 0x62,
	0xab, 0xed, 0x5a, 0x77, 0xf7, 0xe6, 0x9c, 0x39, 0xe7, 0xdc, 0x03, 0x81, 0x4e, 0x9a, 0x31, 0xc1,
	0xce, 0x62, 0x2a, 0xfc, 0x81, 0x1a, 0xf1, 0x7f, 0x9a, 0xb1, 0x20, 0x9f, 0x0b, 0xf7, 0x9d, 0xc0,
	0xde, 0x24, 0xe1, 0x34, 0x13, 0x53, 0x96, 0xdc, 0x51, 0xe1, 0x7b, 0xf4, 0x29, 0xa7, 0x5c, 0x60,
	0x0b, 0x8c, 0x30, 0xb0, 0x49, 0x8f, 0xf4, 0x4d, 0xcf, 0x08, 0x03, 0x74, 0xa0, 0xee, 0xe7, 0x62,
	0x99, 0xf8, 0x31, 0xb5, 0x8d, 0x1e, 0xe9, 0xef, 0x78, 0xd5, 0x8e, 0x08, 0x56, 0xea, 0x8b, 0xa5,
	0x6d, 0xaa, 0xef, 0x6a, 0xc6, 0x4b, 0xa8, 0xcf, 0x97, 0x61, 0x14, 0x64, 0x34, 0xb1, 0xad, 0x9e,
	0xd9, 0x6f, 0x0c, 0x8f, 0x06, 0xda, 0x74, 0xb0, 0xc9, 0xd0, 0xab, 0xe8, 0xee, 0x1b, 0x81, 0x6e,
	0x41, 0xf9, 0x23, 0x81, 0x5e, 0x09, 0xd4, 0x25, 0x32, 0x49, 0x16, 0xec, 0x37, 0x73, 0x1c, 0x43,
	0x5b, 0x02, 0xe3, 0x97, 0x49, 0xb0, 0xa5, 0x15, 0xf7, 0x14, 0xba, 0x05, 0xe5, 0xde, 0x8f, 0x69,
	0x49, 0x42, 0xb0, 0x54, 0x3c, 0x52, 0xc4, 0x90, 0xb3, 0x7b, 0x01, 0xbb, 0x85, 0x09, 0x4f, 0x59,
	0xc2, 0x29, 0x76, 0xc0, 0x8c, 0xf9, 0xa3, 0xa6, 0xc8, 0x11, 0x0f, 0xa0, 0xc6, 0x85, 0x2f, 0x72,
	0xae, 0xce, 0x32, 0x3d, 0xbd, 0xb9, 0x1d, 0x68, 0xc9, 0x97, 0xa3, 0x28, 0xd2, 0xfa, 0xee, 0xac,
	0xc8, 0xa5, 0xbe, 0x68, 0xb9, 0x13, 0xb0, 0xc2, 0x64, 0xc1, 0x6c, 0xa2, 0x2e, 0xec, 0x56, 0x17,
	0x96, 0x35, 0x7a, 0x0a, 0x2e, 0x5d, 0x8d, 0x4d, 0xae, 0xe6, 0xaa, 0xeb, 0xf0, 0x83, 0x40, 0x43,
	0x3e, 0x9e, 0xd2, 0xec, 0x39, 0x9c, 0x53, 0xbc, 0x86, 0xc6, 0x2d, 0x15, 0x65, 0x1d, 0x68, 0xaf,
	0x39, 0xac, 0x34, 0xe4, 0xec, 0xaf, 0x21, 0x65, 0x40, 0xf7, 0x1f, 0xde, 0x40, 0xb3, 0x52, 0x90,
	0x6d, 0xa1, 0xf3, 0x4d, 0x63, 0xa5, 0xc2, 0xed, 0x2a, 0x23, 0x00, 0xad, 0x32, 0x8a, 0x22, 0x3c,
	0x5c, 0xa3, 0x7d, 0x55, 0xe4, 0xd8, 0x3f, 0x81, 0x52, 0x62, 0xdc, 0x7e, 0x68, 0xaa, 0xbf, 0xf2,
	0x4a, 0x53, 0x66, 0x35, 0xb5, 0x9e, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x61, 0xc1, 0x32, 0x00,
	0xb8, 0x03, 0x00, 0x00,
}
