.PHONY: run

run:
	go run puf.go

build:
	go build -v -o puf puf.go