.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server -tags=prod .

.PHONY: run
run:
	go run -tags=dev .
