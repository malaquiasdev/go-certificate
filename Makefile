export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

build:
	- cd handlers/ && go build -a -installsuffix cgo -ldflags '-s -w -extldflags "-static"' -o ../../bin/main *.go
	- chmod +x bin/main
	- cd bin/ && zip -j main.zip main
