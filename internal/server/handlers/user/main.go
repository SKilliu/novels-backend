package user

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log     *logrus.Entry
	usersDB db.UsersQ
	authKey string
}

func New(logger *logrus.Entry, authKey string) *Handler {
	return &Handler{
		log:     logger,
		usersDB: db.Connection().UsersQ(),
		authKey: authKey,
	}
}
