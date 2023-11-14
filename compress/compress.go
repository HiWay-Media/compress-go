package compress

import (
	"github.com/go-resty/resty/v2"
	"fmt"
)
//
type ICompress interface {
	// 
	HealthCheck() error
	IsDebug() bool
	GetUploads(uploadsPaginated UploadsPaginated) ([]VideoUploadInfo, error)
	GetSingleUpload(requestBody jobidProgressRequest) (*VideoUploadInfo , error)
	GetJobidProgress(requestBody jobidProgressRequest) (*VideoUploadInfo , error)
	SetPublishedUpload(requestBody publishedUploadRequest) (*VideoUploadInfo , error)
	GetCategories(requestBody categoriesRequest) ([]Category, error)
	CreateCategory( requestBody createCategoryRequest ) (*Category, error)
	GetRestreamers( requestBody findRestreamersRequest ) ([]Restreamer, error)
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
func NewCompress(customerName, apiKey string,isDebug bool) ( ICompress, error ){
	if customerName == "" {
		return nil, fmt.Errof("customerName is compulsory")
	}
	if apiKey == "" {
		return nil, fmt.Errof("apiKey is compulsory")
	}
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
//