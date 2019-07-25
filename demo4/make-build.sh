#!/bin/sh
docker-compose down
cd gateway
make build
cd ../user-api/
make build
cd ../user-service/
make build
cd ../restaurant-service/
make build
cd ../diner-service/
make build
cd ..
docker-compose build
docker-compose up

