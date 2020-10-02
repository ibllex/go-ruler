package ast

// Target target value
type Target struct {
	ID *Ident
}

func (t *Target) Literal() string {
	return t.ID.Literal()
}

// NewTarget construct target node
func NewTarget(id *Ident) *Target {
	return &Target{
		ID: id,
	}
}
