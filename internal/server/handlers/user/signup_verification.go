package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUpVerification(c echo.Context) error {
	token := c.QueryParam("token")

	user, err := h.usersDB.GetByID(token)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if user.IsVerified {
		h.log.WithError(err).Error("account is already verified")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	user.IsVerified = true

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
