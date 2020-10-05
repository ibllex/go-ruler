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
