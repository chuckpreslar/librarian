package librarian

import (
	"database/sql"
	"fmt"
)

import (
	"github.com/chuckpreslar/codex/managers"
	"github.com/chuckpreslar/codex/nodes"
)

type Record struct {
	model    *Model
	values   map[*Attribute]interface{}
	modified []*Attribute
	pristine bool
}

func (r *Record) Get(name string) interface{} {
	for attr, value := range r.values {
		if attr.name == name {
			return value
		}
	}

	return nil
}

func (r *Record) Set(name string, value interface{}) *Record {
	var attribute *Attribute

	for i := 0; i < len(r.model.attributes); i++ {
		if r.model.attributes[i].name == name {
			attribute = r.model.attributes[i]
			break
		}
	}

	if attribute == nil {
		panic(fmt.Errorf("model for table `%s` has no attribute named `%s`", r.model.table.name, name))
	}

	r.values[attribute] = attribute.ConvertValue(value)

	for i := 0; i < len(r.modified); i++ {
		if attribute == r.modified[i] {
			return r
		}
	}

	r.modified = append(r.modified, attribute)
	return r
}

func (r *Record) Save() error {
	if r.pristine {
		return r.Insert()
	}

	return r.Update()
}

func (r *Record) Insert() error {
	accessor := r.model.table.CodexAccessor()
	manager := managers.Insertion(accessor.Relation())
	binding := nodes.Binding()

	var values []interface{}

	for i := 0; i < len(r.modified); i++ {
		manager.Insert(binding).Into(r.modified[i].name)
		values = append(values, r.values[r.modified[i]])
	}

	manager.Returning(r.model.table.key)

	query, err := manager.SetAdapter(Librarian.adapter).ToSql()

	if nil != err {
		return err
	}

	return r.modify(query, values...)
}

func (r *Record) Update() error {
	return nil
}

func (r *Record) modify(query string, values ...interface{}) error {
	var (
		transaction *sql.Tx
		statement   *sql.Stmt
		rows        *sql.Rows
		columns     []string
		err         error
	)

	if transaction, err = Librarian.handle.Begin(); nil != err {
		return err
	} else if statement, err = transaction.Prepare(query); nil != err {
		return err
	} else if rows, err = statement.Query(values...); nil != err {
		return err
	} else if columns, err = rows.Columns(); nil != err {
		return err
	}

	defer func() {
		if nil != err {
			transaction.Rollback()
		} else {
			transaction.Commit()
		}

		if nil != statement {
			statement.Close()
		}

		if nil != rows {
			rows.Close()
		}
	}()

	// Should be at max returning one row,
	// returns false when constraint errors
	// incountered.
	rows.Next()

	buffer := make([]interface{}, len(columns))

	for i := 0; i < len(buffer); i++ {
		var item interface{}
		buffer[i] = &item
	}

	if err = rows.Scan(buffer...); nil != err {
		return err
	} else {
		for i := 0; i < len(buffer); i++ {
			r.Set(columns[i], (*buffer[i].(*interface{})))
		}

		// Record was persisted to database,
		// set Record's pristine flag to false
		// and clear modified Attribute array.
		r.pristine = false
		r.modified = make([]*Attribute, 0)
	}

	return nil
}

func (r *Record) IsPrestine() bool {
	return r.pristine
}

func (r *Record) IsModified() bool {
	return 0 < len(r.modified)
}

func (r *Record) IsValid() bool {
	for _, attribute := range r.model.attributes {
		for _, validator := range attribute.validators {
			if !validator.Validate(r.values[attribute], attribute, r) {
				return false
			}
		}
	}

	return true
}
