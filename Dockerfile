FROM golang:1.12-alpine

MAINTAINER Tanmay R < Tanmay Renugunta "tanmay17200@gmail.com" >

WORKDIR /go/src/concurrency-9

ADD . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 9000

CMD ["concurrency-9"]
