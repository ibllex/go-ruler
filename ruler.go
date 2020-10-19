package ruler

import (
	"github.com/ibllex/go-ruler/interpreter"
	"github.com/ibllex/go-ruler/lexer"
	"github.com/ibllex/go-ruler/object"
	"github.com/ibllex/go-ruler/parser"
	"github.com/ibllex/go-ruler/spec"
)

// T is alias of interpreter.T
type T = interpreter.T

// O is alias of interpreter.O
type O = interpreter.O

// P is alias of interpreter.P
type P = interpreter.P

// PP is alias of interpreter.PP
type PP = interpreter.PP

// Interpreter is alias of interpreter.Interpreter
type Interpreter = interpreter.Interpreter

// Ruler ruler engine entry
type Ruler struct {
	operators O
}

func makeInterpreter(rule string, ops O) (*Interpreter, error) {
	p := parser.New(lexer.New(rule))
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	i := interpreter.New(tree, ops)
	return &i, nil
}

// Satisfies returns true if the target matches the rules, false otherwise.
func (r *Ruler) Satisfies(target T, rule string, params P, pParams PP) (bool, error) {
	i, err := makeInterpreter(rule, r.operators)
	if err != nil {
		return false, err
	}

	return object.ToNativeBool(i.Exec(target, params, pParams)), nil
}

// SatisfiesSpec satisfies by specification.
func (r *Ruler) SatisfiesSpec(target T, s spec.Specification) (bool, error) {
	return r.Satisfies(target, s.Rule(), s.Params(), s.PositionalParams())
}

// Filter filter and return all targets that match the rules
func (r *Ruler) Filter(targets []T, rule string, params P, pParams PP) (ret []T, err error) {

	ok := false

	for _, t := range targets {
		if ok, err = r.Satisfies(t, rule, params, pParams); ok {
			ret = append(ret, t)
		}
	}

	return
}

// FilterSpec filter by specification
func (r *Ruler) FilterSpec(targets []T, s spec.Specification) (ret []T, err error) {
	return r.Filter(targets, s.Rule(), s.Params(), s.PositionalParams())
}

// New construct a new ruler
func New(operators O) *Ruler {
	return &Ruler{
		operators: operators,
	}
}
