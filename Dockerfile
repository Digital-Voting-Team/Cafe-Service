FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/Cafe-Service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/Cafe-Service /go/src/Cafe-Service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/Cafe-Service /usr/local/bin/Cafe-Service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["Cafe-Service"]
