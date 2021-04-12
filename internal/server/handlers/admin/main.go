package admin

import (
	"database/sql"

	"github.com/SKilliu/users-rest-api/internal/db"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	log     *logrus.Entry
	db      *sql.DB
	authKey string
}

func New(logger *logrus.Entry, authKey string) *Handler {
	return &Handler{
		log:     logger,
		db:      db.Connection(),
		authKey: authKey,
	}
}
