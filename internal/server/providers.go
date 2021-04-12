package server

import (
	"github.com/SKilliu/users-rest-api/internal/server/handlers/admin"
	"github.com/SKilliu/users-rest-api/internal/server/handlers/user"
	"github.com/sirupsen/logrus"
)

// Provider persist handlers service structures.
type Provider struct {
	AdminHandler *admin.Handler
	UserHandler  *user.Handler
}

// NewProvider is provider constructor.
func NewProvider(logger *logrus.Entry, authKey string) *Provider {
	return &Provider{
		AdminHandler: admin.New(logger, authKey),
		UserHandler:  user.New(logger, authKey),
	}
}
