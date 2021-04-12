package models

import "time"

// UsersTableName table name in db.
const UsersTableName = "users"

// User entity in db.
type User struct {
	ID             string    `db:"id"`
	Username       string    `db:"username"`
	HashedPassword string    `db:"hashed_password"`
	Email          string    `db:"email"`
	DateOfBirth    time.Time `db:"date_of_birth"`
	Gender         string    `db:"gender"`
	Membership     string    `db:"membership"`
	AvatarData     string    `db:"avatar_data"`
	DeviceID       string    `db:"device_id"`
	Rate           int       `db:"rate"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t User) TableName() string {
	return UsersTableName
}
