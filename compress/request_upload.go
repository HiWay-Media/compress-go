package compress

type UploadsPaginated struct {
	BaseModel
	StartFrom   int     `json:"start_from"`
	Amount      int     `json:"amount"`
	TitleFilter *string `json:"q"`
	Category    *string `json:"category"`
	Tags        *string `json:"tags"`
	FromDate    *string `json:"from_date", omitempty `
	ToDate      *string `json:"to_date", omitempty `
}