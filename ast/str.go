package ast

import "github.com/ibllex/go-ruler/token"

// Str string const node
type Str struct {
	token token.Token
	value string
}

// NewStr construct a number node
func NewStr(tk token.Token) Str {
	return Str{
		token: tk,
		value: tk.Literal,
	}
}
