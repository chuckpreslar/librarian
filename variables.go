package librarian

import (
  "github.com/chuckpreslar/cartographer"
)

var Database *DatabaseConnection
var Cartographer = cartographer.Initialize("db")
