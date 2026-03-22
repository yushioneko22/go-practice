# 01 Standard Basics (標準ライブラリ `net/http` の基礎)

このトピックでは、**「Goの標準ライブラリのみでWebサーバーを構築し、リクエストの解析とレスポンスの返却」**を行う方法を学びます。

---

## 1. なぜ必要か
フレームワーク（GinやEcho）は非常に便利ですが、その中身はすべてこの `net/http` の上に成り立っています。
標準ライブラリの書き方を知っておくことで、「フレームワークが裏側で何をしているのか」を理解でき、複雑なカスタマイズやトラブルシューティングが可能になります。

## 2. 比喩：`net/http` は「手書きの注文票」
- **Framework**: タブレット注文システム。ボタン一つで色々やってくれる。
- **Standard Lib**: 紙とペン。項目（パス、メソッド、ヘッダー）を一つずつ自分で書き込み、確認する必要がある。手間はかかるが、仕組みが最もよく見える。

## 3. コード解説

### 01. ルーティングの設定 (15-20行目) ★ここが一番重要
```go
http.HandleFunc("/hello", helloHandler)
```
特定のURLパス（`/hello`）と、それに対応する処理（`helloHandler`）を紐付けます。

### 02. メソッドの判定とJSON返却 (25-35行目)
```go
if r.Method != http.MethodGet {
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    return
}
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(response)
```
標準ライブラリでは、HTTPメソッドの判定や、レスポンスヘッダーの設定をすべて **明示的** に行う必要があります。

---

## 4. 重要なTips
> [!TIP]
> 1. **`r.URL.Path` の罠**: 標準のルーティングでは、詳細なパスパラメータ（`/user/:id` など）を自動で抽出してくれません。自分で文字列操作をして取り出す必要があります。
> 2. **`w.Header().Set` のタイミング**: `w.Write()` や `json.Encode()` を呼んだ後にヘッダーをセットしても無視されます。必ず「書く前」にセットしましょう。
> 3. **`http.DefaultServeMux`**: `http.HandleFunc` を使うと、グローバルなルーティングマネージャーに登録されます。小規模なら便利ですが、大規模開発では個別に `ServeMux` を作るのが一般的です。

## 5. まとめ
`net/http` は「不便」に感じるかもしれませんが、GoのWebエンジニアとしての「基礎体力」を作るために最も重要なパーツです。

---

## 演習問題 I/O
### 仕様
- `/greet` にGETリクエストを送ると、JSONで `{ "message": "Hello, Gopher!" }` と返してください。
- POSTなどGET以外のメソッドの場合は `405 Method Not Allowed` を返してください。

## 実行方法
```bash
cd exercise
go run main.go
# 別のターミナルで
curl -X GET http://localhost:8080/greet
```
