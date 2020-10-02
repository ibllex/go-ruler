package ast

import "github.com/ibllex/go-ruler/token"

// Boolean boolean type node
type Boolean struct {
	token token.Token
	value bool
}

// Literal token literal
func (b *Boolean) Literal() string {
	return b.token.Literal
}

// NewBoolean construct boolean node
func NewBoolean(tk token.Token) *Boolean {
	v := false

	if tk.Type == token.TRUE {
		v = true
	}

	return &Boolean{
		token: tk,
		value: v,
	}
}
