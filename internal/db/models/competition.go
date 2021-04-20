package models

const CompetitionsTableName = "novels_pool"

type Competition struct {
	ID                   string `db:"id"`
	NovelOneID           string `db:"novel_one_id"`
	NovelTwoID           string `db:"novel_two_id"`
	UserOneID            string `db:"user_one_id"`
	UserTwoID            string `db:"user_two_id"`
	CompetitionStartedAt int64  `db:"competition_started_at"`
	Status               string `db:"status"`
	NovelOneVotes        int    `db:"novel_one_votes"`
	NovelTwoVotes        int    `db:"novel_two_votes"`
	CreatedAt            int64  `db:"created_at" example:"121342424"`
	UpdatedAt            int64  `db:"updated_at" example:"1654726235"`
}

type CompetitionOpponent struct {
	UserID string
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t Competition) TableName() string {
	return CompetitionsTableName
}
