package main

import "github.com/mburbidg/gkit-parser/gen"

type Visitor struct {
	gen.GQLVisitor
}

func NewVisitor() gen.GQLVisitor {
	return &Visitor{}
}
