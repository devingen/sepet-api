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
func (controller ServiceController) UploadFile(ctx context.Context, req core.Request) (*core.Response, error) {

	_, interceptorStatusCode, interceptorError := controller.InterceptorService.Pre(ctx, req)
	if interceptorError != nil {
		return &core.Response{
			StatusCode: interceptorStatusCode,
			Body:       interceptorError,
		}, nil
	}

	logger, err := log.Of(ctx)
	if err != nil {
		return nil, core.NewStatusError(http.StatusInternalServerError)
	}

	domain := req.PathParameters["domain"]
	if domain == "" {
		return nil, core.NewError(http.StatusBadRequest, "domain-missing")
	}

	bucket, err := controller.DataService.FindBucketWithDomain(ctx, domain)
	if err != nil {
		return nil, err
	}
	if bucket == nil {
		return nil, core.NewError(http.StatusNotFound, "bucket-not-found")
	}

	bucketVersion := bucket.Version
	versionFromHeader, headerHasVersion := req.GetHeader("bucket-version")
	if headerHasVersion {
		if !core.BoolValue(bucket.IsVersioningEnabled()) {
			return nil, core.NewError(http.StatusBadRequest, "versioning-not-enabled")
		}
		bucketVersion = core.String(versionFromHeader)
	}

	files, fileHeaders, err := extractFiles(req)
	if err != nil {
		return nil, err
	}

	locations, err := controller.FileService.UploadFile(ctx, bucket, *bucketVersion, files, fileHeaders)
	if err != nil {
		return nil, err
	}

	logger.WithFields(logrus.Fields{
		"sepet-domain": bucket.Domain,
	}).Info("finished-uploading-file")

	return &core.Response{
		StatusCode: http.StatusOK,
		Body: dto.UploadFileResponse{
			Locations: locations,
		},
	}, nil
}

func extractFiles(req core.Request) (map[string]multipart.File, map[string]*multipart.FileHeader, error) {
	// create http.Request with the request data to parse the multipart form data
	headers := map[string][]string{}
	for name, value := range req.Headers {
		if name == "content-type" {
			// ParseMultipartForm is very case sensitive for Content-Type header
			headers["Content-Type"] = []string{value}
		}
		headers[name] = []string{value}
	}

	httpReq := http.Request{
		Body:   ioutil.NopCloser(strings.NewReader(req.Body)),
		Header: headers,
	}

	err := httpReq.ParseMultipartForm(32 << 20)
	if err != nil {
		return nil, nil, err
	}

	files := map[string]multipart.File{}
	fileHeaders := map[string]*multipart.FileHeader{}
	for key := range httpReq.MultipartForm.Value {
		files[key] = nil
	}
	for key := range httpReq.MultipartForm.File {
		file, header, rfErr := httpReq.FormFile(key)
		if rfErr != nil {
			return nil, nil, rfErr
		}
		files[key] = file
		fileHeaders[key] = header
	}
	return files, fileHeaders, nil
}
