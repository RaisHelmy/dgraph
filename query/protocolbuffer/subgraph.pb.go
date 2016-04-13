// Code generated by protoc-gen-go.
// source: subgraph.proto
// DO NOT EDIT!

/*
Package protocolbuffer is a generated protocol buffer package.

It is generated from these files:
	subgraph.proto

It has these top-level messages:
	UidList
	Result
	SubGraph
*/
package protocolbuffer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type UidList struct {
	Uids []uint64 `protobuf:"fixed64,1,rep,packed,name=uids" json:"uids,omitempty"`
}

func (m *UidList) Reset()                    { *m = UidList{} }
func (m *UidList) String() string            { return proto.CompactTextString(m) }
func (*UidList) ProtoMessage()               {}
func (*UidList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Result struct {
	Values    [][]byte   `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	Uidmatrix []*UidList `protobuf:"bytes,2,rep,name=uidmatrix" json:"uidmatrix,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetUidmatrix() []*UidList {
	if m != nil {
		return m.Uidmatrix
	}
	return nil
}

type SubGraph struct {
	Attr     string      `protobuf:"bytes,1,opt,name=attr" json:"attr,omitempty"`
	Children []*SubGraph `protobuf:"bytes,2,rep,name=children" json:"children,omitempty"`
	Result   *Result     `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *SubGraph) Reset()                    { *m = SubGraph{} }
func (m *SubGraph) String() string            { return proto.CompactTextString(m) }
func (*SubGraph) ProtoMessage()               {}
func (*SubGraph) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SubGraph) GetChildren() []*SubGraph {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *SubGraph) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*UidList)(nil), "protocolbuffer.UidList")
	proto.RegisterType((*Result)(nil), "protocolbuffer.Result")
	proto.RegisterType((*SubGraph)(nil), "protocolbuffer.SubGraph")
}

var fileDescriptor0 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2e, 0x4d, 0x4a,
	0x2f, 0x4a, 0x2c, 0xc8, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x53, 0xc9, 0xf9,
	0x39, 0x49, 0xa5, 0x69, 0x69, 0xa9, 0x45, 0x4a, 0x8a, 0x5c, 0xec, 0xa1, 0x99, 0x29, 0x3e, 0x99,
	0xc5, 0x25, 0x42, 0x62, 0x5c, 0x2c, 0xa5, 0x99, 0x29, 0xc5, 0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x6c,
	0x4e, 0x4c, 0x02, 0x8c, 0x41, 0x60, 0xbe, 0x52, 0x38, 0x17, 0x5b, 0x50, 0x6a, 0x71, 0x69, 0x0e,
	0x48, 0x05, 0x5b, 0x59, 0x62, 0x4e, 0x69, 0x2a, 0x44, 0x0d, 0x4f, 0x10, 0x94, 0x27, 0x64, 0xca,
	0xc5, 0x09, 0x54, 0x99, 0x9b, 0x58, 0x52, 0x94, 0x59, 0x21, 0xc1, 0x04, 0x94, 0xe2, 0x36, 0x12,
	0xd7, 0x43, 0xb5, 0x48, 0x0f, 0x6a, 0x4b, 0x10, 0x42, 0xa5, 0x52, 0x0b, 0x23, 0x17, 0x47, 0x70,
	0x69, 0x92, 0x3b, 0xc8, 0x79, 0x42, 0x42, 0x5c, 0x2c, 0x89, 0x25, 0x25, 0x45, 0x40, 0x93, 0x19,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x13, 0x2e, 0x8e, 0xe4, 0x8c, 0xcc, 0x9c, 0x94, 0xa2, 0xd4,
	0x3c, 0xa8, 0xb1, 0x12, 0xe8, 0xc6, 0xc2, 0xf4, 0x07, 0xc1, 0x55, 0x0a, 0xe9, 0x71, 0xb1, 0x15,
	0x81, 0xdd, 0x2b, 0xc1, 0x0c, 0x34, 0x8b, 0xdb, 0x48, 0x0c, 0x5d, 0x0f, 0xc4, 0x37, 0x41, 0x50,
	0x55, 0x49, 0x6c, 0x60, 0x69, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x1f, 0xc8, 0xb1,
	0x2b, 0x01, 0x00, 0x00,
}
