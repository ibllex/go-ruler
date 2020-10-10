package spec

import (
	"github.com/ibllex/go-ruler/interpreter"
)

// Specification specification base interface
type Specification interface {
	Rule() string
	Params() interpreter.P
	PositionalParams() interpreter.PP
}
