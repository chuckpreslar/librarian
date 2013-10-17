package librarian

import (
	"github.com/chuckpreslar/codex"
	"github.com/chuckpreslar/codex/managers"
)

type Table struct {
	name, key string
}

func (t *Table) Name() string {
	return t.name
}

func (t *Table) PrimaryKey() string {
	return t.key
}

func (t *Table) CodexAccessor() managers.Accessor {
	return codex.Table(t.name)
}
