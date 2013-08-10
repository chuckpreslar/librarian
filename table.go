// Package librarian provides an ORM.
package librarian

import (
  "errors"
)

var (
  ErrNewFailed = errors.New("Failed to create new replica of model.")
)

type Table struct {
  Name       string         // Name of the table the relation connects to.
  PrimaryKey string         // Database column that is the models primary key.
  Model      ModelInterface // Model calling Table's New method generates.
}

// Table errors.
var (
  ErrNoPrimaryKey = errors.New("Table has no primary key to use with Find.")
)

func (self Table) New() (ModelInterface, error) {
  replica, err := Cartographer.CreateReplica(self.Model, CreateModelReplicaFor(self, NEW_MODEL))

  if nil != err {
    err = ErrNewFailed
    return nil, err
  }

  return replica.Interface().(ModelInterface), nil
}

func (self Table) DestroyAll() error {
  return nil
}

func (self Table) Select(columns ...string) *Relation {
  return nil
}

func (self Table) Where(conditions ...interface{}) *Relation {
  return nil

}

func (self Table) Distinct() *Relation {
  return nil
}

func (self Table) Unique() *Relation {
  return nil
}

func (self Table) Order(orderings ...string) *Relation {
  return nil
}

func (self Table) Group() *Relation {
  return nil
}

func (self Table) Having() *Relation {
  return nil
}

func (self Table) Limit(limit int) *Relation {
  return nil
}

func (self Table) Offset(offset int) *Relation {
  return nil
}

func (self Table) Lock() *Relation {
  return nil
}

func (self Table) Find(key interface{}) (interface{}, error) {
  return nil, nil
}

func (self Table) First() (interface{}, error) {
  return nil, nil
}

func (self Table) Last() (interface{}, error) {
  return nil, nil
}

func (self Table) All() ([]interface{}, error) {
  return nil, nil
}
