package main

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/database"
	"github.com/devingen/sepet-api/config"
	mongo_data_service "github.com/devingen/sepet-api/data-service/mongo-data-service"
	"github.com/devingen/sepet-api/dto"
	"github.com/devingen/sepet-api/model"
	"github.com/devingen/sepet-api/server"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

var restyClient = resty.New()

var appConfig = config.App{
	CDNDomain:   "sepet.devingen.io",
	CDNProtocol: "https",
	Port:        "1005",
	Mongo: config.Mongo{
		URI:      "mongodb://localhost:27017",
		Database: "sepet-api-integration-test",
	},
	S3: config.S3{
		Endpoint:    "http://localhost:9000",
		AccessKeyID: "AKIAIOSFODNN7EXAMPLE",
		AccessKey:   "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		Region:      "ca-central-1",
		Bucket:      "sepet-api-integration-test",
	},
}

func TestEntireFeatureSet(t *testing.T) {

	// prepare the database
	resetDatabase()

	// prepare the file server
	err := resetFileServer()
	assert.Nil(t, err, "resetting file server failed")
	if err != nil {
		return
	}

	// run the sepet api server
	serverEndpoint := "http://localhost:" + appConfig.Port
	sepetServer := runServer()

	// there shouldn't be any buckets
	resp, buckets, err := getBuckets(serverEndpoint)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get buckets status not correct")
	assert.Equal(t, 0, len(buckets), "bucket length is not correct")

	// should create a bucket
	bucketToCreate := model.Bucket{
		Domain:              core.String("dvncdn"),
		Region:              core.String("eu-central-1"),
		IndexPagePath:       core.String("/index.html"),
		ErrorPagePath:       core.String("/index.html"),
		IsCacheEnabled:      core.Bool(false),
		IsVersioningEnabled: core.Bool(false),
	}
	resp, createdBucket, err := createBucket(serverEndpoint, bucketToCreate)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "create bucket status not correct")

	// should get the created bucket
	resp, retrievedBucket, err := getBucket(serverEndpoint, createdBucket.ID.Hex())
	assert.Nil(t, err)
	assert.NotNil(t, retrievedBucket)
	assert.Equal(t, createdBucket.ID, retrievedBucket.ID, "bucket IDs don't match")
	assert.Equal(t, bucketToCreate.Domain, retrievedBucket.Domain, "bucket domains don't match")
	assert.Equal(t, bucketToCreate.IndexPagePath, retrievedBucket.IndexPagePath, "bucket domains don't match")
	assert.Equal(t, bucketToCreate.ErrorPagePath, retrievedBucket.ErrorPagePath, "bucket error page don't match")
	assert.Equal(t, bucketToCreate.IsCacheEnabled, retrievedBucket.IsCacheEnabled, "bucket cache don't match")
	assert.Equal(t, bucketToCreate.IsVersioningEnabled, retrievedBucket.IsVersioningEnabled, "bucket versioning don't match")
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get bucket status not correct")

	// get all buckets should contain one bucket
	resp, buckets, err = getBuckets(serverEndpoint)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get buckets status not correct")
	assert.Equal(t, 1, len(buckets), "bucket length is not correct")

	// should update a bucket
	bucketToUpdate := model.Bucket{
		Domain:              core.String("dvncdnupdated"),
		IndexPagePath:       core.String("/index2.html"),
		ErrorPagePath:       core.String("/index3.html"),
		IsCacheEnabled:      core.Bool(true),
		IsVersioningEnabled: core.Bool(true),
	}
	resp, updatedBucket, err := updateBucket(serverEndpoint, createdBucket.ID.Hex(), bucketToUpdate)
	assert.Nil(t, err)
	assert.NotNil(t, updatedBucket)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "update bucket status not correct")

	// should get the updated bucket
	resp, retrievedBucketAfterUpdate, err := getBucket(serverEndpoint, createdBucket.ID.Hex())
	assert.Nil(t, err)
	assert.NotNil(t, retrievedBucketAfterUpdate)
	assert.Equal(t, createdBucket.ID, retrievedBucketAfterUpdate.ID, "bucket IDs don't match")
	assert.Equal(t, *bucketToUpdate.Domain, *retrievedBucketAfterUpdate.Domain, "bucket domains don't match")
	assert.Equal(t, bucketToUpdate.IndexPagePath, retrievedBucketAfterUpdate.IndexPagePath, "bucket domains don't match")
	assert.Equal(t, bucketToUpdate.ErrorPagePath, retrievedBucketAfterUpdate.ErrorPagePath, "bucket error page don't match")
	assert.Equal(t, bucketToUpdate.IsCacheEnabled, retrievedBucketAfterUpdate.IsCacheEnabled, "bucket cache don't match")
	assert.Equal(t, bucketToUpdate.IsVersioningEnabled, retrievedBucketAfterUpdate.IsVersioningEnabled, "bucket versioning don't match")
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get bucket status not correct")

	// shouldn return empty file list
	resp, fileList, err := getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.Equal(t, &dto.GetFileListResponse{
		Results: []string{},
	}, fileList, "file list is not correct")

	// should fail to create empty folder if the path doesn't end with /
	files := map[string]string{
		"empty-folder": "",
	}
	resp, _, err = uploadFile(serverEndpoint, *retrievedBucketAfterUpdate.Domain, files)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode(), "create folder status not correct")

	// should create empty folder
	files = map[string]string{
		"empty-folder/": "",
	}
	resp, createFolderResponse, err := uploadFile(serverEndpoint, *retrievedBucketAfterUpdate.Domain, files)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "create folder status not correct")
	assert.Equal(t, &dto.UploadFileResponse{
		Locations: []string{"https://" + *retrievedBucketAfterUpdate.Domain + ".sepet.devingen.io/empty-folder/"},
	}, createFolderResponse, "create empty folder file response not correct")

	// should return empty folder in the file list after creating folder
	resp, fileList, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.Equal(t, []string{
		"empty-folder/",
	}, fileList.Results, "file list is not correct")

	// should upload file
	files = map[string]string{
		"images/tr/Marmaris.png":  "files/Marmaris.png",
		"images/tr/Eskisehir.png": "files/Eskisehir.png",
		"images/gb/London.png":    "files/London.png",
	}
	resp, uploadResp, err := uploadFile(serverEndpoint, *retrievedBucketAfterUpdate.Domain, files)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "upload file status not correct")
	assert.ElementsMatch(t, []string{
		"https://" + *retrievedBucketAfterUpdate.Domain + ".sepet.devingen.io/images/tr/Marmaris.png",
		"https://" + *retrievedBucketAfterUpdate.Domain + ".sepet.devingen.io/images/tr/Eskisehir.png",
		"https://" + *retrievedBucketAfterUpdate.Domain + ".sepet.devingen.io/images/gb/London.png",
	}, uploadResp.Locations, "upload file response not correct")

	// should return the folder of the new file in the root file list after uploading file
	resp, fileList, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.ElementsMatch(t, []string{
		"empty-folder/",
		"images/",
	}, fileList.Results, "file list is not correct")

	// should return the file of the sub folder
	resp, fileList, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "images")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.ElementsMatch(t, []string{
		"gb/",
		"tr/",
	}, fileList.Results, "file list is not correct")

	// should return the file of the sub sub folder
	resp, fileList, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "images/tr")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.ElementsMatch(t, []string{
		"Marmaris.png",
		"Eskisehir.png",
	}, fileList.Results, "file list is not correct")

	// should delete file
	resp, err = deleteFile(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "images/tr/Marmaris.png")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode(), "delete file status not correct")

	// should return correct file list after deleting the file
	resp, fileList, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "images/tr")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.ElementsMatch(t, []string{
		"Eskisehir.png",
	}, fileList.Results, "file list is not correct")

	// should delete the entire folder
	resp, err = deleteFile(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "images/")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode(), "delete file status not correct")

	// should return correct file list after deleting the entire folder
	resp, fileList, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode(), "get file list status not correct")
	assert.ElementsMatch(t, []string{
		"empty-folder/",
	}, fileList.Results, "file list is not correct")

	// should delete bucket
	resp, err = deleteBucket(serverEndpoint, retrievedBucketAfterUpdate.ID.Hex())
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, resp.StatusCode(), "delete bucket status not correct")

	// should fail to get file list after deleting bucket
	resp, _, err = getFileList(serverEndpoint, *retrievedBucketAfterUpdate.Domain, "")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode(), "get file list status not correct")

	// TODO should not allow passing bucket version when versioning is not enabled
	// TODO should do all of the things above with the versions

	stopServer(sepetServer)
}

// resetDatabase clears all the data in the buckets collection
func resetDatabase() {
	db, err := database.New(appConfig.Mongo.URI)
	if err != nil {
		log.Fatalf("Database connection failed %s", err.Error())
	}

	collection := db.Client.Database(appConfig.Mongo.Database).Collection(mongo_data_service.CollectionBuckets)
	err = collection.Drop(context.Background())
	if err != nil {
		log.Fatalf("Dropping collection failed %s", err.Error())
	}
}

// resetFileServer clears all the folders in the bucket
func resetFileServer() error {
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(appConfig.S3.AccessKeyID, appConfig.S3.AccessKey, ""),
		Endpoint:         aws.String(appConfig.S3.Endpoint),
		Region:           aws.String(appConfig.S3.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	sess := session.New(s3Config)
	s3Client := s3.New(sess, s3Config)

	files := make([]string, 0)
	nextMarker := ""
	for {
		input := &s3.ListObjectsInput{
			Bucket: aws.String(appConfig.S3.Bucket),
		}

		if nextMarker != "" {
			input.Marker = aws.String(nextMarker)
		}

		result, err := s3Client.ListObjects(input)
		if err != nil {
			return err
		}
		resultLength := len(result.Contents)

		for i := 0; i < resultLength; i++ {
			filePath := *result.Contents[i].Key
			files = append(files, filePath)
		}

		if !*result.IsTruncated {
			// break if there are no more files to fetch
			break
		}

		if result.NextMarker != nil {
			// use the next marker returned in the response to skip already fetched files
			nextMarker = *result.NextMarker
		} else {
			// use the last file path as next marker
			nextMarker = *result.Contents[resultLength-1].Key
		}
	}

	for _, filePath := range files {
		_, dfErr := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: aws.String(appConfig.S3.Bucket),
			Key:    aws.String(filePath),
		})
		if dfErr != nil {
			return dfErr
		}
	}
	return nil
}

func runServer() *http.Server {

	db, err := database.New(appConfig.Mongo.URI)
	if err != nil {
		log.Fatalf("Database connection failed %s", err.Error())
	}

	sepetServer := server.New(appConfig, db)
	go func() {
		if err := sepetServer.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Listen and serve failed %s", err.Error())
		}
	}()
	return sepetServer
}

func stopServer(sepetServer *http.Server) {
	if err := sepetServer.Shutdown(context.Background()); err != nil {
		panic(err)
	}
}

func getBucket(host, id string) (*resty.Response, *model.Bucket, error) {
	var bucket model.Bucket
	resp, err := restyClient.R().
		SetResult(&bucket).
		Get(host + "/buckets/" + id)
	return resp, &bucket, err
}

func getBuckets(host string) (*resty.Response, []*model.Bucket, error) {
	var data dto.GetBucketListResponse
	resp, err := restyClient.R().
		SetResult(&data).
		Get(host + "/buckets")
	return resp, data.Results, err
}

func createBucket(host string, bucket model.Bucket) (*resty.Response, *model.Bucket, error) {
	var responseBucket model.Bucket
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bucket).
		SetResult(&responseBucket).
		Post(host + "/buckets")
	return resp, &responseBucket, err
}

func updateBucket(host, id string, bucket model.Bucket) (*resty.Response, *model.Bucket, error) {
	var responseBucket model.Bucket
	resp, err := restyClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(bucket).
		SetResult(&responseBucket).
		Put(host + "/buckets/" + id)
	return resp, &responseBucket, err
}

func deleteBucket(host, id string) (*resty.Response, error) {
	resp, err := restyClient.R().Delete(host + "/buckets/" + id)
	return resp, err
}

func uploadFile(host, bucketDomain string, files map[string]string) (*resty.Response, *dto.UploadFileResponse, error) {

	request := restyClient.R()
	for fileParamName, filePath := range files {
		if filePath == "" {
			// this is a folder
			// adding a form field with empty value creates a folder
			request.SetMultipartField(fileParamName, "", "", strings.NewReader(""))
		} else {
			// this is a file
			path, err := os.Getwd()
			if err != nil {
				return nil, nil, err
			}
			absoluteFilePath := path + "/" + filePath
			fileBytes, err := ioutil.ReadFile(absoluteFilePath)
			if err != nil {
				return nil, nil, err
			}

			request.SetFileReader(fileParamName, fileParamName, bytes.NewReader(fileBytes))
		}
	}

	var response dto.UploadFileResponse
	resp, err := request.
		SetResult(&response).
		Post(host + "/" + bucketDomain)
	return resp, &response, err
}

func getFileList(host, bucketDomain, subfolder string) (*resty.Response, *dto.GetFileListResponse, error) {
	var response dto.GetFileListResponse
	resp, err := restyClient.R().
		SetResult(&response).
		Get(host + "/" + bucketDomain + "/" + subfolder)
	return resp, &response, err
}

func deleteFile(host, bucketDomain, filePath string) (*resty.Response, error) {
	resp, err := restyClient.R().
		Delete(host + "/" + bucketDomain + "/" + filePath)
	return resp, err
}
