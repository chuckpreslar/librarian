package librarian

type Attribute struct {
	name       string
	typ        Type
	validators []Validator
}

func (a *Attribute) ValidatesWith(validator Validator) *Attribute {
	a.validators = append(a.validators, validator)
	return a
}
