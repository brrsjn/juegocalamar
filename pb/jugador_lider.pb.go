// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: pb/jugador_lider.proto

package pb

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

type Jugador struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bot  bool  `protobuf:"varint,2,opt,name=bot,proto3" json:"bot,omitempty"`
	Vive bool  `protobuf:"varint,3,opt,name=vive,proto3" json:"vive,omitempty"`
}

func (x *Jugador) Reset() {
	*x = Jugador{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Jugador) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Jugador) ProtoMessage() {}

func (x *Jugador) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Jugador.ProtoReflect.Descriptor instead.
func (*Jugador) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{0}
}

func (x *Jugador) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Jugador) GetBot() bool {
	if x != nil {
		return x.Bot
	}
	return false
}

func (x *Jugador) GetVive() bool {
	if x != nil {
		return x.Vive
	}
	return false
}

type InscripcionJugador struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bot bool `protobuf:"varint,1,opt,name=bot,proto3" json:"bot,omitempty"`
}

func (x *InscripcionJugador) Reset() {
	*x = InscripcionJugador{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InscripcionJugador) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InscripcionJugador) ProtoMessage() {}

func (x *InscripcionJugador) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InscripcionJugador.ProtoReflect.Descriptor instead.
func (*InscripcionJugador) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{1}
}

func (x *InscripcionJugador) GetBot() bool {
	if x != nil {
		return x.Bot
	}
	return false
}

var File_pb_jugador_lider_proto protoreflect.FileDescriptor

var file_pb_jugador_lider_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x07, 0x4a, 0x75, 0x67, 0x61,
	0x64, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x62, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x76, 0x69, 0x76, 0x65, 0x22, 0x26, 0x0a, 0x12, 0x49, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x62, 0x6f,
	0x74, 0x32, 0x49, 0x0a, 0x13, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x4c, 0x69, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0f, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x61, 0x72, 0x55, 0x6e, 0x69, 0x72, 0x63, 0x65, 0x12, 0x13, 0x2e, 0x49, 0x6e,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72,
	0x1a, 0x08, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x42, 0x05, 0x5a, 0x03,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_jugador_lider_proto_rawDescOnce sync.Once
	file_pb_jugador_lider_proto_rawDescData = file_pb_jugador_lider_proto_rawDesc
)

func file_pb_jugador_lider_proto_rawDescGZIP() []byte {
	file_pb_jugador_lider_proto_rawDescOnce.Do(func() {
		file_pb_jugador_lider_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_jugador_lider_proto_rawDescData)
	})
	return file_pb_jugador_lider_proto_rawDescData
}

var file_pb_jugador_lider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_jugador_lider_proto_goTypes = []interface{}{
	(*Jugador)(nil),            // 0: Jugador
	(*InscripcionJugador)(nil), // 1: InscripcionJugador
}
var file_pb_jugador_lider_proto_depIdxs = []int32{
	1, // 0: JugadorLiderService.SolicitarUnirce:input_type -> InscripcionJugador
	0, // 1: JugadorLiderService.SolicitarUnirce:output_type -> Jugador
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_jugador_lider_proto_init() }
func file_pb_jugador_lider_proto_init() {
	if File_pb_jugador_lider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_jugador_lider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Jugador); i {
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
		file_pb_jugador_lider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InscripcionJugador); i {
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
			RawDescriptor: file_pb_jugador_lider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_jugador_lider_proto_goTypes,
		DependencyIndexes: file_pb_jugador_lider_proto_depIdxs,
		MessageInfos:      file_pb_jugador_lider_proto_msgTypes,
	}.Build()
	File_pb_jugador_lider_proto = out.File
	file_pb_jugador_lider_proto_rawDesc = nil
	file_pb_jugador_lider_proto_goTypes = nil
	file_pb_jugador_lider_proto_depIdxs = nil
}
