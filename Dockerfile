FROM golang:1.12-alpine

MAINTAINER Tanmay R < Tanmay Renugunta "tanmay17200@gmail.com" >

WORKDIR /go/src/app

ADD . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["app"]
