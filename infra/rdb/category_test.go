package rdb

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/uqichi/goec/models"

	"github.com/gobuffalo/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uqichi/goec/usecase/repository"
)

var categoryDB repository.CategoryRepository

func init() {
	conn, err := NewConnection("development")
	if err != nil {
		log.Panic(err)
	}
	categoryDB = NewCategoryDB(conn)
}

func TestCategoryDB_FindByID(t *testing.T) {
	assert := assert.New(t)

	res, err := categoryDB.FindByID(context.TODO(),
		uuid.FromStringOrNil("c9a14f2e-f414-459c-baa5-40f906d793d5"),
		uuid.FromStringOrNil("b552485c-c0a2-46e8-94fa-9b76e44e9e34"))
	assert.NoError(err)
	fmt.Printf("%+v", res)
}

func TestCategoryDB_Create(t *testing.T) {
	t.SkipNow()

	assert := assert.New(t)

	res, err := categoryDB.Create(context.TODO(), &models.Category{
		TenantID: uuid.Must(uuid.NewV4()),
		Name:     "cat1",
	})
	assert.NoError(err)
	fmt.Printf("%+v", res)
}
