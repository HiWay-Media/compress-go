package compress

import (
	"fmt"
	"gopkg.in/validator.v2"
)

/**
* upload video with compress encoding
*
* if destination folder is empty, it will upload to the root of the bucket
*
* remember to specify the folder (usually upload)
*
* @param {string} destination_folder
* @param {file} file
* @param {string} title
* @param {string} tags
* @param {string} location_place
* @param {number} category_id
 */
func (o *compress) GetUploads(uploadsPaginated UploadsPaginated) error {
	//
	if errs := validator.Validate(uploadsPaginated); errs != nil {
		// values not valid, deal with errors here
		return errs
	}
	resp, err := o.restyPost(GET_UPLOADS(), nil)
	if err != nil {
		return err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return fmt.Errorf("")
	}
	return nil
}
