package user

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log                  *logrus.Entry
	usersDB              db.UsersQ
	changePassRequestsDB db.ResetPassRequestsQ
	userSocialsDB        db.UserSocialsQ
	authKey              string
	email                string
	password             string
}

func New(logger *logrus.Entry, authKey, email, password string) *Handler {
	return &Handler{
		log:                  logger,
		usersDB:              db.Connection().UsersQ(),
		changePassRequestsDB: db.Connection().ResetPassRequestsQ(),
		userSocialsDB:        db.Connection().UserSocialsQ(),
		authKey:              authKey,
		email:                email,
		password:             password,
	}
}
