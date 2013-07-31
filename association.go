package librarian

type Association interface {
  ToRelation() *Relation
}

type HasMany struct {
  Table      Table
  Through    Table
  ForiegnKey string
}
