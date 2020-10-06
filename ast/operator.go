package ast

// Operator function call ast node
type Operator struct {
	Name   string
	Params []Node
}

func (op *Operator) Literal() string {
	return op.Name
}

// NewOperator construct function call node
func NewOperator(name string, params []Node) *Operator {
	return &Operator{
		Name:   name,
		Params: params,
	}
}
