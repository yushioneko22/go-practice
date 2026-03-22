# 02 Flag Package

本格的なCLIツールを作る際、`-n name` や `--age 20` といったフラグ（オプション）を解析するために標準ライブラリの `flag` パッケージを使用します。

## 1. 基本解説
### フラグの定義
主に2種類の定義方法があります：
1. `name := flag.String("name", "Guest", "usage description")`: 値への **ポインタ** が返ります。
2. `var age int; flag.IntVar(&age, "age", 0, "usage")`: 既存の変数に値を格納します。

### `flag.Parse()`
フラグの定義が終わったら、必ず `flag.Parse()` を呼び出す必要があります。これを忘れると引数が解析されません。

## 2. 重要なTips
> [!TIP]
> 1. **ポインタに注意**: `flag.String` などはポインタを返すため、値を使うときは `*name` のようにデリファレンスが必要です。
> 2. **ヘルプ画面の自動生成**: プログラムを `-h` または `--help` を付けて実行すると、定義したフラグの一覧と説明が自動で表示されます。
> 3. **非フラグ引数**: フラグ以外の引数（例: `cp source dest` の `source` や `dest`）は、`flag.Args()` で取得できます。

## 演習問題 I/O
### 演習1: 文字列フラグ
- フラグ: `-msg` (初期値: "Hello")
- 入力: `-msg Welcome`
- 出力: `Message: Welcome`

### 演習2: 数値フラグ
- フラグ: `-count` (初期値: 1)
- 入力: `-count 5`
- 出力: 指定された回数分メッセージをループ表示

### 演習3: 真偽値フラグ
- フラグ: `-v` (verboseモード)
- 入力: `-v`
- 出力: `Verbose mode is ON`

### 演習4: 混合利用
- 入力: `-msg Hi -count 2`
- 出力: `Hi` を2回表示

## 実行方法
```bash
cd exercise
go run main.go -msg "Hello" -count 3
```
