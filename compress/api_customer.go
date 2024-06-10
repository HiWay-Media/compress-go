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
func (o *compress) GetCustomerS3Zone() (*CustomerS3, error) {
	//
	queryParam := make(map[string]string)
	queryParam["client_id"] = o.GetCliendId()
	queryParam["api_key"] = o.apiKey
	//
	resp, err := o.restyGet(GET_ZONE(), queryParam)
	if err != nil {
		return nil, err
	}
	//o.debugPrint(resp)
	var obj CustomerS3
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	o.debugPrint(obj)
	//
	return &obj, nil
}
