.PHONY: docker build
docker:
	CGO_ENABLED=0 GOOS=linux go build -o user-service
	docker build . -t chope/micro:latest
	rm api
build:
	CGO_ENABLED=0 GOOS=linux go build -o ../build/gateway/micro