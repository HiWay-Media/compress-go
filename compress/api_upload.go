package compress

import (
	"encoding/json"
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

func (o *compress) GetSingleUpload(jobid int) (*VideoUploadInfo, error) {
	requestBody := jobidProgressRequest{
		BaseModel: BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		JobId:     jobid,
	}
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
func (o *compress) GetJobidProgress(jobid int) (*VideoUploadInfo, error) {
	requestBody := jobidProgressRequest{
		BaseModel: BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		JobId:     jobid,
	}
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
func (o *compress) SetPublishedUpload(requestBody publishedUploadRequest) (*VideoUploadInfo, error) {
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
func (o *compress) UploadS3(destinationFolder string, filename string) error {
	//
	fileDest := destinationFolder + "/" + filename
	o.debugPrint(fileDest)
	bodyPresigned := presignedObject{
		CustomerName: o.customerName,
		FileName:     fileDest,
	}
	resp, err := o.restyPost(PRESIGNED_URL_S3(), bodyPresigned)
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

func (o *compress) Upload(apikey string, file []byte, size int64, customer string, categoryId int, title string, tags string, location string, filename string, targetFolder string) (any, error) {
	bucketFolderDestination := targetFolder + "/" + filename
	responsePresignedUrl, err := o.getMinioURL(bucketFolderDestination, customer)
	if err != nil {
		return nil, err
	}

	response, err := o.restClient.R().
		SetResult(&struct{}{}).
		SetBody(file).
		Put(responsePresignedUrl.Message)

	if err != nil {
		return nil, fmt.Errorf("something went wrong during upload to s3 bucket, err: %s", err.Error())
	}

	if !response.IsSuccess() {
		return nil, fmt.Errorf("upload to s3 bucket failed!, err: %s", err.Error())
	}

	responseCreateUploadAndEncode, err := o.createUpload(apikey, bucketFolderDestination, size, categoryId, title, tags, location, customer)
	if err != nil {
		return nil, err
	}

	if responseCreateUploadAndEncode.Message != "OK" {
		return nil, fmt.Errorf("something went wrong during create upload and encode, err: %s %s", responseCreateUploadAndEncode.Message, responseCreateUploadAndEncode.Response)
	}

	return responseCreateUploadAndEncode, nil
}

func (o *compress) createUpload(apikey string, bucketFolderDestination string, size int64, categoryId int, title string, tags string, location string, customer string) (*responseUpload, error) {
	var ru responseUpload
	_, err := o.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&createUploadByApikeyRequest{
			Filename:      bucketFolderDestination,
			Size:          int(size),
			Category:      categoryId,
			Title:         title,
			Tags:          tags,
			Location:      location,
			ReportedEmail: fmt.Sprintf("%s@tngrm.io", customer),
			Apikey:        apikey,
		}).
		SetResult(&ru).
		Post(CREATE_UPLOAD())
	return &ru, err
}

func (o *compress) getMinioURL(bucketFolderDestination string, customer string) (*responseMinioPresigned, error) {
	response, err := o.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&minioUploadPresignedByApikeyRequest{
			Customer: customer,
			FileName: bucketFolderDestination,
		}).
		SetResult(&responseMinioPresigned{}).
		Post(PRESIGNED_URL_S3())
	if err != nil {
		return nil, err
	}
	responsePresignedUrl := response.Result().(*responseMinioPresigned)

	if responsePresignedUrl.Response != "OK" {
		return nil, fmt.Errorf("something went wrong with getting presigned url minio, err: %s %s", responsePresignedUrl.Response, responsePresignedUrl.Message)
	}

	return responsePresignedUrl, nil
}
