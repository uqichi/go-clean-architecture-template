package handler

import (
	"net/http"

	"github.com/uqichi/goec/models"

	"github.com/gobuffalo/uuid"
	"github.com/labstack/echo"
	"github.com/uqichi/goec/usecase"
)

type ProductHandler struct {
	productUseCase *usecase.ProductUseCase
}

func NewProductHandler(
	productUseCase *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

const tenantID = "b0e9ace2-b10d-4200-b6b6-f55642238d32" // TODO: get from session or sth

func (h *ProductHandler) Get(c echo.Context) error {
	productID := c.Param("id")
	res, err := h.productUseCase.Get(c.Request().Context(),
		uuid.FromStringOrNil(tenantID),
		uuid.FromStringOrNil(productID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) List(c echo.Context) error {
	res, err := h.productUseCase.List(c.Request().Context(),
		uuid.FromStringOrNil(tenantID))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) Create(c echo.Context) error {
	req := new(models.Product)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	req.TenantID = uuid.FromStringOrNil(tenantID)
	res, err := h.productUseCase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *ProductHandler) Update(c echo.Context) error {
	productID := c.Param("id")
	req := new(models.Product)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	req.ID = uuid.FromStringOrNil(productID)
	req.TenantID = uuid.FromStringOrNil(tenantID)
	res, err := h.productUseCase.Update(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) Delete(c echo.Context) error {
	productID := c.Param("id")
	err := h.productUseCase.Delete(c.Request().Context(),
		uuid.FromStringOrNil(tenantID),
		uuid.FromStringOrNil(productID))
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
