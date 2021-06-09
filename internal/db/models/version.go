package models

// UsersTableName table name in db.
const VersionsTableName = "versions"

// User entity in db.
type Versions struct {
	ID      string `db:"id"`
	Android string `db:"android"`
	Ios     string `db:"ios"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t Versions) TableName() string {
	return VersionsTableName
}
