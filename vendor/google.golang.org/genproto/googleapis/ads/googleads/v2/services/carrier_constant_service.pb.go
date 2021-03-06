// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/services/carrier_constant_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v2/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [CarrierConstantService.GetCarrierConstant][google.ads.googleads.v2.services.CarrierConstantService.GetCarrierConstant].
type GetCarrierConstantRequest struct {
	// Resource name of the carrier constant to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCarrierConstantRequest) Reset()         { *m = GetCarrierConstantRequest{} }
func (m *GetCarrierConstantRequest) String() string { return proto.CompactTextString(m) }
func (*GetCarrierConstantRequest) ProtoMessage()    {}
func (*GetCarrierConstantRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b7cfe7382d5c66, []int{0}
}

func (m *GetCarrierConstantRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCarrierConstantRequest.Unmarshal(m, b)
}
func (m *GetCarrierConstantRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCarrierConstantRequest.Marshal(b, m, deterministic)
}
func (m *GetCarrierConstantRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCarrierConstantRequest.Merge(m, src)
}
func (m *GetCarrierConstantRequest) XXX_Size() int {
	return xxx_messageInfo_GetCarrierConstantRequest.Size(m)
}
func (m *GetCarrierConstantRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCarrierConstantRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCarrierConstantRequest proto.InternalMessageInfo

func (m *GetCarrierConstantRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCarrierConstantRequest)(nil), "google.ads.googleads.v2.services.GetCarrierConstantRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/services/carrier_constant_service.proto", fileDescriptor_08b7cfe7382d5c66)
}

var fileDescriptor_08b7cfe7382d5c66 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbf, 0x4a, 0xc3, 0x40,
	0x18, 0x27, 0x11, 0x04, 0x83, 0x2e, 0x19, 0xb4, 0xa6, 0x0e, 0xa5, 0x16, 0x29, 0x0e, 0x77, 0x10,
	0x17, 0xb9, 0x22, 0x9a, 0x76, 0xa8, 0x93, 0x94, 0x0a, 0x1d, 0x24, 0x50, 0xce, 0xe4, 0x08, 0x81,
	0xe6, 0xae, 0xde, 0x97, 0x76, 0x11, 0x97, 0xbe, 0x82, 0x6f, 0xe0, 0xe8, 0xee, 0x4b, 0x74, 0xf5,
	0x15, 0x9c, 0x04, 0xdf, 0x41, 0xd2, 0xcb, 0xa5, 0xb6, 0x36, 0x74, 0xfb, 0x71, 0xdf, 0xef, 0xcf,
	0xf7, 0xfd, 0x12, 0xeb, 0x3a, 0x12, 0x22, 0x1a, 0x31, 0x4c, 0x43, 0xc0, 0x0a, 0x66, 0x68, 0xea,
	0x62, 0x60, 0x72, 0x1a, 0x07, 0x0c, 0x70, 0x40, 0xa5, 0x8c, 0x99, 0x1c, 0x06, 0x82, 0x43, 0x4a,
	0x79, 0x3a, 0xcc, 0x27, 0x68, 0x2c, 0x45, 0x2a, 0xec, 0x9a, 0x52, 0x21, 0x1a, 0x02, 0x2a, 0x0c,
	0xd0, 0xd4, 0x45, 0xda, 0xc0, 0xb9, 0x2c, 0x8b, 0x90, 0x0c, 0xc4, 0x44, 0x6e, 0xca, 0x50, 0xde,
	0xce, 0x89, 0x56, 0x8e, 0x63, 0x4c, 0x39, 0x17, 0x29, 0x4d, 0x63, 0xc1, 0x21, 0x9f, 0x1e, 0xfd,
	0x99, 0x06, 0xa3, 0x98, 0x69, 0x59, 0xfd, 0xc6, 0x3a, 0xee, 0xb2, 0xb4, 0xa3, 0x3c, 0x3b, 0xb9,
	0x65, 0x9f, 0x3d, 0x4d, 0x18, 0xa4, 0xf6, 0xa9, 0x75, 0xa0, 0x73, 0x87, 0x9c, 0x26, 0xac, 0x62,
	0xd4, 0x8c, 0xe6, 0x5e, 0x7f, 0x5f, 0x3f, 0xde, 0xd1, 0x84, 0xb9, 0x3f, 0x86, 0x75, 0xb8, 0xa6,
	0xbf, 0x57, 0xe7, 0xd8, 0x1f, 0x86, 0x65, 0xff, 0x77, 0xb7, 0x5b, 0x68, 0x5b, 0x0f, 0xa8, 0x74,
	0x27, 0xc7, 0x2d, 0x15, 0x17, 0x15, 0xa1, 0x35, 0x69, 0x1d, 0xcd, 0x3e, 0xbf, 0x5e, 0xcd, 0xa6,
	0x7d, 0x96, 0x35, 0xf9, 0xbc, 0x72, 0xd2, 0x55, 0xb0, 0xca, 0x05, 0x7c, 0xfe, 0xe2, 0x54, 0xe7,
	0x5e, 0x65, 0x69, 0x9d, 0xa3, 0x71, 0x0c, 0x28, 0x10, 0x49, 0x7b, 0x66, 0x5a, 0x8d, 0x40, 0x24,
	0x5b, 0x6f, 0x68, 0x57, 0x37, 0xb7, 0xd2, 0xcb, 0x7a, 0xef, 0x19, 0x0f, 0xb7, 0xb9, 0x41, 0x24,
	0x46, 0x94, 0x47, 0x48, 0xc8, 0x08, 0x47, 0x8c, 0x2f, 0xbe, 0x0a, 0x5e, 0x46, 0x96, 0xff, 0x6c,
	0x2d, 0x0d, 0xde, 0xcc, 0x9d, 0xae, 0xe7, 0xbd, 0x9b, 0xb5, 0xae, 0x32, 0xf4, 0x42, 0x40, 0x0a,
	0x66, 0x68, 0xe0, 0xa2, 0x3c, 0x18, 0xe6, 0x9a, 0xe2, 0x7b, 0x21, 0xf8, 0x05, 0xc5, 0x1f, 0xb8,
	0xbe, 0xa6, 0x7c, 0x9b, 0x0d, 0xf5, 0x4e, 0x88, 0x17, 0x02, 0x21, 0x05, 0x89, 0x90, 0x81, 0x4b,
	0x88, 0xa6, 0x3d, 0xee, 0x2e, 0xf6, 0xbc, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x77, 0xb3,
	0x23, 0x13, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CarrierConstantServiceClient is the client API for CarrierConstantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CarrierConstantServiceClient interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error)
}

type carrierConstantServiceClient struct {
	cc *grpc.ClientConn
}

func NewCarrierConstantServiceClient(cc *grpc.ClientConn) CarrierConstantServiceClient {
	return &carrierConstantServiceClient{cc}
}

func (c *carrierConstantServiceClient) GetCarrierConstant(ctx context.Context, in *GetCarrierConstantRequest, opts ...grpc.CallOption) (*resources.CarrierConstant, error) {
	out := new(resources.CarrierConstant)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v2.services.CarrierConstantService/GetCarrierConstant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarrierConstantServiceServer is the server API for CarrierConstantService service.
type CarrierConstantServiceServer interface {
	// Returns the requested carrier constant in full detail.
	GetCarrierConstant(context.Context, *GetCarrierConstantRequest) (*resources.CarrierConstant, error)
}

// UnimplementedCarrierConstantServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCarrierConstantServiceServer struct {
}

func (*UnimplementedCarrierConstantServiceServer) GetCarrierConstant(ctx context.Context, req *GetCarrierConstantRequest) (*resources.CarrierConstant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCarrierConstant not implemented")
}

func RegisterCarrierConstantServiceServer(s *grpc.Server, srv CarrierConstantServiceServer) {
	s.RegisterService(&_CarrierConstantService_serviceDesc, srv)
}

func _CarrierConstantService_GetCarrierConstant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarrierConstantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v2.services.CarrierConstantService/GetCarrierConstant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarrierConstantServiceServer).GetCarrierConstant(ctx, req.(*GetCarrierConstantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CarrierConstantService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v2.services.CarrierConstantService",
	HandlerType: (*CarrierConstantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCarrierConstant",
			Handler:    _CarrierConstantService_GetCarrierConstant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v2/services/carrier_constant_service.proto",
}
