package interpreter

import (
	"github.com/ibllex/go-ruler/ast"
	"github.com/ibllex/go-ruler/object"
)

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

func (i *Interpreter) eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.LogicalOp:
		return i.evalLogicalOp(node)
	}

	return &object.Null{}
}

func (i *Interpreter) evalLogicalOp(node *ast.LogicalOp) object.Object {

	// left := i.eval(node.Left)
	// right := i.eval(node.Right)

	// switch node.Token.Type {
	// case token.AND:
	// 	return left.Cast(object.BOOLEAN).(*object.Boolean).Value && right.Cast(object.BOOLEAN).(*object.Boolean).Value
	// case token.OR:
	// 	return left || right
	// case token.XOR:
	// 	return left != right
	// case token.EQUAL:
	// 	return left == right
	// case token.GT:
	// 	return left > right
	// case token.LT:
	// 	return left < right
	// case token.NOT_EQUAL:
	// 	return left != right
	// case token.GT_OR_EQUAL:
	// 	return left >= right
	// case token.LT_OR_EQUAL:
	// 	return left <= right
	// }

	return &object.Boolean{Value: false}
}

// Exec Execute the expression and return the result
func (i *Interpreter) Exec() object.Object {
	return i.eval(i.root)
}

// New construct Interpreter
func New(tree ast.Node, target T, params P) Interpreter {
	return Interpreter{
		root:   tree,
		target: target,
		params: params,
	}
}
