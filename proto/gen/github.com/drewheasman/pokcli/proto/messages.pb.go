// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: messages.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlayActionType int32

const (
	PlayActionType_CHECK PlayActionType = 0
	PlayActionType_CALL  PlayActionType = 1
	PlayActionType_BET   PlayActionType = 2
	PlayActionType_FOLD  PlayActionType = 3
)

// Enum value maps for PlayActionType.
var (
	PlayActionType_name = map[int32]string{
		0: "CHECK",
		1: "CALL",
		2: "BET",
		3: "FOLD",
	}
	PlayActionType_value = map[string]int32{
		"CHECK": 0,
		"CALL":  1,
		"BET":   2,
		"FOLD":  3,
	}
)

func (x PlayActionType) Enum() *PlayActionType {
	p := new(PlayActionType)
	*p = x
	return p
}

func (x PlayActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PlayActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_messages_proto_enumTypes[0].Descriptor()
}

func (PlayActionType) Type() protoreflect.EnumType {
	return &file_messages_proto_enumTypes[0]
}

func (x PlayActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PlayActionType.Descriptor instead.
func (PlayActionType) EnumDescriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

type Register struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Register) Reset() {
	*x = Register{}
	mi := &file_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Register) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterOutcome struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        bool                   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegisterOutcome) Reset() {
	*x = RegisterOutcome{}
	mi := &file_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterOutcome) ProtoMessage() {}

func (x *RegisterOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterOutcome.ProtoReflect.Descriptor instead.
func (*RegisterOutcome) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterOutcome) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *RegisterOutcome) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PlayAction struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          PlayActionType         `protobuf:"varint,1,opt,name=type,proto3,enum=pokcli.PlayActionType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PlayAction) Reset() {
	*x = PlayAction{}
	mi := &file_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlayAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayAction) ProtoMessage() {}

func (x *PlayAction) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayAction.ProtoReflect.Descriptor instead.
func (*PlayAction) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PlayAction) GetType() PlayActionType {
	if x != nil {
		return x.Type
	}
	return PlayActionType_CHECK
}

type BetOutcome struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PlayerName    string                 `protobuf:"bytes,1,opt,name=player_name,json=playerName,proto3" json:"player_name,omitempty"`
	Value         int32                  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BetOutcome) Reset() {
	*x = BetOutcome{}
	mi := &file_messages_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BetOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BetOutcome) ProtoMessage() {}

func (x *BetOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_messages_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BetOutcome.ProtoReflect.Descriptor instead.
func (*BetOutcome) Descriptor() ([]byte, []int) {
	return file_messages_proto_rawDescGZIP(), []int{3}
}

func (x *BetOutcome) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *BetOutcome) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_messages_proto protoreflect.FileDescriptor

const file_messages_proto_rawDesc = "" +
	"\n" +
	"\x0emessages.proto\x12\x06pokcli\"\x1e\n" +
	"\bRegister\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"C\n" +
	"\x0fRegisterOutcome\x12\x16\n" +
	"\x06status\x18\x01 \x01(\bR\x06status\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"8\n" +
	"\n" +
	"PlayAction\x12*\n" +
	"\x04type\x18\x01 \x01(\x0e2\x16.pokcli.PlayActionTypeR\x04type\"C\n" +
	"\n" +
	"BetOutcome\x12\x1f\n" +
	"\vplayer_name\x18\x01 \x01(\tR\n" +
	"playerName\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value*8\n" +
	"\x0ePlayActionType\x12\t\n" +
	"\x05CHECK\x10\x00\x12\b\n" +
	"\x04CALL\x10\x01\x12\a\n" +
	"\x03BET\x10\x02\x12\b\n" +
	"\x04FOLD\x10\x03B&Z$github.com/drewheasman/pokcli/proto/b\x06proto3"

var (
	file_messages_proto_rawDescOnce sync.Once
	file_messages_proto_rawDescData []byte
)

func file_messages_proto_rawDescGZIP() []byte {
	file_messages_proto_rawDescOnce.Do(func() {
		file_messages_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_messages_proto_rawDesc), len(file_messages_proto_rawDesc)))
	})
	return file_messages_proto_rawDescData
}

var file_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_messages_proto_goTypes = []any{
	(PlayActionType)(0),     // 0: pokcli.PlayActionType
	(*Register)(nil),        // 1: pokcli.Register
	(*RegisterOutcome)(nil), // 2: pokcli.RegisterOutcome
	(*PlayAction)(nil),      // 3: pokcli.PlayAction
	(*BetOutcome)(nil),      // 4: pokcli.BetOutcome
}
var file_messages_proto_depIdxs = []int32{
	0, // 0: pokcli.PlayAction.type:type_name -> pokcli.PlayActionType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_messages_proto_init() }
func file_messages_proto_init() {
	if File_messages_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_messages_proto_rawDesc), len(file_messages_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_proto_goTypes,
		DependencyIndexes: file_messages_proto_depIdxs,
		EnumInfos:         file_messages_proto_enumTypes,
		MessageInfos:      file_messages_proto_msgTypes,
	}.Build()
	File_messages_proto = out.File
	file_messages_proto_goTypes = nil
	file_messages_proto_depIdxs = nil
}
