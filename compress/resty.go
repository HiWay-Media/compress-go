package compress

import (
	"fmt"
	"log"
	"github.com/go-resty/resty/v2"
)

func (o *compress) HealthCheck() error {
	resp, err := o.restyPost("/health", nil)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("")
	}
	return nil
}

func (o *compress) IsDebug() bool {
	return o.debug
}

func(o *compress) GetCliendId() string {
	return o.customerName+"_client"
}

// Resty Methods

func (o *compress) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *compress) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}


func (o *compress) debugPrint(data interface{}) {
	if o.debug {
		log.Println(data)
	}
}

func castPax[T any](body []byte) (*T, error) {
	var obj T
	if err := json.Unmarshal(body, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}
