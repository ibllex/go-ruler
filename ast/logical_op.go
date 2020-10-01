package ast

import "github.com/ibllex/go-ruler/token"

// LogicalOp logical operator, and/or/xor
type LogicalOp struct {
	left  AST
	right AST
	token token.Token
}

// NewLogicalOp construct a logical operator
func NewLogicalOp(left AST, tk token.Token, right AST) LogicalOp {
	return LogicalOp{
		left:  left,
		right: right,
		token: tk,
	}
}
