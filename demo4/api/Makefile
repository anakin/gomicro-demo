.PHONY: docker build
docker:
	CGO_ENABLED=0 GOOS=linux go build -o api main.go
	docker build . -t chope/api:latest
	rm api
build:
	CGO_ENABLED=0 GOOS=linux go build -o api main.go