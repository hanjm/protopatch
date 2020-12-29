// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: tests/lint/lint.proto

package lint

import (
	_ "github.com/alta/protopatch/patch/gopb"
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

type Protocol int32

const (
	// PROTOCOL_INVALID value should lint to ProtocolInvalid.
	ProtocolInvalid Protocol = 0
	// PROTOCOL_IP value should lint to ProtocolIP.
	ProtocolIP Protocol = 1
	// PROTOCOL_UDP value should lint to ProtocolUDP.
	ProtocolUDP Protocol = 2
	// PROTOCOL_TCP value should lint to ProtocolTCP.
	ProtocolTCP Protocol = 3
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "PROTOCOL_INVALID",
		1: "PROTOCOL_IP",
		2: "PROTOCOL_UDP",
		3: "PROTOCOL_TCP",
	}
	Protocol_value = map[string]int32{
		"PROTOCOL_INVALID": 0,
		"PROTOCOL_IP":      1,
		"PROTOCOL_UDP":     2,
		"PROTOCOL_TCP":     3,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_lint_lint_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_tests_lint_lint_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{0}
}

type Basic int32

const (
	// INVALID value should lint to BasicInvalid.
	BasicInvalid Basic = 0
	// A value should lint to BasicInvalid.
	BasicA Basic = 1
	// B value should lint to BasicInvalid.
	BasicB Basic = 2
	// C value should lint to BasicInvalid.
	BasicC Basic = 3
)

// Enum value maps for Basic.
var (
	Basic_name = map[int32]string{
		0: "INVALID",
		1: "A",
		2: "B",
		3: "C",
	}
	Basic_value = map[string]int32{
		"INVALID": 0,
		"A":       1,
		"B":       2,
		"C":       3,
	}
)

func (x Basic) Enum() *Basic {
	p := new(Basic)
	*p = x
	return p
}

func (x Basic) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Basic) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_lint_lint_proto_enumTypes[1].Descriptor()
}

func (Basic) Type() protoreflect.EnumType {
	return &file_tests_lint_lint_proto_enumTypes[1]
}

func (x Basic) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Basic.Descriptor instead.
func (Basic) EnumDescriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{1}
}

// Flavor should lint to OuterMessageFlavor.
type OuterMessageFlavor int32

const (
	OuterMessageFlavorInvalid OuterMessageFlavor = 0
	OuterMessageFlavorA       OuterMessageFlavor = 1
	OuterMessageFlavorB       OuterMessageFlavor = 2
	OuterMessageFlavorC       OuterMessageFlavor = 3
)

// Enum value maps for OuterMessage_Flavor.
var (
	OuterMessageFlavor_name = map[int32]string{
		0: "FLAVOR_INVALID",
		1: "FLAVOR_A",
		2: "FLAVOR_B",
		3: "FLAVOR_C",
	}
	OuterMessageFlavor_value = map[string]int32{
		"FLAVOR_INVALID": 0,
		"FLAVOR_A":       1,
		"FLAVOR_B":       2,
		"FLAVOR_C":       3,
	}
)

func (x OuterMessageFlavor) Enum() *OuterMessageFlavor {
	p := new(OuterMessageFlavor)
	*p = x
	return p
}

func (x OuterMessageFlavor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OuterMessageFlavor) Descriptor() protoreflect.EnumDescriptor {
	return file_tests_lint_lint_proto_enumTypes[2].Descriptor()
}

func (OuterMessageFlavor) Type() protoreflect.EnumType {
	return &file_tests_lint_lint_proto_enumTypes[2]
}

func (x OuterMessageFlavor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OuterMessage_Flavor.Descriptor instead.
func (OuterMessageFlavor) EnumDescriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{4, 0}
}

// URL message type should lint to URL.
type URL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *URL) Reset() {
	*x = URL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *URL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*URL) ProtoMessage() {}

func (x *URL) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Url.ProtoReflect.Descriptor instead.
func (*URL) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{0}
}

// ID message type should lint to ID.
type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{1}
}

// RgbColor should lint to RGBColor.
type RGBColor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RGBColor) Reset() {
	*x = RGBColor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RGBColor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RGBColor) ProtoMessage() {}

func (x *RGBColor) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RGBColor.ProtoReflect.Descriptor instead.
func (*RGBColor) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{2}
}

type OneofMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Contents:
	//	*OneofMessage_Id
	//	*OneofMessage_Url
	Contents isOneofMessage_Contents `protobuf_oneof:"contents"`
}

func (x *OneofMessage) Reset() {
	*x = OneofMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofMessage) ProtoMessage() {}

func (x *OneofMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofMessage.ProtoReflect.Descriptor instead.
func (*OneofMessage) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{3}
}

func (m *OneofMessage) GetContents() isOneofMessage_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *OneofMessage) GetID() *ID {
	if x, ok := x.GetContents().(*OneofMessage_ID); ok {
		return x.ID
	}
	return nil
}

func (x *OneofMessage) GetURL() *URL {
	if x, ok := x.GetContents().(*OneofMessage_URL); ok {
		return x.URL
	}
	return nil
}

type isOneofMessage_Contents interface {
	isOneofMessage_Contents()
}

type OneofMessage_ID struct {
	// ID oneof should lint to ID.
	ID *ID `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type OneofMessage_URL struct {
	// URL oneof should lint to URL.
	URL *URL `protobuf:"bytes,2,opt,name=url,proto3,oneof"`
}

func (*OneofMessage_ID) isOneofMessage_Contents() {}

func (*OneofMessage_URL) isOneofMessage_Contents() {}

type OuterMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id field should lint to ID.
	ID *OuterMessageInnerID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// url field should lint to URL.
	URL *OuterMessageInnerURL `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *OuterMessage) Reset() {
	*x = OuterMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessage) ProtoMessage() {}

func (x *OuterMessage) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage.ProtoReflect.Descriptor instead.
func (*OuterMessage) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{4}
}

func (x *OuterMessage) GetID() *OuterMessageInnerID {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *OuterMessage) GetURL() *OuterMessageInnerURL {
	if x != nil {
		return x.URL
	}
	return nil
}

type Color struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*Color_Rgb
	//	*Color_Rgba
	//	*Color_Hsv
	Value isColor_Value `protobuf_oneof:"value"`
}

func (x *Color) Reset() {
	*x = Color{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Color) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Color) ProtoMessage() {}

func (x *Color) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Color.ProtoReflect.Descriptor instead.
func (*Color) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{5}
}

func (m *Color) GetValue() isColor_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Color) GetRGB() string {
	if x, ok := x.GetValue().(*Color_RGB); ok {
		return x.RGB
	}
	return ""
}

func (x *Color) GetRGBA() string {
	if x, ok := x.GetValue().(*Color_RGBA); ok {
		return x.RGBA
	}
	return ""
}

func (x *Color) GetHSV() string {
	if x, ok := x.GetValue().(*Color_HSV); ok {
		return x.HSV
	}
	return ""
}

type isColor_Value interface {
	isColor_Value()
}

type Color_RGB struct {
	// rgb should lint to RGB.
	RGB string `protobuf:"bytes,1,opt,name=rgb,proto3,oneof"`
}

type Color_RGBA struct {
	// rgba should lint to RGBA.
	RGBA string `protobuf:"bytes,2,opt,name=rgba,proto3,oneof"`
}

type Color_HSV struct {
	// hsv should lint to HSV.
	HSV string `protobuf:"bytes,3,opt,name=hsv,proto3,oneof"`
}

func (*Color_RGB) isColor_Value() {}

func (*Color_RGBA) isColor_Value() {}

func (*Color_HSV) isColor_Value() {}

// IDFieldsTest message should lint to IDFieldsTest.
type IDFieldsTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id field should lint to ID.
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// customer_id field should lint to CustomerID.
	CustomerID string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// api_path field should lint to APIPath.
	APIPath string `protobuf:"bytes,3,opt,name=api_path,json=apiPath,proto3" json:"api_path,omitempty"`
}

func (x *IDFieldsTest) Reset() {
	*x = IDFieldsTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDFieldsTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDFieldsTest) ProtoMessage() {}

func (x *IDFieldsTest) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdFieldsTest.ProtoReflect.Descriptor instead.
func (*IDFieldsTest) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{6}
}

func (x *IDFieldsTest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *IDFieldsTest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *IDFieldsTest) GetAPIPath() string {
	if x != nil {
		return x.APIPath
	}
	return ""
}

// InnerId message should lint to InnerID.
type OuterMessageInnerID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OuterMessageInnerID) Reset() {
	*x = OuterMessageInnerID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessageInnerID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessageInnerID) ProtoMessage() {}

func (x *OuterMessageInnerID) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage_InnerId.ProtoReflect.Descriptor instead.
func (*OuterMessageInnerID) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{4, 0}
}

// InnerId message should lint to InnerURL.
type OuterMessageInnerURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OuterMessageInnerURL) Reset() {
	*x = OuterMessageInnerURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_lint_lint_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OuterMessageInnerURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OuterMessageInnerURL) ProtoMessage() {}

func (x *OuterMessageInnerURL) ProtoReflect() protoreflect.Message {
	mi := &file_tests_lint_lint_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OuterMessage_InnerUrl.ProtoReflect.Descriptor instead.
func (*OuterMessageInnerURL) Descriptor() ([]byte, []int) {
	return file_tests_lint_lint_proto_rawDescGZIP(), []int{4, 1}
}

var File_tests_lint_lint_proto protoreflect.FileDescriptor

var file_tests_lint_lint_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x74, 0x2f, 0x6c, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6c,
	0x69, 0x6e, 0x74, 0x1a, 0x0e, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x05, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x22, 0x04, 0x0a, 0x02, 0x49, 0x64,
	0x22, 0x0a, 0x0a, 0x08, 0x52, 0x47, 0x42, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0x61, 0x0a, 0x0c,
	0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x49, 0x64, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x55, 0x72, 0x6c, 0x48, 0x00, 0x52, 0x03,
	0x75, 0x72, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0xd4, 0x01, 0x0a, 0x0c, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x75, 0x74, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x33, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x6c, 0x69, 0x6e, 0x74, 0x2e, 0x4f, 0x75, 0x74,
	0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x55,
	0x72, 0x6c, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x09, 0x0a, 0x07, 0x49, 0x6e, 0x6e, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x0a, 0x0a, 0x08, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x46,
	0x0a, 0x06, 0x46, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x0e, 0x46, 0x4c, 0x41, 0x56,
	0x4f, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x46, 0x4c, 0x41, 0x56, 0x4f, 0x52, 0x5f, 0x41, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4c,
	0x41, 0x56, 0x4f, 0x52, 0x5f, 0x42, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x4c, 0x41, 0x56,
	0x4f, 0x52, 0x5f, 0x43, 0x10, 0x03, 0x22, 0x4e, 0x0a, 0x05, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x03, 0x72, 0x67, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03,
	0x72, 0x67, 0x62, 0x12, 0x14, 0x0a, 0x04, 0x72, 0x67, 0x62, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x72, 0x67, 0x62, 0x61, 0x12, 0x12, 0x0a, 0x03, 0x68, 0x73, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x68, 0x73, 0x76, 0x42, 0x07, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5a, 0x0a, 0x0c, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x54, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x69, 0x50, 0x61,
	0x74, 0x68, 0x2a, 0x55, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x14,
	0x0a, 0x10, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c,
	0x5f, 0x49, 0x50, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f,
	0x4c, 0x5f, 0x55, 0x44, 0x50, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x54, 0x4f,
	0x43, 0x4f, 0x4c, 0x5f, 0x54, 0x43, 0x50, 0x10, 0x03, 0x2a, 0x29, 0x0a, 0x05, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x05, 0x0a, 0x01, 0x41, 0x10, 0x01, 0x12, 0x05, 0x0a, 0x01, 0x42, 0x10, 0x02, 0x12, 0x05, 0x0a,
	0x01, 0x43, 0x10, 0x03, 0x42, 0x41, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6c, 0x69, 0x6e, 0x74, 0x90, 0xb6, 0x03,
	0x01, 0xe2, 0xb6, 0x03, 0x03, 0x52, 0x47, 0x42, 0xe2, 0xb6, 0x03, 0x04, 0x52, 0x47, 0x42, 0x41,
	0xe2, 0xb6, 0x03, 0x03, 0x48, 0x53, 0x56, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tests_lint_lint_proto_rawDescOnce sync.Once
	file_tests_lint_lint_proto_rawDescData = file_tests_lint_lint_proto_rawDesc
)

func file_tests_lint_lint_proto_rawDescGZIP() []byte {
	file_tests_lint_lint_proto_rawDescOnce.Do(func() {
		file_tests_lint_lint_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_lint_lint_proto_rawDescData)
	})
	return file_tests_lint_lint_proto_rawDescData
}

var file_tests_lint_lint_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_tests_lint_lint_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tests_lint_lint_proto_goTypes = []interface{}{
	(Protocol)(0),                // 0: tests.lint.Protocol
	(Basic)(0),                   // 1: tests.lint.Basic
	(OuterMessageFlavor)(0),      // 2: tests.lint.OuterMessage.Flavor
	(*URL)(nil),                  // 3: tests.lint.Url
	(*ID)(nil),                   // 4: tests.lint.Id
	(*RGBColor)(nil),             // 5: tests.lint.RGBColor
	(*OneofMessage)(nil),         // 6: tests.lint.OneofMessage
	(*OuterMessage)(nil),         // 7: tests.lint.OuterMessage
	(*Color)(nil),                // 8: tests.lint.Color
	(*IDFieldsTest)(nil),         // 9: tests.lint.IdFieldsTest
	(*OuterMessageInnerID)(nil),  // 10: tests.lint.OuterMessage.InnerId
	(*OuterMessageInnerURL)(nil), // 11: tests.lint.OuterMessage.InnerUrl
}
var file_tests_lint_lint_proto_depIdxs = []int32{
	4,  // 0: tests.lint.OneofMessage.id:type_name -> tests.lint.Id
	3,  // 1: tests.lint.OneofMessage.url:type_name -> tests.lint.Url
	10, // 2: tests.lint.OuterMessage.id:type_name -> tests.lint.OuterMessage.InnerId
	11, // 3: tests.lint.OuterMessage.url:type_name -> tests.lint.OuterMessage.InnerUrl
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_tests_lint_lint_proto_init() }
func file_tests_lint_lint_proto_init() {
	if File_tests_lint_lint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_lint_lint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*URL); i {
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
		file_tests_lint_lint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_tests_lint_lint_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RGBColor); i {
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
		file_tests_lint_lint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofMessage); i {
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
		file_tests_lint_lint_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessage); i {
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
		file_tests_lint_lint_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Color); i {
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
		file_tests_lint_lint_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDFieldsTest); i {
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
		file_tests_lint_lint_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessageInnerID); i {
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
		file_tests_lint_lint_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OuterMessageInnerURL); i {
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
	file_tests_lint_lint_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*OneofMessage_ID)(nil),
		(*OneofMessage_URL)(nil),
	}
	file_tests_lint_lint_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*Color_RGB)(nil),
		(*Color_RGBA)(nil),
		(*Color_HSV)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tests_lint_lint_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_lint_lint_proto_goTypes,
		DependencyIndexes: file_tests_lint_lint_proto_depIdxs,
		EnumInfos:         file_tests_lint_lint_proto_enumTypes,
		MessageInfos:      file_tests_lint_lint_proto_msgTypes,
	}.Build()
	File_tests_lint_lint_proto = out.File
	file_tests_lint_lint_proto_rawDesc = nil
	file_tests_lint_lint_proto_goTypes = nil
	file_tests_lint_lint_proto_depIdxs = nil
}
