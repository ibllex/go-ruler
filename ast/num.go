package ast

import "github.com/ibllex/go-ruler/token"

// Num num node, include float and integer
type Num struct {
	token token.Token
	value string
}

func (n *Num) Literal() string {
	return n.token.Literal
}

// NewNum construct a number node
func NewNum(tk token.Token) *Num {
	return &Num{
		token: tk,
		value: tk.Literal,
	}
}
