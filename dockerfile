FROM golang:1.19-alpine AS build

LABEL project="42"

ENV PATH="$PATH:$(go env GOPATH)/bin"
ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build
ENV GOOS linux

WORKDIR /go/src

COPY ./ ./

RUN go build -o ./src/bin/kenbunshoku-haki

ENTRYPOINT ["/go/src/bin/kenbunshoku-haki"]