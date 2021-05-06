package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/labstack/echo/v4"
)

// @Summary Reset user
// @Tags Authentication
// @Security bearerAuth
// @Consume application/json
// @Description Reset all user progress and rate
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/reset-user [put]
func (h *Handler) Reset(c echo.Context) error {
	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
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

	user.AvatarData = ""
	user.Rate = 0

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
