// Code generated by protoc-gen-go.
// source: server/pfs/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/persist/persist.proto

It has these top-level messages:
	Append
	BranchClock
	Repo
	Branch
	Diff
	Commit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Append struct {
	// a document only has one of these fields
	Blockrefs []string        `protobuf:"bytes,1,rep,name=blockrefs" json:"blockrefs,omitempty"`
	Children  map[string]bool `protobuf:"bytes,2,rep,name=children" json:"children,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Delete    bool            `protobuf:"varint,3,opt,name=delete" json:"delete,omitempty"`
}

func (m *Append) Reset()                    { *m = Append{} }
func (m *Append) String() string            { return proto.CompactTextString(m) }
func (*Append) ProtoMessage()               {}
func (*Append) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Append) GetChildren() map[string]bool {
	if m != nil {
		return m.Children
	}
	return nil
}

type BranchClock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
	// or this field, but not both
	Intersection []*BranchClock `protobuf:"bytes,3,rep,name=intersection" json:"intersection,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BranchClock) GetIntersection() []*BranchClock {
	if m != nil {
		return m.Intersection
	}
	return nil
}

type Repo struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type Branch struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Diff struct {
	ID       string    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CommitID string    `protobuf:"bytes,2,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string    `protobuf:"bytes,3,opt,name=path" json:"path,omitempty"`
	Appends  []*Append `protobuf:"bytes,4,rep,name=appends" json:"appends,omitempty"`
	Size     uint64    `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Diff) GetAppends() []*Append {
	if m != nil {
		return m.Appends
	}
	return nil
}

type Commit struct {
	ID         string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo       string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Clock      []*BranchClock             `protobuf:"bytes,3,rep,name=clock" json:"clock,omitempty"`
	Started    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished   *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Provenance []string                   `protobuf:"bytes,6,rep,name=provenance" json:"provenance,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Commit) GetClock() []*BranchClock {
	if m != nil {
		return m.Clock
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func init() {
	proto.RegisterType((*Append)(nil), "Append")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*Branch)(nil), "Branch")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
}

func init() { proto.RegisterFile("server/pfs/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0x34, 0x6d, 0xa7, 0x0b, 0x42, 0x16, 0xa0, 0xa8, 0x20, 0x28, 0x39, 0xed, 0x29,
	0x85, 0x05, 0x21, 0x04, 0x17, 0xa0, 0x70, 0xe0, 0x86, 0x2c, 0xee, 0x28, 0x3f, 0x93, 0x8d, 0xb5,
	0x89, 0x6d, 0xd9, 0xde, 0x4a, 0xcb, 0x81, 0x47, 0xe2, 0x91, 0x78, 0x16, 0x6c, 0xc7, 0xd9, 0xed,
	0x0a, 0xa4, 0x72, 0xca, 0xcc, 0xe4, 0xfb, 0xb1, 0xbf, 0x31, 0x6c, 0x34, 0xaa, 0x3d, 0xaa, 0xad,
	0x6c, 0xf4, 0x56, 0xa2, 0xd2, 0x4c, 0x9b, 0xf1, 0x9b, 0x4b, 0x25, 0x8c, 0x58, 0x3f, 0x3d, 0x17,
	0xe2, 0xbc, 0xc3, 0xad, 0xef, 0xca, 0xcb, 0x66, 0x6b, 0x58, 0x8f, 0xda, 0x14, 0xbd, 0x1c, 0x00,
	0xd9, 0xaf, 0x08, 0x92, 0x0f, 0x52, 0x22, 0xaf, 0xc9, 0x63, 0x58, 0x96, 0x9d, 0xa8, 0x2e, 0x14,
	0x36, 0x3a, 0x8d, 0x36, 0xd3, 0xd3, 0x25, 0xbd, 0x19, 0x90, 0x17, 0xb0, 0xa8, 0x5a, 0xd6, 0xd5,
	0x0a, 0x79, 0x3a, 0xb1, 0x3f, 0x57, 0x67, 0x0f, 0xf2, 0x81, 0x98, 0xef, 0xc2, 0xfc, 0x33, 0x37,
	0xea, 0x8a, 0x5e, 0xc3, 0xc8, 0x43, 0x48, 0x6a, 0xec, 0xd0, 0x60, 0x3a, 0xdd, 0x44, 0xa7, 0x0b,
	0x1a, 0xba, 0xf5, 0x3b, 0xb8, 0x73, 0x8b, 0x42, 0xee, 0xc1, 0xf4, 0x02, 0xaf, 0xac, 0x67, 0x64,
	0x3d, 0x5d, 0x49, 0xee, 0xc3, 0x6c, 0x5f, 0x74, 0x97, 0x68, 0xad, 0x1c, 0x73, 0x68, 0xde, 0x4e,
	0xde, 0x44, 0x59, 0x0f, 0xab, 0x8f, 0xaa, 0xe0, 0x55, 0xbb, 0x73, 0x47, 0x73, 0x1e, 0xa5, 0x6f,
	0x03, 0x3b, 0x74, 0x4e, 0xa0, 0x72, 0x00, 0x2f, 0x10, 0xd3, 0xa1, 0x21, 0xcf, 0xe1, 0x84, 0x71,
	0x63, 0x13, 0xc2, 0xca, 0x30, 0xc1, 0xed, 0xb9, 0xdc, 0x45, 0x4e, 0xf2, 0x03, 0x45, 0x7a, 0x0b,
	0x91, 0x7d, 0x85, 0x98, 0xa2, 0x14, 0x84, 0x40, 0xcc, 0x8b, 0x1e, 0x83, 0x8b, 0xaf, 0xc9, 0x2b,
	0x98, 0x57, 0x0a, 0x0b, 0x83, 0xb5, 0x77, 0x59, 0x9d, 0xad, 0xf3, 0x21, 0xee, 0x7c, 0x8c, 0x3b,
	0xff, 0x36, 0xc6, 0x4d, 0x47, 0x68, 0xf6, 0x1e, 0x92, 0xc1, 0x8e, 0xdc, 0x85, 0x09, 0xab, 0x83,
	0xa2, 0xad, 0x9c, 0x87, 0xb2, 0x5e, 0x5e, 0xcc, 0x7a, 0xa8, 0x43, 0xdf, 0xe9, 0x8d, 0x6f, 0xf6,
	0x13, 0xe2, 0x4f, 0xac, 0x69, 0xfe, 0xe2, 0x3f, 0x82, 0x65, 0x25, 0xfa, 0x9e, 0x99, 0xef, 0xac,
	0x0e, 0x22, 0x8b, 0x61, 0xf0, 0xc5, 0x8b, 0xcb, 0xc2, 0xb4, 0xa3, 0x90, 0xab, 0xc9, 0x33, 0x98,
	0x17, 0x7e, 0x85, 0x3a, 0x8d, 0x7d, 0x12, 0xf3, 0xb0, 0x52, 0x3a, 0xce, 0x1d, 0x4d, 0xb3, 0x1f,
	0x98, 0xce, 0x7c, 0x8c, 0xbe, 0xce, 0x7e, 0xdb, 0x37, 0xb3, 0xf3, 0xba, 0xff, 0x75, 0x85, 0x6c,
	0x5c, 0xc5, 0xbf, 0xd2, 0x0e, 0x8b, 0xb1, 0x51, 0xda, 0x98, 0x94, 0x8b, 0x32, 0x3e, 0x1e, 0x65,
	0x80, 0x92, 0xd7, 0xb0, 0x68, 0x18, 0x67, 0xba, 0xb5, 0xb4, 0xd9, 0x51, 0xda, 0x35, 0x96, 0x3c,
	0x01, 0xb0, 0xff, 0xf7, 0xc8, 0xed, 0x39, 0x30, 0x4d, 0xfc, 0x53, 0x3f, 0x98, 0x94, 0x89, 0x67,
	0xbf, 0xfc, 0x13, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xb1, 0xb8, 0xb9, 0x60, 0x03, 0x00, 0x00,
}
