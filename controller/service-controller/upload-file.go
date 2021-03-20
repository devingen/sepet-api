package srvcont

import (
	"context"
	core "github.com/devingen/api-core"
	"github.com/devingen/api-core/log"
	"github.com/devingen/sepet-api/dto"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

// UploadFile implements IServiceController interface
func (controller ServiceController) UploadFile(ctx context.Context, req core.Request) (interface{}, int, error) {

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	domain := req.PathParameters["domain"]
	if domain == "" {
		return "", http.StatusBadRequest, core.NewError(http.StatusBadRequest, "domain-missing")
	}

	bucket, err := controller.DataService.FindBucketWithDomain(ctx, domain)
	if err != nil {
		return nil, 0, err
	}
	if bucket == nil {
		return nil, 0, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	bucketVersion := bucket.Version
	versionFromHeader, headerHasVersion := req.GetHeader("Bucket-Version")
	if headerHasVersion {
		if !bucket.IsVersioningEnabled {
			return nil, 0, core.NewError(http.StatusBadRequest, "versioning-not-enabled")
		}
		bucketVersion = versionFromHeader
	}

	files, err := extractFiles(req)
	if err != nil {
		return nil, 0, err
	}

	locations, err := controller.FileService.UploadFile(ctx, bucket, bucketVersion, files)
	if err != nil {
		return nil, 0, err
	}

	logger.WithFields(logrus.Fields{
		"sepet-domain": bucket.Domain,
	}).Info("finished-uploading-file")

	return &dto.UploadFileResponse{
		Locations: locations,
	}, http.StatusOK, nil
}

func extractFiles(req core.Request) (map[string]multipart.File, error) {
	// create http.Request with the request data to parse the multipart form data
	headers := map[string][]string{}
	for name, value := range req.Headers {
		headers[name] = []string{value}
	}

	httpReq := http.Request{
		Body:   ioutil.NopCloser(strings.NewReader(req.Body)),
		Header: headers,
	}

	err := httpReq.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, err
	}

	files := map[string]multipart.File{}
	for key := range httpReq.MultipartForm.Value {
		files[key] = nil
	}
	for key := range httpReq.MultipartForm.File {
		file, _, rfErr := httpReq.FormFile(key)
		if rfErr != nil {
			return nil, rfErr
		}
		files[key] = file
	}
	return files, nil
}
