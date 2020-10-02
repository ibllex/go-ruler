package object

import (
	"fmt"
	"math"
	"strconv"
)

// Float float builtin type
type Float struct {
	Value float64
}

// Type type string
func (f *Float) Type() Type {
	return FLOAT
}

// Inspect return value as string
func (f *Float) Inspect() string {
	return fmt.Sprintf("%f", f.Value)
}

// HashKey implements the HashMap interface
// for comparing whether values are the same or not.
func (f *Float) HashKey() HashKey {
	return HashKey{Type: f.Type(), Value: math.Float64bits(f.Value)}
}

// Cast type cast
func (f *Float) Cast(t Type) Object {

	switch t {
	case STRING:
		return &String{strconv.FormatFloat(f.Value, 'g', -1, 64)}
	case INTEGER:
		return &Integer{int64(f.Value)}
	case BOOLEAN:
		return &Boolean{f.Value != 0}
	case FLOAT:
		return f
	}

	return &Null{}
}
