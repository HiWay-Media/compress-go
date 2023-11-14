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

func (o *compress) GetCategories(requestBody categoriesRequest) ([]Category, error) {
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
func (o *compress) CreateCategory( requestBody createCategoryRequest ) (*Category, error) {
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(CREATE_CATEGORY(requestBody.CategoryName), requestBody)
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
