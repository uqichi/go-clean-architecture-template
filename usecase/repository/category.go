package repository

import (
	"context"

	"github.com/gobuffalo/uuid"
	"github.com/uqichi/goec/models"
)

type CategoryRepository interface {
	FindByID(ctx context.Context, tenantID, id uuid.UUID) (*models.Category, error)
	//FindAll(ctx context.Context, tenantID uuid.UUID) (*models.Categories, error)
	Create(ctx context.Context, cat *models.Category) (*models.Category, error)
	//Update(ctx context.Context, cat *models.Category) (*models.Category, error)
	//Delete(ctx context.Context, tenantID, id uuid.UUID) error
}
