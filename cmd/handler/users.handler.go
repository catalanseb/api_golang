package users

import (
	"net/http"

	handle_error "api-go/internal/errors"

	"github.com/labstack/echo/v4"
)

func (h *UsersHandler) GetAll(c echo.Context) error {

	response, err := h.service.GetAll()
	if err != nil {
		e := handle_error.ErrResp(400, "error", err.Error())
		return c.JSON(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, response)
}

func (h *UsersHandler) GetById(c echo.Context) error {
	return nil
}

func (h *UsersHandler) Register(c echo.Context) error {
	return nil
}

func (h *UsersHandler) Update(c echo.Context) error {
	return nil
}

func (h *UsersHandler) Delete(c echo.Context) error {
	return nil
}
