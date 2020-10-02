package ast

// Param param value
type Param struct {
	ID *Ident
}

func (p *Param) Literal() string {
	return p.ID.Literal()
}

// NewParam construct Param node
func NewParam(id *Ident) *Param {
	return &Param{
		ID: id,
	}
}
