package server

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(logger *logrus.Entry) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//m := middlewares.New(authConfig.VerifyKey)

	provider := NewProvider(logger, authConfig.VerifyKey)

	e.GET("/", healthz)
	e.POST("/signup", provider.UserHandler.SignUp)
	// e.GET("/admin/panel", provider.AdminHandler.GetAdminPanel)
	// e.POST("/user/log_event", provider.UserHandler.LogEvent)
	// e.GET("/admin/events", provider.AdminHandler.GetUserEvents)

	return e
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "Server successfully started")
}
