package librarian

import (
	"database/sql"
)

import (
	_ "github.com/lib/pq"
)

var Librarian struct {
	handle  *sql.DB
	adapter string
	options string
	models  []*Model
}

type Configuration struct{}

func (c *Configuration) SetAdapter(adapter string) *Configuration {
	Librarian.adapter = adapter
	return c
}

func (c *Configuration) SetOptions(options string) *Configuration {
	Librarian.options = options
	return c
}

type Configurator func(config *Configuration)

func (c Configurator) Configure(config *Configuration) {
	c(config)
}

func Configure(configurator Configurator) error {
	config := new(Configuration)
	configurator.Configure(config)

	var err error

	Librarian.handle, err = sql.Open(Librarian.adapter, Librarian.options)

	if nil != err {
		panic(err)
	}

	return err
}
