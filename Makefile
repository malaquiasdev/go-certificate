export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

build-checkNewCertificates:
	- cd handlers/checkNewCertificates && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../bin/bootstrap *.go
	- chmod +x bin/bootstrap
	- cd bin/ && zip -j check_new_certificates_lambda.zip bootstrap
