// Package librarian provides a RDMS agnostic ORM.
package librarian

import (
  "reflect"
)

type Table struct {
  Name       string
  PrimaryKey string
  Model      ModelInterface
}

func (t Table) New() (model ModelInterface) {
  value := reflect.New(reflect.TypeOf(t.Model))
  field := value.Elem().FieldByName("Model")

  base := new(Model)
  base.table = t
  base.reference = value.Interface().(ModelInterface)

  field.Set(reflect.ValueOf(base))

  model = value.Interface().(ModelInterface)
  return
}

func (t Table) Find(key interface{}) (ModelInterface, error) {
  return NewRelation(t).Find(key)
}

func (t Table) First() ModelInterface { return nil }
func (t Table) Last() ModelInterface  { return nil }

// func (t Table) Select() *Relation     { return nil }
// func (t Table) Where() *Relation      { return nil }
