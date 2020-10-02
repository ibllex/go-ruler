package ast

// Ident target value
type Ident struct {
	path []string
}

func (id *Ident) Literal() string {
	if len(id.path) > 0 {
		return id.path[0]
	}

	return ""
}

// NewIdent construct ident node
func NewIdent(path []string) *Ident {
	return &Ident{
		path: path,
	}
}
