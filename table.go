package librarian

import (
  "errors"
  "fmt"
)

type Table struct {
  Name        string         // Name of the table the relation connects to.
  PrimaryKey  string         // Database column that is the models primary key.
  Model       ModelInterface // Model calling Table's New method generates.
  Reflections []ReflectionInterface
}

func (self Table) New() ModelInterface {
  var replica, err = Cartographer.CreateReplica(self.Model, createModel(self, true))

  if nil != err {
    panic(err)
  }

  return replica.Interface().(ModelInterface)
}

func (self Table) DestroyAll() error {
  return nil
}

func (self Table) Select(columns ...string) *Relation {
  return InitializeRelation(self).Select(columns...)
}

func (self Table) Where(conditions ...interface{}) *Relation {
  return InitializeRelation(self).Where(conditions...)

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
  return InitializeRelation(self).Limit(limit)
}

func (self Table) Offset(offset int) *Relation {
  return InitializeRelation(self).Offset(offset)
}

func (self Table) Lock() *Relation {
  return nil
}

func (self Table) Find(key interface{}) (interface{}, error) {
  if 0 == len(self.PrimaryKey) {
    return nil, errors.New(fmt.Sprintf("Table %s has no primary key to use with Find", self.Name))
  }

  return InitializeRelation(self).Find(key)
}

func (self Table) First() (interface{}, error) {
  return InitializeRelation(self).First()
}

func (self Table) Last() (interface{}, error) {
  return InitializeRelation(self).Last()
}

func (self Table) All() ([]interface{}, error) {
  return InitializeRelation(self).All()
}
