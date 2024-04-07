// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: worktabd.proto

package worktabdpb

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

type Container struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Container) Reset() {
	*x = Container{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worktabd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Container) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Container) ProtoMessage() {}

func (x *Container) ProtoReflect() protoreflect.Message {
	mi := &file_worktabd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Container.ProtoReflect.Descriptor instead.
func (*Container) Descriptor() ([]byte, []int) {
	return file_worktabd_proto_rawDescGZIP(), []int{0}
}

func (x *Container) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worktabd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_worktabd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_worktabd_proto_rawDescGZIP(), []int{1}
}

type ServerVersionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ServerVersionResponse) Reset() {
	*x = ServerVersionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_worktabd_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerVersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerVersionResponse) ProtoMessage() {}

func (x *ServerVersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_worktabd_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerVersionResponse.ProtoReflect.Descriptor instead.
func (*ServerVersionResponse) Descriptor() ([]byte, []int) {
	return file_worktabd_proto_rawDescGZIP(), []int{2}
}

func (x *ServerVersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_worktabd_proto protoreflect.FileDescriptor

var file_worktabd_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x62, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x62, 0x64, 0x70, 0x62, 0x22, 0x1b, 0x0a, 0x09,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x31, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x8c, 0x01, 0x0a, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61,
	0x62, 0x64, 0x12, 0x47, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x62, 0x64, 0x70, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x21, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61, 0x62,
	0x64, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x74,
	0x61, 0x62, 0x64, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x74, 0x61, 0x62, 0x64, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x74, 0x61,
	0x62, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_worktabd_proto_rawDescOnce sync.Once
	file_worktabd_proto_rawDescData = file_worktabd_proto_rawDesc
)

func file_worktabd_proto_rawDescGZIP() []byte {
	file_worktabd_proto_rawDescOnce.Do(func() {
		file_worktabd_proto_rawDescData = protoimpl.X.CompressGZIP(file_worktabd_proto_rawDescData)
	})
	return file_worktabd_proto_rawDescData
}

var file_worktabd_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_worktabd_proto_goTypes = []interface{}{
	(*Container)(nil),             // 0: worktabdpb.Container
	(*Empty)(nil),                 // 1: worktabdpb.Empty
	(*ServerVersionResponse)(nil), // 2: worktabdpb.ServerVersionResponse
}
var file_worktabd_proto_depIdxs = []int32{
	1, // 0: worktabdpb.worktabd.ServerVersion:input_type -> worktabdpb.Empty
	1, // 1: worktabdpb.worktabd.container:input_type -> worktabdpb.Empty
	2, // 2: worktabdpb.worktabd.ServerVersion:output_type -> worktabdpb.ServerVersionResponse
	0, // 3: worktabdpb.worktabd.container:output_type -> worktabdpb.Container
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_worktabd_proto_init() }
func file_worktabd_proto_init() {
	if File_worktabd_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_worktabd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Container); i {
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
		file_worktabd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_worktabd_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerVersionResponse); i {
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
			RawDescriptor: file_worktabd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_worktabd_proto_goTypes,
		DependencyIndexes: file_worktabd_proto_depIdxs,
		MessageInfos:      file_worktabd_proto_msgTypes,
	}.Build()
	File_worktabd_proto = out.File
	file_worktabd_proto_rawDesc = nil
	file_worktabd_proto_goTypes = nil
	file_worktabd_proto_depIdxs = nil
}
