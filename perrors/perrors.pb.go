// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: protobufs/perrors/perrors.proto

package perrors

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

type Errors int32

const (
	Errors_Ok                  Errors = 0
	Errors_GeneratingToken     Errors = 1
	Errors_WrongUsername       Errors = 2
	Errors_WrongPassword       Errors = 3
	Errors_InvalidFirstName    Errors = 4
	Errors_InvalidLastName     Errors = 5
	Errors_InvalidEmail        Errors = 6
	Errors_InvalidAge          Errors = 7
	Errors_InvalidUsername     Errors = 8
	Errors_AlreadyUsedUsername Errors = 9
	Errors_InvalidPassword     Errors = 10
	Errors_SomethingWrong      Errors = 12
	Errors_NotFound            Errors = 13
	Errors_Unauthorized        Errors = 14
	Errors_AlreadyUsedEmail    Errors = 15
)

// Enum value maps for Errors.
var (
	Errors_name = map[int32]string{
		0:  "Ok",
		1:  "GeneratingToken",
		2:  "WrongUsername",
		3:  "WrongPassword",
		4:  "InvalidFirstName",
		5:  "InvalidLastName",
		6:  "InvalidEmail",
		7:  "InvalidAge",
		8:  "InvalidUsername",
		9:  "AlreadyUsedUsername",
		10: "InvalidPassword",
		12: "SomethingWrong",
		13: "NotFound",
		14: "Unauthorized",
		15: "AlreadyUsedEmail",
	}
	Errors_value = map[string]int32{
		"Ok":                  0,
		"GeneratingToken":     1,
		"WrongUsername":       2,
		"WrongPassword":       3,
		"InvalidFirstName":    4,
		"InvalidLastName":     5,
		"InvalidEmail":        6,
		"InvalidAge":          7,
		"InvalidUsername":     8,
		"AlreadyUsedUsername": 9,
		"InvalidPassword":     10,
		"SomethingWrong":      12,
		"NotFound":            13,
		"Unauthorized":        14,
		"AlreadyUsedEmail":    15,
	}
)

func (x Errors) Enum() *Errors {
	p := new(Errors)
	*p = x
	return p
}

func (x Errors) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Errors) Descriptor() protoreflect.EnumDescriptor {
	return file_protobufs_perrors_perrors_proto_enumTypes[0].Descriptor()
}

func (Errors) Type() protoreflect.EnumType {
	return &file_protobufs_perrors_perrors_proto_enumTypes[0]
}

func (x Errors) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Errors.Descriptor instead.
func (Errors) EnumDescriptor() ([]byte, []int) {
	return file_protobufs_perrors_perrors_proto_rawDescGZIP(), []int{0}
}

var File_protobufs_perrors_perrors_proto protoreflect.FileDescriptor

var file_protobufs_perrors_perrors_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2f, 0x70, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x70, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2a, 0xa5, 0x02, 0x0a, 0x06, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x04, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x10, 0x05, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x41, 0x67, 0x65, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x08, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x6c,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x55, 0x73, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x10, 0x09, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x0a, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x6f, 0x6d, 0x65,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x57, 0x72, 0x6f, 0x6e, 0x67, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08,
	0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10,
	0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x55, 0x73, 0x65, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x10, 0x0f, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobufs_perrors_perrors_proto_rawDescOnce sync.Once
	file_protobufs_perrors_perrors_proto_rawDescData = file_protobufs_perrors_perrors_proto_rawDesc
)

func file_protobufs_perrors_perrors_proto_rawDescGZIP() []byte {
	file_protobufs_perrors_perrors_proto_rawDescOnce.Do(func() {
		file_protobufs_perrors_perrors_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobufs_perrors_perrors_proto_rawDescData)
	})
	return file_protobufs_perrors_perrors_proto_rawDescData
}

var file_protobufs_perrors_perrors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protobufs_perrors_perrors_proto_goTypes = []interface{}{
	(Errors)(0), // 0: perrors.Errors
}
var file_protobufs_perrors_perrors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobufs_perrors_perrors_proto_init() }
func file_protobufs_perrors_perrors_proto_init() {
	if File_protobufs_perrors_perrors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobufs_perrors_perrors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobufs_perrors_perrors_proto_goTypes,
		DependencyIndexes: file_protobufs_perrors_perrors_proto_depIdxs,
		EnumInfos:         file_protobufs_perrors_perrors_proto_enumTypes,
	}.Build()
	File_protobufs_perrors_perrors_proto = out.File
	file_protobufs_perrors_perrors_proto_rawDesc = nil
	file_protobufs_perrors_perrors_proto_goTypes = nil
	file_protobufs_perrors_perrors_proto_depIdxs = nil
}
