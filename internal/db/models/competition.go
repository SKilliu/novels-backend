package models

const CompetitionsTableName = "novels_pool"

type Competition struct {
	ID                   string `db:"id"`
	NovelOneID           string `db:"novel_one_id"`
	NOvelTwoID           string `db:"novel_two_id"`
	CompatitionStartedAt uint64 `db:"competitionStartedAt"`
	Status               string `db:"status"`
	CreatedAt            int64  `db:"createdAt" example:"121342424"`
	UpdatedAt            int64  `db:"updatedAt" example:"1654726235"`
}

type CompetitionOpponent struct {
	UserID string
	
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t Competition) TableName() string {
	return CompetitionsTableName
}
