# modular-monolith-example-go

[Go](https://github.com/aplulu/modular-monolith-example-go) | [TypeScript](https://github.com/aplulu/modular-monolith-example-ts)

GoでのModular Monolith及びサービス間通信のサンプル実装です。

サービス間の通信はgRPC経由で行われますが、ServiceServerとServiceClientをServiceAdapterがReflectionを使用して仲介することにより、ネットワーク接続を介さないgRPCを使ったサービス間通信を実現しています。

gRPCのServiceAdapterの実装は `internal/grpc/service_adapter.go` にあります。

また、このサンプルは[Connect](https://connectrpc.com)にも対応しており、外部及び内部の両方での使用が可能です。ただし、Connectを内部のサービス間通信に使用するためには、ServiceごとにAdapterを作成する必要があります。
そのため、Connectはフロントエンドとの通信のみに使用し、サービス間通信にはgRPC (ConnectのgRPC互換ではなく `google.golang.org/grpc` )を使用することを推奨します。

ConnectのServiceAdapterの実装は `internal/grpc/internal_user_adapter.go` にあります。

サービス間通信にConnectを使用するには、 `compose.yml` 内の環境変数 `INTERNAL_PROTOCOL` を `connect` に設定してください。

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

```shell
$ curl -X POST -H "Content-Type: application/json" -d '{}' http://localhost:8080/example.article.v1.ArticleService/ListArticle
```

Connect (GET)

```shell
$ curl 'http://localhost:8080/example.article.v1.ArticleService/ListArticle?encoding=json&message=\{\}'
```

## 未実装や制約項目

* サービス間通信にConnectを使用する場合に、Service Adapterを個別に実装する必要がある。
  * サービス間通信にはgRPCの方を使った方がよいので、Connectの使用は推奨しない。

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
    ├── component // モジュラモノリスの各サービス
    │   ├── article // Articleサービス
    │   │   ├── domain // ドメイン層
    │   │   │   ├── model // モデル
    │   │   │   └── repository // リポジトリ
    │   │   ├── infrastructure // インフラストラクチャ層
    │   │   │   └── inmemory // インメモリリポジトリ
    │   │   ├── interface // インターフェース層
    │   │   │   ├── connect // Connectサーバー
    │   │   │   └── grpc // gRPCサーバー
    │   │   └── usecase // ユースケース層
    │   └── user // ユーザサービス
    ├── config // 設定
    ├── grpc
    │   └── example // 自動生成されたgRPCコード
    └── infrastructure
        └── http // HTTPサーバー
```