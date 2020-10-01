package ast

// Null null value ast node
type Null struct {
	//
}

// NewNull construct null node
func NewNull() Null {
	return Null{}
}
