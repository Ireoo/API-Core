FROM node:alpine as admin

RUN apk add --no-cache make git
RUN git clone https://github.com/Ireoo/API.admin.git /app
WORKDIR /app
RUN cd /app && npm i && npm run build

FROM golang:alpine as builder

RUN apk add --no-cache make git
WORKDIR /api-core-src
COPY --from=tonistiigi/xx:golang / /
COPY . /api-core-src
RUN go mod download && \
    make docker && \
    mv ./bin/API-Core-docker /API-Core \
    mv ./api-core.conf /api-core.conf

FROM alpine:latest
LABEL org.opencontainers.image.source="https://github.com/Ireoo/API-Core"

RUN apk add --no-cache ca-certificates
COPY --from=builder /API-Core /
COPY --from=builder /api-core.conf /
COPY --from=admin /dist /static
ENTRYPOINT ["/API-Core"]
