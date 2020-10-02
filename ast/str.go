package ast

import "github.com/ibllex/go-ruler/token"

// Str string const node
type Str struct {
	Token token.Token
	Value string
}

func (s *Str) Literal() string {
	return s.Token.Literal
}

// NewStr construct a number node
func NewStr(tk token.Token) *Str {
	return &Str{
		Token: tk,
		Value: tk.Literal,
	}
}
