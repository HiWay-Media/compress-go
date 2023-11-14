package compress

import (
	"fmt"
	"encoding/json"
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
func (o *compress) GetUploads(uploadsPaginated UploadsPaginated) ([]VideoUploadInfo, error) {
	//
	if errs := validator.Validate(uploadsPaginated); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_UPLOADS(), nil)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("")
	}
	var obj []VideoUploadInfo
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	o.debugPrint(obj)
	return obj, nil
}

/**
* jobid is compulsory
* example: get_single_upload(1000)
* @param {string} api_key 
* @param {number} jobid 
* @returns upload list
*/

func (o *compress) GetSingleUpload(requestBody jobidProgressRequest) (*VideoUploadInfo , error){
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_SINGLE_UPLOAD(requestBody.JobId), requestBody)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("")
	}
	var obj VideoUploadInfo
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}
/**
* 
* @param {number} job_id 
* @returns progressStateResponse 
*/
func (o *compress) GetJobidProgress(requestBody jobidProgressRequest) (*VideoUploadInfo , error) {
	//
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(GET_JOBID_PROGRESS(requestBody.JobId), requestBody)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("")
	}
	var obj VideoUploadInfo
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

/**
* jobid is compulsory
* example: set_published_upload(1000)
* @param {string} api_key 
* @param {number} jobid 
*/
func (o *compress) SetPublishedUpload(requestBody publishedUploadRequest) (*VideoUploadInfo , error) {
	//
	if errs := validator.Validate(requestBody); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	resp, err := o.restyPost(SET_PUBLISHED_UPLOAD(requestBody.JobId), requestBody)
	if err != nil {
		return nil, err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return nil, fmt.Errorf("")
	}
	var obj VideoUploadInfo
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

/**
* upload video to minio s3 bucket with a presigned PUT url
*
* videos will not be displayed in compress platform,
*
* this is just a plain upload to s3 storage
* @param {string} destination_folder
* @param {string} filename
* @param {file} file
*/
func (o *compress) UploadS3( destinationFolder string, filename string, ) error {
	var fileDest = destinationFolder + "/" + filename
	o.debugPrint(fileDest)
	resp, err := o.restyPost(PRESIGNED_URL_S3, presignedObject{
		CustomerName: o.customerName,
		FileName: fileDest
	})
	if err != nil {
		return err
	}
	o.debugPrint(resp)
	if resp.IsError() {
		return fmt.Errorf("")
	}
	//
	return nil
}