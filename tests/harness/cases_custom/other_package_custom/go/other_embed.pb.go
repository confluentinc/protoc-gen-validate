// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: other_embed.proto

package other_package_custom

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Embed_Enumerated int32

const (
	Embed_VALUE Embed_Enumerated = 0
)

// Enum value maps for Embed_Enumerated.
var (
	Embed_Enumerated_name = map[int32]string{
		0: "VALUE",
	}
	Embed_Enumerated_value = map[string]int32{
		"VALUE": 0,
	}
)

func (x Embed_Enumerated) Enum() *Embed_Enumerated {
	p := new(Embed_Enumerated)
	*p = x
	return p
}

func (x Embed_Enumerated) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Embed_Enumerated) Descriptor() protoreflect.EnumDescriptor {
	return file_other_embed_proto_enumTypes[0].Descriptor()
}

func (Embed_Enumerated) Type() protoreflect.EnumType {
	return &file_other_embed_proto_enumTypes[0]
}

func (x Embed_Enumerated) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Embed_Enumerated.Descriptor instead.
func (Embed_Enumerated) EnumDescriptor() ([]byte, []int) {
	return file_other_embed_proto_rawDescGZIP(), []int{0, 0}
}

// Validate message embedding across packages.
type Embed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int64 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Embed) Reset() {
	*x = Embed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_embed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Embed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Embed) ProtoMessage() {}

func (x *Embed) ProtoReflect() protoreflect.Message {
	mi := &file_other_embed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Embed.ProtoReflect.Descriptor instead.
func (*Embed) Descriptor() ([]byte, []int) {
	return file_other_embed_proto_rawDescGZIP(), []int{0}
}

func (x *Embed) GetVal() int64 {
	if x != nil {
		return x.Val
	}
	return 0
}

var File_other_embed_proto protoreflect.FileDescriptor

var file_other_embed_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x62, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a,
	0x05, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x22, 0x17, 0x0a, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x09, 0x0a, 0x05, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x10, 0x00, 0x42, 0x16, 0x5a, 0x14, 0x6f, 0x74,
	0x68, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_other_embed_proto_rawDescOnce sync.Once
	file_other_embed_proto_rawDescData = file_other_embed_proto_rawDesc
)

func file_other_embed_proto_rawDescGZIP() []byte {
	file_other_embed_proto_rawDescOnce.Do(func() {
		file_other_embed_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_embed_proto_rawDescData)
	})
	return file_other_embed_proto_rawDescData
}

var file_other_embed_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_other_embed_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_other_embed_proto_goTypes = []interface{}{
	(Embed_Enumerated)(0), // 0: tests.harness.cases_custom.other_package_custom.Embed.Enumerated
	(*Embed)(nil),         // 1: tests.harness.cases_custom.other_package_custom.Embed
}
var file_other_embed_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_other_embed_proto_init() }
func file_other_embed_proto_init() {
	if File_other_embed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_embed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Embed); i {
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
			RawDescriptor: file_other_embed_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_other_embed_proto_goTypes,
		DependencyIndexes: file_other_embed_proto_depIdxs,
		EnumInfos:         file_other_embed_proto_enumTypes,
		MessageInfos:      file_other_embed_proto_msgTypes,
	}.Build()
	File_other_embed_proto = out.File
	file_other_embed_proto_rawDesc = nil
	file_other_embed_proto_goTypes = nil
	file_other_embed_proto_depIdxs = nil
}
