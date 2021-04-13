package user

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/email/content"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) RequestResetPassword(c echo.Context) error {
	var req dto.SignInReq

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse change password request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	uid, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	user, err := h.usersDB.GetByEmail(req.Login)
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

	if uid != user.ID {
		h.log.Error("incorrect email for reset password")
		return c.JSON(http.StatusForbidden, errs.IncorrectAccountTypeErr)
	}

	token := uuid.New().String()

	err = h.changePassRequestsDB.Insert(models.ChangePassRequest{
		ID:     token,
		UserID: user.ID,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert new change pass request into db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	changePassRequestLink := fmt.Sprintf("http://localhost:8081/change_password?token=%s", token)

	err = h.emailClient.SendEmail(user.Username, user.Email, "Change password request",
		fmt.Sprintf(content.ChangePasswordRequestEmailContent, changePassRequestLink))
	if err != nil {
		h.log.WithError(err).Error("failed to send email for changing password")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusAccepted)
}
