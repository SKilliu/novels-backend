package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/labstack/echo/v4"
)

// @Summary Get user info
// @Security bearerAuth
// @Tags Authentication
// @Consume application/json
// @Description Get user info by user ID from bearer token
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/user-info [get]
func (h *Handler) GetInfo(c echo.Context) error {

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
			return c.JSON(http.StatusBadRequest, errs.UserNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get user from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	token, err := utils.GenerateJWT(user.ID, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to create hash for password")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	resp := dto.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		// Email:       user.Email,
		Token:       token,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Membership:  user.Membership,
		AvatarData:  user.AvatarData,
		Rate:        user.Rate,
	}

	return c.JSON(http.StatusOK, resp)
}
