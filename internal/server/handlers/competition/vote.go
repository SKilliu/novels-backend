package competition

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/labstack/echo/v4"
)

// @Summary Vote for a novel
// @Security bearerAuth
// @Tags Competitions
// @Consume application/json
// @Param JSON body dto.VoteRequest true "body for a voting"
// @Description Vote for a one of two novels in competition
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/competition/vote [post]
func (h *Handler) Vote(c echo.Context) error {
	var req dto.VoteRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse vote request")
		return c.JSON(http.StatusBadRequest, errs.BadParamInBodyErr)
	}

	// !!!!!! Maybe if we have empty req.NovelID - it means that user skiped competition
	// Need to discuss it

	if req.NovelID == "" {
		readyForVote, err := h.readyForVoteDB.GetForVote()
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				return c.JSON(http.StatusOK, nil)
			default:
				h.log.WithError(err).Error("failed to get ready for vote from db")
				return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
			}
		}

		readyForVote.ViewsAmount++
		err = h.readyForVoteDB.Update(readyForVote)
		if err != nil {
			h.log.WithError(err).Error("failed to update ready for vote entity in db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}

		return c.JSON(http.StatusAccepted, "competition has skipped")
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	_, err = h.novelsDB.GetByID(req.NovelID)
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

	competition, err := h.competitionsDB.GetByNovelID(req.NovelID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("competiton doesn't exist")
			return c.JSON(http.StatusBadRequest, errs.CompetitonNotFoundErr)
		default:
			h.log.WithError(err).Error("filed to get competition from db")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if competition.Status != dto.StatusStarted {
		h.log.Error("competition is not active right now")
		return c.JSON(http.StatusConflict, errs.CompetitionIsNotActiveErr)
	}

	if competition.UserOneID == userID || competition.UserTwoID == userID {
		h.log.Error("authors of novels do not participate in voting")
		return c.JSON(http.StatusConflict, errs.CompetitionIsNotActiveErr)
	}

	readyForVote, err := h.readyForVoteDB.GetByUserAndCompetitionIDs(userID, competition.ID)
	if err != nil {
		h.log.WithError(err).Error("failed to get ready for vote entity from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if readyForVote.IsVoted {
		h.log.Error("user already voted for this competition")
		return c.JSON(http.StatusConflict, errs.UserAlreadyVotedErr)
	}

	readyForVote.ViewsAmount++
	readyForVote.IsVoted = true

	err = h.readyForVoteDB.Update(readyForVote)
	if err != nil {
		h.log.WithError(err).Error("failed to update ready for vote entity in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	if req.NovelID == competition.NovelOneID {
		competition.NovelOneVotes++
	} else {
		competition.NovelTwoVotes++
	}

	err = h.competitionsDB.Update(competition)
	if err != nil {
		h.log.WithError(err).Error("failed to update competition in db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = h.readyForVoteDB.Delete(readyForVote)
	if err != nil {
		h.log.WithError(err).Error("failed to delete ready for vote entity")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.JSON(http.StatusAccepted, "vote has accepted")
}
