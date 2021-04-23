package novel

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/SKilliu/novels-backend/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

// @Summary Create a new novel
// @Security bearerAuth
// @Tags Novels
// @Consume application/json
// @Param JSON body dto.CreateNovelRequest true "body for a new novel creation"
// @Description Create a new novel with title and content
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.NovelResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/novel/create [post]
func (h *Handler) Create(c echo.Context) error {
	var req dto.CreateNovelRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create novel request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	createdAt := time.Now().Unix()

	novel := models.Novel{
		ID:                        uuid.New().String(),
		UserID:                    userID,
		Title:                     req.Title,
		Data:                      req.Data,
		ParticipatedInCompetition: false,
		CreatedAt:                 createdAt,
		UpdatedAt:                 createdAt,
	}

	err = h.novelsDB.Insert(novel)
	if err != nil {
		h.log.WithError(err).Error("failed to insert a new novel into db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	user, err := h.usersDB.GetByID(userID)
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

	competitionOpponent, err := h.competitionsDB.GetCompetitionOpponent(user.Rate, user.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			err = h.competitionsDB.Insert(models.Competition{
				ID:         uuid.New().String(),
				NovelOneID: novel.ID,
				UserOneID:  user.ID,
				Status:     dto.StatusWaitingForOpponent,
				CreatedAt:  time.Now().Unix(),
				UpdatedAt:  time.Now().Unix(),
			})
			if err != nil {
				h.log.WithError(err).Error("failed to create new competition")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}

			// A job checks the competition and if opponent will not be found - change status to "expired"
			cr := cron.New(cron.WithLocation(time.UTC))
			cr.Start()

			now := time.Now().UTC()
			now = now.Add(time.Minute * 2).UTC()
			_, err = cr.AddFunc(fmt.Sprintf("%d %d %d %d *", now.Minute(), now.Hour(), now.Day(), now.Month()), func() {

				competition, err := h.competitionsDB.GetByNovelOneID(novel.ID)
				if err != nil {
					switch err {
					case sql.ErrNoRows:
						h.log.WithError(err).Error("competition doesn't exist")
						return
					default:
						h.log.WithError(err).Error("failed ot get competition from db")
						return
					}
				}

				if competition.Status == dto.StatusWaitingForOpponent {
					competition.Status = dto.StatusExpired

					err = h.competitionsDB.Update(competition)
					if err != nil {
						h.log.WithError(err).Error("failed to update a competition in db")
						return
					}

					novel.VotingResult = 51
					err = h.novelsDB.Update(novel)
					if err != nil {
						h.log.WithError(err).Error("failed to update novel in db")
					}
				}

				cr.Stop()
			})
			if err != nil {
				h.log.WithError(err).Error("failed to start job")
			}

			return c.JSON(http.StatusOK, dto.NovelResponse{
				ID:                        novel.ID,
				Title:                     req.Title,
				Data:                      req.Data,
				ParticipatedInCompetition: novel.ParticipatedInCompetition,
				CreatedAt:                 novel.CreatedAt,
				UpdatedAt:                 novel.UpdatedAt,
			})

		default:
			h.log.WithError(err).Error("failed to get opponent for novel")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	competitionStart := time.Now()

	competitionOpponent.NovelTwoID = novel.ID
	competitionOpponent.Status = dto.StatusStarted
	competitionOpponent.UserTwoID = user.ID
	competitionOpponent.UpdatedAt = time.Now().Unix()
	competitionOpponent.CompetitionStartedAt = competitionStart.Unix()

	err = h.competitionsDB.Update(competitionOpponent)
	if err != nil {
		h.log.WithError(err).Error("failed to update competition in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	users, err := h.usersDB.GetAllForVote(competitionOpponent.UserOneID, competitionOpponent.UserTwoID)
	if err != nil {
		h.log.WithError(err).Error("failed to get users from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	for _, u := range users {
		err = h.readyForVoteDB.Insert(models.ReadyForVote{
			ID:           uuid.New().String(),
			UserID:       u.ID,
			NovelsPoolID: competitionOpponent.ID,
			ViewsAmount:  0,
			IsVoted:      false,
		})
		if err != nil {
			h.log.WithError(err).Error("failed to insert a new ready for vote entity")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	// after 2 hours voting will be closed
	cr := cron.New(cron.WithLocation(time.UTC))
	cr.Start()

	now := competitionStart.Add(time.Minute * 5).UTC()
	_, err = cr.AddFunc(fmt.Sprintf("%d %d %d %d *", now.Minute(), now.Hour(), now.Day(), now.Month()), func() {

		competition, err := h.competitionsDB.GetByNovelOneID(novel.ID)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				h.log.WithError(err).Error("competition doesn't exist")
				return
			default:
				h.log.WithError(err).Error("failed ot get competition from db")
				return
			}
		}

		competition.Status = dto.StatusFinished

		err = h.competitionsDB.Update(competition)
		if err != nil {
			h.log.WithError(err).Error("failed to update a competition in db")
			return
		}

		// here we need to determine the competition winner
		resNovelOne, resNovelTwo := utils.GetVotingResults(competition.NovelOneVotes, competition.NovelTwoVotes)

		novel.VotingResult = resNovelOne

		err = h.novelsDB.Update(novel)
		if err != nil {
			h.log.WithError(err).Error("failed to update first novel in db")
		}

		opponentNovel, err := h.novelsDB.GetByID(competition.NovelTwoID)
		if err != nil {
			h.log.WithError(err).Error("failed to get opponent novel from db")
		}

		opponentNovel.VotingResult = resNovelTwo

		err = h.novelsDB.Update(opponentNovel)
		if err != nil {
			h.log.WithError(err).Error("failed to update opponent novel in db")
		}

		cr.Stop()
	})
	if err != nil {
		h.log.WithError(err).Error("failed to start job")
	}

	return c.JSON(http.StatusOK, dto.NovelResponse{
		ID:                        novel.ID,
		Title:                     req.Title,
		Data:                      req.Data,
		ParticipatedInCompetition: false,
		CreatedAt:                 createdAt,
		UpdatedAt:                 createdAt,
	})
}
