// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: twitter.proto

package grpcTwitter

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

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_twitter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_twitter_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_twitter_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Tweet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	User          *User                  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Text          string                 `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	Timestamp     int64                  `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tweet) Reset() {
	*x = Tweet{}
	mi := &file_twitter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tweet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tweet) ProtoMessage() {}

func (x *Tweet) ProtoReflect() protoreflect.Message {
	mi := &file_twitter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tweet.ProtoReflect.Descriptor instead.
func (*Tweet) Descriptor() ([]byte, []int) {
	return file_twitter_proto_rawDescGZIP(), []int{1}
}

func (x *Tweet) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tweet) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Tweet) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Tweet) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type CreateTweetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          *User                  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTweetRequest) Reset() {
	*x = CreateTweetRequest{}
	mi := &file_twitter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTweetRequest) ProtoMessage() {}

func (x *CreateTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_twitter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTweetRequest.ProtoReflect.Descriptor instead.
func (*CreateTweetRequest) Descriptor() ([]byte, []int) {
	return file_twitter_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTweetRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *CreateTweetRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreateTweetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tweet         *Tweet                 `protobuf:"bytes,1,opt,name=tweet,proto3" json:"tweet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTweetResponse) Reset() {
	*x = CreateTweetResponse{}
	mi := &file_twitter_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTweetResponse) ProtoMessage() {}

func (x *CreateTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_twitter_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTweetResponse.ProtoReflect.Descriptor instead.
func (*CreateTweetResponse) Descriptor() ([]byte, []int) {
	return file_twitter_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTweetResponse) GetTweet() *Tweet {
	if x != nil {
		return x.Tweet
	}
	return nil
}

type GetTweetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTweetRequest) Reset() {
	*x = GetTweetRequest{}
	mi := &file_twitter_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTweetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetRequest) ProtoMessage() {}

func (x *GetTweetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_twitter_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetRequest.ProtoReflect.Descriptor instead.
func (*GetTweetRequest) Descriptor() ([]byte, []int) {
	return file_twitter_proto_rawDescGZIP(), []int{4}
}

func (x *GetTweetRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTweetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tweet         *Tweet                 `protobuf:"bytes,1,opt,name=tweet,proto3" json:"tweet,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTweetResponse) Reset() {
	*x = GetTweetResponse{}
	mi := &file_twitter_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTweetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTweetResponse) ProtoMessage() {}

func (x *GetTweetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_twitter_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTweetResponse.ProtoReflect.Descriptor instead.
func (*GetTweetResponse) Descriptor() ([]byte, []int) {
	return file_twitter_proto_rawDescGZIP(), []int{5}
}

func (x *GetTweetResponse) GetTweet() *Tweet {
	if x != nil {
		return x.Tweet
	}
	return nil
}

var File_twitter_proto protoreflect.FileDescriptor

var file_twitter_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x22, 0x46, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x6c, 0x0a, 0x05, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4b,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x77, 0x65, 0x65,
	0x74, 0x52, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x74, 0x77, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x05,
	0x74, 0x77, 0x65, 0x65, 0x74, 0x32, 0x9b, 0x01, 0x0a, 0x0e, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x12, 0x18,
	0x2e, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x77, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x77, 0x65, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0e, 0x5a, 0x0c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x54, 0x77, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_twitter_proto_rawDescOnce sync.Once
	file_twitter_proto_rawDescData []byte
)

func file_twitter_proto_rawDescGZIP() []byte {
	file_twitter_proto_rawDescOnce.Do(func() {
		file_twitter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_twitter_proto_rawDesc), len(file_twitter_proto_rawDesc)))
	})
	return file_twitter_proto_rawDescData
}

var file_twitter_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_twitter_proto_goTypes = []any{
	(*User)(nil),                // 0: twitter.User
	(*Tweet)(nil),               // 1: twitter.Tweet
	(*CreateTweetRequest)(nil),  // 2: twitter.CreateTweetRequest
	(*CreateTweetResponse)(nil), // 3: twitter.CreateTweetResponse
	(*GetTweetRequest)(nil),     // 4: twitter.GetTweetRequest
	(*GetTweetResponse)(nil),    // 5: twitter.GetTweetResponse
}
var file_twitter_proto_depIdxs = []int32{
	0, // 0: twitter.Tweet.user:type_name -> twitter.User
	0, // 1: twitter.CreateTweetRequest.user:type_name -> twitter.User
	1, // 2: twitter.CreateTweetResponse.tweet:type_name -> twitter.Tweet
	1, // 3: twitter.GetTweetResponse.tweet:type_name -> twitter.Tweet
	2, // 4: twitter.TwitterService.CreateTweet:input_type -> twitter.CreateTweetRequest
	4, // 5: twitter.TwitterService.GetTweet:input_type -> twitter.GetTweetRequest
	3, // 6: twitter.TwitterService.CreateTweet:output_type -> twitter.CreateTweetResponse
	5, // 7: twitter.TwitterService.GetTweet:output_type -> twitter.GetTweetResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_twitter_proto_init() }
func file_twitter_proto_init() {
	if File_twitter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_twitter_proto_rawDesc), len(file_twitter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_twitter_proto_goTypes,
		DependencyIndexes: file_twitter_proto_depIdxs,
		MessageInfos:      file_twitter_proto_msgTypes,
	}.Build()
	File_twitter_proto = out.File
	file_twitter_proto_goTypes = nil
	file_twitter_proto_depIdxs = nil
}
