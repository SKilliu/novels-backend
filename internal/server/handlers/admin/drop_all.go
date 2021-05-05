package admin

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/labstack/echo/v4"
)

// @Summary Drop all users
// @Security bearerAuth
// @Tags admin
// @Consume application/json
// @Description Drop all users from the database
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/admin/drop_all [delete]
func (h *Handler) DropAll(c echo.Context) error {
	err := h.usersDB.DropAll()
	if err != nil {
		h.log.WithError(err).Error("failed to delete all users from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = h.competitionsDB.DropAll()
	if err != nil {
		h.log.WithError(err).Error("failed to drop all competitions from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
