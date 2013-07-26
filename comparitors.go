package librarian

type Comparison interface {
  comparison()
}

type Eq struct {
  Column string
  Value  interface{}
}

type Neq struct {
  Column string
  Value  interface{}
}

type Gt struct {
  Column string
  Value  interface{}
}

type Gte struct {
  Column string
  Value  interface{}
}

type Lt struct {
  Column string
  Value  interface{}
}

type Lte struct {
  Column string
  Value  interface{}
}

type Like struct {
  Column string
  Value  interface{}
}

type Unlike struct {
  Column string
  Value  interface{}
}

func (self Eq) comparison()     {}
func (self Neq) comparison()    {}
func (self Gt) comparison()     {}
func (self Gte) comparison()    {}
func (self Lt) comparison()     {}
func (self Lte) comparison()    {}
func (self Like) comparison()   {}
func (self Unlike) comparison() {}
