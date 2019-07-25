FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /

## Launch the wait tool and then your application
COPY ./user-service /
COPY .env.json /
ENV MICRO_REGISTRY=consul \
    MICRO_REGISTRY_ADDRESS=consul:8500 \
    MICRO_BROKER=nats \
    MICRO_BROKER_ADDRESS=nats
ENTRYPOINT /user-service
