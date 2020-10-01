package ast

// FunctionCall function call ast node
type FunctionCall struct {
	funcName string
	params   []AST
}

// NewFunctionCall construct function call node
func NewFunctionCall(funcName string, params []AST) FunctionCall {
	return FunctionCall{
		funcName: funcName,
		params:   params,
	}
}
