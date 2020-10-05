package object

import (
	"bytes"
	"strings"
)

// Array Array builtin type
type Array struct {
	Elements []Object
}

// Type type string
func (a *Array) Type() Type {
	return ARRAY
}

// Inspect return value as string
func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// Equals returns if two objects are equal
func (a *Array) Equals(o Object) bool {

	// type cast left object if the two operands are not the same type
	if a.Type() != o.Type() {
		o = o.Cast(a.Type())
	}

	ar, ok := o.(*Array)
	if !ok || len(a.Elements) != len(ar.Elements) {
		return false
	}

	// returns false if any one of the elements are not equal
	for i, e := range a.Elements {
		if !e.Equals(ar.Elements[i]) {
			return false
		}
	}

	return true
}

// Cast type cast
func (a *Array) Cast(t Type) Object {

	switch t {
	case STRING:
		return &String{a.Inspect()}
	case INTEGER:
		return &Integer{int64(len(a.Elements))}
	case BOOLEAN:
		if len(a.Elements) > 0 {
			return &Boolean{true}
		}
		return &Boolean{false}
	case FLOAT:
		return &Float{float64(len(a.Elements))}
	case ARRAY:
		return a
	}

	return &Null{}
}
