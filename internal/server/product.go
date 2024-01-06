package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllProducts(ctx echo.Context) error {
	vendorId := ctx.QueryParam("vendorId")

	result, err := s.DB.GetAllProducts(ctx.Request().Context(), vendorId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, result)

}
