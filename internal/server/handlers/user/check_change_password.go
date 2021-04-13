package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CheckResetPassword(c echo.Context) error {
	var req dto.SignUpRequest
	token := c.QueryParam("token")

	reqResPass, err := h.changePassRequestsDB.GetByID(token)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("reset password request doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		default:
			h.log.WithError(err).Error("failed to get reset password request from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	user, err := h.usersDB.GetByID(reqResPass.UserID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by id")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		h.log.WithError(err).Error("failed to create hash for password")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	user.HashedPassword = string(passwordBytes)

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
