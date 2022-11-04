# Sepet API
Rest API for Sepet.

## Building and running the server

Build and run `./cmd/sepet/sepet.go`.

Check the environment variables defined in `config/app.go`.
They must all have the prefix `SEPET_` like `SEPET_S3_ACCESS_KEY_ID`.

## Running the Docker image

```
// pull the latest image
docker pull devingen/sepet-api:VERSION_HERE

// stop and remove any existing container
docker stop sepet-api && docker rm sepet-api

// run the container
docker run \
  --restart always \
  --name sepet-api \
  -e SEPET_API_PORT=8080 \
  -e SEPET_API_LOG_LEVEL=debug \
  -e SEPET_S3_ACCESS_KEY_ID=ACCESSKEYIDFORTHEFILESERVER \
  -e SEPET_S3_SECRET_ACCESS_KEY=ACCESSKEYFORTHEFILESERVER \
  -e SEPET_S3_REGION=region-of-the-cdn \
  -e SEPET_S3_BUCKET=the-root-bucket-name-in-s3 \
  -e SEPET_MONGO_URI=mongodb://complete.mongo.uri \
  -e SEPET_MONGO_DATABASE=mongo-database-name \
  -p 8080:8080 \
  devingen/sepet-api:VERSION_HERE
```

## Development 

### Releasing new Docker image
```
docker build -t devingen/sepet-api:0.0.5 .
docker push devingen/sepet-api:0.0.5
```
