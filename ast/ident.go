package ast

// Ident target value
type Ident struct {
	Path []string
}

func (id *Ident) Literal() string {
	if len(id.Path) > 0 {
		return id.Path[0]
	}

	return ""
}

// NewIdent construct ident node
func NewIdent(path []string) *Ident {
	return &Ident{
		Path: path,
	}
}
