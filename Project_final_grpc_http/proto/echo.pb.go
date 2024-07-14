package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[0]
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
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{0}
}

type CreateAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Amount   int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateAccountReq) Reset() {
	*x = CreateAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountReq) ProtoMessage() {}

func (x *CreateAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountReq) Descriptor() ([]byte, []int) {
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAccountReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateAccountReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type PatchAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Amount   int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *PatchAccountReq) Reset() {
	*x = PatchAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatchAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatchAccountReq) ProtoMessage() {}

func (x *PatchAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatchAccountRequest.ProtoReflect.Descriptor instead.
func (*PatchAccountReq) Descriptor() ([]byte, []int) {
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{2}
}

func (x *PatchAccountReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *PatchAccountReq) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ChangeAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastName string `protobuf:"bytes,1,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	NewName  string `protobuf:"bytes,2,opt,name=new_name,json=newName,proto3" json:"new_name,omitempty"`
}

func (x *ChangeAccountReq) Reset() {
	*x = ChangeAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeAccountReq) ProtoMessage() {}

func (x *ChangeAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeAccountRequest.ProtoReflect.Descriptor instead.
func (*ChangeAccountReq) Descriptor() ([]byte, []int) {
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{3}
}

func (x *ChangeAccountReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *ChangeAccountReq) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type DeleteAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DeleteAccountReq) Reset() {
	*x = DeleteAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccountReq) ProtoMessage() {}

func (x *DeleteAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccountRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccountReq) Descriptor() ([]byte, []int) {
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteAccountReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GetAccountReq) Reset() {
	*x = GetAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReq) ProtoMessage() {}

func (x *GetAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountRequest.ProtoReflect.Descriptor instead.
func (*GetAccountReq) Descriptor() ([]byte, []int) {
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{5}
}

func (x *GetAccountReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GetAccountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Amount   int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetAccountResp) Reset() {
	*x = GetAccountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fasthttp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountResp) ProtoMessage() {}

func (x *GetAccountResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fasthttp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountResponse.ProtoReflect.Descriptor instead.
func (*GetAccountResp) Descriptor() ([]byte, []int) {
	return file_proto_fasthttp_proto_rawDescGZIP(), []int{6}
}

func (x *GetAccountResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetAccountResp) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_proto_fasthttp_proto protoreflect.FileDescriptor

var file_proto_fasthttp_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x61, 0x73, 0x74, 0x68, 0x74, 0x74, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x4a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x49, 0x0a, 0x13, 0x50, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4e, 0x0a,
	0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a,
	0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc5, 0x02, 0x0a,
	0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x63, 0x68,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x61, 0x74, 0x63, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x43, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x53, 0x75, 0x6d, 0x6d, 0x65, 0x72, 0x48, 0x53,
	0x45, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fasthttp_proto_rawDescOnce sync.Once
	file_proto_fasthttp_proto_rawDescData = file_proto_fasthttp_proto_rawDesc
)

func file_proto_fasthttp_proto_rawDescGZIP() []byte {
	file_proto_fasthttp_proto_rawDescOnce.Do(func() {
		file_proto_fasthttp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fasthttp_proto_rawDescData)
	})
	return file_proto_fasthttp_proto_rawDescData
}

var file_proto_fasthttp_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_fasthttp_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: proto.Empty
	(*CreateAccountReq)(nil), // 1: proto.CreateAccountRequest
	(*PatchAccountReq)(nil),  // 2: proto.PatchAccountRequest
	(*ChangeAccountReq)(nil), // 3: proto.ChangeAccountRequest
	(*DeleteAccountReq)(nil), // 4: proto.DeleteAccountRequest
	(*GetAccountReq)(nil),    // 5: proto.GetAccountRequest
	(*GetAccountResp)(nil),   // 6: proto.GetAccountResponse
}
var file_proto_fasthttp_proto_depIdxs = []int32{
	1, // 0: proto.Accounts.CreateAccount:input_type -> proto.CreateAccountRequest
	2, // 1: proto.Accounts.PatchAccount:input_type -> proto.PatchAccountRequest
	3, // 2: proto.Accounts.ChangeAccount:input_type -> proto.ChangeAccountRequest
	4, // 3: proto.Accounts.DeleteAccount:input_type -> proto.DeleteAccountRequest
	5, // 4: proto.Accounts.GetAccount:input_type -> proto.GetAccountRequest
	0, // 5: proto.Accounts.CreateAccount:output_type -> proto.Empty
	0, // 6: proto.Accounts.PatchAccount:output_type -> proto.Empty
	0, // 7: proto.Accounts.ChangeAccount:output_type -> proto.Empty
	0, // 8: proto.Accounts.DeleteAccount:output_type -> proto.Empty
	6, // 9: proto.Accounts.GetAccount:output_type -> proto.GetAccountResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_fasthttp_proto_init() }
func file_proto_fasthttp_proto_init() {
	if File_proto_fasthttp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_fasthttp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_fasthttp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountReq); i {
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
		file_proto_fasthttp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatchAccountReq); i {
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
		file_proto_fasthttp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeAccountReq); i {
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
		file_proto_fasthttp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccountReq); i {
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
		file_proto_fasthttp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountReq); i {
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
		file_proto_fasthttp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountResp); i {
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
			RawDescriptor: file_proto_fasthttp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fasthttp_proto_goTypes,
		DependencyIndexes: file_proto_fasthttp_proto_depIdxs,
		MessageInfos:      file_proto_fasthttp_proto_msgTypes,
	}.Build()
	File_proto_fasthttp_proto = out.File
	file_proto_fasthttp_proto_rawDesc = nil
	file_proto_fasthttp_proto_goTypes = nil
	file_proto_fasthttp_proto_depIdxs = nil
}
