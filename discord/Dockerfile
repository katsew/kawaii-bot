FROM golang:alpine AS build-machine

RUN apk --update add git openssh ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    rm /var/cache/apk/*

COPY . /go/src/github.com/katsew/kawaii-bot

WORKDIR /go/src/github.com/katsew/kawaii-bot/discord

RUN go-wrapper download
RUN go-wrapper install

FROM alpine:latest

COPY --from=build-machine /go/bin/discord ./kawaii-bot-discord

CMD ["./kawaii-bot-discord"]