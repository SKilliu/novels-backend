package models

const ResetPasswordRequestsTableName = "reset_password_requests"

type ResetPassRequest struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t ResetPassRequest) TableName() string {
	return ResetPasswordRequestsTableName
}
