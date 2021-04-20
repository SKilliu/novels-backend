package competition

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
)

// @Summary Get own competition
// @Security bearerAuth
// @Tags Competitions
// @Consume application/json
// @Param novel_id query string true "novel_id in db"
// @Description Get own competition by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CompetitionResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/competition/own/get/ [get]
func (h *Handler) GetOwnCompetition(c echo.Context) error {
	var (
		novelTwo models.Novel
		userTwo  models.User
	)

	novelID := c.QueryParam("novel_id")

	if novelID == "" {
		h.log.Error("emplty query param")
		return c.JSON(http.StatusBadRequest, errs.EmptyQueryParamErr)
	}

	novel, err := h.novelsDB.GetByID(novelID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("novel doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.NovelNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get novel from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	competition, err := h.competitionsDB.GetByNovelID(novelID)
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

	user, err := h.usersDB.GetByID(novel.UserID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if competition.NovelOneID != novelID {
		novelTwo = novel
		userTwo = user

		novel, err = h.novelsDB.GetByID(competition.NovelOneID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("novel doesn't exist")
				return c.JSON(http.StatusInternalServerError, errs.NovelNotFoundErr)
			default:
				h.log.WithError(err).Error("failed to get novel from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}

		user, err = h.usersDB.GetByID(competition.UserOneID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("user doesn't exist")
				return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
			default:
				h.log.WithError(err).Error("failed to get user from db by email")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}
	} else {
		novelTwo, err = h.novelsDB.GetByID(competition.NovelOneID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("novel doesn't exist")
				return c.JSON(http.StatusInternalServerError, errs.NovelNotFoundErr)
			default:
				h.log.WithError(err).Error("failed to get novel from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}

		userTwo, err = h.usersDB.GetByID(competition.UserOneID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("user doesn't exist")
				return c.JSON(http.StatusInternalServerError, errs.UserNotFoundErr)
			default:
				h.log.WithError(err).Error("failed to get user from db by email")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}
	}

	return c.JSON(http.StatusOK, dto.CompetitionResponse{
		ID: competition.ID,
		NovelOne: &dto.NovelData{
			User: dto.UserData{
				Username:    user.Username,
				DateOfBirth: user.DateOfBirth,
				Gender:      user.Gender,
				Membership:  user.Membership,
				Rate:        user.Rate,
			},
			NovelResponse: dto.NovelResponse{
				ID:                        novel.ID,
				Title:                     novel.Title,
				Data:                      novel.Data,
				ParticipatedInCompetition: novel.ParticipatedInCompetition,
				CreatedAt:                 novel.CreatedAt,
				UpdatedAt:                 novel.UpdatedAt,
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
