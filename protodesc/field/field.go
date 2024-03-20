// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package field

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Descriptor describes a field within a message.
type Descriptor struct {
	desc *descriptorpb.FieldDescriptorProto
}
