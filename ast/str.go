package ast

import "github.com/ibllex/go-ruler/token"

// Str string const node
type Str struct {
	token token.Token
	value string
}

func (s *Str) Literal() string {
	return s.token.Literal
}

// NewStr construct a number node
func NewStr(tk token.Token) *Str {
	return &Str{
		token: tk,
		value: tk.Literal,
	}
}
