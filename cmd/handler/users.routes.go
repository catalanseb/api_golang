package users

import (
	users "api-go/internal/services"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	service users.ServiceUsers
}

func NewUsersHandler(e *echo.Echo, service users.ServiceUsers) *UsersHandler {
	h := &UsersHandler{
		service: service,
	}
	h.RegistryURI(e)
	return h

}

func (h *UsersHandler) RegistryURI(e *echo.Echo) {

	routeUsers := e.Group("/api/v1/users")

	routeUsers.GET("", h.GetAll)
	routeUsers.GET("/:id", h.GetById)
	routeUsers.POST("", h.Register)
	routeUsers.PATCH("", h.Register)
	routeUsers.DELETE("/:id", h.Delete)
}
