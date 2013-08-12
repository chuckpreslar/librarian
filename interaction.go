// Package librarian provides an ORM.
package librarian

// Standard library imports.
import (
  "errors"
)

// Package imports.
import (
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/tree/managers"
  "github.com/chuckpreslar/codex/tree/nodes"
)

// Table errors.
var (
  ErrNoPrimaryKey = errors.New("Table is missing a primary key.")
)

var (
  Accessors   = make(map[Table]managers.Accessor)   // Cached codex Accessors for Tables.
  Relations   = make(map[Table]*nodes.RelationNode) // Cached codex RelationNodes for Tables.
  PrimaryKeys = make(map[Table]interface{})         // Cached primary keys for Tables.
)

// AccessorFor looks up or returns a cached codex Accessor for a Table.
func AccessorFor(table Table) managers.Accessor {
  if accessor, ok := Accessors[table]; ok {
    return accessor
  }

  Accessors[table] = codex.Table(table.Name)
  return Accessors[table]
}

// RelationFor looks up or returns a cached codex RelationNode for a Table.
func RelationFor(table Table) *nodes.RelationNode {
  if relation, ok := Relations[table]; ok {
    return relation
  }

  Relations[table] = AccessorFor(table).Relation()
  return Relations[table]
}

// RelationFor looks up or returns a cached primary key for a Table, or an error if
// one was not found.
func PrimaryKeyFor(table Table) (interface{}, error) {
  if key, ok := PrimaryKeys[table]; ok {
    return key, nil
  }

  key, err := Cartographer.ColumnForField(table.Model, table.PrimaryKey)

  if nil == err {
    PrimaryKeys[table] = key
  }

  return key, err
}

// GenerateBindingsFor generates an slice of codex BindingNodes for the given elements.
func GenerateBindingsFor(elements ...interface{}) (bindings []interface{}) {
  for i := 0; i < len(elements); i++ {
    bindings = append(bindings, Binding)
  }

  return
}

// ModifiedColumnsAndValuesFor returns a slice of columns and a slice of values
// which have been modified on the Model, or an error. Note, columns[i] = values[i].
func ModifiedColumnsAndValuesFor(model *Model) (columns []interface{}, values []interface{}, err error) {
  modified, err := Cartographer.ModifiedColumnsValuesMapFor(model.values, model.definition)

  if nil != err {
    return
  }

  for column, value := range modified {
    columns = append(columns, column)
    values = append(values, value)
  }

  return
}

// PreformModelTransaction attempts to begin a transaction intending to update a single
// Model struct, returning an error if it fails and calls Rollback on the sql Tx.
func PreformModelTransaction(model *Model, query string, arguments ...interface{}) (err error) {
  transaction, err := Connection.driver.Begin()
  if nil != err {
    return
  }

  statement, err := transaction.Prepare(query)

  if nil != err {
    return
  }

  defer statement.Close()
  defer WrtieToLog(query)

  rows, err := statement.Query(arguments...)

  if nil != err {
    transaction.Rollback()
    return
  }

  if err = Cartographer.Sync(rows, model.definition); nil != err {
    transaction.Rollback()
    return
  } else if err = transaction.Commit(); nil != err {
    return
  }

  model.flags = PERSISTED
  model.values, err = Cartographer.FieldValueMapFor(model.definition)

  return
}

// InsertSingleRecord attempts to store a newly created Model into
// a record in the Connection drivers database.
func InsertSingleRecord(model *Model) (err error) {
  columns, values, err := ModifiedColumnsAndValuesFor(model)

  if nil != err {
    return
  }

  manager := managers.Insertion(RelationFor(model.table)).
    Insert(GenerateBindingsFor(values...)...).
    Into(columns...).
    SetAdapter(Connection.adapter)

  if key, err := PrimaryKeyFor(model.table); nil == err {
    manager.Returning(key)
  } else {
    return ErrNoPrimaryKey
  }

  sql, err := manager.ToSql()

  if nil != err {
    return
  }

  return PreformModelTransaction(model, sql, values...)
}

// UpdateSingleRecord attempts to update an existing record
// to the Models current values based on the Models Tables
// PrimaryKey field.
func UpdateSingleRecord(model *Model) (err error) {
  columns, values, err := ModifiedColumnsAndValuesFor(model)

  if nil != err {
    return
  }

  manager := managers.Modification(RelationFor(model.table)).
    Set(columns...).
    To(GenerateBindingsFor(values...)...).
    SetAdapter(Connection.adapter)

  if 0 < len(model.table.PrimaryKey) {
    column, err := PrimaryKeyFor(model.table)

    if nil != err {
      return err
    }

    field, err := Cartographer.FieldForColumn(model.definition, column)

    if nil != err {
      return err
    }

    manager.Where(AccessorFor(model.table)(column).Eq(model.values[field]))

  } else {
    return ErrNoPrimaryKey
  }

  sql, err := manager.ToSql()

  if nil != err {
    return
  }

  return PreformModelTransaction(model, sql, values...)
}
