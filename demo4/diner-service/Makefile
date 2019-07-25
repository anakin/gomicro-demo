.PHONY: docker build
docker:
	CGO_ENABLED=0 GOOS=linux go build -o diner-service
	docker build . -t chope/diner-service:latest
	rm api
build:
	CGO_ENABLED=0 GOOS=linux go build -o ../build/diner-service/diner-service