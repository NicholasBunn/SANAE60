// Code generated by protoc-gen-go. DO NOT EDIT.
// source: powerEstimationSP/proto/powerEstimationAPI.proto

package powerEstimation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ModelTypeEnum int32

const (
	ModelTypeEnum_UNKNOWN   ModelTypeEnum = 0
	ModelTypeEnum_OPENWATER ModelTypeEnum = 1
	ModelTypeEnum_ICE       ModelTypeEnum = 2
)

var ModelTypeEnum_name = map[int32]string{
	0: "UNKNOWN",
	1: "OPENWATER",
	2: "ICE",
}

var ModelTypeEnum_value = map[string]int32{
	"UNKNOWN":   0,
	"OPENWATER": 1,
	"ICE":       2,
}

func (x ModelTypeEnum) String() string {
	return proto.EnumName(ModelTypeEnum_name, int32(x))
}

func (ModelTypeEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_daf15462637bc3b8, []int{0}
}

type ServicePackageRequestMessage struct {
	InputFile            string        `protobuf:"bytes,1,opt,name=input_file,json=inputFile,proto3" json:"input_file,omitempty"`
	ModelType            ModelTypeEnum `protobuf:"varint,2,opt,name=model_type,json=modelType,proto3,enum=ModelTypeEnum" json:"model_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ServicePackageRequestMessage) Reset()         { *m = ServicePackageRequestMessage{} }
func (m *ServicePackageRequestMessage) String() string { return proto.CompactTextString(m) }
func (*ServicePackageRequestMessage) ProtoMessage()    {}
func (*ServicePackageRequestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_daf15462637bc3b8, []int{0}
}

func (m *ServicePackageRequestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServicePackageRequestMessage.Unmarshal(m, b)
}
func (m *ServicePackageRequestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServicePackageRequestMessage.Marshal(b, m, deterministic)
}
func (m *ServicePackageRequestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServicePackageRequestMessage.Merge(m, src)
}
func (m *ServicePackageRequestMessage) XXX_Size() int {
	return xxx_messageInfo_ServicePackageRequestMessage.Size(m)
}
func (m *ServicePackageRequestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ServicePackageRequestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ServicePackageRequestMessage proto.InternalMessageInfo

func (m *ServicePackageRequestMessage) GetInputFile() string {
	if m != nil {
		return m.InputFile
	}
	return ""
}

func (m *ServicePackageRequestMessage) GetModelType() ModelTypeEnum {
	if m != nil {
		return m.ModelType
	}
	return ModelTypeEnum_UNKNOWN
}

type EstimateResponseMessage struct {
	PowerEstimate        []float32 `protobuf:"fixed32,1,rep,packed,name=power_estimate,json=powerEstimate,proto3" json:"power_estimate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EstimateResponseMessage) Reset()         { *m = EstimateResponseMessage{} }
func (m *EstimateResponseMessage) String() string { return proto.CompactTextString(m) }
func (*EstimateResponseMessage) ProtoMessage()    {}
func (*EstimateResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_daf15462637bc3b8, []int{1}
}

func (m *EstimateResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EstimateResponseMessage.Unmarshal(m, b)
}
func (m *EstimateResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EstimateResponseMessage.Marshal(b, m, deterministic)
}
func (m *EstimateResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EstimateResponseMessage.Merge(m, src)
}
func (m *EstimateResponseMessage) XXX_Size() int {
	return xxx_messageInfo_EstimateResponseMessage.Size(m)
}
func (m *EstimateResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EstimateResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EstimateResponseMessage proto.InternalMessageInfo

func (m *EstimateResponseMessage) GetPowerEstimate() []float32 {
	if m != nil {
		return m.PowerEstimate
	}
	return nil
}

type EvaluateResponseMessage struct {
	PowerEstimate        []float32 `protobuf:"fixed32,1,rep,packed,name=power_estimate,json=powerEstimate,proto3" json:"power_estimate,omitempty"`
	PowerActual          []float32 `protobuf:"fixed32,2,rep,packed,name=power_actual,json=powerActual,proto3" json:"power_actual,omitempty"`
	SpeedOverGround      []float32 `protobuf:"fixed32,3,rep,packed,name=speed_over_ground,json=speedOverGround,proto3" json:"speed_over_ground,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *EvaluateResponseMessage) Reset()         { *m = EvaluateResponseMessage{} }
func (m *EvaluateResponseMessage) String() string { return proto.CompactTextString(m) }
func (*EvaluateResponseMessage) ProtoMessage()    {}
func (*EvaluateResponseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_daf15462637bc3b8, []int{2}
}

func (m *EvaluateResponseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvaluateResponseMessage.Unmarshal(m, b)
}
func (m *EvaluateResponseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvaluateResponseMessage.Marshal(b, m, deterministic)
}
func (m *EvaluateResponseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvaluateResponseMessage.Merge(m, src)
}
func (m *EvaluateResponseMessage) XXX_Size() int {
	return xxx_messageInfo_EvaluateResponseMessage.Size(m)
}
func (m *EvaluateResponseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_EvaluateResponseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_EvaluateResponseMessage proto.InternalMessageInfo

func (m *EvaluateResponseMessage) GetPowerEstimate() []float32 {
	if m != nil {
		return m.PowerEstimate
	}
	return nil
}

func (m *EvaluateResponseMessage) GetPowerActual() []float32 {
	if m != nil {
		return m.PowerActual
	}
	return nil
}

func (m *EvaluateResponseMessage) GetSpeedOverGround() []float32 {
	if m != nil {
		return m.SpeedOverGround
	}
	return nil
}

func init() {
	proto.RegisterEnum("ModelTypeEnum", ModelTypeEnum_name, ModelTypeEnum_value)
	proto.RegisterType((*ServicePackageRequestMessage)(nil), "ServicePackageRequestMessage")
	proto.RegisterType((*EstimateResponseMessage)(nil), "EstimateResponseMessage")
	proto.RegisterType((*EvaluateResponseMessage)(nil), "EvaluateResponseMessage")
}

func init() {
	proto.RegisterFile("powerEstimationSP/proto/powerEstimationAPI.proto", fileDescriptor_daf15462637bc3b8)
}

var fileDescriptor_daf15462637bc3b8 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x4b, 0x2a, 0x41,
	0x14, 0xc6, 0xef, 0xae, 0x70, 0x65, 0x8f, 0x57, 0xaf, 0x77, 0xe0, 0xe2, 0x12, 0x09, 0x26, 0x44,
	0x26, 0xb4, 0x86, 0xf5, 0xd6, 0x4b, 0x16, 0x5b, 0x48, 0xb8, 0x2e, 0xab, 0x21, 0xf4, 0xb2, 0x6c,
	0x7a, 0x92, 0xa5, 0xdd, 0x9d, 0x69, 0x66, 0x76, 0xc3, 0xbf, 0xa2, 0x3f, 0xaa, 0x7f, 0x2c, 0x1a,
	0x35, 0x54, 0x32, 0xa8, 0xc7, 0xf9, 0x9d, 0x8f, 0x33, 0xdf, 0x7c, 0xdf, 0xc0, 0x31, 0xa3, 0xcf,
	0xc8, 0x6d, 0x21, 0xc3, 0x38, 0x90, 0x21, 0x4d, 0x06, 0x6e, 0x8b, 0x71, 0x2a, 0x69, 0x6b, 0x83,
	0x77, 0xdc, 0xae, 0xa5, 0x06, 0xf5, 0x08, 0x76, 0x07, 0xc8, 0xb3, 0x70, 0x8c, 0x6e, 0x30, 0x7e,
	0x0c, 0xa6, 0xe8, 0xe1, 0x53, 0x8a, 0x42, 0xf6, 0x50, 0x88, 0x60, 0x8a, 0xa4, 0x0a, 0x10, 0x26,
	0x2c, 0x95, 0xfe, 0x43, 0x18, 0xa1, 0xa9, 0xd5, 0xb4, 0x86, 0xe1, 0x19, 0x8a, 0x5c, 0x85, 0x11,
	0x92, 0x23, 0x80, 0x98, 0x4e, 0x30, 0xf2, 0xe5, 0x8c, 0xa1, 0xa9, 0xd7, 0xb4, 0x46, 0xa9, 0x5d,
	0xb2, 0x7a, 0xef, 0x68, 0x38, 0x63, 0x68, 0x27, 0x69, 0xec, 0x19, 0xf1, 0xf2, 0x58, 0x3f, 0x87,
	0xca, 0xc2, 0x04, 0x7a, 0x28, 0x18, 0x4d, 0x04, 0x2e, 0x2f, 0xda, 0x87, 0x92, 0x32, 0xe9, 0xe3,
	0x42, 0x60, 0x6a, 0xb5, 0x5c, 0x43, 0xf7, 0x8a, 0xab, 0xd6, 0xb1, 0xfe, 0xa2, 0x41, 0xc5, 0xce,
	0x82, 0x28, 0xfd, 0xf1, 0x0a, 0xb2, 0x07, 0x7f, 0xe6, 0xb2, 0x60, 0x2c, 0xd3, 0x20, 0x32, 0x75,
	0x25, 0x2a, 0x28, 0xd6, 0x51, 0x88, 0x34, 0xe1, 0x9f, 0x60, 0x88, 0x13, 0x9f, 0x66, 0xc8, 0xfd,
	0x29, 0xa7, 0x69, 0x32, 0x31, 0x73, 0x4a, 0xf7, 0x57, 0x0d, 0xfa, 0x19, 0xf2, 0x6b, 0x85, 0x9b,
	0xa7, 0x50, 0x5c, 0x7b, 0x2f, 0x29, 0x40, 0xfe, 0xd6, 0xb9, 0x71, 0xfa, 0x23, 0xa7, 0xfc, 0x8b,
	0x14, 0xc1, 0xe8, 0xbb, 0xb6, 0x33, 0xea, 0x0c, 0x6d, 0xaf, 0xac, 0x91, 0x3c, 0xe4, 0xba, 0x97,
	0x76, 0x59, 0x6f, 0xbf, 0x6a, 0x50, 0x75, 0x37, 0xca, 0x5a, 0xeb, 0x81, 0xb8, 0xf0, 0x7f, 0x55,
	0x40, 0xf9, 0x62, 0x4e, 0xaa, 0xd6, 0x57, 0x8d, 0xed, 0x98, 0xd6, 0xb6, 0x88, 0x3f, 0x36, 0xce,
	0xf3, 0xfb, 0xce, 0xc6, 0xcf, 0x13, 0xbf, 0x38, 0xbc, 0x3b, 0xd8, 0xf2, 0xe3, 0xce, 0x36, 0xf8,
	0xfd, 0x6f, 0x85, 0x4f, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0x96, 0x92, 0xea, 0x24, 0xa3, 0x02,
	0x00, 0x00,
}
