package model

type Membership struct {
  ID string
  Version int
  Name string
  Description string
}

type Entity interface {
  ID() string
  Version() int
}
