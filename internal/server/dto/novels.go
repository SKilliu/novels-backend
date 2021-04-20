package dto

const PercentSymbol = "%"

type CreateNovelRequest struct {
	Title string `json:"title" example:"My new novel"`
	Data  string `json:"data" example:"My awesome true story!"`
}

type NovelResponse struct {
	ID                        string `json:"id" example:"some_id"`
	Title                     string `json:"title" example:"My new novel"`
	Data                      string `json:"data" example:"My awesome true story!"`
	ParticipatedInCompetition bool   `json:"participatedInCompetition" example:"false"`
	CreatedAt                 int64  `json:"createdAt" example:"121342424"`
	UpdatedAt                 int64  `json:"updatedAt" example:"1654726235"`
}

type UpdateNovelRequest struct {
	ID string `json:"id" example:"some_id"`
	CreateNovelRequest
}
