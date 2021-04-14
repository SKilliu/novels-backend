package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GuestSignUp(c echo.Context) error {
	var req dto.SignUpRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse signup request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	if req.DeviceID == "" {
		h.log.WithError(err).Error("empty device ID")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	user, err := h.usersDB.GetByDeviceID(req.DeviceID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			uid := uuid.New().String()

			randomName := utils.GenerateName()

			err = h.usersDB.Insert(models.User{
				ID:           uid,
				Username:     randomName,
				Email:        fmt.Sprintf("%s has no email", randomName),
				DeviceID:     req.DeviceID,
				DateOfBirth:  time.Now().Unix(),
				IsRegistered: false,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to create new user")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			token, err := utils.GenerateJWT(uid, "user", h.authKey)
			if err != nil {
				h.log.WithError(err).Error("failed to generate jwt token")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			resp := dto.AuthResponse{
				ID:       uid,
				Username: randomName,
				Email:    fmt.Sprintf("%s has no email", randomName),
				Token:    token,
			}

			return c.JSON(http.StatusOK, resp)
		default:
			h.log.WithError(err).Error("failed to get user form db by device ID")
			return c.JSON(http.StatusBadRequest, errs.InternalServerErr)
		}
	}

	token, err := utils.GenerateJWT(user.ID, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate jwt token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	resp := dto.AuthResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}

	return c.JSON(http.StatusOK, resp)
}
