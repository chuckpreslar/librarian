#librarain

An ORM for Go.  As of right now, this is a simple prototype and is under heavy development.  

## Installation

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/chuckpreslar/librarian

## Usage

An example of how to create a basic model:

```go
import (
  "github.com/chuckpreslar/librarian"
)

type User struct {
  *librarian.Model
  Id        int    `db:"id"`
  FirstName string `db:"first_name"`
  LastName  string `db:"last_name"`
  Email     string `db:"email"`
}

// Now define your `Table` that interacts with this model.

var Users = librarian.Table{
  Name: "users", // The name of the table in your database.
  Model: User{}, // The model that should be initialized/populated when interacting with the table.
}
```


## Documentation

View godoc or visit [godoc.org](http://godoc.org/github.com/chuckpreslar/librarian).

    $ godoc librarian

## License

> The MIT License (MIT)

> Copyright (c) 2013 Chuck Preslar

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
