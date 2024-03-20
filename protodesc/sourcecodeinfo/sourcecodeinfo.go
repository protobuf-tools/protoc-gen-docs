// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package sourcecodeinfo

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Location represents encapsulated information about the original source file from which a [descriptorpb.FileDescriptorProto] was generated.
type Location struct {
	desc *descriptorpb.SourceCodeInfo_Location
}
