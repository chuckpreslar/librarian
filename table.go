// Package librarian provides an ORM.
package librarian

import (
  "errors"
)

var (
  ErrNewFailed = errors.New("Failed to create new replica of model.")
)

type Table struct {
  Name  string         // Name of the table the relation connects to.
  Model ModelInterface // Model calling Table's New method generates.
}
