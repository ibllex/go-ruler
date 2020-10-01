package ast

import "github.com/ibllex/go-ruler/token"

// Boolean boolean type node
type Boolean struct {
	token token.Token
	value bool
}

// NewBoolean construct boolean node
func NewBoolean(tk token.Token) Boolean {
	v := false

	if tk.Type == token.TRUE {
		v = true
	}

	return Boolean{
		token: tk,
		value: v,
	}
}
