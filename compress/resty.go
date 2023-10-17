package compress

import (
	"fmt"

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
