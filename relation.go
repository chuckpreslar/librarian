package librarian

import (
  "errors"
  "reflect"
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/cartographer"
)

type Relation struct {
  table    Table
  accessor managers.Accessor
  query    *managers.SelectManager
}

func (self *Relation) First() (interface{}, error) {
  if nil == connection {
    return nil, errors.New("No connection has been established.")
  }

  sql, err := self.query.Limit(1).ToSql()

  if nil != err {
    return nil, err
  }

  results, err := connection.session.Query(sql)

  instance := cartographer.New()

  mapped, err := instance.Map(results, self.table.Model, func(element reflect.Value) {
    field := element.FieldByName("Model")
    if field.CanSet() {
      base := new(Model)
      
      // Attach the replicated model and self to base model struct.
      base.definition = element.Interface()
      base.table = self.table

      // Set value of embedded Model type to newly created base.
      field.Set(reflect.ValueOf(base))
    }
  })

  return mapped[0], err
}

func createRelation(table Table) (relation *Relation) {
  relation = new(Relation)
  relation.table = table
  relation.accessor = codex.Table(table.Name)
  relation.query = managers.Selection(relation.accessor.Relation())

  return
}