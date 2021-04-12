package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/users-rest-api/internal/errs"
	"github.com/SKilliu/users-rest-api/internal/server/dto"
	"github.com/SKilliu/users-rest-api/utils"
	"golang.org/x/crypto/bcrypt"

	"github.com/SKilliu/users-rest-api/internal/db"
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

	tx, err := h.db.Begin()
	if err != nil {
		h.log.WithError(err).Error("failed to create db transaction")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}
	defer tx.Rollback()

	_, err = db.GetUserByEmail(tx, req.Email)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			passwordBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
			if err != nil {
				h.log.WithError(err).Error("failed to create hash for password")
				return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
			}

			uid := uuid.New().String()

			err = db.InsertUser(tx, db.User{
				ID:             uid,
				Name:           req.Name,
				HashedPassword: string(passwordBytes),
				Email:          req.Email,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to create new user")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			err = tx.Commit()
			if err != nil {
				h.log.WithError(err).Error("failed to commit a transaction")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			token, err := utils.GenerateJWT(uid, "user", h.authKey)
			if err != nil {
				h.log.WithError(err).Error("failed to create hash for password")
				return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
			}

			return c.JSON(http.StatusOK, dto.AuthResponse{
				Token: token,
			})
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	h.log.WithError(err).Error("user with this email already exists")
	return c.JSON(http.StatusBadRequest, errs.UserAlreadyExistErr)
}
