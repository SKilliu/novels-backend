package db

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var dbConn *sql.DB

// New connection opening.
func NewConnection(config string) {
	db, err := sql.Open("postgres", config)
	if err != nil {
		logger.WithError(err).Error("failed to open db connection")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		logger.WithError(err).Error("failed to ping db")
		panic(err)
	}

	logger.Info("db connection successfully created")

	dbConn = db
}

func Init(log *logrus.Entry) {
	setLogger(log)

	loadConfigFromEnvs()

	NewConnection(configuration.Info())

}

func (d Configuration) Info() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.Host, d.Port, d.User, d.Password, d.Name, d.SSL,
	)
}

func Connection() *sql.DB {
	return dbConn
}
