package object

import (
	"fmt"
	"strconv"
)

// Integer Integer builtin type
type Integer struct {
	Value int64
}

// Type type string
func (i *Integer) Type() Type {
	return INTEGER
}

// Inspect return value as string
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// HashKey implements the HashMap interface
// for comparing whether values are the same or not.
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// Equals returns if two objects are equal
func (i *Integer) Equals(o Object) bool {
	// type cast left object if the two operands are not the same type
	if i.Type() != o.Type() {
		o = o.Cast(i.Type())
	}

	return i.HashKey() == o.(Hashable).HashKey()
}

// Cast type cast
func (i *Integer) Cast(t Type) Object {

	switch t {
	case STRING:
		return &String{strconv.FormatInt(i.Value, 10)}
	case INTEGER:
		return i
	case BOOLEAN:
		return &Boolean{i.Value != 0}
	case FLOAT:
		return &Float{float64(i.Value)}
	case ARRAY:
		return &Array{[]Object{i}}
	}

	return &Null{}
}
