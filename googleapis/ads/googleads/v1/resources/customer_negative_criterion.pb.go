// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/customer_negative_criterion.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v1/common"
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

// A negative criterion for exclusions at the customer level.
type CustomerNegativeCriterion struct {
	// The resource name of the customer negative criterion.
	// Customer negative criterion resource names have the form:
	//
	// `customers/{customer_id}/customerNegativeCriteria/{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the criterion.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,3,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The customer negative criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CustomerNegativeCriterion_ContentLabel
	//	*CustomerNegativeCriterion_MobileApplication
	//	*CustomerNegativeCriterion_MobileAppCategory
	//	*CustomerNegativeCriterion_Placement
	//	*CustomerNegativeCriterion_YoutubeVideo
	//	*CustomerNegativeCriterion_YoutubeChannel
	Criterion            isCustomerNegativeCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *CustomerNegativeCriterion) Reset()         { *m = CustomerNegativeCriterion{} }
func (m *CustomerNegativeCriterion) String() string { return proto.CompactTextString(m) }
func (*CustomerNegativeCriterion) ProtoMessage()    {}
func (*CustomerNegativeCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_761ad3c7a7659e34, []int{0}
}

func (m *CustomerNegativeCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerNegativeCriterion.Unmarshal(m, b)
}
func (m *CustomerNegativeCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerNegativeCriterion.Marshal(b, m, deterministic)
}
func (m *CustomerNegativeCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerNegativeCriterion.Merge(m, src)
}
func (m *CustomerNegativeCriterion) XXX_Size() int {
	return xxx_messageInfo_CustomerNegativeCriterion.Size(m)
}
func (m *CustomerNegativeCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerNegativeCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerNegativeCriterion proto.InternalMessageInfo

func (m *CustomerNegativeCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CustomerNegativeCriterion) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isCustomerNegativeCriterion_Criterion interface {
	isCustomerNegativeCriterion_Criterion()
}

type CustomerNegativeCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,4,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CustomerNegativeCriterion_MobileApplication struct {
	MobileApplication *common.MobileApplicationInfo `protobuf:"bytes,5,opt,name=mobile_application,json=mobileApplication,proto3,oneof"`
}

type CustomerNegativeCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,6,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CustomerNegativeCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type CustomerNegativeCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,8,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CustomerNegativeCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,9,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

func (*CustomerNegativeCriterion_ContentLabel) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_MobileApplication) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_MobileAppCategory) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_Placement) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_YoutubeVideo) isCustomerNegativeCriterion_Criterion() {}

func (*CustomerNegativeCriterion_YoutubeChannel) isCustomerNegativeCriterion_Criterion() {}

func (m *CustomerNegativeCriterion) GetCriterion() isCustomerNegativeCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetMobileApplication() *common.MobileApplicationInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_MobileApplication); ok {
		return x.MobileApplication
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CustomerNegativeCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CustomerNegativeCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CustomerNegativeCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CustomerNegativeCriterion_ContentLabel)(nil),
		(*CustomerNegativeCriterion_MobileApplication)(nil),
		(*CustomerNegativeCriterion_MobileAppCategory)(nil),
		(*CustomerNegativeCriterion_Placement)(nil),
		(*CustomerNegativeCriterion_YoutubeVideo)(nil),
		(*CustomerNegativeCriterion_YoutubeChannel)(nil),
	}
}

func init() {
	proto.RegisterType((*CustomerNegativeCriterion)(nil), "google.ads.googleads.v1.resources.CustomerNegativeCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/customer_negative_criterion.proto", fileDescriptor_761ad3c7a7659e34)
}

var fileDescriptor_761ad3c7a7659e34 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0x6d, 0xf6, 0x43, 0x3b, 0xfb, 0x21, 0xc6, 0x9b, 0xb8, 0xca, 0xd2, 0x55, 0x16, 0x0a,
	0xb2, 0x13, 0x5b, 0x3f, 0x2e, 0x22, 0x08, 0x69, 0x90, 0x75, 0xc5, 0x5d, 0x4a, 0x58, 0x2a, 0x4a,
	0x25, 0x4c, 0x92, 0xd3, 0x18, 0x48, 0x66, 0x42, 0x66, 0x52, 0xe9, 0x33, 0xf8, 0x16, 0x5e, 0xfa,
	0x28, 0x3e, 0x8a, 0x2f, 0xa1, 0x64, 0x32, 0x93, 0xfa, 0x41, 0xed, 0x7a, 0x77, 0x7a, 0xce, 0xf9,
	0xff, 0xce, 0x99, 0x3f, 0x3d, 0x41, 0x5e, 0xc2, 0x58, 0x92, 0x81, 0x4d, 0x62, 0x6e, 0x37, 0x61,
	0x1d, 0xcd, 0x07, 0x76, 0x09, 0x9c, 0x55, 0x65, 0x04, 0xdc, 0x8e, 0x2a, 0x2e, 0x58, 0x0e, 0x65,
	0x40, 0x21, 0x21, 0x22, 0x9d, 0x43, 0x10, 0x95, 0xa9, 0x80, 0x32, 0x65, 0x14, 0x17, 0x25, 0x13,
	0xcc, 0x3c, 0x6a, 0x94, 0x98, 0xc4, 0x1c, 0xb7, 0x10, 0x3c, 0x1f, 0xe0, 0x16, 0x72, 0x70, 0xb2,
	0x6a, 0x4e, 0xc4, 0xf2, 0x9c, 0x51, 0x5b, 0x21, 0x49, 0x43, 0x3c, 0x18, 0xae, 0x6a, 0x07, 0x5a,
	0xe5, 0xdc, 0x6e, 0x17, 0x08, 0xc4, 0xa2, 0x00, 0xa5, 0x39, 0x54, 0x1a, 0xf9, 0x2b, 0xac, 0x66,
	0xf6, 0xa7, 0x92, 0x14, 0x05, 0x94, 0x5c, 0xd5, 0xef, 0x69, 0x66, 0x91, 0xda, 0x84, 0x52, 0x26,
	0x88, 0x48, 0x19, 0x55, 0xd5, 0xfb, 0x3f, 0xb6, 0xd0, 0x1d, 0x4f, 0xbd, 0xf4, 0x42, 0x3d, 0xd4,
	0xd3, 0x63, 0xcc, 0x07, 0x68, 0x4f, 0xbf, 0x25, 0xa0, 0x24, 0x07, 0xab, 0xd3, 0xeb, 0xf4, 0xbb,
	0xfe, 0xae, 0x4e, 0x5e, 0x90, 0x1c, 0xcc, 0x87, 0xc8, 0x48, 0x63, 0xcb, 0xe8, 0x75, 0xfa, 0x3b,
	0xc3, 0xbb, 0xca, 0x08, 0xac, 0xb7, 0xc1, 0x67, 0x54, 0x3c, 0x7b, 0x32, 0x21, 0x59, 0x05, 0xbe,
	0x91, 0xc6, 0xa6, 0x8f, 0x36, 0xeb, 0xdd, 0xad, 0x8d, 0x5e, 0xa7, 0xbf, 0x3f, 0x7c, 0x81, 0x57,
	0x59, 0x28, 0x1f, 0x8c, 0xdb, 0x4d, 0x2e, 0x17, 0x05, 0xbc, 0xa4, 0x55, 0xfe, 0x7b, 0xc6, 0x97,
	0x2c, 0xf3, 0x2d, 0xda, 0x8b, 0x18, 0x15, 0x40, 0x45, 0x90, 0x91, 0x10, 0x32, 0x6b, 0x53, 0xee,
	0xf2, 0x68, 0x25, 0xbc, 0x31, 0x1f, 0x7b, 0x8d, 0xe8, 0x4d, 0xad, 0x39, 0xa3, 0x33, 0xf6, 0xea,
	0x9a, 0xbf, 0x1b, 0xfd, 0x92, 0x33, 0x67, 0xc8, 0xcc, 0x59, 0x98, 0x66, 0x10, 0x90, 0xa2, 0xc8,
	0xd2, 0x48, 0x3a, 0x67, 0x6d, 0x49, 0xfa, 0xd3, 0x75, 0xf4, 0x73, 0xa9, 0x74, 0x97, 0x42, 0x35,
	0xe2, 0x56, 0xfe, 0x67, 0xc1, 0x4c, 0xd0, 0xed, 0xe5, 0x9c, 0x20, 0x22, 0x02, 0x12, 0x56, 0x2e,
	0xac, 0xed, 0xff, 0x1c, 0xe4, 0x29, 0xe1, 0x5f, 0x83, 0x74, 0xc1, 0x3c, 0x47, 0xdd, 0x22, 0x23,
	0x11, 0xe4, 0x40, 0x85, 0x75, 0x5d, 0xe2, 0x4f, 0xd6, 0xe1, 0xc7, 0x5a, 0xa0, 0xb0, 0x4b, 0x42,
	0x6d, 0xfc, 0x82, 0x55, 0xa2, 0x0a, 0x21, 0x98, 0xa7, 0x31, 0x30, 0xeb, 0xc6, 0xd5, 0x8c, 0x7f,
	0xc7, 0xaa, 0xcb, 0x2a, 0x84, 0x49, 0xad, 0xd1, 0xc6, 0x2b, 0x90, 0xcc, 0x99, 0x1f, 0xd0, 0x4d,
	0x0d, 0x8e, 0x3e, 0x12, 0x4a, 0x21, 0xb3, 0xba, 0x12, 0x3d, 0xbc, 0x22, 0xda, 0x6b, 0x54, 0x0a,
	0xbe, 0xaf, 0x60, 0x2a, 0x3b, 0xda, 0x41, 0xdd, 0xf6, 0x94, 0x46, 0x9f, 0x0d, 0x74, 0x1c, 0xb1,
	0x1c, 0xaf, 0x3d, 0xe6, 0xd1, 0xe1, 0xca, 0x43, 0x19, 0xd7, 0xff, 0xf6, 0x71, 0xe7, 0xfd, 0x6b,
	0x05, 0x49, 0x58, 0x46, 0x68, 0x82, 0x59, 0x99, 0xd8, 0x09, 0x50, 0x79, 0x0b, 0xfa, 0x9e, 0x8b,
	0x94, 0xff, 0xe3, 0xab, 0xf3, 0xbc, 0x8d, 0xbe, 0x18, 0x1b, 0xa7, 0xae, 0xfb, 0xd5, 0x38, 0x3a,
	0x6d, 0x90, 0x6e, 0xcc, 0x71, 0x13, 0xd6, 0xd1, 0x64, 0x80, 0x7d, 0xdd, 0xf9, 0x4d, 0xf7, 0x4c,
	0xdd, 0x98, 0x4f, 0xdb, 0x9e, 0xe9, 0x64, 0x30, 0x6d, 0x7b, 0xbe, 0x1b, 0xc7, 0x4d, 0xc1, 0x71,
	0xdc, 0x98, 0x3b, 0x4e, 0xdb, 0xe5, 0x38, 0x93, 0x81, 0xe3, 0xb4, 0x7d, 0xe1, 0xb6, 0x5c, 0xf6,
	0xf1, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x01, 0xad, 0x51, 0x21, 0x05, 0x00, 0x00,
}
