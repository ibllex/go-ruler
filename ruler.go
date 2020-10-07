package ruler

import (
	"github.com/ibllex/go-ruler/interpreter"
	"github.com/ibllex/go-ruler/lexer"
	"github.com/ibllex/go-ruler/object"
	"github.com/ibllex/go-ruler/parser"
	"github.com/ibllex/go-ruler/spec"
)

// Ruler ruler engine entry
type Ruler struct {
	operators interpreter.O
}

func makeInterpreter(rule string, ops interpreter.O) (*interpreter.Interpreter, error) {
	p := parser.New(lexer.New(rule))
	tree, err := p.Parse()
	if err != nil {
		return nil, err
	}

	i := interpreter.New(tree, ops)
	return &i, nil
}

// Satisfies returns true if the target matches the rules, false otherwise.
func (r *Ruler) Satisfies(target interpreter.T, rule string, params interpreter.P) (bool, error) {
	i, err := makeInterpreter(rule, r.operators)
	if err != nil {
		return false, err
	}

	return object.ToNativeBool(i.Exec(target, params)), nil
}

// SatisfiesSpec satisfies by specification.
func (r *Ruler) SatisfiesSpec(target interpreter.T, s spec.Specification) (bool, error) {
	return r.Satisfies(target, s.Rule(), s.Params())
}

// Filter filter and return all targets that match the rules
func (r *Ruler) Filter(targets []interpreter.T, rule string, params interpreter.P) (ret []interpreter.T, err error) {

	ok := false

	for _, t := range targets {
		if ok, err = r.Satisfies(t, rule, params); ok {
			ret = append(ret, t)
		}
	}

	return
}

// FilterSpec filter by specification
func (r *Ruler) FilterSpec(targets []interpreter.T, s spec.Specification) (ret []interpreter.T, err error) {
	return r.Filter(targets, s.Rule(), s.Params())
}

// New construct a new ruler
func New(operators interpreter.O) *Ruler {
	return &Ruler{
		operators: operators,
	}
}
