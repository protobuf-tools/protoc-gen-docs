# protoc-gen-docs

[![Go Reference](https://pkg.go.dev/badge/github.com/protobuf-tools/protoc-gen-docs.svg)](https://pkg.go.dev/github.com/protobuf-tools/protoc-gen-docs)

Command `protoc-gen-docs` is the protoc plugin to generate documentation from the Protocol Buffers definitions.

## Philosophy

> [!NOTE]
>
> - [AIP-192: Documentation](https://google.aip.dev/192)
>
> Documentation is one of the most critical aspects of API design.
>
> Users of your API are unable to dig into the implementation to understand the API better; often, the API surface definition
> and its corresponding documentation will be the only things a user has.
>
> Therefore, it is important that documentation be as clear, complete, and unambiguous as possible.

---

> [!NOTE]
>
> - [API Best Practices | Protocol Buffers Documentation](https://protobuf.dev/programming-guides/api/#precisely-concisely)
>
> Precisely, Concisely Document Most Fields and Messages

---

> [!NOTE]
>
> - [Buf - Style guide](https://buf.build/docs/best-practices/style-guide#files-and-packages)
>
> Over-document, and use complete sentences for comments.
