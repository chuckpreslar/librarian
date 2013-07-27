package librarian

import (
  "github.com/chuckpreslar/cartographer"
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/tree/managers"
)

var CARTOGRAPHER = cartographer.Initialize("db")

type Relation struct {
  Table    Table
  Mananger *managers.SelectManager
  Accessor managers.Accessor
}

func (self *Relation) Select(columns ...string) *Relation {
  for _, column := range columns {
    column, err := CARTOGRAPHER.ColumnForField(self.Table.Model, column)

    if nil != err {
      panic(err)
    }

    self.Mananger.Project(self.Accessor(column))
  }

  return self
}

func (self *Relation) Where(conditions ...interface{}) *Relation {
  if 0 == len(conditions) {
    return self
  }

  switch condition := conditions[0]; condition.(type) {
  case string:
    self.Mananger.Where(parseStringBinding(condition.(string), conditions[1:]...))
  default:
    panic("Unable to parse Where conditions supplied.")
  }

  return self
}

func (self *Relation) Distinct() *Relation {
  return self
}

func (self *Relation) Unique() *Relation {
  return self
}

func (self *Relation) Order(orderings ...string) *Relation {
  return self
}

func (self *Relation) Group() *Relation {
  return self
}

func (self *Relation) Having() *Relation {
  return self
}

func (self *Relation) Limit(limit int) *Relation {
  return self
}

func (self *Relation) Offset(offset int) *Relation {
  return self
}

func (self *Relation) Lock() *Relation {
  return self
}

func (self *Relation) First() (interface{}, error) {
  self.Mananger.Limit(1)

  if 0 < len(self.Table.PrimaryKey) {
    column, err := CARTOGRAPHER.ColumnForField(self.Table.Model, self.Table.PrimaryKey)

    if nil != err {
      panic(err)
    }

    self.Mananger.Order(self.Accessor(column).Asc())
  }

  return self.Mananger.ToSql()
}

func (self *Relation) Last() (interface{}, error) {
  self.Mananger.Limit(1)

  if 0 < len(self.Table.PrimaryKey) {
    column, err := CARTOGRAPHER.ColumnForField(self.Table.Model, self.Table.PrimaryKey)

    if nil != err {
      panic(err)
    }

    self.Mananger.Order(self.Accessor(column).Desc())
  }

  return self.Mananger.ToSql()
}

func (self *Relation) All() ([]interface{}, error) {
  return nil, nil
}

func InitializeRelation(table Table) (relation *Relation) {
  relation = new(Relation)
  relation.Table = table
  relation.Accessor = codex.Table(table.Name)
  relation.Mananger = managers.Selection(relation.Accessor.Relation())

  return
}
