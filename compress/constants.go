package compress

const (
	TNGRM_BASE_URL 	= "https://api-compress.hiway.media/api/v4.0"
	CATEGORIES     	= "/external/upload/categories"
	RESTREAMERS    	= "/external/restreamers"
	UPLOADS        	= "/external/upload"
	CREDENTIALS    	= "/external/credentials/"
	EVENTS 			= "/external/events"
	CUSTOMERS 		= "/external/customers"
)

var (
	//
	GET_CATEGORIES = func() string {
		return TNGRM_BASE_URL + CATEGORIES
	}
	CREATE_CATEGORY = func() string {
		return TNGRM_BASE_URL + CATEGORIES + "/create"
	}
	CREATE_UPLOAD = func() string {
		return TNGRM_BASE_URL + UPLOADS + "/create"
	}
	GET_RESTREAMERS = func() string {
		return TNGRM_BASE_URL + RESTREAMERS
	}
	GET_RUNNING_INSTANCES = func() string {
		return TNGRM_BASE_URL + RESTREAMERS + "/running_instances"
	}
	GET_RUNNING_SINGLE_INSTANCE = func() string {
		return TNGRM_BASE_URL + RESTREAMERS + "/single_instance"
	}
	SCALE_INSTANCE = func() string {
		return TNGRM_BASE_URL + RESTREAMERS + "/scale_instance"
	}
	BULK_EVENTS_CREATE = func() string {
		return TNGRM_BASE_URL + EVENTS + "/create_bulk"
	}
	RESTREAMER_HLS_START = func() string { // instanceName string, streamProtocol string
		return TNGRM_BASE_URL + RESTREAMERS + "/hls/start"
	}
	RESTREAMER_HLS_STOP = func() string { // instanceName string, streamProtocol string
		return TNGRM_BASE_URL + RESTREAMERS + "/hls/stop"
	}
	RESTREAMER_PUSH_START = func() string { // instanceName string, streamProtocol string
		return TNGRM_BASE_URL + RESTREAMERS + "/push/start"
	}
	RESTREAMER_PUSH_STOP = func() string { // instanceName string, streamProtocol string
		return TNGRM_BASE_URL + RESTREAMERS + "/push/stop"
	}
	RESTREAMER_PULL_START = func() string { // instanceName string, streamProtocol string
		return TNGRM_BASE_URL + RESTREAMERS + "/pull/start"
	}
	RESTREAMER_PULL_STOP = func() string { // instanceName string, streamProtocol string
		return TNGRM_BASE_URL + RESTREAMERS + "/pull/stop"
	}
	GET_UPLOADS = func() string {
		return TNGRM_BASE_URL + UPLOADS
	}
	GET_SINGLE_UPLOAD = func(jobId int) string {
		return TNGRM_BASE_URL + UPLOADS + "/jobid"
	}
	GET_JOBID_PROGRESS = func(jobId int) string {
		return TNGRM_BASE_URL + UPLOADS + "/job_progress/jobid"
	}
	SET_PUBLISHED_UPLOAD = func(jobId int) string {
		return TNGRM_BASE_URL + UPLOADS + "/set_published"
	}
	PRESIGNED_URL_S3 = func() string {
		return TNGRM_BASE_URL + UPLOADS + "/presignedUrl"
	}
	GET_ZONE = func() string {
		return TNGRM_BASE_URL + CUSTOMERS + "/s3"
	}
	//
)
