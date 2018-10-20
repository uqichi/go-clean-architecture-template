package rdb

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/gobuffalo/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uqichi/goec/models"
	"github.com/uqichi/goec/usecase/repository"
)

var productDB repository.ProductRepository

func init() {
	conn, err := NewConnection("development")
	if err != nil {
		log.Panic(err)
	}
	productDB = NewProductDB(conn)
}

func TestProductDB_FindByID(t *testing.T) {
	assert := assert.New(t)

	res, err := productDB.FindByID(context.TODO(),
		uuid.FromStringOrNil("b0e9ace2-b10d-4200-b6b6-f55642238d32"),
		uuid.FromStringOrNil("424bce30-ff2a-4364-afa3-d9fcf607dbd0"))
	assert.NoError(err)
	fmt.Printf("%+v", res)
}

func TestProductDB_FindAll(t *testing.T) {
	assert := assert.New(t)

	res, err := productDB.FindAll(context.TODO(),
		uuid.FromStringOrNil("b0e9ace2-b10d-4200-b6b6-f55642238d32"))
	assert.NoError(err)
	fmt.Printf("%+v", res)
}

func TestProductDB_Create(t *testing.T) {
	t.SkipNow()

	assert := assert.New(t)

	res, err := productDB.Create(context.TODO(), &models.Product{
		TenantID:   uuid.Must(uuid.NewV4()),
		CategoryID: uuid.Must(uuid.NewV4()),
		Name:       "hoge",
		Code:       "xyzxyzxyz",
		StockNum:   99,
		Amount:     18900,
	})
	assert.NoError(err)
	fmt.Printf("%+v", res)
}

func TestProductDB_Update(t *testing.T) {
	t.SkipNow()

	assert := assert.New(t)

	found, _ := productDB.FindByID(context.TODO(),
		uuid.FromStringOrNil("b0e9ace2-b10d-4200-b6b6-f55642238d32"),
		uuid.FromStringOrNil("424bce30-ff2a-4364-afa3-d9fcf607dbd0"))

	found.Name = "fugafuga"
	found.Code = "iiiiiiiii"
	res, err := productDB.Update(context.TODO(), found)
	assert.NoError(err)
	fmt.Printf("%+v", res)
}

func TestProductDB_Delete(t *testing.T) {
	t.SkipNow()

	assert := assert.New(t)

	err := productDB.Delete(context.TODO(),
		uuid.FromStringOrNil("b0e9ace2-b10d-4200-b6b6-f55642238d32"),
		uuid.FromStringOrNil("fd1b5820-edc1-4b35-a662-bd9f7025b699"))
	assert.NoError(err)
}
