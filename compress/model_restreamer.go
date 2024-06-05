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
	StartFrom int    `json:"start_from" validate:"nonil" required:"true"`
	Amount    int    `json:"amount"  validate:"nonnil" required:"true"`
}

type restreamerRequest struct {
	BaseModel
	InstanceName string `json:"instance_name"  validate:"nonzero,min=1" required:"true"`
}

type scaleRestreamerRequest struct {
	BaseModel
	InstanceName string `json:"instance_name"  validate:"nonzero,min=1" required:"true"`
	Scale        int    `json:"scale" required:"true"`
}

type InstancesEventCreate struct {
	InstanceName 	string `json:"instance_name"  validate:"nonzero,min=1" required:"true"`
	EventName 		string `json:"event_name"  validate:"nonzero,min=1" required:"true"`
	Protocol 		string `json:"protocol"  validate:"nonzero,min=1" required:"true"`
}

/*type instancesEventCreateRequest struct {
	InstancesEventCreate
	Customer 
}*/

type bulkRestreamerRequest struct {
	BaseModel
	Instances []InstancesEventCreate
}


type HlsResponse struct {
	Message string `json:"message" `
	Result  string `json:"result" `
}

type hlsBodyRequest struct {
	BaseModel
	InstanceName   string `json:"instance_name" `
	StreamProtocol string `json:"stream_protocol" `
}


type startPushRequest struct {
	BaseModel
	InstanceName    string `json:"instance_name" validate:"min=0" required:"true"`
	ExternalServers []struct {
		ExternalServer string `json:"external_server" `
		IngestProtocol string `json:"ingest_protocol" `
		AudioChannel   string `json:"audio_channel" `
	} `json:"external_servers"`
}


type stopPushRequest struct {
	BaseModel
	InstanceName    string `json:"instance_name" `
	ExternalServers []struct {
		ProcessID string `json:"process_id" `
	} `json:"external_servers"`
}


type startPullRequest struct {
	BaseModel
	InstanceName   string `json:"instance_name" `
	ExternalServer string `json:"external_server" `
	IngestProtocol string `json:"ingest_protocol" `
	AudioChannel   string `json:"audio_channel" `
	Type           string `json:"type" `
	Encoding       string `json:"encoding" `
}


type stopPullRequest struct {
	BaseModel
	InstanceName string `json:"instance_name" `
	ProcessID    string `json:"process_id" `
}

// call restreamer api for generate vod from event
type generateVodRequest struct {
	BaseModel
	Category         string  `json:"category" `
	CustomerID       *int    `json:"customer_id" `
	Description      string  `json:"description" `
	EventID          string  `json:"event_id" `
	Instance         string  `json:"instance" `
	Location         string  `json:"location" `
	SkipMin          string  `json:"skip_min" `
	Tags             string  `json:"tags" `
	Title            string  `json:"title" `
	ThumbnailsNumber string  `json:"thumbnails_number" `
	Protocol         string  `json:"protocol" `
}


type eventsHistoryRequest struct {
	BaseModel
	StartFrom int    `json:"start_from" `
	Amount    int    `json:"amount" `
}

type RestreamerEvent struct {
	Id           int        `json:"id" `
	EventId      string     `json:"event_id" `
	StartAt      time.Time  `json:"start_at" gorm:"type:timestamp;"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"type:timestamp;"`
	Active       int        `json:"active" `
	Completed    int        `json:"completed" `
	InstanceName string     `json:"instance_name" `
	CustomerId   int        `json:"customer_id" `
	EventName    string     `json:"event_name" `
	EndAt        *time.Time `json:"end_at" gorm:"type:timestamp;"`
	Protocol     string     `json:"protocol" `
	IsLive       int        `json:"is_live" `
	Deleted      int        `json:"deleted" `
}

type generateVodResponse struct {
	JobId  string  `json:"jobid" `
	Result string  `json:"result" `
	Status *string `json:"status" `
}