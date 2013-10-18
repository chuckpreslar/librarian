package librarian

import (
	"database/sql"
	"fmt"
	"strings"
)

import (
	"github.com/chuckpreslar/codex/managers"
)

type Relation struct {
	model      *Model
	manager    *managers.SelectManager
	accessor   managers.Accessor
	parameters []interface{}
}

func (r *Relation) First() (*Record, error) {
	query, err := r.manager.
		Order(r.accessor(r.model.table.key).Asc()).
		Limit(1).SetAdapter(Librarian.adapter).ToSql()
	if nil != err {
		return nil, err
	}

	return r.one(query)
}

func (r *Relation) Last() (*Record, error) {
	query, err := r.manager.
		Order(r.accessor(r.model.table.key).Desc()).
		Limit(1).SetAdapter(Librarian.adapter).ToSql()

	if nil != err {
		return nil, err
	}

	return r.one(query)
}

func (r *Relation) Find(key interface{}) (*Record, error) {
	r.manager.Where(r.accessor(r.model.table.key).Eq(key))
	return r.First()
}

func (r *Relation) Select(columns ...string) *Relation {
	for i := 0; i < len(columns); i++ {
		r.manager.Project(r.accessor(columns[i]))
	}

	return r
}

func (r *Relation) Where(formater string, parameters ...interface{}) *Relation {
	if "postgres" == Librarian.adapter {
		for count := len(r.parameters) + 1; 0 < strings.Index(formater, "?"); count++ {
			formater = strings.Replace(formater, "?", fmt.Sprintf("$%d", count), 1)
		}
	}

	r.manager.Where(formater)
	r.parameters = append(r.parameters, parameters...)
	return r
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
	} else if rows, err = statement.Query(r.parameters...); nil != err {
		return nil, err
	} else if columns, err = rows.Columns(); nil != err {
		return nil, err
	}

	for rows.Next() {
		buffer := make([]interface{}, len(columns))

		for i := 0; i < len(buffer); i++ {
			var item interface{}
			buffer[i] = &item
		}

		if err = rows.Scan(buffer...); nil != err {
			return nil, err
		} else {
			record := r.model.New()

			for i := 0; i < len(buffer); i++ {
				record.Set(columns[i], (*buffer[i].(*interface{})))
			}

			// Record was obtained from database,
			// set Record's pristine flag to false
			// and clear modified Attribute array.
			record.pristine = false
			record.modified = make([]*Attribute, 0)

			records = append(records, record)
		}
	}

	return records, nil
}

func NewRelation(model *Model) *Relation {
	relation := new(Relation)
	relation.model = model
	relation.accessor = model.Table().CodexAccessor()
	relation.manager = managers.Selection(relation.accessor.Relation())

	return relation
}
