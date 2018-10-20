package repository

import (
	"context"

	"github.com/gobuffalo/uuid"
	"github.com/uqichi/goec/models"
)

type ProductRepository interface {
	FindByID(ctx context.Context, tenantID, id uuid.UUID) (*models.Product, error)
	FindAll(ctx context.Context, tenantID uuid.UUID) (*models.Products, error)
	Create(ctx context.Context, prod *models.Product) (*models.Product, error)
	Update(ctx context.Context, prod *models.Product) (*models.Product, error)
	Delete(ctx context.Context, tenantID, id uuid.UUID) error
}
