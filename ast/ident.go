package ast

// Ident target value
type Ident struct {
	path []string
}

// NewIdent construct ident node
func NewIdent(path []string) Ident {
	return Ident{
		path: path,
	}
}
