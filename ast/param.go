package ast

// Param param value
type Param struct {
	id *Ident
}

func (p *Param) Literal() string {
	return p.id.Literal()
}

// NewParam construct Param node
func NewParam(id *Ident) *Param {
	return &Param{
		id: id,
	}
}
