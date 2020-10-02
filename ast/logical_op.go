package ast

import "github.com/ibllex/go-ruler/token"

// LogicalOp logical operator, and/or/xor
type LogicalOp struct {
	left  Node
	right Node
	token token.Token
}

func (l *LogicalOp) Literal() string {
	return l.token.Literal
}

// NewLogicalOp construct a logical operator
func NewLogicalOp(left Node, tk token.Token, right Node) *LogicalOp {
	return &LogicalOp{
		left:  left,
		right: right,
		token: tk,
	}
}
