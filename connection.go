// Package librarian provides a RDMS agnostic ORM.
package librarian

import (
  "database/sql"
  _ "encoding/json"
  "fmt"
  _ "io/ioutil"
  "os"
  "os/signal"
)

import (
  _ "github.com/lib/pq" // PostgreSQL driver.
)

type DatabaseEngine string

const (
  Postgres DatabaseEngine = "postgres"
)

var database struct {
  name    string
  options string
  engine  string
  handle  *sql.DB
}

func EstablishConnection(engine, name, options string) (err error) {
  if database.handle, err = sql.Open(engine, DataSourceFor(engine, name, options)); nil != err {
    return
  }

  database.engine = engine
  database.name = name
  database.options = options

  sigint := make(chan os.Signal, 1)
  signal.Notify(sigint, os.Interrupt)

  go func() {
    for s := range sigint {
      if nil != s {
        database.handle.Close()
      }
    }
  }()

  return
}

func PingDatabase() error {
  return database.handle.Ping()
}

func DataSourceFor(engine, name, options string) string {
  return fmt.Sprintf("dbname=%s %s", name, options)
}
