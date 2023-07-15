# goバージョン
FROM golang:1.19.1-alpine
# アップデートとgitのインストール
RUN apk update && apk add --no-cache git gcc musl-dev

WORKDIR /go/src/github.com/todos-restapi
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/github.com/todos-restapi

# パッケージのインポート
RUN go install golang.org/x/tools/cmd/goimports@latest
ENV PATH $PATH:$(go env GOPATH)/bin
