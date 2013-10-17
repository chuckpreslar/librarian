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

func (m *Model) First() (*Record, error) {
	return NewRelation(m).First()
}

func (m *Model) Last() (*Record, error) {
	return NewRelation(m).Last()
}

func (m *Model) Find(key interface{}) (*Record, error) {
	return NewRelation(m).Find(key)
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

	Librarian.models = append(Librarian.models, model)

	return model
}
