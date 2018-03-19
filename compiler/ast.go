package compiler

// AST is an abstract syntax tree.
type AST struct {
	Descendants []*AST
	Type        uint
	Location
}
