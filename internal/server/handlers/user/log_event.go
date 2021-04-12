package user

import (
	"net/http"
	"time"

	"github.com/SKilliu/users-rest-api/internal/db"
	"github.com/SKilliu/users-rest-api/internal/server/dto"
	"github.com/google/uuid"

	"github.com/SKilliu/users-rest-api/internal/errs"
	"github.com/labstack/echo/v4"
)

func (h *Handler) LogEvent(c echo.Context) error {
	var req dto.LogEventRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse log event request")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	tx, err := h.db.Begin()
	if err != nil {
		h.log.WithError(err).Error("failed to create db transaction")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}
	defer tx.Rollback()

	err = db.InsertEvent(tx, db.Event{
		ID:       uuid.New().String(),
		UserID:   req.UserID,
		DeviceID: req.DeviceID,
		Data:     req.Data,
		Time:     time.Unix(req.Time, 0),
	})

	if err != nil {
		h.log.WithError(err).Error("failed to insert event into db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	err = tx.Commit()
	if err != nil {
		h.log.WithError(err).Error("failed to commit a transaction")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	return c.NoContent(http.StatusOK)
}
