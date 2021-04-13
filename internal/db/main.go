package db

import (
	"fmt"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

var dbConn QInterface

type QInterface interface {
	DBX() *dbx.DB

	UsersQ() UsersQ
	ChangePassRequestsQ() ChangePassRequestsQ
}

// DB wraps dbx interface.
type DB struct {
	db *dbx.DB
}

// DBX returns db client.
func (d DB) DBX() *dbx.DB {
	return d.db
}

// New connection opening.
func New(link string) (QInterface, error) {
	db, err := dbx.Open("postgres", link)

	logger.Info("Database connection succesfully opened")

	return &DB{db: db}, err
}

func Init(log *logrus.Entry) {
	setLogger(log)

	loadConfigFromEnvs()

	connect, err := New(configuration.Info())
	if err != nil {
		logger.WithError(err).Error("failed to setup db")
		panic(err)
	}

	dbConn = connect
}

func (d Configuration) Info() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSL,
	)
}

func Connection() QInterface {
	return dbConn
}
