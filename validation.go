package librarian

type Validator interface {
	Validate(value interface{}, attribute *Attribute, record *Record) bool
}

type ValidatorFunc func(value interface{}, attribute *Attribute, record *Record) bool

func (v ValidatorFunc) Validate(value interface{}, attribute *Attribute, record *Record) bool {
	return v(value, attribute, record)
}

type PresenceValidator struct{}

func (p PresenceValidator) Validate(value interface{}, attribute *Attribute, record *Record) bool {
	if value == nil {
		return false
	}

	return true
}

type TypeValidator struct {
	typ Type
}

func (t TypeValidator) Validate(value interface{}, attribute *Attribute, record *Record) bool {
	// If value is nil, allow PresenceValidator to catch invalid records.
	if nil == value {
		return true
	}

	switch value.(type) {
	case int, int8, int16, int32, int64:
		return t.typ == Integer
	case string:
		return t.typ == String
	}

	return false
}
