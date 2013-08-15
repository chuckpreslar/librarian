// Package librarian provides an ORM.
package librarian

import (
  "errors"
)

import (
  "fmt"
  "strings"
)

import (
  "github.com/chuckpreslar/codex"
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

type Direction uint8

// Migration direction constants.
const (
  UP Direction = iota
  DOWN
)

// Expose Codex package type constants.
const (
  STRING    = codex.STRING
  TEXT      = codex.TEXT
  BOOLEAN   = codex.BOOLEAN
  INTEGER   = codex.INTEGER
  FLOAT     = codex.FLOAT
  DECIMAL   = codex.DECIMAL
  DATE      = codex.DATE
  TIME      = codex.TIME
  DATETIME  = codex.DATETIME
  TIMESTAMP = codex.TIMESTAMP
)

// Expose Codex package constraint constants.
const (
  NOT_NULL    = codex.NOT_NULL
  UNIQUE      = codex.UNIQUE
  PRIMARY_KEY = codex.PRIMARY_KEY
  FOREIGN_KEY = codex.FOREIGN_KEY
  CHECK       = codex.CHECK
  DEFAULT     = codex.DEFAULT
)

var (
  ErrNoDirection = errors.New("No action defined for direction.")
)

type IndexOptions struct {
  Name      string
  Reference interface{}
  Type      sql.Constraint
}

type ColumnOptions struct {
  Null    bool
  Default interface{}
}

type Migrator struct {
  alterations []*TableAlteration
}

type TableAlteration struct {
  relation  *nodes.RelationNode
  statement *managers.AlterManager
}

func (self *Migrator) CreateTable(name string) (alteration *TableAlteration) {
  alteration = new(TableAlteration)
  alteration.statement = codex.CreateTable(name)
  alteration.relation = nodes.Relation(name)

  self.alterations = append(self.alterations, alteration)

  return
}

func (self *TableAlteration) String(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.STRING, options...)
}

func (self *TableAlteration) Text(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.TEXT, options...)
}

func (self *TableAlteration) Boolean(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.BOOLEAN, options...)
}

func (self *TableAlteration) Integer(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.INTEGER, options...)
}

func (self *TableAlteration) Float(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.FLOAT, options...)
}

func (self *TableAlteration) Decimal(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.DECIMAL, options...)
}

func (self *TableAlteration) Date(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.DATE, options...)
}

func (self *TableAlteration) Time(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.TIME, options...)
}

func (self *TableAlteration) Datetime(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.DATETIME, options...)
}

func (self *TableAlteration) Timestamp(name string, options ...ColumnOptions) *TableAlteration {
  return self.AddColumn(name, codex.TIMESTAMP, options...)
}

func (self *TableAlteration) IncludeTimestamps() {}

func (self *TableAlteration) AddIndex(column string, options IndexOptions) *TableAlteration {
  switch options.Type {
  case UNIQUE:
    return self.AddUniqueIndex(column, options.Name)
  case PRIMARY_KEY:
    return self.AddPrimaryKeyIndex(column, options.Name)
  case FOREIGN_KEY:
    return self.AddForeignKeyIndex(column, options.Name, options.Reference)
  }

  return self
}

func (self *TableAlteration) AddUniqueIndex(column string, options ...string) *TableAlteration {
  var name string

  if 0 < len(options) {
    name = options[0]
  } else {
    name = fmt.Sprintf("%v_unique_%v", strings.ToLower(self.relation.Name), strings.ToLower(column))
  }

  self.statement.AddConstraint(column, UNIQUE, name)

  return self
}

func (self *TableAlteration) AddPrimaryKeyIndex(column string, options ...string) *TableAlteration {
  var name string

  if 0 < len(options) {
    name = options[0]
  } else {
    name = fmt.Sprintf("%v_pkey_%v", strings.ToLower(self.relation.Name), strings.ToLower(column))
  }

  self.statement.AddConstraint(column, PRIMARY_KEY, name)

  return self
}

func (self *TableAlteration) AddForeignKeyIndex(column string, options ...interface{}) *TableAlteration {
  var (
    relation *nodes.RelationNode
    name     string
  )

  switch len(options) {
  case 0:
    return self
  case 1:
    name = fmt.Sprintf("%v_fkey_%v", strings.ToLower(self.relation.Name), strings.ToLower(column))
    relation = CodexRelationFor(options[0])
  default:
    name = fmt.Sprintf("%v", options[0])
    relation = CodexRelationFor(options[1])
  }

  self.statement.AddConstraint(column, FOREIGN_KEY, name, relation)
  return self
}

func (self *TableAlteration) AddColumn(name string, typ sql.Type, options ...ColumnOptions) *TableAlteration {
  self.statement.AddColumn(name, typ)

  if 0 < len(options) {
    if false == options[0].Null {
      self.statement.AddConstraint(name, codex.NOT_NULL)
    }

    if nil != options[0].Default {
      self.statement.AddConstraint(name, codex.DEFAULT, options[0].Default)
    }
  }

  return self
}

func (self *TableAlteration) Sequalize() (string, error) {
  return self.statement.ToSql()
}

func (self *Migrator) Execute(sql string) error {
  result, err := Connection.driver.Exec(sql)

  if nil != err {
    return err
  }
}

func (self *Migrator) preform() error {
  for _, alteration := range self.alterations {
    sql, err := alteration.Sequalize()

    if nil != err {
      return err
    } else if err = self.Execute(sql); nil != err {
      return err
    }
  }

  return nil
}

type Migration struct {
  Up   func(*Migrator) error
  Down func(*Migrator) error
}

func (self Migration) run(direction Direction) (err error) {
  switch direction {
  case UP:
    err = self.up()
  case DOWN:
    err = self.down()
  default:
    return ErrNoDirection
  }

  return
}

func (self Migration) up() error {
  if nil == self.Up {
    return ErrNoDirection
  }

  migrator := new(Migrator)

  if err := self.Up(migrator); nil != err {
    return err
  }

  return migrator.preform()
}

func (self Migration) down() error {
  if nil == self.Down {
    return ErrNoDirection
  }

  migrator := new(Migrator)

  if err := self.Down(migrator); nil != err {
    return err
  }

  return migrator.preform()
}
