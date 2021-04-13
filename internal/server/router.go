package server

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(logger *logrus.Entry) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	m := middlewares.New(authConfig.VerifyKey)

	provider := NewProvider(logger, authConfig.VerifyKey)

	e.GET("/", healthz)
	e.POST("/signup", provider.UserHandler.SignUp)
	e.POST("/signin", provider.UserHandler.SignIn)
	e.POST("/guest_signup", provider.UserHandler.GuestSignUp)
	e.GET("/info", provider.UserHandler.GetInfo, m.ParseToken)
	e.PATCH("/verify_signup", provider.UserHandler.SignUpVerification)
	e.POST("/reset_password", provider.UserHandler.RequestResetPassword)
	e.POST("/check_password", provider.UserHandler.CheckResetPassword)

	return e
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
