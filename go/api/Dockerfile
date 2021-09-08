FROM golang:1.17-alpine as base

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

### 開発環境
FROM base AS dev

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN go get -u github.com/cosmtrek/air && \
    go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8888
EXPOSE 2345

CMD ["air", "-c", ".air.toml"]

### ビルドステージ
FROM base AS builder
WORKDIR /app

COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o main -a .

### 本番環境
FROM alpine:latest

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    curl \
    tzdata \
    && update-ca-certificates

COPY --from=builder /app/main .
EXPOSE 8080

ENTRYPOINT ["main"]
