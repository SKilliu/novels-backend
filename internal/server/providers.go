package server

import (
	"github.com/SKilliu/novels-backend/internal/server/handlers/user"
	"github.com/sirupsen/logrus"
)

// Provider persist handlers service structures.
type Provider struct {
	// AdminHandler *admin.Handler
	UserHandler *user.Handler
}

// NewProvider is provider constructor.
func NewProvider(logger *logrus.Entry, auth Auth) *Provider {
	return &Provider{
		// AdminHandler: admin.New(logger, authKey),
		UserHandler: user.New(logger, auth.VerifyKey, auth.EmailAddress, auth.EmailPassword),
	}
}
