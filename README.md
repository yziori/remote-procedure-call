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
├── Dockerfile
└── go.mod

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
