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
	GET_RESTREAMERS = func() string {
		return TNGRM_BASE_URL+RESTREAMERS
	}
	GET_UPLOADS = func() string {
		return TNGRM_BASE_URL+UPLOADS
	}
	//
)