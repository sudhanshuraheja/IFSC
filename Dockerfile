FROM golang:1.9.2-alpine3.7 AS build
LABEL maintainer="sudhanshu@go-jek.com"

# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN apk add --no-cache bash
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers will only be re-built when Gopkg files are updated
RUN mkdir -p /go/src/github.com/sudhanshuraheja/ifsc
ADD . /go/src/github.com/sudhanshuraheja/ifsc/
WORKDIR /go/src/github.com/sudhanshuraheja/ifsc

RUN dep ensure -vendor-only
RUN go build -o /go/bin/ifsc

CMD [ "ifsc", "start" ]