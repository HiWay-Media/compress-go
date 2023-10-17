package compress

import (
	"fmt"
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
func(o *compress) GetUploads(startFrom int, amount int, title *string, categoryName *string, tags *string) error{
	resp, err := o.restyPost(GET_UPLOADS(), nil)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return fmt.Errorf("")
	}
	return nil
}