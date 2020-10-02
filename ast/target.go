package ast

// Target target value
type Target struct {
	id *Ident
}

func (t *Target) Literal() string {
	return t.id.Literal()
}

// NewTarget construct target node
func NewTarget(id *Ident) *Target {
	return &Target{
		id: id,
	}
}
