package tasks

import (
  "fmt"
  "os"
  "path"
  "strings"
)

import (
  "github.com/chuckpreslar/gofer"
)

const (
  DATABASE_DIRECTORY = "database"
  MIGRATIONS_DIRECTORY = "migrations"
  SCHEMA_FILE = "schema.json"
  CONFIGURATION_FILE = "configuration.json"
)

const (
  CONFIGURATION_FILE_TEMPLATE = `
{
  "development": {
    "adapter": "",
    "database": "",
    "host": "",
    "post": ""
  },
  "staging": {
    "adapter": "",
    "database": "",
    "host": "",
    "post": ""
  },
  "production": {
    "adapter": "",
    "database": "",
    "host": "",
    "post": ""
  }
}
`
)

var directory string

var Initialize = gofer.Register(gofer.Task{
  Namespace: "librarian",
  Label: "init",
  Description: "Preforms initial setup to begin using the Librarian package.",
  Dependencies: []string{ "librarian:init:directories", "librarian:init:files" },
})

var InitializeDirectories = gofer.Register(gofer.Task{
  Namespace: "librarian:init",
  Label: "directories",
  Description: "Creates initial directories to to be used by the Librarian package.",
  Action: func() (err error) {
    fmt.Fprint(os.Stdout, "Creating initial directories.\n")

    directory, err = os.Getwd()

    if nil != err {
      return
    }

    err = os.MkdirAll(path.Join(directory, DATABASE_DIRECTORY, MIGRATIONS_DIRECTORY), os.ModePerm)

    return
  },
})

var InitializeFiles = gofer.Register(gofer.Task{
  Namespace: "librarian:init",
  Label: "files",
  Description: "Creates initial files to be used by the Librarian package.",
  Dependencies: []string{ "librarian:init:directories" },
  Action: func() (err error) {
    fmt.Fprint(os.Stdout, "Creating initial files.\n")

    schema, err := os.Create(path.Join(directory, DATABASE_DIRECTORY, SCHEMA_FILE))
    defer schema.Close()

    if nil != err {
      return
    }

    cfg, err := os.Create(path.Join(directory, DATABASE_DIRECTORY, CONFIGURATION_FILE))
    defer cfg.Close()

    if nil != err {
      return
    }

    _, err = fmt.Fprint(cfg, strings.TrimLeft(CONFIGURATION_FILE_TEMPLATE, "\n" ))
    return
  },
})