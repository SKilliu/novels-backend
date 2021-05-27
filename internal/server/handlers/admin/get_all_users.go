package admin

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
)

// @Summary Get all users
// @Security bearerAuth
// @Tags admin
// @Consume application/json
// @Description Get all users from db
// @Accept  json
// @Produce  json
// @Success 200 {} http.StatusOK
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/admin/all_users [get]
func (h *Handler) GetAllUsers(c echo.Context) error {
	var resp []dto.GetAllUsersResponse
	users, err := h.usersDB.GetAll()
	if err != nil {
		h.log.WithError(err).Error("failed to get all users from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)

	}

	for _, u := range users {
		resp = append(resp, dto.GetAllUsersResponse{
			ID:       u.ID,
			Username: u.Username,
			// Email:    u.Email,
			DeviceID: u.DeviceID,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
