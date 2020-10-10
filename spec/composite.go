package spec

import (
	"fmt"
	"strings"

	"github.com/ibllex/go-ruler/interpreter"
)

// Composite join the specifications
type Composite struct {
	operator       string
	specifications []Specification
}

// Rule combine specifications into a rule string and return
func (c *Composite) Rule() string {
	rules := []string{}
	for _, s := range c.specifications {
		rules = append(rules, fmt.Sprintf("(%s)", s.Rule()))
	}
	return strings.Join(rules, fmt.Sprintf(" %s ", c.operator))
}

// Params combine all specifications's params
func (c *Composite) Params() interpreter.P {
	params := interpreter.P{}

	for _, s := range c.specifications {
		for k, p := range s.Params() {
			if params[k] != nil {
				panic(fmt.Sprintf("param %v is alread set.", k))
			}
			params[k] = p
		}
	}

	return params
}

// PositionalParams combine all specifications's positional params
func (c *Composite) PositionalParams() interpreter.PP {
	pParams := interpreter.PP{}

	for _, s := range c.specifications {
		pp := s.PositionalParams()
		for !pp.IsEmpty() {
			pParams.Push(pp.Pop())
		}
	}

	return pParams
}

// AddSpecification add specification after construct
func (c *Composite) AddSpecification(spec Specification) {
	c.specifications = append(c.specifications, spec)
}
