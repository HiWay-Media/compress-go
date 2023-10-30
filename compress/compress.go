package compress

import (
	"github.com/go-resty/resty/v2"
)
//
type ICompress interface {
	// 
	HealthCheck() error
	IsDebug() bool
	GetUploads(uploadsPaginated UploadsPaginated) error
	//
}
//
type compress struct {
	restClient    	*resty.Client
	debug 			bool
	customerName 	string
	apiKey 			string
}
//
func NewCompress(customerName, apiKey string,isDebug bool) ICompress{
	c := &compress{
		debug: isDebug,
		restClient: resty.New(),
		customerName: customerName,
		apiKey: apiKey,
	}
	c.restClient.SetBaseURL(TNGRM_BASE_URL)
	c.restClient.SetDebug(isDebug)
	//
	return c
}