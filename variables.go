package librarian

import (
  "github.com/chuckpreslar/cartographer"
  "github.com/chuckpreslar/codex/tree/nodes"
)

var Database *DatabaseConnection
var Binding = nodes.Binding()
var Cartographer = cartographer.Initialize("db")
