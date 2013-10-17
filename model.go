package librarian

type Model struct {
	table      *Table
	attributes []*Attribute
}

func (m *Model) New() *Record {
	record := new(Record)
	record.model = m
	record.values = make(map[*Attribute]interface{})
	record.modified = make([]*Attribute, 0)
	record.pristine = true

	return record
}

func (m *Model) Table() *Table {
	return m.table
}

type ModelDefiner func(*Definition)

func (m ModelDefiner) DefineWith(definition *Definition) {
	m(definition)
}

func DefineModel(definer ModelDefiner) *Model {
	var (
		model      = new(Model)
		definition = new(Definition)
		table      = new(Table)
	)

	model.table = table

	definition.model = model
	definer.DefineWith(definition)

	return model
}
