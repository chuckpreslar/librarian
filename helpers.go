package librarian

import (
  "fmt"
  "strings"
)

const BIND = `?`

func parseStringBinding(formatter string, bindings ...interface{}) string {
  for _, binding := range bindings {
    formatter = strings.Replace(formatter, BIND, tagBindingVariable(binding), 1)
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