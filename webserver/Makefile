.PHONY: tidy build docker

tidy:
	go mod tidy

build: tidy
	go build -o gowebserver main.go

docker:
	docker build -f ./Dockerfile -t gowebserver:1.0.0 --no-cache .
