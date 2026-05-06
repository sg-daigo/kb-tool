# kb-tool

[Kanboard](https://kanboard.org/) の API を利用して各種情報を取得・表示する CLI ツールです。

## 機能

| サブコマンド | 説明 |
|---|---|
| `projects` | 全プロジェクトの一覧を取得する |
| `tags` | 全タグの一覧を取得する |
| `tasktags` | 指定したタスクに紐づくタグを取得する |
| `users` | 全ユーザーの一覧を取得する |

## 必要要件

- Go 1.26 以上
- Kanboard サーバーへのアクセス権限および API トークン

## インストール

```bash
go install github.com/sg-daigo/kb-tool@latest
```

またはリポジトリをクローンしてビルドする場合:

```bash
git clone https://github.com/<your-name>/kb-tool.git
cd kb-tool
go build -o kb-tool .
```

## 設定

Kanboard の API トークンを環境変数 `KB_TOKEN` に設定してください。

```bash
export KB_TOKEN="your_kanboard_api_token"
```

API トークンは Kanboard の管理画面 **アカウント設定 > API** から確認できます。

## 使い方

### 基本構文

```bash
kb-tools [グローバルオプション] <サブコマンド> [オプション]
```

### グローバルオプション

| オプション      | 短縮形  | デフォルト | 説明                  |
|------------|------| --- |---------------------|
| `--server` | `-s` | `http://localhost` | Kanboard サーバーの URL  |
| `--debug`  | `-d` | `false` | デバッグモードを有効にする       |
| `--json`   | `-j` | `false` | API の戻り値をJSO形式で表示する |

### サブコマンド

#### `projects` - 全プロジェクトの取得

```bash
kb-tools projects
```

#### `tags` - 全タグの取得

```bash
kb-tools tags
```

#### `tasktags` - タスクのタグを取得

| オプション | 短縮形 | 説明 |
| --- | --- | --- |
| `--task` | `-t` | タグを取得するタスクの ID |

```bash
kb-tools tasktags -t 42
```

#### `users` - 全ユーザーの取得

```bash
kb-tools users
```

各コマンドの結果は標準出力に出力されます。

## 使用ライブラリ

| ライブラリ                                                                      | バージョン   | 用途 |
|----------------------------------------------------------------------------|---------| --- |
| [github.com/sg-daigo/kanboard-go](https://github.com/sg-daigo/kanboard-go) | v1.0.0  | Kanboard API クライアント |
| [github.com/spf13/cobra](https://github.com/spf13/cobra)                   | v1.10.2 | CLI フレームワーク |
| [github.com/google/uuid](https://github.com/google/uuid)                   | v1.6.0  | UUID 生成 |

## ライセンス

[MIT License](./LICENSE)
