// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/message/proto/message.proto

package proto

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

type DataNode_Type int32

const (
	DataNode_UNKNOWN DataNode_Type = 0
	DataNode_CLIENT  DataNode_Type = 1
	DataNode_SERVER  DataNode_Type = 2
)

var DataNode_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLIENT",
	2: "SERVER",
}

var DataNode_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"CLIENT":  1,
	"SERVER":  2,
}

func (x DataNode_Type) Enum() *DataNode_Type {
	p := new(DataNode_Type)
	*p = x
	return p
}

func (x DataNode_Type) String() string {
	return proto.EnumName(DataNode_Type_name, int32(x))
}

func (x *DataNode_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DataNode_Type_value, data, "DataNode_Type")
	if err != nil {
		return err
	}
	*x = DataNode_Type(value)
	return nil
}

func (DataNode_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_fa2fcc833ce07832, []int{1, 0}
}

// Constants defines constants with default values.
type Constants struct {
	Magic                *uint64  `protobuf:"varint,1,opt,name=magic,def=257787339638762" json:"magic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Constants) Reset()         { *m = Constants{} }
func (m *Constants) String() string { return proto.CompactTextString(m) }
func (*Constants) ProtoMessage()    {}
func (*Constants) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa2fcc833ce07832, []int{0}
}

func (m *Constants) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Constants.Unmarshal(m, b)
}
func (m *Constants) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Constants.Marshal(b, m, deterministic)
}
func (m *Constants) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Constants.Merge(m, src)
}
func (m *Constants) XXX_Size() int {
	return xxx_messageInfo_Constants.Size(m)
}
func (m *Constants) XXX_DiscardUnknown() {
	xxx_messageInfo_Constants.DiscardUnknown(m)
}

var xxx_messageInfo_Constants proto.InternalMessageInfo

const Default_Constants_Magic uint64 = 257787339638762

func (m *Constants) GetMagic() uint64 {
	if m != nil && m.Magic != nil {
		return *m.Magic
	}
	return Default_Constants_Magic
}

// Datanode is something that see's a message AND can modify it.
type DataNode struct {
	Type *DataNode_Type `protobuf:"varint,1,opt,name=type,enum=message.DataNode_Type,def=1" json:"type,omitempty"`
	Name *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Port *string        `protobuf:"bytes,4,opt,name=port" json:"port,omitempty"`
	// 8 bytes of timestamp in pcap-friendly network byte order.
	TimestampUsec        []byte   `protobuf:"bytes,3,opt,name=timestamp_usec,json=timestampUsec" json:"timestamp_usec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataNode) Reset()         { *m = DataNode{} }
func (m *DataNode) String() string { return proto.CompactTextString(m) }
func (*DataNode) ProtoMessage()    {}
func (*DataNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa2fcc833ce07832, []int{1}
}

func (m *DataNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataNode.Unmarshal(m, b)
}
func (m *DataNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataNode.Marshal(b, m, deterministic)
}
func (m *DataNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataNode.Merge(m, src)
}
func (m *DataNode) XXX_Size() int {
	return xxx_messageInfo_DataNode.Size(m)
}
func (m *DataNode) XXX_DiscardUnknown() {
	xxx_messageInfo_DataNode.DiscardUnknown(m)
}

var xxx_messageInfo_DataNode proto.InternalMessageInfo

const Default_DataNode_Type DataNode_Type = DataNode_CLIENT

func (m *DataNode) GetType() DataNode_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_DataNode_Type
}

func (m *DataNode) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DataNode) GetPort() string {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return ""
}

func (m *DataNode) GetTimestampUsec() []byte {
	if m != nil {
		return m.TimestampUsec
	}
	return nil
}

// Msg is a message sent over the network.
// magic, seq, src and dst are required fields.
type Msg struct {
	Magic *uint64 `protobuf:"fixed64,1,opt,name=magic" json:"magic,omitempty"`
	// 8 bytes of sequence in pcap-friendly network byte order.
	Seq []byte `protobuf:"bytes,2,opt,name=seq" json:"seq,omitempty"`
	// Datanodes seen by this message.
	Src   *DataNode   `protobuf:"bytes,3,opt,name=src" json:"src,omitempty"`
	Dst   *DataNode   `protobuf:"bytes,4,opt,name=dst" json:"dst,omitempty"`
	Nodes []*DataNode `protobuf:"bytes,5,rep,name=nodes" json:"nodes,omitempty"`
	// Optional payload.
	Payload              []byte   `protobuf:"bytes,99,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_fa2fcc833ce07832, []int{2}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetMagic() uint64 {
	if m != nil && m.Magic != nil {
		return *m.Magic
	}
	return 0
}

func (m *Msg) GetSeq() []byte {
	if m != nil {
		return m.Seq
	}
	return nil
}

func (m *Msg) GetSrc() *DataNode {
	if m != nil {
		return m.Src
	}
	return nil
}

func (m *Msg) GetDst() *DataNode {
	if m != nil {
		return m.Dst
	}
	return nil
}

func (m *Msg) GetNodes() []*DataNode {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *Msg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("message.DataNode_Type", DataNode_Type_name, DataNode_Type_value)
	proto.RegisterType((*Constants)(nil), "message.Constants")
	proto.RegisterType((*DataNode)(nil), "message.DataNode")
	proto.RegisterType((*Msg)(nil), "message.Msg")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/message/proto/message.proto", fileDescriptor_fa2fcc833ce07832)
}

var fileDescriptor_fa2fcc833ce07832 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x4f, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0xb4, 0xb4, 0x80, 0x3c, 0x10, 0xeb, 0xc6, 0x98, 0x3d, 0x36, 0x18, 0x22, 0x89, 0x49, 0x9b,
	0x94, 0x48, 0x81, 0x2b, 0x72, 0x30, 0x6a, 0x4d, 0x56, 0xd0, 0xa3, 0x59, 0xda, 0x4d, 0x25, 0xa1,
	0xdd, 0xb5, 0xbb, 0x1c, 0xf8, 0x6b, 0xfe, 0x3a, 0xb3, 0xfd, 0x30, 0x26, 0x72, 0x9b, 0x79, 0x33,
	0x2f, 0x33, 0x03, 0xb3, 0x64, 0xab, 0x3e, 0xf7, 0x1b, 0x37, 0xe2, 0xa9, 0x97, 0x70, 0x9e, 0xec,
	0x98, 0x17, 0xed, 0xf8, 0x3e, 0x16, 0x39, 0xdf, 0xb0, 0xdc, 0x4b, 0x99, 0x94, 0x34, 0x61, 0x9e,
	0xc8, 0xb9, 0xe2, 0x35, 0x73, 0x0b, 0x86, 0xda, 0x15, 0x1d, 0xf8, 0xd0, 0x59, 0xf0, 0x4c, 0x2a,
	0x9a, 0x29, 0x89, 0x86, 0xd0, 0x4c, 0x69, 0xb2, 0x8d, 0xb0, 0xe1, 0x18, 0x23, 0x6b, 0x7e, 0xee,
	0xdf, 0x05, 0xc1, 0x34, 0x18, 0x8f, 0x67, 0x93, 0xf1, 0x34, 0x98, 0xf8, 0xa4, 0x54, 0x07, 0xdf,
	0x06, 0x9c, 0xde, 0x53, 0x45, 0x43, 0x1e, 0x33, 0xe4, 0x83, 0xa5, 0x0e, 0x82, 0x15, 0x2f, 0x7d,
	0xff, 0xca, 0xad, 0x73, 0x6a, 0x83, 0xbb, 0x3a, 0x08, 0x36, 0x6f, 0x2d, 0x9e, 0x1e, 0x96, 0xe1,
	0x8a, 0x14, 0x5e, 0x84, 0xc0, 0xca, 0x68, 0xca, 0x70, 0xc3, 0x31, 0x46, 0x1d, 0x52, 0x60, 0x7d,
	0x13, 0x3c, 0x57, 0xd8, 0x2a, 0x6f, 0x1a, 0xa3, 0x21, 0xf4, 0xd5, 0x36, 0x65, 0x52, 0xd1, 0x54,
	0x7c, 0xec, 0x25, 0x8b, 0xb0, 0xe9, 0x18, 0xa3, 0x1e, 0x39, 0xfb, 0xbd, 0xae, 0x25, 0x8b, 0x06,
	0xb7, 0x60, 0xe9, 0x10, 0xd4, 0x85, 0xf6, 0x3a, 0x7c, 0x0c, 0x5f, 0xde, 0x43, 0xfb, 0x04, 0x01,
	0x54, 0x99, 0xb6, 0xa1, 0xf1, 0xeb, 0x92, 0xbc, 0x2d, 0x89, 0xdd, 0xd0, 0xe5, 0xcd, 0x67, 0x99,
	0xa0, 0xcb, 0xbf, 0x5b, 0x5b, 0xd5, 0x34, 0x64, 0x83, 0x29, 0xd9, 0x57, 0x51, 0xac, 0x47, 0x34,
	0x44, 0xd7, 0x60, 0xca, 0xbc, 0x0c, 0xee, 0xfa, 0x17, 0xff, 0xe6, 0x11, 0xad, 0x6a, 0x53, 0x2c,
	0xcb, 0xee, 0xc7, 0x4d, 0xb1, 0x54, 0xe8, 0x06, 0x9a, 0x19, 0x8f, 0x99, 0xc4, 0x4d, 0xc7, 0x3c,
	0x6e, 0x2b, 0x75, 0x84, 0xa1, 0x2d, 0xe8, 0x61, 0xc7, 0x69, 0x8c, 0xa3, 0xa2, 0x48, 0x4d, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xc2, 0x91, 0x4d, 0xf2, 0x01, 0x00, 0x00,
}