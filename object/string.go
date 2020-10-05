package object

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

// String String builtin type
type String struct {
	Value string
}

// Type type string
func (s *String) Type() Type {
	return STRING
}

// Inspect return value as string
func (s *String) Inspect() string {
	return fmt.Sprintf("%s", s.Value)
}

// HashKey implements the HashMap interface
// for comparing whether values are the same or not.
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// Equals returns if two objects are equal
func (s *String) Equals(o Object) bool {
	// type cast left object if the two operands are not the same type
	if s.Type() != o.Type() {
		o = o.Cast(s.Type())
	}

	return s.HashKey() == o.(Hashable).HashKey()
}

// Cast type cast
func (s *String) Cast(t Type) Object {

	switch t {
	case STRING:
		return s
	case INTEGER:
		v, _ := strconv.ParseInt(s.Value, 10, 64)
		return &Integer{v}
	case BOOLEAN:
		if s.Value == "0" || s.Value == "" {
			return &Boolean{false}
		}
		return &Boolean{true}
	case FLOAT:
		v, _ := strconv.ParseFloat(s.Value, 64)
		return &Float{v}
	case ARRAY:
		return &Array{[]Object{s}}
	}

	return &Null{}
}
