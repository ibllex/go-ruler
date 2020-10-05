package object

// Null null builtin type
type Null struct {
}

// Type type string
func (n *Null) Type() Type {
	return NULL
}

// Inspect return value as string
func (n *Null) Inspect() string {
	return "null"
}

// Equals returns if two objects are equal
func (n *Null) Equals(o Object) bool {
	return IsNull(o)
}

// Cast type cast
func (n *Null) Cast(t Type) Object {

	switch t {
	case STRING:
		return &String{""}
	case INTEGER:
		return &Integer{0}
	case BOOLEAN:
		return &Boolean{false}
	case FLOAT:
		return &Float{0}
	case ARRAY:
		return &Array{}
	}

	return n
}
