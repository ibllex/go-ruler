package ast

// PositionalParam param by index
type PositionalParam struct {
	//
}

func (p *PositionalParam) Literal() string {
	return "?"
}

// NewPositionalParam construct Param node
func NewPositionalParam() *PositionalParam {
	return &PositionalParam{}
}
