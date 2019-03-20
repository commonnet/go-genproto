// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/shopping_performance_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Request message for
// [ShoppingPerformanceViewService.GetShoppingPerformanceView][google.ads.googleads.v1.services.ShoppingPerformanceViewService.GetShoppingPerformanceView].
type GetShoppingPerformanceViewRequest struct {
	// The resource name of the Shopping performance view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetShoppingPerformanceViewRequest) Reset()         { *m = GetShoppingPerformanceViewRequest{} }
func (m *GetShoppingPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetShoppingPerformanceViewRequest) ProtoMessage()    {}
func (*GetShoppingPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d6c4cf22050d6d3, []int{0}
}

func (m *GetShoppingPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.Merge(m, src)
}
func (m *GetShoppingPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetShoppingPerformanceViewRequest.Size(m)
}
func (m *GetShoppingPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetShoppingPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetShoppingPerformanceViewRequest proto.InternalMessageInfo

func (m *GetShoppingPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetShoppingPerformanceViewRequest)(nil), "google.ads.googleads.v1.services.GetShoppingPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/shopping_performance_view_service.proto", fileDescriptor_4d6c4cf22050d6d3)
}

var fileDescriptor_4d6c4cf22050d6d3 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc6, 0x49, 0x5e, 0x78, 0xe1, 0x0d, 0xaf, 0x4b, 0x26, 0x09, 0x22, 0xb5, 0xed, 0x20, 0x1d,
	0xee, 0x88, 0x4e, 0x5e, 0x71, 0x48, 0x1d, 0xda, 0x49, 0x4a, 0x0b, 0x19, 0x24, 0x10, 0xce, 0xe4,
	0x8c, 0x81, 0xe6, 0x2e, 0xde, 0x3f, 0x4d, 0x07, 0x71, 0xd1, 0xcd, 0xd5, 0x6f, 0xe0, 0xe8, 0x47,
	0x71, 0xf5, 0x0b, 0x38, 0x38, 0xf9, 0x29, 0x24, 0xbd, 0x5c, 0x44, 0x21, 0xed, 0xf6, 0x90, 0x3c,
	0xf7, 0x7b, 0xee, 0xff, 0xdc, 0xdf, 0x9a, 0x24, 0x42, 0x24, 0x0b, 0x86, 0x69, 0x0c, 0x58, 0xc9,
	0x4a, 0x95, 0x2e, 0x06, 0x26, 0xcb, 0x34, 0x62, 0x80, 0xe1, 0x5a, 0xe4, 0x79, 0xca, 0x93, 0x30,
	0x67, 0xf2, 0x4a, 0xc8, 0x8c, 0xf2, 0x88, 0x85, 0x65, 0xca, 0x56, 0x61, 0x6d, 0x41, 0xb9, 0x14,
	0x85, 0xb0, 0x3b, 0xea, 0x38, 0xa2, 0x31, 0xa0, 0x86, 0x84, 0x4a, 0x17, 0x69, 0x92, 0xe3, 0xb5,
	0x65, 0x49, 0x06, 0x62, 0x29, 0x37, 0x86, 0xa9, 0x10, 0x67, 0x4f, 0x23, 0xf2, 0x14, 0x53, 0xce,
	0x45, 0x41, 0x8b, 0x54, 0x70, 0x50, 0x7f, 0xbb, 0x13, 0xeb, 0x60, 0xcc, 0x8a, 0x79, 0xcd, 0x98,
	0x7e, 0x23, 0xfc, 0x94, 0xad, 0x66, 0xec, 0x66, 0xc9, 0xa0, 0xb0, 0x7b, 0xd6, 0x8e, 0xce, 0x0b,
	0x39, 0xcd, 0xd8, 0xae, 0xd1, 0x31, 0x0e, 0xff, 0xcd, 0xfe, 0xeb, 0x8f, 0xe7, 0x34, 0x63, 0x47,
	0x0f, 0xa6, 0xb5, 0xdf, 0xc2, 0x99, 0xab, 0x71, 0xec, 0x77, 0xc3, 0x72, 0xda, 0xd3, 0xec, 0x33,
	0xb4, 0xad, 0x0f, 0xb4, 0xf5, 0xae, 0x0e, 0x69, 0x85, 0x34, 0x95, 0xa1, 0x16, 0x44, 0xd7, 0xbb,
	0x7f, 0xfb, 0x78, 0x32, 0x87, 0xf6, 0x49, 0xd5, 0xf0, 0xed, 0x8f, 0x91, 0x4f, 0xa3, 0x25, 0x14,
	0x22, 0x63, 0x12, 0xf0, 0xa0, 0xa9, 0xfc, 0xd7, 0x79, 0x3c, 0xb8, 0x1b, 0x3d, 0x9a, 0x56, 0x3f,
	0x12, 0xd9, 0xd6, 0x49, 0x46, 0xbd, 0xcd, 0x5d, 0x4d, 0xab, 0xd7, 0x99, 0x1a, 0x17, 0xf5, 0xb2,
	0xa1, 0x44, 0x2c, 0x28, 0x4f, 0x90, 0x90, 0x09, 0x4e, 0x18, 0x5f, 0xbf, 0x9d, 0x5e, 0x88, 0x3c,
	0x85, 0xf6, 0x5d, 0x1c, 0x6a, 0xf1, 0x6c, 0xfe, 0x19, 0x7b, 0xde, 0x8b, 0xd9, 0x19, 0x2b, 0xa0,
	0x17, 0x03, 0x52, 0xb2, 0x52, 0xbe, 0x8b, 0xea, 0x60, 0x78, 0xd5, 0x96, 0xc0, 0x8b, 0x21, 0x68,
	0x2c, 0x81, 0xef, 0x06, 0xda, 0xf2, 0x69, 0xf6, 0xd5, 0x77, 0x42, 0xbc, 0x18, 0x08, 0x69, 0x4c,
	0x84, 0xf8, 0x2e, 0x21, 0xda, 0x76, 0xf9, 0x77, 0x7d, 0xcf, 0xe3, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x0b, 0x90, 0x3f, 0x35, 0x32, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShoppingPerformanceViewServiceClient is the client API for ShoppingPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShoppingPerformanceViewServiceClient interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error)
}

type shoppingPerformanceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewShoppingPerformanceViewServiceClient(cc *grpc.ClientConn) ShoppingPerformanceViewServiceClient {
	return &shoppingPerformanceViewServiceClient{cc}
}

func (c *shoppingPerformanceViewServiceClient) GetShoppingPerformanceView(ctx context.Context, in *GetShoppingPerformanceViewRequest, opts ...grpc.CallOption) (*resources.ShoppingPerformanceView, error) {
	out := new(resources.ShoppingPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShoppingPerformanceViewServiceServer is the server API for ShoppingPerformanceViewService service.
type ShoppingPerformanceViewServiceServer interface {
	// Returns the requested Shopping performance view in full detail.
	GetShoppingPerformanceView(context.Context, *GetShoppingPerformanceViewRequest) (*resources.ShoppingPerformanceView, error)
}

func RegisterShoppingPerformanceViewServiceServer(s *grpc.Server, srv ShoppingPerformanceViewServiceServer) {
	s.RegisterService(&_ShoppingPerformanceViewService_serviceDesc, srv)
}

func _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShoppingPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.ShoppingPerformanceViewService/GetShoppingPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShoppingPerformanceViewServiceServer).GetShoppingPerformanceView(ctx, req.(*GetShoppingPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ShoppingPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.ShoppingPerformanceViewService",
	HandlerType: (*ShoppingPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShoppingPerformanceView",
			Handler:    _ShoppingPerformanceViewService_GetShoppingPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/shopping_performance_view_service.proto",
}
