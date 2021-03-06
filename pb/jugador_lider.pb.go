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

	Id          int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bot         bool  `protobuf:"varint,2,opt,name=bot,proto3" json:"bot,omitempty"`
	Vive        bool  `protobuf:"varint,3,opt,name=vive,proto3" json:"vive,omitempty"`
	SumaJugada1 int32 `protobuf:"varint,4,opt,name=SumaJugada1,proto3" json:"SumaJugada1,omitempty"`
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

func (x *Jugador) GetSumaJugada1() int32 {
	if x != nil {
		return x.SumaJugada1
	}
	return 0
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

type SolicitarInicioJuego struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SolicitarInicioJuego) Reset() {
	*x = SolicitarInicioJuego{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitarInicioJuego) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitarInicioJuego) ProtoMessage() {}

func (x *SolicitarInicioJuego) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitarInicioJuego.ProtoReflect.Descriptor instead.
func (*SolicitarInicioJuego) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{2}
}

func (x *SolicitarInicioJuego) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type EsperandoJugadores struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Confirmacion bool   `protobuf:"varint,2,opt,name=confirmacion,proto3" json:"confirmacion,omitempty"`
}

func (x *EsperandoJugadores) Reset() {
	*x = EsperandoJugadores{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EsperandoJugadores) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EsperandoJugadores) ProtoMessage() {}

func (x *EsperandoJugadores) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EsperandoJugadores.ProtoReflect.Descriptor instead.
func (*EsperandoJugadores) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{3}
}

func (x *EsperandoJugadores) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *EsperandoJugadores) GetConfirmacion() bool {
	if x != nil {
		return x.Confirmacion
	}
	return false
}

type JugadaEtapa1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Step int32 `protobuf:"varint,1,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *JugadaEtapa1) Reset() {
	*x = JugadaEtapa1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JugadaEtapa1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JugadaEtapa1) ProtoMessage() {}

func (x *JugadaEtapa1) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JugadaEtapa1.ProtoReflect.Descriptor instead.
func (*JugadaEtapa1) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{4}
}

func (x *JugadaEtapa1) GetStep() int32 {
	if x != nil {
		return x.Step
	}
	return 0
}

type Etapa struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Etapa) Reset() {
	*x = Etapa{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Etapa) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Etapa) ProtoMessage() {}

func (x *Etapa) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Etapa.ProtoReflect.Descriptor instead.
func (*Etapa) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{5}
}

func (x *Etapa) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type JugadaCliente struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Jugador *Jugador `protobuf:"bytes,2,opt,name=jugador,proto3" json:"jugador,omitempty"`
	Message int32    `protobuf:"varint,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *JugadaCliente) Reset() {
	*x = JugadaCliente{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JugadaCliente) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JugadaCliente) ProtoMessage() {}

func (x *JugadaCliente) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JugadaCliente.ProtoReflect.Descriptor instead.
func (*JugadaCliente) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{6}
}

func (x *JugadaCliente) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JugadaCliente) GetJugador() *Jugador {
	if x != nil {
		return x.Jugador
	}
	return nil
}

func (x *JugadaCliente) GetMessage() int32 {
	if x != nil {
		return x.Message
	}
	return 0
}

type JugadaLider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Message    int32 `protobuf:"varint,2,opt,name=message,proto3" json:"message,omitempty"`
	ReadyEtapa bool  `protobuf:"varint,3,opt,name=readyEtapa,proto3" json:"readyEtapa,omitempty"`
}

func (x *JugadaLider) Reset() {
	*x = JugadaLider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_jugador_lider_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JugadaLider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JugadaLider) ProtoMessage() {}

func (x *JugadaLider) ProtoReflect() protoreflect.Message {
	mi := &file_pb_jugador_lider_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JugadaLider.ProtoReflect.Descriptor instead.
func (*JugadaLider) Descriptor() ([]byte, []int) {
	return file_pb_jugador_lider_proto_rawDescGZIP(), []int{7}
}

func (x *JugadaLider) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *JugadaLider) GetMessage() int32 {
	if x != nil {
		return x.Message
	}
	return 0
}

func (x *JugadaLider) GetReadyEtapa() bool {
	if x != nil {
		return x.ReadyEtapa
	}
	return false
}

var File_pb_jugador_lider_proto protoreflect.FileDescriptor

var file_pb_jugador_lider_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x5f, 0x6c, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x07, 0x4a, 0x75, 0x67, 0x61,
	0x64, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x03, 0x62, 0x6f, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x69, 0x76, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x04, 0x76, 0x69, 0x76, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x75, 0x6d,
	0x61, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x53, 0x75, 0x6d, 0x61, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x31, 0x22, 0x26, 0x0a, 0x12, 0x49,
	0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03,
	0x62, 0x6f, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x69, 0x63, 0x69, 0x6f, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x52, 0x0a, 0x12, 0x45,
	0x73, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x22,
	0x22, 0x0a, 0x0c, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x45, 0x74, 0x61, 0x70, 0x61, 0x31, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x74, 0x65, 0x70, 0x22, 0x17, 0x0a, 0x05, 0x45, 0x74, 0x61, 0x70, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5d, 0x0a, 0x0d,
	0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a,
	0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x52, 0x07, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x57, 0x0a, 0x0b, 0x6a,
	0x75, 0x67, 0x61, 0x64, 0x61, 0x4c, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x74, 0x61,
	0x70, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45,
	0x74, 0x61, 0x70, 0x61, 0x32, 0xe8, 0x01, 0x0a, 0x13, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72,
	0x4c, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0f,
	0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x55, 0x6e, 0x69, 0x72, 0x63, 0x65, 0x12,
	0x13, 0x2e, 0x49, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x63, 0x69, 0x6f, 0x6e, 0x4a, 0x75, 0x67,
	0x61, 0x64, 0x6f, 0x72, 0x1a, 0x08, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x00,
	0x12, 0x3e, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x72, 0x45, 0x74, 0x61, 0x70, 0x61,
	0x12, 0x15, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x61, 0x72, 0x49, 0x6e, 0x69, 0x63,
	0x69, 0x6f, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x1a, 0x13, 0x2e, 0x45, 0x73, 0x70, 0x65, 0x72, 0x61,
	0x6e, 0x64, 0x6f, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x33, 0x0a, 0x0f, 0x4c, 0x75, 0x7a, 0x52, 0x6f, 0x6a, 0x61, 0x4c, 0x75, 0x7a, 0x56, 0x65,
	0x72, 0x64, 0x65, 0x12, 0x0e, 0x2e, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x65, 0x1a, 0x0c, 0x2e, 0x6a, 0x75, 0x67, 0x61, 0x64, 0x61, 0x4c, 0x69, 0x64, 0x65,
	0x72, 0x22, 0x00, 0x30, 0x01, 0x12, 0x28, 0x0a, 0x10, 0x45, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x44,
	0x65, 0x6c, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x12, 0x08, 0x2e, 0x4a, 0x75, 0x67, 0x61,
	0x64, 0x6f, 0x72, 0x1a, 0x08, 0x2e, 0x4a, 0x75, 0x67, 0x61, 0x64, 0x6f, 0x72, 0x22, 0x00, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_pb_jugador_lider_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pb_jugador_lider_proto_goTypes = []interface{}{
	(*Jugador)(nil),              // 0: Jugador
	(*InscripcionJugador)(nil),   // 1: InscripcionJugador
	(*SolicitarInicioJuego)(nil), // 2: SolicitarInicioJuego
	(*EsperandoJugadores)(nil),   // 3: EsperandoJugadores
	(*JugadaEtapa1)(nil),         // 4: JugadaEtapa1
	(*Etapa)(nil),                // 5: Etapa
	(*JugadaCliente)(nil),        // 6: jugadaCliente
	(*JugadaLider)(nil),          // 7: jugadaLider
}
var file_pb_jugador_lider_proto_depIdxs = []int32{
	0, // 0: jugadaCliente.jugador:type_name -> Jugador
	1, // 1: JugadorLiderService.SolicitarUnirce:input_type -> InscripcionJugador
	2, // 2: JugadorLiderService.IniciarEtapa:input_type -> SolicitarInicioJuego
	6, // 3: JugadorLiderService.LuzRojaLuzVerde:input_type -> jugadaCliente
	0, // 4: JugadorLiderService.EstadoDelJugador:input_type -> Jugador
	0, // 5: JugadorLiderService.SolicitarUnirce:output_type -> Jugador
	3, // 6: JugadorLiderService.IniciarEtapa:output_type -> EsperandoJugadores
	7, // 7: JugadorLiderService.LuzRojaLuzVerde:output_type -> jugadaLider
	0, // 8: JugadorLiderService.EstadoDelJugador:output_type -> Jugador
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_pb_jugador_lider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitarInicioJuego); i {
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
		file_pb_jugador_lider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EsperandoJugadores); i {
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
		file_pb_jugador_lider_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JugadaEtapa1); i {
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
		file_pb_jugador_lider_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Etapa); i {
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
		file_pb_jugador_lider_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JugadaCliente); i {
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
		file_pb_jugador_lider_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JugadaLider); i {
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
			NumMessages:   8,
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
