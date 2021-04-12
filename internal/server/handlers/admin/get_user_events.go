package admin

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/SKilliu/novels-backend/internal/errs"

	"github.com/SKilliu/novels-backend/internal/server/dto"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUserEvents(c echo.Context) error {
	var resp []dto.GetUserEventsResponse

	userid := c.QueryParam("user_id")
	if userid == "" {
		h.log.Error("user_id parameter is empty")
		return c.JSON(http.StatusBadRequest, errs.EmptyQueryParamErr)
	}

	tx, err := h.db.Begin()
	if err != nil {
		h.log.WithError(err).Error("failed to create db transaction")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}
	defer tx.Rollback()

	events, err := db.GetEventsByUserID(tx, userid)
	if err != nil {
		h.log.WithError(err).Error("failed to get user events from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = tx.Commit()
	if err != nil {
		h.log.WithError(err).Error("failed to commit a transaction")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	for _, e := range events {
		resp = append(resp, dto.GetUserEventsResponse{
			EventID:  e.ID,
			UserID:   e.UserID,
			DeviceID: e.DeviceID,
			Data:     e.Data,
			Time:     e.Time,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
