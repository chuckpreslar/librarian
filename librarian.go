// Package librarian provides an ORM.
package librarian

// Standard library imports.
import (
  "database/sql"
  "errors"
  "fmt"
  "io"
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
  "github.com/chuckpreslar/cartographer"
  "github.com/chuckpreslar/codex/tree/nodes"
)

// Driver errors.
var (
  ErrConnectFailed     = errors.New("Failed to connect to the database.")
  ErrClosingConnection = errors.New("Failed to close connection to the database.")
)

// Logging setup.
var log io.Writer = os.Stdout

func WrtieToLog(line string) {
  fmt.Fprintf(log, "%s\n", line)
}

func SetLogWriter(writer io.Writer) {
  log = writer
}

// Database connection.
var Connection struct {
  adapter string  // The adapter type of the connection, eg. 'postgres', 'mysql'...
  driver  *sql.DB // The SQL driver.
}

// Exposed variables.
var (
  Cartographer = cartographer.Initialize("db")
  Binding      = nodes.Binding()
)

// Exposed functions.
func CreateModelReplicaForHook(table Table, flag uint16) cartographer.Hook {
  return func(replica reflect.Value) (err error) {
    base := new(Model)
    embedded := replica.Elem().FieldByName("Model")

    base.definition = replica.Interface().(ModelInterface)
    base.table = table
    base.values, err = Cartographer.FieldValueMapFor(replica.Interface())

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
