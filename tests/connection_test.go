package test

import (
  "testing"
)

import (
  "github.com/chuckpreslar/librarian"
)

func TestEstablishConnection(t *testing.T) {
  err := librarian.EstablishConnection("postgres", "cheetah_development", "user=chuckpreslar sslmode=disable")

  if nil != err {
    t.Errorf("Failed to establish connection to database, error `%s` was returned.", err)
  }
}
