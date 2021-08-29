// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: person.proto

package validator

import (
	fmt "fmt"
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

type Person struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c9e10cf24b1156d, []int{0}
}
func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Person)(nil), "main.Person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor_4c9e10cf24b1156d) }

var fileDescriptor_4c9e10cf24b1156d = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x93, 0x72,
	0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x4d, 0x4c, 0xcc, 0x4b,
	0x2c, 0x4e, 0x2c, 0x2e, 0x4d, 0x2a, 0x4a, 0xcc, 0xc8, 0x4d, 0xcc, 0xab, 0x4c, 0xcc, 0xd5, 0x2d,
	0x4e, 0xd1, 0x4f, 0xcf, 0xd7, 0x05, 0xeb, 0xd0, 0x2d, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0xc9,
	0x2f, 0x2a, 0xd6, 0x87, 0x33, 0x21, 0x86, 0x29, 0x99, 0x73, 0xb1, 0x05, 0x80, 0x0d, 0x17, 0x92,
	0xe6, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x62, 0x7f, 0x74,
	0x5f, 0x9e, 0x79, 0x06, 0x23, 0x63, 0x10, 0x58, 0x50, 0x48, 0x80, 0x8b, 0x39, 0x31, 0x3d, 0x55,
	0x82, 0x49, 0x81, 0x51, 0x83, 0x35, 0x08, 0xc4, 0x74, 0xe2, 0x8d, 0xe2, 0xd6, 0x43, 0x98, 0x96,
	0xc4, 0x06, 0x36, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x83, 0x55, 0x73, 0x67, 0xab, 0x00,
	0x00, 0x00,
}
