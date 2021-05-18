package user

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Guest sign in
// @Tags Authentication
// @Consume application/json
// @Param JSON body dto.GuestSignInRequest true "Body for guest sign in"
// @Description Sign in like a guest (without progress saving)
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.AuthResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/guest-registration [post]
func (h *Handler) GuestSignIn(c echo.Context) error {
	var req dto.GuestSignInRequest

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

			randomName, err := utils.GenerateName("guest")
			if err != nil {
				h.log.WithError(err).Error("failed to generate random name")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			err = h.usersDB.Insert(models.User{
				ID:       uid,
				Username: randomName,
				// Email:        fmt.Sprintf("%s has no email", randomName),
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

			// here we fill up ready for vote list for a new user
			competitions, err := h.competitionsDB.GetAllStarted()
			if err != nil {
				h.log.WithError(err).Error("failed to get all started competitions for a new user")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			if len(competitions) > 0 {
				for _, co := range competitions {
					err = h.readyForVoteDB.Insert(models.ReadyForVote{
						ID:           uuid.New().String(),
						UserID:       uid,
						NovelsPoolID: co.ID,
						IsViewed:     false,
					})
					if err != nil {
						h.log.WithError(err).Error("failed to insert a new ready for vote entity")
						return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
					}
				}
			}

			resp := dto.AuthResponse{
				ID:       uid,
				Username: randomName,
				// Email:    fmt.Sprintf("%s has no email", randomName),
				Token: token,
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
		// Email:    user.Email,
		Token:       token,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Membership:  user.Membership,
		AvatarData:  user.AvatarData,
		Rate:        user.Rate,
	}

	return c.JSON(http.StatusOK, resp)
}
