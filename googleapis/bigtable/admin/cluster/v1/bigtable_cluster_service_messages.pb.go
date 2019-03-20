// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/bigtable/admin/cluster/v1/bigtable_cluster_service_messages.proto

package cluster

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Request message for BigtableClusterService.ListZones.
type ListZonesRequest struct {
	// The unique name of the project for which a list of supported zones is
	// requested.
	// Values are of the form projects/<project>
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListZonesRequest) Reset()         { *m = ListZonesRequest{} }
func (m *ListZonesRequest) String() string { return proto.CompactTextString(m) }
func (*ListZonesRequest) ProtoMessage()    {}
func (*ListZonesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{0}
}

func (m *ListZonesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListZonesRequest.Unmarshal(m, b)
}
func (m *ListZonesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListZonesRequest.Marshal(b, m, deterministic)
}
func (m *ListZonesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListZonesRequest.Merge(m, src)
}
func (m *ListZonesRequest) XXX_Size() int {
	return xxx_messageInfo_ListZonesRequest.Size(m)
}
func (m *ListZonesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListZonesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListZonesRequest proto.InternalMessageInfo

func (m *ListZonesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message for BigtableClusterService.ListZones.
type ListZonesResponse struct {
	// The list of requested zones.
	Zones                []*Zone  `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListZonesResponse) Reset()         { *m = ListZonesResponse{} }
func (m *ListZonesResponse) String() string { return proto.CompactTextString(m) }
func (*ListZonesResponse) ProtoMessage()    {}
func (*ListZonesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{1}
}

func (m *ListZonesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListZonesResponse.Unmarshal(m, b)
}
func (m *ListZonesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListZonesResponse.Marshal(b, m, deterministic)
}
func (m *ListZonesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListZonesResponse.Merge(m, src)
}
func (m *ListZonesResponse) XXX_Size() int {
	return xxx_messageInfo_ListZonesResponse.Size(m)
}
func (m *ListZonesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListZonesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListZonesResponse proto.InternalMessageInfo

func (m *ListZonesResponse) GetZones() []*Zone {
	if m != nil {
		return m.Zones
	}
	return nil
}

// Request message for BigtableClusterService.GetCluster.
type GetClusterRequest struct {
	// The unique name of the requested cluster.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClusterRequest) Reset()         { *m = GetClusterRequest{} }
func (m *GetClusterRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterRequest) ProtoMessage()    {}
func (*GetClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{2}
}

func (m *GetClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClusterRequest.Unmarshal(m, b)
}
func (m *GetClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClusterRequest.Marshal(b, m, deterministic)
}
func (m *GetClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterRequest.Merge(m, src)
}
func (m *GetClusterRequest) XXX_Size() int {
	return xxx_messageInfo_GetClusterRequest.Size(m)
}
func (m *GetClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterRequest proto.InternalMessageInfo

func (m *GetClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableClusterService.ListClusters.
type ListClustersRequest struct {
	// The unique name of the project for which a list of clusters is requested.
	// Values are of the form projects/<project>
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClustersRequest) Reset()         { *m = ListClustersRequest{} }
func (m *ListClustersRequest) String() string { return proto.CompactTextString(m) }
func (*ListClustersRequest) ProtoMessage()    {}
func (*ListClustersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{3}
}

func (m *ListClustersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClustersRequest.Unmarshal(m, b)
}
func (m *ListClustersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClustersRequest.Marshal(b, m, deterministic)
}
func (m *ListClustersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClustersRequest.Merge(m, src)
}
func (m *ListClustersRequest) XXX_Size() int {
	return xxx_messageInfo_ListClustersRequest.Size(m)
}
func (m *ListClustersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClustersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListClustersRequest proto.InternalMessageInfo

func (m *ListClustersRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Response message for BigtableClusterService.ListClusters.
type ListClustersResponse struct {
	// The list of requested Clusters.
	Clusters []*Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// The zones for which clusters could not be retrieved.
	FailedZones          []*Zone  `protobuf:"bytes,2,rep,name=failed_zones,json=failedZones,proto3" json:"failed_zones,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListClustersResponse) Reset()         { *m = ListClustersResponse{} }
func (m *ListClustersResponse) String() string { return proto.CompactTextString(m) }
func (*ListClustersResponse) ProtoMessage()    {}
func (*ListClustersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{4}
}

func (m *ListClustersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListClustersResponse.Unmarshal(m, b)
}
func (m *ListClustersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListClustersResponse.Marshal(b, m, deterministic)
}
func (m *ListClustersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListClustersResponse.Merge(m, src)
}
func (m *ListClustersResponse) XXX_Size() int {
	return xxx_messageInfo_ListClustersResponse.Size(m)
}
func (m *ListClustersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListClustersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListClustersResponse proto.InternalMessageInfo

func (m *ListClustersResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *ListClustersResponse) GetFailedZones() []*Zone {
	if m != nil {
		return m.FailedZones
	}
	return nil
}

// Request message for BigtableClusterService.CreateCluster.
type CreateClusterRequest struct {
	// The unique name of the zone in which to create the cluster.
	// Values are of the form projects/<project>/zones/<zone>
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The id to be used when referring to the new cluster within its zone,
	// e.g. just the "test-cluster" section of the full name
	// "projects/<project>/zones/<zone>/clusters/test-cluster".
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The cluster to create.
	// The "name", "delete_time", and "current_operation" fields must be left
	// blank.
	Cluster              *Cluster `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateClusterRequest) Reset()         { *m = CreateClusterRequest{} }
func (m *CreateClusterRequest) String() string { return proto.CompactTextString(m) }
func (*CreateClusterRequest) ProtoMessage()    {}
func (*CreateClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{5}
}

func (m *CreateClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClusterRequest.Unmarshal(m, b)
}
func (m *CreateClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClusterRequest.Marshal(b, m, deterministic)
}
func (m *CreateClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClusterRequest.Merge(m, src)
}
func (m *CreateClusterRequest) XXX_Size() int {
	return xxx_messageInfo_CreateClusterRequest.Size(m)
}
func (m *CreateClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClusterRequest proto.InternalMessageInfo

func (m *CreateClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateClusterRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *CreateClusterRequest) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

// Metadata type for the operation returned by
// BigtableClusterService.CreateCluster.
type CreateClusterMetadata struct {
	// The request which prompted the creation of this operation.
	OriginalRequest *CreateClusterRequest `protobuf:"bytes,1,opt,name=original_request,json=originalRequest,proto3" json:"original_request,omitempty"`
	// The time at which original_request was received.
	RequestTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	FinishTime           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateClusterMetadata) Reset()         { *m = CreateClusterMetadata{} }
func (m *CreateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*CreateClusterMetadata) ProtoMessage()    {}
func (*CreateClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{6}
}

func (m *CreateClusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateClusterMetadata.Unmarshal(m, b)
}
func (m *CreateClusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateClusterMetadata.Marshal(b, m, deterministic)
}
func (m *CreateClusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateClusterMetadata.Merge(m, src)
}
func (m *CreateClusterMetadata) XXX_Size() int {
	return xxx_messageInfo_CreateClusterMetadata.Size(m)
}
func (m *CreateClusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateClusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_CreateClusterMetadata proto.InternalMessageInfo

func (m *CreateClusterMetadata) GetOriginalRequest() *CreateClusterRequest {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *CreateClusterMetadata) GetRequestTime() *timestamp.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *CreateClusterMetadata) GetFinishTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// Metadata type for the operation returned by
// BigtableClusterService.UpdateCluster.
type UpdateClusterMetadata struct {
	// The request which prompted the creation of this operation.
	OriginalRequest *Cluster `protobuf:"bytes,1,opt,name=original_request,json=originalRequest,proto3" json:"original_request,omitempty"`
	// The time at which original_request was received.
	RequestTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// The time at which this operation was cancelled. If set, this operation is
	// in the process of undoing itself (which is guaranteed to succeed) and
	// cannot be cancelled again.
	CancelTime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=cancel_time,json=cancelTime,proto3" json:"cancel_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	FinishTime           *timestamp.Timestamp `protobuf:"bytes,4,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateClusterMetadata) Reset()         { *m = UpdateClusterMetadata{} }
func (m *UpdateClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterMetadata) ProtoMessage()    {}
func (*UpdateClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{7}
}

func (m *UpdateClusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateClusterMetadata.Unmarshal(m, b)
}
func (m *UpdateClusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateClusterMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateClusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateClusterMetadata.Merge(m, src)
}
func (m *UpdateClusterMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateClusterMetadata.Size(m)
}
func (m *UpdateClusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateClusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateClusterMetadata proto.InternalMessageInfo

func (m *UpdateClusterMetadata) GetOriginalRequest() *Cluster {
	if m != nil {
		return m.OriginalRequest
	}
	return nil
}

func (m *UpdateClusterMetadata) GetRequestTime() *timestamp.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UpdateClusterMetadata) GetCancelTime() *timestamp.Timestamp {
	if m != nil {
		return m.CancelTime
	}
	return nil
}

func (m *UpdateClusterMetadata) GetFinishTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// Request message for BigtableClusterService.DeleteCluster.
type DeleteClusterRequest struct {
	// The unique name of the cluster to be deleted.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteClusterRequest) Reset()         { *m = DeleteClusterRequest{} }
func (m *DeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteClusterRequest) ProtoMessage()    {}
func (*DeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{8}
}

func (m *DeleteClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteClusterRequest.Unmarshal(m, b)
}
func (m *DeleteClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteClusterRequest.Marshal(b, m, deterministic)
}
func (m *DeleteClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteClusterRequest.Merge(m, src)
}
func (m *DeleteClusterRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteClusterRequest.Size(m)
}
func (m *DeleteClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteClusterRequest proto.InternalMessageInfo

func (m *DeleteClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for BigtableClusterService.UndeleteCluster.
type UndeleteClusterRequest struct {
	// The unique name of the cluster to be un-deleted.
	// Values are of the form projects/<project>/zones/<zone>/clusters/<cluster>
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UndeleteClusterRequest) Reset()         { *m = UndeleteClusterRequest{} }
func (m *UndeleteClusterRequest) String() string { return proto.CompactTextString(m) }
func (*UndeleteClusterRequest) ProtoMessage()    {}
func (*UndeleteClusterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{9}
}

func (m *UndeleteClusterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndeleteClusterRequest.Unmarshal(m, b)
}
func (m *UndeleteClusterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndeleteClusterRequest.Marshal(b, m, deterministic)
}
func (m *UndeleteClusterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndeleteClusterRequest.Merge(m, src)
}
func (m *UndeleteClusterRequest) XXX_Size() int {
	return xxx_messageInfo_UndeleteClusterRequest.Size(m)
}
func (m *UndeleteClusterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UndeleteClusterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UndeleteClusterRequest proto.InternalMessageInfo

func (m *UndeleteClusterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Metadata type for the operation returned by
// BigtableClusterService.UndeleteCluster.
type UndeleteClusterMetadata struct {
	// The time at which the original request was received.
	RequestTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=request_time,json=requestTime,proto3" json:"request_time,omitempty"`
	// The time at which this operation failed or was completed successfully.
	FinishTime           *timestamp.Timestamp `protobuf:"bytes,2,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UndeleteClusterMetadata) Reset()         { *m = UndeleteClusterMetadata{} }
func (m *UndeleteClusterMetadata) String() string { return proto.CompactTextString(m) }
func (*UndeleteClusterMetadata) ProtoMessage()    {}
func (*UndeleteClusterMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{10}
}

func (m *UndeleteClusterMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UndeleteClusterMetadata.Unmarshal(m, b)
}
func (m *UndeleteClusterMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UndeleteClusterMetadata.Marshal(b, m, deterministic)
}
func (m *UndeleteClusterMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UndeleteClusterMetadata.Merge(m, src)
}
func (m *UndeleteClusterMetadata) XXX_Size() int {
	return xxx_messageInfo_UndeleteClusterMetadata.Size(m)
}
func (m *UndeleteClusterMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UndeleteClusterMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UndeleteClusterMetadata proto.InternalMessageInfo

func (m *UndeleteClusterMetadata) GetRequestTime() *timestamp.Timestamp {
	if m != nil {
		return m.RequestTime
	}
	return nil
}

func (m *UndeleteClusterMetadata) GetFinishTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinishTime
	}
	return nil
}

// Metadata type for operations initiated by the V2 BigtableAdmin service.
// More complete information for such operations is available via the V2 API.
type V2OperationMetadata struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *V2OperationMetadata) Reset()         { *m = V2OperationMetadata{} }
func (m *V2OperationMetadata) String() string { return proto.CompactTextString(m) }
func (*V2OperationMetadata) ProtoMessage()    {}
func (*V2OperationMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a8715cfb8408734, []int{11}
}

func (m *V2OperationMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_V2OperationMetadata.Unmarshal(m, b)
}
func (m *V2OperationMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_V2OperationMetadata.Marshal(b, m, deterministic)
}
func (m *V2OperationMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_V2OperationMetadata.Merge(m, src)
}
func (m *V2OperationMetadata) XXX_Size() int {
	return xxx_messageInfo_V2OperationMetadata.Size(m)
}
func (m *V2OperationMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_V2OperationMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_V2OperationMetadata proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListZonesRequest)(nil), "google.bigtable.admin.cluster.v1.ListZonesRequest")
	proto.RegisterType((*ListZonesResponse)(nil), "google.bigtable.admin.cluster.v1.ListZonesResponse")
	proto.RegisterType((*GetClusterRequest)(nil), "google.bigtable.admin.cluster.v1.GetClusterRequest")
	proto.RegisterType((*ListClustersRequest)(nil), "google.bigtable.admin.cluster.v1.ListClustersRequest")
	proto.RegisterType((*ListClustersResponse)(nil), "google.bigtable.admin.cluster.v1.ListClustersResponse")
	proto.RegisterType((*CreateClusterRequest)(nil), "google.bigtable.admin.cluster.v1.CreateClusterRequest")
	proto.RegisterType((*CreateClusterMetadata)(nil), "google.bigtable.admin.cluster.v1.CreateClusterMetadata")
	proto.RegisterType((*UpdateClusterMetadata)(nil), "google.bigtable.admin.cluster.v1.UpdateClusterMetadata")
	proto.RegisterType((*DeleteClusterRequest)(nil), "google.bigtable.admin.cluster.v1.DeleteClusterRequest")
	proto.RegisterType((*UndeleteClusterRequest)(nil), "google.bigtable.admin.cluster.v1.UndeleteClusterRequest")
	proto.RegisterType((*UndeleteClusterMetadata)(nil), "google.bigtable.admin.cluster.v1.UndeleteClusterMetadata")
	proto.RegisterType((*V2OperationMetadata)(nil), "google.bigtable.admin.cluster.v1.V2OperationMetadata")
}

func init() {
	proto.RegisterFile("google/bigtable/admin/cluster/v1/bigtable_cluster_service_messages.proto", fileDescriptor_2a8715cfb8408734)
}

var fileDescriptor_2a8715cfb8408734 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xd5, 0x26, 0xe5, 0xa3, 0xe3, 0x4a, 0xb4, 0x6e, 0x02, 0x51, 0x24, 0x44, 0x64, 0x50, 0x69,
	0x11, 0xb2, 0xd5, 0x20, 0x71, 0x69, 0xb9, 0x24, 0xa0, 0x52, 0x89, 0x8a, 0x12, 0x5a, 0x0e, 0xbd,
	0x58, 0x9b, 0x78, 0x62, 0x56, 0xb2, 0x77, 0x8d, 0x77, 0x93, 0x03, 0x3f, 0x82, 0x1b, 0xfc, 0x04,
	0xc4, 0x2f, 0xe4, 0x8c, 0xec, 0xdd, 0x8d, 0x68, 0x95, 0xd6, 0xb1, 0x10, 0xb7, 0xdd, 0x99, 0xf7,
	0x66, 0xde, 0x9b, 0x1d, 0x69, 0xe1, 0x6d, 0x2c, 0x44, 0x9c, 0x60, 0x30, 0x66, 0xb1, 0xa2, 0xe3,
	0x04, 0x03, 0x1a, 0xa5, 0x8c, 0x07, 0x93, 0x64, 0x26, 0x15, 0xe6, 0xc1, 0x7c, 0x7f, 0x91, 0x09,
	0x4d, 0x2c, 0x94, 0x98, 0xcf, 0xd9, 0x04, 0xc3, 0x14, 0xa5, 0xa4, 0x31, 0x4a, 0x3f, 0xcb, 0x85,
	0x12, 0x6e, 0x4f, 0x57, 0xf2, 0x2d, 0xde, 0x2f, 0x2b, 0xf9, 0x86, 0xe5, 0xcf, 0xf7, 0xbb, 0x87,
	0xf5, 0x7b, 0x45, 0x54, 0x51, 0x5d, 0xbf, 0xfb, 0xc8, 0xb0, 0xcb, 0xdb, 0x78, 0x36, 0x0d, 0x14,
	0x4b, 0x51, 0x2a, 0x9a, 0x66, 0x1a, 0xe0, 0xed, 0xc0, 0xe6, 0x3b, 0x26, 0xd5, 0x85, 0xe0, 0x28,
	0x47, 0xf8, 0x65, 0x86, 0x52, 0xb9, 0x2e, 0xac, 0x71, 0x9a, 0x62, 0x87, 0xf4, 0xc8, 0xee, 0xfa,
	0xa8, 0x3c, 0x7b, 0x1f, 0x60, 0xeb, 0x2f, 0x9c, 0xcc, 0x04, 0x97, 0xe8, 0x1e, 0xc2, 0xad, 0xaf,
	0x45, 0xa0, 0x43, 0x7a, 0xcd, 0x5d, 0xa7, 0xbf, 0xe3, 0x57, 0xb9, 0xf1, 0x0b, 0xfe, 0x48, 0x93,
	0xbc, 0xa7, 0xb0, 0x75, 0x84, 0x6a, 0xa8, 0x93, 0x37, 0xf5, 0xde, 0x83, 0xed, 0xa2, 0xb7, 0x41,
	0xde, 0x28, 0xf3, 0x17, 0x81, 0xd6, 0x65, 0xac, 0x91, 0xfa, 0x06, 0xee, 0x1a, 0x19, 0x56, 0xed,
	0x5e, 0xb5, 0x5a, 0xab, 0x6d, 0x41, 0x75, 0x8f, 0x61, 0x63, 0x4a, 0x59, 0x82, 0x51, 0xa8, 0x8d,
	0x37, 0x6a, 0x19, 0x77, 0x34, 0xb7, 0x1c, 0xa2, 0xf7, 0x8d, 0x40, 0x6b, 0x98, 0x23, 0x55, 0x58,
	0x3d, 0x02, 0xf7, 0x21, 0x80, 0x7d, 0x5d, 0x16, 0x75, 0x1a, 0x65, 0x66, 0xdd, 0x44, 0x8e, 0x23,
	0x77, 0x08, 0x77, 0xcc, 0xa5, 0xd3, 0xec, 0x91, 0x7a, 0xe6, 0x2c, 0xd3, 0xfb, 0x4d, 0xa0, 0x7d,
	0x49, 0xd0, 0x09, 0x2a, 0x5a, 0xec, 0x92, 0x4b, 0x61, 0x53, 0xe4, 0x2c, 0x66, 0x9c, 0x26, 0x61,
	0xae, 0x55, 0x96, 0xea, 0x9c, 0xfe, 0xcb, 0x15, 0xfa, 0x2c, 0xf1, 0x38, 0xba, 0x67, 0xeb, 0x59,
	0xd3, 0xaf, 0x60, 0xc3, 0x54, 0x0e, 0x8b, 0x15, 0x2d, 0x2d, 0x3a, 0xfd, 0xae, 0x2d, 0x6f, 0xf7,
	0xd7, 0x3f, 0xb3, 0xfb, 0x3b, 0x72, 0x0c, 0xbe, 0x88, 0xb8, 0x07, 0xe0, 0x4c, 0x19, 0x67, 0xf2,
	0xb3, 0x66, 0x37, 0x2b, 0xd9, 0xa0, 0xe1, 0x45, 0xc0, 0xfb, 0xd9, 0x80, 0xf6, 0x79, 0x16, 0x2d,
	0x31, 0x7e, 0x76, 0xad, 0xf1, 0x1a, 0x03, 0xfe, 0x0f, 0x5e, 0x27, 0x94, 0x4f, 0x30, 0x59, 0xd9,
	0xab, 0x86, 0x2f, 0x1b, 0xd4, 0x5a, 0xad, 0x41, 0x3d, 0x83, 0xd6, 0x6b, 0x4c, 0x70, 0x95, 0x8d,
	0xf5, 0x9e, 0xc3, 0xfd, 0x73, 0x1e, 0xad, 0x8a, 0xfe, 0x4e, 0xe0, 0xc1, 0x15, 0xf8, 0xe2, 0x11,
	0xae, 0x8e, 0x8b, 0xfc, 0xd3, 0x6a, 0x34, 0x6a, 0x39, 0x6e, 0xc3, 0xf6, 0xa7, 0xfe, 0xfb, 0x0c,
	0x73, 0xaa, 0x98, 0xe0, 0x56, 0xd2, 0xe0, 0x07, 0x81, 0x27, 0x13, 0x91, 0x56, 0xee, 0xc0, 0xe0,
	0xf1, 0xc0, 0xa4, 0x8c, 0xa9, 0x8f, 0xfa, 0x1b, 0x38, 0x31, 0xbf, 0xc0, 0x69, 0xd1, 0xfd, 0x94,
	0x5c, 0x1c, 0x99, 0x42, 0xb1, 0x48, 0x28, 0x8f, 0x7d, 0x91, 0xc7, 0x41, 0x8c, 0xbc, 0xd4, 0x16,
	0xe8, 0x14, 0xcd, 0x98, 0xbc, 0xfe, 0x0f, 0x38, 0x30, 0xc7, 0xf1, 0xed, 0x92, 0xf3, 0xe2, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x75, 0x68, 0x13, 0xa2, 0x06, 0x00, 0x00,
}
