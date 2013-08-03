#librarian

An ORM for Go. As of right now, this is a simple prototype and is under heavy development.

## Dependencies

[Codex](http://www.github.com/chuckpreslar/codex)
[Cartographer](http://www.github.com/chuckpreslar/cartographer)

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
  Name: "users",   // The name of the table in your database.
  Model: User{},   // The model that should be initialized/populated when interacting with the table.
  PrimaryKey: "id" // The primary key for the table.
}
```

## Basics

### Querying existing records.

Retrieving records stored within your table is simple.

```go
// Continuing from out `Usage` example above.
users, err := Users.All()

if nil != err {
  // Handle your error.
}

for _, user := range users {
  fmt.Println(user.(*User).FirstName)
}
```

To select only certain columns from a table, you may pass their names to your table's Select method. You may pass their names as how they are represented on your struct, or their column represented by the field's `db` tag.

```go
users, err := Users.Select("id", "first_name").All()
```

To filter the results returned, use the Where method. Currently, the Where method only accepts strings, replacing all binding charaters (`?`) with the additional arguments supplied.  Columns listed within the formatter string must be as they appear in the database.

```go
users, err := Users.Where("first_name = ? AND email LIKE ?", "Jon", "%@example.com").All()
```



### Creating and Persisting Records

To create a new records, there's one small gotcha' that you may have seen coming - you must type assert your models when you initially create them.

```go
user := Users.New().(*User)
user.FirstName = "Jon"
user.LastName = "Doe"
user.Email = "jon@example.com"

// user.IsNew() => true

if err := user.Save(); nil != err {
  // Time for some error handling.
}
```

If you have assigned a value to your Table's `PrimaryKey` field, the field will be populated on a successful save.  Once your record has been persisted, it is just as easily updated with another call to `Save`.


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
