package librarian

import (
  "database/sql"
  _ "github.com/lib/pq"
)

type Connection struct {
  engine     string
  options    string
  session    *sql.DB
}

var connection *Connection

func Connect(engine string, options string) (*Connection, error){
  if nil != connection {
    if err := connection.session.Close(); nil != err {
      return nil, err
    }
  }

  database, err := sql.Open(engine, options)

  if nil != err {
    return nil, err
  }

  connection = &Connection{engine, options, database}

  return connection, nil
}