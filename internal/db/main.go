package db

import (
	"fmt"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var dbConn QInterface

type QInterface interface {
	DBX() *dbx.DB

	UsersQ() UsersQ
}

// // New connection opening.
// func NewConnection(config string) {
// 	db, err := sql.Open("postgres", config)
// 	if err != nil {
// 		logger.WithError(err).Error("failed to open db connection")
// 		panic(err)
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		logger.WithError(err).Error("failed to ping db")
// 		panic(err)
// 	}

// 	logger.Info("db connection successfully created")

// 	dbConn = db
// }

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
