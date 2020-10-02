package ast

// FunctionCall function call ast node
type FunctionCall struct {
	funcName string
	params   []Node
}

func (fn *FunctionCall) Literal() string {
	return fn.funcName
}

// NewFunctionCall construct function call node
func NewFunctionCall(funcName string, params []Node) *FunctionCall {
	return &FunctionCall{
		funcName: funcName,
		params:   params,
	}
}
