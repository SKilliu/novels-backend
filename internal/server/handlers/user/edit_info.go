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

// @Summary Edit info
// @Tags Authentication
// @Security bearerAuth
// @Consume application/json
// @Param JSON body dto.EditInfoRequest true "body for edit info"
// @Description Edit user info
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/user-info [put]
func (h *Handler) EditInfo(c echo.Context) error {
	var req dto.EditInfoRequest

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

	if req.AvatarData != "" {
		user.AvatarData = req.AvatarData
	}

	if req.Rate != 0 {
		user.Rate = req.Rate
	}

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	token, err := utils.GenerateJWT(user.ID, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate jwt token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, dto.AuthResponse{
		ID:          user.ID,
		Username:    user.Username,
		// Email:       user.Email,
		Token:       token,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Membership:  user.Membership,
		AvatarData:  user.AvatarData,
		Rate:        user.Rate,
	})
}
