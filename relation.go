// Package librarian provides a RDMS agnostic ORM.
package librarian

import (
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/managers"
)

type Relation struct {
  managers.Accessor
  table Table
}

func (r *Relation) Find(key interface{}) (m ModelInterface, err error) {
  sql, err := r.Where(r.Accessor(r.table.PrimaryKey).Eq(key)).ToSql()

  if nil != err {
    return
  } else if err = PingDatabase(); nil != err {
    return
  }

  stmt, err := database.handle.Prepare(sql)

  if nil != err {
    return
  }

  rows, err := stmt.Query()

  if nil != err {
    return
  }

  return
}

func NewRelation(t Table) (r *Relation) {
  r = new(Relation)
  r.table = table
  r.Accessor = codex.Table(t.Name)
}
