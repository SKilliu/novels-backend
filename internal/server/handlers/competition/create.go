package competition

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	var req dto.CreateCompetitionReqeust

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create competition request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	

	return c.JSON(http.StatusOK, dto.CompetitionResponse{})
}
