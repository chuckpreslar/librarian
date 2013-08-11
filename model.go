// Package librarian provides an ORM.
package librarian

import (
  "errors"
)

type ModelInterface interface {
  IsNew() bool
  IsModified() (bool, error)
  IsValid() bool
  Save() error
  Destroy() error
}

const (
  PERSISTED_MODEL = 0x0001
  NEW_MODEL       = 0x0010
  DIRTY_MODEL     = 0x0100
)

var (
  ErrModificationCheck = errors.New("Failed to determine if the model has been modified.")
)

type Model struct {
  table      Table
  definition ModelInterface
  values     map[interface{}]interface{}
  flags      uint16
}

func (self *Model) IsNew() bool {
  return 0 < (NEW_MODEL & self.flags)
}

func (self *Model) IsModified() (bool, error) {
  if 0 < (DIRTY_MODEL & self.flags) {
    return true, nil
  }

  values, err := Cartographer.FieldValueMapFor(self.definition)

  if nil != err {
    return false, ErrModificationCheck
  }

  for key, value := range self.values {
    if values[key] != value {
      self.flags = self.flags | DIRTY_MODEL
      return true, nil
    }
  }

  return false, nil
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
