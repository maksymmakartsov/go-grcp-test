// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: countries.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CountriesRepositoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountriesRepositoryResponse) Reset() {
	*x = CountriesRepositoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryResponse) ProtoMessage() {}

func (x *CountriesRepositoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryResponse.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryResponse) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0}
}

type CountriesRepositoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountriesRepositoryRequest) Reset() {
	*x = CountriesRepositoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryRequest) ProtoMessage() {}

func (x *CountriesRepositoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryRequest.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryRequest) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1}
}

type CountriesRepositoryResponse_Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountriesRepositoryResponse_Country) Reset() {
	*x = CountriesRepositoryResponse_Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryResponse_Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryResponse_Country) ProtoMessage() {}

func (x *CountriesRepositoryResponse_Country) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryResponse_Country.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryResponse_Country) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0}
}

type CountriesRepositoryResponse_Country_Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Likes int32  `protobuf:"varint,3,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *CountriesRepositoryResponse_Country_Single) Reset() {
	*x = CountriesRepositoryResponse_Country_Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryResponse_Country_Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryResponse_Country_Single) ProtoMessage() {}

func (x *CountriesRepositoryResponse_Country_Single) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryResponse_Country_Single.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryResponse_Country_Single) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *CountriesRepositoryResponse_Country_Single) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CountriesRepositoryResponse_Country_Single) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CountriesRepositoryResponse_Country_Single) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

type CountriesRepositoryResponse_Country_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta      *CountriesRepositoryResponse_Country_List_Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Countries []*CountriesRepositoryResponse_Country_Single  `protobuf:"bytes,2,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *CountriesRepositoryResponse_Country_List) Reset() {
	*x = CountriesRepositoryResponse_Country_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryResponse_Country_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryResponse_Country_List) ProtoMessage() {}

func (x *CountriesRepositoryResponse_Country_List) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryResponse_Country_List.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryResponse_Country_List) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *CountriesRepositoryResponse_Country_List) GetMeta() *CountriesRepositoryResponse_Country_List_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *CountriesRepositoryResponse_Country_List) GetCountries() []*CountriesRepositoryResponse_Country_Single {
	if x != nil {
		return x.Countries
	}
	return nil
}

type CountriesRepositoryResponse_Country_List_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CountriesRepositoryResponse_Country_List_Meta) Reset() {
	*x = CountriesRepositoryResponse_Country_List_Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryResponse_Country_List_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryResponse_Country_List_Meta) ProtoMessage() {}

func (x *CountriesRepositoryResponse_Country_List_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryResponse_Country_List_Meta.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryResponse_Country_List_Meta) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{0, 0, 1, 0}
}

func (x *CountriesRepositoryResponse_Country_List_Meta) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CountriesRepositoryResponse_Country_List_Meta) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type CountriesRepositoryRequest_Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CountriesRepositoryRequest_Country) Reset() {
	*x = CountriesRepositoryRequest_Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryRequest_Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryRequest_Country) ProtoMessage() {}

func (x *CountriesRepositoryRequest_Country) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryRequest_Country.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryRequest_Country) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0}
}

type CountriesRepositoryRequest_Country_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset uint32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CountriesRepositoryRequest_Country_List) Reset() {
	*x = CountriesRepositoryRequest_Country_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryRequest_Country_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryRequest_Country_List) ProtoMessage() {}

func (x *CountriesRepositoryRequest_Country_List) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryRequest_Country_List.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryRequest_Country_List) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *CountriesRepositoryRequest_Country_List) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CountriesRepositoryRequest_Country_List) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type CountriesRepositoryRequest_Country_Like struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CountriesRepositoryRequest_Country_Like) Reset() {
	*x = CountriesRepositoryRequest_Country_Like{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryRequest_Country_Like) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryRequest_Country_Like) ProtoMessage() {}

func (x *CountriesRepositoryRequest_Country_Like) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryRequest_Country_Like.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryRequest_Country_Like) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *CountriesRepositoryRequest_Country_Like) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CountriesRepositoryRequest_Country_Dislike struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CountriesRepositoryRequest_Country_Dislike) Reset() {
	*x = CountriesRepositoryRequest_Country_Dislike{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryRequest_Country_Dislike) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryRequest_Country_Dislike) ProtoMessage() {}

func (x *CountriesRepositoryRequest_Country_Dislike) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryRequest_Country_Dislike.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryRequest_Country_Dislike) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 2}
}

func (x *CountriesRepositoryRequest_Country_Dislike) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CountriesRepositoryRequest_Country_Single struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CountriesRepositoryRequest_Country_Single) Reset() {
	*x = CountriesRepositoryRequest_Country_Single{}
	if protoimpl.UnsafeEnabled {
		mi := &file_countries_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountriesRepositoryRequest_Country_Single) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountriesRepositoryRequest_Country_Single) ProtoMessage() {}

func (x *CountriesRepositoryRequest_Country_Single) ProtoReflect() protoreflect.Message {
	mi := &file_countries_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountriesRepositoryRequest_Country_Single.ProtoReflect.Descriptor instead.
func (*CountriesRepositoryRequest_Country_Single) Descriptor() ([]byte, []int) {
	return file_countries_proto_rawDescGZIP(), []int{1, 0, 3}
}

func (x *CountriesRepositoryRequest_Country_Single) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_countries_proto protoreflect.FileDescriptor

var file_countries_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xcf, 0x02, 0x0a,
	0x1b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0xaf, 0x02, 0x0a,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x42, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x1a, 0xdf, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x12, 0x53, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x34, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xab,
	0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x8c, 0x01,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x34, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x1a,
	0x16, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x19, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x6c, 0x69,
	0x6b, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x1a, 0x18, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x32, 0xf1, 0x03, 0x0a,
	0x13, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x77, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x12, 0x34,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69,
	0x6e, 0x67, 0x6c, 0x65, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x73, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x32, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x1a, 0x35, 0x2e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x07, 0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65,
	0x12, 0x35, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x44, 0x69, 0x73, 0x6c, 0x69, 0x6b, 0x65, 0x1a, 0x35, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x00,
	0x42, 0x05, 0x5a, 0x03, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_countries_proto_rawDescOnce sync.Once
	file_countries_proto_rawDescData = file_countries_proto_rawDesc
)

func file_countries_proto_rawDescGZIP() []byte {
	file_countries_proto_rawDescOnce.Do(func() {
		file_countries_proto_rawDescData = protoimpl.X.CompressGZIP(file_countries_proto_rawDescData)
	})
	return file_countries_proto_rawDescData
}

var file_countries_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_countries_proto_goTypes = []interface{}{
	(*CountriesRepositoryResponse)(nil),                   // 0: countries.CountriesRepositoryResponse
	(*CountriesRepositoryRequest)(nil),                    // 1: countries.CountriesRepositoryRequest
	(*CountriesRepositoryResponse_Country)(nil),           // 2: countries.CountriesRepositoryResponse.Country
	(*CountriesRepositoryResponse_Country_Single)(nil),    // 3: countries.CountriesRepositoryResponse.Country.Single
	(*CountriesRepositoryResponse_Country_List)(nil),      // 4: countries.CountriesRepositoryResponse.Country.List
	(*CountriesRepositoryResponse_Country_List_Meta)(nil), // 5: countries.CountriesRepositoryResponse.Country.List.Meta
	(*CountriesRepositoryRequest_Country)(nil),            // 6: countries.CountriesRepositoryRequest.Country
	(*CountriesRepositoryRequest_Country_List)(nil),       // 7: countries.CountriesRepositoryRequest.Country.List
	(*CountriesRepositoryRequest_Country_Like)(nil),       // 8: countries.CountriesRepositoryRequest.Country.Like
	(*CountriesRepositoryRequest_Country_Dislike)(nil),    // 9: countries.CountriesRepositoryRequest.Country.Dislike
	(*CountriesRepositoryRequest_Country_Single)(nil),     // 10: countries.CountriesRepositoryRequest.Country.Single
}
var file_countries_proto_depIdxs = []int32{
	5,  // 0: countries.CountriesRepositoryResponse.Country.List.meta:type_name -> countries.CountriesRepositoryResponse.Country.List.Meta
	3,  // 1: countries.CountriesRepositoryResponse.Country.List.countries:type_name -> countries.CountriesRepositoryResponse.Country.Single
	10, // 2: countries.CountriesRepository.Single:input_type -> countries.CountriesRepositoryRequest.Country.Single
	7,  // 3: countries.CountriesRepository.List:input_type -> countries.CountriesRepositoryRequest.Country.List
	8,  // 4: countries.CountriesRepository.Like:input_type -> countries.CountriesRepositoryRequest.Country.Like
	9,  // 5: countries.CountriesRepository.Dislike:input_type -> countries.CountriesRepositoryRequest.Country.Dislike
	3,  // 6: countries.CountriesRepository.Single:output_type -> countries.CountriesRepositoryResponse.Country.Single
	4,  // 7: countries.CountriesRepository.List:output_type -> countries.CountriesRepositoryResponse.Country.List
	3,  // 8: countries.CountriesRepository.Like:output_type -> countries.CountriesRepositoryResponse.Country.Single
	3,  // 9: countries.CountriesRepository.Dislike:output_type -> countries.CountriesRepositoryResponse.Country.Single
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_countries_proto_init() }
func file_countries_proto_init() {
	if File_countries_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_countries_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryResponse_Country); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryResponse_Country_Single); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryResponse_Country_List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryResponse_Country_List_Meta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryRequest_Country); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryRequest_Country_List); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryRequest_Country_Like); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryRequest_Country_Dislike); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_countries_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CountriesRepositoryRequest_Country_Single); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_countries_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_countries_proto_goTypes,
		DependencyIndexes: file_countries_proto_depIdxs,
		MessageInfos:      file_countries_proto_msgTypes,
	}.Build()
	File_countries_proto = out.File
	file_countries_proto_rawDesc = nil
	file_countries_proto_goTypes = nil
	file_countries_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CountriesRepositoryClient is the client API for CountriesRepository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CountriesRepositoryClient interface {
	Single(ctx context.Context, in *CountriesRepositoryRequest_Country_Single, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_Single, error)
	List(ctx context.Context, in *CountriesRepositoryRequest_Country_List, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_List, error)
	Like(ctx context.Context, in *CountriesRepositoryRequest_Country_Like, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_Single, error)
	Dislike(ctx context.Context, in *CountriesRepositoryRequest_Country_Dislike, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_Single, error)
}

type countriesRepositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewCountriesRepositoryClient(cc grpc.ClientConnInterface) CountriesRepositoryClient {
	return &countriesRepositoryClient{cc}
}

func (c *countriesRepositoryClient) Single(ctx context.Context, in *CountriesRepositoryRequest_Country_Single, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_Single, error) {
	out := new(CountriesRepositoryResponse_Country_Single)
	err := c.cc.Invoke(ctx, "/countries.CountriesRepository/Single", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesRepositoryClient) List(ctx context.Context, in *CountriesRepositoryRequest_Country_List, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_List, error) {
	out := new(CountriesRepositoryResponse_Country_List)
	err := c.cc.Invoke(ctx, "/countries.CountriesRepository/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesRepositoryClient) Like(ctx context.Context, in *CountriesRepositoryRequest_Country_Like, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_Single, error) {
	out := new(CountriesRepositoryResponse_Country_Single)
	err := c.cc.Invoke(ctx, "/countries.CountriesRepository/Like", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *countriesRepositoryClient) Dislike(ctx context.Context, in *CountriesRepositoryRequest_Country_Dislike, opts ...grpc.CallOption) (*CountriesRepositoryResponse_Country_Single, error) {
	out := new(CountriesRepositoryResponse_Country_Single)
	err := c.cc.Invoke(ctx, "/countries.CountriesRepository/Dislike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CountriesRepositoryServer is the server API for CountriesRepository service.
type CountriesRepositoryServer interface {
	Single(context.Context, *CountriesRepositoryRequest_Country_Single) (*CountriesRepositoryResponse_Country_Single, error)
	List(context.Context, *CountriesRepositoryRequest_Country_List) (*CountriesRepositoryResponse_Country_List, error)
	Like(context.Context, *CountriesRepositoryRequest_Country_Like) (*CountriesRepositoryResponse_Country_Single, error)
	Dislike(context.Context, *CountriesRepositoryRequest_Country_Dislike) (*CountriesRepositoryResponse_Country_Single, error)
}

// UnimplementedCountriesRepositoryServer can be embedded to have forward compatible implementations.
type UnimplementedCountriesRepositoryServer struct {
}

func (*UnimplementedCountriesRepositoryServer) Single(context.Context, *CountriesRepositoryRequest_Country_Single) (*CountriesRepositoryResponse_Country_Single, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Single not implemented")
}
func (*UnimplementedCountriesRepositoryServer) List(context.Context, *CountriesRepositoryRequest_Country_List) (*CountriesRepositoryResponse_Country_List, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCountriesRepositoryServer) Like(context.Context, *CountriesRepositoryRequest_Country_Like) (*CountriesRepositoryResponse_Country_Single, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Like not implemented")
}
func (*UnimplementedCountriesRepositoryServer) Dislike(context.Context, *CountriesRepositoryRequest_Country_Dislike) (*CountriesRepositoryResponse_Country_Single, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dislike not implemented")
}

func RegisterCountriesRepositoryServer(s *grpc.Server, srv CountriesRepositoryServer) {
	s.RegisterService(&_CountriesRepository_serviceDesc, srv)
}

func _CountriesRepository_Single_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountriesRepositoryRequest_Country_Single)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesRepositoryServer).Single(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.CountriesRepository/Single",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesRepositoryServer).Single(ctx, req.(*CountriesRepositoryRequest_Country_Single))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountriesRepository_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountriesRepositoryRequest_Country_List)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesRepositoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.CountriesRepository/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesRepositoryServer).List(ctx, req.(*CountriesRepositoryRequest_Country_List))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountriesRepository_Like_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountriesRepositoryRequest_Country_Like)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesRepositoryServer).Like(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.CountriesRepository/Like",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesRepositoryServer).Like(ctx, req.(*CountriesRepositoryRequest_Country_Like))
	}
	return interceptor(ctx, in, info, handler)
}

func _CountriesRepository_Dislike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountriesRepositoryRequest_Country_Dislike)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CountriesRepositoryServer).Dislike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/countries.CountriesRepository/Dislike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CountriesRepositoryServer).Dislike(ctx, req.(*CountriesRepositoryRequest_Country_Dislike))
	}
	return interceptor(ctx, in, info, handler)
}

var _CountriesRepository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "countries.CountriesRepository",
	HandlerType: (*CountriesRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Single",
			Handler:    _CountriesRepository_Single_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CountriesRepository_List_Handler,
		},
		{
			MethodName: "Like",
			Handler:    _CountriesRepository_Like_Handler,
		},
		{
			MethodName: "Dislike",
			Handler:    _CountriesRepository_Dislike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "countries.proto",
}
