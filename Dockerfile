# 公式イメージを利用
## Golang 1.6.3
## OS Alpine Linux
FROM golang:1.6.3-alpine

# 連絡先
MAINTAINER mediba-kitada <kitada@mediba.jp>

# assetsファイルをコピー
COPY ./assets $GOPATH/assets 

# アプリケーションをコピー
COPY ./src $GOPATH/src 

# アプリケーションのビルド及びインストール
RUN go-wrapper install main

# アプリケーションの起動
ENTRYPOINT exec $GOPATH/bin/main

# 8080番ポートを待受
EXPOSE 8080
