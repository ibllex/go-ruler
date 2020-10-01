package ast

// Target target value
type Target struct {
	id Ident
}

// NewTarget construct target node
func NewTarget(id Ident) Target {
	return Target{
		id: id,
	}
}
