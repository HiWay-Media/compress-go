package compress

import "time"
/*
*/
type Credential struct {
	ID                       int    `json:"id"`
	CustomerID               int    `json:"customer_id" `
	S3Host                   string `json:"s3_host" `
	S3Bucket                 string `json:"s3_bucket" `
	S3AccessKey              string `json:"s3_access_key" `
	S3SecretKey              string `json:"s3_secret_key" `
	SftpUsername             string `json:"sftp_username" `
	SftpPassword             string `json:"sftp_password" `
	SftpRemoteHost           string `json:"sftp_remote_host" `
	SftpRemotePort           int    `json:"sftp_remote_port" `
	SftpBasePath             string `json:"sftp_base_path" `
	Protocol                 string `json:"protocol" `
	EndpointVisualization    string `json:"endpoint_visualization" `
	NfsMount                 string `json:"nfs_mount" `
	VodesBaseURL             string `json:"vodes_base_url" `
	LiveBaseURL              string `json:"live_base_url" `
	LiveBaseURLDedicated     string `json:"live_base_url_dedicated" `
	ExternalClusters         string `json:"external_clusters" `
	RestreamerHlsBasePath    string `json:"restreamer_hls_base_path" `
	RestreamerHlstmpBasePath string `json:"restreamer_hlstmp_base_path" `
	ClientID                 string `json:"client_id" `
	SSDMount                 string `json:"ssd_mount" `
	S3mount                  string `json:"s3_mount" `
	SlaveMount               int    `json:"slave_mount"`
	LumenS3Host              string `json:"lumen_s3_host" `
	LumenS3Bucket            string `json:"lumen_s3_bucket" `
	LumenS3AccessKey         string `json:"lumen_s3_access_key" `
	LumenS3SecretKey         string `json:"lumen_s3_secret_key" `
}


type ResponseServerCredential struct {
	Message  string     `json:"message"`
	Response string     `json:"response"`
	Data     Credential `json:"data"`
}


type CustomerS3 struct {
	ID                     int        `json:"id" `
	CustomerID             int        `json:"customer_id" gorm:"column:customer_id"`
	S3Name                 string     `json:"s3_name" `
	Host                   string     `json:"host" `
	Bucket                 string     `json:"bucket" `
	AccessKey              string     `json:"access_key" `
	SecretKey              string     `json:"secret_key" `
	IsActive               int        `json:"is_active" `
	UpdateEnabled          int        `json:"update_enabled" `
	CreatedAt              *time.Time `json:"created_at" gorm:"type:timestamp;"`
	UpdatedAt              *time.Time `json:"updated_at" gorm:"type:timestamp;"`
	DeletedAt              *time.Time `json:"deleted_at" gorm:"type:timestamp;"`
	LiveNFSMount           string     `json:"live_nfs_mount"`
	VodNFSMount            string     `json:"vod_nfs_mount"`
	VodCDNBaseURL          string     `json:"vod_cdn_base_url"`
	LiveCDNBaseURL         string     `json:"live_cdn_base_url"`
	LiveUnprotectedBaseURL string     `json:"live_unprotected_base_url"`
	VodUnprotectedBaseURL  string     `json:"vod_unprotected_base_url"`
	Zone                   string     `json:"zone"`
	BucketUpload           string     `json:"bucket_upload"`
}