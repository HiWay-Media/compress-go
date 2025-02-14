package compress

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/minio/minio-go"
)

type ICompress interface {
	//
	HealthCheck() error
	IsDebug() bool
	GetCredentials() (*Credential, error)
	GetUploads(uploadsPaginated UploadsPaginated) ([]VideoUploadInfo, error) // need to change it with args
	GetSingleUpload(jobid int) (*VideoUploadInfo, error)
	GetJobidProgress(jobid int) (*VideoUploadInfo, error)
	SetPublishedUpload(jobid int, published int) (*VideoUploadInfo, error)
	Upload(file []byte, filename string, size int64, categoryId int, title string, tags string, location string) (*ResponseUpload, error)
	//UploadMultipart(reader io.Reader, size int64, categoryId int, title string, tags string, location string, filename string, targetFolder string) (*ResponseUpload, error) // need to add this feature
	GetCategories() ([]Category, error)
	CreateCategory(name string) (*Category, error)
	GetRestreamers(startFrom int, amount int) ([]Restreamer, error)                  // startFrom int, amount int
	GetRestreamersOttAll(startFrom int, amount int) (*RestreamersOTTResponse, error) // startFrom int, amount int
	GetSingleRestreamer(instanceName string) (*Restreamer, error)
	ScaleRestreamer(instanceName string, scale int) (*ResponseServer, error)
	CreateEventsBulk(request []InstancesEventCreate) (*ResponseServer, error)
	RestreamerHlsStart(instanceName string, streamProtocol string) (*HlsResponse, error)
	RestreamerHlsStop(instanceName string, streamProtocol string) (*HlsResponse, error)
	RestreamerEventsHistory(startFrom int, amount int) ([]RestreamerEvent, error)
	GetCustomerS3Zone() (*ResponseCustomerS3, error)
	GenerateVodProxy(eventId string, instanceName string, title string) (*generateVodResponse, error)
	//
}

type compress struct {
	restClient   *resty.Client
	bucket       string
	debug        bool
	customerName string
	clientId     string
	apiKey       string
	customerId   int
}

func NewCompress(customerName, apiKey string, isDebug bool) (ICompress, error) {
	if customerName == "" {
		return nil, fmt.Errorf("customerName is compulsory")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("apiKey is compulsory")
	}
	//customerName := strings.ReplaceAll(clientId, "_client", "")
	clientId := customerName + "_client"
	c := &compress{
		debug:        isDebug,
		restClient:   resty.New(),
		customerName: customerName,
		clientId:     clientId,
		apiKey:       apiKey,
	}
	c.restClient.SetBaseURL(TNGRM_BASE_URL)
	c.restClient.SetDebug(isDebug)
	//
	cred, err := c.GetCredentials()
	if err != nil {
		return nil, err
	}
	c.customerId = cred.CustomerID

	return c, nil
}

func getMinioClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
