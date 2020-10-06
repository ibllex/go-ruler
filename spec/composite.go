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
			params[k] = p
		}
	}

	return params
}

// AddSpecification add specification after construct
func (c *Composite) AddSpecification(spec Specification) {
	c.specifications = append(c.specifications, spec)
}
