package rdb

import (
	"context"

	"github.com/pkg/errors"

	"github.com/go-sql-driver/mysql"

	"github.com/gobuffalo/pop"

	"github.com/gobuffalo/uuid"
	"github.com/uqichi/goec/models"
)

type ProductDB struct {
	conn *pop.Connection
}

func NewProductDB(conn *pop.Connection) *ProductDB {
	return &ProductDB{conn}
}

func (r *ProductDB) FindByID(ctx context.Context, tenantID, id uuid.UUID) (*models.Product, error) {
	prod := models.Product{}

	err := r.conn.Eager().Where("tenant_id = ?", tenantID).Find(&prod, id)
	if err != nil {
		if err.Error() == ErrNotFound.Error() {
			return nil, models.NewError(models.ErrNotFound)
		}
		return nil, err
	}

	return &prod, nil
}

func (r *ProductDB) FindAll(ctx context.Context, tenantID uuid.UUID) (*models.Products, error) {
	prods := models.Products{}

	err := r.conn.Eager().Where("tenant_id = ?", tenantID).All(&prods)
	if err != nil {
		return nil, err
	}

	return &prods, nil
}

func (r *ProductDB) Create(ctx context.Context, prod *models.Product) (*models.Product, error) {
	verr, err := r.conn.ValidateAndCreate(prod)
	if verr.HasAny() {
		return nil, models.NewError(models.ErrInvalid, verr)
	}
	if err != nil {
		if me, ok := errors.Cause(err).(*mysql.MySQLError); ok && me.Number == 1062 {
			return nil, models.NewError(models.ErrDuplicate, err)
		}
		return nil, err
	}

	_ = r.conn.Load(prod)

	return prod, err
}

func (r *ProductDB) Update(ctx context.Context, prod *models.Product) (*models.Product, error) {
	verr, err := r.conn.ValidateAndUpdate(prod, "tenant_id", "code")
	if verr.HasAny() {
		return nil, models.NewError(models.ErrInvalid, verr)
	}
	if err != nil {
		return nil, err
	}

	return prod, err
}

func (r *ProductDB) Delete(ctx context.Context, tenantID, id uuid.UUID) error {
	return r.conn.Destroy(&models.Product{
		ID:       id,
		TenantID: tenantID,
	})
}
