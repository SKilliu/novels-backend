package models

const ReadyForVoteTableName = "ready_for_vote"

type ReadyForVote struct {
	ID           string `db:"id"`
	UserID       string `db:"user_id"`
	NovelsPoolID string `db:"novels_pool_id"`
	ViewsAmount  int    `db:"views_amount"`
	IsVoted      bool   `db:"is_voted"`
}

// TableName override function from DBX for notice which db relates to provided struct.
func (t ReadyForVote) TableName() string {
	return ReadyForVoteTableName
}
