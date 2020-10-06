package object

const (
	NULL  = "NULL"
	ARRAY = "ARRAY"

	INTEGER = "INTEGER"
	FLOAT   = "FLOAT"
	BOOLEAN = "BOOLEAN"
	STRING  = "STRING"

	BUILTIN = "BUILTIN"
)

// Type object type name
type Type string

// HashKey object hash key
type HashKey struct {
	Type  Type
	Value uint64
}

// Hashable base type
type Hashable interface {
	HashKey() HashKey
}

// Object base type
type Object interface {
	Type() Type
	Inspect() string
	Cast(Type) Object
	Equals(Object) bool
}

// IsNull returns an object is Null or not
func IsNull(o Object) bool {
	_, ok := o.(*Null)
	return ok
}

// IsEqual returns if two objects are equal
func IsEqual(l Object, r Object) bool {
	return l.Equals(r)
}

// IsEmpty returns if an object is empty
func IsEmpty(o Object) bool {
	switch o := o.(type) {
	case *String:
		return o.Value == "" || o.Value == "0"
	case *Integer:
		return o.Value == 0
	case *Float:
		return o.Value == 0
	case *Array:
		return len(o.Elements) == 0
	case *Boolean:
		return o.Value == false
	}

	return true
}
