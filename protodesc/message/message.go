// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package message

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/protobuf-tools/protoc-gen-docs/protodesc"
	"github.com/protobuf-tools/protoc-gen-docs/protodesc/enum"
	"github.com/protobuf-tools/protoc-gen-docs/protodesc/field"
	"github.com/protobuf-tools/protoc-gen-docs/protodesc/oneof"
)

// Descriptor describes a message type.
type Descriptor struct {
	desc    *descriptorpb.DescriptorProto
	comment *protodesc.Comment

	Fields []*field.Descriptor
	Enums  []*enum.Descriptor
	Oneofs []*oneof.Descriptor
}

// NewDescriptor returns the new named [*Descriptor].
func NewDescriptor(name string) *Descriptor {
	return &Descriptor{
		desc: &descriptorpb.DescriptorProto{
			Name: proto.String(name),
		},
	}
}

// GetName gets the message name.
func (d *Descriptor) GetName() string {
	return d.desc.GetName()
}

// GetComment gets the enum comment.
func (d *Descriptor) GetComment() *protodesc.Comment {
	return d.comment
}

// SetLeadingComment sets the enum leading comment.
func (d *Descriptor) SetLeadingComment(fn, leading string) {
	// TODO(zchee): normalize with fn
	_ = fn
	d.comment.Leading = leading
}

// SetTrailingComment sets the enum trailing comment.
func (d *Descriptor) SetTrailingComment(trailing string) {
	d.comment.Trailing = trailing
}

// SetDetachedComment sets the enum detached comments.
func (d *Descriptor) SetDetachedComment(detached []string) {
	d.comment.Detached = detached
}

// GetOptions gets the fields within an message.
func (d *Descriptor) GetFields() []*field.Descriptor {
	return d.Fields
}

// TODO(zchee): implements
// func (d *Descriptor) GetExtension() []*descriptorpb.FieldDescriptorProto {
// 	return d.Extension
// }

// TODO(zchee): implements
// func (d *Descriptor) GetNestedType() []*Descriptor {
// 	return d.NestedType
// }

// GetOptions gets the enums within an message.
func (d *Descriptor) GetEnums() []*enum.Descriptor {
	return d.Enums
}

// TODO(zchee): implements
// func (d *Descriptor) GetExtensionRange() []*descriptorpb.Descriptor_ExtensionRange {
// 	return x.ExtensionRange
// }

// TODO(zchee): implements.
// func (d *Descriptor) GetOneof() []*oneof.Descriptor {
// 	return d.desc.GetOneofDecl()
// }

// SetOneof sets the oneof declaration.
func (d *Descriptor) SetOneof(decl *oneof.Descriptor) {
	d.Oneofs = append(d.Oneofs, decl)
}

// GetOptions gets the message options.
func (d *Descriptor) GetOptions() *descriptorpb.MessageOptions {
	return d.desc.GetOptions()
}

// TODO(zchee): implements
// func (d *Descriptor) GetReservedRange() []*descriptorpb.Descriptor_ReservedRange {
// 	return d.ReservedRange
// }

// TODO(zchee): implements
// func (d *Descriptor) GetReservedName() []string {
// 	return d.ReservedName
// }

// SetDeprecated sets the deprecated itself.
func (d *Descriptor) SetDeprecated(deprecated bool) {
	d.desc.Options.Deprecated = proto.Bool(deprecated)
}

func (d *Descriptor) IsEmptyField() bool {
	return len(d.desc.Field) == 0 && len(d.desc.EnumType) == 0 && len(d.desc.NestedType) == 0
}

// Descriptor returns the message descriptor.
func (d *Descriptor) Descriptor() *descriptorpb.DescriptorProto {
	return d.desc
}
