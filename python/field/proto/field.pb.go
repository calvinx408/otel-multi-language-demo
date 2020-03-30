// Code generated by protoc-gen-go. DO NOT EDIT.
// source: field.proto

package field

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type FieldRequest struct {
	Slow                 bool     `protobuf:"varint,1,opt,name=slow,proto3" json:"slow,omitempty"`
	Unreliable           bool     `protobuf:"varint,2,opt,name=unreliable,proto3" json:"unreliable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldRequest) Reset()         { *m = FieldRequest{} }
func (m *FieldRequest) String() string { return proto.CompactTextString(m) }
func (*FieldRequest) ProtoMessage()    {}
func (*FieldRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_04234ff7fdd53e6e, []int{0}
}

func (m *FieldRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldRequest.Unmarshal(m, b)
}
func (m *FieldRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldRequest.Marshal(b, m, deterministic)
}
func (m *FieldRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldRequest.Merge(m, src)
}
func (m *FieldRequest) XXX_Size() int {
	return xxx_messageInfo_FieldRequest.Size(m)
}
func (m *FieldRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FieldRequest proto.InternalMessageInfo

func (m *FieldRequest) GetSlow() bool {
	if m != nil {
		return m.Slow
	}
	return false
}

func (m *FieldRequest) GetUnreliable() bool {
	if m != nil {
		return m.Unreliable
	}
	return false
}

type FieldReply struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldReply) Reset()         { *m = FieldReply{} }
func (m *FieldReply) String() string { return proto.CompactTextString(m) }
func (*FieldReply) ProtoMessage()    {}
func (*FieldReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_04234ff7fdd53e6e, []int{1}
}

func (m *FieldReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldReply.Unmarshal(m, b)
}
func (m *FieldReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldReply.Marshal(b, m, deterministic)
}
func (m *FieldReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldReply.Merge(m, src)
}
func (m *FieldReply) XXX_Size() int {
	return xxx_messageInfo_FieldReply.Size(m)
}
func (m *FieldReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldReply.DiscardUnknown(m)
}

var xxx_messageInfo_FieldReply proto.InternalMessageInfo

func (m *FieldReply) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func init() {
	proto.RegisterType((*FieldRequest)(nil), "field.FieldRequest")
	proto.RegisterType((*FieldReply)(nil), "field.FieldReply")
}

func init() { proto.RegisterFile("field.proto", fileDescriptor_04234ff7fdd53e6e) }

var fileDescriptor_04234ff7fdd53e6e = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcb, 0x4c, 0xcd,
	0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x73, 0x94, 0x9c, 0xb8, 0x78, 0xdc,
	0x40, 0x8c, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xe2, 0x9c, 0xfc,
	0x72, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x8e, 0x20, 0x30, 0x5b, 0x48, 0x8e, 0x8b, 0xab, 0x34, 0xaf,
	0x28, 0x35, 0x27, 0x33, 0x31, 0x29, 0x27, 0x55, 0x82, 0x09, 0x2c, 0x83, 0x24, 0xa2, 0xa4, 0xc4,
	0xc5, 0x05, 0x35, 0xa3, 0x20, 0xa7, 0x52, 0x48, 0x84, 0x0b, 0x62, 0x34, 0xd8, 0x08, 0xce, 0x20,
	0x08, 0xc7, 0xc8, 0x96, 0x8b, 0x15, 0xac, 0x46, 0xc8, 0x84, 0x8b, 0xc3, 0x3d, 0xb5, 0x04, 0xc2,
	0x16, 0xd6, 0x83, 0xb8, 0x08, 0xd9, 0x05, 0x52, 0x82, 0xa8, 0x82, 0x05, 0x39, 0x95, 0x4a, 0x0c,
	0x49, 0x6c, 0x60, 0x47, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa1, 0x5d, 0x4f, 0xba, 0xc3,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FieldClient is the client API for Field service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FieldClient interface {
	GetField(ctx context.Context, in *FieldRequest, opts ...grpc.CallOption) (*FieldReply, error)
}

type fieldClient struct {
	cc *grpc.ClientConn
}

func NewFieldClient(cc *grpc.ClientConn) FieldClient {
	return &fieldClient{cc}
}

func (c *fieldClient) GetField(ctx context.Context, in *FieldRequest, opts ...grpc.CallOption) (*FieldReply, error) {
	out := new(FieldReply)
	err := c.cc.Invoke(ctx, "/field.Field/GetField", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FieldServer is the server API for Field service.
type FieldServer interface {
	GetField(context.Context, *FieldRequest) (*FieldReply, error)
}

// UnimplementedFieldServer can be embedded to have forward compatible implementations.
type UnimplementedFieldServer struct {
}

func (*UnimplementedFieldServer) GetField(ctx context.Context, req *FieldRequest) (*FieldReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetField not implemented")
}

func RegisterFieldServer(s *grpc.Server, srv FieldServer) {
	s.RegisterService(&_Field_serviceDesc, srv)
}

func _Field_GetField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FieldServer).GetField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/field.Field/GetField",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FieldServer).GetField(ctx, req.(*FieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Field_serviceDesc = grpc.ServiceDesc{
	ServiceName: "field.Field",
	HandlerType: (*FieldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetField",
			Handler:    _Field_GetField_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "field.proto",
}
