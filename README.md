# Go Master Roadmap: 基礎から実践（Web/WS）まで

このプロジェクトは、Go言語の基礎から始まり、CLIツールの開発、モダンなWebフレームワークの比較、そしてリアルタイムなWebSocket通信までを体系的に学ぶための技術教材です。

---

## 1. プロジェクト構成

```text
.
├── phase1_basics/        # 基礎編: Go言語のコア（変数、ポインタ、IF等）
├── phase2_cli/           # 基礎編: 標準ライブラリとCUI（File IO, JSON, 並行処理）
├── phase3_web/           # 応用編: Web開発（net/http, Gin, Echo, カレンダー比較）
└── phase4_websocket/     # 応用編: リアルタイム通信（接続, Hub, ブロードキャスト）
└── template.md           # 教材作成の「秘伝のレシピ」
```

---

## 2. 学習の流れ

### ステップ1：基礎固め (Phase 1 & 2)
まずはGoの文法と、標準ライブラリを使ったCLIツール開発を学びます。
- `01_hello_world` 〜 `05_json_processing` で基本を習得
- `06_todo_project` 等の実践課題で知識を統合

### ステップ2：Web開発への招待 (Phase 3)
標準ライブラリ（フルスクラッチ）と、人気フレームワーク（Gin, Echo）を比較しながら学びます。
- 全く同じ仕様のカレンダーAPIを3パターンで実装し、トレードオフを理解する

### ステップ3：リアルタイム通信の極意 (Phase 4)
双方向通信の仕組みを学び、チャットアプリを構築します。
- HTTPの限界を超えた、モダンなWebシステムの裏側をマスター

---

## 3. セットアップ手順

1. **依存ライブラリのインストール**
   ```bash
   go mod tidy
   ```

2. **各トピックの実行**
   各フォルダ内の `example/main.go` を実行して動作を確認します。
   ```bash
   # 例: Gin版カレンダーを起動
   cd phase3_web/04_calendar_project/gin
   go run main.go
   ```

---

## 4. 必要な環境
- **Go**: 1.25.1 以上推奨
- **APIキー**: 今回の教材では外部サービスのAPIキーは不要です。

---

## 5. 利用ライブラリの役割一覧

| ライブラリ | 役割 | なぜ必要か |
| :--- | :--- | :--- |
| `net/http` | 標準Webサーバー | GoのWeb開発のすべての基礎であり、仕組みを理解するため |
| `github.com/gin-gonic/gin` | Webフレームワーク | 高速なルーティングと生産性の高いAPI開発のため |
| `github.com/labstack/echo/v4` | Webフレームワーク | シンプルな設計とミドルウェアの拡張性を学ぶため |
| `github.com/gorilla/websocket` | WebSocket | Goでデファクトスタンダードの双方向通信を実現するため |

---

## 6. 学習を成功させる5つのコツ

1. **「なぜ」を意識する**: READMEの「なぜ必要か」セクションを読み込み、背景を理解する。
2. **写経ではなく改造**: サンプルコードを動かしたら、少しだけ数字や文字を変えて挙動を見る。
3. **比喩で覚える**: 技術概念を日常の出来事に例えた説明を大切にする。
4. **ターミナルを使い倒す**: `curl` や `ls` など、コマンドを使ってサーバーと対話する。
5. **焦らず3層構造で**: 基礎編(example)で納得してから応用・実践へ進む。

---

## 7. 参考リンク
- [Go Documentation (Official)](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Gin Web Framework Documentation](https://gin-gonic.com/docs/)
- [Echo Web Framework Documentation](https://echo.labstack.com/docs)
- [Gorilla WebSocket Documentation](https://pkg.go.dev/github.com/gorilla/websocket)
