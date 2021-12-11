package model

type entity struct {
	id      string
	version int
}

type User struct {
	entity
	Role         string
	PasswordHash string
	IsActive     bool
}

func (e User) ID() string {
	return e.id
}

func (e User) Version() int {
	return e.version
}

type Membership struct {
	entity
	Name           string
	Description    string
	Price          float32
	DurationInDays int
}

func (e Membership) ID() string {
	return e.id
}

func (e Membership) Version() int {
	return e.version
}

type TimeStamp int64

type Member struct {
	entity
	User        *User
	Membership  Membership
	RenewalDate TimeStamp
}

func (e Member) ID() string {
	return e.id
}

func (e Member) Version() int {
	return e.version
}
