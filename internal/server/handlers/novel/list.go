package novel

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

// @Summary Novels list
// @Security bearerAuth
// @Tags Novels
// @Consume application/json
// @Param search query string false "search by any fields in datagrid"
// @Param sort_field query string false "name of sorting field"
// @Param sort_order query string false "asc or desc"
// @Param page query string false "page number"
// @Param limit query string false "limit of items on page"
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

	// setup default params if any empty
	if sortField == "" {
		sortField = "created_at"
	}

	if sortOrder == "" {
		sortOrder = "DESC"
	}

	if limitParam == "" {
		limitParam = "1000"
	}

	if pageParam == "" {
		pageParam = "1"
	}

	err := paramValidation(sortField, sortOrder)
	if err != nil {
		h.log.WithError(err).Error("not valid params in query")
		return c.JSON(http.StatusBadRequest, errs.QueryParamIsNotValidErr)
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

// paramValidation validates all params from request
func paramValidation(sortField, sortOrder string) error {
	err := validation.Validate(sortField,
		validation.Required,
		validation.In("data", "created_at", "updated_at", "title"),
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
