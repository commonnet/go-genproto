// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/ad_schedule_view_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	resources "google.golang.org/genproto/googleapis/ads/googleads/v0/resources"
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
// [AdScheduleViewService.GetAdScheduleView][google.ads.googleads.v0.services.AdScheduleViewService.GetAdScheduleView].
type GetAdScheduleViewRequest struct {
	// The resource name of the ad schedule view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdScheduleViewRequest) Reset()         { *m = GetAdScheduleViewRequest{} }
func (m *GetAdScheduleViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdScheduleViewRequest) ProtoMessage()    {}
func (*GetAdScheduleViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_43e521f19e21f9cb, []int{0}
}

func (m *GetAdScheduleViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdScheduleViewRequest.Unmarshal(m, b)
}
func (m *GetAdScheduleViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdScheduleViewRequest.Marshal(b, m, deterministic)
}
func (m *GetAdScheduleViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdScheduleViewRequest.Merge(m, src)
}
func (m *GetAdScheduleViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdScheduleViewRequest.Size(m)
}
func (m *GetAdScheduleViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdScheduleViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdScheduleViewRequest proto.InternalMessageInfo

func (m *GetAdScheduleViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdScheduleViewRequest)(nil), "google.ads.googleads.v0.services.GetAdScheduleViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/ad_schedule_view_service.proto", fileDescriptor_43e521f19e21f9cb)
}

var fileDescriptor_43e521f19e21f9cb = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x49, 0x2e, 0x5c, 0xb8, 0xe1, 0xba, 0x30, 0x20, 0x94, 0xe0, 0xa2, 0xd4, 0x2e, 0xa4,
	0x8b, 0x99, 0xd4, 0x6e, 0x74, 0x44, 0x4a, 0xba, 0xa9, 0x2b, 0x29, 0x2d, 0x64, 0x21, 0x81, 0x30,
	0x66, 0x86, 0x18, 0x68, 0x32, 0x35, 0x27, 0x49, 0x17, 0xe2, 0x42, 0x5f, 0xc1, 0x37, 0x70, 0xe9,
	0x3b, 0xf8, 0x02, 0x6e, 0x5d, 0xf8, 0x02, 0xae, 0x7c, 0x0a, 0x49, 0x27, 0x13, 0x08, 0x36, 0x74,
	0xf7, 0x31, 0xe7, 0xfb, 0x7d, 0xe7, 0x4f, 0x62, 0x8c, 0x43, 0x21, 0xc2, 0x25, 0xc7, 0x94, 0x01,
	0x96, 0xb2, 0x54, 0x85, 0x8d, 0x81, 0xa7, 0x45, 0x14, 0x70, 0xc0, 0x94, 0xf9, 0x10, 0xdc, 0x72,
	0x96, 0x2f, 0xb9, 0x5f, 0x44, 0x7c, 0xed, 0x57, 0x15, 0xb4, 0x4a, 0x45, 0x26, 0xcc, 0xae, 0xa4,
	0x10, 0x65, 0x80, 0xea, 0x00, 0x54, 0xd8, 0x48, 0x05, 0x58, 0xa7, 0x6d, 0x2d, 0x52, 0x0e, 0x22,
	0x4f, 0xb7, 0xf5, 0x90, 0xd9, 0xd6, 0xa1, 0x22, 0x57, 0x11, 0xa6, 0x49, 0x22, 0x32, 0x9a, 0x45,
	0x22, 0x01, 0x59, 0xed, 0x8d, 0x8d, 0xce, 0x94, 0x67, 0x0e, 0x5b, 0x54, 0xa4, 0x1b, 0xf1, 0xf5,
	0x9c, 0xdf, 0xe5, 0x1c, 0x32, 0xf3, 0xc8, 0xd8, 0x53, 0xe9, 0x7e, 0x42, 0x63, 0xde, 0xd1, 0xba,
	0xda, 0xf1, 0xbf, 0xf9, 0x7f, 0xf5, 0x78, 0x45, 0x63, 0x7e, 0xf2, 0xa9, 0x19, 0x07, 0x4d, 0x7c,
	0x21, 0x67, 0x36, 0xdf, 0x34, 0x63, 0xff, 0x57, 0xb6, 0x49, 0xd0, 0xae, 0x5d, 0x51, 0xdb, 0x40,
	0xd6, 0xb0, 0x95, 0xad, 0xaf, 0x80, 0x9a, 0x64, 0xef, 0xec, 0xe9, 0xe3, 0xeb, 0x59, 0x1f, 0x99,
	0xc3, 0xf2, 0x56, 0xf7, 0x8d, 0x75, 0x2e, 0x82, 0x1c, 0x32, 0x11, 0xf3, 0x14, 0xf0, 0x00, 0xd3,
	0x06, 0x06, 0x78, 0xf0, 0x30, 0x79, 0xd4, 0x8d, 0x7e, 0x20, 0xe2, 0x9d, 0xf3, 0x4e, 0xac, 0xad,
	0xfb, 0xcf, 0xca, 0xfb, 0xce, 0xb4, 0xeb, 0xcb, 0x8a, 0x0f, 0xc5, 0x92, 0x26, 0x21, 0x12, 0x69,
	0x88, 0x43, 0x9e, 0x6c, 0xae, 0xaf, 0xbe, 0xe4, 0x2a, 0x82, 0xf6, 0x7f, 0xe7, 0x5c, 0x89, 0x17,
	0xfd, 0xcf, 0xd4, 0x71, 0x5e, 0xf5, 0xee, 0x54, 0x06, 0x3a, 0x0c, 0x90, 0x94, 0xa5, 0x72, 0x6d,
	0x54, 0x35, 0x86, 0x77, 0x65, 0xf1, 0x1c, 0x06, 0x5e, 0x6d, 0xf1, 0x5c, 0xdb, 0x53, 0x96, 0x6f,
	0xbd, 0x2f, 0xdf, 0x09, 0x71, 0x18, 0x10, 0x52, 0x9b, 0x08, 0x71, 0x6d, 0x42, 0x94, 0xed, 0xe6,
	0xef, 0x66, 0xce, 0xd1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xfb, 0x20, 0x21, 0xe2, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdScheduleViewServiceClient is the client API for AdScheduleViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdScheduleViewServiceClient interface {
	// Returns the requested ad schedule view in full detail.
	GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error)
}

type adScheduleViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdScheduleViewServiceClient(cc *grpc.ClientConn) AdScheduleViewServiceClient {
	return &adScheduleViewServiceClient{cc}
}

func (c *adScheduleViewServiceClient) GetAdScheduleView(ctx context.Context, in *GetAdScheduleViewRequest, opts ...grpc.CallOption) (*resources.AdScheduleView, error) {
	out := new(resources.AdScheduleView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdScheduleViewService/GetAdScheduleView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdScheduleViewServiceServer is the server API for AdScheduleViewService service.
type AdScheduleViewServiceServer interface {
	// Returns the requested ad schedule view in full detail.
	GetAdScheduleView(context.Context, *GetAdScheduleViewRequest) (*resources.AdScheduleView, error)
}

func RegisterAdScheduleViewServiceServer(s *grpc.Server, srv AdScheduleViewServiceServer) {
	s.RegisterService(&_AdScheduleViewService_serviceDesc, srv)
}

func _AdScheduleViewService_GetAdScheduleView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdScheduleViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdScheduleViewService/GetAdScheduleView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdScheduleViewServiceServer).GetAdScheduleView(ctx, req.(*GetAdScheduleViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdScheduleViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AdScheduleViewService",
	HandlerType: (*AdScheduleViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdScheduleView",
			Handler:    _AdScheduleViewService_GetAdScheduleView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/ad_schedule_view_service.proto",
}
