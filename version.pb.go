// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: github.com/dosquad/go-cliversion/version.proto

package cliversion

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VersionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Build *BuildInfo `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	Git   *GitInfo   `protobuf:"bytes,2,opt,name=git,proto3" json:"git,omitempty"`
}

func (x *VersionInfo) Reset() {
	*x = VersionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dosquad_go_cliversion_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionInfo) ProtoMessage() {}

func (x *VersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dosquad_go_cliversion_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionInfo.ProtoReflect.Descriptor instead.
func (*VersionInfo) Descriptor() ([]byte, []int) {
	return file_github_com_dosquad_go_cliversion_version_proto_rawDescGZIP(), []int{0}
}

func (x *VersionInfo) GetBuild() *BuildInfo {
	if x != nil {
		return x.Build
	}
	return nil
}

func (x *VersionInfo) GetGit() *GitInfo {
	if x != nil {
		return x.Git
	}
	return nil
}

type BuildInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Debug     bool                   `protobuf:"varint,1,opt,name=debug,proto3" json:"debug,omitempty"`
	Method    string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Date      *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Version   string                 `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	GoVersion string                 `protobuf:"bytes,5,opt,name=go_version,proto3" json:"go_version,omitempty"`
}

func (x *BuildInfo) Reset() {
	*x = BuildInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dosquad_go_cliversion_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildInfo) ProtoMessage() {}

func (x *BuildInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dosquad_go_cliversion_version_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildInfo.ProtoReflect.Descriptor instead.
func (*BuildInfo) Descriptor() ([]byte, []int) {
	return file_github_com_dosquad_go_cliversion_version_proto_rawDescGZIP(), []int{1}
}

func (x *BuildInfo) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *BuildInfo) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *BuildInfo) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *BuildInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BuildInfo) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

type GitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo     string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Slug     string `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Commit   string `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	Tag      string `protobuf:"bytes,4,opt,name=tag,proto3" json:"tag,omitempty"`
	ExactTag string `protobuf:"bytes,5,opt,name=exact_tag,proto3" json:"exact_tag,omitempty"`
}

func (x *GitInfo) Reset() {
	*x = GitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_dosquad_go_cliversion_version_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitInfo) ProtoMessage() {}

func (x *GitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_dosquad_go_cliversion_version_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitInfo.ProtoReflect.Descriptor instead.
func (*GitInfo) Descriptor() ([]byte, []int) {
	return file_github_com_dosquad_go_cliversion_version_proto_rawDescGZIP(), []int{2}
}

func (x *GitInfo) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *GitInfo) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *GitInfo) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *GitInfo) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *GitInfo) GetExactTag() string {
	if x != nil {
		return x.ExactTag
	}
	return ""
}

var File_github_com_dosquad_go_cliversion_version_proto protoreflect.FileDescriptor

var file_github_com_dosquad_go_cliversion_version_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x73,
	0x71, 0x75, 0x61, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x64, 0x6f, 0x73, 0x71, 0x75, 0x61, 0x64, 0x2e, 0x63, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x64, 0x6f, 0x73, 0x71, 0x75, 0x61, 0x64, 0x2e, 0x63, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x2d, 0x0a, 0x03, 0x67, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x6f, 0x73, 0x71, 0x75, 0x61, 0x64,
	0x2e, 0x63, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x69, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x03, 0x67, 0x69, 0x74, 0x22, 0xa3, 0x01, 0x0a, 0x09, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x67, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x67, 0x6f, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x79,
	0x0a, 0x07, 0x47, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6c, 0x75,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x65,
	0x78, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x61, 0x67, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x73, 0x71, 0x75, 0x61, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x63, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x63, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_dosquad_go_cliversion_version_proto_rawDescOnce sync.Once
	file_github_com_dosquad_go_cliversion_version_proto_rawDescData = file_github_com_dosquad_go_cliversion_version_proto_rawDesc
)

func file_github_com_dosquad_go_cliversion_version_proto_rawDescGZIP() []byte {
	file_github_com_dosquad_go_cliversion_version_proto_rawDescOnce.Do(func() {
		file_github_com_dosquad_go_cliversion_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_dosquad_go_cliversion_version_proto_rawDescData)
	})
	return file_github_com_dosquad_go_cliversion_version_proto_rawDescData
}

var file_github_com_dosquad_go_cliversion_version_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_dosquad_go_cliversion_version_proto_goTypes = []any{
	(*VersionInfo)(nil),           // 0: dosquad.cliversion.VersionInfo
	(*BuildInfo)(nil),             // 1: dosquad.cliversion.BuildInfo
	(*GitInfo)(nil),               // 2: dosquad.cliversion.GitInfo
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_github_com_dosquad_go_cliversion_version_proto_depIdxs = []int32{
	1, // 0: dosquad.cliversion.VersionInfo.build:type_name -> dosquad.cliversion.BuildInfo
	2, // 1: dosquad.cliversion.VersionInfo.git:type_name -> dosquad.cliversion.GitInfo
	3, // 2: dosquad.cliversion.BuildInfo.date:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_dosquad_go_cliversion_version_proto_init() }
func file_github_com_dosquad_go_cliversion_version_proto_init() {
	if File_github_com_dosquad_go_cliversion_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_dosquad_go_cliversion_version_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*VersionInfo); i {
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
		file_github_com_dosquad_go_cliversion_version_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BuildInfo); i {
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
		file_github_com_dosquad_go_cliversion_version_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GitInfo); i {
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
			RawDescriptor: file_github_com_dosquad_go_cliversion_version_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_dosquad_go_cliversion_version_proto_goTypes,
		DependencyIndexes: file_github_com_dosquad_go_cliversion_version_proto_depIdxs,
		MessageInfos:      file_github_com_dosquad_go_cliversion_version_proto_msgTypes,
	}.Build()
	File_github_com_dosquad_go_cliversion_version_proto = out.File
	file_github_com_dosquad_go_cliversion_version_proto_rawDesc = nil
	file_github_com_dosquad_go_cliversion_version_proto_goTypes = nil
	file_github_com_dosquad_go_cliversion_version_proto_depIdxs = nil
}
