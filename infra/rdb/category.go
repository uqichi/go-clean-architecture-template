package rdb

import (
	"context"

	"github.com/gobuffalo/pop"

	"github.com/gobuffalo/uuid"
	"github.com/uqichi/goec/models"
)

type CategoryDB struct {
	conn *pop.Connection
}

func NewCategoryDB(conn *pop.Connection) *CategoryDB {
	return &CategoryDB{conn}
}

func (r *CategoryDB) FindByID(ctx context.Context, tenantID, id uuid.UUID) (*models.Category, error) {
	cat := models.Category{}

	err := r.conn.Where("tenant_id = ?", tenantID).Find(&cat, id)
	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (r *CategoryDB) Create(ctx context.Context, cat *models.Category) (*models.Category, error) {
	verr, err := r.conn.ValidateAndCreate(cat)
	if verr != nil {
		return nil, err // TODO: handle as invalid
	}
	if err != nil {
		return nil, err
	}

	return cat, err
}
