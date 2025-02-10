# Gkit-parser
The Gkit parser is a language independent, standalone parser for GQL. It runs as an executable and communicates with other components of a GQL Implementation using pipes. It takes a GQL-Program, and produces an Abstract Syntax Tree (AST), symbol table and error results. It parses a GQL program and applies the Syntax Rules defined by the GQL Standard.

The Gkit-parser uses the [Protobuf](https://protobuf.dev/) to serialize the AST, symbol table and error results.

The Gkit-parser can be used by a GQL Implementation as the frontend for the implementation. The implementation can be written in any language supported by Protobuf.

The Protobuf type definitions are defined [gkit-ast](https://github.com/mburbidg/gkit-ast). That repository also contains the compiled type definitions for all the languages supported by Protobuf.

For developmemnt purposes the Gkit-parser can be run as a CLI, returning the AST, symbol table and results to stdout encoded as Protobuf produced JSON. See below for command options.
