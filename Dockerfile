FROM golang:1.18-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o server ./

FROM debian:stretch-slim
COPY ./wait-for-it.sh /
COPY ./configs /configs

COPY --from=builder /build/configs /configs

COPY --from=builder /build/server /

RUN set -eux; \
	apt-get update; \
	apt-get install -y \
		--no-install-recommends \
		netcat; \
    apt-get install -y  \
        --no-install-recommends  \
        dos2unix; \
    dos2unix wait-for-it.sh; \
        chmod 755 wait-for-it.sh \
