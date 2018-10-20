package cache

import (
	"github.com/gobuffalo/uuid"
)

type SessionCache struct{}

func NewSessionCache() *SessionCache {
	return &SessionCache{}
}

func (r *SessionCache) Exists(userID uuid.UUID) (bool, error) {
	// TODO: impl using like redis, memcache or any
	return true, nil
}
