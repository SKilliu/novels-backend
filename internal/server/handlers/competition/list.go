package competition

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo/v4"
)

// @Summary Competitions list
// @Security bearerAuth
// @Tags Competitions
// @Consume application/json
// @Param status query string true "can be <b>waiting_for_opponent</b>, <b>started</b>, <b>expired</b>, <b>finished</b> or can be skipped"
// @Param sort_field query string true "name of sorting field"
// @Param sort_order query string true "asc or desc"
// @Param page query string true "page number"
// @Param limit query string true "limit of items on page"
// @Description Get competitions list by status, sorting and pagination
// @Accept  json
// @Produce  json
// @Success 200 {object} []dto.CompetitionResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/competition/own/list [get]
func (h *Handler) List(c echo.Context) error {
	var resp []dto.CompetitionResponse

	status := c.QueryParam("status")
	sortField := c.QueryParam("sort_field")
	sortOrder := c.QueryParam("sort_order")
	limitParam := c.QueryParam("limit")
	pageParam := c.QueryParam("page")

	err := paramValidation(status, sortField, sortOrder)
	if err != nil {
		h.log.WithError(err).Error("not valid params in query")
		return c.JSON(http.StatusBadRequest, errs.QueryParamIsNotValidErr)
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		h.log.WithError(err).Error("failed to parse page string value to int")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		h.log.WithError(err).Error("failed to parse limit string value to int")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	// paramForQuery := fmt.Sprintf("'%s%s%s'", dto.PercentSymbol, status, dto.PercentSymbol)

	offset := page*limit - limit

	competitions, err := h.competitionsDB.GetListWithParam(status, userID, sortField, sortOrder, offset, limit)
	if err != nil {
		h.log.WithError(err).Error("failed ot get competitions from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	fmt.Println(competitions)

	for _, comp := range competitions {
		var (
			novelTwoData *dto.NovelData
			novelOneData *dto.NovelData
		)
		novel, err := h.novelsDB.GetByID(comp.NovelOneID)
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

		user, err := h.usersDB.GetByID(comp.UserOneID)
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

		novelOneData = &dto.NovelData{
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
		}

		if comp.NovelTwoID == "" {
			novelTwoData = nil
		} else {
			novel, err := h.novelsDB.GetByID(comp.NovelTwoID)
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

			user, err := h.usersDB.GetByID(comp.UserTwoID)
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

			novelTwoData = &dto.NovelData{
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
			}
		}

		resp = append(resp, dto.CompetitionResponse{
			ID:                   comp.ID,
			NovelOne:             novelOneData,
			NovelTwo:             novelTwoData,
			CompetitionStartedAt: comp.CompetitionStartedAt,
			Status:               comp.Status,
			CreatedAt:            comp.CreatedAt,
			UpdatedAt:            comp.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, resp)
}

// paramValidation validates all params from request
func paramValidation(status, sortField, sortOrder string) error {
	err := validation.Validate(status,
		validation.Required,
		validation.In(dto.StatusWaitingForOpponent, dto.StatusStarted, dto.StatusExpired, dto.StatusFinished),
	)
	if err != nil {
		return err
	}

	err = validation.Validate(sortField,
		validation.Required,
		validation.In("competition_started_at", "created_at", "updated_at", "status"),
	)
	if err != nil {
		return err
	}

	err = validation.Validate(sortOrder,
		validation.Required,
		validation.In("asc", "desc", "ASC", "DESC"),
	)
	if err != nil {
		return err
	}

	return nil
}
