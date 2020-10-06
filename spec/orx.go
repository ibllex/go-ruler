package spec

// OrX create or composite specification
func OrX(specs []Specification) *Composite {
	return &Composite{
		operator:       "or",
		specifications: specs,
	}
}
