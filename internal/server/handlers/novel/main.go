package novel

import (
	"github.com/SKilliu/novels-backend/internal/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log            *logrus.Entry
	usersDB        db.UsersQ
	novelsDB       db.NovelsQ
	competitionsDB db.CompetitionsQ
	readyForVoteDB db.ReadyForVoteQ
	authKey        string
}

func New(logger *logrus.Entry, authKey string) *Handler {
	return &Handler{
		log:            logger,
		usersDB:        db.Connection().UsersQ(),
		novelsDB:       db.Connection().NovelsQ(),
		competitionsDB: db.Connection().CompetitionsQ(),
		readyForVoteDB: db.Connection().ReadyForVoteQ(),
		authKey:        authKey,
	}
}
