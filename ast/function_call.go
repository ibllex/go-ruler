package ast

// FunctionCall function call ast node
type FunctionCall struct {
	FuncName string
	Params   []Node
}

func (fn *FunctionCall) Literal() string {
	return fn.FuncName
}

// NewFunctionCall construct function call node
func NewFunctionCall(funcName string, params []Node) *FunctionCall {
	return &FunctionCall{
		FuncName: funcName,
		Params:   params,
	}
}
