// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/hotel_group_view_service.proto

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
// [HotelGroupViewService.GetHotelGroupView][google.ads.googleads.v0.services.HotelGroupViewService.GetHotelGroupView].
type GetHotelGroupViewRequest struct {
	// Resource name of the Hotel Group View to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHotelGroupViewRequest) Reset()         { *m = GetHotelGroupViewRequest{} }
func (m *GetHotelGroupViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetHotelGroupViewRequest) ProtoMessage()    {}
func (*GetHotelGroupViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ebbcd60766d696a, []int{0}
}

func (m *GetHotelGroupViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHotelGroupViewRequest.Unmarshal(m, b)
}
func (m *GetHotelGroupViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHotelGroupViewRequest.Marshal(b, m, deterministic)
}
func (m *GetHotelGroupViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHotelGroupViewRequest.Merge(m, src)
}
func (m *GetHotelGroupViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetHotelGroupViewRequest.Size(m)
}
func (m *GetHotelGroupViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHotelGroupViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHotelGroupViewRequest proto.InternalMessageInfo

func (m *GetHotelGroupViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetHotelGroupViewRequest)(nil), "google.ads.googleads.v0.services.GetHotelGroupViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/hotel_group_view_service.proto", fileDescriptor_2ebbcd60766d696a)
}

var fileDescriptor_2ebbcd60766d696a = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0x49, 0x04, 0xc1, 0xa0, 0x83, 0x01, 0xa1, 0x04, 0x87, 0x52, 0x3b, 0x48, 0x87, 0xbb,
	0xd4, 0x2e, 0x7a, 0x22, 0x25, 0x5d, 0xd2, 0x49, 0x4a, 0x85, 0x0c, 0x12, 0x08, 0xb1, 0xf9, 0x88,
	0x81, 0x26, 0x17, 0xf3, 0x25, 0xe9, 0x20, 0x0e, 0xfa, 0x0a, 0xbe, 0x81, 0xa3, 0xef, 0xe0, 0x0b,
	0xb8, 0x3a, 0xf8, 0x02, 0x4e, 0x3e, 0x85, 0xa4, 0xd7, 0x0b, 0x14, 0x1b, 0xba, 0xfd, 0xb9, 0xfb,
	0xff, 0xfe, 0xf7, 0x7d, 0xff, 0x44, 0x1b, 0x86, 0x9c, 0x87, 0x73, 0xa0, 0x7e, 0x80, 0x54, 0xc8,
	0x4a, 0x95, 0x26, 0x45, 0xc8, 0xca, 0x68, 0x06, 0x48, 0xef, 0x79, 0x0e, 0x73, 0x2f, 0xcc, 0x78,
	0x91, 0x7a, 0x65, 0x04, 0x0b, 0x6f, 0x75, 0x43, 0xd2, 0x8c, 0xe7, 0x5c, 0x6f, 0x0b, 0x8a, 0xf8,
	0x01, 0x92, 0x3a, 0x80, 0x94, 0x26, 0x91, 0x01, 0xc6, 0x79, 0xd3, 0x13, 0x19, 0x20, 0x2f, 0xb2,
	0x4d, 0x6f, 0x88, 0x6c, 0xe3, 0x58, 0x92, 0x69, 0x44, 0xfd, 0x24, 0xe1, 0xb9, 0x9f, 0x47, 0x3c,
	0x41, 0x71, 0xdb, 0x19, 0x6a, 0x2d, 0x1b, 0xf2, 0x71, 0x85, 0xda, 0x15, 0xe9, 0x44, 0xb0, 0x98,
	0xc2, 0x43, 0x01, 0x98, 0xeb, 0x27, 0xda, 0x81, 0x4c, 0xf7, 0x12, 0x3f, 0x86, 0x96, 0xd2, 0x56,
	0x4e, 0xf7, 0xa6, 0xfb, 0xf2, 0xf0, 0xda, 0x8f, 0xe1, 0xec, 0x5b, 0xd1, 0x8e, 0xd6, 0xf1, 0x1b,
	0x31, 0xb3, 0xfe, 0xa1, 0x68, 0x87, 0xff, 0xb2, 0x75, 0x46, 0xb6, 0xed, 0x4a, 0x9a, 0x06, 0x32,
	0xfa, 0x8d, 0x6c, 0xdd, 0x02, 0x59, 0x27, 0x3b, 0x17, 0x2f, 0x5f, 0x3f, 0xaf, 0xea, 0x40, 0xef,
	0x57, 0x5d, 0x3d, 0xae, 0xad, 0x73, 0x35, 0x2b, 0x30, 0xe7, 0x31, 0x64, 0x48, 0x7b, 0xa2, 0xbc,
	0x1a, 0x43, 0xda, 0x7b, 0x1a, 0x3d, 0xab, 0x5a, 0x77, 0xc6, 0xe3, 0xad, 0xf3, 0x8e, 0x8c, 0x8d,
	0xfb, 0x4f, 0xaa, 0x7e, 0x27, 0xca, 0xed, 0x78, 0xc5, 0x87, 0x7c, 0xee, 0x27, 0x21, 0xe1, 0x59,
	0x48, 0x43, 0x48, 0x96, 0xed, 0xcb, 0x2f, 0x99, 0x46, 0xd8, 0xfc, 0xef, 0x5c, 0x4a, 0xf1, 0xa6,
	0xee, 0xd8, 0x96, 0xf5, 0xae, 0xb6, 0x6d, 0x11, 0x68, 0x05, 0x48, 0x84, 0xac, 0x94, 0x63, 0x92,
	0xd5, 0xc3, 0xf8, 0x29, 0x2d, 0xae, 0x15, 0xa0, 0x5b, 0x5b, 0x5c, 0xc7, 0x74, 0xa5, 0xe5, 0x57,
	0xed, 0x8a, 0x73, 0xc6, 0xac, 0x00, 0x19, 0xab, 0x4d, 0x8c, 0x39, 0x26, 0x63, 0xd2, 0x76, 0xb7,
	0xbb, 0x9c, 0x73, 0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x2d, 0x0c, 0x9f, 0xe2, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HotelGroupViewServiceClient is the client API for HotelGroupViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HotelGroupViewServiceClient interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error)
}

type hotelGroupViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewHotelGroupViewServiceClient(cc *grpc.ClientConn) HotelGroupViewServiceClient {
	return &hotelGroupViewServiceClient{cc}
}

func (c *hotelGroupViewServiceClient) GetHotelGroupView(ctx context.Context, in *GetHotelGroupViewRequest, opts ...grpc.CallOption) (*resources.HotelGroupView, error) {
	out := new(resources.HotelGroupView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.HotelGroupViewService/GetHotelGroupView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HotelGroupViewServiceServer is the server API for HotelGroupViewService service.
type HotelGroupViewServiceServer interface {
	// Returns the requested Hotel Group View in full detail.
	GetHotelGroupView(context.Context, *GetHotelGroupViewRequest) (*resources.HotelGroupView, error)
}

func RegisterHotelGroupViewServiceServer(s *grpc.Server, srv HotelGroupViewServiceServer) {
	s.RegisterService(&_HotelGroupViewService_serviceDesc, srv)
}

func _HotelGroupViewService_GetHotelGroupView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHotelGroupViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.HotelGroupViewService/GetHotelGroupView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HotelGroupViewServiceServer).GetHotelGroupView(ctx, req.(*GetHotelGroupViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HotelGroupViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.HotelGroupViewService",
	HandlerType: (*HotelGroupViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHotelGroupView",
			Handler:    _HotelGroupViewService_GetHotelGroupView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/hotel_group_view_service.proto",
}
