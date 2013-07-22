package librarian

import (
  "fmt"
)

type Model struct {
  table      Table
  definition interface{}
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
