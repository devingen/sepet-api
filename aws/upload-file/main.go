package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/devingen/sepet-api/aws"
)

// TODO this function is not working properly in Lambda. The uploaded file is malformed.
func main() {
	serviceController, wrap := aws.InitDeps()
	lambda.Start(wrap(serviceController.UploadFile))
}
