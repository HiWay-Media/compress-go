package compress

import (
	"fmt"
	"gopkg.in/validator.v2"
)

/**
*
* @param {string} apikey
* @param {string} customer
* @returns list of categories of the customer
 */

func (o *compress) GetCategories() ([]Category, error) {
	requestBody := &categoriesRequest{
		BaseModel: BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
	}
	//
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_CATEGORIES(), requestBody)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("")
	}
	//
	return nil, nil
}

/**
*
* @param {string} apikey
* @param {string} customer
* @returns list of categories of the customer
 */
func (o *compress) CreateCategory(name string) (*Category, error) {
	requestBody := &createCategoryRequest{
		BaseModel:    BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		CategoryName: name,
	}
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(CREATE_CATEGORY(), requestBody)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("")
	}
	//
	return nil, nil
}
