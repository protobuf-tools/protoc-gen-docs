// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/protobuf-tools/protoc-gen-docs/protodesc"
)

// Descriptor describes an enum type.
type Descriptor struct {
	desc    *descriptorpb.EnumDescriptorProto
	comment *protodesc.Comment

	Values ValueDescriptors
}

// NewDescriptor returns the new named [*Descriptor].
func NewDescriptor(name string) *Descriptor {
	return &Descriptor{
		desc: &descriptorpb.EnumDescriptorProto{
			Name: proto.String(name),
		},
	}
}

// GetName gets the enum name.
func (d *Descriptor) GetName() string {
	return d.desc.GetName()
}

// GetValue gets the enum values.
func (d *Descriptor) GetValue() ValueDescriptors {
	return d.Values
}

// GetOptions gets the enum options.
func (d *Descriptor) GetOptions() *descriptorpb.EnumOptions {
	return d.desc.GetOptions()
}

// GetReservedRange gets the enum reserved range.
func (d *Descriptor) GetReservedRange() []*descriptorpb.EnumDescriptorProto_EnumReservedRange {
	return d.desc.GetReservedRange()
}

// GetReservedName gets the enum reserved names.
func (d *Descriptor) GetReservedName() []string {
	return d.desc.GetReservedName()
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
	// TODO(zchee): normalize
	d.comment.Trailing = trailing
}

// SetDetachedComment sets the enum detached comments.
func (d *Descriptor) SetDetachedComment(detached []string) {
	// TODO(zchee): normalize
	d.comment.Detached = detached
}

// Descriptor returns the enum descriptor.
func (d *Descriptor) Descriptor() *descriptorpb.EnumDescriptorProto {
	return d.desc
}

// ValueDescriptor describes a value within an enum.
type ValueDescriptor struct {
	desc *descriptorpb.EnumValueDescriptorProto
}

// NewValueDescriptor returns the new named [*ValueDescriptor] with number.
func NewValueDescriptor(name string, number int32) *ValueDescriptor {
	return &ValueDescriptor{
		desc: &descriptorpb.EnumValueDescriptorProto{
			Name:   proto.String(name),
			Number: proto.Int32(number),
		},
	}
}

// GetName gets the enum name.
func (d *ValueDescriptor) GetName() string {
	return d.desc.GetName()
}

// SetDeprecated sets the deprecated itself.
func (d *ValueDescriptor) SetDeprecated() {
	d.desc.Options.Deprecated = proto.Bool(true)
}

// Descriptor returns the enum value descriptor.
func (d *ValueDescriptor) Descriptor() *descriptorpb.EnumValueDescriptorProto {
	return d.desc
}

// ValueDescriptors describes a slice of value within an enum.
type ValueDescriptors []*ValueDescriptor

// Descriptor returns the enum value descriptors.
func (s ValueDescriptors) Descriptor() []*descriptorpb.EnumValueDescriptorProto {
	values := make([]*descriptorpb.EnumValueDescriptorProto, len(s))
	for i, ev := range s {
		values[i] = ev.Descriptor()
	}
	return values
}
