package test

import (
  "testing"
)

import (
  "github.com/chuckpreslar/librarian"
)

type Test struct { 
  *librarian.Model
  Key int
  Value string
}

var TestTable = librarian.Table{
  Name: "test",
  Model: Test{},
  PrimaryKey: "Key",
}

func TestTableNew(t *testing.T) {
  table := librarian.Table{
    Model: Test{},
  }

  replica := table.New()

  if _, ok := replica.(*Test); !ok {
    t.Errorf("Table failed to create replica of model, returned type %T, exected %T\n.", replica, &Test{})
  }
}

func TestTableFind(t *testing.T) {
  if result, err := TestTable.Find(1); nil != err {
    t.Errorf("Failed to find record in database, error `%s` was returned.\n", err)
  } else if result.(*Test).Key != 1 {
    t.Errorf("Failed to find record in appropriate record, recieved `%v`.\n", result.(*Test))
  }
}
