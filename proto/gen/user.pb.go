// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: user.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CloseUserAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CloseUserAccountResponse) Reset() {
	*x = CloseUserAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseUserAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseUserAccountResponse) ProtoMessage() {}

func (x *CloseUserAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseUserAccountResponse.ProtoReflect.Descriptor instead.
func (*CloseUserAccountResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *CloseUserAccountResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CloseUserAccountResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdatePasscodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPasscode string `protobuf:"bytes,1,opt,name=old_passcode,json=oldPasscode,proto3" json:"old_passcode,omitempty"`
	NewPasscode string `protobuf:"bytes,2,opt,name=new_passcode,json=newPasscode,proto3" json:"new_passcode,omitempty"`
}

func (x *UpdatePasscodeRequest) Reset() {
	*x = UpdatePasscodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasscodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasscodeRequest) ProtoMessage() {}

func (x *UpdatePasscodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasscodeRequest.ProtoReflect.Descriptor instead.
func (*UpdatePasscodeRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePasscodeRequest) GetOldPasscode() string {
	if x != nil {
		return x.OldPasscode
	}
	return ""
}

func (x *UpdatePasscodeRequest) GetNewPasscode() string {
	if x != nil {
		return x.NewPasscode
	}
	return ""
}

type UpdatePasscodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdatePasscodeResponse) Reset() {
	*x = UpdatePasscodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePasscodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePasscodeResponse) ProtoMessage() {}

func (x *UpdatePasscodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePasscodeResponse.ProtoReflect.Descriptor instead.
func (*UpdatePasscodeResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePasscodeResponse) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdatePasscodeResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type IsUserExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email       string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,2,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *IsUserExistRequest) Reset() {
	*x = IsUserExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExistRequest) ProtoMessage() {}

func (x *IsUserExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExistRequest.ProtoReflect.Descriptor instead.
func (*IsUserExistRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *IsUserExistRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *IsUserExistRequest) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type IsUserExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsExist bool     `protobuf:"varint,1,opt,name=is_exist,json=isExist,proto3" json:"is_exist,omitempty"`
	Reasons []string `protobuf:"bytes,2,rep,name=reasons,proto3" json:"reasons,omitempty"`
}

func (x *IsUserExistResponse) Reset() {
	*x = IsUserExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsUserExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsUserExistResponse) ProtoMessage() {}

func (x *IsUserExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsUserExistResponse.ProtoReflect.Descriptor instead.
func (*IsUserExistResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{4}
}

func (x *IsUserExistResponse) GetIsExist() bool {
	if x != nil {
		return x.IsExist
	}
	return false
}

func (x *IsUserExistResponse) GetReasons() []string {
	if x != nil {
		return x.Reasons
	}
	return nil
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x66, 0x66,
	0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x18,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x5d, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x49, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x4d, 0x0a, 0x12, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x4a, 0x0a, 0x13, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x32, 0xb1,
	0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59,
	0x0a, 0x10, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2b, 0x2e, 0x66, 0x66, 0x61,
	0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43,
	0x6c, 0x6f, 0x73, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x2e, 0x66, 0x66,
	0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e,
	0x2e, 0x63, 0x61, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x0b, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x12, 0x25, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x55, 0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a,
	0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2f, 0x43, 0x41, 0x49, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_proto_goTypes = []interface{}{
	(*CloseUserAccountResponse)(nil), // 0: ffauzann.cai.user.CloseUserAccountResponse
	(*UpdatePasscodeRequest)(nil),    // 1: ffauzann.cai.user.UpdatePasscodeRequest
	(*UpdatePasscodeResponse)(nil),   // 2: ffauzann.cai.user.UpdatePasscodeResponse
	(*IsUserExistRequest)(nil),       // 3: ffauzann.cai.user.IsUserExistRequest
	(*IsUserExistResponse)(nil),      // 4: ffauzann.cai.user.IsUserExistResponse
	(*emptypb.Empty)(nil),            // 5: google.protobuf.Empty
}
var file_user_proto_depIdxs = []int32{
	5, // 0: ffauzann.cai.user.UserService.CloseUserAccount:input_type -> google.protobuf.Empty
	1, // 1: ffauzann.cai.user.UserService.UpdatePasscode:input_type -> ffauzann.cai.user.UpdatePasscodeRequest
	3, // 2: ffauzann.cai.user.UserService.IsUserExist:input_type -> ffauzann.cai.user.IsUserExistRequest
	0, // 3: ffauzann.cai.user.UserService.CloseUserAccount:output_type -> ffauzann.cai.user.CloseUserAccountResponse
	2, // 4: ffauzann.cai.user.UserService.UpdatePasscode:output_type -> ffauzann.cai.user.UpdatePasscodeResponse
	4, // 5: ffauzann.cai.user.UserService.IsUserExist:output_type -> ffauzann.cai.user.IsUserExistResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseUserAccountResponse); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasscodeRequest); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePasscodeResponse); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserExistRequest); i {
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
		file_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsUserExistResponse); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
