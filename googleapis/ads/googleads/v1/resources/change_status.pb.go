// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/change_status.proto

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

// Describes the status of returned resource.
type ChangeStatus struct {
	// The resource name of the change status.
	// Change status resource names have the form:
	//
	// `customers/{customer_id}/changeStatus/{change_status_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Time at which the most recent change has occurred on this resource.
	LastChangeDateTime *wrappers.StringValue `protobuf:"bytes,3,opt,name=last_change_date_time,json=lastChangeDateTime,proto3" json:"last_change_date_time,omitempty"`
	// Represents the type of the changed resource. This dictates what fields
	// will be set. For example, for AD_GROUP, campaign and ad_group fields will
	// be set.
	ResourceType enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType `protobuf:"varint,4,opt,name=resource_type,json=resourceType,proto3,enum=google.ads.googleads.v1.enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType" json:"resource_type,omitempty"`
	// The Campaign affected by this change.
	Campaign *wrappers.StringValue `protobuf:"bytes,5,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The AdGroup affected by this change.
	AdGroup *wrappers.StringValue `protobuf:"bytes,6,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Represents the status of the changed resource.
	ResourceStatus enums.ChangeStatusOperationEnum_ChangeStatusOperation `protobuf:"varint,8,opt,name=resource_status,json=resourceStatus,proto3,enum=google.ads.googleads.v1.enums.ChangeStatusOperationEnum_ChangeStatusOperation" json:"resource_status,omitempty"`
	// The AdGroupAd affected by this change.
	AdGroupAd *wrappers.StringValue `protobuf:"bytes,9,opt,name=ad_group_ad,json=adGroupAd,proto3" json:"ad_group_ad,omitempty"`
	// The AdGroupCriterion affected by this change.
	AdGroupCriterion *wrappers.StringValue `protobuf:"bytes,10,opt,name=ad_group_criterion,json=adGroupCriterion,proto3" json:"ad_group_criterion,omitempty"`
	// The CampaignCriterion affected by this change.
	CampaignCriterion *wrappers.StringValue `protobuf:"bytes,11,opt,name=campaign_criterion,json=campaignCriterion,proto3" json:"campaign_criterion,omitempty"`
	// The Feed affected by this change.
	Feed *wrappers.StringValue `protobuf:"bytes,12,opt,name=feed,proto3" json:"feed,omitempty"`
	// The FeedItem affected by this change.
	FeedItem *wrappers.StringValue `protobuf:"bytes,13,opt,name=feed_item,json=feedItem,proto3" json:"feed_item,omitempty"`
	// The AdGroupFeed affected by this change.
	AdGroupFeed *wrappers.StringValue `protobuf:"bytes,14,opt,name=ad_group_feed,json=adGroupFeed,proto3" json:"ad_group_feed,omitempty"`
	// The CampaignFeed affected by this change.
	CampaignFeed *wrappers.StringValue `protobuf:"bytes,15,opt,name=campaign_feed,json=campaignFeed,proto3" json:"campaign_feed,omitempty"`
	// The AdGroupBidModifier affected by this change.
	AdGroupBidModifier   *wrappers.StringValue `protobuf:"bytes,16,opt,name=ad_group_bid_modifier,json=adGroupBidModifier,proto3" json:"ad_group_bid_modifier,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ChangeStatus) Reset()         { *m = ChangeStatus{} }
func (m *ChangeStatus) String() string { return proto.CompactTextString(m) }
func (*ChangeStatus) ProtoMessage()    {}
func (*ChangeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a34c24bda75f035, []int{0}
}

func (m *ChangeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeStatus.Unmarshal(m, b)
}
func (m *ChangeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeStatus.Marshal(b, m, deterministic)
}
func (m *ChangeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeStatus.Merge(m, src)
}
func (m *ChangeStatus) XXX_Size() int {
	return xxx_messageInfo_ChangeStatus.Size(m)
}
func (m *ChangeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeStatus proto.InternalMessageInfo

func (m *ChangeStatus) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *ChangeStatus) GetLastChangeDateTime() *wrappers.StringValue {
	if m != nil {
		return m.LastChangeDateTime
	}
	return nil
}

func (m *ChangeStatus) GetResourceType() enums.ChangeStatusResourceTypeEnum_ChangeStatusResourceType {
	if m != nil {
		return m.ResourceType
	}
	return enums.ChangeStatusResourceTypeEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *ChangeStatus) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *ChangeStatus) GetResourceStatus() enums.ChangeStatusOperationEnum_ChangeStatusOperation {
	if m != nil {
		return m.ResourceStatus
	}
	return enums.ChangeStatusOperationEnum_UNSPECIFIED
}

func (m *ChangeStatus) GetAdGroupAd() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupAd
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupCriterion() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupCriterion
	}
	return nil
}

func (m *ChangeStatus) GetCampaignCriterion() *wrappers.StringValue {
	if m != nil {
		return m.CampaignCriterion
	}
	return nil
}

func (m *ChangeStatus) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *ChangeStatus) GetFeedItem() *wrappers.StringValue {
	if m != nil {
		return m.FeedItem
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupFeed() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupFeed
	}
	return nil
}

func (m *ChangeStatus) GetCampaignFeed() *wrappers.StringValue {
	if m != nil {
		return m.CampaignFeed
	}
	return nil
}

func (m *ChangeStatus) GetAdGroupBidModifier() *wrappers.StringValue {
	if m != nil {
		return m.AdGroupBidModifier
	}
	return nil
}

func init() {
	proto.RegisterType((*ChangeStatus)(nil), "google.ads.googleads.v1.resources.ChangeStatus")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/change_status.proto", fileDescriptor_9a34c24bda75f035)
}

var fileDescriptor_9a34c24bda75f035 = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6a, 0xdb, 0x3e,
	0x14, 0xc7, 0x6d, 0xff, 0xfd, 0x27, 0x4a, 0xd2, 0x0f, 0x41, 0xc1, 0x94, 0x32, 0xda, 0x8d, 0x42,
	0xaf, 0xe4, 0xa5, 0x63, 0x6c, 0x73, 0x07, 0x9b, 0xdb, 0x6d, 0x65, 0x1d, 0x6b, 0x4b, 0x5a, 0x72,
	0x31, 0x02, 0x46, 0x8d, 0x54, 0x4f, 0x10, 0x4b, 0x46, 0x92, 0x5b, 0xf2, 0x00, 0x7b, 0x91, 0x5d,
	0xee, 0x51, 0xf6, 0x28, 0x7b, 0x88, 0x31, 0x2c, 0x4b, 0x6a, 0xd8, 0xc8, 0xe6, 0x5e, 0xe5, 0x38,
	0xe7, 0xf7, 0x75, 0x8e, 0x6c, 0x81, 0xa7, 0x99, 0x10, 0xd9, 0x84, 0x46, 0x98, 0xa8, 0xa8, 0x2e,
	0xab, 0xea, 0xa6, 0x1f, 0x49, 0xaa, 0x44, 0x29, 0xc7, 0x54, 0x45, 0xe3, 0xcf, 0x98, 0x67, 0x34,
	0x55, 0x1a, 0xeb, 0x52, 0xa1, 0x42, 0x0a, 0x2d, 0xe0, 0x4e, 0x8d, 0x45, 0x98, 0x28, 0xe4, 0x69,
	0xe8, 0xa6, 0x8f, 0x3c, 0x6d, 0xf3, 0x60, 0x9e, 0x32, 0xe5, 0x65, 0xfe, 0x9b, 0x6a, 0x2a, 0x0a,
	0x2a, 0xb1, 0x66, 0x82, 0xd7, 0xfa, 0x9b, 0xaf, 0xee, 0x43, 0x76, 0x9e, 0xa9, 0x9e, 0x16, 0xd4,
	0x0a, 0x3c, 0xb0, 0x02, 0xe6, 0xe9, 0xaa, 0xbc, 0x8e, 0x6e, 0x25, 0x2e, 0x0a, 0x2a, 0xed, 0x00,
	0x9b, 0x5b, 0xce, 0xa0, 0x60, 0x11, 0xe6, 0x5c, 0x68, 0xe3, 0x6e, 0xbb, 0x0f, 0xbf, 0xb4, 0x40,
	0xf7, 0xc8, 0x78, 0x5c, 0x18, 0x0b, 0xf8, 0x08, 0xf4, 0xbc, 0x0b, 0xc7, 0x39, 0x0d, 0x83, 0xed,
	0x60, 0xaf, 0x3d, 0xe8, 0xba, 0x3f, 0x4f, 0x71, 0x4e, 0xe1, 0x19, 0xd8, 0x98, 0x60, 0xa5, 0x53,
	0x9b, 0x8e, 0x60, 0x4d, 0x53, 0xcd, 0x72, 0x1a, 0x2e, 0x6e, 0x07, 0x7b, 0x9d, 0xfd, 0x2d, 0xbb,
	0x29, 0xe4, 0x32, 0xa1, 0x0b, 0x2d, 0x19, 0xcf, 0x86, 0x78, 0x52, 0xd2, 0x01, 0xac, 0xa8, 0xb5,
	0xe7, 0x1b, 0xac, 0xe9, 0x25, 0xcb, 0x29, 0x9c, 0xce, 0xb8, 0x56, 0xb3, 0x85, 0x4b, 0xdb, 0xc1,
	0xde, 0xca, 0xfe, 0x25, 0x9a, 0xb7, 0x7d, 0xb3, 0x1d, 0x34, 0x9b, 0x7c, 0x60, 0xf9, 0x97, 0xd3,
	0x82, 0xbe, 0xe5, 0x65, 0x3e, 0xb7, 0x79, 0x37, 0x4b, 0xf5, 0x04, 0x9f, 0x83, 0xd6, 0x18, 0xe7,
	0x05, 0x66, 0x19, 0x0f, 0xff, 0x6b, 0x10, 0xdf, 0xa3, 0xe1, 0x33, 0xd0, 0xc2, 0x24, 0xcd, 0xa4,
	0x28, 0x8b, 0x70, 0xb9, 0x01, 0xf3, 0x7f, 0x4c, 0x8e, 0x2b, 0x30, 0xbc, 0x05, 0xab, 0x7e, 0xda,
	0xfa, 0x64, 0xc3, 0x96, 0x99, 0xf7, 0xf4, 0x1e, 0xf3, 0x9e, 0xb9, 0x17, 0xe9, 0x8f, 0x61, 0x7d,
	0x67, 0xb0, 0xe2, 0x6c, 0xec, 0xe1, 0xbe, 0x04, 0x1d, 0x97, 0x38, 0xc5, 0x24, 0x6c, 0x37, 0x08,
	0xdd, 0xb6, 0xa1, 0x13, 0x02, 0x4f, 0x00, 0xf4, 0xec, 0xb1, 0x64, 0x9a, 0x4a, 0x26, 0x78, 0x08,
	0x1a, 0x88, 0xac, 0x59, 0x91, 0x23, 0xc7, 0x82, 0x1f, 0x00, 0x74, 0x7b, 0x9c, 0xd1, 0xea, 0x34,
	0xd0, 0x5a, 0x77, 0xbc, 0x3b, 0xb1, 0xc7, 0x60, 0xe9, 0x9a, 0x52, 0x12, 0x76, 0x1b, 0xd0, 0x0d,
	0x12, 0xbe, 0x00, 0xed, 0xea, 0x37, 0x65, 0x9a, 0xe6, 0x61, 0xaf, 0xc9, 0xa9, 0x57, 0xf0, 0xf7,
	0x9a, 0xe6, 0xf0, 0x35, 0xe8, 0xf9, 0x2d, 0x18, 0xd7, 0x95, 0x06, 0xf4, 0x8e, 0x5d, 0xc0, 0xbb,
	0xca, 0x3c, 0x01, 0x3d, 0x3f, 0xbb, 0x51, 0x58, 0x6d, 0xa0, 0xd0, 0x75, 0x14, 0x23, 0x71, 0x06,
	0x36, 0x7c, 0x88, 0x2b, 0x46, 0xd2, 0x5c, 0x10, 0x76, 0xcd, 0xa8, 0x0c, 0xd7, 0x9a, 0x7c, 0x80,
	0x36, 0xcc, 0x21, 0x23, 0x1f, 0x2d, 0xef, 0xf0, 0x67, 0x00, 0x76, 0xc7, 0x22, 0x47, 0xff, 0xbc,
	0xed, 0x0e, 0xd7, 0x67, 0x5f, 0xb5, 0xf3, 0x4a, 0xff, 0x3c, 0xf8, 0x74, 0x62, 0x79, 0x99, 0x98,
	0x60, 0x9e, 0x21, 0x21, 0xb3, 0x28, 0xa3, 0xdc, 0xb8, 0xbb, 0x5b, 0xad, 0x60, 0xea, 0x2f, 0x77,
	0xef, 0x81, 0xaf, 0xbe, 0x2e, 0x2c, 0x1e, 0x27, 0xc9, 0xb7, 0x85, 0x9d, 0xe3, 0x5a, 0x32, 0x21,
	0x0a, 0xd5, 0x65, 0x55, 0x0d, 0xfb, 0xc8, 0x7d, 0xcb, 0xea, 0xbb, 0xc3, 0x8c, 0x12, 0xa2, 0x46,
	0x1e, 0x33, 0x1a, 0xf6, 0x47, 0x1e, 0xf3, 0x63, 0x61, 0xb7, 0x6e, 0xc4, 0x71, 0x42, 0x54, 0x1c,
	0x7b, 0x54, 0x1c, 0x0f, 0xfb, 0x71, 0xec, 0x71, 0x57, 0xcb, 0x26, 0xec, 0x93, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xe0, 0xa2, 0x09, 0x27, 0x27, 0x06, 0x00, 0x00,
}
