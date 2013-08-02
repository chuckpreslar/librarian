package librarian

import (
  "errors"
  "fmt"
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/tree/managers"
)

type Relation struct {
  Table    Table
  Mananger *managers.SelectManager
  Accessor managers.Accessor
}

func (self *Relation) Select(columns ...string) *Relation {
  for _, column := range columns {
    column, err := Cartographer.ColumnForField(self.Table.Model, column)

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

func (self *Relation) Order(orderings ...interface{}) *Relation {
  for _, ordering := range orderings {
    self.Mananger.Order(ordering)
  }

  return self
}

func (self *Relation) Group() *Relation {
  return self
}

func (self *Relation) Having() *Relation {
  return self
}

func (self *Relation) Limit(limit int) *Relation {
  self.Mananger.Limit(limit)
  return self
}

func (self *Relation) Offset(offset int) *Relation {
  self.Mananger.Offset(offset)
  return self
}

func (self *Relation) Lock() *Relation {
  return self
}

func (self *Relation) Find(key interface{}) (interface{}, error) {
  column, err := Cartographer.ColumnForField(self.Table.Model, self.Table.PrimaryKey)

  if nil != err {
    return nil, err
  }

  accessor := accessorFor(self.Table)
  self.Mananger.Where(accessor(column).Eq(key))

  return self.First()
}

func (self *Relation) First() (interface{}, error) {
  self.Mananger.Limit(1)

  if 0 < len(self.Table.PrimaryKey) {
    column, err := Cartographer.ColumnForField(self.Table.Model, self.Table.PrimaryKey)

    if nil != err {
      return nil, err
    }

    self.Mananger.Order(self.Accessor(column).Asc())
  }

  results, err := self.All()

  if nil != err {
    return nil, err
  } else if 0 >= len(results) {
    return nil, nil
  }

  return results[0], nil
}

func (self *Relation) Last() (interface{}, error) {
  self.Mananger.Limit(1)

  if 0 < len(self.Table.PrimaryKey) {
    column, err := Cartographer.ColumnForField(self.Table.Model, self.Table.PrimaryKey)

    if nil != err {
      return nil, err
    }

    self.Mananger.Order(self.Accessor(column).Desc())
  }

  results, err := self.All()

  if nil != err {
    return nil, err
  } else if 0 >= len(results) {
    return nil, nil
  }

  return results[0], nil
}

func (self *Relation) All() (results []interface{}, err error) {
  sql, err := self.Mananger.Engine(Database.engine).ToSql()

  if err != nil {
    return
  }

  statement, err := Database.Prepare(sql)

  if err != nil {
    return
  }

  defer statement.Close()

  rows, err := statement.Query()

  if err != nil {
    return
  }

  return Cartographer.Map(rows, self.Table.Model, createModel(self.Table, false))
}

func InitializeRelation(table Table) (relation *Relation) {
  relation = new(Relation)
  relation.Table = table
  relation.Accessor = codex.Table(table.Name)
  relation.Mananger = managers.Selection(relation.Accessor.Relation())

  return
}

// FIXME: Lots of duplicated code between insert and update, no bueno.
func Insert(values, columns []interface{}, model *Model) error {
  accessor := accessorFor(model.table)
  manager := managers.Insertion(accessor.Relation()).Insert(bindingsFor(values)...)

  for _, column := range columns {
    column, err := Cartographer.ColumnForField(model.definition, column.(string))

    if nil != err {
      return err
    }

    manager.Into(column)
  }

  if 0 < len(model.table.PrimaryKey) {
    column, err := Cartographer.ColumnForField(model.definition, model.table.PrimaryKey)

    if nil != err {
      return err
    }

    manager.Returning(column)
  }

  sql, err := manager.Engine(Database.engine).ToSql()

  if nil != err {
    return err
  }

  fmt.Println(sql)

  // FIXME: This should be a transaction.
  statement, err := Database.Prepare(sql)

  if nil != err {
    return err
  }

  defer statement.Close()

  rows, err := statement.Query(values...)

  if nil != err {
    return err
  }

  err = Cartographer.Sync(rows, model.definition)

  if nil != err {
    return err
  }

  model.isNew = false
  model.values, err = Cartographer.FieldValueMapFor(model.definition)

  return err
}

// FIXME: Lots of duplicated code between insert and update, no bueno.
func Update(values, columns []interface{}, model *Model) error {
  accessor := accessorFor(model.table)
  manager := managers.Modification(accessor.Relation())

  for _, column := range columns {
    column, err := Cartographer.ColumnForField(model.definition, column.(string))

    if nil != err {
      return err
    }

    manager.Set(column)
  }

  manager.To(values...)

  if 0 < len(model.table.PrimaryKey) {
    column, err := Cartographer.ColumnForField(model.definition, model.table.PrimaryKey)

    if nil != err {
      return err
    }

    field, err := Cartographer.FieldForColumn(model.definition, model.table.PrimaryKey)

    if nil != err {
      return err
    }

    manager.Where(accessor(column).Eq(model.values[field]))

  } else {
    return errors.New("Unable to update record missing value for primary key field.")
  }

  sql, err := manager.Engine(Database.engine).ToSql()

  if nil != err {
    return err
  }

  // FIXME: This should be a transaction.
  statement, err := Database.Prepare(sql)

  if nil != err {
    return err
  }

  defer statement.Close()

  _, err = statement.Exec()

  return err
}
