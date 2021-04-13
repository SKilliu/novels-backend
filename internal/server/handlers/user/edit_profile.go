package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EditProfile(c echo.Context) error {
	// var req dto.EditProfileRequest

	return c.NoContent(http.StatusOK)
}
