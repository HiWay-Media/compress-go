package compress

import "time"

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


// VideoUploadInfo single video upload representation for response, made with Upload, AuthUser and Category tables
type VideoUploadInfo struct {
	Id                  int         `json:"id" `
	Title               string      `json:"title" `
	Location            string      `json:"location"`
	ReporterFirstName   string      `json:"reporter_first_name" `
	ReporterLastName    string      `json:"reporter_last_name" `
	Created             string      `json:"created" `
	CategoryName        string      `json:"category_name" `
	Token               string      `json:"token" `
	JobId               int         `json:"job_id" ` // changed with job_id
	DownloadUrl         string      `json:"download_url" `
	RemoteUrl           string      `json:"remote_url" `
	Thumbnail           string      `json:"thumbnail" `
	Thumbnails          string      `json:"thumbnails" `
	ThumbnailsV2        string      `json:"thumbnailsV2" gorm:"column:thumbnailsV2"` //TODO to deprecate with new models --> preview_urls
	WatermarkUrl        string      `json:"watermark_url" `
	Tags                string      `json:"tags" `
	ExtendedData        string      `json:"extended_data" `
	OriginalUrl         string      `json:"original_url" `
	Duration            int         `json:"duration" `
	DurationMillisecond int         `json:"duration_msec" gorm:"column:duration_msec"`
	Size                int64       `json:"size" `
	Published           int         `json:"published" `
	Width               int         `json:"width"`
	Height              int         `json:"height"`
	DisplayAspectRatio  string      `json:"display_aspect_ratio"`
	Is360               int         `json:"360degree"  gorm:"column:360degree"`
	ClientId            string      `json:"client_id" `
	DeletedAt           *time.Time  `json:"deleted_at,omitempty" gorm:"type:timestamp;"`
	MiniClips           interface{} `json:"mini_clips" gorm:"-"`
}

type jobidProgressRequest struct {
	BaseModel
	JobId int `json:"job_id" validate:"nonzero,min=1" required:"true"`
}