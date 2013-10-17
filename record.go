package librarian

import (
	"fmt"
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

	r.values[attribute] = value

	for i := 0; i < len(r.modified); i++ {
		if attribute == r.modified[i] {
			return r
		}
	}

	r.modified = append(r.modified, attribute)
	return r
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
