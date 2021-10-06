// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: watermarksvc.proto

package watermark

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

type StatusReply_Status int32

const (
	StatusReply_PENDING     StatusReply_Status = 0
	StatusReply_STARTED     StatusReply_Status = 1
	StatusReply_IN_PROGRESS StatusReply_Status = 2
	StatusReply_FINISHED    StatusReply_Status = 3
	StatusReply_FAILED      StatusReply_Status = 4
)

// Enum value maps for StatusReply_Status.
var (
	StatusReply_Status_name = map[int32]string{
		0: "PENDING",
		1: "STARTED",
		2: "IN_PROGRESS",
		3: "FINISHED",
		4: "FAILED",
	}
	StatusReply_Status_value = map[string]int32{
		"PENDING":     0,
		"STARTED":     1,
		"IN_PROGRESS": 2,
		"FINISHED":    3,
		"FAILED":      4,
	}
)

func (x StatusReply_Status) Enum() *StatusReply_Status {
	p := new(StatusReply_Status)
	*p = x
	return p
}

func (x StatusReply_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusReply_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_watermarksvc_proto_enumTypes[0].Descriptor()
}

func (StatusReply_Status) Type() protoreflect.EnumType {
	return &file_watermarksvc_proto_enumTypes[0]
}

func (x StatusReply_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusReply_Status.Descriptor instead.
func (StatusReply_Status) EnumDescriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{4, 0}
}

type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content   string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author    string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Topic     string `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Watermark string `protobuf:"bytes,5,opt,name=watermark,proto3" json:"watermark,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{0}
}

func (x *Document) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Document) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Document) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Document) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Document) GetWatermark() string {
	if x != nil {
		return x.Watermark
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filters []*GetRequest_Filters `protobuf:"bytes,1,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetFilters() []*GetRequest_Filters {
	if x != nil {
		return x.Filters
	}
	return nil
}

type GetReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Documents []*Document `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *GetReply) Reset() {
	*x = GetReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReply) ProtoMessage() {}

func (x *GetReply) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReply.ProtoReflect.Descriptor instead.
func (*GetReply) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{2}
}

func (x *GetReply) GetDocuments() []*Document {
	if x != nil {
		return x.Documents
	}
	return nil
}

func (x *GetReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketID string `protobuf:"bytes,1,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{3}
}

func (x *StatusRequest) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status StatusReply_Status `protobuf:"varint,1,opt,name=status,proto3,enum=pb.StatusReply_Status" json:"status,omitempty"`
	Err    string             `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{4}
}

func (x *StatusReply) GetStatus() StatusReply_Status {
	if x != nil {
		return x.Status
	}
	return StatusReply_PENDING
}

func (x *StatusReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type WatermarkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketID string `protobuf:"bytes,1,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
	Mark     string `protobuf:"bytes,2,opt,name=mark,proto3" json:"mark,omitempty"`
}

func (x *WatermarkRequest) Reset() {
	*x = WatermarkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatermarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatermarkRequest) ProtoMessage() {}

func (x *WatermarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatermarkRequest.ProtoReflect.Descriptor instead.
func (*WatermarkRequest) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{5}
}

func (x *WatermarkRequest) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

func (x *WatermarkRequest) GetMark() string {
	if x != nil {
		return x.Mark
	}
	return ""
}

type WatermarkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *WatermarkReply) Reset() {
	*x = WatermarkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatermarkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatermarkReply) ProtoMessage() {}

func (x *WatermarkReply) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatermarkReply.ProtoReflect.Descriptor instead.
func (*WatermarkReply) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{6}
}

func (x *WatermarkReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *WatermarkReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type AddDocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Document *Document `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *AddDocumentRequest) Reset() {
	*x = AddDocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDocumentRequest) ProtoMessage() {}

func (x *AddDocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDocumentRequest.ProtoReflect.Descriptor instead.
func (*AddDocumentRequest) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{7}
}

func (x *AddDocumentRequest) GetDocument() *Document {
	if x != nil {
		return x.Document
	}
	return nil
}

type AddDcoumentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TicketID string `protobuf:"bytes,1,opt,name=ticketID,proto3" json:"ticketID,omitempty"`
	Err      string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *AddDcoumentReply) Reset() {
	*x = AddDcoumentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDcoumentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDcoumentReply) ProtoMessage() {}

func (x *AddDcoumentReply) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDcoumentReply.ProtoReflect.Descriptor instead.
func (*AddDcoumentReply) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{8}
}

func (x *AddDcoumentReply) GetTicketID() string {
	if x != nil {
		return x.TicketID
	}
	return ""
}

func (x *AddDcoumentReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type ServiceStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServiceStatusRequest) Reset() {
	*x = ServiceStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatusRequest) ProtoMessage() {}

func (x *ServiceStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatusRequest.ProtoReflect.Descriptor instead.
func (*ServiceStatusRequest) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{9}
}

type ServiceStatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err  string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ServiceStatusReply) Reset() {
	*x = ServiceStatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceStatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceStatusReply) ProtoMessage() {}

func (x *ServiceStatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceStatusReply.ProtoReflect.Descriptor instead.
func (*ServiceStatusReply) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{10}
}

func (x *ServiceStatusReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ServiceStatusReply) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type GetRequest_Filters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetRequest_Filters) Reset() {
	*x = GetRequest_Filters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_watermarksvc_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest_Filters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest_Filters) ProtoMessage() {}

func (x *GetRequest_Filters) ProtoReflect() protoreflect.Message {
	mi := &file_watermarksvc_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest_Filters.ProtoReflect.Descriptor instead.
func (*GetRequest_Filters) Descriptor() ([]byte, []int) {
	return file_watermarksvc_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetRequest_Filters) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetRequest_Filters) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_watermarksvc_proto protoreflect.FileDescriptor

var file_watermarksvc_proto_rawDesc = []byte{
	0x0a, 0x12, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x76, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0x71, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x30, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x73, 0x1a, 0x31, 0x0a, 0x07, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x48, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2a, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x45, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x22, 0x2b,
	0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x22, 0x9e, 0x01, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x45,
	0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x22, 0x4d, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x22, 0x42, 0x0a, 0x10,
	0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x61, 0x72, 0x6b,
	0x22, 0x36, 0x0a, 0x0e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x3e, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x40, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x44,
	0x63, 0x6f, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x3a, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x95,
	0x02, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x23, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x35, 0x0a, 0x09, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x14,
	0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x61, 0x74, 0x65, 0x72, 0x6d,
	0x61, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x63, 0x6f, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x41, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x77, 0x61, 0x74,
	0x65, 0x72, 0x6d, 0x61, 0x72, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_watermarksvc_proto_rawDescOnce sync.Once
	file_watermarksvc_proto_rawDescData = file_watermarksvc_proto_rawDesc
)

func file_watermarksvc_proto_rawDescGZIP() []byte {
	file_watermarksvc_proto_rawDescOnce.Do(func() {
		file_watermarksvc_proto_rawDescData = protoimpl.X.CompressGZIP(file_watermarksvc_proto_rawDescData)
	})
	return file_watermarksvc_proto_rawDescData
}

var file_watermarksvc_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_watermarksvc_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_watermarksvc_proto_goTypes = []interface{}{
	(StatusReply_Status)(0),      // 0: pb.StatusReply.Status
	(*Document)(nil),             // 1: pb.Document
	(*GetRequest)(nil),           // 2: pb.GetRequest
	(*GetReply)(nil),             // 3: pb.GetReply
	(*StatusRequest)(nil),        // 4: pb.StatusRequest
	(*StatusReply)(nil),          // 5: pb.StatusReply
	(*WatermarkRequest)(nil),     // 6: pb.WatermarkRequest
	(*WatermarkReply)(nil),       // 7: pb.WatermarkReply
	(*AddDocumentRequest)(nil),   // 8: pb.AddDocumentRequest
	(*AddDcoumentReply)(nil),     // 9: pb.AddDcoumentReply
	(*ServiceStatusRequest)(nil), // 10: pb.ServiceStatusRequest
	(*ServiceStatusReply)(nil),   // 11: pb.ServiceStatusReply
	(*GetRequest_Filters)(nil),   // 12: pb.GetRequest.Filters
}
var file_watermarksvc_proto_depIdxs = []int32{
	12, // 0: pb.GetRequest.filters:type_name -> pb.GetRequest.Filters
	1,  // 1: pb.GetReply.documents:type_name -> pb.Document
	0,  // 2: pb.StatusReply.status:type_name -> pb.StatusReply.Status
	1,  // 3: pb.AddDocumentRequest.document:type_name -> pb.Document
	2,  // 4: pb.Watermark.Get:input_type -> pb.GetRequest
	6,  // 5: pb.Watermark.Watermark:input_type -> pb.WatermarkRequest
	4,  // 6: pb.Watermark.Status:input_type -> pb.StatusRequest
	8,  // 7: pb.Watermark.AddDocument:input_type -> pb.AddDocumentRequest
	10, // 8: pb.Watermark.ServiceStatus:input_type -> pb.ServiceStatusRequest
	3,  // 9: pb.Watermark.Get:output_type -> pb.GetReply
	7,  // 10: pb.Watermark.Watermark:output_type -> pb.WatermarkReply
	5,  // 11: pb.Watermark.Status:output_type -> pb.StatusReply
	9,  // 12: pb.Watermark.AddDocument:output_type -> pb.AddDcoumentReply
	11, // 13: pb.Watermark.ServiceStatus:output_type -> pb.ServiceStatusReply
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_watermarksvc_proto_init() }
func file_watermarksvc_proto_init() {
	if File_watermarksvc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_watermarksvc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Document); i {
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
		file_watermarksvc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_watermarksvc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReply); i {
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
		file_watermarksvc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_watermarksvc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_watermarksvc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatermarkRequest); i {
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
		file_watermarksvc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatermarkReply); i {
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
		file_watermarksvc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDocumentRequest); i {
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
		file_watermarksvc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDcoumentReply); i {
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
		file_watermarksvc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatusRequest); i {
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
		file_watermarksvc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceStatusReply); i {
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
		file_watermarksvc_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest_Filters); i {
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
			RawDescriptor: file_watermarksvc_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_watermarksvc_proto_goTypes,
		DependencyIndexes: file_watermarksvc_proto_depIdxs,
		EnumInfos:         file_watermarksvc_proto_enumTypes,
		MessageInfos:      file_watermarksvc_proto_msgTypes,
	}.Build()
	File_watermarksvc_proto = out.File
	file_watermarksvc_proto_rawDesc = nil
	file_watermarksvc_proto_goTypes = nil
	file_watermarksvc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WatermarkClient is the client API for Watermark service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WatermarkClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkReply, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error)
	AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDcoumentReply, error)
	ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error)
}

type watermarkClient struct {
	cc grpc.ClientConnInterface
}

func NewWatermarkClient(cc grpc.ClientConnInterface) WatermarkClient {
	return &watermarkClient{cc}
}

func (c *watermarkClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) Watermark(ctx context.Context, in *WatermarkRequest, opts ...grpc.CallOption) (*WatermarkReply, error) {
	out := new(WatermarkReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/Watermark", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusReply, error) {
	out := new(StatusReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) AddDocument(ctx context.Context, in *AddDocumentRequest, opts ...grpc.CallOption) (*AddDcoumentReply, error) {
	out := new(AddDcoumentReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/AddDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *watermarkClient) ServiceStatus(ctx context.Context, in *ServiceStatusRequest, opts ...grpc.CallOption) (*ServiceStatusReply, error) {
	out := new(ServiceStatusReply)
	err := c.cc.Invoke(ctx, "/pb.Watermark/ServiceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WatermarkServer is the server API for Watermark service.
type WatermarkServer interface {
	Get(context.Context, *GetRequest) (*GetReply, error)
	Watermark(context.Context, *WatermarkRequest) (*WatermarkReply, error)
	Status(context.Context, *StatusRequest) (*StatusReply, error)
	AddDocument(context.Context, *AddDocumentRequest) (*AddDcoumentReply, error)
	ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error)
}

// UnimplementedWatermarkServer can be embedded to have forward compatible implementations.
type UnimplementedWatermarkServer struct {
}

func (*UnimplementedWatermarkServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedWatermarkServer) Watermark(context.Context, *WatermarkRequest) (*WatermarkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watermark not implemented")
}
func (*UnimplementedWatermarkServer) Status(context.Context, *StatusRequest) (*StatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedWatermarkServer) AddDocument(context.Context, *AddDocumentRequest) (*AddDcoumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDocument not implemented")
}
func (*UnimplementedWatermarkServer) ServiceStatus(context.Context, *ServiceStatusRequest) (*ServiceStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceStatus not implemented")
}

func RegisterWatermarkServer(s *grpc.Server, srv WatermarkServer) {
	s.RegisterService(&_Watermark_serviceDesc, srv)
}

func _Watermark_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_Watermark_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WatermarkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Watermark(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Watermark",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Watermark(ctx, req.(*WatermarkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_AddDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).AddDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/AddDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).AddDocument(ctx, req.(*AddDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Watermark_ServiceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WatermarkServer).ServiceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Watermark/ServiceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WatermarkServer).ServiceStatus(ctx, req.(*ServiceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Watermark_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Watermark",
	HandlerType: (*WatermarkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Watermark_Get_Handler,
		},
		{
			MethodName: "Watermark",
			Handler:    _Watermark_Watermark_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Watermark_Status_Handler,
		},
		{
			MethodName: "AddDocument",
			Handler:    _Watermark_AddDocument_Handler,
		},
		{
			MethodName: "ServiceStatus",
			Handler:    _Watermark_ServiceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "watermarksvc.proto",
}