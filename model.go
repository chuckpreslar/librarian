// Package librarian provides an ORM.
package librarian

type ModelInterface interface {
  IsNew() bool               // Is the Model new?
  IsModified() (bool, error) // Has the Model been modified?
  Save() error
  Destroy() error
}

type Flag uint32

type Model struct {
  // Public exposed fields.
  Id int // The primary key all models will have.

  // Private struct fields used by the librarian package.
  table      Table                       // The Table the Model belongs to.
  definition ModelInterface              // The definition of the created model.
  values     map[interface{}]interface{} // Cached values contained by the fields of the definition struct.
  flags      Flag                        // Useful bitmask for flags.
}

func (self *Model) IsNew() bool {
  return false
}

func (self *Model) IsModified() (bool, error) {
  return false, nil
}

func (self *Model) Save() error {
  return nil
}

func (self *Model) Destroy() error {
  return nil
}
