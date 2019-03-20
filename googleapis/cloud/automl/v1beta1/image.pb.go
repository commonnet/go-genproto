// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/image.proto

package automl

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Dataset metadata that is specific to image classification.
type ImageClassificationDatasetMetadata struct {
	// Required.
	// Type of the classification problem.
	ClassificationType   ClassificationType `protobuf:"varint,1,opt,name=classification_type,json=classificationType,proto3,enum=google.cloud.automl.v1beta1.ClassificationType" json:"classification_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ImageClassificationDatasetMetadata) Reset()         { *m = ImageClassificationDatasetMetadata{} }
func (m *ImageClassificationDatasetMetadata) String() string { return proto.CompactTextString(m) }
func (*ImageClassificationDatasetMetadata) ProtoMessage()    {}
func (*ImageClassificationDatasetMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b9f2bc900da869, []int{0}
}

func (m *ImageClassificationDatasetMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageClassificationDatasetMetadata.Unmarshal(m, b)
}
func (m *ImageClassificationDatasetMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageClassificationDatasetMetadata.Marshal(b, m, deterministic)
}
func (m *ImageClassificationDatasetMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageClassificationDatasetMetadata.Merge(m, src)
}
func (m *ImageClassificationDatasetMetadata) XXX_Size() int {
	return xxx_messageInfo_ImageClassificationDatasetMetadata.Size(m)
}
func (m *ImageClassificationDatasetMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageClassificationDatasetMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ImageClassificationDatasetMetadata proto.InternalMessageInfo

func (m *ImageClassificationDatasetMetadata) GetClassificationType() ClassificationType {
	if m != nil {
		return m.ClassificationType
	}
	return ClassificationType_CLASSIFICATION_TYPE_UNSPECIFIED
}

// Model metadata for image classification.
type ImageClassificationModelMetadata struct {
	// Optional. The ID of the `base` model. If it is specified, the new model
	// will be created based on the `base` model. Otherwise, the new model will be
	// created from scratch. The `base` model is expected to be in the same
	// `project` and `location` as the new model to create.
	BaseModelId string `protobuf:"bytes,1,opt,name=base_model_id,json=baseModelId,proto3" json:"base_model_id,omitempty"`
	// Required. The train budget of creating this model. The actual
	// `train_cost` will be equal or less than this value.
	TrainBudget int64 `protobuf:"varint,2,opt,name=train_budget,json=trainBudget,proto3" json:"train_budget,omitempty"`
	// Output only. The actual train cost of creating this model. If this
	// model is created from a `base` model, the train cost used to create the
	// `base` model are not included.
	TrainCost int64 `protobuf:"varint,3,opt,name=train_cost,json=trainCost,proto3" json:"train_cost,omitempty"`
	// Output only. The reason that this create model operation stopped,
	// e.g. BUDGET_REACHED, CONVERGED.
	StopReason           string   `protobuf:"bytes,5,opt,name=stop_reason,json=stopReason,proto3" json:"stop_reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageClassificationModelMetadata) Reset()         { *m = ImageClassificationModelMetadata{} }
func (m *ImageClassificationModelMetadata) String() string { return proto.CompactTextString(m) }
func (*ImageClassificationModelMetadata) ProtoMessage()    {}
func (*ImageClassificationModelMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_29b9f2bc900da869, []int{1}
}

func (m *ImageClassificationModelMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageClassificationModelMetadata.Unmarshal(m, b)
}
func (m *ImageClassificationModelMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageClassificationModelMetadata.Marshal(b, m, deterministic)
}
func (m *ImageClassificationModelMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageClassificationModelMetadata.Merge(m, src)
}
func (m *ImageClassificationModelMetadata) XXX_Size() int {
	return xxx_messageInfo_ImageClassificationModelMetadata.Size(m)
}
func (m *ImageClassificationModelMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageClassificationModelMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ImageClassificationModelMetadata proto.InternalMessageInfo

func (m *ImageClassificationModelMetadata) GetBaseModelId() string {
	if m != nil {
		return m.BaseModelId
	}
	return ""
}

func (m *ImageClassificationModelMetadata) GetTrainBudget() int64 {
	if m != nil {
		return m.TrainBudget
	}
	return 0
}

func (m *ImageClassificationModelMetadata) GetTrainCost() int64 {
	if m != nil {
		return m.TrainCost
	}
	return 0
}

func (m *ImageClassificationModelMetadata) GetStopReason() string {
	if m != nil {
		return m.StopReason
	}
	return ""
}

func init() {
	proto.RegisterType((*ImageClassificationDatasetMetadata)(nil), "google.cloud.automl.v1beta1.ImageClassificationDatasetMetadata")
	proto.RegisterType((*ImageClassificationModelMetadata)(nil), "google.cloud.automl.v1beta1.ImageClassificationModelMetadata")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/image.proto", fileDescriptor_29b9f2bc900da869)
}

var fileDescriptor_29b9f2bc900da869 = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0x8b, 0x42, 0xb7, 0xea, 0x21, 0x5e, 0x42, 0xab, 0xb4, 0xe6, 0x62, 0x4f, 0x59,
	0xab, 0x47, 0x4f, 0x6d, 0x05, 0xe9, 0xa1, 0x50, 0x82, 0x78, 0x90, 0x42, 0x9c, 0x24, 0xdb, 0x65,
	0x21, 0xc9, 0x84, 0xec, 0x44, 0xe8, 0x0b, 0x78, 0xf6, 0x35, 0x7c, 0x15, 0x9f, 0x4a, 0xb2, 0x1b,
	0x84, 0x62, 0xe9, 0x71, 0xbf, 0x7c, 0xff, 0xec, 0x3f, 0x59, 0x76, 0x2b, 0x11, 0x65, 0x26, 0x78,
	0x92, 0x61, 0x9d, 0x72, 0xa8, 0x09, 0xf3, 0x8c, 0x7f, 0x4c, 0x63, 0x41, 0x30, 0xe5, 0x2a, 0x07,
	0x29, 0x82, 0xb2, 0x42, 0x42, 0x77, 0x68, 0xc5, 0xc0, 0x88, 0x81, 0x15, 0x83, 0x56, 0x1c, 0x5c,
	0xb5, 0x53, 0xa0, 0x54, 0x1c, 0x8a, 0x02, 0x09, 0x48, 0x61, 0xa1, 0x6d, 0x74, 0x70, 0x77, 0xec,
	0x8e, 0x24, 0x03, 0xad, 0xd5, 0x56, 0x25, 0x26, 0xd2, 0x26, 0x46, 0x6d, 0xc2, 0x9c, 0xe2, 0x7a,
	0xcb, 0x49, 0xe5, 0x42, 0x13, 0xe4, 0xa5, 0x15, 0xfc, 0x4f, 0x87, 0xf9, 0xcb, 0xa6, 0xdd, 0x62,
	0x2f, 0xfe, 0x04, 0x04, 0x5a, 0xd0, 0x4a, 0x10, 0xa4, 0x40, 0xe0, 0xbe, 0xb3, 0xcb, 0xfd, 0xf9,
	0x11, 0xed, 0x4a, 0xe1, 0x39, 0x63, 0x67, 0x72, 0x71, 0xcf, 0x83, 0x23, 0x2b, 0x05, 0xfb, 0x83,
	0x5f, 0x76, 0xa5, 0x08, 0xdd, 0xe4, 0x1f, 0xf3, 0xbf, 0x1d, 0x36, 0x3e, 0x50, 0x64, 0x85, 0xa9,
	0xc8, 0xfe, 0x6a, 0xf8, 0xec, 0x3c, 0x06, 0x2d, 0xa2, 0xbc, 0xa1, 0x91, 0x4a, 0x4d, 0x81, 0x5e,
	0xd8, 0x6f, 0xa0, 0x31, 0x97, 0xa9, 0x7b, 0xc3, 0xce, 0xa8, 0x02, 0x55, 0x44, 0x71, 0x9d, 0x4a,
	0x41, 0x5e, 0x67, 0xec, 0x4c, 0xba, 0x61, 0xdf, 0xb0, 0xb9, 0x41, 0xee, 0x35, 0x63, 0x56, 0x49,
	0x50, 0x93, 0xd7, 0x35, 0x42, 0xcf, 0x90, 0x05, 0x6a, 0x72, 0x47, 0xac, 0xaf, 0x09, 0xcb, 0xa8,
	0x12, 0xa0, 0xb1, 0xf0, 0x4e, 0xcc, 0x1d, 0xac, 0x41, 0xa1, 0x21, 0xf3, 0x2f, 0x87, 0x8d, 0x12,
	0xcc, 0x8f, 0xad, 0x3d, 0x67, 0x66, 0x99, 0x75, 0xf3, 0x93, 0xd7, 0xce, 0xdb, 0xac, 0x55, 0x25,
	0x66, 0x50, 0xc8, 0x00, 0x2b, 0xc9, 0xa5, 0x28, 0xcc, 0x13, 0x70, 0xfb, 0x09, 0x4a, 0xa5, 0x0f,
	0x3e, 0xec, 0xa3, 0x3d, 0xfe, 0x74, 0x86, 0xcf, 0x46, 0xdc, 0x2c, 0x1a, 0x69, 0x33, 0xab, 0x09,
	0x57, 0xd9, 0xe6, 0xd5, 0x4a, 0xf1, 0xa9, 0x99, 0xf5, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x0c,
	0x97, 0x73, 0xa0, 0x87, 0x02, 0x00, 0x00,
}
