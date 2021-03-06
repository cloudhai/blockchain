// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction.proto

package entity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransactionType int32

const (
	TransactionType_TRANF TransactionType = 0
	TransactionType_ISSUE TransactionType = 1
)

var TransactionType_name = map[int32]string{
	0: "TRANF",
	1: "ISSUE",
}
var TransactionType_value = map[string]int32{
	"TRANF": 0,
	"ISSUE": 1,
}

func (x TransactionType) String() string {
	return proto.EnumName(TransactionType_name, int32(x))
}
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_transaction_96119ffde5683cdf, []int{0}
}

type Transaction struct {
	Type                 TransactionType `protobuf:"varint,1,opt,name=type,proto3,enum=entity.TransactionType" json:"type,omitempty"`
	Hash                 []byte          `protobuf:"bytes,2,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Sign                 []byte          `protobuf:"bytes,3,opt,name=Sign,proto3" json:"Sign,omitempty"`
	Sender               []byte          `protobuf:"bytes,4,opt,name=Sender,proto3" json:"Sender,omitempty"`
	Timestamp            uint64          `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Data                 []byte          `protobuf:"bytes,6,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_96119ffde5683cdf, []int{0}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetType() TransactionType {
	if m != nil {
		return m.Type
	}
	return TransactionType_TRANF
}

func (m *Transaction) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Transaction) GetSign() []byte {
	if m != nil {
		return m.Sign
	}
	return nil
}

func (m *Transaction) GetSender() []byte {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *Transaction) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Transaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BlockHeader struct {
	Version              uint32   `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	MagicCode            uint32   `protobuf:"varint,2,opt,name=MagicCode,proto3" json:"MagicCode,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Height               uint32   `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	Timestamp            uint64   `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	PreHash              []byte   `protobuf:"bytes,6,opt,name=PreHash,proto3" json:"PreHash,omitempty"`
	Miner                []byte   `protobuf:"bytes,7,opt,name=Miner,proto3" json:"Miner,omitempty"`
	Count                uint32   `protobuf:"varint,8,opt,name=Count,proto3" json:"Count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_96119ffde5683cdf, []int{1}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (dst *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(dst, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetMagicCode() uint32 {
	if m != nil {
		return m.MagicCode
	}
	return 0
}

func (m *BlockHeader) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockHeader) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BlockHeader) GetPreHash() []byte {
	if m != nil {
		return m.PreHash
	}
	return nil
}

func (m *BlockHeader) GetMiner() []byte {
	if m != nil {
		return m.Miner
	}
	return nil
}

func (m *BlockHeader) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Block struct {
	Header               *BlockHeader   `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Trans                []*Transaction `protobuf:"bytes,2,rep,name=Trans,proto3" json:"Trans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_transaction_96119ffde5683cdf, []int{2}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTrans() []*Transaction {
	if m != nil {
		return m.Trans
	}
	return nil
}

func init() {
	proto.RegisterType((*Transaction)(nil), "entity.Transaction")
	proto.RegisterType((*BlockHeader)(nil), "entity.BlockHeader")
	proto.RegisterType((*Block)(nil), "entity.Block")
	proto.RegisterEnum("entity.TransactionType", TransactionType_name, TransactionType_value)
}

func init() { proto.RegisterFile("transaction.proto", fileDescriptor_transaction_96119ffde5683cdf) }

var fileDescriptor_transaction_96119ffde5683cdf = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x4e, 0x32, 0x31,
	0x14, 0x85, 0xff, 0xc2, 0xcc, 0xf0, 0x73, 0x11, 0xc5, 0x6a, 0xb0, 0x0b, 0x17, 0x13, 0x36, 0xa2,
	0x24, 0x2c, 0xf0, 0x09, 0x14, 0x35, 0xb8, 0xc0, 0x98, 0xce, 0xe8, 0xd6, 0x54, 0x68, 0xa0, 0x51,
	0xda, 0x49, 0xa7, 0x2e, 0x78, 0x27, 0x5f, 0xc7, 0xf7, 0x31, 0xbd, 0x05, 0x99, 0xb8, 0x71, 0x77,
	0xcf, 0x99, 0x93, 0x33, 0xdf, 0xcd, 0x2d, 0x1c, 0x3a, 0x2b, 0x74, 0x29, 0x66, 0x4e, 0x19, 0x3d,
	0x2c, 0xac, 0x71, 0x86, 0x26, 0x52, 0x3b, 0xe5, 0xd6, 0xbd, 0x4f, 0x02, 0xad, 0x7c, 0xf7, 0x95,
	0x0e, 0x20, 0x72, 0xeb, 0x42, 0x32, 0x92, 0x92, 0xfe, 0xfe, 0xe8, 0x64, 0x18, 0x62, 0xc3, 0x4a,
	0x24, 0x5f, 0x17, 0x92, 0x63, 0x88, 0x52, 0x88, 0x26, 0xa2, 0x5c, 0xb2, 0x5a, 0x4a, 0xfa, 0x7b,
	0x1c, 0x67, 0xef, 0x65, 0x6a, 0xa1, 0x59, 0x3d, 0x78, 0x7e, 0xa6, 0x5d, 0x48, 0x32, 0xa9, 0xe7,
	0xd2, 0xb2, 0x08, 0xdd, 0x8d, 0xa2, 0xa7, 0xd0, 0xcc, 0xd5, 0x4a, 0x96, 0x4e, 0xac, 0x0a, 0x16,
	0xa7, 0xa4, 0x1f, 0xf1, 0x9d, 0xe1, 0x9b, 0x6e, 0x84, 0x13, 0x2c, 0x09, 0x4d, 0x7e, 0xee, 0x7d,
	0x11, 0x68, 0x5d, 0xbf, 0x9b, 0xd9, 0xdb, 0x44, 0x0a, 0xdf, 0xc0, 0xa0, 0xf1, 0x2c, 0x6d, 0xa9,
	0x8c, 0x46, 0xe2, 0x36, 0xdf, 0x4a, 0xdf, 0x3d, 0x15, 0x0b, 0x35, 0x1b, 0x9b, 0xb9, 0x44, 0xc0,
	0x36, 0xdf, 0x19, 0x3f, 0xe4, 0xf5, 0x0a, 0x79, 0x17, 0x92, 0x89, 0x54, 0x8b, 0xa5, 0x43, 0xca,
	0x36, 0xdf, 0xa8, 0x3f, 0x28, 0x19, 0x34, 0x1e, 0xad, 0xc4, 0xb2, 0x00, 0xba, 0x95, 0xf4, 0x18,
	0xe2, 0xa9, 0xd2, 0xd2, 0xb2, 0x06, 0xfa, 0x41, 0x78, 0x77, 0x6c, 0x3e, 0xb4, 0x63, 0xff, 0xf1,
	0x27, 0x41, 0xf4, 0x5e, 0x20, 0xc6, 0xb5, 0xe8, 0xc0, 0x43, 0xf8, 0xd5, 0x70, 0x9f, 0xd6, 0xe8,
	0x68, 0x7b, 0x81, 0xca, 0xd6, 0x7c, 0x13, 0xa1, 0xe7, 0x10, 0xe3, 0x61, 0x58, 0x2d, 0xad, 0x57,
	0xb3, 0x95, 0x6b, 0xf1, 0x90, 0xb8, 0x38, 0x83, 0x83, 0x5f, 0x37, 0xa4, 0x4d, 0x88, 0x73, 0x7e,
	0xf5, 0x70, 0xd7, 0xf9, 0xe7, 0xc7, 0xfb, 0x2c, 0x7b, 0xba, 0xed, 0x90, 0xd7, 0x04, 0xdf, 0xc7,
	0xe5, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x71, 0x96, 0x53, 0x34, 0x02, 0x00, 0x00,
}
