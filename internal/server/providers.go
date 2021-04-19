package server

import (
	"github.com/SKilliu/novels-backend/internal/server/handlers/competition"
	"github.com/SKilliu/novels-backend/internal/server/handlers/novel"
	"github.com/SKilliu/novels-backend/internal/server/handlers/user"
	"github.com/sirupsen/logrus"
)

// Provider persist handlers service structures.
type Provider struct {
	NovelHandler       *novel.Handler
	UserHandler        *user.Handler
	CompetitionHandler *competition.Handler
}

// NewProvider is provider constructor.
func NewProvider(logger *logrus.Entry, auth Auth) *Provider {
	return &Provider{
		// AdminHandler: admin.New(logger, authKey),
		UserHandler:        user.New(logger, auth.VerifyKey, auth.EmailAddress, auth.EmailPassword),
		NovelHandler:       novel.New(logger, auth.VerifyKey),
		CompetitionHandler: competition.New(logger, auth.VerifyKey),
	}
}
