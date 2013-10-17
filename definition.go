package librarian

type Definition struct {
	model *Model
}

func (d *Definition) SetTableName(name string) *Definition {
	d.model.table.name = name
	return d
}

func (d *Definition) SetTablePrimaryKey(key string) *Definition {
	d.model.table.key = key
	return d
}

func (d *Definition) Attribute(name string, typ Type) *Attribute {
	attribute := new(Attribute)
	attribute.name = name
	attribute.typ = typ
	attribute.validators = make([]Validator, 0)
	attribute.validators = append(attribute.validators, TypeValidator{typ})

	d.model.attributes = append(d.model.attributes, attribute)

	return attribute
}
