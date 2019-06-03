FROM golang:1.12.5-alpine3.9 AS build

ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN apk --update add git

RUN mkdir -p /adam/src && mkdir -p /adam/bin
WORKDIR /adam/src
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . /adam/src

ARG GOOS=linux
ARG GOARCH=amd64

RUN go build -o /adam/bin/adam main.go

FROM scratch

COPY --from=build /adam/bin/adam /adam/bin/adam
ENTRYPOINT ["/adam/bin/adam"]

