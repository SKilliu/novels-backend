package novel

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/SKilliu/novels-backend/internal/db/models"
	"github.com/SKilliu/novels-backend/internal/errs"
	"github.com/SKilliu/novels-backend/internal/server/dto"
	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// @Summary Create a new novel
// @Security bearerAuth
// @Tags Novels
// @Consume application/json
// @Param JSON body dto.CreateNovelRequest true "body for a new novel creation"
// @Description Create a new novel with title and content
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.NovelResponse
// @Failure 400 {object} errs.ErrResp
// @Failure 500 {object} errs.ErrResp
// @Router /api/novel/create [post]
func (h *Handler) Create(c echo.Context) error {
	var req dto.CreateNovelRequest

	err := c.Bind(&req)
	if err != nil {
		h.log.WithError(err).Error("failed to parse create novel request")
		return c.JSON(http.StatusBadRequest, "bad param in body")
	}

	userID, _, err := middlewares.GetUserIDFromJWT(c.Request(), h.authKey)
	if err != nil {
		h.log.WithError(err).Error("failed to get user ID from token")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	nid := uuid.New().String()
	createdAt := time.Now().Unix()

	err = h.novelsDB.Insert(models.Novel{
		ID:                        nid,
		UserID:                    userID,
		Title:                     req.Title,
		Data:                      req.Data,
		ParticipatedInCompetition: false,
		CreatedAt:                 createdAt,
		UpdatedAt:                 createdAt,
	})
	if err != nil {
		h.log.WithError(err).Error("failed to insert a new novel into db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	user, err := h.usersDB.GetByID(userID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			h.log.WithError(err).Error("user doesn't exist")
			return c.JSON(http.StatusInternalServerError, errs.UserDoesntExistErr)
		default:
			h.log.WithError(err).Error("failed to get user from db by email")
			return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
		}
	}

	opponents, err := h.usersDB.GetOpponentsByRate(user.Rate)
	if err != nil {
		h.log.WithError(err).Error("Failed to get opponents for competition from db")
		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
	}

	for _, u := range opponents {
		err = h.competitionsDB.GetWaitingForOpponentByNovelID(u.ID)
	}

	return c.JSON(http.StatusOK, dto.NovelResponse{
		ID:                        nid,
		Title:                     req.Title,
		Data:                      req.Data,
		ParticipatedInCompetition: false,
		CreatedAt:                 createdAt,
		UpdatedAt:                 createdAt,
	})
}
