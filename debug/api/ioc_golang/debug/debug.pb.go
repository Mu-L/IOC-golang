// EDIT IT, change to your package, service and message

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: debug/api/ioc_golang/debug/debug.proto

package debug

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"

	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceMetadata []*ServiceMetadata `protobuf:"bytes,1,rep,name=serviceMetadata,proto3" json:"serviceMetadata,omitempty"`
}

func (x *ListServiceResponse) Reset() {
	*x = ListServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceResponse) ProtoMessage() {}

func (x *ListServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceResponse.ProtoReflect.Descriptor instead.
func (*ListServiceResponse) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{0}
}

func (x *ListServiceResponse) GetServiceMetadata() []*ServiceMetadata {
	if x != nil {
		return x.ServiceMetadata
	}
	return nil
}

type ServiceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InterfaceName      string   `protobuf:"bytes,1,opt,name=interfaceName,proto3" json:"interfaceName,omitempty"`
	ImplementationName string   `protobuf:"bytes,2,opt,name=implementationName,proto3" json:"implementationName,omitempty"`
	Methods            []string `protobuf:"bytes,3,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (x *ServiceMetadata) Reset() {
	*x = ServiceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceMetadata) ProtoMessage() {}

func (x *ServiceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceMetadata.ProtoReflect.Descriptor instead.
func (*ServiceMetadata) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceMetadata) GetInterfaceName() string {
	if x != nil {
		return x.InterfaceName
	}
	return ""
}

func (x *ServiceMetadata) GetImplementationName() string {
	if x != nil {
		return x.ImplementationName
	}
	return ""
}

func (x *ServiceMetadata) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

type WatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdid     string     `protobuf:"bytes,1,opt,name=sdid,proto3" json:"sdid,omitempty"`
	Method   string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Matchers []*Matcher `protobuf:"bytes,3,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *WatchRequest) Reset() {
	*x = WatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchRequest) ProtoMessage() {}

func (x *WatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchRequest.ProtoReflect.Descriptor instead.
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{2}
}

func (x *WatchRequest) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *WatchRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *WatchRequest) GetMatchers() []*Matcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type TraceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdid     string     `protobuf:"bytes,1,opt,name=sdid,proto3" json:"sdid,omitempty"`
	Method   string     `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Matchers []*Matcher `protobuf:"bytes,3,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *TraceRequest) Reset() {
	*x = TraceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceRequest) ProtoMessage() {}

func (x *TraceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceRequest.ProtoReflect.Descriptor instead.
func (*TraceRequest) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{3}
}

func (x *TraceRequest) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *TraceRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *TraceRequest) GetMatchers() []*Matcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index      int64  `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	MatchPath  string `protobuf:"bytes,2,opt,name=matchPath,proto3" json:"matchPath,omitempty"`
	MatchValue string `protobuf:"bytes,3,opt,name=matchValue,proto3" json:"matchValue,omitempty"`
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{4}
}

func (x *Matcher) GetIndex() int64 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Matcher) GetMatchPath() string {
	if x != nil {
		return x.MatchPath
	}
	return ""
}

func (x *Matcher) GetMatchValue() string {
	if x != nil {
		return x.MatchValue
	}
	return ""
}

type WatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdid         string   `protobuf:"bytes,1,opt,name=sdid,proto3" json:"sdid,omitempty"`
	MethodName   string   `protobuf:"bytes,2,opt,name=methodName,proto3" json:"methodName,omitempty"`
	Params       []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	ReturnValues []string `protobuf:"bytes,4,rep,name=returnValues,proto3" json:"returnValues,omitempty"`
}

func (x *WatchResponse) Reset() {
	*x = WatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchResponse) ProtoMessage() {}

func (x *WatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchResponse.ProtoReflect.Descriptor instead.
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{5}
}

func (x *WatchResponse) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *WatchResponse) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *WatchResponse) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *WatchResponse) GetReturnValues() []string {
	if x != nil {
		return x.ReturnValues
	}
	return nil
}

type TraceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectorAddress string `protobuf:"bytes,1,opt,name=collectorAddress,proto3" json:"collectorAddress,omitempty"`
}

func (x *TraceResponse) Reset() {
	*x = TraceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TraceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceResponse) ProtoMessage() {}

func (x *TraceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceResponse.ProtoReflect.Descriptor instead.
func (*TraceResponse) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{6}
}

func (x *TraceResponse) GetCollectorAddress() string {
	if x != nil {
		return x.CollectorAddress
	}
	return ""
}

// todo return who span links, now it's usless
type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sdid         string   `protobuf:"bytes,1,opt,name=sdid,proto3" json:"sdid,omitempty"`
	MethodName   string   `protobuf:"bytes,2,opt,name=methodName,proto3" json:"methodName,omitempty"`
	Params       []string `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty"`
	ReturnValues []string `protobuf:"bytes,4,rep,name=returnValues,proto3" json:"returnValues,omitempty"`
	StartTime    int64    `protobuf:"varint,7,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime      int64    `protobuf:"varint,8,opt,name=endTime,proto3" json:"endTime,omitempty"`
	ChildSpans   []*Span  `protobuf:"bytes,9,rep,name=childSpans,proto3" json:"childSpans,omitempty"`
	ParentSpan   *Span    `protobuf:"bytes,10,opt,name=parentSpan,proto3" json:"parentSpan,omitempty"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_debug_api_ioc_golang_debug_debug_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP(), []int{7}
}

func (x *Span) GetSdid() string {
	if x != nil {
		return x.Sdid
	}
	return ""
}

func (x *Span) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *Span) GetParams() []string {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *Span) GetReturnValues() []string {
	if x != nil {
		return x.ReturnValues
	}
	return nil
}

func (x *Span) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Span) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *Span) GetChildSpans() []*Span {
	if x != nil {
		return x.ChildSpans
	}
	return nil
}

func (x *Span) GetParentSpan() *Span {
	if x != nil {
		return x.ParentSpan
	}
	return nil
}

var File_debug_api_ioc_golang_debug_debug_proto protoreflect.FileDescriptor

var file_debug_api_ioc_golang_debug_debug_proto_rawDesc = []byte{
	0x0a, 0x26, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6f, 0x63, 0x5f,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x0f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22,
	0x71, 0x0a, 0x0c, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x73, 0x22, 0x71, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x35,
	0x0a, 0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x08, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x5d, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x50,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x7f, 0x0a, 0x0d, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x64, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x9e, 0x02, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x64, 0x69, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x61, 0x6e,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52,
	0x0a, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x0a, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x61, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62,
	0x75, 0x67, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x70, 0x61, 0x6e, 0x32, 0xfb, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1e, 0x2e,
	0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67,
	0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x4c, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x69, 0x6f,
	0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f,
	0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x4f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x25, 0x2e, 0x69, 0x6f, 0x63, 0x5f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x12, 0x5a, 0x10, 0x69, 0x6f, 0x63, 0x5f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x64, 0x65, 0x62, 0x75, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_debug_api_ioc_golang_debug_debug_proto_rawDescOnce sync.Once
	file_debug_api_ioc_golang_debug_debug_proto_rawDescData = file_debug_api_ioc_golang_debug_debug_proto_rawDesc
)

func file_debug_api_ioc_golang_debug_debug_proto_rawDescGZIP() []byte {
	file_debug_api_ioc_golang_debug_debug_proto_rawDescOnce.Do(func() {
		file_debug_api_ioc_golang_debug_debug_proto_rawDescData = protoimpl.X.CompressGZIP(file_debug_api_ioc_golang_debug_debug_proto_rawDescData)
	})
	return file_debug_api_ioc_golang_debug_debug_proto_rawDescData
}

var file_debug_api_ioc_golang_debug_debug_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_debug_api_ioc_golang_debug_debug_proto_goTypes = []interface{}{
	(*ListServiceResponse)(nil), // 0: ioc_golang.debug.ListServiceResponse
	(*ServiceMetadata)(nil),     // 1: ioc_golang.debug.ServiceMetadata
	(*WatchRequest)(nil),        // 2: ioc_golang.debug.WatchRequest
	(*TraceRequest)(nil),        // 3: ioc_golang.debug.TraceRequest
	(*Matcher)(nil),             // 4: ioc_golang.debug.Matcher
	(*WatchResponse)(nil),       // 5: ioc_golang.debug.WatchResponse
	(*TraceResponse)(nil),       // 6: ioc_golang.debug.TraceResponse
	(*Span)(nil),                // 7: ioc_golang.debug.Span
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_debug_api_ioc_golang_debug_debug_proto_depIdxs = []int32{
	1, // 0: ioc_golang.debug.ListServiceResponse.serviceMetadata:type_name -> ioc_golang.debug.ServiceMetadata
	4, // 1: ioc_golang.debug.WatchRequest.matchers:type_name -> ioc_golang.debug.Matcher
	4, // 2: ioc_golang.debug.TraceRequest.matchers:type_name -> ioc_golang.debug.Matcher
	7, // 3: ioc_golang.debug.Span.childSpans:type_name -> ioc_golang.debug.Span
	7, // 4: ioc_golang.debug.Span.parentSpan:type_name -> ioc_golang.debug.Span
	2, // 5: ioc_golang.debug.DebugService.Watch:input_type -> ioc_golang.debug.WatchRequest
	3, // 6: ioc_golang.debug.DebugService.Trace:input_type -> ioc_golang.debug.TraceRequest
	8, // 7: ioc_golang.debug.DebugService.ListServices:input_type -> google.protobuf.Empty
	5, // 8: ioc_golang.debug.DebugService.Watch:output_type -> ioc_golang.debug.WatchResponse
	6, // 9: ioc_golang.debug.DebugService.Trace:output_type -> ioc_golang.debug.TraceResponse
	0, // 10: ioc_golang.debug.DebugService.ListServices:output_type -> ioc_golang.debug.ListServiceResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_debug_api_ioc_golang_debug_debug_proto_init() }
func file_debug_api_ioc_golang_debug_debug_proto_init() {
	if File_debug_api_ioc_golang_debug_debug_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceResponse); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceMetadata); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchRequest); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceRequest); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchResponse); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TraceResponse); i {
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
		file_debug_api_ioc_golang_debug_debug_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
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
			RawDescriptor: file_debug_api_ioc_golang_debug_debug_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_debug_api_ioc_golang_debug_debug_proto_goTypes,
		DependencyIndexes: file_debug_api_ioc_golang_debug_debug_proto_depIdxs,
		MessageInfos:      file_debug_api_ioc_golang_debug_debug_proto_msgTypes,
	}.Build()
	File_debug_api_ioc_golang_debug_debug_proto = out.File
	file_debug_api_ioc_golang_debug_debug_proto_rawDesc = nil
	file_debug_api_ioc_golang_debug_debug_proto_goTypes = nil
	file_debug_api_ioc_golang_debug_debug_proto_depIdxs = nil
}
