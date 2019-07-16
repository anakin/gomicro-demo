.PHONY: docker build
docker:
	CGO_ENABLED=0 GOOS=linux go build -o demo2
	docker build . -t chope/demo2:latest
	rm api
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./build/demo2
