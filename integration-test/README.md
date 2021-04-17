# Sepet API Integration Tests

This folder contains the test files and assets used for the integration tests.

`main_test.go` is the main test file that contains all the test scenarios. It includes CRUD operations for
bucket management as well as the file operations. It covers most of the scenarios, but it's not 100%.

Sadly, the test cases are not atomic. Meaning that they'll all start to fail if a scenario fails.

## Prerequisite

### Mongo DB
You must have a MongoDB instance running on the local machine 
that is accessible with the URI `mongodb://localhost:27017`.
The test process will create a new database named `sepet-api-integration-test` in the MongoDB instance.

### MinIO
You must have a MinIO server running locally. File server credentials
are set on top of the `main_test.go`, so you need to replace those values 
after you run the MinIO server locally.

You must also have a bucket in MinIO named `sepet-api-integration-test` that'll 
be used by the integration tests to do file operations.

## Running the tests

Run this command to execute the integration tests.

```
go test ./integration-test/...
```

## TODO

* Bucket versioning must also be covered by the tests.