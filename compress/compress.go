package compress

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/minio/minio-go"
	"io"
	"net/url"
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
	Upload(file []byte, size int64, categoryId int, title string, tags string, location string, filename string, targetFolder string) (*ResponseUpload, error)
	UploadMultipart(reader io.Reader, size int64, categoryId int, title string, tags string, location string, filename string, targetFolder string) (*ResponseUpload, error)
	GetCategories() ([]Category, error)
	CreateCategory(name string) (*Category, error)
	GetRestreamers() ([]Restreamer, error) // startFrom int, amount int
	GetSingleRestreamer(instanceName string) (*Restreamer, error)
	ScaleRestreamer(instanceName string, scale int) (*ResponseServer, error)
	CreateEventsBulk(request []InstancesEventCreate) (*ResponseServer, error)
	//
}

type compress struct {
	restClient   *resty.Client
	minioClient  *minio.Client
	bucket       string
	debug        bool
	customerName string
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
	c := &compress{
		debug:        isDebug,
		restClient:   resty.New(),
		customerName: customerName,
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

	u, err := url.Parse(cred.S3Host)
	if err != nil {
		return nil, err
	}
	c.minioClient, err = getMinioClient(
		u.Host,
		cred.S3AccessKey,
		cred.S3SecretKey,
		false,
	)
	c.bucket = cred.S3Bucket
	if err != nil {
		return nil, err
	}
	return c, nil
}

func getMinioClient(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}
