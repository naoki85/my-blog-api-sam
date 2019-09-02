.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./handler/handler
	
build:
	GOOS=linux GOARCH=amd64 go build -o handler/handler ./handler