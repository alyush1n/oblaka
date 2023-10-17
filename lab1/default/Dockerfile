# bad practice 1 - latest
FROM golang:latest as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base

COPY . /go/src/sport-admin

WORKDIR /go/src/sport-admin

ENV TZ=Europe/Moscow

RUN go build cmd/main.go

FROM alpine:3.14

RUN apk add --no-cache curl tzdata

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/src/sport-admin/main /app/sport-admin

RUN wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" -O /YandexCA.crt

WORKDIR /app

ENTRYPOINT ["./sport-admin"]

# bad practice 2 - multiple services inside 1 container
FROM golang:1.20-alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh build-base

# bad practice 3 - environment variables in Dockerfile
ENV USER=admin
ENV PASSWORD=password

COPY . /go/src/sport-admin

WORKDIR /go/src/sport-admin

ENV TZ=Europe/Moscow

RUN go build cmd/main.go

FROM alpine:3.14

RUN apk add --no-cache curl tzdata

ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY --from=builder /go/src/sport-admin/main /app/sport-admin

RUN wget "https://storage.yandexcloud.net/cloud-certs/CA.pem" -O /YandexCA.crt

WORKDIR /app

ENTRYPOINT ["./sport-admin"]