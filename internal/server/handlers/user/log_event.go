package user

// import (
// 	"net/http"
// 	"time"

// 	"github.com/SKilliu/novels-backend/internal/db"
// 	"github.com/SKilliu/novels-backend/internal/server/dto"
// 	"github.com/google/uuid"

// 	"github.com/SKilliu/novels-backend/internal/errs"
// 	"github.com/labstack/echo/v4"
// )

// func (h *Handler) LogEvent(c echo.Context) error {
// 	var req dto.LogEventRequest

// 	err := c.Bind(&req)
// 	if err != nil {
// 		h.log.WithError(err).Error("failed to parse log event request")
// 		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
// 	}

// 	err = db.InsertEvent(tx, db.Event{
// 		ID:       uuid.New().String(),
// 		UserID:   req.UserID,
// 		DeviceID: req.DeviceID,
// 		Data:     req.Data,
// 		Time:     time.Unix(req.Time, 0),
// 	})

// 	if err != nil {
// 		h.log.WithError(err).Error("failed to insert event into db")
// 		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		h.log.WithError(err).Error("failed to commit a transaction")
// 		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
// 	}

// 	return c.NoContent(http.StatusOK)
// }
