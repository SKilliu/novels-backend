package user

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/SKilliu/novels-backend/email/content"
	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Reset password request
// @Security bearerAuth
// @Tags Authentication
// @Consume application/json
// @Param JSON body dto.ResetPasswordRequest true "email for reset password"
// @Description Reset your account password
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/reset_password_request [post]
func (h *Handler) RequestResetPassword(c echo.Context) error {
	var req dto.ResetPasswordRequest
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

	user, err := h.usersDB.GetByID(uid)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by ID")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	token := uuid.New().String()

	err = h.changePassRequestsDB.Insert(models.ResetPassRequest{
		ID:     token,
		UserID: user.ID,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert new change pass request into db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	t, err := template.New("email").Parse(content.ChangePasswordRequestEmailContent)
	if err != nil {
		h.log.WithError(err).Error("failed to parse template")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	data := struct {
		URL  string
		Name string
	}{
		URL:  fmt.Sprintf("https://165.227.207.77:8000/api/check_password?token=%s", token),
		Name: user.Username,
	}

	var parsedHTML bytes.Buffer
	err = t.Execute(&parsedHTML, data)
	if err != nil {
		h.log.WithError(err).Error("failed to execute template")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = utils.SendEmail(h.email, h.password, req.Email, parsedHTML.String(), "Reset password request", "text/html")
	if err != nil {
		h.log.WithError(err).WithField("user_email", req.Email).Error("failed to send confirmation email to user")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	// changePassRequestLink := fmt.Sprintf("http://localhost:8081/change_password?token=%s", token)

	// err = h.emailClient.SendEmail(user.Username, user.Email, "Change password request",
	// 	fmt.Sprintf(content.ChangePasswordRequestEmailContent, changePassRequestLink))
	// if err != nil {
	// 	h.log.WithError(err).Error("failed to send email for changing password")
	// 	return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	// }

	return c.NoContent(http.StatusAccepted)
}
