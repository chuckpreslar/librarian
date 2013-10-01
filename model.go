// Package librarian provides a RDMS agnostic ORM.
package librarian

type ModelInterface interface {
  Save() error
}

type Model struct {
  table     Table
  reference ModelInterface
}

func (m Model) Save() error { return nil }
