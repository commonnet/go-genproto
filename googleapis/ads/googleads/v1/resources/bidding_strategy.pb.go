// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/bidding_strategy.proto

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

// A bidding strategy.
type BiddingStrategy struct {
	// The resource name of the bidding strategy.
	// Bidding strategy resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the bidding strategy.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the bidding strategy.
	// All bidding strategies within an account must be named distinctly.
	//
	// The length of this string should be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the bidding strategy.
	//
	// This field is read-only.
	Status enums.BiddingStrategyStatusEnum_BiddingStrategyStatus `protobuf:"varint,15,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.BiddingStrategyStatusEnum_BiddingStrategyStatus" json:"status,omitempty"`
	// The type of the bidding strategy.
	// Create a bidding strategy by setting the bidding scheme.
	//
	// This field is read-only.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v1.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// The number of campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	CampaignCount *wrappers.Int64Value `protobuf:"bytes,13,opt,name=campaign_count,json=campaignCount,proto3" json:"campaign_count,omitempty"`
	// The number of non-removed campaigns attached to this bidding strategy.
	//
	// This field is read-only.
	NonRemovedCampaignCount *wrappers.Int64Value `protobuf:"bytes,14,opt,name=non_removed_campaign_count,json=nonRemovedCampaignCount,proto3" json:"non_removed_campaign_count,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are valid to be assigned to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_PageOnePromoted
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetImpressionShare
	//	*BiddingStrategy_TargetOutrankShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme               isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BiddingStrategy) Reset()         { *m = BiddingStrategy{} }
func (m *BiddingStrategy) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategy) ProtoMessage()    {}
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_45db3d8386a3943d, []int{0}
}

func (m *BiddingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategy.Unmarshal(m, b)
}
func (m *BiddingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategy.Marshal(b, m, deterministic)
}
func (m *BiddingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategy.Merge(m, src)
}
func (m *BiddingStrategy) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategy.Size(m)
}
func (m *BiddingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategy proto.InternalMessageInfo

func (m *BiddingStrategy) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BiddingStrategy) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BiddingStrategy) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BiddingStrategy) GetStatus() enums.BiddingStrategyStatusEnum_BiddingStrategyStatus {
	if m != nil {
		return m.Status
	}
	return enums.BiddingStrategyStatusEnum_UNSPECIFIED
}

func (m *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if m != nil {
		return m.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

func (m *BiddingStrategy) GetCampaignCount() *wrappers.Int64Value {
	if m != nil {
		return m.CampaignCount
	}
	return nil
}

func (m *BiddingStrategy) GetNonRemovedCampaignCount() *wrappers.Int64Value {
	if m != nil {
		return m.NonRemovedCampaignCount
	}
	return nil
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_PageOnePromoted struct {
	PageOnePromoted *common.PageOnePromoted `protobuf:"bytes,8,opt,name=page_one_promoted,json=pageOnePromoted,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetImpressionShare struct {
	TargetImpressionShare *common.TargetImpressionShare `protobuf:"bytes,48,opt,name=target_impression_share,json=targetImpressionShare,proto3,oneof"`
}

type BiddingStrategy_TargetOutrankShare struct {
	TargetOutrankShare *common.TargetOutrankShare `protobuf:"bytes,10,opt,name=target_outrank_share,json=targetOutrankShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_PageOnePromoted) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetImpressionShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetOutrankShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (m *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := m.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (m *BiddingStrategy) GetPageOnePromoted() *common.PageOnePromoted {
	if x, ok := m.GetScheme().(*BiddingStrategy_PageOnePromoted); ok {
		return x.PageOnePromoted
	}
	return nil
}

func (m *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (m *BiddingStrategy) GetTargetImpressionShare() *common.TargetImpressionShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetImpressionShare); ok {
		return x.TargetImpressionShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetOutrankShare() *common.TargetOutrankShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetOutrankShare); ok {
		return x.TargetOutrankShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (m *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BiddingStrategy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_PageOnePromoted)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetImpressionShare)(nil),
		(*BiddingStrategy_TargetOutrankShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
}

func init() {
	proto.RegisterType((*BiddingStrategy)(nil), "google.ads.googleads.v1.resources.BiddingStrategy")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/bidding_strategy.proto", fileDescriptor_45db3d8386a3943d)
}

var fileDescriptor_45db3d8386a3943d = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdd, 0x4e, 0xdb, 0x30,
	0x14, 0xc7, 0xdb, 0xc2, 0x18, 0xb8, 0x7c, 0x68, 0x11, 0x13, 0x11, 0x43, 0x13, 0x6c, 0x42, 0x62,
	0x63, 0x72, 0x28, 0xfb, 0xd0, 0x16, 0xae, 0xda, 0x0a, 0x51, 0x90, 0x06, 0x55, 0x8a, 0xaa, 0x69,
	0xea, 0x16, 0x99, 0xe4, 0x10, 0xa2, 0x11, 0xdb, 0xb2, 0x1d, 0x26, 0x2e, 0xf7, 0x2a, 0xbb, 0xdc,
	0xa3, 0xec, 0x51, 0xf6, 0x0c, 0xbb, 0x98, 0xe2, 0x38, 0x41, 0x7c, 0x74, 0x2d, 0x77, 0x3e, 0xf6,
	0xff, 0xff, 0xfb, 0xbb, 0xa7, 0x27, 0x09, 0x7a, 0x1f, 0x31, 0x16, 0x9d, 0x83, 0x43, 0x42, 0xe9,
	0xe4, 0xcb, 0x6c, 0x75, 0xd1, 0x70, 0x04, 0x48, 0x96, 0x8a, 0x00, 0xa4, 0x73, 0x12, 0x87, 0x61,
	0x4c, 0x23, 0x5f, 0x2a, 0x41, 0x14, 0x44, 0x97, 0x98, 0x0b, 0xa6, 0x98, 0xb5, 0x96, 0xcb, 0x31,
	0x09, 0x25, 0x2e, 0x9d, 0xf8, 0xa2, 0x81, 0x4b, 0xe7, 0xf2, 0xab, 0x61, 0xf0, 0x80, 0x25, 0x09,
	0xa3, 0x05, 0x39, 0x07, 0x2e, 0xef, 0x0c, 0x53, 0x03, 0x4d, 0x93, 0xdb, 0xd7, 0xf0, 0xa5, 0x22,
	0x2a, 0x95, 0xc6, 0xfc, 0xe1, 0x9e, 0x66, 0x75, 0xc9, 0xc1, 0x58, 0x9f, 0x1a, 0xab, 0xae, 0x4e,
	0xd2, 0x53, 0xe7, 0xbb, 0x20, 0x9c, 0x83, 0x28, 0xd0, 0x2b, 0x05, 0x9a, 0xc7, 0x0e, 0xa1, 0x94,
	0x29, 0xa2, 0x62, 0x46, 0xcd, 0xe9, 0xb3, 0xbf, 0xd3, 0x68, 0xa1, 0x95, 0xd3, 0x7b, 0x06, 0x6e,
	0x3d, 0x47, 0x73, 0x45, 0x13, 0x7c, 0x4a, 0x12, 0xb0, 0xab, 0xab, 0xd5, 0x8d, 0x19, 0x6f, 0xb6,
	0xd8, 0x3c, 0x24, 0x09, 0x58, 0x9b, 0xa8, 0x16, 0x87, 0xf6, 0xc4, 0x6a, 0x75, 0xa3, 0xbe, 0xfd,
	0xc4, 0x74, 0x10, 0x17, 0x77, 0xc0, 0xfb, 0x54, 0xbd, 0x7b, 0xd3, 0x27, 0xe7, 0x29, 0x78, 0xb5,
	0x38, 0xb4, 0xb6, 0xd0, 0xa4, 0x06, 0x4d, 0x6a, 0xf9, 0xca, 0x2d, 0x79, 0x4f, 0x89, 0x98, 0x46,
	0xb9, 0x5e, 0x2b, 0xad, 0x53, 0x34, 0x95, 0x37, 0xc8, 0x5e, 0x58, 0xad, 0x6e, 0xcc, 0x6f, 0x1f,
	0xe2, 0x61, 0xff, 0x97, 0xee, 0x10, 0xbe, 0xf1, 0x1b, 0x7a, 0xda, 0xbb, 0x4b, 0xd3, 0xe4, 0xee,
	0x13, 0xcf, 0xd0, 0xad, 0xaf, 0x68, 0x32, 0xeb, 0xa5, 0xfd, 0x40, 0xa7, 0x1c, 0xdc, 0x2f, 0xe5,
	0xf8, 0x92, 0xc3, 0x5d, 0x19, 0xd9, 0xbe, 0xa7, 0xb9, 0x56, 0x0b, 0xcd, 0x07, 0x24, 0xe1, 0x24,
	0x8e, 0xa8, 0x1f, 0xb0, 0x94, 0x2a, 0x7b, 0x6e, 0x74, 0xcb, 0xe6, 0x0a, 0x4b, 0x3b, 0x73, 0x58,
	0x9f, 0xd0, 0x32, 0x65, 0xd4, 0x17, 0x90, 0xb0, 0x0b, 0x08, 0xfd, 0x1b, 0xbc, 0xf9, 0xd1, 0xbc,
	0x25, 0xca, 0xa8, 0x97, 0xbb, 0xdb, 0xd7, 0xc8, 0x5d, 0x34, 0x0b, 0xf4, 0x8c, 0xd0, 0x20, 0xc3,
	0xf2, 0xc0, 0x7e, 0xa8, 0x59, 0x9b, 0x43, 0xbb, 0x90, 0x0f, 0x3e, 0xde, 0x35, 0x9e, 0x36, 0x0f,
	0x3a, 0x15, 0xaf, 0x0e, 0x57, 0xa5, 0xf5, 0x05, 0x3d, 0xe2, 0x24, 0x02, 0x9f, 0x51, 0xf0, 0xb9,
	0x60, 0x09, 0x53, 0x10, 0xda, 0xd3, 0x1a, 0xeb, 0x8c, 0xc2, 0x76, 0x49, 0x04, 0x47, 0x14, 0xba,
	0xc6, 0xd6, 0xa9, 0x78, 0x0b, 0xfc, 0xfa, 0x96, 0x75, 0x80, 0x90, 0x22, 0x22, 0x02, 0xe5, 0x07,
	0x9c, 0xd8, 0x33, 0x9a, 0xfb, 0x62, 0x14, 0xf7, 0x58, 0x3b, 0xda, 0x9c, 0x74, 0x2a, 0xde, 0x8c,
	0x2a, 0x0a, 0x8b, 0xa1, 0x25, 0xc3, 0x8a, 0x13, 0x2e, 0x40, 0xca, 0x98, 0x51, 0x5f, 0x9e, 0x11,
	0x01, 0xf6, 0x96, 0x06, 0xbf, 0x1d, 0x0f, 0xbc, 0x5f, 0xba, 0x7b, 0x99, 0xb9, 0x53, 0xf1, 0x1e,
	0xab, 0xbb, 0x0e, 0xac, 0x53, 0xb4, 0x68, 0x02, 0x59, 0xaa, 0x04, 0xa1, 0xdf, 0x4c, 0x1a, 0xd2,
	0x69, 0xdb, 0xe3, 0xa5, 0x1d, 0xe5, 0xd6, 0x22, 0xca, 0x52, 0xb7, 0x76, 0xad, 0x8f, 0xa8, 0x6e,
	0x72, 0x04, 0x23, 0xd2, 0xae, 0x6b, 0xfc, 0xcb, 0xf1, 0xf0, 0x1e, 0x23, 0xb2, 0x53, 0xf1, 0x4c,
	0x97, 0xb3, 0x2a, 0x1b, 0x12, 0x83, 0x93, 0x1c, 0x68, 0x68, 0xcf, 0x8e, 0x37, 0x24, 0x39, 0xaf,
	0x97, 0x59, 0xb2, 0x21, 0x51, 0x57, 0x65, 0x6b, 0x1a, 0x4d, 0xc9, 0xe0, 0x0c, 0x12, 0x68, 0xfd,
	0xa8, 0xa1, 0xf5, 0x80, 0x25, 0x78, 0xe4, 0xcb, 0xb8, 0xb5, 0x78, 0xe3, 0x19, 0xeb, 0x66, 0x73,
	0xde, 0xad, 0x7e, 0x3e, 0x30, 0xd6, 0x88, 0x9d, 0x13, 0x1a, 0x61, 0x26, 0x22, 0x27, 0x02, 0xaa,
	0x9f, 0x82, 0xe2, 0x4d, 0xca, 0x63, 0xf9, 0x9f, 0x0f, 0xc4, 0x4e, 0xb9, 0xfa, 0x59, 0x9b, 0xd8,
	0x6b, 0x36, 0x7f, 0xd5, 0xd6, 0xf6, 0x72, 0x64, 0x33, 0x94, 0x38, 0x5f, 0x66, 0xab, 0x7e, 0x03,
	0x7b, 0x85, 0xf2, 0x77, 0xa1, 0x19, 0x34, 0x43, 0x39, 0x28, 0x35, 0x83, 0x7e, 0x63, 0x50, 0x6a,
	0xfe, 0xd4, 0xd6, 0xf3, 0x03, 0xd7, 0x6d, 0x86, 0xd2, 0x75, 0x4b, 0x95, 0xeb, 0xf6, 0x1b, 0xae,
	0x5b, 0xea, 0x4e, 0xa6, 0xf4, 0x65, 0x5f, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x35, 0xc9, 0x5d,
	0xf0, 0xcc, 0x06, 0x00, 0x00,
}
