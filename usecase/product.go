package usecase

import (
	"context"

	"github.com/gobuffalo/uuid"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	"github.com/uqichi/go-pay-grpc/proto"
	"github.com/uqichi/goec/models"
	"github.com/uqichi/goec/usecase/repository"
	"google.golang.org/grpc"
)

type ProductUseCase struct {
	productRepository repository.ProductRepository
}

func NewProductUseCase(
	productRepository repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
	}
}

func (uc *ProductUseCase) Get(ctx context.Context, tenantID, id uuid.UUID) (*models.Product, error) {
	prod, err := uc.productRepository.FindByID(ctx, tenantID, id)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func (uc *ProductUseCase) List(ctx context.Context, tenantID uuid.UUID) (*models.Products, error) {
	prods, err := uc.productRepository.FindAll(ctx, tenantID)
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func (uc *ProductUseCase) Create(ctx context.Context, prod *models.Product) (*models.Product, error) {
	prod, err := uc.productRepository.Create(ctx, prod)
	if err != nil {
		return nil, err
	}
	return prod, err
}

func (uc *ProductUseCase) Update(ctx context.Context, prod *models.Product) (*models.Product, error) {
	ent, err := uc.productRepository.FindByID(ctx, prod.TenantID, prod.ID)
	if err != nil {
		return nil, err
	}

	// set values to update
	ent.Name = prod.Name

	prod, err = uc.productRepository.Update(ctx, ent)
	if err != nil {
		return nil, err
	}
	return prod, err
}

func (uc *ProductUseCase) Delete(ctx context.Context, tenantID, id uuid.UUID) error {
	return uc.productRepository.Delete(ctx, tenantID, id)
}

type PurchaseProductInput struct {
	Token    string
	Products models.Products
}

func (uc *ProductUseCase) PurchaseProducts(ctx context.Context, input *PurchaseProductInput) error {
	var totalAmount uint
	for _, v := range input.Products {
		totalAmount += v.Amount
	}

	req := &pb.ChargeRequest{
		Id:          "hogehoge", // TODO: pass multiple product's ids
		Token:       input.Token,
		Amount:      int32(totalAmount),
		Name:        "fugafuga", // TODO: name properly
		Description: "piyopiyo", // TODO: describe properly
	}

	log.Infof("send charge request: %+v", req)

	const addr = "localhost:50051" // TODO: move

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		println(err)
		return err
	}
	defer conn.Close()
	client := pb.NewPayServiceClient(conn)

	res, err := client.Charge(ctx, req)
	if err != nil {
		return err
	}
	if !res.Paid {
		return errors.New("not paid")
	}

	return nil
}
