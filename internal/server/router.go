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

	provider := NewProvider(logger, authConfig)

	e.GET("/", healthz)
	e.POST("/api/registration", provider.UserHandler.SignUp)
	e.POST("/api/login", provider.UserHandler.SignIn)
	e.POST("/api/guest-registration", provider.UserHandler.GuestSignIn)
	e.GET("/api/verify_signup", provider.UserHandler.SignUpVerification)
	e.POST("/api/check_password", provider.UserHandler.CheckResetPassword)
	e.GET("/api/check_password", provider.UserHandler.CheckResetPassword)
	e.POST("/api/socials-registration", provider.UserHandler.SocialsSignUp)
	e.POST("/api/socials-login", provider.UserHandler.SocialsSignIn)

	// with bearer token
	e.GET("/api/user-info", provider.UserHandler.GetInfo, m.ParseToken)
	e.POST("/api/reset_password_request", provider.UserHandler.RequestResetPassword, m.ParseToken)

	return e
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
