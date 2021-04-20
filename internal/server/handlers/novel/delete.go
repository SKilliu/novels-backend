package novel

import (
	"database/sql"
	"net/http"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/labstack/echo/v4"
)

// @Summary Delete novel
// @Security bearerAuth
// @Tags Novels
// @Consume application/json
// @Param id query string true "novel id"
// @Description Delete user novels by ID
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/novel/delete [delete]
func (h *Handler) Delete(c echo.Context) error {
	novelID := c.QueryParam("id")

	novel, err := h.novelsDB.GetByID(novelID)
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

	err = h.novelsDB.Delete(novel)
	if err != nil {
		h.log.WithError(err).Error("failed to delete novel from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
