package ruler

import (
	"github.com/ibllex/go-ruler/interpreter"
	"github.com/ibllex/go-ruler/lexer"
	"github.com/ibllex/go-ruler/object"
	"github.com/ibllex/go-ruler/parser"
)

// Ruler ruler engine entry
type Ruler struct {
	//
}

// Satisfies returns true if the target matches the rules, false otherwise.
func (r *Ruler) Satisfies(target interpreter.T, rule string, params interpreter.P) (bool, error) {
	p := parser.New(lexer.New(rule))
	tree, err := p.Parse()
	if err != nil {
		return false, err
	}

	i := interpreter.New(tree, target, params)
	return object.ToNativeBool(i.Exec()), nil
}

// New construct a new ruler
func New() *Ruler {
	return &Ruler{}
}
