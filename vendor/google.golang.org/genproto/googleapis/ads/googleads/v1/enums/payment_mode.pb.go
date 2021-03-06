// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/payment_mode.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible payment modes.
type PaymentModeEnum_PaymentMode int32

const (
	// Not specified.
	PaymentModeEnum_UNSPECIFIED PaymentModeEnum_PaymentMode = 0
	// Used for return value only. Represents value unknown in this version.
	PaymentModeEnum_UNKNOWN PaymentModeEnum_PaymentMode = 1
	// Pay per click.
	PaymentModeEnum_CLICKS PaymentModeEnum_PaymentMode = 4
	// Pay per conversion value.
	PaymentModeEnum_CONVERSION_VALUE PaymentModeEnum_PaymentMode = 5
)

var PaymentModeEnum_PaymentMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	4: "CLICKS",
	5: "CONVERSION_VALUE",
}
var PaymentModeEnum_PaymentMode_value = map[string]int32{
	"UNSPECIFIED":      0,
	"UNKNOWN":          1,
	"CLICKS":           4,
	"CONVERSION_VALUE": 5,
}

func (x PaymentModeEnum_PaymentMode) String() string {
	return proto.EnumName(PaymentModeEnum_PaymentMode_name, int32(x))
}
func (PaymentModeEnum_PaymentMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_payment_mode_2af63ed3f55af206, []int{0, 0}
}

// Container for enum describing possible payment modes.
type PaymentModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaymentModeEnum) Reset()         { *m = PaymentModeEnum{} }
func (m *PaymentModeEnum) String() string { return proto.CompactTextString(m) }
func (*PaymentModeEnum) ProtoMessage()    {}
func (*PaymentModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_mode_2af63ed3f55af206, []int{0}
}
func (m *PaymentModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentModeEnum.Unmarshal(m, b)
}
func (m *PaymentModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentModeEnum.Marshal(b, m, deterministic)
}
func (dst *PaymentModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentModeEnum.Merge(dst, src)
}
func (m *PaymentModeEnum) XXX_Size() int {
	return xxx_messageInfo_PaymentModeEnum.Size(m)
}
func (m *PaymentModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PaymentModeEnum)(nil), "google.ads.googleads.v1.enums.PaymentModeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.PaymentModeEnum_PaymentMode", PaymentModeEnum_PaymentMode_name, PaymentModeEnum_PaymentMode_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/payment_mode.proto", fileDescriptor_payment_mode_2af63ed3f55af206)
}

var fileDescriptor_payment_mode_2af63ed3f55af206 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0xcd, 0x4a, 0x33, 0x31,
	0x14, 0xfd, 0xda, 0x4f, 0x2b, 0xa4, 0x8b, 0x86, 0xc1, 0x95, 0xd8, 0x45, 0xfb, 0x00, 0x89, 0x83,
	0xbb, 0xb8, 0x4a, 0xc7, 0xb1, 0x0c, 0x6d, 0xd3, 0xc1, 0xd2, 0x11, 0x64, 0xa0, 0x46, 0x13, 0x42,
	0xa1, 0x93, 0x0c, 0xcd, 0xb4, 0xe0, 0xeb, 0xb8, 0xf4, 0x51, 0x7c, 0x0f, 0x37, 0x3e, 0x85, 0x4c,
	0x62, 0x87, 0x6e, 0x74, 0x13, 0x0e, 0xf7, 0xfc, 0xe4, 0xdc, 0x0b, 0xae, 0x94, 0x31, 0x6a, 0x23,
	0x31, 0x17, 0x16, 0x7b, 0x58, 0xa3, 0x7d, 0x88, 0xa5, 0xde, 0x15, 0x16, 0x97, 0xfc, 0xb5, 0x90,
	0xba, 0x5a, 0x15, 0x46, 0x48, 0x54, 0x6e, 0x4d, 0x65, 0x82, 0xbe, 0x97, 0x21, 0x2e, 0x2c, 0x6a,
	0x1c, 0x68, 0x1f, 0x22, 0xe7, 0xb8, 0xb8, 0x3c, 0x04, 0x96, 0x6b, 0xcc, 0xb5, 0x36, 0x15, 0xaf,
	0xd6, 0x46, 0x5b, 0x6f, 0x1e, 0x3e, 0x81, 0x5e, 0xea, 0x23, 0x67, 0x46, 0xc8, 0x58, 0xef, 0x8a,
	0xe1, 0x0c, 0x74, 0x8f, 0x46, 0x41, 0x0f, 0x74, 0x97, 0x6c, 0x91, 0xc6, 0x51, 0x72, 0x97, 0xc4,
	0xb7, 0xf0, 0x5f, 0xd0, 0x05, 0x67, 0x4b, 0x36, 0x61, 0xf3, 0x07, 0x06, 0x5b, 0x01, 0x00, 0x9d,
	0x68, 0x9a, 0x44, 0x93, 0x05, 0x3c, 0x09, 0xce, 0x01, 0x8c, 0xe6, 0x2c, 0x8b, 0xef, 0x17, 0xc9,
	0x9c, 0xad, 0x32, 0x3a, 0x5d, 0xc6, 0xf0, 0x74, 0xf4, 0xd9, 0x02, 0x83, 0x17, 0x53, 0xa0, 0x3f,
	0x5b, 0x8e, 0xe0, 0xd1, 0x97, 0x69, 0xdd, 0x2c, 0x6d, 0x3d, 0x8e, 0x7e, 0x2c, 0xca, 0x6c, 0xb8,
	0x56, 0xc8, 0x6c, 0x15, 0x56, 0x52, 0xbb, 0xde, 0x87, 0xd3, 0x94, 0x6b, 0xfb, 0xcb, 0xa5, 0x6e,
	0xdc, 0xfb, 0xd6, 0xfe, 0x3f, 0xa6, 0xf4, 0xbd, 0xdd, 0x1f, 0xfb, 0x28, 0x2a, 0x2c, 0xf2, 0xb0,
	0x46, 0x59, 0x88, 0xea, 0x8d, 0xed, 0xc7, 0x81, 0xcf, 0xa9, 0xb0, 0x79, 0xc3, 0xe7, 0x59, 0x98,
	0x3b, 0xfe, 0xab, 0x3d, 0xf0, 0x43, 0x42, 0xa8, 0xb0, 0x84, 0x34, 0x0a, 0x42, 0xb2, 0x90, 0x10,
	0xa7, 0x79, 0xee, 0xb8, 0x62, 0xd7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x92, 0xa7, 0x9f,
	0xc1, 0x01, 0x00, 0x00,
}
