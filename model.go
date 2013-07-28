package librarian

type ModelInterface interface {
  IsNew() bool
  IsModified() bool
  IsValid() bool
  Save() error
  Destroy() error
}

type ModelInterfaces []ModelInterface

type Model struct {
  table      Table
  definition ModelInterface
  values     map[string]interface{}
  isNew      bool
}

func (self *Model) IsNew() bool {
  return self.isNew
}

func (self *Model) IsModified() bool {
  return false
}

func (self *Model) IsValid() bool {
  return false
}

func (self *Model) Save() error {
  modified, err := CARTOGRAPHER.ModifiedColumnsValuesMapFor(self.values, self.definition)

  if nil != err || 0 == len(modified) {
    return err
  }

  var columns, values []interface{}

  for column, value := range modified {
    columns = append(columns, column)
    values = append(values, value)
  }

  if self.IsNew() {
    return Insert(values, columns, self)
  }

  return Update(values, columns, self)
}

func (self *Model) Destroy() error {
  return nil
}
