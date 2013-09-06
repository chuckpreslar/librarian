// Package librarian provides a RDMS agnostic ORM.
package librarian

// File `migration.go` provides types and functionality
// to assist in generation migrations to alter tables
// stored in a RDMS.

import (
  _ "fmt"
)

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
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

type Modifier interface {
  ToSql() (string, error)
}

type ColumnOption uint8

const (
  Null ColumnOption = iota
  Default
)

type ColumnOptions map[ColumnOption]interface{}

type Creation struct {
  relation *nodes.RelationNode
  manager  *managers.CreateManager
}

func (c *Creation) AddColumn(column string, typ sql.Type, options ...ColumnOptions) *Creation {
  c.manager.AddColumn(column, typ)

  if 0 < len(options) {
    if null, ok := options[0][Null]; ok {
      if null, ok := null.(bool); ok && !null {
        c.manager.AddConstraint(column, NOT_NULL)
        return c
      }
    }

    if def, ok := options[0][Default]; ok {
      c.manager.AddConstraint(column, DEFAULT, def)
    }
  }

  return c
}

func (c *Creation) String(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, STRING, options...)
}

func (c *Creation) Text(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, TEXT, options...)
}

func (c *Creation) Boolean(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, BOOLEAN, options...)
}

func (c *Creation) Integer(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, INTEGER, options...)
}

func (c *Creation) Float(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, FLOAT, options...)
}

func (c *Creation) Decimal(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, DECIMAL, options...)
}

func (c *Creation) Date(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, DATE, options...)
}

func (c *Creation) Time(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, TIME, options...)
}

func (c *Creation) Datetime(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, DATETIME, options...)
}

func (c *Creation) Timestamp(column string, options ...ColumnOptions) *Creation {
  return c.AddColumn(column, TIMESTAMP, options...)
}

func (c *Creation) AddIndex(name string, kind sql.Constraint, columns []string, options ...interface{}) *Creation {
  return c.manager.AddConstraint(column, kind, ...)
}

func (c *Creation) Unique(name string, columns ...string) *Creation {
  return c
}

func (c *Creation) PrimaryKey(name, column string) *Creation {
  return c
}

func (c *Creation) ForeignKey(name string, column string, reference string) *Creation {
  return c
}

func (c *Creation) ToSql() (string, error) {
  return c.manager.ToSql()
}

func NewCreation(table string) (creation *Creation) {
  creation = new(Creation)
  creation.relation = nodes.Relation(table)
  creation.manager = managers.Creation(creation.relation)
  return
}

type Alteration struct {
  relation *nodes.RelationNode
  manager  *managers.AlterManager
}

func (a *Alteration) ToSql() (string, error) {
  return a.manager.ToSql()
}

func NewAlteration(table string) (alteration *Alteration) {
  alteration = new(Alteration)
  alteration.relation = nodes.Relation(table)
  alteration.manager = managers.Alteration(alteration.relation)
  return
}

type Migrator struct {
  modifiers []Modifier
}

func (m *Migrator) CreateTable(table string) (creation *Creation) {
  creation = NewCreation(table)
  m.modifiers = append(m.modifiers, creation)
  return
}

func (m *Migrator) AlterTable(table string) (alteration *Alteration) {
  alteration = NewAlteration(table)
  m.modifiers = append(m.modifiers, alteration)
  return
}

type MigrationRunner func(*Migrator)

func (m MigrationRunner) Run(migrator *Migrator) {
  m(migrator)
}

type Migration struct {
  Up   MigrationRunner
  Down MigrationRunner
}
