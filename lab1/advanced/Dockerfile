FROM golang:1.20-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base

COPY . /go/src/oblaka-advanced-1

WORKDIR /go/src/oblaka-advanced-1

ENV TZ=Europe/Moscow

RUN go mod download

RUN go build -o oblaka-advanced-1

CMD ["./oblaka-advanced-1", "Bruh"]