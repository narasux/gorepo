.PHONY: tidy build docker

tidy:
	go mod tidy

build: tidy
	go build -o goloadgen main.go

docker:
	docker build -f ./Dockerfile -t goloadgen:1.0.0 --no-cache .
