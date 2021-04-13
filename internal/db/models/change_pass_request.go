package models

const ChangePasswordRequestsTableName = "change_password_requests"

type ChangePassRequest struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t ChangePassRequest) TableName() string {
	return ChangePasswordRequestsTableName
}
