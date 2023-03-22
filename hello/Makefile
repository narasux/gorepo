.PHONY: tidy build docker

tidy:
	go mod tidy

build: tidy
	go build -o gohello main.go

docker:
	docker build -f ./Dockerfile -t gohello:1.0.0 --no-cache .
