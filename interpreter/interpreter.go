package interpreter

import "github.com/ibllex/go-ruler/ast"

// T means Target which is a shortcut for map[string]interface{}
type T map[string]interface{}

// P means params which is a shortcut for map[string]interface{}
type P map[string]interface{}

// Interpreter exec rule and return true or false
type Interpreter struct {
	root   ast.Node
	target T
	params P
}

// New construct Interpreter
func New(tree ast.Node, target T, params P) Interpreter {
	return Interpreter{
		root:   tree,
		target: target,
		params: params,
	}
}
