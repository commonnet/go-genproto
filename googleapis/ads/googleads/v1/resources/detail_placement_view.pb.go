// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/detail_placement_view.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	// The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The automatic placement string at detail level, e. g. website URL, mobile
	// application ID, or a YouTube video ID.
	Placement *wrappers.StringValue `protobuf:"bytes,2,opt,name=placement,proto3" json:"placement,omitempty"`
	// The display name is URL name for websites, YouTube video name for YouTube
	// videos, and translated mobile app name for mobile apps.
	DisplayName *wrappers.StringValue `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// URL of the group placement, e.g. domain, link to the mobile application in
	// app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3" json:"group_placement_target_url,omitempty"`
	// URL of the placement, e.g. website, link to the mobile application in app
	// store, or a YouTube video URL.
	TargetUrl *wrappers.StringValue `protobuf:"bytes,5,opt,name=target_url,json=targetUrl,proto3" json:"target_url,omitempty"`
	// Type of the placement, e.g. Website, YouTube Video, and Mobile Application.
	PlacementType        enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v1.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *DetailPlacementView) Reset()         { *m = DetailPlacementView{} }
func (m *DetailPlacementView) String() string { return proto.CompactTextString(m) }
func (*DetailPlacementView) ProtoMessage()    {}
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return fileDescriptor_9703f91bccf33f03, []int{0}
}

func (m *DetailPlacementView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailPlacementView.Unmarshal(m, b)
}
func (m *DetailPlacementView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailPlacementView.Marshal(b, m, deterministic)
}
func (m *DetailPlacementView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailPlacementView.Merge(m, src)
}
func (m *DetailPlacementView) XXX_Size() int {
	return xxx_messageInfo_DetailPlacementView.Size(m)
}
func (m *DetailPlacementView) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailPlacementView.DiscardUnknown(m)
}

var xxx_messageInfo_DetailPlacementView proto.InternalMessageInfo

func (m *DetailPlacementView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DetailPlacementView) GetPlacement() *wrappers.StringValue {
	if m != nil {
		return m.Placement
	}
	return nil
}

func (m *DetailPlacementView) GetDisplayName() *wrappers.StringValue {
	if m != nil {
		return m.DisplayName
	}
	return nil
}

func (m *DetailPlacementView) GetGroupPlacementTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.GroupPlacementTargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetTargetUrl() *wrappers.StringValue {
	if m != nil {
		return m.TargetUrl
	}
	return nil
}

func (m *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if m != nil {
		return m.PlacementType
	}
	return enums.PlacementTypeEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*DetailPlacementView)(nil), "google.ads.googleads.v1.resources.DetailPlacementView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/detail_placement_view.proto", fileDescriptor_9703f91bccf33f03)
}

var fileDescriptor_9703f91bccf33f03 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x8a, 0xdb, 0x30,
	0x14, 0xc5, 0x4e, 0x3b, 0x30, 0x9a, 0xc7, 0xc2, 0x5d, 0xd4, 0x84, 0xa1, 0x64, 0xfa, 0x80, 0xac,
	0x64, 0x9c, 0xee, 0x34, 0xb4, 0xc5, 0x43, 0xcb, 0x40, 0x17, 0x25, 0x64, 0x5a, 0x43, 0x4b, 0x20,
	0x68, 0xe2, 0x5b, 0x21, 0xb0, 0x25, 0x21, 0xc9, 0x0e, 0xd9, 0xf7, 0x4b, 0xba, 0xec, 0x07, 0xf4,
	0x23, 0xfa, 0x29, 0xfd, 0x8a, 0x12, 0xdb, 0x52, 0x32, 0xd0, 0xb4, 0xd9, 0x5d, 0xe9, 0x9e, 0x73,
	0x74, 0x8e, 0xee, 0x45, 0xaf, 0x98, 0x94, 0xac, 0x84, 0x84, 0x16, 0x26, 0xe9, 0xca, 0x4d, 0xd5,
	0xa4, 0x89, 0x06, 0x23, 0x6b, 0xbd, 0x04, 0x93, 0x14, 0x60, 0x29, 0x2f, 0x17, 0xaa, 0xa4, 0x4b,
	0xa8, 0x40, 0xd8, 0x45, 0xc3, 0x61, 0x85, 0x95, 0x96, 0x56, 0x46, 0x97, 0x1d, 0x07, 0xd3, 0xc2,
	0x60, 0x4f, 0xc7, 0x4d, 0x8a, 0x3d, 0x7d, 0x38, 0xd9, 0xf7, 0x02, 0x88, 0xba, 0x32, 0xc9, 0x56,
	0xd6, 0xae, 0x15, 0x74, 0xb2, 0xc3, 0x27, 0x3d, 0xa7, 0x3d, 0xdd, 0xd5, 0x5f, 0x93, 0x95, 0xa6,
	0x4a, 0x81, 0x36, 0x7d, 0xff, 0xc2, 0x69, 0x2a, 0x9e, 0x50, 0x21, 0xa4, 0xa5, 0x96, 0x4b, 0xd1,
	0x77, 0x9f, 0xfe, 0x1c, 0xa0, 0x47, 0x6f, 0x5b, 0xd3, 0x53, 0x27, 0x9e, 0x73, 0x58, 0x45, 0xcf,
	0xd0, 0x99, 0xb3, 0xb5, 0x10, 0xb4, 0x82, 0x38, 0x18, 0x05, 0xe3, 0xe3, 0xd9, 0xa9, 0xbb, 0xfc,
	0x40, 0x2b, 0x88, 0x08, 0x3a, 0xf6, 0x96, 0xe2, 0x70, 0x14, 0x8c, 0x4f, 0x26, 0x17, 0x7d, 0x34,
	0xec, 0xec, 0xe0, 0x5b, 0xab, 0xb9, 0x60, 0x39, 0x2d, 0x6b, 0x98, 0x6d, 0xe1, 0xd1, 0x1b, 0x74,
	0x5a, 0x70, 0xa3, 0x4a, 0xba, 0xee, 0xf4, 0x07, 0x07, 0xd0, 0x4f, 0x7a, 0x46, 0xfb, 0xf8, 0x67,
	0x34, 0x64, 0x5a, 0xd6, 0x6a, 0xe7, 0xb3, 0x2d, 0xd5, 0x0c, 0xec, 0xa2, 0xd6, 0x65, 0xfc, 0xe0,
	0x00, 0xb9, 0xc7, 0x2d, 0xdf, 0xe7, 0xfe, 0xd8, 0xb2, 0x3f, 0xe9, 0x32, 0xba, 0x42, 0x68, 0x47,
	0xea, 0xe1, 0x21, 0xc1, 0xac, 0x27, 0x03, 0x3a, 0xbf, 0x3f, 0xa7, 0xf8, 0x68, 0x14, 0x8c, 0xcf,
	0x27, 0xaf, 0xf1, 0xbe, 0xf9, 0xb7, 0xc3, 0xc5, 0x5b, 0x1f, 0x6b, 0x05, 0xef, 0x44, 0x5d, 0xdd,
	0xbf, 0x99, 0x9d, 0xa9, 0xdd, 0xe3, 0xf5, 0xb7, 0x10, 0xbd, 0x58, 0xca, 0x0a, 0xff, 0x77, 0xa9,
	0xae, 0xe3, 0xbf, 0xcc, 0x77, 0xba, 0x09, 0x31, 0x0d, 0xbe, 0xbc, 0xef, 0xe9, 0x4c, 0x96, 0x54,
	0x30, 0x2c, 0x35, 0x4b, 0x18, 0x88, 0x36, 0xa2, 0x5b, 0x40, 0xc5, 0xcd, 0x3f, 0x36, 0xfe, 0xca,
	0x57, 0xdf, 0xc3, 0xc1, 0x4d, 0x96, 0xfd, 0x08, 0x2f, 0x6f, 0x3a, 0xc9, 0xac, 0x30, 0xb8, 0x2b,
	0x37, 0x55, 0x9e, 0xe2, 0x99, 0x43, 0xfe, 0x72, 0x98, 0x79, 0x56, 0x98, 0xb9, 0xc7, 0xcc, 0xf3,
	0x74, 0xee, 0x31, 0xbf, 0xc3, 0xe7, 0x5d, 0x83, 0x90, 0xac, 0x30, 0x84, 0x78, 0x14, 0x21, 0x79,
	0x4a, 0xc8, 0x2d, 0xe8, 0x86, 0x2f, 0xc1, 0xdc, 0x1d, 0xb5, 0x5e, 0x5f, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xe3, 0x2c, 0x66, 0xeb, 0x9c, 0x03, 0x00, 0x00,
}
