// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package oneof

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Descriptor describes a oneof.
type Descriptor struct {
	desc *descriptorpb.OneofDescriptorProto
}
