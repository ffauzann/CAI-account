// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: account.proto

package gen

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

type AccountCategory int32

const (
	AccountCategory_AC_UNSPECIFIED AccountCategory = 0
	AccountCategory_AC_DEBIT       AccountCategory = 1
	AccountCategory_AC_CREDIT      AccountCategory = 2
	AccountCategory_AC_LOAN        AccountCategory = 3
)

// Enum value maps for AccountCategory.
var (
	AccountCategory_name = map[int32]string{
		0: "AC_UNSPECIFIED",
		1: "AC_DEBIT",
		2: "AC_CREDIT",
		3: "AC_LOAN",
	}
	AccountCategory_value = map[string]int32{
		"AC_UNSPECIFIED": 0,
		"AC_DEBIT":       1,
		"AC_CREDIT":      2,
		"AC_LOAN":        3,
	}
)

func (x AccountCategory) Enum() *AccountCategory {
	p := new(AccountCategory)
	*p = x
	return p
}

func (x AccountCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_account_proto_enumTypes[0].Descriptor()
}

func (AccountCategory) Type() protoreflect.EnumType {
	return &file_account_proto_enumTypes[0]
}

func (x AccountCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountCategory.Descriptor instead.
func (AccountCategory) EnumDescriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64          `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category AccountCategory `protobuf:"varint,2,opt,name=category,proto3,enum=ffauzann.cai.account.AccountCategory" json:"category,omitempty"`
	Balance  float32         `protobuf:"fixed32,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetCategory() AccountCategory {
	if x != nil {
		return x.Category
	}
	return AccountCategory_AC_UNSPECIFIED
}

func (x *Account) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type GetListAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetListAccountRequest) Reset() {
	*x = GetListAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAccountRequest) ProtoMessage() {}

func (x *GetListAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAccountRequest.ProtoReflect.Descriptor instead.
func (*GetListAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{1}
}

type GetListAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GetListAccountResponse) Reset() {
	*x = GetListAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAccountResponse) ProtoMessage() {}

func (x *GetListAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAccountResponse.ProtoReflect.Descriptor instead.
func (*GetListAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{2}
}

func (x *GetListAccountResponse) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type UpdateBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceAccountId      uint64  `protobuf:"varint,1,opt,name=source_account_id,json=sourceAccountId,proto3" json:"source_account_id,omitempty"`
	DestinationAccountId uint64  `protobuf:"varint,2,opt,name=destination_account_id,json=destinationAccountId,proto3" json:"destination_account_id,omitempty"`
	Amount               float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *UpdateBalanceRequest) Reset() {
	*x = UpdateBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBalanceRequest) ProtoMessage() {}

func (x *UpdateBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBalanceRequest.ProtoReflect.Descriptor instead.
func (*UpdateBalanceRequest) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBalanceRequest) GetSourceAccountId() uint64 {
	if x != nil {
		return x.SourceAccountId
	}
	return 0
}

func (x *UpdateBalanceRequest) GetDestinationAccountId() uint64 {
	if x != nil {
		return x.DestinationAccountId
	}
	return 0
}

func (x *UpdateBalanceRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type UpdateBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateBalanceResponse) Reset() {
	*x = UpdateBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBalanceResponse) ProtoMessage() {}

func (x *UpdateBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBalanceResponse.ProtoReflect.Descriptor instead.
func (*UpdateBalanceResponse) Descriptor() ([]byte, []int) {
	return file_account_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBalanceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_account_proto protoreflect.FileDescriptor

var file_account_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x76, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x41, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61,
	0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x53, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61,
	0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2f,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a,
	0x4f, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x43, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x5f, 0x44, 0x45, 0x42,
	0x49, 0x54, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x43, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49,
	0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x5f, 0x4c, 0x4f, 0x41, 0x4e, 0x10, 0x03,
	0x32, 0xeb, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e,
	0x2e, 0x63, 0x61, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61,
	0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6a, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x2a, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63,
	0x61, 0x69, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x66, 0x66, 0x61, 0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2e, 0x63, 0x61, 0x69, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x66, 0x61,
	0x75, 0x7a, 0x61, 0x6e, 0x6e, 0x2f, 0x43, 0x41, 0x49, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_account_proto_rawDescOnce sync.Once
	file_account_proto_rawDescData = file_account_proto_rawDesc
)

func file_account_proto_rawDescGZIP() []byte {
	file_account_proto_rawDescOnce.Do(func() {
		file_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_proto_rawDescData)
	})
	return file_account_proto_rawDescData
}

var file_account_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_account_proto_goTypes = []interface{}{
	(AccountCategory)(0),           // 0: ffauzann.cai.account.AccountCategory
	(*Account)(nil),                // 1: ffauzann.cai.account.Account
	(*GetListAccountRequest)(nil),  // 2: ffauzann.cai.account.GetListAccountRequest
	(*GetListAccountResponse)(nil), // 3: ffauzann.cai.account.GetListAccountResponse
	(*UpdateBalanceRequest)(nil),   // 4: ffauzann.cai.account.UpdateBalanceRequest
	(*UpdateBalanceResponse)(nil),  // 5: ffauzann.cai.account.UpdateBalanceResponse
}
var file_account_proto_depIdxs = []int32{
	0, // 0: ffauzann.cai.account.Account.category:type_name -> ffauzann.cai.account.AccountCategory
	1, // 1: ffauzann.cai.account.GetListAccountResponse.accounts:type_name -> ffauzann.cai.account.Account
	2, // 2: ffauzann.cai.account.AccountService.GetListAccount:input_type -> ffauzann.cai.account.GetListAccountRequest
	4, // 3: ffauzann.cai.account.AccountService.UpdateBalance:input_type -> ffauzann.cai.account.UpdateBalanceRequest
	3, // 4: ffauzann.cai.account.AccountService.GetListAccount:output_type -> ffauzann.cai.account.GetListAccountResponse
	5, // 5: ffauzann.cai.account.AccountService.UpdateBalance:output_type -> ffauzann.cai.account.UpdateBalanceResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_proto_init() }
func file_account_proto_init() {
	if File_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAccountRequest); i {
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
		file_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAccountResponse); i {
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
		file_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBalanceRequest); i {
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
		file_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBalanceResponse); i {
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
			RawDescriptor: file_account_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_proto_goTypes,
		DependencyIndexes: file_account_proto_depIdxs,
		EnumInfos:         file_account_proto_enumTypes,
		MessageInfos:      file_account_proto_msgTypes,
	}.Build()
	File_account_proto = out.File
	file_account_proto_rawDesc = nil
	file_account_proto_goTypes = nil
	file_account_proto_depIdxs = nil
}