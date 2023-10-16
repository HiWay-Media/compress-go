package compress


import (
	"github.com/go-resty/resty/v2"
)


type ICompress interface {
	// 
	//
}

type compress struct {
	restClient    *resty.Client
}


func NewCompress() {
	fk := &fakeyou{
		configuration: configuration,
	}
	fk.restClient = resty.New()
}