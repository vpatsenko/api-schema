// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config_sync.proto

package dialog

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/types"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Syncing Parameter
type Parameter struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{0}
}

func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (m *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(m, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Parameter) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Getting Parameters
type RequestGetParameters struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestGetParameters) Reset()         { *m = RequestGetParameters{} }
func (m *RequestGetParameters) String() string { return proto.CompactTextString(m) }
func (*RequestGetParameters) ProtoMessage()    {}
func (*RequestGetParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{1}
}

func (m *RequestGetParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestGetParameters.Unmarshal(m, b)
}
func (m *RequestGetParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestGetParameters.Marshal(b, m, deterministic)
}
func (m *RequestGetParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestGetParameters.Merge(m, src)
}
func (m *RequestGetParameters) XXX_Size() int {
	return xxx_messageInfo_RequestGetParameters.Size(m)
}
func (m *RequestGetParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestGetParameters.DiscardUnknown(m)
}

var xxx_messageInfo_RequestGetParameters proto.InternalMessageInfo

type ResponseGetParameters struct {
	Parameters           []*Parameter `protobuf:"bytes,1,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ResponseGetParameters) Reset()         { *m = ResponseGetParameters{} }
func (m *ResponseGetParameters) String() string { return proto.CompactTextString(m) }
func (*ResponseGetParameters) ProtoMessage()    {}
func (*ResponseGetParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{2}
}

func (m *ResponseGetParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseGetParameters.Unmarshal(m, b)
}
func (m *ResponseGetParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseGetParameters.Marshal(b, m, deterministic)
}
func (m *ResponseGetParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseGetParameters.Merge(m, src)
}
func (m *ResponseGetParameters) XXX_Size() int {
	return xxx_messageInfo_ResponseGetParameters.Size(m)
}
func (m *ResponseGetParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseGetParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseGetParameters proto.InternalMessageInfo

func (m *ResponseGetParameters) GetParameters() []*Parameter {
	if m != nil {
		return m.Parameters
	}
	return nil
}

// Change parameter value
type RequestEditParameter struct {
	Key                  string                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *wrappers.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RequestEditParameter) Reset()         { *m = RequestEditParameter{} }
func (m *RequestEditParameter) String() string { return proto.CompactTextString(m) }
func (*RequestEditParameter) ProtoMessage()    {}
func (*RequestEditParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{3}
}

func (m *RequestEditParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestEditParameter.Unmarshal(m, b)
}
func (m *RequestEditParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestEditParameter.Marshal(b, m, deterministic)
}
func (m *RequestEditParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestEditParameter.Merge(m, src)
}
func (m *RequestEditParameter) XXX_Size() int {
	return xxx_messageInfo_RequestEditParameter.Size(m)
}
func (m *RequestEditParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestEditParameter.DiscardUnknown(m)
}

var xxx_messageInfo_RequestEditParameter proto.InternalMessageInfo

func (m *RequestEditParameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestEditParameter) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

type FeatureFlag struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Clock                int64    `protobuf:"varint,3,opt,name=clock,proto3" json:"clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureFlag) Reset()         { *m = FeatureFlag{} }
func (m *FeatureFlag) String() string { return proto.CompactTextString(m) }
func (*FeatureFlag) ProtoMessage()    {}
func (*FeatureFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{4}
}

func (m *FeatureFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureFlag.Unmarshal(m, b)
}
func (m *FeatureFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureFlag.Marshal(b, m, deterministic)
}
func (m *FeatureFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureFlag.Merge(m, src)
}
func (m *FeatureFlag) XXX_Size() int {
	return xxx_messageInfo_FeatureFlag.Size(m)
}
func (m *FeatureFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureFlag.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureFlag proto.InternalMessageInfo

func (m *FeatureFlag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *FeatureFlag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FeatureFlag) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

type RequestFeatureFlags struct {
	Clock                int64    `protobuf:"varint,1,opt,name=clock,proto3" json:"clock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestFeatureFlags) Reset()         { *m = RequestFeatureFlags{} }
func (m *RequestFeatureFlags) String() string { return proto.CompactTextString(m) }
func (*RequestFeatureFlags) ProtoMessage()    {}
func (*RequestFeatureFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{5}
}

func (m *RequestFeatureFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestFeatureFlags.Unmarshal(m, b)
}
func (m *RequestFeatureFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestFeatureFlags.Marshal(b, m, deterministic)
}
func (m *RequestFeatureFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestFeatureFlags.Merge(m, src)
}
func (m *RequestFeatureFlags) XXX_Size() int {
	return xxx_messageInfo_RequestFeatureFlags.Size(m)
}
func (m *RequestFeatureFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestFeatureFlags.DiscardUnknown(m)
}

var xxx_messageInfo_RequestFeatureFlags proto.InternalMessageInfo

func (m *RequestFeatureFlags) GetClock() int64 {
	if m != nil {
		return m.Clock
	}
	return 0
}

type ResponseFeatureFlags struct {
	FeatureConfig        []*FeatureFlag `protobuf:"bytes,1,rep,name=feature_config,json=featureConfig,proto3" json:"feature_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ResponseFeatureFlags) Reset()         { *m = ResponseFeatureFlags{} }
func (m *ResponseFeatureFlags) String() string { return proto.CompactTextString(m) }
func (*ResponseFeatureFlags) ProtoMessage()    {}
func (*ResponseFeatureFlags) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{6}
}

func (m *ResponseFeatureFlags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseFeatureFlags.Unmarshal(m, b)
}
func (m *ResponseFeatureFlags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseFeatureFlags.Marshal(b, m, deterministic)
}
func (m *ResponseFeatureFlags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseFeatureFlags.Merge(m, src)
}
func (m *ResponseFeatureFlags) XXX_Size() int {
	return xxx_messageInfo_ResponseFeatureFlags.Size(m)
}
func (m *ResponseFeatureFlags) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseFeatureFlags.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseFeatureFlags proto.InternalMessageInfo

func (m *ResponseFeatureFlags) GetFeatureConfig() []*FeatureFlag {
	if m != nil {
		return m.FeatureConfig
	}
	return nil
}

type UpdateFeatureFlagChanged struct {
	Feature              *FeatureFlag `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateFeatureFlagChanged) Reset()         { *m = UpdateFeatureFlagChanged{} }
func (m *UpdateFeatureFlagChanged) String() string { return proto.CompactTextString(m) }
func (*UpdateFeatureFlagChanged) ProtoMessage()    {}
func (*UpdateFeatureFlagChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{7}
}

func (m *UpdateFeatureFlagChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeatureFlagChanged.Unmarshal(m, b)
}
func (m *UpdateFeatureFlagChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeatureFlagChanged.Marshal(b, m, deterministic)
}
func (m *UpdateFeatureFlagChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeatureFlagChanged.Merge(m, src)
}
func (m *UpdateFeatureFlagChanged) XXX_Size() int {
	return xxx_messageInfo_UpdateFeatureFlagChanged.Size(m)
}
func (m *UpdateFeatureFlagChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeatureFlagChanged.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeatureFlagChanged proto.InternalMessageInfo

func (m *UpdateFeatureFlagChanged) GetFeature() *FeatureFlag {
	if m != nil {
		return m.Feature
	}
	return nil
}

// Update about parameter change
type UpdateParameterChanged struct {
	Key                  string                `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *wrappers.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateParameterChanged) Reset()         { *m = UpdateParameterChanged{} }
func (m *UpdateParameterChanged) String() string { return proto.CompactTextString(m) }
func (*UpdateParameterChanged) ProtoMessage()    {}
func (*UpdateParameterChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_760cd360e0cd3e2d, []int{8}
}

func (m *UpdateParameterChanged) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateParameterChanged.Unmarshal(m, b)
}
func (m *UpdateParameterChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateParameterChanged.Marshal(b, m, deterministic)
}
func (m *UpdateParameterChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateParameterChanged.Merge(m, src)
}
func (m *UpdateParameterChanged) XXX_Size() int {
	return xxx_messageInfo_UpdateParameterChanged.Size(m)
}
func (m *UpdateParameterChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateParameterChanged.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateParameterChanged proto.InternalMessageInfo

func (m *UpdateParameterChanged) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdateParameterChanged) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameter)(nil), "dialog.Parameter")
	proto.RegisterType((*RequestGetParameters)(nil), "dialog.RequestGetParameters")
	proto.RegisterType((*ResponseGetParameters)(nil), "dialog.ResponseGetParameters")
	proto.RegisterType((*RequestEditParameter)(nil), "dialog.RequestEditParameter")
	proto.RegisterType((*FeatureFlag)(nil), "dialog.FeatureFlag")
	proto.RegisterType((*RequestFeatureFlags)(nil), "dialog.RequestFeatureFlags")
	proto.RegisterType((*ResponseFeatureFlags)(nil), "dialog.ResponseFeatureFlags")
	proto.RegisterType((*UpdateFeatureFlagChanged)(nil), "dialog.UpdateFeatureFlagChanged")
	proto.RegisterType((*UpdateParameterChanged)(nil), "dialog.UpdateParameterChanged")
}

func init() { proto.RegisterFile("config_sync.proto", fileDescriptor_760cd360e0cd3e2d) }

var fileDescriptor_760cd360e0cd3e2d = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x1b, 0x91, 0x92, 0x49, 0x83, 0x14, 0x27, 0x2d, 0xae, 0x49, 0x21, 0x18, 0x81, 0xa2,
	0x82, 0xec, 0x36, 0x70, 0x8a, 0x90, 0x82, 0x5a, 0xd1, 0x8a, 0x5b, 0xe5, 0x08, 0xc4, 0xad, 0xda,
	0xd8, 0x13, 0x77, 0x55, 0x67, 0xed, 0x78, 0x9d, 0x40, 0xe0, 0xc6, 0xb1, 0x57, 0x0e, 0x7c, 0x18,
	0x7f, 0x10, 0x71, 0xea, 0x91, 0x2f, 0x40, 0xf6, 0xda, 0xa9, 0xed, 0x46, 0x44, 0x1c, 0x38, 0x59,
	0xde, 0xf7, 0x66, 0xde, 0xcc, 0x9b, 0x19, 0xa8, 0x5b, 0x1e, 0x1b, 0x51, 0xe7, 0x9c, 0xcf, 0x99,
	0xa5, 0xfb, 0x81, 0x17, 0x7a, 0x72, 0xd9, 0xa6, 0xc4, 0xf5, 0x1c, 0xf5, 0xa1, 0xe3, 0x79, 0x8e,
	0x8b, 0x46, 0xfc, 0x3a, 0x9c, 0x8e, 0x8c, 0x4f, 0x01, 0xf1, 0x7d, 0x0c, 0xb8, 0xe0, 0xa9, 0xad,
	0x04, 0x27, 0x3e, 0x35, 0x08, 0x63, 0x5e, 0x48, 0x42, 0xea, 0xb1, 0x14, 0xad, 0xdb, 0x38, 0xa2,
	0x8c, 0x66, 0x9f, 0x1a, 0x63, 0xca, 0x2d, 0x74, 0x5d, 0xc2, 0xd0, 0x9b, 0xa6, 0x8f, 0xdb, 0xdc,
	0x22, 0x2e, 0xf1, 0x87, 0x46, 0xf2, 0x15, 0xcf, 0xda, 0x19, 0x54, 0xce, 0x48, 0x40, 0xc6, 0x18,
	0x62, 0x20, 0x3f, 0x82, 0xd2, 0x25, 0xce, 0x15, 0xa9, 0x2d, 0x75, 0x2a, 0x47, 0xb5, 0xab, 0xeb,
	0x83, 0x0a, 0x6c, 0xce, 0x28, 0xa7, 0x43, 0x17, 0xcd, 0x08, 0x91, 0x35, 0xb8, 0x33, 0x23, 0xee,
	0x14, 0x95, 0x8d, 0x98, 0xb2, 0x75, 0x75, 0x7d, 0x70, 0x17, 0xca, 0x17, 0xd4, 0xb6, 0x91, 0x99,
	0x02, 0xd2, 0x5e, 0x41, 0xd3, 0xc4, 0xc9, 0x14, 0x79, 0x78, 0x8a, 0xe1, 0x32, 0x37, 0xef, 0xb5,
	0x16, 0xfd, 0x5d, 0xb8, 0x4f, 0xc7, 0xba, 0xed, 0x3a, 0xba, 0x13, 0xf8, 0x96, 0x7e, 0x1a, 0xf8,
	0x56, 0x42, 0xd5, 0x28, 0x6c, 0x9b, 0xc8, 0x7d, 0x8f, 0x71, 0xcc, 0x85, 0xc9, 0x87, 0x00, 0xfe,
	0xf2, 0x4f, 0x91, 0xda, 0xa5, 0x4e, 0xb5, 0x5b, 0xd7, 0x85, 0x75, 0xfa, 0x92, 0x67, 0x66, 0x48,
	0xbd, 0xbd, 0x45, 0x5f, 0x05, 0xe5, 0xb6, 0x92, 0x48, 0xaf, 0xfd, 0x90, 0x96, 0x15, 0xbe, 0xb5,
	0x69, 0xf8, 0x0f, 0xed, 0xbf, 0xc9, 0xb6, 0x5f, 0xed, 0xb6, 0x74, 0x31, 0x19, 0x3d, 0x9d, 0x9c,
	0x3e, 0x08, 0x03, 0xca, 0x9c, 0x0f, 0x11, 0x67, 0xa5, 0x39, 0x6b, 0x4c, 0xf8, 0x0c, 0xd5, 0x13,
	0x24, 0xe1, 0x34, 0xc0, 0x13, 0x97, 0x38, 0xeb, 0xeb, 0x79, 0x92, 0x1f, 0x47, 0x81, 0x22, 0xb0,
	0x88, 0x64, 0xb9, 0x9e, 0x75, 0xa9, 0x94, 0xda, 0x52, 0xa7, 0x74, 0x8b, 0x14, 0x63, 0xda, 0x3b,
	0x68, 0x24, 0x45, 0x64, 0x0a, 0xe0, 0x72, 0x33, 0x8d, 0x8d, 0x6a, 0x28, 0x25, 0xe4, 0x35, 0x4d,
	0x4c, 0x22, 0x77, 0x85, 0xd5, 0xb9, 0x5c, 0x3d, 0xb8, 0x37, 0x12, 0xff, 0xe7, 0xe2, 0x16, 0x92,
	0x61, 0x36, 0xd2, 0x61, 0x66, 0xd8, 0x66, 0x2d, 0xa1, 0x1e, 0xc7, 0xcc, 0x75, 0x13, 0xfd, 0x08,
	0xca, 0x7b, 0xdf, 0x26, 0x61, 0x56, 0xf0, 0xf8, 0x82, 0x30, 0x07, 0x6d, 0xf9, 0x35, 0x6c, 0x26,
	0xb9, 0xe2, 0x26, 0x56, 0xeb, 0x15, 0x5d, 0x49, 0x43, 0xb4, 0xaf, 0xb0, 0x23, 0x32, 0x2f, 0xb7,
	0x24, 0xcd, 0xfb, 0xff, 0x97, 0xa5, 0xfb, 0x7b, 0x03, 0x40, 0x18, 0x30, 0x98, 0x33, 0x4b, 0xfe,
	0x02, 0xb5, 0xfc, 0x69, 0xb4, 0xd2, 0x4e, 0x56, 0xdd, 0x9b, 0xba, 0x77, 0x83, 0xae, 0xb8, 0x2b,
	0xed, 0xc5, 0xb7, 0x9f, 0xbf, 0xbe, 0x6f, 0x3c, 0xd3, 0x1e, 0x1b, 0xb3, 0x43, 0x23, 0xb2, 0xd4,
	0xb8, 0x91, 0x32, 0xf2, 0x97, 0x2b, 0xed, 0xcb, 0x3e, 0xd4, 0xf2, 0xb7, 0x52, 0xd4, 0xce, 0xa1,
	0x6a, 0xa3, 0xa8, 0x3d, 0xc0, 0xc9, 0xdf, 0x15, 0x73, 0xf1, 0x91, 0xe2, 0x0c, 0xb6, 0x72, 0xeb,
	0xf3, 0xa0, 0x20, 0x98, 0x05, 0xd5, 0x56, 0x51, 0x2f, 0x8b, 0x6a, 0xcf, 0x63, 0xe1, 0xa7, 0x5a,
	0x7b, 0x95, 0x70, 0x96, 0xd9, 0x93, 0xf6, 0x8f, 0x76, 0x17, 0xfd, 0x1d, 0x68, 0x66, 0x57, 0x8d,
	0x63, 0x30, 0xa3, 0x16, 0xf2, 0x61, 0x39, 0x1e, 0xdd, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0xc0, 0x1d, 0x9c, 0xcc, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigSyncClient is the client API for ConfigSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigSyncClient interface {
	GetParameters(ctx context.Context, in *RequestGetParameters, opts ...grpc.CallOption) (*ResponseGetParameters, error)
	EditParameter(ctx context.Context, in *RequestEditParameter, opts ...grpc.CallOption) (*ResponseSeq, error)
	FeatureFlags(ctx context.Context, in *RequestFeatureFlags, opts ...grpc.CallOption) (*ResponseFeatureFlags, error)
}

type configSyncClient struct {
	cc *grpc.ClientConn
}

func NewConfigSyncClient(cc *grpc.ClientConn) ConfigSyncClient {
	return &configSyncClient{cc}
}

func (c *configSyncClient) GetParameters(ctx context.Context, in *RequestGetParameters, opts ...grpc.CallOption) (*ResponseGetParameters, error) {
	out := new(ResponseGetParameters)
	err := c.cc.Invoke(ctx, "/dialog.ConfigSync/GetParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configSyncClient) EditParameter(ctx context.Context, in *RequestEditParameter, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := c.cc.Invoke(ctx, "/dialog.ConfigSync/EditParameter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configSyncClient) FeatureFlags(ctx context.Context, in *RequestFeatureFlags, opts ...grpc.CallOption) (*ResponseFeatureFlags, error) {
	out := new(ResponseFeatureFlags)
	err := c.cc.Invoke(ctx, "/dialog.ConfigSync/FeatureFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigSyncServer is the server API for ConfigSync service.
type ConfigSyncServer interface {
	GetParameters(context.Context, *RequestGetParameters) (*ResponseGetParameters, error)
	EditParameter(context.Context, *RequestEditParameter) (*ResponseSeq, error)
	FeatureFlags(context.Context, *RequestFeatureFlags) (*ResponseFeatureFlags, error)
}

// UnimplementedConfigSyncServer can be embedded to have forward compatible implementations.
type UnimplementedConfigSyncServer struct {
}

func (*UnimplementedConfigSyncServer) GetParameters(ctx context.Context, req *RequestGetParameters) (*ResponseGetParameters, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetParameters not implemented")
}
func (*UnimplementedConfigSyncServer) EditParameter(ctx context.Context, req *RequestEditParameter) (*ResponseSeq, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditParameter not implemented")
}
func (*UnimplementedConfigSyncServer) FeatureFlags(ctx context.Context, req *RequestFeatureFlags) (*ResponseFeatureFlags, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FeatureFlags not implemented")
}

func RegisterConfigSyncServer(s *grpc.Server, srv ConfigSyncServer) {
	s.RegisterService(&_ConfigSync_serviceDesc, srv)
}

func _ConfigSync_GetParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigSyncServer).GetParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.ConfigSync/GetParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigSyncServer).GetParameters(ctx, req.(*RequestGetParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigSync_EditParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestEditParameter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigSyncServer).EditParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.ConfigSync/EditParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigSyncServer).EditParameter(ctx, req.(*RequestEditParameter))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigSync_FeatureFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestFeatureFlags)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigSyncServer).FeatureFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.ConfigSync/FeatureFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigSyncServer).FeatureFlags(ctx, req.(*RequestFeatureFlags))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.ConfigSync",
	HandlerType: (*ConfigSyncServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetParameters",
			Handler:    _ConfigSync_GetParameters_Handler,
		},
		{
			MethodName: "EditParameter",
			Handler:    _ConfigSync_EditParameter_Handler,
		},
		{
			MethodName: "FeatureFlags",
			Handler:    _ConfigSync_FeatureFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config_sync.proto",
}
