// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/sequencer/operating_status.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// OperatingStatus defines the operating status of a sequencer
type OperatingStatus int32

const (
	// OPERATING_STATUS_UNBONDED defines a sequencer that is not active and won't be scheduled
	Unbonded OperatingStatus = 0
	// UNBONDING defines a sequencer that is currently unbonding.
	Unbonding OperatingStatus = 1
	// OPERATING_STATUS_BONDED defines a sequencer that is bonded and can be scheduled
	Bonded OperatingStatus = 2
)

var OperatingStatus_name = map[int32]string{
	0: "OPERATING_STATUS_UNBONDED",
	1: "OPERATING_STATUS_UNBONDING",
	2: "OPERATING_STATUS_BONDED",
}

var OperatingStatus_value = map[string]int32{
	"OPERATING_STATUS_UNBONDED":  0,
	"OPERATING_STATUS_UNBONDING": 1,
	"OPERATING_STATUS_BONDED":    2,
}

func (x OperatingStatus) String() string {
	return proto.EnumName(OperatingStatus_name, int32(x))
}

func (OperatingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_406c52b8df3dd90f, []int{0}
}

func init() {
	proto.RegisterEnum("dymensionxyz.dymension.sequencer.OperatingStatus", OperatingStatus_name, OperatingStatus_value)
}

func init() {
	proto.RegisterFile("dymension/sequencer/operating_status.proto", fileDescriptor_406c52b8df3dd90f)
}

var fileDescriptor_406c52b8df3dd90f = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0x4e, 0x2d, 0x2c, 0x4d, 0xcd, 0x4b, 0x4e, 0x2d, 0xd2,
	0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0x4b, 0x8f, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x80, 0xab, 0xad, 0xa8, 0xac, 0xd2, 0x83, 0x73,
	0xf4, 0xe0, 0x1a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x8a, 0xf5, 0x41, 0x2c, 0x88, 0x3e,
	0xad, 0x39, 0x8c, 0x5c, 0xfc, 0xfe, 0x30, 0x23, 0x83, 0xc1, 0x26, 0x0a, 0x69, 0x73, 0x49, 0xfa,
	0x07, 0xb8, 0x06, 0x39, 0x86, 0x78, 0xfa, 0xb9, 0xc7, 0x07, 0x87, 0x38, 0x86, 0x84, 0x06, 0xc7,
	0x87, 0xfa, 0x39, 0xf9, 0xfb, 0xb9, 0xb8, 0xba, 0x08, 0x30, 0x48, 0xf1, 0x74, 0xcd, 0x55, 0xe0,
	0x08, 0xcd, 0x4b, 0xca, 0xcf, 0x4b, 0x49, 0x4d, 0x11, 0xd2, 0xe5, 0x92, 0xc2, 0xa1, 0xd8, 0xd3,
	0xcf, 0x5d, 0x80, 0x51, 0x8a, 0xb7, 0x6b, 0xae, 0x02, 0x27, 0x44, 0x75, 0x66, 0x5e, 0xba, 0x90,
	0x3a, 0x97, 0x38, 0x86, 0x72, 0xa8, 0xc9, 0x4c, 0x52, 0x5c, 0x5d, 0x73, 0x15, 0xd8, 0x9c, 0xc0,
	0xe6, 0x4a, 0xb1, 0x74, 0x2c, 0x96, 0x63, 0x70, 0x0a, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23,
	0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6,
	0x63, 0x39, 0x86, 0x28, 0xb3, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d,
	0x64, 0xbf, 0x23, 0x38, 0xfa, 0x65, 0xc6, 0xfa, 0x15, 0x48, 0x21, 0x57, 0x52, 0x59, 0x90, 0x5a,
	0x9c, 0xc4, 0x06, 0xf6, 0xb7, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x7a, 0x7d, 0xf3, 0x5d,
	0x01, 0x00, 0x00,
}
