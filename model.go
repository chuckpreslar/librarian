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

type Flag uint32

const (
  PERSISTED Flag = iota
  NEW
  DIRTY
)

var (
  ErrModificationCheck = errors.New("Failed to determine if the model has been modified.")
)

type Model struct {
  table      Table
  definition ModelInterface
  values     map[interface{}]interface{}
  flags      Flag
}

func (self *Model) IsNew() bool {
  return 0 < (NEW & self.flags)
}

func (self *Model) IsModified() (bool, error) {
  if 0 < (DIRTY & self.flags) {
    return true, nil
  }

  values, err := Cartographer.FieldValueMapFor(self.definition)

  if nil != err {
    return false, ErrModificationCheck
  }

  for key, value := range self.values {
    if values[key] != value {
      self.flags = self.flags | DIRTY
      return true, nil
    }
  }

  return false, nil
}

func (self *Model) IsValid() bool {
  return false
}

func (self *Model) Save() error {
  if modified, err := self.IsModified(); !modified || nil != err {
    return err
  }

  if self.IsNew() {
    return InsertSingleRecord(self)
  }

  return UpdateSingleRecord(self)
}

func (self *Model) Destroy() error {
  return nil
}
