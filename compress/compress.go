package compress

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ICompress interface {
	//
	HealthCheck() error
	IsDebug() bool
	GetCredentials() (*ResponseServer, error)
	GetUploads(uploadsPaginated UploadsPaginated) ([]VideoUploadInfo, error) // need to change it with args
	GetSingleUpload(jobid int) (*VideoUploadInfo, error)
	GetJobidProgress(jobid int) (*VideoUploadInfo, error)
	SetPublishedUpload(jobid, published int) (*VideoUploadInfo, error)
	UploadS3(destinationFolder string, filename string) error
	GetCategories() ([]Category, error)
	CreateCategory(name string) (*Category, error)
	GetRestreamers(startFrom, amount int) ([]Restreamer, error)
	GetSingleRestreamer(instanceName string) (*Restreamer, error)
	//
}

type compress struct {
	restClient   *resty.Client
	debug        bool
	customerName string
	apiKey       string
}

func NewCompress(customerName, apiKey string, isDebug bool) (ICompress, error) {
	if customerName == "" {
		return nil, fmt.Errorf("customerName is compulsory")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("apiKey is compulsory")
	}
	c := &compress{
		debug:        isDebug,
		restClient:   resty.New(),
		customerName: customerName,
		apiKey:       apiKey,
	}
	c.restClient.SetBaseURL(TNGRM_BASE_URL)
	c.restClient.SetDebug(isDebug)
	//
	_, err := c.GetCredentials()
	if err != nil {
		return nil, err
	}
	//
	return c, nil
}

//
