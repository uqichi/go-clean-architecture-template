package api

import (
	"fmt"
	"net/http"

	"github.com/uqichi/goec/models"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/uqichi/goec/api/handler"
	"github.com/uqichi/goec/infra/cache"
	"github.com/uqichi/goec/infra/rdb"
	"github.com/uqichi/goec/usecase"
)

const port = 1323

func Serve() {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Debug = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	conn, err := rdb.NewConnection("development")
	if err != nil {
		log.Panic(err)
	}

	productRepository := rdb.NewProductDB(conn)
	_ = cache.NewSessionCache()

	productUseCase := usecase.NewProductUseCase(productRepository)

	productHandler := handler.NewProductHandler(productUseCase)

	{
		product := e.Group("/products")
		product.GET("/:id", productHandler.Get)
		product.GET("", productHandler.List)
		product.POST("", productHandler.Create)
		product.PUT("/:id", productHandler.Update)
		product.DELETE("/:id", productHandler.Delete)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

// customHTTPErrorHandler is the custom HTTP error handler. It sends a JSON response
// with status code.
func customHTTPErrorHandler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
		if he.Internal != nil {
			err = fmt.Errorf("%v, %v", err, he.Internal)
		}
	} else if ae, ok := err.(*models.Error); ok {
		// Handle application error to convert to HTTP status code
		switch ae.Code {
		case models.ErrNotFound:
			code = http.StatusNotFound
		case models.ErrInvalid:
			code = http.StatusBadRequest
		case models.ErrDuplicate:
			code = http.StatusConflict
		}
		msg = ae
	} else if c.Echo().Debug {
		msg = err.Error()
	} else {
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = echo.Map{"message": msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead { // Issue #608
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			c.Echo().Logger.Error(err)
		}
	}
}
