// Package librarian provides an ORM.
package librarian

// Standard library imports.
import (
  "database/sql"
  "errors"
  "os"
  "os/signal"
  "reflect"
  "runtime"
)

// SQL driver imports.
import (
  _ "github.com/lib/pq" // Postgres driver.
)

// Package imports.
import (
  "github.com/chuckpreslar/codex/nodes"
)

// Driver errors.
var (
  ErrConnectFailed     = errors.New("Failed to connect to the database.")
  ErrClosingConnection = errors.New("Failed to close connection to the database.")
)

// Database connection.
var Connection struct {
  adapter string  // The adapter type of the connection, eg. 'postgres', 'mysql'...
  driver  *sql.DB // The SQL driver.
}

// Exposed variables.
var (
  Binding = nodes.Binding()
)

// CodexRelationFor returns a pointer to a RelationNode from the codex package.
func CodexRelationFor(o interface{}) (relation *nodes.RelationNode) {
  switch o.(type) {
  case Table:
    relation = nodes.Relation(o.(Table).Name)
  case string:
    relation = nodes.Relation(o.(string))
  case *nodes.RelationNode:
  }

  return
}

type Hook func(reflect.Value) error

// Hook to use when creating new model replicas.
func CreateModelReplicaForHook(table Table, flag Flag) Hook {
  return func(replica reflect.Value) (err error) {
    base := new(Model)
    embedded := replica.Elem().FieldByName("Model")

    base.definition = replica.Interface().(ModelInterface)
    base.table = table
    base.flags = flag

    embedded.Set(reflect.ValueOf(base))
    return
  }
}

// Initialize attempts to open a connection to a database, using
// the driver specified by the `adapter` and calling the sql
// packages `Open` method with it and the `options`.
func Initialize(adapter, options string) (err error) {
  if Connection.driver, err = sql.Open(adapter, options); nil != err {
    return
  }

  Connection.adapter = adapter

  // Attempt to close connection on interrupt, or kill.
  go func() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, os.Kill)

    _ = <-c

    CloseConnection()
  }()

  // Before exit, attempt to close connection.
  runtime.SetFinalizer(Connection.driver, func(driver *sql.DB) {
    driver.Close()
  })

  return
}

// CloseConnection attempts to close the `Connection`'s opened
// driver connections.
func CloseConnection() (err error) {
  if nil == Connection.driver {
    return
  }

  if err = Connection.driver.Close(); nil != err {
    err = ErrClosingConnection
  }
  return
}
