package admin

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/labstack/echo/v4"
)

// @Summary Drop user
// @Security bearerAuth
// @Tags admin
// @Consume application/json
// @Description Drop user by ID from the database
// @Param user_id query string true "user_id"
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOK
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /admin/drop_user [delete]
func (h *Handler) DropUser(c echo.Context) error {

	userID := c.QueryParam("user_id")

	err := h.usersDB.DeleteByID(userID)
	if err != nil {
		h.log.WithError(err).Error("failed to drop user from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}
	return c.NoContent(http.StatusOK)
}
