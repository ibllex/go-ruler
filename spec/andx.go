package spec

// AndX create and composite specification
func AndX(specs []Specification) *Composite {
	return &Composite{
		operator:       "and",
		specifications: specs,
	}
}
