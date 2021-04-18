module github.com/devingen/sepet-api

go 1.12

//replace github.com/devingen/api-core => ../api-core

require (
	github.com/aws/aws-lambda-go v1.16.0
	github.com/aws/aws-sdk-go v1.0.0
	github.com/devingen/api-core v0.0.22
	github.com/go-ini/ini v1.62.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1
	github.com/go-resty/resty/v2 v2.4.0
	github.com/gorilla/mux v1.7.4
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/kisielk/errcheck v1.6.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/stretchr/testify v1.4.0
	go.mongodb.org/mongo-driver v1.3.2
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5 // indirect
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/sys v0.0.0-20210319071255-635bc2c9138d // indirect
	golang.org/x/tools v0.1.0 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
)
