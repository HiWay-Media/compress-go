package compress

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