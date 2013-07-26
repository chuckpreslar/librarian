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
  table         Table
  definition    ModelInterface
  initialValues map[string]interface{}
  isNew         bool
}

func (self *Model) IsNew() bool {
  return self.isNew
}

func (self *Model) IsModified() bool {
  currentValues, err := mapper.FieldValueMapFor(self.definition)

  if nil != err {
    panic(err)
  }

  for key, value := range currentValues {
    if self.initialValues[key] != value {
      return true
    }
  }

  return false
}

func (self *Model) IsValid() bool {
  return false
}

func (self *Model) Save() error {
  return nil
}

func (self *Model) Destroy() error {
  return nil
}
