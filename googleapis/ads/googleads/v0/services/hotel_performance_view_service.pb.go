// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/hotel_performance_view_service.proto

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
// [HotelPerformanceViewService.GetHotelPerformanceView][google.ads.googleads.v0.services.HotelPerformanceViewService.GetHotelPerformanceView].
type GetHotelPerformanceViewRequest struct {
	// Resource name of the Hotel Performance View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotelPerformanceViewRequest) Reset()         { *m = GetHotelPerformanceViewRequest{} }
func (m *GetHotelPerformanceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotelPerformanceViewRequest) ProtoMessage()    {}
func (*GetHotelPerformanceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_19330a473f167e68, []int{0}
}

func (m *GetHotelPerformanceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Unmarshal(m, b)
}
func (m *GetHotelPerformanceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetHotelPerformanceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotelPerformanceViewRequest.Merge(m, src)
}
func (m *GetHotelPerformanceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotelPerformanceViewRequest.Size(m)
}
func (m *GetHotelPerformanceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotelPerformanceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotelPerformanceViewRequest proto.InternalMessageInfo

func (m *GetHotelPerformanceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHotelPerformanceViewRequest)(nil), "google.ads.googleads.v0.services.GetHotelPerformanceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/hotel_performance_view_service.proto", fileDescriptor_19330a473f167e68)
}

var fileDescriptor_19330a473f167e68 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0xf9, 0xe0, 0x03, 0x83, 0x6e, 0xb2, 0x51, 0xaa, 0x48, 0xa8, 0x5d, 0x88, 0x8b, 0x49,
	0x50, 0x51, 0x18, 0x7f, 0x30, 0x85, 0xd2, 0xae, 0xa4, 0x54, 0xc8, 0x42, 0x02, 0x61, 0x4c, 0xae,
	0x31, 0xd0, 0x64, 0xe2, 0xcc, 0x34, 0x5d, 0x88, 0x1b, 0x37, 0x3e, 0x80, 0x6f, 0xe0, 0xd2, 0x47,
	0x71, 0xa9, 0xaf, 0xe0, 0xca, 0xa5, 0x4f, 0x20, 0xe9, 0x74, 0x22, 0x4a, 0xd3, 0xee, 0x0e, 0x33,
	0xe7, 0xe7, 0xce, 0xb9, 0x63, 0x74, 0x62, 0x4a, 0xe3, 0x21, 0xd8, 0x24, 0xe2, 0xb6, 0x84, 0x25,
	0x2a, 0x1c, 0x9b, 0x03, 0x2b, 0x92, 0x10, 0xb8, 0x7d, 0x43, 0x05, 0x0c, 0x83, 0x1c, 0xd8, 0x35,
	0x65, 0x29, 0xc9, 0x42, 0x08, 0x8a, 0x04, 0xc6, 0xc1, 0xf4, 0x1e, 0xe5, 0x8c, 0x0a, 0x6a, 0x5a,
	0x52, 0x8b, 0x48, 0xc4, 0x51, 0x65, 0x83, 0x0a, 0x07, 0x29, 0x9b, 0xc6, 0x69, 0x5d, 0x10, 0x03,
	0x4e, 0x47, 0xac, 0x3e, 0x49, 0x26, 0x34, 0x36, 0x94, 0x3e, 0x4f, 0x6c, 0x92, 0x65, 0x54, 0x10,
	0x91, 0xd0, 0x8c, 0xcb, 0xdb, 0x66, 0xc7, 0xd8, 0xec, 0x82, 0xe8, 0x95, 0x06, 0xfd, 0x1f, 0xbd,
	0x97, 0xc0, 0x78, 0x00, 0xb7, 0x23, 0xe0, 0xc2, 0xdc, 0x32, 0x56, 0x54, 0x52, 0x90, 0x91, 0x14,
	0xd6, 0x34, 0x4b, 0xdb, 0x5e, 0x1a, 0x2c, 0xab, 0xc3, 0x73, 0x92, 0xc2, 0xee, 0x97, 0x66, 0xac,
	0xcf, 0x32, 0xb9, 0x90, 0xaf, 0x30, 0xdf, 0x34, 0x63, 0xb5, 0x26, 0xc7, 0x3c, 0x43, 0x8b, 0x3a,
	0x40, 0xf3, 0x47, 0x6c, 0x1c, 0xd6, 0x3a, 0x54, 0x1d, 0xa1, 0x59, 0xfa, 0xe6, 0xf1, 0xc3, 0xfb,
	0xc7, 0x93, 0x7e, 0x60, 0xee, 0x97, 0x7d, 0xde, 0xfd, 0x7a, 0xe6, 0x49, 0x38, 0xe2, 0x82, 0xa6,
	0xc0, 0xb8, 0xbd, 0x23, 0x0b, 0xfe, 0x23, 0xbe, 0x6f, 0x3f, 0xea, 0x46, 0x2b, 0xa4, 0xe9, 0xc2,
	0xf1, 0xdb, 0xd6, 0x9c, 0x6a, 0xfa, 0xe5, 0x1a, 0xfa, 0xda, 0x65, 0x6f, 0xea, 0x12, 0xd3, 0x21,
	0xc9, 0x62, 0x44, 0x59, 0x6c, 0xc7, 0x90, 0x4d, 0x96, 0xa4, 0xd6, 0x9e, 0x27, 0xbc, 0xfe, 0xbb,
	0x1d, 0x29, 0xf0, 0xac, 0xff, 0xeb, 0xba, 0xee, 0x8b, 0x6e, 0x75, 0xa5, 0xa1, 0x1b, 0x71, 0x24,
	0x61, 0x89, 0x3c, 0x07, 0x4d, 0x83, 0xf9, 0xab, 0xa2, 0xf8, 0x6e, 0xc4, 0xfd, 0x8a, 0xe2, 0x7b,
	0x8e, 0xaf, 0x28, 0x9f, 0x7a, 0x4b, 0x9e, 0x63, 0xec, 0x46, 0x1c, 0xe3, 0x8a, 0x84, 0xb1, 0xe7,
	0x60, 0xac, 0x68, 0x57, 0xff, 0x27, 0x73, 0xee, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x59, 0x25,
	0xb5, 0xba, 0x15, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotelPerformanceViewServiceClient is the client API for HotelPerformanceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelPerformanceViewServiceClient interface {
	// Returns the requested Hotel Performance View in full detail.
	GetHotelPerformanceView(ctx context.Context, in *GetHotelPerformanceViewRequest, opts ...grpc.CallOption) (*resources.HotelPerformanceView, error)
}

type hotelPerformanceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewHotelPerformanceViewServiceClient(cc *grpc.ClientConn) HotelPerformanceViewServiceClient {
	return &hotelPerformanceViewServiceClient{cc}
}

func (c *hotelPerformanceViewServiceClient) GetHotelPerformanceView(ctx context.Context, in *GetHotelPerformanceViewRequest, opts ...grpc.CallOption) (*resources.HotelPerformanceView, error) {
	out := new(resources.HotelPerformanceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.HotelPerformanceViewService/GetHotelPerformanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelPerformanceViewServiceServer is the server API for HotelPerformanceViewService service.
type HotelPerformanceViewServiceServer interface {
	// Returns the requested Hotel Performance View in full detail.
	GetHotelPerformanceView(context.Context, *GetHotelPerformanceViewRequest) (*resources.HotelPerformanceView, error)
}

func RegisterHotelPerformanceViewServiceServer(s *grpc.Server, srv HotelPerformanceViewServiceServer) {
	s.RegisterService(&_HotelPerformanceViewService_serviceDesc, srv)
}

func _HotelPerformanceViewService_GetHotelPerformanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelPerformanceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelPerformanceViewServiceServer).GetHotelPerformanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.HotelPerformanceViewService/GetHotelPerformanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelPerformanceViewServiceServer).GetHotelPerformanceView(ctx, req.(*GetHotelPerformanceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelPerformanceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.HotelPerformanceViewService",
	HandlerType: (*HotelPerformanceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelPerformanceView",
			Handler:    _HotelPerformanceViewService_GetHotelPerformanceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/hotel_performance_view_service.proto",
}
