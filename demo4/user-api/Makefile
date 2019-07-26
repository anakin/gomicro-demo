.PHONY: docker build
docker:
	CGO_ENABLED=0 GOOS=linux go build -o user-api
	docker build . -t chope/user-api:latest
	rm api
build:
	CGO_ENABLED=0 GOOS=linux go build -o ../build/user-api/user-api