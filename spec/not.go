package spec

import (
	"fmt"

	"github.com/ibllex/go-ruler/interpreter"
)

// NotSpec not specification
type NotSpec struct {
	spec Specification
}

// Rule not specification rule string
func (s *NotSpec) Rule() string {
	return fmt.Sprintf("not(%s)", s.spec.Rule())
}

// Params not specification params
func (s *NotSpec) Params() interpreter.P {
	return s.spec.Params()
}

// PositionalParams not specification positional params
func (s *NotSpec) PositionalParams() interpreter.PP {
	return s.spec.PositionalParams()
}

// Not create not specification
func Not(spec Specification) *NotSpec {
	return &NotSpec{
		spec: spec,
	}
}
