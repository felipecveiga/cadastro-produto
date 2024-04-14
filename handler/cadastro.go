package handler

import (
	"net/http"

	"github.com/felipecveiga/cadastro-produto/model"
	"github.com/felipecveiga/cadastro-produto/service"
	"github.com/labstack/echo"
)

type Handler struct {
	Service *service.Service
}

func NewUserHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}

}

func (h *Handler) CreateProductHandler(c echo.Context) error{
	
	product := new(model.Produtos)
	if err := c.Bind(product); err!= nil {
        return c.JSON(400, err.Error())
    }
	err := h.Service.CreateProduct(product)
	if err!= nil {
        return c.JSON(400, err.Error())
    }
	return c.JSON(http.StatusOK, product)
}
