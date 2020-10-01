package ast

// Param param value
type Param struct {
	id Ident
}

// NewParam construct Param node
func NewParam(id Ident) Param {
	return Param{
		id: id,
	}
}
