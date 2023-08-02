.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-buckets aws/get-buckets/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/create-bucket aws/create-bucket/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-bucket aws/get-bucket/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-bucket-version-list aws/get-bucket-version-list/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/get-file-list-or-delete-file aws/get-file-list-or-delete-file/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/update-bucket aws/update-bucket/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/delete-bucket aws/delete-bucket/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/upload-file aws/upload-file/main.go

lint:
	golint ./...

checkerr:
	errcheck ./...

scan: lint checkerr

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	serverless deploy --stage prod --region eu-central-1 --verbose

teardown: clean
	serverless remove --stage prod --region eu-central-1 --verbose

deploy-dev: clean build
	serverless deploy --stage dev --region ca-central-1 --verbose

teardown-dev: clean
	serverless remove --stage dev --region ca-central-1 --verbose

tools:
	go get github.com/kisielk/errcheck
	go get github.com/golang/lint/golint