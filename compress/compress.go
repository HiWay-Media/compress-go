package compress


import (
	"github.com/go-resty/resty/v2"
)


type ICompress interface {
	// 
	HealthCheck() error
	//
}

type compress struct {
	restClient    *resty.Client
	debug bool
}


func NewCompress(isDebug bool) ICompress{
	c := &compress{
		debug: isDebug,
		restClient: resty.New(),
	}
	c.restClient.SetBaseURL(TNGRM_BASE_URL)
	//
	return c
}