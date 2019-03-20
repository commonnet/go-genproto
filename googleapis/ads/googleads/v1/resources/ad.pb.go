// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/ad.proto

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

// An ad.
type Ad struct {
	// The ID of the ad.
	Id *wrappers.Int64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,2,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// The list of possible final mobile URLs after all cross-domain redirects
	// for the ad.
	FinalMobileUrls []*wrappers.StringValue `protobuf:"bytes,16,rep,name=final_mobile_urls,json=finalMobileUrls,proto3" json:"final_mobile_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,12,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings that can be used to substitute custom parameter tags
	// in a `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,10,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The URL that appears in the ad description for some ad formats.
	DisplayUrl *wrappers.StringValue `protobuf:"bytes,4,opt,name=display_url,json=displayUrl,proto3" json:"display_url,omitempty"`
	// The type of ad.
	Type enums.AdTypeEnum_AdType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.AdTypeEnum_AdType" json:"type,omitempty"`
	// Indicates if this ad was automatically added by Google Ads and not by a
	// user. For example, this could happen when ads are automatically created as
	// suggestions for new ads based on knowledge of how existing ads are
	// performing.
	AddedByGoogleAds *wrappers.BoolValue `protobuf:"bytes,19,opt,name=added_by_google_ads,json=addedByGoogleAds,proto3" json:"added_by_google_ads,omitempty"`
	// The device preference for the ad. You can only specify a preference for
	// mobile devices. When this preference is set the ad will be preferred over
	// other ads when being displayed on a mobile device. The ad can still be
	// displayed on other device types, e.g. if no other ads are available.
	// If unspecified (no device preference), all devices are targeted.
	// This is only supported by some ad types.
	DevicePreference enums.DeviceEnum_Device `protobuf:"varint,20,opt,name=device_preference,json=devicePreference,proto3,enum=google.ads.googleads.v1.enums.DeviceEnum_Device" json:"device_preference,omitempty"`
	// Additional URLs for the ad that are tagged with a unique identifier that
	// can be referenced from other fields in the ad.
	UrlCollections []*common.UrlCollection `protobuf:"bytes,26,rep,name=url_collections,json=urlCollections,proto3" json:"url_collections,omitempty"`
	// The name of the ad. This is only used to be able to identify the ad. It
	// does not need to be unique and does not affect the served ad.
	Name *wrappers.StringValue `protobuf:"bytes,23,opt,name=name,proto3" json:"name,omitempty"`
	// If this ad is system managed, then this field will indicate the source.
	// This field is read-only.
	SystemManagedResourceSource enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource `protobuf:"varint,27,opt,name=system_managed_resource_source,json=systemManagedResourceSource,proto3,enum=google.ads.googleads.v1.enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource" json:"system_managed_resource_source,omitempty"`
	// Details pertinent to the ad type. Exactly one value must be set.
	//
	// Types that are valid to be assigned to AdData:
	//	*Ad_TextAd
	//	*Ad_ExpandedTextAd
	//	*Ad_CallOnlyAd
	//	*Ad_ExpandedDynamicSearchAd
	//	*Ad_HotelAd
	//	*Ad_ShoppingSmartAd
	//	*Ad_ShoppingProductAd
	//	*Ad_GmailAd
	//	*Ad_ImageAd
	//	*Ad_VideoAd
	//	*Ad_ResponsiveSearchAd
	//	*Ad_LegacyResponsiveDisplayAd
	AdData               isAd_AdData `protobuf_oneof:"ad_data"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Ad) Reset()         { *m = Ad{} }
func (m *Ad) String() string { return proto.CompactTextString(m) }
func (*Ad) ProtoMessage()    {}
func (*Ad) Descriptor() ([]byte, []int) {
	return fileDescriptor_2cfecfd7772c305a, []int{0}
}

func (m *Ad) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ad.Unmarshal(m, b)
}
func (m *Ad) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ad.Marshal(b, m, deterministic)
}
func (m *Ad) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ad.Merge(m, src)
}
func (m *Ad) XXX_Size() int {
	return xxx_messageInfo_Ad.Size(m)
}
func (m *Ad) XXX_DiscardUnknown() {
	xxx_messageInfo_Ad.DiscardUnknown(m)
}

var xxx_messageInfo_Ad proto.InternalMessageInfo

func (m *Ad) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Ad) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *Ad) GetFinalMobileUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalMobileUrls
	}
	return nil
}

func (m *Ad) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *Ad) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

func (m *Ad) GetDisplayUrl() *wrappers.StringValue {
	if m != nil {
		return m.DisplayUrl
	}
	return nil
}

func (m *Ad) GetType() enums.AdTypeEnum_AdType {
	if m != nil {
		return m.Type
	}
	return enums.AdTypeEnum_UNSPECIFIED
}

func (m *Ad) GetAddedByGoogleAds() *wrappers.BoolValue {
	if m != nil {
		return m.AddedByGoogleAds
	}
	return nil
}

func (m *Ad) GetDevicePreference() enums.DeviceEnum_Device {
	if m != nil {
		return m.DevicePreference
	}
	return enums.DeviceEnum_UNSPECIFIED
}

func (m *Ad) GetUrlCollections() []*common.UrlCollection {
	if m != nil {
		return m.UrlCollections
	}
	return nil
}

func (m *Ad) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Ad) GetSystemManagedResourceSource() enums.SystemManagedResourceSourceEnum_SystemManagedResourceSource {
	if m != nil {
		return m.SystemManagedResourceSource
	}
	return enums.SystemManagedResourceSourceEnum_UNSPECIFIED
}

type isAd_AdData interface {
	isAd_AdData()
}

type Ad_TextAd struct {
	TextAd *common.TextAdInfo `protobuf:"bytes,6,opt,name=text_ad,json=textAd,proto3,oneof"`
}

type Ad_ExpandedTextAd struct {
	ExpandedTextAd *common.ExpandedTextAdInfo `protobuf:"bytes,7,opt,name=expanded_text_ad,json=expandedTextAd,proto3,oneof"`
}

type Ad_CallOnlyAd struct {
	CallOnlyAd *common.CallOnlyAdInfo `protobuf:"bytes,13,opt,name=call_only_ad,json=callOnlyAd,proto3,oneof"`
}

type Ad_ExpandedDynamicSearchAd struct {
	ExpandedDynamicSearchAd *common.ExpandedDynamicSearchAdInfo `protobuf:"bytes,14,opt,name=expanded_dynamic_search_ad,json=expandedDynamicSearchAd,proto3,oneof"`
}

type Ad_HotelAd struct {
	HotelAd *common.HotelAdInfo `protobuf:"bytes,15,opt,name=hotel_ad,json=hotelAd,proto3,oneof"`
}

type Ad_ShoppingSmartAd struct {
	ShoppingSmartAd *common.ShoppingSmartAdInfo `protobuf:"bytes,17,opt,name=shopping_smart_ad,json=shoppingSmartAd,proto3,oneof"`
}

type Ad_ShoppingProductAd struct {
	ShoppingProductAd *common.ShoppingProductAdInfo `protobuf:"bytes,18,opt,name=shopping_product_ad,json=shoppingProductAd,proto3,oneof"`
}

type Ad_GmailAd struct {
	GmailAd *common.GmailAdInfo `protobuf:"bytes,21,opt,name=gmail_ad,json=gmailAd,proto3,oneof"`
}

type Ad_ImageAd struct {
	ImageAd *common.ImageAdInfo `protobuf:"bytes,22,opt,name=image_ad,json=imageAd,proto3,oneof"`
}

type Ad_VideoAd struct {
	VideoAd *common.VideoAdInfo `protobuf:"bytes,24,opt,name=video_ad,json=videoAd,proto3,oneof"`
}

type Ad_ResponsiveSearchAd struct {
	ResponsiveSearchAd *common.ResponsiveSearchAdInfo `protobuf:"bytes,25,opt,name=responsive_search_ad,json=responsiveSearchAd,proto3,oneof"`
}

type Ad_LegacyResponsiveDisplayAd struct {
	LegacyResponsiveDisplayAd *common.LegacyResponsiveDisplayAdInfo `protobuf:"bytes,28,opt,name=legacy_responsive_display_ad,json=legacyResponsiveDisplayAd,proto3,oneof"`
}

func (*Ad_TextAd) isAd_AdData() {}

func (*Ad_ExpandedTextAd) isAd_AdData() {}

func (*Ad_CallOnlyAd) isAd_AdData() {}

func (*Ad_ExpandedDynamicSearchAd) isAd_AdData() {}

func (*Ad_HotelAd) isAd_AdData() {}

func (*Ad_ShoppingSmartAd) isAd_AdData() {}

func (*Ad_ShoppingProductAd) isAd_AdData() {}

func (*Ad_GmailAd) isAd_AdData() {}

func (*Ad_ImageAd) isAd_AdData() {}

func (*Ad_VideoAd) isAd_AdData() {}

func (*Ad_ResponsiveSearchAd) isAd_AdData() {}

func (*Ad_LegacyResponsiveDisplayAd) isAd_AdData() {}

func (m *Ad) GetAdData() isAd_AdData {
	if m != nil {
		return m.AdData
	}
	return nil
}

func (m *Ad) GetTextAd() *common.TextAdInfo {
	if x, ok := m.GetAdData().(*Ad_TextAd); ok {
		return x.TextAd
	}
	return nil
}

func (m *Ad) GetExpandedTextAd() *common.ExpandedTextAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedTextAd); ok {
		return x.ExpandedTextAd
	}
	return nil
}

func (m *Ad) GetCallOnlyAd() *common.CallOnlyAdInfo {
	if x, ok := m.GetAdData().(*Ad_CallOnlyAd); ok {
		return x.CallOnlyAd
	}
	return nil
}

func (m *Ad) GetExpandedDynamicSearchAd() *common.ExpandedDynamicSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ExpandedDynamicSearchAd); ok {
		return x.ExpandedDynamicSearchAd
	}
	return nil
}

func (m *Ad) GetHotelAd() *common.HotelAdInfo {
	if x, ok := m.GetAdData().(*Ad_HotelAd); ok {
		return x.HotelAd
	}
	return nil
}

func (m *Ad) GetShoppingSmartAd() *common.ShoppingSmartAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingSmartAd); ok {
		return x.ShoppingSmartAd
	}
	return nil
}

func (m *Ad) GetShoppingProductAd() *common.ShoppingProductAdInfo {
	if x, ok := m.GetAdData().(*Ad_ShoppingProductAd); ok {
		return x.ShoppingProductAd
	}
	return nil
}

func (m *Ad) GetGmailAd() *common.GmailAdInfo {
	if x, ok := m.GetAdData().(*Ad_GmailAd); ok {
		return x.GmailAd
	}
	return nil
}

func (m *Ad) GetImageAd() *common.ImageAdInfo {
	if x, ok := m.GetAdData().(*Ad_ImageAd); ok {
		return x.ImageAd
	}
	return nil
}

func (m *Ad) GetVideoAd() *common.VideoAdInfo {
	if x, ok := m.GetAdData().(*Ad_VideoAd); ok {
		return x.VideoAd
	}
	return nil
}

func (m *Ad) GetResponsiveSearchAd() *common.ResponsiveSearchAdInfo {
	if x, ok := m.GetAdData().(*Ad_ResponsiveSearchAd); ok {
		return x.ResponsiveSearchAd
	}
	return nil
}

func (m *Ad) GetLegacyResponsiveDisplayAd() *common.LegacyResponsiveDisplayAdInfo {
	if x, ok := m.GetAdData().(*Ad_LegacyResponsiveDisplayAd); ok {
		return x.LegacyResponsiveDisplayAd
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Ad) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Ad_TextAd)(nil),
		(*Ad_ExpandedTextAd)(nil),
		(*Ad_CallOnlyAd)(nil),
		(*Ad_ExpandedDynamicSearchAd)(nil),
		(*Ad_HotelAd)(nil),
		(*Ad_ShoppingSmartAd)(nil),
		(*Ad_ShoppingProductAd)(nil),
		(*Ad_GmailAd)(nil),
		(*Ad_ImageAd)(nil),
		(*Ad_VideoAd)(nil),
		(*Ad_ResponsiveSearchAd)(nil),
		(*Ad_LegacyResponsiveDisplayAd)(nil),
	}
}

func init() {
	proto.RegisterType((*Ad)(nil), "google.ads.googleads.v1.resources.Ad")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/ad.proto", fileDescriptor_2cfecfd7772c305a)
}

var fileDescriptor_2cfecfd7772c305a = []byte{
	// 1018 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x4e, 0xe4, 0x36,
	0x14, 0xc7, 0x3b, 0xb3, 0x14, 0xba, 0x66, 0xcb, 0x87, 0x59, 0xba, 0x59, 0x40, 0x2b, 0xb6, 0xd2,
	0x4a, 0x08, 0xd4, 0xcc, 0x02, 0xdd, 0xbd, 0x18, 0xb4, 0x52, 0xc3, 0x82, 0x80, 0xaa, 0xab, 0x8e,
	0x32, 0x30, 0x17, 0x2b, 0xda, 0xc8, 0xc4, 0x26, 0xa4, 0x75, 0xec, 0xc8, 0x76, 0xa6, 0xa4, 0x57,
	0x7d, 0x93, 0x4a, 0xbd, 0xec, 0x4d, 0xdf, 0xa3, 0x8f, 0xd2, 0xeb, 0x3e, 0x40, 0x65, 0x3b, 0x31,
	0x5f, 0x9d, 0xc9, 0xdc, 0x8c, 0xec, 0xf8, 0xff, 0xff, 0x9d, 0xe3, 0x93, 0x63, 0x4f, 0xc0, 0x66,
	0xc2, 0x79, 0x42, 0x49, 0x07, 0x61, 0xd9, 0xb1, 0x43, 0x3d, 0x1a, 0x6e, 0x77, 0x04, 0x91, 0xbc,
	0x10, 0x31, 0x91, 0x1d, 0x84, 0xfd, 0x5c, 0x70, 0xc5, 0xe1, 0x4b, 0x2b, 0xf0, 0x11, 0x96, 0xbe,
	0xd3, 0xfa, 0xc3, 0x6d, 0xdf, 0x69, 0x57, 0x76, 0x46, 0xe1, 0x62, 0x9e, 0x65, 0x9c, 0x75, 0x10,
	0x8e, 0x54, 0x99, 0x93, 0x28, 0x65, 0x97, 0x5c, 0x5a, 0xec, 0xca, 0x9b, 0x06, 0x4f, 0x5c, 0x48,
	0xc5, 0xb3, 0x28, 0x47, 0x02, 0x65, 0x44, 0x11, 0x51, 0xd9, 0x76, 0x1b, 0x6c, 0x85, 0xa0, 0x51,
	0xcc, 0x29, 0x25, 0xb1, 0x4a, 0x39, 0xab, 0x4c, 0x5b, 0xa3, 0x4c, 0x84, 0x15, 0x99, 0xac, 0xd3,
	0xab, 0xc4, 0x9b, 0xe3, 0xc5, 0x98, 0x0c, 0xd3, 0xb8, 0xd6, 0x7e, 0x33, 0x5e, 0x2b, 0x4b, 0xa9,
	0x48, 0x16, 0x65, 0x88, 0xa1, 0x84, 0xe0, 0x88, 0x30, 0x95, 0xaa, 0x32, 0xb2, 0x55, 0xab, 0x08,
	0x2f, 0x2a, 0x82, 0x99, 0x5d, 0x14, 0x97, 0x9d, 0x5f, 0x04, 0xca, 0x73, 0x22, 0xea, 0x32, 0xad,
	0xd5, 0x11, 0xf2, 0xb4, 0x83, 0x18, 0xe3, 0x0a, 0xe9, 0x7d, 0x55, 0xab, 0x5f, 0xfe, 0xb5, 0x00,
	0xda, 0x01, 0x86, 0x5b, 0xa0, 0x9d, 0x62, 0xaf, 0xb5, 0xde, 0xda, 0x98, 0xdd, 0x59, 0xad, 0x5e,
	0x92, 0x5f, 0x13, 0xfd, 0x13, 0xa6, 0xde, 0x7e, 0x3d, 0x40, 0xb4, 0x20, 0x61, 0x3b, 0xc5, 0x70,
	0x0f, 0x80, 0xcb, 0x94, 0x21, 0x1a, 0x15, 0x82, 0x4a, 0xaf, 0xbd, 0xfe, 0x68, 0x63, 0x76, 0x67,
	0xed, 0x81, 0xa9, 0xaf, 0x44, 0xca, 0x12, 0xeb, 0x7a, 0x6c, 0xf4, 0x67, 0x82, 0x4a, 0x78, 0x0c,
	0x16, 0xad, 0x39, 0xe3, 0x17, 0x29, 0x25, 0x96, 0xb1, 0x30, 0x01, 0x63, 0xde, 0xd8, 0x3e, 0x18,
	0x97, 0x21, 0xf5, 0xc0, 0xb2, 0x12, 0x28, 0xfe, 0x39, 0x65, 0x89, 0xa6, 0x44, 0x8a, 0x64, 0x39,
	0x45, 0x8a, 0x78, 0x4f, 0xcc, 0x36, 0xc6, 0xd3, 0x96, 0x6a, 0xeb, 0x99, 0xa0, 0xa7, 0x95, 0x11,
	0xc6, 0x60, 0xd9, 0xbc, 0xfd, 0x7b, 0x8d, 0x23, 0x3d, 0x60, 0xf2, 0xeb, 0xf8, 0xa3, 0x1a, 0xd9,
	0xb6, 0x8e, 0xff, 0xde, 0x18, 0x7b, 0xb5, 0x2f, 0x5c, 0x2a, 0x04, 0xbd, 0xf7, 0x4c, 0xc2, 0x77,
	0x60, 0x16, 0xa7, 0x32, 0xa7, 0xa8, 0xd4, 0x59, 0x7b, 0x53, 0x13, 0x24, 0x0b, 0x2a, 0xc3, 0x99,
	0xa0, 0xf0, 0x00, 0x4c, 0xe9, 0x56, 0xf3, 0x3e, 0x5d, 0x6f, 0x6d, 0xcc, 0xed, 0xbc, 0x1e, 0x99,
	0x92, 0xe9, 0x1f, 0x3f, 0xc0, 0xa7, 0x65, 0x4e, 0x0e, 0x59, 0x91, 0x55, 0xc3, 0xd0, 0xb8, 0xe1,
	0x09, 0x58, 0x42, 0x18, 0x13, 0x1c, 0x5d, 0x94, 0x91, 0xb5, 0x45, 0x08, 0x4b, 0x6f, 0xc9, 0x24,
	0xb3, 0xf2, 0x20, 0x99, 0x7d, 0xce, 0xa9, 0x4d, 0x65, 0xc1, 0xd8, 0xf6, 0xcb, 0x23, 0xa3, 0x08,
	0xb0, 0x84, 0x3f, 0x80, 0x45, 0xdb, 0xd1, 0x51, 0x2e, 0xc8, 0x25, 0x11, 0x84, 0xc5, 0xc4, 0x7b,
	0x3a, 0x51, 0x76, 0x07, 0xc6, 0x67, 0xb2, 0xb3, 0xc3, 0x70, 0xc1, 0xa2, 0x7a, 0x8e, 0x04, 0x07,
	0x60, 0xfe, 0xee, 0x89, 0x94, 0xde, 0x8a, 0x79, 0x1b, 0x5f, 0x35, 0xbd, 0x8d, 0x33, 0x41, 0xdf,
	0x3b, 0x57, 0x38, 0x57, 0xdc, 0x9e, 0x4a, 0xf8, 0x1a, 0x4c, 0x31, 0x94, 0x11, 0xef, 0xd9, 0x04,
	0xf5, 0x37, 0x4a, 0xf8, 0x7b, 0x0b, 0xbc, 0xb8, 0x77, 0x1e, 0xeb, 0x0b, 0xac, 0x3a, 0x91, 0xde,
	0xaa, 0xd9, 0xf6, 0xc7, 0x86, 0x6d, 0xf7, 0x0d, 0xe4, 0x83, 0x65, 0x84, 0x15, 0xa2, 0x6f, 0x7e,
	0x4d, 0x2d, 0xc6, 0xac, 0x87, 0xab, 0x72, 0xf4, 0x22, 0x3c, 0x04, 0x33, 0x8a, 0x5c, 0xab, 0x08,
	0x61, 0x6f, 0xda, 0x6c, 0x6b, 0xb3, 0xa9, 0x46, 0xa7, 0xe4, 0x5a, 0x05, 0xf8, 0x84, 0x5d, 0xf2,
	0xe3, 0x4f, 0xc2, 0x69, 0x65, 0x66, 0xf0, 0x47, 0xb0, 0x40, 0xae, 0x73, 0xc4, 0x74, 0x7f, 0xd4,
	0xbc, 0x19, 0xc3, 0xdb, 0x69, 0xe2, 0x1d, 0x56, 0xbe, 0x3b, 0xdc, 0x39, 0x72, 0xe7, 0x29, 0x0c,
	0xc1, 0x93, 0x18, 0x51, 0x1a, 0x71, 0x46, 0x4b, 0xcd, 0xfe, 0xdc, 0xb0, 0xfd, 0xc6, 0xd3, 0x85,
	0x28, 0xfd, 0x9e, 0xd1, 0xd2, 0x71, 0x41, 0xec, 0x9e, 0xc0, 0x5f, 0xc1, 0x8a, 0xcb, 0x19, 0x97,
	0x0c, 0x65, 0x69, 0x1c, 0x49, 0x82, 0x44, 0x7c, 0xa5, 0x23, 0xcc, 0x99, 0x08, 0x7b, 0x93, 0x66,
	0x7f, 0x60, 0x01, 0x7d, 0xe3, 0x77, 0xe1, 0x9e, 0x91, 0xff, 0x5f, 0x86, 0xc7, 0xe0, 0xb3, 0x2b,
	0xae, 0x08, 0xd5, 0x91, 0xe6, 0x4d, 0xa4, 0xad, 0xa6, 0x48, 0xc7, 0x5a, 0xef, 0xc8, 0x33, 0x57,
	0x76, 0x0a, 0x11, 0x58, 0x94, 0x57, 0x3c, 0xcf, 0xf5, 0x95, 0x26, 0x33, 0x24, 0x4c, 0xe9, 0x17,
	0x0d, 0x72, 0xb7, 0x09, 0xd9, 0xaf, 0x8c, 0x7d, 0xed, 0x73, 0xe8, 0x79, 0x79, 0xf7, 0x31, 0x4c,
	0xc0, 0x92, 0x0b, 0x91, 0x0b, 0x8e, 0x8b, 0xd8, 0x04, 0x81, 0x26, 0xc8, 0x9b, 0x49, 0x83, 0xf4,
	0xac, 0xd3, 0x85, 0x71, 0x69, 0xbb, 0x05, 0x5d, 0x95, 0x24, 0x43, 0xa9, 0xa9, 0xca, 0xf2, 0x64,
	0x55, 0x39, 0xd2, 0xfa, 0x9b, 0xaa, 0x24, 0x76, 0xaa, 0x49, 0x69, 0x86, 0x12, 0x7d, 0x45, 0x79,
	0x5f, 0x4c, 0x46, 0x3a, 0xd1, 0xfa, 0x1b, 0x52, 0x6a, 0xa7, 0x9a, 0x34, 0x4c, 0x31, 0xe1, 0x9a,
	0xe4, 0x4d, 0x46, 0x1a, 0x68, 0xfd, 0x0d, 0x69, 0x68, 0xa7, 0xf0, 0x27, 0xf0, 0x54, 0x10, 0x99,
	0x73, 0x26, 0xd3, 0x21, 0xb9, 0xd5, 0x69, 0xcf, 0x0d, 0xf5, 0x6d, 0x13, 0x35, 0x74, 0xde, 0x7b,
	0x4d, 0x06, 0xc5, 0x83, 0x15, 0xf8, 0x5b, 0x0b, 0xac, 0x51, 0x92, 0xa0, 0xb8, 0x8c, 0x6e, 0xc5,
	0xac, 0xff, 0x44, 0x10, 0xf6, 0xd6, 0x4c, 0xd0, 0x77, 0x4d, 0x41, 0xbf, 0x33, 0x8c, 0x9b, 0xd0,
	0x07, 0x96, 0xe0, 0x62, 0x3f, 0xa7, 0xa3, 0x04, 0xfb, 0x8f, 0xc1, 0x0c, 0xc2, 0x11, 0x46, 0x0a,
	0xed, 0xff, 0xdb, 0x02, 0xaf, 0x62, 0x9e, 0xf9, 0x8d, 0x1f, 0x75, 0xfb, 0x33, 0x01, 0xee, 0xe9,
	0xeb, 0xb4, 0xd7, 0xfa, 0xf8, 0x6d, 0xa5, 0x4e, 0x38, 0x45, 0x2c, 0xf1, 0xb9, 0x48, 0x3a, 0x09,
	0x61, 0xe6, 0xb2, 0xad, 0x3f, 0x7b, 0xf2, 0x54, 0x8e, 0xf9, 0x9a, 0xdc, 0x73, 0xa3, 0x3f, 0xda,
	0x8f, 0x8e, 0x82, 0xe0, 0xcf, 0xf6, 0x4b, 0xfb, 0x17, 0xe4, 0x07, 0x58, 0xfa, 0xee, 0xdf, 0xc8,
	0x1f, 0x6c, 0xfb, 0xf5, 0xa5, 0x28, 0xff, 0xae, 0x35, 0xe7, 0x01, 0x96, 0xe7, 0x4e, 0x73, 0x3e,
	0xd8, 0x3e, 0x77, 0x9a, 0x7f, 0xda, 0xaf, 0xec, 0x42, 0xb7, 0x1b, 0x60, 0xd9, 0xed, 0x3a, 0x55,
	0xb7, 0x3b, 0xd8, 0xee, 0x76, 0x9d, 0xee, 0x62, 0xda, 0x24, 0xbb, 0xfb, 0x5f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xa3, 0x26, 0x2f, 0x66, 0xf9, 0x0a, 0x00, 0x00,
}
