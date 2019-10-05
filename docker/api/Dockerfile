FROM golang:1.12-alpine

RUN apk --update add git gcc groff curl musl-dev

RUN mkdir /app
WORKDIR /app
ADD go.mod /app/go.mod
ADD go.sum /app/go.sum
RUN go mod download

ADD . /app
