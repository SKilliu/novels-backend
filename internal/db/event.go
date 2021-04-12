package db

import (
	"database/sql"
	"time"

	"github.com/SKilliu/novels-backend/internal/server/dto"
)

type Event struct {
	ID       string    `db:"id"`
	UserID   string    `db:"user_id"`
	DeviceID string    `db:"device_id"`
	Data     dto.Data  `db:"data"`
	Time     time.Time `db:"time"`
}

func InsertEvent(tx *sql.Tx, event Event) error {
	_, err := tx.Exec(`INSERT INTO user_events (id, user_id, device_id, data, time) VALUES ($1, $2, $3, $4, $5)`,
		event.ID, event.UserID, event.DeviceID, event.Data, event.Time)

	if err != nil {
		return err
	}

	return err
}

func GetEventsByUserID(tx *sql.Tx, userid string) (events []Event, err error) {
	var (
		rows  *sql.Rows
		event Event
	)

	rows, err = tx.Query(`SELECT * FROM user_events WHERE user_id = $1`, userid)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&event.ID, &event.UserID, &event.DeviceID, &event.Data, &event.Time)
		if err != nil {
			return
		}
		events = append(events, event)
	}

	return
}
