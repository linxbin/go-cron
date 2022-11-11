FROM golang:1.18-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn

WORKDIR /build

COPY . .

RUN go build -o cron-server ./main.go

FROM scratch

COPY --from=builder /build/configs /configs

COPY --from=builder /build/cron-server /

# 需要运行的命令
ENTRYPOINT ["/cron-server"]