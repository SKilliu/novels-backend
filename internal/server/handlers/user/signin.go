package user

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignIn(c echo.Context) error {
	var (
		req  dto.SignInReq
		resp dto.AuthResponse
	)

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse signup request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
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

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		h.log.WithError(err).Error("incorrect email or password")
		return c.JSON(http.StatusBadRequest, errs.WrongCredentialsErr)
	}

	token, err := utils.GenerateJWT(user.ID, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate jwt token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	resp = dto.AuthResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Token:       token,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Membership:  user.Membership,
		AvatarData:  user.AvatarData,
		Rate:        user.Rate,
	}

	return c.JSON(http.StatusOK, resp)
}