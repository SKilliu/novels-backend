package models

const NovelsTableName = "novels"

type Novel struct {
	ID                        string `db:"id"`
	UserID                    string `db:"user_id"`
	Title                     string `db:"title"`
	Data                      string `db:"data"`
	ParticipatedInCompetition bool   `db:"participated_in_competiton"`
	VotingResult              int32  `db:"voting_result"`
	UpdatedAt                 int64  `db:"updated_at"`
	CreatedAt                 int64  `db:"created_at"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t Novel) TableName() string {
	return NovelsTableName
}
