package dto

const (
	StatusWaitingForOpponent = "waiting_for_opponent"
	StatusStarted            = "started"
	StatusExpired            = "expired"
	StatusFinished           = "finished"
	TimeForAction            = 2 // 2 hours for opponent finding
)

type CreateCompetitionReqeust struct {
	NovelID string `json:"novel_id"`
}

type CompetitionResponse struct {
	ID                   string    `json:"id"`
	NovelOne             NovelData `json:"novel1"`
	NovelTwo             NovelData `json:"novel2"`
	CompatitionStartedAt uint64    `json:"competitionStartedAt"`
	Status               string    `json:"status"`
	CreatedAt            int64     `json:"createdAt" example:"121342424"`
	UpdatedAt            int64     `json:"updatedAt" example:"1654726235"`
}

type NovelData struct {
	NovelResponse
	User UserData `json:"user"`
}

type UserData struct {
	Username    string `json:"username" example:"awesome_user"`
	DateOfBirth int64  `json:"dateOfBith" example:"12345672"`
	Gender      string `json:"gender" example:"male"`
	Membership  string `json:"membership" example:"some_info"`
	Rate        int    `json:"rate" example:"0"`
}
