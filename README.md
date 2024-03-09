# modular-monolith-example-go

GoでのModular Monolith及びモジュール間通信のサンプル実装です。

モジュール間の通信はgRPC経由としていますが、ServiceServerとServiceClientをServiceAdapterがreflectionで仲介することにより、ネットワーク接続がない状態でgRPCを使ったモジュール間通信を実現しています。

ServiceAdapterの実装: `internal/grpc/service_adapter.go`

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

```shell
$ grpcurl --plaintext localhost:8080 example.article.v1.ArticleService.ListArticle 
```

## ディレクトリ構成

```
.
├── api
│   └── proto // gRPCのプロトコル定義
├── cmd
│   └── serve // サーバーのエントリーポイント
├── docker // Dockerfile
└── internal
    ├── components // モジュラモノリスの各モジュール
    │   ├── article // Articleモジュール
    │   │   ├── domain // ドメイン層
    │   │   │   ├── model // モデル
    │   │   │   └── repository // リポジトリ
    │   │   ├── infrastructure // インフラストラクチャ層
    │   │   │   └── inmemory // インメモリリポジトリ
    │   │   ├── interface // インターフェース層
    │   │   │   └── grpc // gRPCサーバー
    │   │   └── usecase // ユースケース層
    │   └── user // ユーザモジュール
    ├── config // 設定
    ├── grpc
    │   └── example // 自動生成されたgRPCコード
    └── infrastructure
        └── http // HTTPサーバー

```