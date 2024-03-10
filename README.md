# modular-monolith-example-go

GoでのModular Monolith及びモジュール間通信のサンプル実装です。

モジュール間の通信はgRPC経由で行われますが、ServiceServerとServiceClientをServiceAdapterがReflectionを使用して仲介することにより、ネットワーク接続を介さないgRPCを使ったモジュール間通信を実現しています。
gRPCのServiceAdapterの実装は `internal/grpc/service_adapter.go` にあります。

また、このサンプルは[Connect](https://connectrpc.com)にも対応しており、外部及び内部の両方での使用が可能です。ただし、Connectを内部のモジュール間通信に使用するためには、Reflectionが使用できないため、ServiceごとにAdapterを作成する必要があります。
そのため、Connectはフロントエンドとの通信のみに使用し、モジュール間通信にはgRPC (ConnectのgRPC互換ではなく `google.golang.org/grpc` )を使用することを推奨します。
ConnectのServiceAdapterの実装は `internal/grpc/internal_user_adapter.go` にあります。

モジュール間通信にConnectを使用するには、 `compose.yml` 内の環境変数 `INTERNAL_PROTOCOL` を `connect` に設定してください。

## 前提環境

* Docker
* Docker Compose
* GNU Make

## ローカル環境の起動

```shell
$ make up
```

## ローカル環境の停止

```shell
$ make down
```

## gRPCのコード生成

```shell
$ make buf-generate
```

## 確認用のgRPCリクエスト

gRPC
```shell
$ grpcurl --plaintext localhost:8080 example.article.v1.ArticleService.ListArticle 
```

Connect

http://localhost:8080/example.article.v1.ArticleService/ListArticle?encoding=json&message={}
```shell
$ curl -X POST -H "Content-Type: application/json" -d '{}' http://localhost:8080/example.article.v1.ArticleService/ListArticle
```

Connect (GET)

```shell
$ curl 'http://localhost:8080/example.article.v1.ArticleService/ListArticle?encoding=json&message=\{\}'
```


## ディレクトリ構成

```
.
├── api
│   └── proto // gRPCのプロトコル定義
├── cmd
│   └── serve
│       └── main.go // サーバーのエントリーポイント
├── docker // Docker関連のファイル
└── internal
    ├── components // モジュラモノリスの各モジュール
    │   ├── article // Articleモジュール
    │   │   ├── domain // ドメイン層
    │   │   │   ├── model // モデル
    │   │   │   └── repository // リポジトリ
    │   │   ├── infrastructure // インフラストラクチャ層
    │   │   │   └── inmemory // インメモリリポジトリ
    │   │   ├── interface // インターフェース層
    │   │   │   ├── connect // Connectサーバー
    │   │   │   └── grpc // gRPCサーバー
    │   │   └── usecase // ユースケース層
    │   └── user // ユーザモジュール
    ├── config // 設定
    ├── grpc
    │   └── example // 自動生成されたgRPCコード
    └── infrastructure
        └── http // HTTPサーバー

```