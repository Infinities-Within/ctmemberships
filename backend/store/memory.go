package store

import "github.com/infinities-within/ctmemberships/backend/model"

type Entity interface {
	ID() string
	Version() int
}

type MemoryStore struct {
	users       map[string]Entity
	members     map[string]Entity
	memberships map[string]Entity
}

type UnsupportedEntityError struct {
}

func (e UnsupportedEntityError) Error() string {
	return "supplied Entity is not supported by this store"
}

func (store *MemoryStore) Save(e Entity) error {
	switch e.(type) {
	case model.User:
		store.users[e.ID()] = e
	case model.Member:
		store.members[e.ID()] = e
	case model.Membership:
		store.memberships[e.ID()] = e
	default:
		return UnsupportedEntityError{}
	}

	return nil
}
