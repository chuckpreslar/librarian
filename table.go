package librarian

import (
  "github.com/chuckpreslar/cartographer"
  "reflect"
)

type Table struct {
  Name       string         // Name of the table the relation connects to.
  PrimaryKey string         // Database column that is the models primary key.
  Model      ModelInterface // Model calling Table's New method generates.
}

// createModel returns a cartographer.Hook that embeds a Model type into
// the Table's Model field struct, seeting the Model's fields appropriately
// based on the `isNew` bool that is passed.
func (self Table) createModel(isNew bool) cartographer.Hook {
  return func(replica reflect.Value) (err error) {
    base := new(Model)
    embedded := replica.Elem().FieldByName("Model")

    base.definition = replica.Interface().(ModelInterface)
    base.table = self
    base.initialValues, err = mapper.FieldValueMapFor(replica.Interface())

    if isNew {
      base.isNew = true
    }

    embedded.Set(reflect.ValueOf(base))

    return
  }
}

// New returns a pointer to a new instance of the Table's embedded Model field.
func (self Table) New() ModelInterface {
  var replica, err = mapper.CreateReplica(self.Model, self.createModel(true))

  if nil != err {
    panic(err)
  }

  return replica.Interface().(ModelInterface)
}

// All uses an established database connection to query for all table
// records, returning an array of pointers to the Table's embedded
// Model field or an error if one occurs.
func (self Table) All() (results []interface{}, err error) {
  return initializeQuery(self).All()
}

func (self Table) Select(columns ...string) *Query {
  return initializeQuery(self).Select(columns...)
}

func (self Table) Where(comparisons ...Comparison) *Query {
  return initializeQuery(self).Where(comparisons...)
}

// func (self Table) Save(m *Model) (err error) {
//   modified, err := mapper.ModifiedColumnsValuesMapFor(m.initialValues, m.definition)

//   if nil != err {
//     return
//   }

//   if m.IsNew() {
//     var columns, values []interface{}

//     for column, value := range modified {
//       columns = append(columns, column)
//       values = append(values, value)
//     }

//     relation := AccessorFor(self).Insert(values...).
//       Into(columns...)

//     if 0 < len(self.PrimaryKey) {
//       relation.Returning(self.PrimaryKey)
//     }

//     sql, _ := relation.ToSql()
//     stmt, _ := connection.session.Prepare(sql)
//     rows, _ := stmt.Query()
//     _ = mapper.Sync(rows, m.definition)
//   }

//   return
// }
