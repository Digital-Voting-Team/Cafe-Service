FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/cafe-service
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/cafe-service /go/src/cafe-service


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/cafe-service /usr/local/bin/cafe-service
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["cafe-service"]
