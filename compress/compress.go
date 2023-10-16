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


func NewCompress() ICompress{
	c := &compress{
	}
	c.restClient = resty.New()
	return c
}