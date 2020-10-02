package ast

// Null null value ast node
type Null struct {
	//
}

func (n *Null) Literal() string {
	return "null"
}

// NewNull construct null node
func NewNull() *Null {
	return &Null{}
}
