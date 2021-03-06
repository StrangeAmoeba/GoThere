FROM golang:1.12-alpine

MAINTAINER Tanmay R < Tanmay Renugunta "tanmay17200@gmail.com" >

WORKDIR /go/src/concurrency-9

ADD . .

RUN go get -d -v ./...

RUN go install -v ./...

ENV PORT=80

EXPOSE 80

CMD ["concurrency-9"]
