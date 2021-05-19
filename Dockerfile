FROM golang:1.12-alpine

WORKDIR /go/src/tcp_echo_client
COPY . .
RUN apk update \
    && apk --no-cache --update add build-base
RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -o /usr/local/bin/tcp_echo_client

###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/tcp_echo_client /usr/local/bin/tcp_echo_client
RUN apk add --no-cache ca-certificates


