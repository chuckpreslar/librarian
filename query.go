package librarian

import (
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/cartographer"
  "github.com/chuckpreslar/codex/tree/managers"
  "fmt"
)

var mapper = cartographer.Initialize("db")

func AccessorFor(table Table) managers.Accessor {
  return codex.Table(table.Name)
}

type Query struct {
  table    Table
  accessor managers.Accessor
  manager  *managers.SelectManager
}

func (self *Query) Where(comparisons ...Comparison) *Query {
  for _, comparison := range comparisons {
    switch comparison.(type) {
    case Eq:
      self.manager.Where(self.accessor(comparison.(Eq).Column).Eq(comparison.(Eq).Value))
    case Neq:
      self.manager.Where(self.accessor(comparison.(Neq).Column).Neq(comparison.(Neq).Value))
    case Gt:
      self.manager.Where(self.accessor(comparison.(Gt).Column).Gt(comparison.(Gt).Value))
    case Gte:
      self.manager.Where(self.accessor(comparison.(Gte).Column).Gte(comparison.(Gte).Value))
    case Lt:
      self.manager.Where(self.accessor(comparison.(Lt).Column).Lt(comparison.(Lt).Value))
    case Lte:
      self.manager.Where(self.accessor(comparison.(Lte).Column).Lte(comparison.(Lte).Value))
    case Like:
      self.manager.Where(self.accessor(comparison.(Like).Column).Like(comparison.(Like).Value))
    case Unlike:
      self.manager.Where(self.accessor(comparison.(Unlike).Column).Unlike(comparison.(Unlike).Value))
    }
  }

  return self
}

func (self *Query) Select(columns ...string) *Query {
  for _, column := range columns {
    self.manager.Project(self.accessor(column))
  }

  return self
}

func (self *Query) All() (results []interface{}, err error) {
  sql, err := self.manager.ToSql()
  fmt.Println(sql)
  if err != nil {
    return
  }

  statement, err := connection.session.Prepare(sql)

  if err != nil {
    return
  }

  rows, err := statement.Query()

  if err != nil {
    return
  }

  results, err = mapper.Map(rows, self.table.Model, self.table.createModel(false))

  return
}

func initializeQuery(table Table) (query *Query) {
  query = new(Query)
  query.table = table
  query.accessor = AccessorFor(table)
  query.manager = managers.Selection(query.accessor.Relation())

  return
}