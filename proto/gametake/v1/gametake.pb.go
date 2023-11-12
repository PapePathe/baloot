// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: proto/gametake/v1/gametake.proto

package proto

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gametake_v1_gametake_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gametake_v1_gametake_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_proto_gametake_v1_gametake_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Card) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type RecommendGameTakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cards []*Card `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
}

func (x *RecommendGameTakeRequest) Reset() {
	*x = RecommendGameTakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gametake_v1_gametake_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendGameTakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendGameTakeRequest) ProtoMessage() {}

func (x *RecommendGameTakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gametake_v1_gametake_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendGameTakeRequest.ProtoReflect.Descriptor instead.
func (*RecommendGameTakeRequest) Descriptor() ([]byte, []int) {
	return file_proto_gametake_v1_gametake_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendGameTakeRequest) GetCards() []*Card {
	if x != nil {
		return x.Cards
	}
	return nil
}

type RecommendGameTakeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvailableTakes []*RecommendedGameTake `protobuf:"bytes,1,rep,name=available_takes,json=availableTakes,proto3" json:"available_takes,omitempty"`
}

func (x *RecommendGameTakeResponse) Reset() {
	*x = RecommendGameTakeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gametake_v1_gametake_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendGameTakeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendGameTakeResponse) ProtoMessage() {}

func (x *RecommendGameTakeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gametake_v1_gametake_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendGameTakeResponse.ProtoReflect.Descriptor instead.
func (*RecommendGameTakeResponse) Descriptor() ([]byte, []int) {
	return file_proto_gametake_v1_gametake_proto_rawDescGZIP(), []int{2}
}

func (x *RecommendGameTakeResponse) GetAvailableTakes() []*RecommendedGameTake {
	if x != nil {
		return x.AvailableTakes
	}
	return nil
}

type Flag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Flag) Reset() {
	*x = Flag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gametake_v1_gametake_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flag) ProtoMessage() {}

func (x *Flag) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gametake_v1_gametake_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flag.ProtoReflect.Descriptor instead.
func (*Flag) Descriptor() ([]byte, []int) {
	return file_proto_gametake_v1_gametake_proto_rawDescGZIP(), []int{3}
}

func (x *Flag) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RecommendedGameTake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Take  string  `protobuf:"bytes,1,opt,name=take,proto3" json:"take,omitempty"`
	Flags []*Flag `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (x *RecommendedGameTake) Reset() {
	*x = RecommendedGameTake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gametake_v1_gametake_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendedGameTake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendedGameTake) ProtoMessage() {}

func (x *RecommendedGameTake) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gametake_v1_gametake_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendedGameTake.ProtoReflect.Descriptor instead.
func (*RecommendedGameTake) Descriptor() ([]byte, []int) {
	return file_proto_gametake_v1_gametake_proto_rawDescGZIP(), []int{4}
}

func (x *RecommendedGameTake) GetTake() string {
	if x != nil {
		return x.Take
	}
	return ""
}

func (x *RecommendedGameTake) GetFlags() []*Flag {
	if x != nil {
		return x.Flags
	}
	return nil
}

var File_proto_gametake_v1_gametake_proto protoreflect.FileDescriptor

var file_proto_gametake_v1_gametake_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x61, 0x6b, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x61,
	0x6b, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x49, 0x0a, 0x18, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x74,
	0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x64, 0x73, 0x22, 0x6c, 0x0a, 0x19, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x47,
	0x61, 0x6d, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4f, 0x0a, 0x0f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x6b,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x6b, 0x65,
	0x52, 0x0e, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x73,
	0x22, 0x1a, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x58, 0x0a, 0x13,
	0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x54,
	0x61, 0x6b, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x61, 0x6d, 0x65, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x52,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x32, 0x82, 0x01, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x54,
	0x61, 0x6b, 0x65, 0x4c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x6e, 0x0a, 0x11, 0x52,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x61, 0x6b, 0x65,
	0x12, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x61, 0x6b,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x47, 0x61,
	0x6d, 0x65, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x47, 0x61, 0x6d, 0x65, 0x54,
	0x61, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x7a,
	0x69, 0x6e, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_gametake_v1_gametake_proto_rawDescOnce sync.Once
	file_proto_gametake_v1_gametake_proto_rawDescData = file_proto_gametake_v1_gametake_proto_rawDesc
)

func file_proto_gametake_v1_gametake_proto_rawDescGZIP() []byte {
	file_proto_gametake_v1_gametake_proto_rawDescOnce.Do(func() {
		file_proto_gametake_v1_gametake_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gametake_v1_gametake_proto_rawDescData)
	})
	return file_proto_gametake_v1_gametake_proto_rawDescData
}

var file_proto_gametake_v1_gametake_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_gametake_v1_gametake_proto_goTypes = []interface{}{
	(*Card)(nil),                      // 0: proto.gametake.v1.Card
	(*RecommendGameTakeRequest)(nil),  // 1: proto.gametake.v1.RecommendGameTakeRequest
	(*RecommendGameTakeResponse)(nil), // 2: proto.gametake.v1.RecommendGameTakeResponse
	(*Flag)(nil),                      // 3: proto.gametake.v1.flag
	(*RecommendedGameTake)(nil),       // 4: proto.gametake.v1.RecommendedGameTake
}
var file_proto_gametake_v1_gametake_proto_depIdxs = []int32{
	0, // 0: proto.gametake.v1.RecommendGameTakeRequest.cards:type_name -> proto.gametake.v1.Card
	4, // 1: proto.gametake.v1.RecommendGameTakeResponse.available_takes:type_name -> proto.gametake.v1.RecommendedGameTake
	3, // 2: proto.gametake.v1.RecommendedGameTake.flags:type_name -> proto.gametake.v1.flag
	1, // 3: proto.gametake.v1.GameTakeLearning.RecommendGameTake:input_type -> proto.gametake.v1.RecommendGameTakeRequest
	2, // 4: proto.gametake.v1.GameTakeLearning.RecommendGameTake:output_type -> proto.gametake.v1.RecommendGameTakeResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_gametake_v1_gametake_proto_init() }
func file_proto_gametake_v1_gametake_proto_init() {
	if File_proto_gametake_v1_gametake_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gametake_v1_gametake_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_proto_gametake_v1_gametake_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendGameTakeRequest); i {
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
		file_proto_gametake_v1_gametake_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendGameTakeResponse); i {
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
		file_proto_gametake_v1_gametake_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flag); i {
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
		file_proto_gametake_v1_gametake_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendedGameTake); i {
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
			RawDescriptor: file_proto_gametake_v1_gametake_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gametake_v1_gametake_proto_goTypes,
		DependencyIndexes: file_proto_gametake_v1_gametake_proto_depIdxs,
		MessageInfos:      file_proto_gametake_v1_gametake_proto_msgTypes,
	}.Build()
	File_proto_gametake_v1_gametake_proto = out.File
	file_proto_gametake_v1_gametake_proto_rawDesc = nil
	file_proto_gametake_v1_gametake_proto_goTypes = nil
	file_proto_gametake_v1_gametake_proto_depIdxs = nil
}
