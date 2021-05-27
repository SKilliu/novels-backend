package server

import (
	"net/http"

	"github.com/SKilliu/novels-backend/internal/server/middlewares"
	"github.com/sirupsen/logrus"

	_ "github.com/SKilliu/novels-backend/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewRouter(logger *logrus.Entry) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	m := middlewares.New(authConfig.VerifyKey)

	provider := NewProvider(logger, authConfig)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", healthz)
	e.POST("/api/registration", provider.UserHandler.SignUp)
	e.POST("/api/login", provider.UserHandler.SignIn)
	e.POST("/api/guest-registration", provider.UserHandler.GuestSignIn)
	e.GET("/api/verify_signup", provider.UserHandler.SignUpVerification)
	e.POST("/api/check_password", provider.UserHandler.CheckResetPassword)
	e.GET("/api/check_password", provider.UserHandler.CheckResetPassword)
	e.POST("/api/socials-login", provider.UserHandler.SocialsSignIn)

	// with bearer token

	// user handlers
	e.GET("/api/user-info", provider.UserHandler.GetInfo, m.ParseToken)
	e.PUT("/api/user-info", provider.UserHandler.EditInfo, m.ParseToken)
	e.POST("/api/reset_password_request", provider.UserHandler.RequestResetPassword, m.ParseToken)
	e.POST("/api/edit-username", provider.UserHandler.EditUsername, m.ParseToken)
	e.PUT("/api/reset-user", provider.UserHandler.Reset, m.ParseToken)

	// novel handlers
	e.POST("/api/novel/create", provider.NovelHandler.Create, m.ParseToken)
	e.DELETE("/api/novel/delete", provider.NovelHandler.Delete, m.ParseToken)
	e.PUT("/api/novel/update", provider.NovelHandler.Update, m.ParseToken)
	e.GET("/api/novel/list", provider.NovelHandler.List, m.ParseToken)

	// competition handlers
	e.GET("/api/competition/own/get/", provider.CompetitionHandler.GetOwnCompetition, m.ParseToken)
	e.GET("/api/competition/own/list", provider.CompetitionHandler.List, m.ParseToken)
	e.GET("/api/competition/ready_for_vote", provider.CompetitionHandler.ReadyForVote, m.ParseToken)
	e.POST("/api/competition/vote", provider.CompetitionHandler.Vote, m.ParseToken)

	// admin handlers
	e.DELETE("/api/admin/drop_all", provider.AdminHandler.DropAll)
	e.POST("/api/admin/version", provider.AdminHandler.UpdateVersion)
	e.GET("/api/admin/version", provider.AdminHandler.GetVersion)
	e.GET("/api/admin/all_users", provider.AdminHandler.GetAllUsers)
	e.DELETE("/api/admin/drop_user", provider.AdminHandler.DropUser)

	return e
}

func healthz(c echo.Context) error {
	return c.JSON(http.StatusOK, "ok")
}
