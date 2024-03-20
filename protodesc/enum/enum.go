// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package enum

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Descriptor describes an enum type.
type Descriptor struct {
	desc  *descriptorpb.EnumDescriptorProto
	value *descriptorpb.EnumValueDescriptorProto
}
