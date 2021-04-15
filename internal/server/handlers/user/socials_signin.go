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

// @Summary Socials sign in
// @Tags authentication
// @Consume application/json
// @Param JSON body dto.SocialsSignInRequest true "body for sign up"
// @Description User login by socials (Facebook, Google, Apple, etc.). If user doesn't exist in DB, new account will be created.
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/socials-login [post]
func (h *Handler) SocialsSignIn(c echo.Context) error {
	var (
		req      dto.SocialsSignInRequest
		uid      string
		email    string
		username string
	)

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse socials sign in request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	userSocial, err := h.userSocialsDB.GetByID(req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			// create an account
			uid = uuid.New().String()
			username = fmt.Sprintf("%s-%s", req.Social, utils.GenerateName())
			email = fmt.Sprintf("%s - no email", username)
			err = h.usersDB.Insert(models.User{
				ID:             uid,
				Username:       username,
				HashedPassword: "no_pass_social_registration",
				Email:          email,
				DateOfBirth:    time.Now().Unix(),
				Gender:         "none",
				Membership:     "none",
				AvatarData:     "none",
				DeviceID:       "registered",
				Rate:           0,
				IsRegistered:   true,
				IsVerified:     true,
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

		default:
			h.log.WithError(err).Error("failed to get user social from db by ID")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	} else {
		// get existed account
		user, err := h.usersDB.GetByID(userSocial.UserID)
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

		uid = user.ID
		username = user.Username
		email = user.Email
	}

	token, err := utils.GenerateJWT(uid, "user", h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to generate jwt token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, dto.AuthResponse{
		ID:          uid,
		Username:    username,
		Email:       email,
		Token:       token,
		DateOfBirth: time.Now().Unix(),
		Gender:      "none",
		Membership:  "none",
		AvatarData:  "none",
		Rate:        0,
	})
}
