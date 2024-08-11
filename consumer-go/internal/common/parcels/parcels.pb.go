// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: parcels.proto

package parcels

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ParcelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingCode string `protobuf:"bytes,1,opt,name=trackingCode,proto3" json:"trackingCode,omitempty"`
}

func (x *ParcelReq) Reset() {
	*x = ParcelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parcels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParcelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParcelReq) ProtoMessage() {}

func (x *ParcelReq) ProtoReflect() protoreflect.Message {
	mi := &file_parcels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParcelReq.ProtoReflect.Descriptor instead.
func (*ParcelReq) Descriptor() ([]byte, []int) {
	return file_parcels_proto_rawDescGZIP(), []int{0}
}

func (x *ParcelReq) GetTrackingCode() string {
	if x != nil {
		return x.TrackingCode
	}
	return ""
}

type ParcelMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrackingCode string  `protobuf:"bytes,2,opt,name=trackingCode,proto3" json:"trackingCode,omitempty"`
	Uuid         *string `protobuf:"bytes,1,opt,name=uuid,proto3,oneof" json:"uuid,omitempty"`
	Name         *string `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Origin       *string `protobuf:"bytes,4,opt,name=origin,proto3,oneof" json:"origin,omitempty"`
	Destination  *string `protobuf:"bytes,5,opt,name=destination,proto3,oneof" json:"destination,omitempty"`
	LastUpdate   *string `protobuf:"bytes,6,opt,name=lastUpdate,proto3,oneof" json:"lastUpdate,omitempty"`
	Status       *string `protobuf:"bytes,7,opt,name=status,proto3,oneof" json:"status,omitempty"`
	ZipCode      *string `protobuf:"bytes,8,opt,name=zipCode,proto3,oneof" json:"zipCode,omitempty"`
	IsDone       *bool   `protobuf:"varint,9,opt,name=isDone,proto3,oneof" json:"isDone,omitempty"`
}

func (x *ParcelMessage) Reset() {
	*x = ParcelMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parcels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParcelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParcelMessage) ProtoMessage() {}

func (x *ParcelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_parcels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParcelMessage.ProtoReflect.Descriptor instead.
func (*ParcelMessage) Descriptor() ([]byte, []int) {
	return file_parcels_proto_rawDescGZIP(), []int{1}
}

func (x *ParcelMessage) GetTrackingCode() string {
	if x != nil {
		return x.TrackingCode
	}
	return ""
}

func (x *ParcelMessage) GetUuid() string {
	if x != nil && x.Uuid != nil {
		return *x.Uuid
	}
	return ""
}

func (x *ParcelMessage) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ParcelMessage) GetOrigin() string {
	if x != nil && x.Origin != nil {
		return *x.Origin
	}
	return ""
}

func (x *ParcelMessage) GetDestination() string {
	if x != nil && x.Destination != nil {
		return *x.Destination
	}
	return ""
}

func (x *ParcelMessage) GetLastUpdate() string {
	if x != nil && x.LastUpdate != nil {
		return *x.LastUpdate
	}
	return ""
}

func (x *ParcelMessage) GetStatus() string {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return ""
}

func (x *ParcelMessage) GetZipCode() string {
	if x != nil && x.ZipCode != nil {
		return *x.ZipCode
	}
	return ""
}

func (x *ParcelMessage) GetIsDone() bool {
	if x != nil && x.IsDone != nil {
		return *x.IsDone
	}
	return false
}

var File_parcels_proto protoreflect.FileDescriptor

var file_parcels_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x09, 0x50, 0x61,
	0x72, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x85, 0x03, 0x0a, 0x0d,
	0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x17, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x88, 0x01, 0x01,
	0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x7a, 0x69, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x07, 0x7a, 0x69,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x69, 0x73, 0x44, 0x6f,
	0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x48, 0x07, 0x52, 0x06, 0x69, 0x73, 0x44, 0x6f,
	0x6e, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0a, 0x0a, 0x08,
	0x5f, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x69, 0x73, 0x44,
	0x6f, 0x6e, 0x65, 0x32, 0xe8, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x12,
	0x45, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x72,
	0x63, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x63, 0x65, 0x6c, 0x42, 0x79, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73,
	0x2e, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a, 0x53, 0x61, 0x76,
	0x65, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x2e, 0x50, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x3c,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x2e, 0x6c, 0x69, 0x62, 0x42, 0x0c, 0x50, 0x61,
	0x72, 0x63, 0x65, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x0e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parcels_proto_rawDescOnce sync.Once
	file_parcels_proto_rawDescData = file_parcels_proto_rawDesc
)

func file_parcels_proto_rawDescGZIP() []byte {
	file_parcels_proto_rawDescOnce.Do(func() {
		file_parcels_proto_rawDescData = protoimpl.X.CompressGZIP(file_parcels_proto_rawDescData)
	})
	return file_parcels_proto_rawDescData
}

var file_parcels_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_parcels_proto_goTypes = []interface{}{
	(*ParcelReq)(nil),     // 0: grpc.parcels.ParcelReq
	(*ParcelMessage)(nil), // 1: grpc.parcels.ParcelMessage
	(*emptypb.Empty)(nil), // 2: google.protobuf.Empty
}
var file_parcels_proto_depIdxs = []int32{
	2, // 0: grpc.parcels.Parcels.GetParcels:input_type -> google.protobuf.Empty
	0, // 1: grpc.parcels.Parcels.GetParcelByTrackingCode:input_type -> grpc.parcels.ParcelReq
	1, // 2: grpc.parcels.Parcels.SaveParcel:input_type -> grpc.parcels.ParcelMessage
	1, // 3: grpc.parcels.Parcels.GetParcels:output_type -> grpc.parcels.ParcelMessage
	1, // 4: grpc.parcels.Parcels.GetParcelByTrackingCode:output_type -> grpc.parcels.ParcelMessage
	2, // 5: grpc.parcels.Parcels.SaveParcel:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_parcels_proto_init() }
func file_parcels_proto_init() {
	if File_parcels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parcels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParcelReq); i {
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
		file_parcels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParcelMessage); i {
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
	file_parcels_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parcels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parcels_proto_goTypes,
		DependencyIndexes: file_parcels_proto_depIdxs,
		MessageInfos:      file_parcels_proto_msgTypes,
	}.Build()
	File_parcels_proto = out.File
	file_parcels_proto_rawDesc = nil
	file_parcels_proto_goTypes = nil
	file_parcels_proto_depIdxs = nil
}
