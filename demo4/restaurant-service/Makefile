.PHONY: docker build
docker:
	CGO_ENABLED=0 GOOS=linux go build -o restaurant-service
	docker build . -t chope/user-service:latest
	rm api
build:
	CGO_ENABLED=0 GOOS=linux go build -o ../build/restaurant-service/restaurant-service