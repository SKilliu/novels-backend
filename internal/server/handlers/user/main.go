package user

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/SKilliu/novels-backend/internal/email"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log         *logrus.Entry
	usersDB     db.UsersQ
	authKey     string
	emailClient *email.EmailClient
}

func New(logger *logrus.Entry, authKey string) *Handler {
	return &Handler{
		log:         logger,
		usersDB:     db.Connection().UsersQ(),
		authKey:     authKey,
		emailClient: email.NewClient(),
	}
}
