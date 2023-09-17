.PHONY: run

run:
	go run gobfs.go

build:
	GOOS=linux GOARCH=amd64 go build -v -o gobfs gobfs.go