package interpreter

import (
	"strconv"

	"github.com/ibllex/go-ruler/ast"
	"github.com/ibllex/go-ruler/object"
	"github.com/ibllex/go-ruler/token"
	"github.com/ibllex/go-ruler/utils"
)

// Operator func type
type Operator func(args ...object.Object) object.Object

// T means Target which is a shortcut for map[string]interface{}
type T map[string]interface{}

// P means params which is a shortcut for map[string]interface{}
type P map[string]interface{}

// O means operators which is a shortcut for map[string]operator.Operator
type O map[string]Operator

// Interpreter exec rule and return true or false
type Interpreter struct {
	root      ast.Node
	target    T
	params    P
	Operators O
}

func (i *Interpreter) eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.LogicalOp:
		return i.evalLogicalOp(node)
	case *ast.Target:
		return i.evalIdent(node.ID, i.target)
	case *ast.Param:
		return i.evalIdent(node.ID, i.params)
	case *ast.Num:
		return i.evalNum(node)
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}
	case *ast.Str:
		return &object.String{Value: node.Value}
	case *ast.Null:
		return &object.Null{}
	}

	return &object.Null{}
}

func (i *Interpreter) evalNum(node *ast.Num) object.Object {
	if node.Token.Type == token.INTEGER_CONST {
		v, _ := strconv.ParseInt(node.Value, 10, 64)
		return &object.Integer{Value: v}
	}

	v, _ := strconv.ParseFloat(node.Value, 64)
	return &object.Float{Value: v}
}

func (i *Interpreter) evalIdent(node *ast.Ident, source map[string]interface{}) object.Object {
	v := utils.QueryInMapInter(source, node.Path)

	switch v := v.(type) {
	case int:
		return &object.Integer{Value: int64(v)}
	case int8:
		return &object.Integer{Value: int64(v)}
	case int16:
		return &object.Integer{Value: int64(v)}
	case int32:
		return &object.Integer{Value: int64(v)}
	case int64:
		return &object.Integer{Value: v}
	case float32:
		return &object.Float{Value: float64(v)}
	case float64:
		return &object.Float{Value: v}
	case string:
		return &object.String{Value: v}
	case bool:
		return &object.Boolean{Value: v}
	}

	return &object.Null{}
}

func (i *Interpreter) evalLogicalOp(node *ast.LogicalOp) object.Object {

	left := i.eval(node.Left)
	right := i.eval(node.Right)

	value := false

	switch node.Token.Type {
	case token.AND:
		value = object.ToNativeBool(left) && object.ToNativeBool(right)
	case token.OR:
		value = object.ToNativeBool(left) || object.ToNativeBool(right)
	case token.XOR:
		value = !object.IsEqual(left, right)
	case token.NOT_EQUAL:
		value = !object.IsEqual(left, right)
	case token.EQUAL:
		value = object.IsEqual(left, right)
	case token.GT:
		value = object.ToNativeFloat64(left) > object.ToNativeFloat64(right)
	case token.LT:
		value = object.ToNativeFloat64(left) < object.ToNativeFloat64(right)
	case token.GT_OR_EQUAL:
		value = object.ToNativeFloat64(left) >= object.ToNativeFloat64(right)
	case token.LT_OR_EQUAL:
		value = object.ToNativeFloat64(left) <= object.ToNativeFloat64(right)
	}

	return &object.Boolean{Value: value}
}

// Exec Execute the expression and return the result
func (i *Interpreter) Exec(target T, params P) object.Object {
	i.target = target
	i.params = params
	return i.eval(i.root)
}

// New construct Interpreter
func New(tree ast.Node, ops O) Interpreter {
	return Interpreter{
		root:      tree,
		Operators: ops,
	}
}
