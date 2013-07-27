package librarian

import (
  "github.com/chuckpreslar/cartographer"
  "reflect"
  "fmt"
  "strings"
)

const BINDING_CHARACTER = `?`

func parseStringBinding(formatter string, bindings ...interface{}) string {
  for _, binding := range bindings {
    formatter = strings.Replace(formatter, BINDING_CHARACTER, tagBindingVariable(binding), 1)
  }
  
  return formatter
}

func tagBindingVariable(binding interface{}) string {
  switch binding.(type) {
  case string, bool:
    return fmt.Sprintf("'%v'", binding)
  default:
    return fmt.Sprintf("%v", binding)
  }
}

func createModel(table Table, isNew bool) cartographer.Hook {
  return func(replica reflect.Value) (err error) {
    base := new(Model)
    embedded := replica.Elem().FieldByName("Model")

    base.definition = replica.Interface().(ModelInterface)
    base.table = table
    base.values, err = CARTOGRAPHER.FieldValueMapFor(replica.Interface())

    if isNew {
      base.isNew = true
    }

    embedded.Set(reflect.ValueOf(base))

    return
  }
}
