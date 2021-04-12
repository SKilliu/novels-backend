package dto

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type LogEventRequest struct {
	UserID   string `json:"user_id"`
	DeviceID string `json:"device_id"`
	Data     Data   `json:"data"`
	Time     int64  `json:"time"`
}

type GetUserEventsResponse struct {
	EventID  string    `json:"event_id"`
	UserID   string    `json:"user_id"`
	DeviceID string    `json:"device_id"`
	Data     Data      `json:"data"`
	Time     time.Time `json:"time"`
}

type Data map[string]interface{}

// We need this for saving and getting JSON from db
func (d Data) Value() (driver.Value, error) {
	return json.Marshal(d)
}

func (d *Data) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &d)
}
