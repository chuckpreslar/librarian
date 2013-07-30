package librarian

import (
  "database/sql"
  _ "github.com/lib/pq"
)

type DatabaseConnection struct {
  engine     string
  options    string
  connection *sql.DB
}

func (database *DatabaseConnection) Prepare(sql string) (*sql.Stmt, error) {
  return database.connection.Prepare(sql)
}

func (database *DatabaseConnection) Begin(sql string) (*sql.Tx, error) {
  return database.connection.Begin()
}

func Connect(engine string, options string) (*DatabaseConnection, error) {
  if nil != Database {
    if err := Database.connection.Close(); nil != err {
      return nil, err
    }
  }

  connection, err := sql.Open(engine, options)

  if nil != err {
    return nil, err
  }

  Database = &DatabaseConnection{engine, options, connection}

  return Database, nil
}

func Disconnect() {
  if nil == Database {
    return
  }

  Database.connection.Close()
}
