// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus", DO NOT EDIT.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pkg/apis/walrusutil/v1/generated.proto

package v1

import (
	fmt "fmt"

	io "io"

	proto "github.com/gogo/protobuf/proto"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

func (m *ScheduleTask) Reset()      { *m = ScheduleTask{} }
func (*ScheduleTask) ProtoMessage() {}
func (*ScheduleTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fde28f20e37081, []int{0}
}
func (m *ScheduleTask) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduleTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleTask.Merge(m, src)
}
func (m *ScheduleTask) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleTask) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleTask.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleTask proto.InternalMessageInfo

func (m *ScheduleTaskList) Reset()      { *m = ScheduleTaskList{} }
func (*ScheduleTaskList) ProtoMessage() {}
func (*ScheduleTaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fde28f20e37081, []int{1}
}
func (m *ScheduleTaskList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleTaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduleTaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleTaskList.Merge(m, src)
}
func (m *ScheduleTaskList) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleTaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleTaskList.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleTaskList proto.InternalMessageInfo

func (m *ScheduleTaskSpec) Reset()      { *m = ScheduleTaskSpec{} }
func (*ScheduleTaskSpec) ProtoMessage() {}
func (*ScheduleTaskSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fde28f20e37081, []int{2}
}
func (m *ScheduleTaskSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleTaskSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduleTaskSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleTaskSpec.Merge(m, src)
}
func (m *ScheduleTaskSpec) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleTaskSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleTaskSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleTaskSpec proto.InternalMessageInfo

func (m *ScheduleTaskStatus) Reset()      { *m = ScheduleTaskStatus{} }
func (*ScheduleTaskStatus) ProtoMessage() {}
func (*ScheduleTaskStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0fde28f20e37081, []int{3}
}
func (m *ScheduleTaskStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduleTaskStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ScheduleTaskStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduleTaskStatus.Merge(m, src)
}
func (m *ScheduleTaskStatus) XXX_Size() int {
	return m.Size()
}
func (m *ScheduleTaskStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduleTaskStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduleTaskStatus proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScheduleTask)(nil), "github.com.seal_io.walrus.pkg.apis.walrusutil.v1.ScheduleTask")
	proto.RegisterType((*ScheduleTaskList)(nil), "github.com.seal_io.walrus.pkg.apis.walrusutil.v1.ScheduleTaskList")
	proto.RegisterType((*ScheduleTaskSpec)(nil), "github.com.seal_io.walrus.pkg.apis.walrusutil.v1.ScheduleTaskSpec")
	proto.RegisterType((*ScheduleTaskStatus)(nil), "github.com.seal_io.walrus.pkg.apis.walrusutil.v1.ScheduleTaskStatus")
}

func init() {
	proto.RegisterFile("pkg/apis/walrusutil/v1/generated.proto", fileDescriptor_a0fde28f20e37081)
}

var fileDescriptor_a0fde28f20e37081 = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x4f, 0x6b, 0x13, 0x4d,
	0x1c, 0xce, 0xa6, 0x49, 0x9b, 0x4e, 0xfb, 0xe6, 0x4d, 0x87, 0xb7, 0x2f, 0x49, 0xc1, 0x4d, 0x89,
	0x20, 0x51, 0xec, 0x6c, 0x5b, 0x44, 0xc4, 0x83, 0x90, 0xa5, 0x28, 0x42, 0xaa, 0x30, 0x29, 0x1e,
	0xaa, 0xa0, 0xd3, 0xdd, 0x69, 0x32, 0x66, 0xff, 0x91, 0x99, 0x4d, 0x11, 0x2f, 0x7e, 0x02, 0xf1,
	0xe6, 0x17, 0xf2, 0xd0, 0x63, 0x2f, 0x42, 0x4f, 0x8b, 0x5d, 0xbf, 0x88, 0xcc, 0xec, 0x36, 0x9b,
	0x26, 0x29, 0x36, 0xf4, 0x94, 0xcc, 0x33, 0xbf, 0xe7, 0x79, 0x7e, 0xf3, 0xcc, 0x8f, 0x1d, 0x70,
	0x2f, 0xe8, 0x77, 0x0d, 0x12, 0x30, 0x6e, 0x9c, 0x10, 0x67, 0x10, 0xf2, 0x50, 0x30, 0xc7, 0x18,
	0xee, 0x18, 0x5d, 0xea, 0xd1, 0x01, 0x11, 0xd4, 0x46, 0xc1, 0xc0, 0x17, 0x3e, 0xdc, 0xee, 0x32,
	0xd1, 0x0b, 0x8f, 0x90, 0xe5, 0xbb, 0x88, 0x53, 0xe2, 0xbc, 0x67, 0x3e, 0x4a, 0x18, 0x28, 0xe8,
	0x77, 0x91, 0x54, 0x40, 0x99, 0x02, 0x1a, 0xee, 0x6c, 0x6c, 0x65, 0x0c, 0xa3, 0xeb, 0x77, 0x7d,
	0x43, 0x09, 0x1d, 0x85, 0xc7, 0x6a, 0xa5, 0x16, 0xea, 0x5f, 0x62, 0xb0, 0xf1, 0xa8, 0xff, 0x84,
	0x23, 0xe6, 0xcb, 0x5e, 0x5c, 0x62, 0xf5, 0x98, 0x47, 0x07, 0x9f, 0x8c, 0x51, 0x73, 0x2e, 0x15,
	0x64, 0x46, 0x5b, 0x1b, 0xc6, 0x75, 0xac, 0x41, 0xe8, 0x09, 0xe6, 0xd2, 0x29, 0xc2, 0xe3, 0xbf,
	0x11, 0xb8, 0xd5, 0xa3, 0x2e, 0x99, 0xe4, 0x35, 0x7e, 0xe4, 0xc1, 0x6a, 0xc7, 0xea, 0x51, 0x3b,
	0x74, 0xe8, 0x01, 0xe1, 0x7d, 0xf8, 0x01, 0x94, 0x64, 0x53, 0x36, 0x11, 0xa4, 0xaa, 0x6d, 0x6a,
	0xcd, 0x95, 0xdd, 0x6d, 0x94, 0x68, 0xa3, 0x71, 0xed, 0x2c, 0x1d, 0x59, 0x8d, 0x86, 0x3b, 0xe8,
	0xf5, 0xd1, 0x47, 0x6a, 0x89, 0x7d, 0x2a, 0x88, 0x09, 0x4f, 0xa3, 0x7a, 0x2e, 0x8e, 0xea, 0x20,
	0xc3, 0xf0, 0x48, 0x15, 0xda, 0xa0, 0xc0, 0x03, 0x6a, 0x55, 0xf3, 0x4a, 0xdd, 0x44, 0xf3, 0xde,
	0x00, 0x1a, 0xef, 0xb7, 0x13, 0x50, 0xcb, 0x5c, 0x4d, 0xfd, 0x0a, 0x72, 0x85, 0x95, 0x3a, 0x74,
	0xc0, 0x22, 0x17, 0x44, 0x84, 0xbc, 0xba, 0xa0, 0x7c, 0xf6, 0x6e, 0xe9, 0xa3, 0xb4, 0xcc, 0x72,
	0xea, 0xb4, 0x98, 0xac, 0x71, 0xea, 0xd1, 0xf8, 0xa9, 0x81, 0xca, 0x78, 0x79, 0x9b, 0x71, 0x01,
	0xdf, 0x4d, 0x45, 0x89, 0x6e, 0x16, 0xa5, 0x64, 0xab, 0x20, 0x2b, 0xa9, 0x5d, 0xe9, 0x12, 0x19,
	0x8b, 0xd1, 0x02, 0x45, 0x26, 0xa8, 0xcb, 0xab, 0xf9, 0xcd, 0x85, 0xe6, 0xca, 0xee, 0xb3, 0xdb,
	0x9d, 0xcf, 0xfc, 0x27, 0xb5, 0x2a, 0xbe, 0x94, 0xa2, 0x38, 0xd1, 0x6e, 0x7c, 0xcf, 0x5f, 0x3d,
	0x97, 0x0c, 0x18, 0x3e, 0x04, 0x25, 0x9e, 0x62, 0xea, 0x5c, 0xcb, 0x59, 0x9f, 0x97, 0xb5, 0x78,
	0x54, 0x01, 0xef, 0x83, 0x25, 0x1e, 0xf2, 0x80, 0x7a, 0xb6, 0xba, 0xf1, 0x92, 0xf9, 0x6f, 0x5a,
	0xbc, 0xd4, 0x49, 0x60, 0x7c, 0xb9, 0x0f, 0xef, 0x82, 0x22, 0x39, 0x21, 0x4c, 0xa8, 0x2b, 0x2b,
	0x65, 0x2d, 0xb5, 0x24, 0x88, 0x93, 0x3d, 0xd8, 0x04, 0x25, 0x39, 0xd0, 0x87, 0xbe, 0x47, 0xab,
	0x05, 0xe5, 0xbe, 0x2a, 0x9d, 0x0f, 0x52, 0x0c, 0x8f, 0x76, 0xe1, 0x5b, 0x50, 0x23, 0x96, 0x60,
	0x43, 0xd5, 0xf9, 0x1e, 0x25, 0xb6, 0xc3, 0x3c, 0xda, 0xa1, 0x96, 0xef, 0xd9, 0xbc, 0x5a, 0xdc,
	0xd4, 0x9a, 0x0b, 0xe6, 0x9d, 0x38, 0xaa, 0xd7, 0x5a, 0xd7, 0x15, 0xe1, 0xeb, 0xf9, 0x8d, 0xaf,
	0x45, 0x00, 0xa7, 0x07, 0x64, 0xce, 0x6c, 0x9e, 0x82, 0x72, 0xe6, 0xf0, 0x8a, 0xb8, 0x54, 0x45,
	0xb4, 0x6c, 0xc2, 0x38, 0xaa, 0x97, 0x5b, 0x57, 0x76, 0xf0, 0x44, 0x25, 0xfc, 0x0c, 0xd6, 0x33,
	0x44, 0x9e, 0xde, 0x0f, 0x85, 0xfc, 0x49, 0xe7, 0xfd, 0xc1, 0xcd, 0x46, 0x4d, 0x32, 0xcc, 0x5a,
	0x1c, 0xd5, 0xd7, 0x5b, 0xb3, 0xc4, 0xf0, 0x6c, 0x0f, 0xe8, 0x80, 0x8a, 0x43, 0xb8, 0x18, 0x05,
	0x20, 0x7d, 0x0b, 0x73, 0xfb, 0xfe, 0x17, 0x47, 0xf5, 0x4a, 0x7b, 0x42, 0x07, 0x4f, 0x29, 0xc3,
	0x01, 0x80, 0x0a, 0x0b, 0x2d, 0x8b, 0x72, 0x7e, 0x1c, 0x3a, 0xca, 0xaf, 0x38, 0xb7, 0xdf, 0xff,
	0x71, 0x54, 0x87, 0xed, 0x29, 0x25, 0x3c, 0x43, 0x1d, 0x1e, 0x83, 0xb2, 0x44, 0x9f, 0x13, 0xe6,
	0x50, 0x5b, 0xf9, 0x2d, 0xce, 0xed, 0xa7, 0xae, 0xb1, 0x7d, 0x45, 0x05, 0x4f, 0xa8, 0xc2, 0x17,
	0x60, 0x2d, 0x43, 0xf6, 0x29, 0xe7, 0xa4, 0x4b, 0xab, 0x4b, 0x6a, 0x0a, 0x6a, 0xe9, 0xe4, 0xac,
	0xb5, 0x27, 0x0b, 0xf0, 0x34, 0xc7, 0x7c, 0x73, 0x7a, 0xa1, 0xe7, 0xce, 0x2e, 0xf4, 0xdc, 0xf9,
	0x85, 0x9e, 0xfb, 0x12, 0xeb, 0xda, 0x69, 0xac, 0x6b, 0x67, 0xb1, 0xae, 0x9d, 0xc7, 0xba, 0xf6,
	0x2b, 0xd6, 0xb5, 0x6f, 0xbf, 0xf5, 0xdc, 0xe1, 0xd8, 0x7b, 0x67, 0xc8, 0xaf, 0xc4, 0x16, 0xf3,
	0xd3, 0x17, 0xd2, 0x98, 0xfd, 0x62, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x02, 0xd0, 0x6b, 0xca,
	0x4a, 0x07, 0x00, 0x00,
}

func (m *ScheduleTask) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleTask) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleTask) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Spec.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.ObjectMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ScheduleTaskList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleTaskList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleTaskList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Items) > 0 {
		for iNdEx := len(m.Items) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Items[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenerated(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.ListMeta.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenerated(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ScheduleTaskSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleTaskSpec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleTaskSpec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ActiveTaskDeadlineSeconds != nil {
		i = encodeVarintGenerated(dAtA, i, uint64(*m.ActiveTaskDeadlineSeconds))
		i--
		dAtA[i] = 0x28
	}
	if m.TimeZone != nil {
		i -= len(*m.TimeZone)
		copy(dAtA[i:], *m.TimeZone)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.TimeZone)))
		i--
		dAtA[i] = 0x22
	}
	i--
	if m.Await {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x18
	i--
	if m.Suspend {
		dAtA[i] = 1
	} else {
		dAtA[i] = 0
	}
	i--
	dAtA[i] = 0x10
	i -= len(m.Schedule)
	copy(dAtA[i:], m.Schedule)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Schedule)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ScheduleTaskStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduleTaskStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduleTaskStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	i -= len(m.LastFailedMessage)
	copy(dAtA[i:], m.LastFailedMessage)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.LastFailedMessage)))
	i--
	dAtA[i] = 0x3a
	if m.LastFailedTime != nil {
		{
			size, err := m.LastFailedTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.LastSuccessfulTime != nil {
		{
			size, err := m.LastSuccessfulTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.LastScheduleTime != nil {
		{
			size, err := m.LastScheduleTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.ActiveTaskTimeoutTime != nil {
		{
			size, err := m.ActiveTaskTimeoutTime.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.ActiveTaskName != nil {
		i -= len(*m.ActiveTaskName)
		copy(dAtA[i:], *m.ActiveTaskName)
		i = encodeVarintGenerated(dAtA, i, uint64(len(*m.ActiveTaskName)))
		i--
		dAtA[i] = 0x12
	}
	i -= len(m.Schedule)
	copy(dAtA[i:], m.Schedule)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Schedule)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduleTask) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ScheduleTaskList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *ScheduleTaskSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Schedule)
	n += 1 + l + sovGenerated(uint64(l))
	n += 2
	n += 2
	if m.TimeZone != nil {
		l = len(*m.TimeZone)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.ActiveTaskDeadlineSeconds != nil {
		n += 1 + sovGenerated(uint64(*m.ActiveTaskDeadlineSeconds))
	}
	return n
}

func (m *ScheduleTaskStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Schedule)
	n += 1 + l + sovGenerated(uint64(l))
	if m.ActiveTaskName != nil {
		l = len(*m.ActiveTaskName)
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.ActiveTaskTimeoutTime != nil {
		l = m.ActiveTaskTimeoutTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LastScheduleTime != nil {
		l = m.LastScheduleTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LastSuccessfulTime != nil {
		l = m.LastSuccessfulTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	if m.LastFailedTime != nil {
		l = m.LastFailedTime.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = len(m.LastFailedMessage)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ScheduleTask) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ScheduleTask{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ObjectMeta), "ObjectMeta", "v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "ScheduleTaskSpec", "ScheduleTaskSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "ScheduleTaskStatus", "ScheduleTaskStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ScheduleTaskList) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForItems := "[]ScheduleTask{"
	for _, f := range this.Items {
		repeatedStringForItems += strings.Replace(strings.Replace(f.String(), "ScheduleTask", "ScheduleTask", 1), `&`, ``, 1) + ","
	}
	repeatedStringForItems += "}"
	s := strings.Join([]string{`&ScheduleTaskList{`,
		`ListMeta:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.ListMeta), "ListMeta", "v1.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + repeatedStringForItems + `,`,
		`}`,
	}, "")
	return s
}
func (this *ScheduleTaskSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ScheduleTaskSpec{`,
		`Schedule:` + fmt.Sprintf("%v", this.Schedule) + `,`,
		`Suspend:` + fmt.Sprintf("%v", this.Suspend) + `,`,
		`Await:` + fmt.Sprintf("%v", this.Await) + `,`,
		`TimeZone:` + valueToStringGenerated(this.TimeZone) + `,`,
		`ActiveTaskDeadlineSeconds:` + valueToStringGenerated(this.ActiveTaskDeadlineSeconds) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ScheduleTaskStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ScheduleTaskStatus{`,
		`Schedule:` + fmt.Sprintf("%v", this.Schedule) + `,`,
		`ActiveTaskName:` + valueToStringGenerated(this.ActiveTaskName) + `,`,
		`ActiveTaskTimeoutTime:` + strings.Replace(fmt.Sprintf("%v", this.ActiveTaskTimeoutTime), "Time", "v1.Time", 1) + `,`,
		`LastScheduleTime:` + strings.Replace(fmt.Sprintf("%v", this.LastScheduleTime), "Time", "v1.Time", 1) + `,`,
		`LastSuccessfulTime:` + strings.Replace(fmt.Sprintf("%v", this.LastSuccessfulTime), "Time", "v1.Time", 1) + `,`,
		`LastFailedTime:` + strings.Replace(fmt.Sprintf("%v", this.LastFailedTime), "Time", "v1.Time", 1) + `,`,
		`LastFailedMessage:` + fmt.Sprintf("%v", this.LastFailedMessage) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ScheduleTask) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleTask: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleTask: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduleTaskList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleTaskList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleTaskList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, ScheduleTask{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduleTaskSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleTaskSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleTaskSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schedule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schedule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Suspend", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Suspend = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Await", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Await = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeZone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.TimeZone = &s
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveTaskDeadlineSeconds", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ActiveTaskDeadlineSeconds = &v
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ScheduleTaskStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ScheduleTaskStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduleTaskStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schedule", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schedule = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveTaskName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.ActiveTaskName = &s
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveTaskTimeoutTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ActiveTaskTimeoutTime == nil {
				m.ActiveTaskTimeoutTime = &v1.Time{}
			}
			if err := m.ActiveTaskTimeoutTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastScheduleTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastScheduleTime == nil {
				m.LastScheduleTime = &v1.Time{}
			}
			if err := m.LastScheduleTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastSuccessfulTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastSuccessfulTime == nil {
				m.LastSuccessfulTime = &v1.Time{}
			}
			if err := m.LastSuccessfulTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastFailedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastFailedTime == nil {
				m.LastFailedTime = &v1.Time{}
			}
			if err := m.LastFailedTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastFailedMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastFailedMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenerated
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenerated
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenerated        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenerated = fmt.Errorf("proto: unexpected end of group")
)