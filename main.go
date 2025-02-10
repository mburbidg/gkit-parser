package main

import (
	"fmt"
	"github.com/mburbidg/gkit-ast/go/ast"
)

func main() {
	foo := ast.CatalogParentAndName{}
	v := Visitor{}
	fmt.Printf("%v %v\n", foo, v)
}
