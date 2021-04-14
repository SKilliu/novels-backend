package user

import (
	"database/sql"
	"net/http"
	"text/template"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CheckResetPassword(c echo.Context) error {
	tmpl := template.Must(template.ParseFiles("./static/web/reset_password_form.html"))

	if c.Request().Method != http.MethodPost {
		err := tmpl.Execute(c.Response(), nil)
		if err != nil {
			h.log.WithError(err).Error("failed to execute template")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
		return c.NoContent(http.StatusOK)
	}

	password := c.FormValue("password")

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

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
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

	err = h.changePassRequestsDB.Delete(reqResPass)
	if err != nil {
		h.log.WithError(err).Error("failed to delete reset password request")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = tmpl.Execute(c.Response(), struct{ Success bool }{true})
	if err != nil {
		h.log.WithError(err).Error("failed to execute template")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
