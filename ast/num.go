package ast

import "github.com/ibllex/go-ruler/token"

// Num num node, include float and integer
type Num struct {
	Token token.Token
	Value string
}

func (n *Num) Literal() string {
	return n.Token.Literal
}

// NewNum construct a number node
func NewNum(tk token.Token) *Num {
	return &Num{
		Token: tk,
		Value: tk.Literal,
	}
}
