package admin

// import (
// 	"html/template"
// 	"net/http"
// 	"os"

// 	"github.com/SKilliu/novels-backend/internal/errs"

// 	"github.com/SKilliu/novels-backend/internal/db"

// 	"github.com/labstack/echo/v4"
// )

// func (h *Handler) GetAdminPanel(c echo.Context) error {

// 	file, err := os.Open("./web/static/admin.html")
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	t, err := template.ParseFiles("./web/static/admin.html")
// 	if err != nil {
// 		return err
// 	}

// 	tx, err := h.db.Begin()
// 	if err != nil {
// 		h.log.WithError(err).Error("failed to create db transaction")
// 		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
// 	}
// 	defer tx.Rollback()

// 	// we get all users from db
// 	users, err := db.GetAllUsers(tx)
// 	if err != nil {
// 		h.log.WithError(err).Error("failed to get all users from db")
// 		return c.JSON(http.StatusInternalServerError, errs.InternalServerErr)
// 	}

// 	data := struct {
// 		UsersList []db.User
// 	}{
// 		UsersList: users,
// 	}

// 	err = t.Execute(c.Response(), data)
// 	if err != nil {
// 		return err
// 	}

// 	return c.NoContent(http.StatusOK)
// }
