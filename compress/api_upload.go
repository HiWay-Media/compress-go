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
	requestBody := &jobidProgressRequest{
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
	requestBody := &jobidProgressRequest{
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
func (o *compress) SetPublishedUpload(jobid int, published int) (*VideoUploadInfo, error) {
	if published > 1 && published < 0 {
		return nil, fmt.Errorf("published need to be 0 or 1")
	}
	requestBody := &publishedUploadRequest{
		BaseModel: BaseModel{ClientId: o.GetCliendId(), ApiKey: o.apiKey},
		JobId:     jobid,
		Published: published,
	}
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
* upload file to compress
* gets signed url from minio, uploads and finally creates the upload record.
* @param {string} file
* @param {string} size
* @param {int} categoryId
* @param {string} title
* @param {string} tags
* @param {string} location
* @param {string} filename
* @param {string} targetFolder
 */
func (o *compress) Upload(file []byte, size int64, categoryId int, title string, tags string, location string, filename string) (*ResponseUpload, error) {
	//
	respCustomerS3, err := o.GetCustomerS3Zone()
	if err != nil {
		return nil, err
	}
	if respCustomerS3.Response != "OK" {
		return nil, fmt.Errorf("error %s", respCustomerS3.Message)
	}
	zone := respCustomerS3.Data.Zone
	o.debugPrint("zone ", zone)
	fmt.Println("bucketUpload: ", respCustomerS3.Data.BucketUpload)
	bucketFolderDestination := NormalizeURL(respCustomerS3.Data.BucketUpload + filename)
	//o.debugPrint("bucketFolderDestination " + respCustomerS3.Data.BucketUpload)
	fmt.Println("bucketFolderDestination", bucketFolderDestination)
	//
	responsePresignedUrl, err := o.getMinioURL(bucketFolderDestination, o.customerName)
	if err != nil {
		return nil, err
	}
	//fmt.Println("responsePresignedUrl", responsePresignedUrl)
	response, err := o.restClient.R().
		SetResult(&struct{}{}).
		SetBody(file).
		Put(responsePresignedUrl.Message)

	if err != nil {
		return nil, fmt.Errorf("something went wrong during upload to s3 bucket, err: %s", err.Error())
	}
	if response.IsError() {
		return nil, fmt.Errorf("something went wrong during upload to s3 bucket resp Error, err: %s", response.Error())
	}
	/*if !response.IsSuccess() {
		return nil, fmt.Errorf("upload to s3 bucket failed!, err: %s", err.Error())
	}*/

	responseCreateUploadAndEncode, err := o.createUpload(o.apiKey, bucketFolderDestination, size, categoryId, title, tags, location, o.customerName, zone)
	if err != nil {
		return nil, err
	}

	if responseCreateUploadAndEncode.Response != "OK" {
		return nil, fmt.Errorf("something went wrong during create upload and encode, err: %s %s", responseCreateUploadAndEncode.Message, responseCreateUploadAndEncode.Response)
	}

	return responseCreateUploadAndEncode, nil
}

func (o *compress) createUpload(apikey string, bucketFolderDestination string, size int64, categoryId int, title string, tags string, location string, customer string, zone string) (*ResponseUpload, error) {
	var ru ResponseUpload
	fmt.Println("createUpload ", bucketFolderDestination)
	rCreate, err := o.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&createUploadByApikeyRequest{
			Filename:      bucketFolderDestination,
			Size:          int(size),
			Category:      categoryId,
			Title:         title,
			Tags:          tags,
			Location:      location,
			ReportedEmail: fmt.Sprintf("%s@tngrm.io", customer),
			//ReportedEmail: "",
			Apikey: apikey,
		}).
		SetResult(&ru).
		Post(CREATE_UPLOAD())
	if err != nil {
		return nil, err
	}
	if rCreate.IsError() {
		return nil, fmt.Errorf("something went wrong during create upload and encode, err: %s", rCreate.Error())
	}
	if ru.Response != "OK" {
		return nil, fmt.Errorf("something went wrong during create upload and encode, err: %s %s", ru.Message, ru.Response)
	}
	return &ru, err
}

func (o *compress) getMinioURL(bucketFolderDestination string, customer string) (*responseMinioPresigned, error) {
	response, err := o.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(&minioUploadPresignedByApikeyRequest{
			//Customer: customer,
			ApiKey:   o.apiKey,
			ClientId: o.clientId,
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
	o.debugPrint(responsePresignedUrl)
	return responsePresignedUrl, nil
}
