package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/labstack/echo/v4"
)

// @Summary Edit username
// @Tags Authentication
// @Security bearerAuth
// @Consume application/json
// @Param JSON body dto.EditUsernameRequest true "body for edit username"
// @Description Edit username
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/edit-username [post]
func (h *Handler) EditUsername(c echo.Context) error {
	var req dto.EditUsernameRequest

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse signup request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	user, err := h.usersDB.GetByID(userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by id")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if req.Username != "" {
		user.Username = req.Username

		err = h.usersDB.Update(user)
		if err != nil {
			h.log.WithError(err).Error("failed to update user in db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	return c.NoContent(http.StatusOK)
}
