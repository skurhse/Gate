// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: services/brokerage/api/src/resolution.proto

package api

import (
	api "github.com/mrNobody95/Gate/api"
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

type GetResolutionByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"id"
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"id"`
}

func (x *GetResolutionByIDRequest) Reset() {
	*x = GetResolutionByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_resolution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResolutionByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResolutionByIDRequest) ProtoMessage() {}

func (x *GetResolutionByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_resolution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResolutionByIDRequest.ProtoReflect.Descriptor instead.
func (*GetResolutionByIDRequest) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_resolution_proto_rawDescGZIP(), []int{0}
}

func (x *GetResolutionByIDRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

type GetResolutionByDurationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"brokerage_id"
	BrokerageID string `protobuf:"bytes,1,opt,name=BrokerageID,proto3" json:"brokerage_id"`
	// @inject_tag: json:"duration"
	Duration int64 `protobuf:"varint,2,opt,name=Duration,proto3" json:"duration"`
}

func (x *GetResolutionByDurationRequest) Reset() {
	*x = GetResolutionByDurationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_resolution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResolutionByDurationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResolutionByDurationRequest) ProtoMessage() {}

func (x *GetResolutionByDurationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_resolution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResolutionByDurationRequest.ProtoReflect.Descriptor instead.
func (*GetResolutionByDurationRequest) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_resolution_proto_rawDescGZIP(), []int{1}
}

func (x *GetResolutionByDurationRequest) GetBrokerageID() string {
	if x != nil {
		return x.BrokerageID
	}
	return ""
}

func (x *GetResolutionByDurationRequest) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type GetResolutionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: json:"brokerage_id"
	BrokerageID string `protobuf:"bytes,1,opt,name=BrokerageID,proto3" json:"brokerage_id"`
}

func (x *GetResolutionListRequest) Reset() {
	*x = GetResolutionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_brokerage_api_src_resolution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResolutionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResolutionListRequest) ProtoMessage() {}

func (x *GetResolutionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_brokerage_api_src_resolution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResolutionListRequest.ProtoReflect.Descriptor instead.
func (*GetResolutionListRequest) Descriptor() ([]byte, []int) {
	return file_services_brokerage_api_src_resolution_proto_rawDescGZIP(), []int{2}
}

func (x *GetResolutionListRequest) GetBrokerageID() string {
	if x != nil {
		return x.BrokerageID
	}
	return ""
}

var File_services_brokerage_api_src_resolution_proto protoreflect.FileDescriptor

var file_services_brokerage_api_src_resolution_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x1a, 0x12, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x69, 0x73, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x5e, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x3c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x0b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x49, 0x44, 0x32,
	0x8c, 0x02, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x0f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x09, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x42, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65,
	0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x72, 0x4e,
	0x6f, 0x62, 0x6f, 0x64, 0x79, 0x39, 0x35, 0x2f, 0x47, 0x61, 0x74, 0x65, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x61, 0x67, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_brokerage_api_src_resolution_proto_rawDescOnce sync.Once
	file_services_brokerage_api_src_resolution_proto_rawDescData = file_services_brokerage_api_src_resolution_proto_rawDesc
)

func file_services_brokerage_api_src_resolution_proto_rawDescGZIP() []byte {
	file_services_brokerage_api_src_resolution_proto_rawDescOnce.Do(func() {
		file_services_brokerage_api_src_resolution_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_brokerage_api_src_resolution_proto_rawDescData)
	})
	return file_services_brokerage_api_src_resolution_proto_rawDescData
}

var file_services_brokerage_api_src_resolution_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_services_brokerage_api_src_resolution_proto_goTypes = []interface{}{
	(*GetResolutionByIDRequest)(nil),       // 0: brokerageApi.GetResolutionByIDRequest
	(*GetResolutionByDurationRequest)(nil), // 1: brokerageApi.GetResolutionByDurationRequest
	(*GetResolutionListRequest)(nil),       // 2: brokerageApi.GetResolutionListRequest
	(*api.Resolution)(nil),                 // 3: api.Resolution
	(*api.Void)(nil),                       // 4: api.Void
	(*api.Resolutions)(nil),                // 5: api.Resolutions
}
var file_services_brokerage_api_src_resolution_proto_depIdxs = []int32{
	3, // 0: brokerageApi.ResolutionService.Set:input_type -> api.Resolution
	0, // 1: brokerageApi.ResolutionService.GetByID:input_type -> brokerageApi.GetResolutionByIDRequest
	1, // 2: brokerageApi.ResolutionService.GetByDuration:input_type -> brokerageApi.GetResolutionByDurationRequest
	2, // 3: brokerageApi.ResolutionService.List:input_type -> brokerageApi.GetResolutionListRequest
	4, // 4: brokerageApi.ResolutionService.Set:output_type -> api.Void
	3, // 5: brokerageApi.ResolutionService.GetByID:output_type -> api.Resolution
	3, // 6: brokerageApi.ResolutionService.GetByDuration:output_type -> api.Resolution
	5, // 7: brokerageApi.ResolutionService.List:output_type -> api.Resolutions
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_brokerage_api_src_resolution_proto_init() }
func file_services_brokerage_api_src_resolution_proto_init() {
	if File_services_brokerage_api_src_resolution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_brokerage_api_src_resolution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResolutionByIDRequest); i {
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
		file_services_brokerage_api_src_resolution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResolutionByDurationRequest); i {
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
		file_services_brokerage_api_src_resolution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResolutionListRequest); i {
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
			RawDescriptor: file_services_brokerage_api_src_resolution_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_brokerage_api_src_resolution_proto_goTypes,
		DependencyIndexes: file_services_brokerage_api_src_resolution_proto_depIdxs,
		MessageInfos:      file_services_brokerage_api_src_resolution_proto_msgTypes,
	}.Build()
	File_services_brokerage_api_src_resolution_proto = out.File
	file_services_brokerage_api_src_resolution_proto_rawDesc = nil
	file_services_brokerage_api_src_resolution_proto_goTypes = nil
	file_services_brokerage_api_src_resolution_proto_depIdxs = nil
}
