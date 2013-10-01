package test

import (
  "testing"
)

import (
  "github.com/chuckpreslar/librarian"
)

type M struct { 
  *librarian.Model
}

func TestTableNew(t *testing.T) {
  table := librarian.Table{
    Model: M{},
  }

  replica := table.New()
  
  if _, ok := replica.(*M); !ok {
    t.Errorf("Table failed to create replica of model, returned type %T, exected %T\n.", replica, &M{})
  }
}
