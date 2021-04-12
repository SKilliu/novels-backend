package models

// UsersTableName table name in db.
const UsersTableName = "users"

// User entity in db.
type User struct {
	ID             string `db:"id"`
	Name           string `db:"name"`
	HashedPassword string `db:"hashed_password"`
	Email          string `db:"email"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t User) TableName() string {
	return UsersTableName
}
