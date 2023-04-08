// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: obd/v1/car/update_rpc.proto

package obdv1

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

type CarUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId           int64 `protobuf:"varint,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	CarBrandModelId int64 `protobuf:"varint,2,opt,name=car_brand_model_id,json=carBrandModelId,proto3" json:"car_brand_model_id,omitempty"`
	ModelYear       int32 `protobuf:"varint,3,opt,name=model_year,json=modelYear,proto3" json:"model_year,omitempty"`
}

func (x *CarUpdateRequest) Reset() {
	*x = CarUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_obd_v1_car_update_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarUpdateRequest) ProtoMessage() {}

func (x *CarUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_obd_v1_car_update_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarUpdateRequest.ProtoReflect.Descriptor instead.
func (*CarUpdateRequest) Descriptor() ([]byte, []int) {
	return file_obd_v1_car_update_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *CarUpdateRequest) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *CarUpdateRequest) GetCarBrandModelId() int64 {
	if x != nil {
		return x.CarBrandModelId
	}
	return 0
}

func (x *CarUpdateRequest) GetModelYear() int32 {
	if x != nil {
		return x.ModelYear
	}
	return 0
}

type CarUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CarUpdateResponse) Reset() {
	*x = CarUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_obd_v1_car_update_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarUpdateResponse) ProtoMessage() {}

func (x *CarUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_obd_v1_car_update_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarUpdateResponse.ProtoReflect.Descriptor instead.
func (*CarUpdateResponse) Descriptor() ([]byte, []int) {
	return file_obd_v1_car_update_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *CarUpdateResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

var File_obd_v1_car_update_rpc_proto protoreflect.FileDescriptor

var file_obd_v1_car_update_rpc_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f,
	0x62, 0x64, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x72, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x75,
	0x0a, 0x10, 0x43, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x12, 0x63, 0x61, 0x72,
	0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x59, 0x65, 0x61, 0x72, 0x22, 0x32, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x63, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x42, 0x8c, 0x01, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x72, 0x77, 0x69, 0x73, 0x68, 0x64, 0x65,
	0x76, 0x2f, 0x6f, 0x62, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62,
	0x2f, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x3b, 0x6f, 0x62, 0x64, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4f, 0x62, 0x64, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x06, 0x4f, 0x62, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x4f, 0x62, 0x64, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x07, 0x4f, 0x62, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_obd_v1_car_update_rpc_proto_rawDescOnce sync.Once
	file_obd_v1_car_update_rpc_proto_rawDescData = file_obd_v1_car_update_rpc_proto_rawDesc
)

func file_obd_v1_car_update_rpc_proto_rawDescGZIP() []byte {
	file_obd_v1_car_update_rpc_proto_rawDescOnce.Do(func() {
		file_obd_v1_car_update_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_obd_v1_car_update_rpc_proto_rawDescData)
	})
	return file_obd_v1_car_update_rpc_proto_rawDescData
}

var file_obd_v1_car_update_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_obd_v1_car_update_rpc_proto_goTypes = []interface{}{
	(*CarUpdateRequest)(nil),  // 0: obd.v1.CarUpdateRequest
	(*CarUpdateResponse)(nil), // 1: obd.v1.CarUpdateResponse
	(*Car)(nil),               // 2: obd.v1.Car
}
var file_obd_v1_car_update_rpc_proto_depIdxs = []int32{
	2, // 0: obd.v1.CarUpdateResponse.car:type_name -> obd.v1.Car
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_obd_v1_car_update_rpc_proto_init() }
func file_obd_v1_car_update_rpc_proto_init() {
	if File_obd_v1_car_update_rpc_proto != nil {
		return
	}
	file_obd_v1_car_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_obd_v1_car_update_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarUpdateRequest); i {
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
		file_obd_v1_car_update_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarUpdateResponse); i {
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
			RawDescriptor: file_obd_v1_car_update_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_obd_v1_car_update_rpc_proto_goTypes,
		DependencyIndexes: file_obd_v1_car_update_rpc_proto_depIdxs,
		MessageInfos:      file_obd_v1_car_update_rpc_proto_msgTypes,
	}.Build()
	File_obd_v1_car_update_rpc_proto = out.File
	file_obd_v1_car_update_rpc_proto_rawDesc = nil
	file_obd_v1_car_update_rpc_proto_goTypes = nil
	file_obd_v1_car_update_rpc_proto_depIdxs = nil
}