package librarian

import (
  "fmt"
  "reflect"
  "strings"
)

type Model struct {
  table      Table
  definition interface{}
}

func (self *Model) Save() error {
  var (
    table = self.table.Name
    model = reflect.TypeOf(self.definition)
  )

  fmt.Println(table)

  if reflect.Ptr == model.Kind() {
    model = model.Elem()
  }

  var numFields = model.NumField()

  for i := 0; i < numFields; i++ {
    var (
      field = model.Field(i)
      name  = field.Name
    )

    if "model" != strings.ToLower(name) {
      fmt.Println(name)
    }
  }

  return nil
}

func (self *Model) Destroy() error {
  return nil
}
