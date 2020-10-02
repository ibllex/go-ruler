package object

const (
	NULL = "NULL"

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

// HashTable base type
type HashTable interface {
	HashKey() HashKey
}

// Object base type
type Object interface {
	Type() Type
	Inspect() string
	Cast(Type) Object
}

// IsNull returns an object is Null or not
func IsNull(o Object) bool {
	_, ok := o.(*Null)
	return ok
}

// IsEqual returns if two objects are equal
func IsEqual(l Object, r Object) bool {
	if IsNull(l) || IsNull(r) {
		return IsNull(l) && IsNull(r)
	}

	// type cast left object if the two operands are not the same type
	if l.Type() != r.Type() {
		l = l.Cast(r.Type())
	}

	return l.(HashTable).HashKey() == r.(HashTable).HashKey()
}
