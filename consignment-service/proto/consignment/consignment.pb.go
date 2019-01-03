// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package consignment

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

//货轮承运的一批货物
type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type Response struct {
	Created              bool         `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
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

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "Consignment")
	proto.RegisterType((*Container)(nil), "Container")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0x24, 0xfd, 0xce, 0x33, 0x02, 0xe1, 0x01, 0x2c, 0x18, 0x88, 0xc2, 0x52, 0x31, 0x18, 0xa9,
	0xfc, 0x02, 0xd4, 0xa9, 0xab, 0xcb, 0x8e, 0x8a, 0xfd, 0x94, 0x3e, 0x89, 0xda, 0x91, 0xed, 0x96,
	0xbf, 0xc3, 0x4f, 0x45, 0x75, 0x1a, 0x30, 0x62, 0xcb, 0xdd, 0xe5, 0xde, 0x9d, 0xce, 0xf0, 0xd0,
	0x7a, 0x17, 0xdd, 0x93, 0x76, 0x36, 0x50, 0x63, 0x77, 0x68, 0x63, 0xfe, 0x2d, 0x93, 0x5a, 0x7f,
	0x15, 0xc0, 0x96, 0xbf, 0x2c, 0xbf, 0x80, 0x01, 0x19, 0x51, 0x54, 0xc5, 0xbc, 0x54, 0x03, 0x32,
	0xbc, 0x02, 0x66, 0x30, 0x68, 0x4f, 0x6d, 0x24, 0x67, 0xc5, 0x20, 0x09, 0x39, 0xc5, 0xaf, 0x61,
	0xf2, 0x89, 0xd4, 0x6c, 0xa3, 0x18, 0x56, 0xc5, 0x7c, 0xac, 0x4e, 0x88, 0x3f, 0x02, 0x68, 0x67,
	0xe3, 0x86, 0x2c, 0xfa, 0x20, 0x46, 0xd5, 0x70, 0xce, 0x16, 0x20, 0x97, 0x3d, 0xa5, 0x32, 0x95,
	0xdf, 0x41, 0x79, 0xc0, 0x10, 0xf0, 0xe3, 0x8d, 0x8c, 0x18, 0xa7, 0x8c, 0x59, 0x47, 0xac, 0x4c,
	0xbd, 0x83, 0xf2, 0xc7, 0xf5, 0xaf, 0xdf, 0x3d, 0x30, 0xbd, 0x0f, 0xd1, 0xed, 0xd0, 0x1f, 0xbd,
	0x5d, 0x3f, 0xe8, 0xa9, 0x95, 0x39, 0xd6, 0x73, 0x9e, 0x1a, 0xb2, 0xa9, 0x5e, 0xa9, 0x4e, 0x88,
	0xdf, 0xc0, 0x74, 0x1f, 0x3a, 0xd3, 0xa8, 0x13, 0x8e, 0x70, 0x65, 0xea, 0x57, 0x98, 0x29, 0x0c,
	0xad, 0xb3, 0x01, 0xb9, 0x80, 0xa9, 0xf6, 0xb8, 0x89, 0xd8, 0x45, 0xce, 0x54, 0x0f, 0xb9, 0x04,
	0x96, 0x8d, 0x99, 0x72, 0xd9, 0xe2, 0x5c, 0x66, 0x53, 0xaa, 0xfc, 0x87, 0xc5, 0x0b, 0x5c, 0xae,
	0xb7, 0xd4, 0xb6, 0x64, 0x9b, 0x35, 0xfa, 0x03, 0x69, 0xe4, 0x12, 0xae, 0x96, 0xe9, 0x5a, 0xbe,
	0xff, 0x9f, 0x13, 0xb7, 0xa5, 0xec, 0xab, 0xd4, 0x67, 0xef, 0x93, 0xf4, 0x62, 0xcf, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x17, 0x5a, 0x77, 0xb5, 0xd8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	cc *grpc.ClientConn
}

func NewShippingServiceClient(cc *grpc.ClientConn) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/ShippingService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
type ShippingServiceServer interface {
	CreateConsignment(context.Context, *Consignment) (*Response, error)
}

func RegisterShippingServiceServer(s *grpc.Server, srv ShippingServiceServer) {
	s.RegisterService(&_ShippingService_serviceDesc, srv)
}

func _ShippingService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Consignment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ShippingService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).CreateConsignment(ctx, req.(*Consignment))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShippingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ShippingService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/consignment/consignment.proto",
}