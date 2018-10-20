package repository

import (
	"context"

	"github.com/gobuffalo/uuid"
)

type SessionRepository interface {
	Exists(ctx context.Context, userID uuid.UUID) (bool, error)
}
