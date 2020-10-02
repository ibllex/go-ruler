package ast

import "github.com/ibllex/go-ruler/token"

// LogicalOp logical operator, and/or/xor
type LogicalOp struct {
	Left  Node
	Right Node
	Token token.Token
}

func (l *LogicalOp) Literal() string {
	return l.Token.Literal
}

// NewLogicalOp construct a logical operator
func NewLogicalOp(left Node, tk token.Token, right Node) *LogicalOp {
	return &LogicalOp{
		Left:  left,
		Right: right,
		Token: tk,
	}
}
