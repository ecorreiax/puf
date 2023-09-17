.PHONY: run

run:
	go run gobfs.go

build:
	go build -v -o gobfs gobfs.go