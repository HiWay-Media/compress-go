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
		return nil, fmt.Errorf("Error %s", obj.Message)
	}
	o.debugPrint(obj)
	//
	return &obj.Data, nil
}
