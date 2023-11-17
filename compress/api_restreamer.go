package compress

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

/**
*
* @param {number} start_from
* @param {number} amount
* @returns restreamer list
*/
func (o *compress) GetRestreamers() ([]Restreamer, error) {
	requestBody := BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey}
	/*&findRestreamersRequest{
		BaseModel: BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		StartFrom: startFrom,
		Amount:    amount,
	}*/
	//
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_RESTREAMERS(), requestBody)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("restreamers error")
	}
	var obj []Restreamer
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	o.debugPrint(obj)
	return obj, nil
}

/**
*
* @returns restreamer object
 */
func (o *compress) GetSingleRestreamer(instanceName string) (*Restreamer, error) {
	//
	requestBody := &restreamerRequest{
		BaseModel:    BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		InstanceName: instanceName,
	}
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_RUNNING_SINGLE_INSTANCE(), requestBody)
	if err != nil {
		return nil, err
	}
	var obj Restreamer
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	o.debugPrint(obj)
	return &obj, nil
}

/*

*/
func (o *compress) ScaleRestreamer(instanceName string, scale int) (*ResponseServer, error){
	// 
	requestBody := &scaleRestreamerRequest{
		BaseModel:    BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		InstanceName: instanceName,
		Scale: scale,
	}
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_RUNNING_SINGLE_INSTANCE(), requestBody)
	if err != nil {
		return nil, err
	}
	var obj ResponseServer
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	if obj.Data == "KO" {
		return nil, fmt.Errorf("Error %s", obj.Message)
	}
	o.debugPrint(obj)
	return &obj, nil
	//
}

/*

*/
func (o *compress) CreateEventsBulk(request []InstancesEventCreate) (*ResponseServer, error) {
	requestBody := &bulkRestreamerRequest{
		BaseModel: BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		Instances: request,
	}
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(BULK_EVENTS_CREATE(), requestBody)
	if err != nil {
		return nil, err
	}
	var obj ResponseServer
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	if obj.Data == "KO" {
		return nil, fmt.Errorf("Error %s", obj.Message)
	}
	o.debugPrint(obj)
	return &obj, nil
}