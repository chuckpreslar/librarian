// Package librarian provides a RDMS agnostic ORM.
package librarian

// File `migration.go` provides types and functionality
// to assist in generation migrations to alter tables
// stored in a RDMS.

import (
  "fmt"
)

import (
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/sql"
)

// SQL Constraint constants from codex's sql package.
const (
  NOT_NULL    = sql.NOT_NULL
  UNIQUE      = sql.UNIQUE
  PRIMARY_KEY = sql.PRIMARY_KEY
  FOREIGN_KEY = sql.FOREIGN_KEY
  CHECK       = sql.CHECK
  DEFAULT     = sql.DEFAULT
)

var constaintToString = map[sql.Constraint]string{
  NOT_NULL:    "not_null",
  UNIQUE:      "unique",
  PRIMARY_KEY: "pkey",
  FOREIGN_KEY: "fkey",
  CHECK:       "check",
  DEFAULT:     "default",
}

const (
  indexNameFormatter = "%v_%v_%v"
)

// SQL Type constants from codex's sql package.
const (
  STRING    = sql.STRING
  TEXT      = sql.TEXT
  BOOLEAN   = sql.BOOLEAN
  INTEGER   = sql.INTEGER
  FLOAT     = sql.FLOAT
  DECIMAL   = sql.DECIMAL
  DATE      = sql.DATE
  TIME      = sql.TIME
  DATETIME  = sql.DATETIME
  TIMESTAMP = sql.TIMESTAMP
)

type MigrationRunner func(*Migrator)

func (m MigrationRunner) Run(migrator *Migrator) {
  m(migrator)
}

type Migration struct {
  Up   MigrationRunner
  Down MigrationRunner
}

type Modifier interface {
  ToSql() (string, error)
}

type Migrator struct {
  modifiers []Modifier // Slice containing the migrations TableModifiers.
}

type Alteration struct {
  manager *managers.AlterManager
}

func (a *Alteration) AddColumn(c string, t sql.Type) *Alteration {
  a.manager.AddColumn(c, t)
  return a
}

func (a *Alteration) ChangeColumn(c string, t sql.Type) *Alteration {
  a.manager.ModifyColumn(c, t)
  return a
}

func (a *Alteration) RenameColumn() *Alteration {
  return a
}

func (a *Alteration) RemoveColumn(c string) *Alteration {
  a.manager.RemoveColumn(c)
  return a
}

func (a *Alteration) AddIndex(c string, k sql.Constraint, o ...interface{}) *Alteration {
  a.manager.AddConstraint(c, k, o...)
  return a
}

func (a *Alteration) RemoveIndexByColumn() *Alteration {
  return a
}

func (a *Alteration) RemoveIndexByName(i string) *Alteration {
  a.manager.RemoveIndex(i)
  return a
}

func (a *Alteration) ToSql() (string, error) {
  return a.manager.ToSql()
}

func NewAlteration(table string) (alteration *Alteration) {
  relation := codex.Table(table).Relation()
  alteration = new(Alteration)
  alteration.manager = managers.Alteration(relation)
  return
}
