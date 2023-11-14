package compress

const (
	TNGRM_BASE_URL                 	= "https://api.tngrm.io/api/v3.0"
	CATEGORIES                 	   	= "/external/upload/categories"
	RESTREAMERS                		= "/external/restreamers"
	UPLOADS                    		= "/external/upload"
)
//
var (
	//
	GET_CATEGORIES = func() string {
		return TNGRM_BASE_URL+CATEGORIES
	}
	CREATE_CATEGORY = func(name string) string{
		return TNGRM_BASE_URL+CATEGORIES + "/create"
	}
	GET_RESTREAMERS = func() string {
		return TNGRM_BASE_URL+RESTREAMERS
	}
	GET_RUNNING_INSTANCES = func() string {
		return TNGRM_BASE_URL+RESTREAMERS+"/running_instances"
	}
	GET_RUNNING_SINGLE_INSTANCE = func() string {
		return TNGRM_BASE_URL+RESTREAMERS+"/single_instance"
	}
	GET_UPLOADS = func() string {
		return TNGRM_BASE_URL+UPLOADS
	}
	GET_SINGLE_UPLOAD = func(jobId int ) string {
		return TNGRM_BASE_URL+UPLOADS+"/jobid"
	}
	GET_JOBID_PROGRESS = func(jobId int) string {
		return TNGRM_BASE_URL+UPLOADS+"job_progress/jobid"
	}
	//
)