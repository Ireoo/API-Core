FROM golang:alpine as builder

WORKDIR /api-core-src
COPY --from=tonistiigi/xx:golang / /
COPY . /api-core-src
RUN go mod download && \
    make docker && \
    mv ./bin/api-core-docker /api-core

FROM alpine:latest
LABEL org.opencontainers.image.source="https://github.com/Ireoo/API-Core"

RUN apk add --no-cache ca-certificates
COPY --from=builder /api-core /
ENTRYPOINT ["/api-core"]
