export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

build-importer:
	- cd cmd/aws_lambda/importer && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j importer_lambda.zip bootstrap

build-generator:
	- cd cmd/aws_lambda/generator && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j generator_lambda.zip bootstrap

build-indexer:
	- cd cmd/aws_lambda/indexer && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j indexer_lambda.zip bootstrap

build-api:
	- cd cmd/aws_lambda/apigateway && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j apigateway_lambda.zip bootstrap

deploy:
	- make build-generator
	- make build-importer
	- make build-indexer
	- cd terraform && terraform apply -var-file='dev.tfvars' -auto-approve

deploy-fast:
	- cd terraform && terraform apply -var-file='dev.tfvars' -auto-approve