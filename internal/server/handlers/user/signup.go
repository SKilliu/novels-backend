package user

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/db/models"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SignUp(c echo.Context) error {
	var req dto.SignUpRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse signup request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	_, err = h.usersDB.GetByEmail(req.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:

			amount, err := h.usersDB.CheckUserByUsername(req.Username)
			if err != nil {
				h.log.WithError(err).Error("failed to get user from db by username")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			if amount.Exists {
				h.log.WithError(err).Error("username already exists")
				return c.JSON(http.StatusInternalServerError, errs.UsernameAlreadyExistsErr)
			}

			passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
			if err != nil {
				h.log.WithError(err).Error("failed to create hash for password")
				return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
			}

			uid := uuid.New().String()

			err = h.usersDB.Insert(models.User{
				ID:             uid,
				Username:       req.Username,
				HashedPassword: string(passwordBytes),
				Email:          req.Email,
				DeviceID:       req.DeviceID,
				DateOfBirth:    time.Now().Unix(),
			})
			if err != nil {
				h.log.WithError(err).Error("failed to create new user")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			token, err := utils.GenerateJWT(uid, "user", h.authKey)
			if err != nil {
				h.log.WithError(err).Error("failed to create hash for password")
				return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
			}

			return c.JSON(http.StatusOK, dto.AuthResponse{
				ID:       uid,
				Username: req.Username,
				Email:    req.Email,
				Token:    token,
			})
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	h.log.WithError(err).Error("user with this email already exists")
	return c.JSON(http.StatusBadRequest, errs.EmailAlreadyExistErr)
}
