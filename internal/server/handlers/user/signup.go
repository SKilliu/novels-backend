package user

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/SKilliu/novels-backend/email/content"
	"github.com/SKilliu/novels-backend/internal/db/models"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

// @Summary Sign up
// @Tags authentication
// @Consume application/json
// @Param JSON body dto.SignUpRequest true "Body for sign up"
// @Description User registraton by email and password
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/registration [post]
func (h *Handler) SignUp(c echo.Context) error {
	var (
		req dto.SignUpRequest
		uid string
	)

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse signup request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	// Here we need to check the user's email and username for already existing in DB
	info, err := h.usersDB.CheckUserByUsername(req.Username)
	if err != nil {
		h.log.WithError(err).Error("failed to get user from db by username")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if info.Exists {
		h.log.WithError(err).Error("username already exists")
		return c.JSON(http.StatusInternalServerError, errs.UsernameAlreadyExistsErr)
	}

	info, err = h.usersDB.CheckUserByEmail(req.Email)
	if err != nil {
		h.log.WithError(err).Error("failed to get user from db by email")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if info.Exists {
		h.log.WithError(err).Error("user with this email already exists")
		return c.JSON(http.StatusBadRequest, errs.EmailAlreadyExistErr)
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		h.log.WithError(err).Error("failed to create hash for password")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	user, err := h.usersDB.GetByDeviceID(req.DeviceID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			uid = uuid.New().String()

			err = h.usersDB.Insert(models.User{
				ID:             uid,
				Username:       req.Username,
				HashedPassword: string(passwordBytes),
				Email:          req.Email,
				DeviceID:       "registered",
				DateOfBirth:    time.Now().Unix(),
				IsRegistered:   true,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to create new user")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	} else {
		uid = user.ID
	}

	user.Username = req.Username
	user.Email = req.Email
	user.HashedPassword = string(passwordBytes)
	user.DeviceID = "registered"

	err = h.usersDB.Update(user)
	if err != nil {
		h.log.WithError(err).Error("failed to update user in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	token, err := utils.GenerateJWT(uid, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate jwt token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	t, err := template.New("email").Parse(content.SignUpVerificationEmailContent)
	if err != nil {
		h.log.WithError(err).Error("failed to parse template")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	data := struct {
		URL  string
		Name string
	}{
		URL:  fmt.Sprintf("https://165.227.207.77:8000/verify_signup?token=%s", uid),
		Name: req.Username,
	}

	var parsedHTML bytes.Buffer
	err = t.Execute(&parsedHTML, data)
	if err != nil {
		h.log.WithError(err).Error("failed to execute template")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	go func() {
		err = utils.SendEmail(h.email, h.password, req.Email, parsedHTML.String(), "Account verification", "text/html")
		if err != nil {
			h.log.WithError(err).WithField("user_email", req.Email).Error("failed to send confirmation email to user")
			// return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}()

	return c.JSON(http.StatusOK, dto.AuthResponse{
		ID:       uid,
		Username: req.Username,
		Email:    req.Email,
		Token:    token,
	})
}
