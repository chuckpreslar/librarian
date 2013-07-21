package librarian

import (
  "github.com/chuckpreslar/librarian/interfaces"
  "reflect"
)

type Table struct {
  Name  string                    // Name of the table the relation connects to.
  Model interfaces.ModelInterface // Model calling Table's New method generates.
}

func (self Table) New() interfaces.ModelInterface {
  var typ = reflect.TypeOf(self.Model)

  if reflect.Ptr == typ.Kind() {
    typ = typ.Elem()
  }

  var (
    replica  = reflect.New(typ)
    element  = replica.Elem()
    embedded = element.FieldByName("Model")
    base     = new(Model)
  )

  // Attach the replicated model and self to base model struct.
  base.definition = replica.Interface()
  base.table = self

  // Set value of embedded Model type to newly created base.
  embedded.Set(reflect.ValueOf(base))

  return replica.Interface().(interfaces.ModelInterface)
}
