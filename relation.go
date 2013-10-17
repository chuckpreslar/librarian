package librarian

import (
	"database/sql"
	"fmt"
)

import (
	"github.com/chuckpreslar/codex"
	"github.com/chuckpreslar/codex/managers"
)

type Relation struct {
	model    *Model
	manager  *managers.SelectManager
	accessor managers.Accessor
}

func (r *Relation) First() (*Record, error) {
	query, err := r.manager.Order(r.accessor(r.model.table.key).Asc()).Limit(1).ToSql()

	if nil != err {
		return nil, err
	}

	return r.one(query)
}

func (r *Relation) Last() (*Record, error) {
	query, err := r.manager.Order(r.accessor(r.model.table.key).Desc()).Limit(1).ToSql()

	if nil != err {
		return nil, err
	}

	return r.one(query)
}

func (r *Relation) Find(key interface{}) (*Record, error) {
	r.manager.Where(r.accessor(r.model.table.key).Eq(key))
	return r.First()
}

func (r *Relation) one(query string) (*Record, error) {
	results, err := r.search(query)

	if nil != err {
		return nil, err
	} else if count := len(results); 1 != count {
		return nil, fmt.Errorf("expected query to return only 1 result, returned %d", count)
	}

	return results[0], nil
}

func (r *Relation) search(query string) ([]*Record, error) {
	var (
		statement *sql.Stmt
		rows      *sql.Rows
		err       error
		columns   []string
		records   []*Record
	)

	if statement, err = Librarian.handle.Prepare(query); nil != err {
		return nil, err
	} else if rows, err = statement.Query(); nil != err {
		return nil, err
	} else if columns, err = rows.Columns(); nil != err {
		return nil, err
	}

	for rows.Next() {
		ptr := make([]interface{}, len(columns))

		for i := 0; i < len(ptr); i++ {
			var buffer interface{}
			ptr[i] = &buffer
		}

		if err = rows.Scan(ptr...); nil != err {
			return nil, err
		} else {
			record := r.model.New()
			for i := 0; i < len(ptr); i++ {
				record.Set(columns[i], (*ptr[i].(*interface{})))
			}

			records = append(records, record)
		}
	}

	return records, nil
}

func NewRelation(model *Model) *Relation {
	relation := new(Relation)
	relation.model = model
	relation.accessor = codex.Table(model.table.name)
	relation.manager = managers.Selection(relation.accessor.Relation())

	return relation
}
