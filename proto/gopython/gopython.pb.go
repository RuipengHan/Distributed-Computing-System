// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/gopython/gopython.proto

package gopython

import (
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

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchSize int32  `protobuf:"varint,1,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	ModelType string `protobuf:"bytes,2,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	JobId     int32  `protobuf:"varint,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gopython_gopython_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gopython_gopython_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_proto_gopython_gopython_proto_rawDescGZIP(), []int{0}
}

func (x *InitializeRequest) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *InitializeRequest) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *InitializeRequest) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gopython_gopython_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gopython_gopython_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_proto_gopython_gopython_proto_rawDescGZIP(), []int{1}
}

func (x *InitializeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type InferenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId               int32  `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	BatchId             int32  `protobuf:"varint,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	InferenceSize       int32  `protobuf:"varint,3,opt,name=inference_size,json=inferenceSize,proto3" json:"inference_size,omitempty"`
	InferenceDataFolder string `protobuf:"bytes,4,opt,name=inference_data_folder,json=inferenceDataFolder,proto3" json:"inference_data_folder,omitempty"`
}

func (x *InferenceRequest) Reset() {
	*x = InferenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gopython_gopython_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceRequest) ProtoMessage() {}

func (x *InferenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gopython_gopython_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceRequest.ProtoReflect.Descriptor instead.
func (*InferenceRequest) Descriptor() ([]byte, []int) {
	return file_proto_gopython_gopython_proto_rawDescGZIP(), []int{2}
}

func (x *InferenceRequest) GetJobId() int32 {
	if x != nil {
		return x.JobId
	}
	return 0
}

func (x *InferenceRequest) GetBatchId() int32 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *InferenceRequest) GetInferenceSize() int32 {
	if x != nil {
		return x.InferenceSize
	}
	return 0
}

func (x *InferenceRequest) GetInferenceDataFolder() string {
	if x != nil {
		return x.InferenceDataFolder
	}
	return ""
}

type InferenceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	InferenceResult []byte `protobuf:"bytes,2,opt,name=inference_result,json=inferenceResult,proto3" json:"inference_result,omitempty"`
}

func (x *InferenceResponse) Reset() {
	*x = InferenceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gopython_gopython_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InferenceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InferenceResponse) ProtoMessage() {}

func (x *InferenceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gopython_gopython_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InferenceResponse.ProtoReflect.Descriptor instead.
func (*InferenceResponse) Descriptor() ([]byte, []int) {
	return file_proto_gopython_gopython_proto_rawDescGZIP(), []int{3}
}

func (x *InferenceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InferenceResponse) GetInferenceResult() []byte {
	if x != nil {
		return x.InferenceResult
	}
	return nil
}

var File_proto_gopython_gopython_proto protoreflect.FileDescriptor

var file_proto_gopython_gopython_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x2f, 0x67, 0x6f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x67, 0x6f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x22, 0x68, 0x0a, 0x11, 0x49, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f,
	0x62, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x12, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x32, 0x0a, 0x15, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x46, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x22, 0x56, 0x0a, 0x11, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x69, 0x6e, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa7, 0x01, 0x0a, 0x08,
	0x47, 0x6f, 0x50, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x0f, 0x49, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x6f,
	0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x70, 0x79, 0x74,
	0x68, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x6f, 0x70,
	0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x6f, 0x70, 0x79, 0x74, 0x68, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x63, 0x73, 0x34, 0x32, 0x35, 0x6d, 0x70,
	0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gopython_gopython_proto_rawDescOnce sync.Once
	file_proto_gopython_gopython_proto_rawDescData = file_proto_gopython_gopython_proto_rawDesc
)

func file_proto_gopython_gopython_proto_rawDescGZIP() []byte {
	file_proto_gopython_gopython_proto_rawDescOnce.Do(func() {
		file_proto_gopython_gopython_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gopython_gopython_proto_rawDescData)
	})
	return file_proto_gopython_gopython_proto_rawDescData
}

var file_proto_gopython_gopython_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_gopython_gopython_proto_goTypes = []interface{}{
	(*InitializeRequest)(nil),  // 0: gopython.InitializeRequest
	(*InitializeResponse)(nil), // 1: gopython.InitializeResponse
	(*InferenceRequest)(nil),   // 2: gopython.InferenceRequest
	(*InferenceResponse)(nil),  // 3: gopython.InferenceResponse
}
var file_proto_gopython_gopython_proto_depIdxs = []int32{
	0, // 0: gopython.GoPython.InitializeModel:input_type -> gopython.InitializeRequest
	2, // 1: gopython.GoPython.ModelInference:input_type -> gopython.InferenceRequest
	1, // 2: gopython.GoPython.InitializeModel:output_type -> gopython.InitializeResponse
	3, // 3: gopython.GoPython.ModelInference:output_type -> gopython.InferenceResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_gopython_gopython_proto_init() }
func file_proto_gopython_gopython_proto_init() {
	if File_proto_gopython_gopython_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gopython_gopython_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_proto_gopython_gopython_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
		file_proto_gopython_gopython_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceRequest); i {
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
		file_proto_gopython_gopython_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InferenceResponse); i {
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
			RawDescriptor: file_proto_gopython_gopython_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gopython_gopython_proto_goTypes,
		DependencyIndexes: file_proto_gopython_gopython_proto_depIdxs,
		MessageInfos:      file_proto_gopython_gopython_proto_msgTypes,
	}.Build()
	File_proto_gopython_gopython_proto = out.File
	file_proto_gopython_gopython_proto_rawDesc = nil
	file_proto_gopython_gopython_proto_goTypes = nil
	file_proto_gopython_gopython_proto_depIdxs = nil
}