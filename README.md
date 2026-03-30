# Go Master Roadmap: 基礎から面接対策まで

このプロジェクトは、Go言語の基礎から始まり、CLIツール開発、Web開発、リアルタイム通信、そしてGoogle等の技術面接対策までを体系的に学ぶための技術教材です。

---

## 1. プロジェクト構成

```text
.
├── phase1_basics/          # 基礎編: Go言語のコア（変数、ポインタ、IF等）
├── phase2_cli/             # 基礎編: 標準ライブラリとCUI（File IO, JSON, 並行処理）
├── phase3_web/             # 応用編: Web開発（net/http, Gin, Echo, カレンダー比較）
├── phase4_websocket/       # 応用編: リアルタイム通信（接続, Hub, ブロードキャスト）
├── phase5_data_structures/ # 面接対策: データ構造（配列, 連結リスト, 木, グラフ等）
├── phase6_algorithms/      # 面接対策: アルゴリズム（ソート, DP, 貪欲法, Two Pointers等）
├── phase7_practice/        # 面接対策: LeetCode頻出問題（Easy〜Hard）
├── phase8_system_design/   # 面接対策: システムデザイン（URL短縮, Cache, LB等）
└── template.md             # 教材作成の「秘伝のレシピ」
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

### ステップ4：面接対策 - データ構造 & アルゴリズム (Phase 5 & 6)
Google等の技術面接で問われるデータ構造とアルゴリズムをGoで実装します。
- 配列、連結リスト、スタック、ハッシュマップ、二分木、ヒープ、グラフ
- ソート、二分探索、再帰、DFS/BFS、動的計画法、貪欲法、スライディングウィンドウ、Two Pointers

### ステップ5：面接対策 - 実践問題 (Phase 7)
LeetCode頻出問題をEasy→Medium→Hardの段階で解きます。
- 配列/文字列、木/グラフ、DP の各カテゴリ別に整理

### ステップ6：面接対策 - システムデザイン (Phase 8)
大規模システムの設計面接に備え、代表的なシステムをGoでプロトタイプ実装します。
- URL短縮、Rate Limiter、LRU Cache、メッセージキュー、ロードバランサー

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

   # 例: Two Sum問題のサンプル
   cd phase5_data_structures/01_array_slice/example
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
| `container/heap` | 優先度キュー | ヒープ/Top K問題の実装に使用（標準ライブラリ） |

---

## 6. 面接対策トピック一覧

### Phase 5: データ構造
| # | トピック | 主な問題 |
|:--|:--|:--|
| 01 | 配列・スライス | Two Sum, 回転, 重複除去 |
| 02 | 連結リスト | 反転, サイクル検出, マージ |
| 03 | スタック・キュー | 有効な括弧, MinStack |
| 04 | ハッシュマップ | アナグラム, 頻度カウント |
| 05 | 二分木 | 走査, 深さ, 反転, レベル順 |
| 06 | ヒープ | Top K, 中央値 |
| 07 | グラフ | 島の数, トポロジカルソート |

### Phase 6: アルゴリズム
| # | トピック | 主な問題 |
|:--|:--|:--|
| 01 | ソート | マージソート, クイックソート |
| 02 | 二分探索 | 回転配列, 平方根 |
| 03 | 再帰 | べき乗, 全順列 |
| 04 | DFS/BFS | 迷路の最短経路, 全経路列挙 |
| 05 | 動的計画法 | ナップサック, LCS |
| 06 | 貪欲法 | 区間スケジューリング, ジャンプゲーム |
| 07 | スライディングウィンドウ | 最長部分文字列, 最小被覆 |
| 08 | Two Pointers | 3Sum, Container With Most Water |

### Phase 7: 実践問題
| # | 難易度 | 主な問題 |
|:--|:--|:--|
| 01 | Easy | Merge Lists, Stock, Anagram |
| 02 | Medium (配列) | Product Except Self, 3Sum, Group Anagrams |
| 03 | Medium (木/グラフ) | Valid BST, Islands, Course Schedule |
| 04 | Medium (DP) | LIS, Word Break, Unique Paths |
| 05 | Hard | Trapping Rain Water, Min Window, Serialize Tree |

### Phase 8: システムデザイン
| # | トピック | 学べること |
|:--|:--|:--|
| 01 | URL短縮 | Base62, キャッシュ, スケーリング |
| 02 | Rate Limiter | Token Bucket, Sliding Window |
| 03 | キャッシュ(LRU) | 双方向連結リスト + ハッシュマップ |
| 04 | メッセージキュー | Pub/Sub, 永続化 |
| 05 | ロードバランサー | Round Robin, ヘルスチェック |

---

## 7. 学習を成功させる5つのコツ

1. **「なぜ」を意識する**: READMEの「なぜ必要か」セクションを読み込み、背景を理解する。
2. **写経ではなく改造**: サンプルコードを動かしたら、少しだけ数字や文字を変えて挙動を見る。
3. **比喩で覚える**: 技術概念を日常の出来事に例えた説明を大切にする。
4. **ターミナルを使い倒す**: `curl` や `ls` など、コマンドを使ってサーバーと対話する。
5. **焦らず3層構造で**: 基礎編(example)で納得してから応用・実践へ進む。

---

## 8. 参考リンク
- [Go Documentation (Official)](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Gin Web Framework Documentation](https://gin-gonic.com/docs/)
- [Echo Web Framework Documentation](https://echo.labstack.com/docs)
- [Gorilla WebSocket Documentation](https://pkg.go.dev/github.com/gorilla/websocket)
- [LeetCode](https://leetcode.com/) - 面接問題の練習
- [NeetCode](https://neetcode.io/) - 面接問題のパターン別整理
