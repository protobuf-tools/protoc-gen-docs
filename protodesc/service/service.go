// Copyright 2024 The protobuf-tools Authors
// SPDX-License-Identifier: Apache-2.0

package service

import (
	"google.golang.org/protobuf/types/descriptorpb"
)

// Descriptor describes a service.
type Descriptor struct {
	desc   *descriptorpb.ServiceDescriptorProto
	method *descriptorpb.MethodDescriptorProto
}
