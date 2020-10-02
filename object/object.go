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
