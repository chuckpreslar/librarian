package librarian

type Type uint8

const (
	Integer Type = iota
	String
	Float
	Datetime
	Timestamp
)
