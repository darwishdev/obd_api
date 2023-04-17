// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: obd/v1/obd_service.proto

package obdv1

import (
	car "github.com/darwishdev/obd_api/pkg/pb/obd/v1/car"
	center "github.com/darwishdev/obd_api/pkg/pb/obd/v1/center"
	review "github.com/darwishdev/obd_api/pkg/pb/obd/v1/review"
	session "github.com/darwishdev/obd_api/pkg/pb/obd/v1/session"
	user "github.com/darwishdev/obd_api/pkg/pb/obd/v1/user"
	winch "github.com/darwishdev/obd_api/pkg/pb/obd/v1/winch"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_obd_v1_obd_service_proto protoreflect.FileDescriptor

var file_obd_v1_obd_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x62, 0x64, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x62, 0x64, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6f, 0x62,
	0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x6c, 0x69, 0x73, 0x74,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x62, 0x64, 0x2f,
	0x76, 0x31, 0x2f, 0x77, 0x69, 0x6e, 0x63, 0x68, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61,
	0x72, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x2f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x2f, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1d, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xbd,
	0x09, 0x0a, 0x03, 0x4f, 0x62, 0x64, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a,
	0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x18, 0x2e, 0x6f, 0x62, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x45, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x62, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1d, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09,
	0x43, 0x61, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x6f, 0x62, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e,
	0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42,
	0x0a, 0x09, 0x57, 0x69, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x6f, 0x62,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x69, 0x6e, 0x63, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x48, 0x0a, 0x0b, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1a, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x6f,
	0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x62, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x6f, 0x62, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x11, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20,
	0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x2e, 0x6f, 0x62,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x74, 0x74, 0x61,
	0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x4b, 0x0a, 0x0c, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x1b, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x89,
	0x01, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x62, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x4f,
	0x62, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x72,
	0x77, 0x69, 0x73, 0x68, 0x64, 0x65, 0x76, 0x2f, 0x6f, 0x62, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x6f, 0x62, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x62,
	0x64, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x4f, 0x62, 0x64, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x06, 0x4f, 0x62, 0x64, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x12, 0x4f, 0x62,
	0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x07, 0x4f, 0x62, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var file_obd_v1_obd_service_proto_goTypes = []interface{}{
	(*user.UserCreateRequest)(nil),             // 0: obd.v1.UserCreateRequest
	(*user.UserLoginRequest)(nil),              // 1: obd.v1.UserLoginRequest
	(*user.UserUpdateRequest)(nil),             // 2: obd.v1.UserUpdateRequest
	(*emptypb.Empty)(nil),                      // 3: google.protobuf.Empty
	(*car.BrandsListRequest)(nil),              // 4: obd.v1.BrandsListRequest
	(*car.CarCreateRequest)(nil),               // 5: obd.v1.CarCreateRequest
	(*car.CarUpdateRequest)(nil),               // 6: obd.v1.CarUpdateRequest
	(*center.CentersListRequest)(nil),          // 7: obd.v1.CentersListRequest
	(*winch.WinchListRequest)(nil),             // 8: obd.v1.WinchListRequest
	(*review.ReviewsListRequest)(nil),          // 9: obd.v1.ReviewsListRequest
	(*review.ReviewCreateRequest)(nil),         // 10: obd.v1.ReviewCreateRequest
	(*session.SessionCreateRequest)(nil),       // 11: obd.v1.SessionCreateRequest
	(*session.SessionAttachCodeRequest)(nil),   // 12: obd.v1.SessionAttachCodeRequest
	(*session.SessionAttachValueRequest)(nil),  // 13: obd.v1.SessionAttachValueRequest
	(*session.SessionsListRequest)(nil),        // 14: obd.v1.SessionsListRequest
	(*session.SessionCloseRequest)(nil),        // 15: obd.v1.SessionCloseRequest
	(*user.UserCreateResponse)(nil),            // 16: obd.v1.UserCreateResponse
	(*user.UserLoginResponse)(nil),             // 17: obd.v1.UserLoginResponse
	(*user.UserUpdateResponse)(nil),            // 18: obd.v1.UserUpdateResponse
	(*user.UserAuthorizeResponse)(nil),         // 19: obd.v1.UserAuthorizeResponse
	(*car.BrandsListResponse)(nil),             // 20: obd.v1.BrandsListResponse
	(*car.CarCreateResponse)(nil),              // 21: obd.v1.CarCreateResponse
	(*car.CarUpdateResponse)(nil),              // 22: obd.v1.CarUpdateResponse
	(*center.CentersListResponse)(nil),         // 23: obd.v1.CentersListResponse
	(*winch.WinchListResponse)(nil),            // 24: obd.v1.WinchListResponse
	(*review.ReviewsListResponse)(nil),         // 25: obd.v1.ReviewsListResponse
	(*review.ReviewCreateResponse)(nil),        // 26: obd.v1.ReviewCreateResponse
	(*session.SessionCreateResponse)(nil),      // 27: obd.v1.SessionCreateResponse
	(*session.SessionAttachCodeResponse)(nil),  // 28: obd.v1.SessionAttachCodeResponse
	(*session.SessionAttachValueResponse)(nil), // 29: obd.v1.SessionAttachValueResponse
	(*session.SessionsListResponse)(nil),       // 30: obd.v1.SessionsListResponse
	(*session.SessionCloseResponse)(nil),       // 31: obd.v1.SessionCloseResponse
}
var file_obd_v1_obd_service_proto_depIdxs = []int32{
	0,  // 0: obd.v1.Obd.UserCreate:input_type -> obd.v1.UserCreateRequest
	1,  // 1: obd.v1.Obd.UserLogin:input_type -> obd.v1.UserLoginRequest
	2,  // 2: obd.v1.Obd.UserUpdate:input_type -> obd.v1.UserUpdateRequest
	3,  // 3: obd.v1.Obd.UserAuthorize:input_type -> google.protobuf.Empty
	4,  // 4: obd.v1.Obd.CarBrandsList:input_type -> obd.v1.BrandsListRequest
	5,  // 5: obd.v1.Obd.CarCreate:input_type -> obd.v1.CarCreateRequest
	6,  // 6: obd.v1.Obd.CarUpdate:input_type -> obd.v1.CarUpdateRequest
	7,  // 7: obd.v1.Obd.CentersList:input_type -> obd.v1.CentersListRequest
	8,  // 8: obd.v1.Obd.WinchList:input_type -> obd.v1.WinchListRequest
	9,  // 9: obd.v1.Obd.ReviewsList:input_type -> obd.v1.ReviewsListRequest
	10, // 10: obd.v1.Obd.ReviewCreate:input_type -> obd.v1.ReviewCreateRequest
	11, // 11: obd.v1.Obd.SessionCreate:input_type -> obd.v1.SessionCreateRequest
	12, // 12: obd.v1.Obd.SessionAttachCode:input_type -> obd.v1.SessionAttachCodeRequest
	13, // 13: obd.v1.Obd.SessionAttachValue:input_type -> obd.v1.SessionAttachValueRequest
	14, // 14: obd.v1.Obd.SessionsList:input_type -> obd.v1.SessionsListRequest
	15, // 15: obd.v1.Obd.SessionClose:input_type -> obd.v1.SessionCloseRequest
	16, // 16: obd.v1.Obd.UserCreate:output_type -> obd.v1.UserCreateResponse
	17, // 17: obd.v1.Obd.UserLogin:output_type -> obd.v1.UserLoginResponse
	18, // 18: obd.v1.Obd.UserUpdate:output_type -> obd.v1.UserUpdateResponse
	19, // 19: obd.v1.Obd.UserAuthorize:output_type -> obd.v1.UserAuthorizeResponse
	20, // 20: obd.v1.Obd.CarBrandsList:output_type -> obd.v1.BrandsListResponse
	21, // 21: obd.v1.Obd.CarCreate:output_type -> obd.v1.CarCreateResponse
	22, // 22: obd.v1.Obd.CarUpdate:output_type -> obd.v1.CarUpdateResponse
	23, // 23: obd.v1.Obd.CentersList:output_type -> obd.v1.CentersListResponse
	24, // 24: obd.v1.Obd.WinchList:output_type -> obd.v1.WinchListResponse
	25, // 25: obd.v1.Obd.ReviewsList:output_type -> obd.v1.ReviewsListResponse
	26, // 26: obd.v1.Obd.ReviewCreate:output_type -> obd.v1.ReviewCreateResponse
	27, // 27: obd.v1.Obd.SessionCreate:output_type -> obd.v1.SessionCreateResponse
	28, // 28: obd.v1.Obd.SessionAttachCode:output_type -> obd.v1.SessionAttachCodeResponse
	29, // 29: obd.v1.Obd.SessionAttachValue:output_type -> obd.v1.SessionAttachValueResponse
	30, // 30: obd.v1.Obd.SessionsList:output_type -> obd.v1.SessionsListResponse
	31, // 31: obd.v1.Obd.SessionClose:output_type -> obd.v1.SessionCloseResponse
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_obd_v1_obd_service_proto_init() }
func file_obd_v1_obd_service_proto_init() {
	if File_obd_v1_obd_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_obd_v1_obd_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_obd_v1_obd_service_proto_goTypes,
		DependencyIndexes: file_obd_v1_obd_service_proto_depIdxs,
	}.Build()
	File_obd_v1_obd_service_proto = out.File
	file_obd_v1_obd_service_proto_rawDesc = nil
	file_obd_v1_obd_service_proto_goTypes = nil
	file_obd_v1_obd_service_proto_depIdxs = nil
}
