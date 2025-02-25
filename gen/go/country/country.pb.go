// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: country.proto

package country

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

type Add_Country_Request struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CountryTitle   string                 `protobuf:"bytes,1,opt,name=country_title,json=countryTitle,proto3" json:"country_title,omitempty"`
	CountryCapital string                 `protobuf:"bytes,2,opt,name=country_capital,json=countryCapital,proto3" json:"country_capital,omitempty"`
	CountryArea    string                 `protobuf:"bytes,3,opt,name=country_area,json=countryArea,proto3" json:"country_area,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Add_Country_Request) Reset() {
	*x = Add_Country_Request{}
	mi := &file_country_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Add_Country_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Add_Country_Request) ProtoMessage() {}

func (x *Add_Country_Request) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Add_Country_Request.ProtoReflect.Descriptor instead.
func (*Add_Country_Request) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{0}
}

func (x *Add_Country_Request) GetCountryTitle() string {
	if x != nil {
		return x.CountryTitle
	}
	return ""
}

func (x *Add_Country_Request) GetCountryCapital() string {
	if x != nil {
		return x.CountryCapital
	}
	return ""
}

func (x *Add_Country_Request) GetCountryArea() string {
	if x != nil {
		return x.CountryArea
	}
	return ""
}

type Add_Country_Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CountryId     int64                  `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Add_Country_Response) Reset() {
	*x = Add_Country_Response{}
	mi := &file_country_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Add_Country_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Add_Country_Response) ProtoMessage() {}

func (x *Add_Country_Response) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Add_Country_Response.ProtoReflect.Descriptor instead.
func (*Add_Country_Response) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{1}
}

func (x *Add_Country_Response) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

type Get_CountryById_Requset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CountryId     int64                  `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Get_CountryById_Requset) Reset() {
	*x = Get_CountryById_Requset{}
	mi := &file_country_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get_CountryById_Requset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_CountryById_Requset) ProtoMessage() {}

func (x *Get_CountryById_Requset) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_CountryById_Requset.ProtoReflect.Descriptor instead.
func (*Get_CountryById_Requset) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{2}
}

func (x *Get_CountryById_Requset) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

type Get_CountryById_Response struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CountryTitle   string                 `protobuf:"bytes,1,opt,name=country_title,json=countryTitle,proto3" json:"country_title,omitempty"`
	CountryCapital string                 `protobuf:"bytes,2,opt,name=country_capital,json=countryCapital,proto3" json:"country_capital,omitempty"`
	CountryArea    string                 `protobuf:"bytes,3,opt,name=country_area,json=countryArea,proto3" json:"country_area,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Get_CountryById_Response) Reset() {
	*x = Get_CountryById_Response{}
	mi := &file_country_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get_CountryById_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_CountryById_Response) ProtoMessage() {}

func (x *Get_CountryById_Response) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_CountryById_Response.ProtoReflect.Descriptor instead.
func (*Get_CountryById_Response) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{3}
}

func (x *Get_CountryById_Response) GetCountryTitle() string {
	if x != nil {
		return x.CountryTitle
	}
	return ""
}

func (x *Get_CountryById_Response) GetCountryCapital() string {
	if x != nil {
		return x.CountryCapital
	}
	return ""
}

func (x *Get_CountryById_Response) GetCountryArea() string {
	if x != nil {
		return x.CountryArea
	}
	return ""
}

type Get_All_Country_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Get_All_Country_Request) Reset() {
	*x = Get_All_Country_Request{}
	mi := &file_country_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get_All_Country_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_All_Country_Request) ProtoMessage() {}

func (x *Get_All_Country_Request) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_All_Country_Request.ProtoReflect.Descriptor instead.
func (*Get_All_Country_Request) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{4}
}

type Get_All_Country_Response struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	Countries     []*Get_CountryById_Response `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Get_All_Country_Response) Reset() {
	*x = Get_All_Country_Response{}
	mi := &file_country_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Get_All_Country_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Get_All_Country_Response) ProtoMessage() {}

func (x *Get_All_Country_Response) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Get_All_Country_Response.ProtoReflect.Descriptor instead.
func (*Get_All_Country_Response) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{5}
}

func (x *Get_All_Country_Response) GetCountries() []*Get_CountryById_Response {
	if x != nil {
		return x.Countries
	}
	return nil
}

type Update_CountryById_Request struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CountryId      int64                  `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	CountryTitle   string                 `protobuf:"bytes,2,opt,name=country_title,json=countryTitle,proto3" json:"country_title,omitempty"`
	CountryCapital string                 `protobuf:"bytes,3,opt,name=country_capital,json=countryCapital,proto3" json:"country_capital,omitempty"`
	CountryArea    string                 `protobuf:"bytes,4,opt,name=country_area,json=countryArea,proto3" json:"country_area,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Update_CountryById_Request) Reset() {
	*x = Update_CountryById_Request{}
	mi := &file_country_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Update_CountryById_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update_CountryById_Request) ProtoMessage() {}

func (x *Update_CountryById_Request) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update_CountryById_Request.ProtoReflect.Descriptor instead.
func (*Update_CountryById_Request) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{6}
}

func (x *Update_CountryById_Request) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

func (x *Update_CountryById_Request) GetCountryTitle() string {
	if x != nil {
		return x.CountryTitle
	}
	return ""
}

func (x *Update_CountryById_Request) GetCountryCapital() string {
	if x != nil {
		return x.CountryCapital
	}
	return ""
}

func (x *Update_CountryById_Request) GetCountryArea() string {
	if x != nil {
		return x.CountryArea
	}
	return ""
}

type Update_CountryById_Response struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CountryTitle   string                 `protobuf:"bytes,1,opt,name=country_title,json=countryTitle,proto3" json:"country_title,omitempty"`
	CountryCapital string                 `protobuf:"bytes,2,opt,name=country_capital,json=countryCapital,proto3" json:"country_capital,omitempty"`
	CountryArea    string                 `protobuf:"bytes,3,opt,name=country_area,json=countryArea,proto3" json:"country_area,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Update_CountryById_Response) Reset() {
	*x = Update_CountryById_Response{}
	mi := &file_country_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Update_CountryById_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Update_CountryById_Response) ProtoMessage() {}

func (x *Update_CountryById_Response) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Update_CountryById_Response.ProtoReflect.Descriptor instead.
func (*Update_CountryById_Response) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{7}
}

func (x *Update_CountryById_Response) GetCountryTitle() string {
	if x != nil {
		return x.CountryTitle
	}
	return ""
}

func (x *Update_CountryById_Response) GetCountryCapital() string {
	if x != nil {
		return x.CountryCapital
	}
	return ""
}

func (x *Update_CountryById_Response) GetCountryArea() string {
	if x != nil {
		return x.CountryArea
	}
	return ""
}

type Delete_CountryById_Request struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CountryId     int64                  `protobuf:"varint,1,opt,name=country_id,json=countryId,proto3" json:"country_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Delete_CountryById_Request) Reset() {
	*x = Delete_CountryById_Request{}
	mi := &file_country_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Delete_CountryById_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delete_CountryById_Request) ProtoMessage() {}

func (x *Delete_CountryById_Request) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delete_CountryById_Request.ProtoReflect.Descriptor instead.
func (*Delete_CountryById_Request) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{8}
}

func (x *Delete_CountryById_Request) GetCountryId() int64 {
	if x != nil {
		return x.CountryId
	}
	return 0
}

type Delete_CountryById_Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CountryTitle  string                 `protobuf:"bytes,1,opt,name=country_title,json=countryTitle,proto3" json:"country_title,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Delete_CountryById_Response) Reset() {
	*x = Delete_CountryById_Response{}
	mi := &file_country_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Delete_CountryById_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Delete_CountryById_Response) ProtoMessage() {}

func (x *Delete_CountryById_Response) ProtoReflect() protoreflect.Message {
	mi := &file_country_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Delete_CountryById_Response.ProtoReflect.Descriptor instead.
func (*Delete_CountryById_Response) Descriptor() ([]byte, []int) {
	return file_country_proto_rawDescGZIP(), []int{9}
}

func (x *Delete_CountryById_Response) GetCountryTitle() string {
	if x != nil {
		return x.CountryTitle
	}
	return ""
}

var File_country_proto protoreflect.FileDescriptor

var file_country_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64,
	0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x72, 0x65,
	0x61, 0x22, 0x35, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x5f,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x73, 0x65, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x22, 0x8d, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x5f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61,
	0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x65,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x41, 0x72, 0x65, 0x61, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x5f,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x5d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x5f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0xae,
	0x01, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63,
	0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x72, 0x65, 0x61, 0x22,
	0x90, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x61, 0x70, 0x69, 0x74, 0x61, 0x6c, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x41, 0x72,
	0x65, 0x61, 0x22, 0x3d, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49,
	0x64, 0x22, 0x44, 0x0a, 0x1d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x32, 0xd9, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x4a, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64,
	0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x5f, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x73, 0x65, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x5f, 0x41, 0x6c, 0x6c, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x20, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x5f, 0x41,
	0x6c, 0x6c, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x5f, 0x41, 0x6c, 0x6c, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x25, 0x2e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x12, 0x25, 0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x4c, 0x69, 0x76, 0x65, 0x69, 0x73, 0x46, 0x50, 0x56,
	0x2e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_country_proto_rawDescOnce sync.Once
	file_country_proto_rawDescData []byte
)

func file_country_proto_rawDescGZIP() []byte {
	file_country_proto_rawDescOnce.Do(func() {
		file_country_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_country_proto_rawDesc), len(file_country_proto_rawDesc)))
	})
	return file_country_proto_rawDescData
}

var file_country_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_country_proto_goTypes = []any{
	(*Add_Country_Request)(nil),         // 0: country.Add_Country_Request
	(*Add_Country_Response)(nil),        // 1: country.Add_Country_Response
	(*Get_CountryById_Requset)(nil),     // 2: country.Get_Country_by_id_Requset
	(*Get_CountryById_Response)(nil),    // 3: country.Get_Country_by_id_Response
	(*Get_All_Country_Request)(nil),     // 4: country.Get_All_Country_Request
	(*Get_All_Country_Response)(nil),    // 5: country.Get_All_Country_Response
	(*Update_CountryById_Request)(nil),  // 6: country.Update_Country_by_id_Request
	(*Update_CountryById_Response)(nil), // 7: country.Update_Country_by_id_Response
	(*Delete_CountryById_Request)(nil),  // 8: country.Delete_Country_by_id_Request
	(*Delete_CountryById_Response)(nil), // 9: country.Delete_Country_by_id_Response
}
var file_country_proto_depIdxs = []int32{
	3, // 0: country.Get_All_Country_Response.countries:type_name -> country.Get_Country_by_id_Response
	0, // 1: country.Country.Add_Country:input_type -> country.Add_Country_Request
	2, // 2: country.Country.Get_Country_by_id:input_type -> country.Get_Country_by_id_Requset
	4, // 3: country.Country.Get_All_Country:input_type -> country.Get_All_Country_Request
	6, // 4: country.Country.Update_Country_by_id:input_type -> country.Update_Country_by_id_Request
	8, // 5: country.Country.Delete_Country_by_id:input_type -> country.Delete_Country_by_id_Request
	1, // 6: country.Country.Add_Country:output_type -> country.Add_Country_Response
	3, // 7: country.Country.Get_Country_by_id:output_type -> country.Get_Country_by_id_Response
	5, // 8: country.Country.Get_All_Country:output_type -> country.Get_All_Country_Response
	7, // 9: country.Country.Update_Country_by_id:output_type -> country.Update_Country_by_id_Response
	9, // 10: country.Country.Delete_Country_by_id:output_type -> country.Delete_Country_by_id_Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_country_proto_init() }
func file_country_proto_init() {
	if File_country_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_country_proto_rawDesc), len(file_country_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_country_proto_goTypes,
		DependencyIndexes: file_country_proto_depIdxs,
		MessageInfos:      file_country_proto_msgTypes,
	}.Build()
	File_country_proto = out.File
	file_country_proto_goTypes = nil
	file_country_proto_depIdxs = nil
}
