// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package file

import (
	"google.golang.org/protobuf/types/descriptorpb"

	"github.com/protobuf-tools/protoc-gen-docs/protodesc"
	"github.com/protobuf-tools/protoc-gen-docs/protodesc/enum"
	"github.com/protobuf-tools/protoc-gen-docs/protodesc/message"
	"github.com/protobuf-tools/protoc-gen-docs/protodesc/service"
)

// Descriptor describes a complete .proto file.
type Descriptor struct {
	desc *descriptorpb.FileDescriptorProto

	Comments       *protodesc.Comment
	SyntaxComments *protodesc.Comment

	Services []*service.Descriptor
	Enums    []*enum.Descriptor
	Messages []*message.Descriptor
}
