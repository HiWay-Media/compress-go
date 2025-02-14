package compress

import (
	"encoding/json"
	"fmt"
)

/*
 */
func (o *compress) GetCredentials() (*Credential, error) {
	//
	resp, err := o.restyPost(CREDENTIALS, BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey})
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	var obj ResponseServerCredential
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	if obj.Response == "KO" {
		return nil, fmt.Errorf("error %s", obj.Message)
	}
	o.debugPrint(obj)
	//
	return &obj.Data, nil
}

/**
* Need to call before upload s3
*
* @returns customer_s3
 */
func (o *compress) GetCustomerS3Zone() (*ResponseCustomerS3, error) {
	//
	resp, err := o.restyPost(GET_ZONE(), BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey})
	if err != nil {
		return nil, err
	}
	//o.debugPrint(resp)
	var obj ResponseCustomerS3
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}

	if obj.Response == "KO" {
		return nil, fmt.Errorf("error %s", obj.Message)
	}
	o.debugPrint(obj)
	//
	return &obj, nil
}
