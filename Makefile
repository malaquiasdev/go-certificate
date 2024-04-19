export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

build-importer:
	- cd handlers/importer && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j importer_lambda.zip bootstrap

build-generator:
	- cd handlers/generator && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j generator_lambda.zip bootstrap