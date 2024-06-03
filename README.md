# remote-procedure-call

サーバーとクライアントが Unix domain socket を使用して通信する RPC システムのデモストレーション。サーバーはクライアントのリクエストに基づいて関数を実行し、その結果を返却する。

## Project Structure

```bash
server/
├── cmd/
│   └── server/
│       └── main.go
├── internal/                 # アプリケーション内部で使用されるパッケージを含むディレクトリ
│   ├── handlers/             # リクエスト処理に関するロジックを含むディレクトリ
│   │   └── rpc_handler.go
│   ├── models/               # データモデルを定義するディレクトリ
│   │   ├── request.go
│   │   └── response.go
│   └── services/             # ビジネスロジックを含むサービスを定義するディレクトリ
│       └── methods.go
└── go.mod

client/
├── biome.json
├── bun.lockb
├── package.json
├── src/
│   ├── __tests__/
│   │   ├── client.test.ts
│   │   └── utils.test.ts
│   ├── client.ts             # サーバーとの通信ロジック
│   ├── config.ts             # 設定関連
│   ├── index.ts              # エントリーポイント
│   ├── models.ts             # データモデル
│   ├── requestHandler.ts     # リクエストの処理ロジック
│   └── utils.ts              # ユーティリティ関数
└── tsconfig.json

```

## Communication Format

リクエストとレスポンスメッセージは JSON 形式を用いる。

### Request

```json
{
    "method": "methodName",
    "params": [param1, param2],
    "param_types": ["type1", "type2"],
    "id": 1
}
```

### Response

```json
{
  "result": "result",
  "result_type": "type",
  "id": 1,
  "error": "error message"
}
```

## Supported Methods

| Method                                     | Description                                                                                              | Return Type |
| ------------------------------------------ | -------------------------------------------------------------------------------------------------------- | ----------- |
| **floor(double x)**                        | 10 進数 x を最も近い整数に切り捨て、その結果を整数で返す。                                               | `int`       |
| **nroot(int n, int x)**                    | 方程式 r^n = x における、r の値を計算する。                                                              | `float`     |
| **reverse(string s)**                      | 文字列 s を入力として受け取り、入力文字列の逆である新しい文字列を返す。                                  | `string`    |
| **validAnagram(string str1, string str2)** | 2 つの文字列を入力として受け取り、2 つの入力文字列が互いにアナグラムであるかどうかを示すブール値を返す。 | `bool`      |
| **sort(string[] strArr)**                  | 文字列の配列を入力として受け取り、その配列をソートして、ソート後の文字列の配列を返す。                   | `string[]`  |

## Execution Instructions

サーバーとクライアントの両方を起動する手順は以下の通り。

### Prerequisites

1. Go 言語のインストール

   - 動作確認済みバージョン: 1.21+
   - インストール手順: https://go.dev/doc/install

2. Bun のインストール

   - 動作確認済みバージョン: 1.1.12+
   - インストール手順: https://bun.sh/docs/installation

3. Unix 系 OS

   - 本プロジェクトは Unix domain socket を使用するため、Unix 系 OS（例: Linux, macOS）が必要。

### Starting the Server

```bash
$ cd server
$ go run cmd/server/main.go
```

### Starting the Client

```bash
$ cd client
$ bun install
$ bun run src/index.ts
```

## Usage

1. サーバーの起動: 上記の手順に従ってサーバーを起動する。
1. クライアントの起動: 上記の手順に従ってクライアントを起動する。
1. リクエストの送信:
   - プロンプトが表示されたら、メソッドとパラメータを入力する。

```bash
$ Enter method and parameters: floor 10.7
```

4. レスポンスの受信:
   - サーバーはリクエストを処理し、レスポンスを返す。
   - クライアントはレスポンスを受信し、表示する。

```bash
$ Response: 10
```
