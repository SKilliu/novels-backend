package models

const UserSocialsTableName = "user_socials"

type UserSocial struct {
	ID       string `db:"id"`
	UserID   string `db:"user_id"`
	Social   string `db:"social"`
	SocialID string `db:"social_id"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t UserSocial) TableName() string {
	return UserSocialsTableName
}
