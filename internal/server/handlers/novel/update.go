package novel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
)

// @Summary Update a novel
// @Security bearerAuth
// @Tags Novels
// @Consume application/json
// @Param JSON body dto.UpdateNovelRequest true "body for a novel updating"
// @Description Update novel title or data
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.NovelResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/novel/update [put]
func (h *Handler) Update(c echo.Context) error {
	var req dto.UpdateNovelRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse update novel request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	novel, err := h.novelsDB.GetByID(req.ID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("novel doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.NovelNotFoundErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	if req.Title != "" {
		novel.Title = req.Title
	}

	if req.Data != "" {
		novel.Data = req.Data
	}

	novel.UpdatedAt = time.Now().Unix()

	return c.JSON(http.StatusOK, dto.NovelResponse{
		ID:                        novel.ID,
		Title:                     novel.Title,
		Data:                      novel.Data,
		ParticipatedInCompetition: novel.ParticipatedInCompetition,
		CreatedAt:                 novel.CreatedAt,
		UpdatedAt:                 novel.UpdatedAt,
	})
}
