FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /api-core-src
COPY --from=tonistiigi/xx:golang / /
COPY . /api-core-src
RUN go mod download && \
    make docker && \
    mv ./bin/API-Core-docker /API-Core

FROM alpine:latest
LABEL org.opencontainers.image.source="https://github.com/Ireoo/API-Core"

RUN apk add --no-cache ca-certificates
COPY --from=builder /API-Core /
ENTRYPOINT ["/API-Core"]
