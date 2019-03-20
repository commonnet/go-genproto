// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/campaign_audience_view_service.proto

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
// [CampaignAudienceViewService.GetCampaignAudienceView][google.ads.googleads.v0.services.CampaignAudienceViewService.GetCampaignAudienceView].
type GetCampaignAudienceViewRequest struct {
	// The resource name of the campaign audience view to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCampaignAudienceViewRequest) Reset()         { *m = GetCampaignAudienceViewRequest{} }
func (m *GetCampaignAudienceViewRequest) String() string { return proto.CompactTextString(m) }
func (*GetCampaignAudienceViewRequest) ProtoMessage()    {}
func (*GetCampaignAudienceViewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a136887176def0f8, []int{0}
}

func (m *GetCampaignAudienceViewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Unmarshal(m, b)
}
func (m *GetCampaignAudienceViewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Marshal(b, m, deterministic)
}
func (m *GetCampaignAudienceViewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCampaignAudienceViewRequest.Merge(m, src)
}
func (m *GetCampaignAudienceViewRequest) XXX_Size() int {
	return xxx_messageInfo_GetCampaignAudienceViewRequest.Size(m)
}
func (m *GetCampaignAudienceViewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCampaignAudienceViewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCampaignAudienceViewRequest proto.InternalMessageInfo

func (m *GetCampaignAudienceViewRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCampaignAudienceViewRequest)(nil), "google.ads.googleads.v0.services.GetCampaignAudienceViewRequest")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/campaign_audience_view_service.proto", fileDescriptor_a136887176def0f8)
}

var fileDescriptor_a136887176def0f8 = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x6b, 0xdb, 0x40,
	0x18, 0x46, 0x2a, 0x14, 0x2a, 0xda, 0x45, 0x4b, 0x8b, 0x5b, 0x8a, 0x70, 0x3d, 0x14, 0x0f, 0x77,
	0xa2, 0x1d, 0x4c, 0x2e, 0xe4, 0x43, 0x0e, 0xc6, 0x99, 0x82, 0x71, 0x40, 0x43, 0x10, 0x88, 0x8b,
	0x74, 0x88, 0x03, 0xeb, 0x4e, 0xd1, 0x2b, 0xc9, 0x43, 0xc8, 0x92, 0x25, 0x3f, 0x20, 0xff, 0x20,
	0x63, 0x7e, 0x4a, 0xd6, 0x90, 0x7f, 0x90, 0x29, 0x7b, 0xf6, 0x20, 0x9f, 0x4e, 0x24, 0x60, 0xd9,
	0xdb, 0xc3, 0xdd, 0xf3, 0x71, 0xef, 0xf3, 0x9e, 0x35, 0x49, 0xa4, 0x4c, 0x16, 0x0c, 0xd3, 0x18,
	0xb0, 0x82, 0x35, 0xaa, 0x5c, 0x0c, 0x2c, 0xaf, 0x78, 0xc4, 0x00, 0x47, 0x34, 0xcd, 0x28, 0x4f,
	0x44, 0x48, 0xcb, 0x98, 0x33, 0x11, 0xb1, 0xb0, 0xe2, 0x6c, 0x19, 0x36, 0xf7, 0x28, 0xcb, 0x65,
	0x21, 0x6d, 0x47, 0x69, 0x11, 0x8d, 0x01, 0xb5, 0x36, 0xa8, 0x72, 0x91, 0xb6, 0xe9, 0xed, 0x77,
	0x05, 0xe5, 0x0c, 0x64, 0x99, 0x77, 0x27, 0xa9, 0x84, 0xde, 0x2f, 0xad, 0xcf, 0x38, 0xa6, 0x42,
	0xc8, 0x82, 0x16, 0x5c, 0x0a, 0x50, 0xb7, 0xfd, 0x89, 0xf5, 0x7b, 0xca, 0x8a, 0xa3, 0xc6, 0xc0,
	0x6b, 0xf4, 0x3e, 0x67, 0xcb, 0x39, 0xbb, 0x28, 0x19, 0x14, 0xf6, 0x1f, 0xeb, 0x9b, 0x4e, 0x0a,
	0x05, 0x4d, 0xd9, 0x0f, 0xc3, 0x31, 0xfe, 0x7e, 0x99, 0x7f, 0xd5, 0x87, 0x27, 0x34, 0x65, 0xff,
	0x5e, 0x0d, 0xeb, 0xe7, 0x3a, 0x93, 0x53, 0x35, 0x85, 0xfd, 0x64, 0x58, 0xdf, 0x3b, 0x72, 0xec,
	0x43, 0xb4, 0xad, 0x03, 0xb4, 0xf9, 0x89, 0xbd, 0x51, 0xa7, 0x43, 0xdb, 0x11, 0x5a, 0xa7, 0xef,
	0x1f, 0x5c, 0x3f, 0x3e, 0xdf, 0x9a, 0x3b, 0xf6, 0xa8, 0xee, 0xf3, 0xf2, 0xc3, 0x98, 0x7b, 0x51,
	0x09, 0x85, 0x4c, 0x59, 0x0e, 0x78, 0xd8, 0x16, 0xfc, 0x5e, 0x0c, 0x78, 0x78, 0x35, 0xbe, 0x31,
	0xad, 0x41, 0x24, 0xd3, 0xad, 0x13, 0x8c, 0x9d, 0x0d, 0xed, 0xcc, 0xea, 0x4d, 0xcc, 0x8c, 0xb3,
	0xe3, 0xc6, 0x25, 0x91, 0x0b, 0x2a, 0x12, 0x24, 0xf3, 0x04, 0x27, 0x4c, 0xac, 0xf6, 0xa4, 0x37,
	0x9f, 0x71, 0xe8, 0xfe, 0x71, 0xbb, 0x1a, 0xdc, 0x99, 0x9f, 0xa6, 0x9e, 0x77, 0x6f, 0x3a, 0x53,
	0x65, 0xe8, 0xc5, 0x80, 0x14, 0xac, 0x91, 0xef, 0xa2, 0x26, 0x18, 0x1e, 0x34, 0x25, 0xf0, 0x62,
	0x08, 0x5a, 0x4a, 0xe0, 0xbb, 0x81, 0xa6, 0xbc, 0x98, 0x03, 0x75, 0x4e, 0x88, 0x17, 0x03, 0x21,
	0x2d, 0x89, 0x10, 0xdf, 0x25, 0x44, 0xd3, 0xce, 0x3f, 0xaf, 0xde, 0xf9, 0xff, 0x2d, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x06, 0x9c, 0xfc, 0x18, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CampaignAudienceViewServiceClient is the client API for CampaignAudienceViewService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CampaignAudienceViewServiceClient interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error)
}

type campaignAudienceViewServiceClient struct {
	cc *grpc.ClientConn
}

func NewCampaignAudienceViewServiceClient(cc *grpc.ClientConn) CampaignAudienceViewServiceClient {
	return &campaignAudienceViewServiceClient{cc}
}

func (c *campaignAudienceViewServiceClient) GetCampaignAudienceView(ctx context.Context, in *GetCampaignAudienceViewRequest, opts ...grpc.CallOption) (*resources.CampaignAudienceView, error) {
	out := new(resources.CampaignAudienceView)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.CampaignAudienceViewService/GetCampaignAudienceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CampaignAudienceViewServiceServer is the server API for CampaignAudienceViewService service.
type CampaignAudienceViewServiceServer interface {
	// Returns the requested campaign audience view in full detail.
	GetCampaignAudienceView(context.Context, *GetCampaignAudienceViewRequest) (*resources.CampaignAudienceView, error)
}

func RegisterCampaignAudienceViewServiceServer(s *grpc.Server, srv CampaignAudienceViewServiceServer) {
	s.RegisterService(&_CampaignAudienceViewService_serviceDesc, srv)
}

func _CampaignAudienceViewService_GetCampaignAudienceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCampaignAudienceViewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.CampaignAudienceViewService/GetCampaignAudienceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CampaignAudienceViewServiceServer).GetCampaignAudienceView(ctx, req.(*GetCampaignAudienceViewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CampaignAudienceViewService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.CampaignAudienceViewService",
	HandlerType: (*CampaignAudienceViewServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCampaignAudienceView",
			Handler:    _CampaignAudienceViewService_GetCampaignAudienceView_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/campaign_audience_view_service.proto",
}
