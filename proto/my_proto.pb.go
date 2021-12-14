// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: my_proto.proto

package t3

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

type Comando struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//string COMMAND = 1; // comando
	Planeta string `protobuf:"bytes,1,opt,name=planeta,proto3" json:"planeta,omitempty"` // nombre planeta
	Ciudad  string `protobuf:"bytes,2,opt,name=ciudad,proto3" json:"ciudad,omitempty"`   // nombre ciudad
	Valor   string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`     // valor de comando de existir, -1 en caso contrario
	Vector  string `protobuf:"bytes,4,opt,name=vector,proto3" json:"vector,omitempty"`   // reloj vector, en caso de enviarlo
}

func (x *Comando) Reset() {
	*x = Comando{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comando) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comando) ProtoMessage() {}

func (x *Comando) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comando.ProtoReflect.Descriptor instead.
func (*Comando) Descriptor() ([]byte, []int) {
	return file_my_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Comando) GetPlaneta() string {
	if x != nil {
		return x.Planeta
	}
	return ""
}

func (x *Comando) GetCiudad() string {
	if x != nil {
		return x.Ciudad
	}
	return ""
}

func (x *Comando) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

func (x *Comando) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

type RespuestaBroker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP     string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`         // IP del Fulcrum
	Vector string `protobuf:"bytes,2,opt,name=vector,proto3" json:"vector,omitempty"` // reloj vector (para Leia)
	Valor  string `protobuf:"bytes,3,opt,name=valor,proto3" json:"valor,omitempty"`   // reloj vector (para Leia)
}

func (x *RespuestaBroker) Reset() {
	*x = RespuestaBroker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaBroker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaBroker) ProtoMessage() {}

func (x *RespuestaBroker) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaBroker.ProtoReflect.Descriptor instead.
func (*RespuestaBroker) Descriptor() ([]byte, []int) {
	return file_my_proto_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaBroker) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *RespuestaBroker) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

func (x *RespuestaBroker) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

type RespuestaReplica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vector string `protobuf:"bytes,1,opt,name=vector,proto3" json:"vector,omitempty"` // reloj vector
	Valor  string `protobuf:"bytes,2,opt,name=valor,proto3" json:"valor,omitempty"`   // valor solicitado (por Leia)
}

func (x *RespuestaReplica) Reset() {
	*x = RespuestaReplica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaReplica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaReplica) ProtoMessage() {}

func (x *RespuestaReplica) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaReplica.ProtoReflect.Descriptor instead.
func (*RespuestaReplica) Descriptor() ([]byte, []int) {
	return file_my_proto_proto_rawDescGZIP(), []int{2}
}

func (x *RespuestaReplica) GetVector() string {
	if x != nil {
		return x.Vector
	}
	return ""
}

func (x *RespuestaReplica) GetValor() string {
	if x != nil {
		return x.Valor
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_my_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_my_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_my_proto_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_my_proto_proto protoreflect.FileDescriptor

var file_my_proto_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x74, 0x33, 0x22, 0x69, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x69, 0x75,
	0x64, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x69, 0x75, 0x64, 0x61,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22,
	0x4f, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x50, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72,
	0x22, 0x40, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x6f, 0x72, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x32, 0xb6, 0x02, 0x0a, 0x0d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12,
	0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x74,
	0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61,
	0x6e, 0x64, 0x6f, 0x1a, 0x13, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f,
	0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x13, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x74,
	0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x13, 0x2e, 0x74, 0x33, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00,
	0x12, 0x30, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0b,
	0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x13, 0x2e, 0x74, 0x33,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61,
	0x6e, 0x64, 0x6f, 0x1a, 0x13, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73,
	0x74, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x00, 0x32, 0xf3, 0x02, 0x0a, 0x0e, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a,
	0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x14, 0x2e,
	0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f,
	0x1a, 0x14, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f,
	0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x14, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x0b, 0x2e, 0x74, 0x33,
	0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x14, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x0b, 0x2e, 0x74, 0x33, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64,
	0x6f, 0x1a, 0x14, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0b, 0x2e, 0x74, 0x33,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x74, 0x33, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x22, 0x00,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x74, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_my_proto_proto_rawDescOnce sync.Once
	file_my_proto_proto_rawDescData = file_my_proto_proto_rawDesc
)

func file_my_proto_proto_rawDescGZIP() []byte {
	file_my_proto_proto_rawDescOnce.Do(func() {
		file_my_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_my_proto_proto_rawDescData)
	})
	return file_my_proto_proto_rawDescData
}

var file_my_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_my_proto_proto_goTypes = []interface{}{
	(*Comando)(nil),          // 0: t3.Comando
	(*RespuestaBroker)(nil),  // 1: t3.RespuestaBroker
	(*RespuestaReplica)(nil), // 2: t3.RespuestaReplica
	(*Message)(nil),          // 3: t3.Message
}
var file_my_proto_proto_depIdxs = []int32{
	3,  // 0: t3.BrokerService.SayHello:input_type -> t3.Message
	0,  // 1: t3.BrokerService.AddCity:input_type -> t3.Comando
	0,  // 2: t3.BrokerService.UpdateName:input_type -> t3.Comando
	0,  // 3: t3.BrokerService.UpdateNumber:input_type -> t3.Comando
	0,  // 4: t3.BrokerService.DeleteCity:input_type -> t3.Comando
	0,  // 5: t3.BrokerService.GetNumberRebelds:input_type -> t3.Comando
	3,  // 6: t3.FulcrumService.SayHello:input_type -> t3.Message
	0,  // 7: t3.FulcrumService.AddCity:input_type -> t3.Comando
	0,  // 8: t3.FulcrumService.UpdateName:input_type -> t3.Comando
	0,  // 9: t3.FulcrumService.UpdateNumber:input_type -> t3.Comando
	0,  // 10: t3.FulcrumService.DeleteCity:input_type -> t3.Comando
	0,  // 11: t3.FulcrumService.GetNumberRebelds:input_type -> t3.Comando
	3,  // 12: t3.FulcrumService.GetClockVector:input_type -> t3.Message
	3,  // 13: t3.BrokerService.SayHello:output_type -> t3.Message
	1,  // 14: t3.BrokerService.AddCity:output_type -> t3.RespuestaBroker
	1,  // 15: t3.BrokerService.UpdateName:output_type -> t3.RespuestaBroker
	1,  // 16: t3.BrokerService.UpdateNumber:output_type -> t3.RespuestaBroker
	1,  // 17: t3.BrokerService.DeleteCity:output_type -> t3.RespuestaBroker
	1,  // 18: t3.BrokerService.GetNumberRebelds:output_type -> t3.RespuestaBroker
	3,  // 19: t3.FulcrumService.SayHello:output_type -> t3.Message
	2,  // 20: t3.FulcrumService.AddCity:output_type -> t3.RespuestaReplica
	2,  // 21: t3.FulcrumService.UpdateName:output_type -> t3.RespuestaReplica
	2,  // 22: t3.FulcrumService.UpdateNumber:output_type -> t3.RespuestaReplica
	2,  // 23: t3.FulcrumService.DeleteCity:output_type -> t3.RespuestaReplica
	2,  // 24: t3.FulcrumService.GetNumberRebelds:output_type -> t3.RespuestaReplica
	2,  // 25: t3.FulcrumService.GetClockVector:output_type -> t3.RespuestaReplica
	13, // [13:26] is the sub-list for method output_type
	0,  // [0:13] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_my_proto_proto_init() }
func file_my_proto_proto_init() {
	if File_my_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_my_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comando); i {
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
		file_my_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaBroker); i {
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
		file_my_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaReplica); i {
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
		file_my_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_my_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_my_proto_proto_goTypes,
		DependencyIndexes: file_my_proto_proto_depIdxs,
		MessageInfos:      file_my_proto_proto_msgTypes,
	}.Build()
	File_my_proto_proto = out.File
	file_my_proto_proto_rawDesc = nil
	file_my_proto_proto_goTypes = nil
	file_my_proto_proto_depIdxs = nil
}
