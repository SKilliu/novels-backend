package novel

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
)

// @Summary Novels list
// @Security bearerAuth
// @Tags Novels
// @Consume application/json
// @Param search query string true "search by any fields in datagrid"
// @Param sort_field query string true "name of sorting field"
// @Param sort_order query string true "asc or desc"
// @Param page query string true "page number"
// @Param limit query string true "limit of items on page"
// @Description Get novels list by search parameter, sorting and pagination
// @Accept  json
// @Produce  json
// @Success 200 {object} []dto.NovelResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/novel/list [get]
func (h *Handler) List(c echo.Context) error {
	var resp []dto.NovelResponse

	search := c.QueryParam("search")
	sortField := c.QueryParam("sort_field")
	sortOrder := c.QueryParam("sort_order")
	limitParam := c.QueryParam("limit")
	pageParam := c.QueryParam("page")

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

	paramForQuery := fmt.Sprintf("'%s%s%s'", dto.PercentSymbol, search, dto.PercentSymbol)

	offset := page*limit - limit

	novels, err := h.novelsDB.GetListWithParam(paramForQuery, sortField, sortOrder, offset, limit)
	if err != nil {
		h.log.WithError(err).Error("failed to get novels list from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	for _, n := range novels {
		resp = append(resp, dto.NovelResponse{
			ID:                        n.ID,
			Title:                     n.Title,
			Data:                      n.Data,
			ParticipatedInCompetition: n.ParticipatedInCompetition,
			CreatedAt:                 n.CreatedAt,
			UpdatedAt:                 n.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
