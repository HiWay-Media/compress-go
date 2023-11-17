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
	var obj ResponseServer
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	if obj.Data == "KO" {
		return nil, fmt.Errorf("Error %s", obj.Message)
	}
	o.debugPrint(obj)
	//
	var cred Credential
	if err := json.Unmarshal(obj.Data, &cred); err != nil {
		return nil, err
	}
	return &cred, nil
}
