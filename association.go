package librarian

type Association interface {
  ToRelation() *Relation
}

type HasMany struct {
  Table      Table
  Through    Table
  ForeignKey string
}

func (self HasMany) ToRelation() *Relation {
  return nil
}
