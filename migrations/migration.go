package migrations

import (
  "github.com/chuckpreslar/codex/managers"
  "github.com/chuckpreslar/codex/nodes"
  "github.com/chuckpreslar/codex/sql"
)

// Expose codex package's SQL Constraint type constants
const (
  NotNull    = sql.NotNull
  Unique     = sql.Unique
  PrimaryKey = sql.PrimaryKey
  ForeignKey = sql.ForeignKey
  Check      = sql.Check
  Default    = sql.Default
)

// Expose codex package's SQL Type type constants
const (
  String    = sql.String
  Text      = sql.Text
  Boolean   = sql.Boolean
  Integer   = sql.Integer
  Float     = sql.Float
  Decimal   = sql.Decimal
  Date      = sql.Date
  Time      = sql.Time
  Datetime  = sql.Datetime
  Timestamp = sql.Timestamp
)

// Map of Option types to any value.
type Constraints map[sql.Constraint]interface{}
type Columns []interface{}

type Modifier interface {
  ToSql() (string, error)
}

type Creation struct {
  relation *nodes.RelationNode
  manager  *managers.CreateManager
}

func (c *Creation) AddColumn(column string, typ sql.Type, constraints ...Constraints) *Creation {
  c.manager.AddColumn(column, typ)
  return c
}

func (c *Creation) String(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, String, constraints...)
}

func (c *Creation) Text(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Text, constraints...)
}

func (c *Creation) Boolean(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Boolean, constraints...)
}

func (c *Creation) Integer(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Integer, constraints...)
}

func (c *Creation) Float(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Float, constraints...)
}

func (c *Creation) Decimal(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Decimal, constraints...)
}

func (c *Creation) Date(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Date, constraints...)
}

func (c *Creation) Time(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Time, constraints...)
}

func (c *Creation) Datetime(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Datetime, constraints...)
}

func (c *Creation) Timestamp(column string, constraints ...Constraints) *Creation {
  return c.AddColumn(column, Timestamp, constraints...)
}

func (c *Creation) AddIndex(columns Columns, kind sql.Constraint, name string, options ...interface{}) *Creation {
  c.manager.AddConstraint(columns, kind, append(options, name)...)
  return c
}

func (c *Creation) Unique(columns Columns, name ...interface{}) *Creation {
  return c
}

func (c *Creation) ForeignKey(columns Columns, reference string, name ...interface{}) *Creation {
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

func NewAlteration(table string) (alteration *Alteration) {
  alteration = new(Alteration)
  alteration.relation = nodes.Relation(table)
  alteration.manager = managers.Alteration(alteration.relation)
  return
}

func (a *Alteration) ToSql() (string, error) {
  return a.manager.ToSql()
}

type Director struct {
  modifiers []Modifier
}

func (d *Director) CreateTable(table string) (creation *Creation) {
  creation = NewCreation(table)
  d.modifiers = append(d.modifiers, creation)
  return
}

func (d *Director) AlterTable(table string) (alteration *Alteration) {
  alteration = NewAlteration(table)
  d.modifiers = append(d.modifiers, alteration)
  return
}

type Preformer func(*Director)

type Migration struct {
  Up   Preformer
  Down Preformer
}
