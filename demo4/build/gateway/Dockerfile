FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY micro /micro
COPY .env.json /
ENV MICRO_REGISTRY=consul \
    MICRO_REGISTRY_ADDRESS=consul:8500 \
    MICRO_API_HANDLER=http \
    MICRO_API_NAMESPACE=chope.co.api
ENTRYPOINT /micro api
LABEL Name=micro Version=0.0.1
EXPOSE 8080 81
