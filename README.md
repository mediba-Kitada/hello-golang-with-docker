hello-golang-with-docker
====

- DockerによるGo言語を用いてWEBアプリケーション構築サンプル

## Requirement

- Docker(>=1.12)
- glide(>=0.10.2)

```
$ brew install glide
```

### ローカル(Mac)で開発する場合

- Go(>=1.6.2)
- direnv(>=2.8.1)

## Install

### ライブラリ(サードパーティのパッケージ)をインストールする

```
$ cd /path/to/hello-golang-with-docker/golang/src/persona
$ glide install
```

### コンテナを立ち上げる/停止させる

```
$ cd /path/to/hello-golang-with-docker
# 起動
$ docker-compose up

# 停止(コンテナのプロセスをkill)
$ 

# 停止(DockerのAPIを利用)
$ docker-compose stop
```

## Usage

- ブラウザからGETリクエスト
	- http://localhost/
		- Hello Golang with Docker!!がレスポンスされる
	- http://localhost/persona?id=3
		- 3 : 男性社会人がレスポンスされる
- GET以外をリクエスト
	- ```curl -sS "http://localhost/persona" -X POST --trace-ascii trace-ascii.log -o /dev/null```
	- コンテナのログ(標準出力)を確認

```
go-container     | 2016/07/29 07:23:20 Fobidden Request: POST
````
