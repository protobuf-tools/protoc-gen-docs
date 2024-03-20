// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package message

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Descriptor describes a message type.
type Descriptor struct {
	desc  *descriptorpb.DescriptorProto
	field *descriptorpb.FieldDescriptorProto
}
