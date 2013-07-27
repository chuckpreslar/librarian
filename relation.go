package librarian

import (
  "github.com/chuckpreslar/codex/tree/managers"
)

type Relation struct {
  Table    Table
  Mananger *managers.SelectManager
  Accessor managers.Accessor
}

func (self *Relation) Select() *Relation {
  return self
}

func (self *Relation) Where() *Relation {
  return self
}

func (self *Relation) Distinct() *Relation {
  return self
}

func (self *Relation) Unique() *Relation {
  return self
}

func (self *Relation) Order() *Relation {
  return self
}

func (self *Relation) Group() *Relation {
  return self
}

func (self *Relation) Having() *Relation {
  return self
}

func (self *Relation) Limit() *Relation {
  return self
}

func (self *Relation) Offset() *Relation {
  return self
}

func (self *Relation) Lock() *Relation {
  return self
}

func (self *Relation) First() ([]interface{}, error) {
  return nil, nil
}

func (self *Relation) Last() ([]interface{}, error) {
  return nil, nil
}

func (self *Relation) All() ([]interface{}, error) {
  return nil, nil
}
