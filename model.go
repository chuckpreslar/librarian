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
  return false
}

func (self *Model) IsModified() bool {
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
