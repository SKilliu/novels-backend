package competition

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/labstack/echo/v4"
)

// @Summary Get novels for vote
// @Security bearerAuth
// @Tags Competitions
// @Consume application/json
// @Description Get novels pair for vote
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CompetitionResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/competition/ready_for_vote [get]
func (h *Handler) ReadyForVote(c echo.Context) error {
	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	readyForVote, err := h.readyForVoteDB.GetForVote(userID)
	fmt.Println("============================>")
	fmt.Println(readyForVote)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return c.JSON(http.StatusOK, nil)
		default:
			h.log.WithError(err).Error("failed to get ready for vote from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	competition, err := h.competitionsDB.GetByID(readyForVote.NovelsPoolID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("competition doesn't exist")
			return c.JSON(http.StatusBadRequest, errs.CompetitonNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get competition from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	novelOne, err := h.novelsDB.GetByID(competition.NovelOneID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("first novel doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.NovelNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get first novel from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	userOne, err := h.usersDB.GetByID(competition.UserOneID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("first user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get first user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	novelTwo, err := h.novelsDB.GetByID(competition.NovelTwoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("second novel doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.NovelNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get second novel from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	userTwo, err := h.usersDB.GetByID(competition.UserTwoID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("second user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get second user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	readyForVote.IsViewed = true
	err = h.readyForVoteDB.Update(readyForVote)
	if err != nil {
		h.log.WithError(err).Error("failed to update ready for vote entity in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusOK, dto.CompetitionResponse{
		ID: competition.ID,
		NovelOne: &dto.NovelData{
			User: dto.UserData{
				Username:    userOne.Username,
				DateOfBirth: userOne.DateOfBirth,
				Gender:      userOne.Gender,
				Membership:  userOne.Membership,
				Rate:        userOne.Rate,
			},
			NovelResponse: dto.NovelResponse{
				ID:                        novelOne.ID,
				Title:                     novelOne.Title,
				Data:                      novelOne.Data,
				ParticipatedInCompetition: novelOne.ParticipatedInCompetition,
				CreatedAt:                 novelOne.CreatedAt,
				UpdatedAt:                 novelOne.UpdatedAt,
			},
		},
		NovelTwo: &dto.NovelData{
			User: dto.UserData{
				Username:    userTwo.Username,
				DateOfBirth: userTwo.DateOfBirth,
				Gender:      userTwo.Gender,
				Membership:  userTwo.Membership,
				Rate:        userTwo.Rate,
			},
			NovelResponse: dto.NovelResponse{
				ID:                        novelTwo.ID,
				Title:                     novelTwo.Title,
				Data:                      novelTwo.Data,
				ParticipatedInCompetition: novelTwo.ParticipatedInCompetition,
				CreatedAt:                 novelTwo.CreatedAt,
				UpdatedAt:                 novelTwo.UpdatedAt,
			},
		},
		CompetitionStartedAt: competition.CompetitionStartedAt,
		Status:               competition.Status,
		CreatedAt:            competition.CreatedAt,
		UpdatedAt:            competition.UpdatedAt,
	})
}
