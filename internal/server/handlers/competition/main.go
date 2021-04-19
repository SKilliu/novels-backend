package competition

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log           *logrus.Entry
	usersDB       db.UsersQ
	novelsDB      db.NovelsQ
	competitionDB db.CompetitionsQ
	authKey       string
}

func New(logger *logrus.Entry, authKey string) *Handler {
	return &Handler{
		log:      logger,
		usersDB:  db.Connection().UsersQ(),
		novelsDB: db.Connection().NovelsQ(),
		authKey:  authKey,
	}
}
