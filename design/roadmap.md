# Roadmap for a protocol compiler MVP

The goal of this roadmap is to plan the development of an MVP for the protocol compiler.

## MVP requirements

We think an MVP RPC compiler should deliver usable RPC codegeneration. The shortest path to this goal is to postpone compiler usability features and focus on the minimum required to accomplish RPC generation.

Compiler usability features that can be postponed are:
- support for custom syntax
- support for packages
- support for automated interoperability verification

As a result, the aspects that must be implemented for an MVP are:
- type system definition in Go, which users will use directly to define APIs
- code generation of type and RPC bindings in Go

## Implementation plan

The resulting implementation plan:

[ ] Type system definitions in Go
[ ] Code generation of data type bindings in Go, which decodes from IPLD Data Model
[ ] Code generation of RPC bindings in Go
