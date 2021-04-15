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

func (h *Handler) SocialsSignUp(c echo.Context) error {
	var req dto.SocialsSignInRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse socials sign in request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	_, err = h.userSocialsDB.GetByID(req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// create an account
			uid := uuid.New().String()
			username := fmt.Sprintf("%s-%s", req.Social, utils.GenerateName())
			err = h.usersDB.Insert(models.User{
				ID:             uid,
				Username:       username,
				HashedPassword: "no_pass_social_registration",
				Email:          fmt.Sprintf("%s - no email", username),
				DateOfBirth:    time.Now().Unix(),
				Gender:         "none",
				Membership:     "none",
				AvatarData:     "none",
				DeviceID:       "registered",
				Rate:           0,
				IsRegistered:   true,
				IsVerified:     false,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to insert new user into db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			// create user socials
			err = h.userSocialsDB.Insert(models.UserSocial{
				ID:       uuid.New().String(),
				UserID:   uid,
				Social:   req.Social,
				SocialID: req.ID,
			})
			if err != nil {
				h.log.WithError(err).Error("failed to insert user social into db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			token, err := utils.GenerateJWT(uid, "user", h.authKey)
			if err != nil {
				h.log.WithError(err).Error("failed to generate jwt token")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			return c.JSON(http.StatusOK, dto.AuthResponse{
				ID:          uid,
				Username:    username,
				Email:       fmt.Sprintf("%s - no email", username),
				Token:       token,
				DateOfBirth: time.Now().Unix(),
				Gender:      "none",
				Membership:  "none",
				AvatarData:  "none",
				Rate:        0,
			})

		default:
			h.log.WithError(err).Error("failed to get user social from db by ID")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	h.log.Error("user social already exists")
	return c.JSON(http.StatusConflict, errs.UserSocialAlreadyExistsErr)
}
