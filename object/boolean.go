package object

import (
	"fmt"
)

// Boolean Boolean builtin type
type Boolean struct {
	Value bool
}

// Type type string
func (b *Boolean) Type() Type {
	return BOOLEAN
}

// Inspect return value as string
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// HashKey implements the HashMap interface
// for comparing whether values are the same or not.
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// Equals returns if two objects are equal
func (b *Boolean) Equals(o Object) bool {

	// type cast left object if the two operands are not the same type
	if b.Type() != o.Type() {
		o = o.Cast(b.Type())
	}

	return b.HashKey() == o.(Hashable).HashKey()
}

// Cast type cast
func (b *Boolean) Cast(t Type) Object {

	switch t {
	case STRING:
		if b.Value {
			return &String{"true"}
		}
		return &String{"false"}
	case INTEGER:
		if b.Value {
			return &Integer{1}
		}
		return &Integer{0}
	case BOOLEAN:
		return b
	case FLOAT:
		if b.Value {
			return &Float{1}
		}
		return &Float{0}
	case ARRAY:
		return &Array{[]Object{b}}
	}

	return &Null{}
}
