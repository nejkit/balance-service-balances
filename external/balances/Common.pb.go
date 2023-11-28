// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: Common.proto

package balances

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

type LockBalanceStatus int32

const (
	LockBalanceStatus_UNRECOGNIZED LockBalanceStatus = 0
	LockBalanceStatus_IN_PROCESS   LockBalanceStatus = 1
	LockBalanceStatus_DONE         LockBalanceStatus = 2
	LockBalanceStatus_REJECTED     LockBalanceStatus = 3
)

// Enum value maps for LockBalanceStatus.
var (
	LockBalanceStatus_name = map[int32]string{
		0: "UNRECOGNIZED",
		1: "IN_PROCESS",
		2: "DONE",
		3: "REJECTED",
	}
	LockBalanceStatus_value = map[string]int32{
		"UNRECOGNIZED": 0,
		"IN_PROCESS":   1,
		"DONE":         2,
		"REJECTED":     3,
	}
)

func (x LockBalanceStatus) Enum() *LockBalanceStatus {
	p := new(LockBalanceStatus)
	*p = x
	return p
}

func (x LockBalanceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockBalanceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[0].Descriptor()
}

func (LockBalanceStatus) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[0]
}

func (x LockBalanceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LockBalanceStatus.Descriptor instead.
func (LockBalanceStatus) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{0}
}

type TransferState int32

const (
	TransferState_TRANSFER_STATE_NONE        TransferState = 0
	TransferState_TRANSFER_STATE_NEW         TransferState = 1
	TransferState_TRANSFER_STATE_IN_PROGRESS TransferState = 2
	TransferState_TRANSFER_STATE_DONE        TransferState = 3
	TransferState_TRANSFER_STATE_REJECT      TransferState = 4
)

// Enum value maps for TransferState.
var (
	TransferState_name = map[int32]string{
		0: "TRANSFER_STATE_NONE",
		1: "TRANSFER_STATE_NEW",
		2: "TRANSFER_STATE_IN_PROGRESS",
		3: "TRANSFER_STATE_DONE",
		4: "TRANSFER_STATE_REJECT",
	}
	TransferState_value = map[string]int32{
		"TRANSFER_STATE_NONE":        0,
		"TRANSFER_STATE_NEW":         1,
		"TRANSFER_STATE_IN_PROGRESS": 2,
		"TRANSFER_STATE_DONE":        3,
		"TRANSFER_STATE_REJECT":      4,
	}
)

func (x TransferState) Enum() *TransferState {
	p := new(TransferState)
	*p = x
	return p
}

func (x TransferState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransferState) Descriptor() protoreflect.EnumDescriptor {
	return file_Common_proto_enumTypes[1].Descriptor()
}

func (TransferState) Type() protoreflect.EnumType {
	return &file_Common_proto_enumTypes[1]
}

func (x TransferState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransferState.Descriptor instead.
func (TransferState) EnumDescriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{1}
}

type BalanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Currency      string  `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	ActualBalance float64 `protobuf:"fixed64,3,opt,name=actual_balance,json=actualBalance,proto3" json:"actual_balance,omitempty"`
	FreezeBalance float64 `protobuf:"fixed64,4,opt,name=freeze_balance,json=freezeBalance,proto3" json:"freeze_balance,omitempty"`
}

func (x *BalanceInfo) Reset() {
	*x = BalanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceInfo) ProtoMessage() {}

func (x *BalanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceInfo.ProtoReflect.Descriptor instead.
func (*BalanceInfo) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{0}
}

func (x *BalanceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BalanceInfo) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *BalanceInfo) GetActualBalance() float64 {
	if x != nil {
		return x.ActualBalance
	}
	return 0
}

func (x *BalanceInfo) GetFreezeBalance() float64 {
	if x != nil {
		return x.FreezeBalance
	}
	return 0
}

type WalletInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Created      uint64         `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	IsDeleted    bool           `protobuf:"varint,3,opt,name=is_deleted,json=isDeleted,proto3" json:"is_deleted,omitempty"`
	BalanceInfos []*BalanceInfo `protobuf:"bytes,4,rep,name=balance_infos,json=balanceInfos,proto3" json:"balance_infos,omitempty"`
}

func (x *WalletInfo) Reset() {
	*x = WalletInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WalletInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WalletInfo) ProtoMessage() {}

func (x *WalletInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WalletInfo.ProtoReflect.Descriptor instead.
func (*WalletInfo) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{1}
}

func (x *WalletInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *WalletInfo) GetCreated() uint64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *WalletInfo) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *WalletInfo) GetBalanceInfos() []*BalanceInfo {
	if x != nil {
		return x.BalanceInfos
	}
	return nil
}

type GetWalletInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *GetWalletInfoRequest) Reset() {
	*x = GetWalletInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletInfoRequest) ProtoMessage() {}

func (x *GetWalletInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletInfoRequest.ProtoReflect.Descriptor instead.
func (*GetWalletInfoRequest) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{2}
}

func (x *GetWalletInfoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetWalletInfoRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetWalletInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WalletInfo *WalletInfo `protobuf:"bytes,2,opt,name=wallet_info,json=walletInfo,proto3" json:"wallet_info,omitempty"`
}

func (x *GetWalletInfoResponse) Reset() {
	*x = GetWalletInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWalletInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWalletInfoResponse) ProtoMessage() {}

func (x *GetWalletInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWalletInfoResponse.ProtoReflect.Descriptor instead.
func (*GetWalletInfoResponse) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{3}
}

func (x *GetWalletInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetWalletInfoResponse) GetWalletInfo() *WalletInfo {
	if x != nil {
		return x.WalletInfo
	}
	return nil
}

type LockBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address  string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Currency string  `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float32 `protobuf:"fixed32,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *LockBalanceRequest) Reset() {
	*x = LockBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockBalanceRequest) ProtoMessage() {}

func (x *LockBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockBalanceRequest.ProtoReflect.Descriptor instead.
func (*LockBalanceRequest) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{4}
}

func (x *LockBalanceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LockBalanceRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *LockBalanceRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *LockBalanceRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type LockBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	State        LockBalanceStatus `protobuf:"varint,2,opt,name=state,proto3,enum=LockBalanceStatus" json:"state,omitempty"`
	ErrorMessage *ErrorMessage     `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
}

func (x *LockBalanceResponse) Reset() {
	*x = LockBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockBalanceResponse) ProtoMessage() {}

func (x *LockBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockBalanceResponse.ProtoReflect.Descriptor instead.
func (*LockBalanceResponse) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{5}
}

func (x *LockBalanceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LockBalanceResponse) GetState() LockBalanceStatus {
	if x != nil {
		return x.State
	}
	return LockBalanceStatus_UNRECOGNIZED
}

func (x *LockBalanceResponse) GetErrorMessage() *ErrorMessage {
	if x != nil {
		return x.ErrorMessage
	}
	return nil
}

type TransferOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Currency string  `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	Amount   float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransferOptions) Reset() {
	*x = TransferOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferOptions) ProtoMessage() {}

func (x *TransferOptions) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferOptions.ProtoReflect.Descriptor instead.
func (*TransferOptions) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{6}
}

func (x *TransferOptions) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TransferOptions) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferOptions) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateTransferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderData    *TransferOptions `protobuf:"bytes,2,opt,name=sender_data,json=senderData,proto3" json:"sender_data,omitempty"`
	RecepientData *TransferOptions `protobuf:"bytes,3,opt,name=recepient_data,json=recepientData,proto3" json:"recepient_data,omitempty"`
}

func (x *CreateTransferRequest) Reset() {
	*x = CreateTransferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTransferRequest) ProtoMessage() {}

func (x *CreateTransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTransferRequest.ProtoReflect.Descriptor instead.
func (*CreateTransferRequest) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTransferRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTransferRequest) GetSenderData() *TransferOptions {
	if x != nil {
		return x.SenderData
	}
	return nil
}

func (x *CreateTransferRequest) GetRecepientData() *TransferOptions {
	if x != nil {
		return x.RecepientData
	}
	return nil
}

type Transfer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderData    *TransferOptions `protobuf:"bytes,2,opt,name=sender_data,json=senderData,proto3" json:"sender_data,omitempty"`
	RecepientData *TransferOptions `protobuf:"bytes,3,opt,name=recepient_data,json=recepientData,proto3" json:"recepient_data,omitempty"`
	State         TransferState    `protobuf:"varint,4,opt,name=state,proto3,enum=TransferState" json:"state,omitempty"`
	Error         *ErrorMessage    `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Transfer) Reset() {
	*x = Transfer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Common_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transfer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transfer) ProtoMessage() {}

func (x *Transfer) ProtoReflect() protoreflect.Message {
	mi := &file_Common_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transfer.ProtoReflect.Descriptor instead.
func (*Transfer) Descriptor() ([]byte, []int) {
	return file_Common_proto_rawDescGZIP(), []int{8}
}

func (x *Transfer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Transfer) GetSenderData() *TransferOptions {
	if x != nil {
		return x.SenderData
	}
	return nil
}

func (x *Transfer) GetRecepientData() *TransferOptions {
	if x != nil {
		return x.RecepientData
	}
	return nil
}

func (x *Transfer) GetState() TransferState {
	if x != nil {
		return x.State
	}
	return TransferState_TRANSFER_STATE_NONE
}

func (x *Transfer) GetError() *ErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_Common_proto protoreflect.FileDescriptor

var file_Common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a,
	0x0b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x63, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x61, 0x63, 0x74, 0x75, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x40, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x55, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x72, 0x0a, 0x12, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x13, 0x4c, 0x6f, 0x63,
	0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x12, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5f,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x93, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0e,
	0x72, 0x65, 0x63, 0x65, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x70, 0x69, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0xd1, 0x01, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x37, 0x0a, 0x0e, 0x72, 0x65, 0x63, 0x65, 0x70, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0d, 0x72, 0x65, 0x63, 0x65, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x4d, 0x0a, 0x11, 0x4c, 0x6f, 0x63,
	0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10,
	0x0a, 0x0c, 0x55, 0x4e, 0x52, 0x45, 0x43, 0x4f, 0x47, 0x4e, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x44, 0x4f, 0x4e, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45,
	0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x94, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x52,
	0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x57, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e,
	0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x54,
	0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x4f,
	0x4e, 0x45, 0x10, 0x03, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x45, 0x52,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x10, 0x04, 0x42,
	0x0b, 0x5a, 0x09, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Common_proto_rawDescOnce sync.Once
	file_Common_proto_rawDescData = file_Common_proto_rawDesc
)

func file_Common_proto_rawDescGZIP() []byte {
	file_Common_proto_rawDescOnce.Do(func() {
		file_Common_proto_rawDescData = protoimpl.X.CompressGZIP(file_Common_proto_rawDescData)
	})
	return file_Common_proto_rawDescData
}

var file_Common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Common_proto_goTypes = []interface{}{
	(LockBalanceStatus)(0),        // 0: LockBalanceStatus
	(TransferState)(0),            // 1: TransferState
	(*BalanceInfo)(nil),           // 2: BalanceInfo
	(*WalletInfo)(nil),            // 3: WalletInfo
	(*GetWalletInfoRequest)(nil),  // 4: GetWalletInfoRequest
	(*GetWalletInfoResponse)(nil), // 5: GetWalletInfoResponse
	(*LockBalanceRequest)(nil),    // 6: LockBalanceRequest
	(*LockBalanceResponse)(nil),   // 7: LockBalanceResponse
	(*TransferOptions)(nil),       // 8: TransferOptions
	(*CreateTransferRequest)(nil), // 9: CreateTransferRequest
	(*Transfer)(nil),              // 10: Transfer
	(*ErrorMessage)(nil),          // 11: ErrorMessage
}
var file_Common_proto_depIdxs = []int32{
	2,  // 0: WalletInfo.balance_infos:type_name -> BalanceInfo
	3,  // 1: GetWalletInfoResponse.wallet_info:type_name -> WalletInfo
	0,  // 2: LockBalanceResponse.state:type_name -> LockBalanceStatus
	11, // 3: LockBalanceResponse.error_message:type_name -> ErrorMessage
	8,  // 4: CreateTransferRequest.sender_data:type_name -> TransferOptions
	8,  // 5: CreateTransferRequest.recepient_data:type_name -> TransferOptions
	8,  // 6: Transfer.sender_data:type_name -> TransferOptions
	8,  // 7: Transfer.recepient_data:type_name -> TransferOptions
	1,  // 8: Transfer.state:type_name -> TransferState
	11, // 9: Transfer.error:type_name -> ErrorMessage
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_Common_proto_init() }
func file_Common_proto_init() {
	if File_Common_proto != nil {
		return
	}
	file_errors_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceInfo); i {
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
		file_Common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WalletInfo); i {
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
		file_Common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWalletInfoRequest); i {
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
		file_Common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWalletInfoResponse); i {
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
		file_Common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockBalanceRequest); i {
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
		file_Common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockBalanceResponse); i {
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
		file_Common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferOptions); i {
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
		file_Common_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTransferRequest); i {
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
		file_Common_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transfer); i {
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
			RawDescriptor: file_Common_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Common_proto_goTypes,
		DependencyIndexes: file_Common_proto_depIdxs,
		EnumInfos:         file_Common_proto_enumTypes,
		MessageInfos:      file_Common_proto_msgTypes,
	}.Build()
	File_Common_proto = out.File
	file_Common_proto_rawDesc = nil
	file_Common_proto_goTypes = nil
	file_Common_proto_depIdxs = nil
}