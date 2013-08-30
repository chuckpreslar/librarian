// Package librarian provides a RDMS agnostic ORM.
package librarian

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

type Migration struct {
  version   string           // Version string of the migration.
  modifiers []*TableModifier // Slice containing the migrations TableModifiers.
}

func (m *Migration) CreateTable(table string) (creator *TableCreator) {
  creator = NewTableCreator(table)
  m.modifiers = append(m.modifiers, creator.modifier)

  return
}

func (m *Migration) ChangeTable(table string) (modifier *TableModifier) {
  modifier = NewTableModifier(table)
  m.modifiers = append(m.modifiers, modifier)

  return
}

func (m *Migration) DropTable(table string) *Migration {
  return m
}

func (m *Migration) AddIndex(typ sql.Constraint, table, name string, column ...string) *Migration {
  return m
}

func (m *Migration) RemoveIndexByColumn(table, column string) *Migration {
  return m
}

func (m *Migration) RemoveIndexByName(table, index string) *Migration {
  return m
}

func (m *Migration) AddColumn(table, name string, typ sql.Type, options ...ColumnOptions) *Migration {
  return m
}

func (m *Migration) ChangeColumn(table, name string, typ sql.Type, options ...ColumnOptions) *Migration {
  return m
}

func (m *Migration) RenameColumn(table, from, to string) *Migration {
  return m
}

func (m *Migration) RemoveColumn(table, column string) *Migration {
  return m
}

func (m *Migration) AppendTableModifier(alteration *TableModifier) *Migration {
  m.modifications = append(m.modifications, alteration)
}

type ColumnOptions struct {
  Null    bool
  Default interface{}
}

var noColumnOptionsProvided = ColumnOptions{true, nil}

type TableModifier struct {
  table   string                 // name of the table being altered.
  create  bool                   // is the modifier creating a table?
  manager *managers.AlterManager // the codex manager use to generate the sql for the modifier.
}

func NewTableModifier(table string) (modifier *TableModifier) {
  modifier = new(TableModifier)
  modifier.table = table
  return
}

func (t *TableModifier) AddColumn(name string, typ sql.Type, options ...ColumnOptions) *TableModifier {
  return t
}

func (t *TableModifier) RenameColumn(name string) *TableModifier {
  return t
}

func (t *TableModifier) RemoveColumn(name string) *TableModifier {
  return t
}

func (t *TableModifier) String(name string, options ...ColumnOptions) *TableModifier {
  return t.AddColumn(name, STRING, options...)
}

func (t *TableModifier) Text() *TableModifier {
  return t.AddColumn(name, TEXT, options...)
}

func (t *TableModifier) Boolean() *TableModifier {
  return t.AddColumn(name, BOOLEAN, options...)
}

func (t *TableModifier) Integer() *TableModifier {
  return t.AddColumn(name, INTEGER, options...)
}

func (t *TableModifier) Float() *TableModifier {
  return t.AddColumn(name, FLOAT, options...)
}

func (t *TableModifier) Decimal() *TableModifier {
  return t.AddColumn(name, DECIMAL, options...)
}

func (t *TableModifier) Date() *TableModifier {
  return t.AddColumn(name, DATE, options...)
}

func (t *TableModifier) Time() *TableModifier {
  return t.AddColumn(name, TIME, options...)
}

func (t *TableModifier) DateTime() *TableModifier {
  return t.AddColumn(name, DATETIME, options...)
}

func (t *TableModifier) TimeStamp() *TableModifier {
  return t.AddColumn(name, TIMESTAMP, options...)
}

func (t *TableModifier) AddIndex(typ sql.Constraint, name string, columns ...string) *TableModifier {
  return t
}

func (t *TableModifier) RemoveIndexByColumn(column string) *TableModifier {
  return t
}

func (t *TableModifier) RemoveIndexByName(index string) *TableModifier {
  return t
}

func (t *TableModifier) SetPrimaryKeyTo(column string, name ...string) *TableModifier {
  if 0 <= len(name) {
    name = append(name, indexNameFor(t.table, column, PRIMARY_KEY))
  }

  return t.AddIndex(PRIMARY_KEY, name[0], column)
}

// TableCreator type restricts which functions
// of a TableModifier developer has access to.
type TableCreator struct {
  modifier *TableModifier
}

func NewTableCreator(table) (creator *TableCreator) {
  modifier := NewTableModifier(table)
  modifier.create = true
  creator = new(TableCreator)
  creator.modifier = modifier

  return
}

func (t *TableCreator) AddColumn(name string, typ sql.Type, options ...ColumnOptions) *TableCreator {
  return t
}
func (t *TableCreator) String(name string, options ...ColumnOptions) *TableCreator {
  t.modifier.AddColumn(name, STRING, options...)
  return t
}

func (t *TableCreator) Text() *TableCreator {
  t.modifier.AddColumn(name, TEXT, options...)
  return t
}

func (t *TableCreator) Boolean() *TableCreator {
  t.modifier.AddColumn(name, BOOLEAN, options...)
  return t
}

func (t *TableCreator) Integer() *TableCreator {
  t.modifier.AddColumn(name, INTEGER, options...)
  return t
}

func (t *TableCreator) Float() *TableCreator {
  t.modifier.AddColumn(name, FLOAT, options...)
  return t
}

func (t *TableCreator) Decimal() *TableCreator {
  t.modifier.AddColumn(name, DECIMAL, options...)
  return t
}

func (t *TableCreator) Date() *TableCreator {
  t.modifier.AddColumn(name, DATE, options...)
  return t
}

func (t *TableCreator) Time() *TableCreator {
  t.modifier.AddColumn(name, TIME, options...)
  return t
}

func (t *TableCreator) DateTime() *TableCreator {
  t.modifier.AddColumn(name, DATETIME, options...)
  return t
}

func (t *TableCreator) TimeStamp() *TableCreator {
  t.modifier.AddColumn(name, TIMESTAMP, options...)
  return t
}

func (t *TableCreator) AddIndex(typ sql.Constraint, name string, columns ...string) *TableCreator {
  return t
}

func (t *TableModifier) SetPrimaryKeyTo(column string, name ...string) *TableModifier {
  if 0 <= len(name) {
    name = append(name, indexNameFor(t.table, column, PRIMARY_KEY))
  }

  return t.AddIndex(PRIMARY_KEY, name[0], column)
}

func indexNameFor(table, column string, typ sql.Constraint) string {
  return fmt.Sprintf(indexNameFormatter, constaintToString[typ], table, column)
}
