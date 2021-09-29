FROM node:alpine as admin

RUN apk add --no-cache make git
RUN mkdir -p /app && git clone https://github.com/Ireoo/API.admin.git /app
WORKDIR /app
RUN npm install
RUN npm run build

FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /api-core-src
COPY --from=tonistiigi/xx:golang / /
COPY . /api-core-src
RUN go mod download && \
    make docker && \
    mv ./bin/API-Core-docker /API-Core
RUN /API-Core

FROM alpine:latest
LABEL org.opencontainers.image.source="https://github.com/Ireoo/API-Core"

RUN apk add --no-cache ca-certificates
COPY --from=builder /API-Core /
COPY --from=builder /api-core.conf /
COPY --from=builder /dist /static
ENTRYPOINT ["/API-Core"]
