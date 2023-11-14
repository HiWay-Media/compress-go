package compress 

import "time"

type Restreamer struct {
	Id                   int                 `json:"id" `
	Enabled              int                 `json:"enabled" `
	Name                 string              `json:"name" `
	CustomerId           int                 `json:"customer_id" `
	Dnsname              string              `json:"dnsname" `
	MarathonUrl          string              `json:"marathon_url" `
	NpmUrl               string              `json:"npm_url" `
	ExternalRtmp         string              `json:"external_rtmp" `
	LoopbackRtmp         string              `json:"loopback_rtmp" `
	Md5Generator         *string             `json:"md5_generator" gorm:"column:md5-generator"`
	MarathonPath         string              `json:"marathon_path" `
	Token                string              `json:"token" `
	DockerTemplate       string              `json:"docker_template" `
	Status               int                 `json:"status" `
	Owner                string              `json:"owner" `
	Title                string              `json:"title" `
	Description          string              `json:"description" `
	OwnerId              int                 `json:"owner_id" `
	CdnPath              string              `json:"cdn_path" `
	ExtendedData         string              `json:"extended_data" `
	LatestEvent          string              `json:"latest_event" `
	Label                string              `json:"label" `
	Dedicated            string              `json:"dedicated" `
	StaticIp             string              `json:"static_ip" `
	NvidiaVisibleDevices string              `json:"nvidia_visible_devices" `
	Transcode            int                 `json:"transcode" `
	UpdateAt             time.Time           `json:"update_at" `
	PassiveToken         string              `json:"passive_token" `
	Zone                 string              `json:"zone"`
	Protocol             string              `json:"protocol"`
	AWSUrl               string              `json:"aws_url"`
	RestreamerSettings   *RestreamerSettings `json:"settings" gorm:"foreignKey:id"`
	RestreamerSrt        *RestreamerSrt      `json:"srt" gorm:"foreignKey:restreamer_id"`
}


type RestreamerSettings struct {
	Id              int    `json:"id" `
	EnableEncoding  int    `json:"enable_encoding" `
	EncodingFormat  string `json:"encoding_format" `
	Encoding1080    int    `json:"1080_encoding" gorm:"column:1080_encoding"`
	EnableLive24    int    `json:"enable_live24" `
	ResetTime       int    `json:"reset_time" `
	Bookable        bool   `json:"bookable" `
	AutoStart       bool   `json:"auto_start" `
	AudioFilter     bool   `json:"audio_filter" `
	FPS50           bool   `json:"50_fps" gorm:"column:50_fps"`
	DVR             bool   `json:"dvr" `
	Enable360Degree bool   `json:"enable_360degree" gorm:"column:enable_360degree"`
	EnableDRM       bool   `json:"enable_drm" gorm:"column:enable_drm"`
	// added enable drm
}

type RestreamerSrt struct {
	Id              int     `json:"id" `
	RestreamerId    int64   `json:"restreamer_id"`
	Port            int     `json:"port" `
	ExternalSrt     string  `json:"external_srt" `
	Token           string  `json:"token" `
	CdnPath         string  `json:"cdn_path" `
	Owner           string  `json:"owner" `
	LatestEvent     string  `json:"latest_event" `
	InternalIp      *string `json:"internal_ip" `
	Mode            string  `json:"mode" `
	TranscodingType *string `json:"transcoding_type" `
	StreamId        string  `json:"streamid" gorm:"column:streamid"`
	Enabled         int     `json:"enabled" `
	PassiveToken    string  `json:"passive_token" `
}


type findRestreamersRequest struct {
	BaseModel
	StartFrom int    `json:"start_from"  validate:"nonil,min=0" required:"true"`
	Amount    int    `json:"amount"  validate:"nonnil,min=0" required:"true"`
}