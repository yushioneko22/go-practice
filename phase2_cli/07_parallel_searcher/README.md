# 07 Parallel Searcher (並行検索ツール)

Goの最大の武器である「並行処理（Concurrency）」を実戦投入します。TerraformやDockerが高速なのも、この仕組みをフル活用しているからです。

## 1. 基本解説
### Goroutine（ゴルーチン）
`go` キーワードを関数の前に付けるだけで、その関数を「裏側（バックグラウンド）」で実行できます。非常に軽量で、数万個同時に動かすことも可能です。

### Channel（チャネル）
並行して動いているゴルーチン同士が「データを安全にやり取りする」ためのパイプです。
- `ch := make(chan string)`: チャネル作成
- `ch <- "data"`: データを送る
- `data := <-ch`: データを受け取る

### WaitGroup
すべてのゴルーチンが終了するのを待つための仕組みです（`sync` パッケージ）。これを忘れると、メインプログラムが先に終わってしまい、裏側の処理が途中で強制終了されます。

## 2. 重要なTips
> [!TIP]
> 1. **並行 vs 並列**: Goは「並行（Concurrency）」、つまり複数の仕事をうまく交通整理して進めるのが得意です（実際のCPUコア数に応じて「並列」にもなります）。
> 2. **チャネルのクローズ**: データの送信が終わったら `close(ch)` することで、受け取り側に「もう終わりだよ」と伝えることができます。
> 3. **Race Conditionに注意**: 複数のゴルーチンから同時に一つの変数に書き込むと、計算がおかしくなります（競合状態）。チャネルを使ってデータの受け渡しをするのが安全なやり方（Do not communicate by sharing memory; instead, share memory by communicating）です。

## 演習問題 I/O
### 動作イメージ
```bash
# 特定の単語をカレントディレクトリのファイルから並行検索
go run main.go "keyword"
# -> 出力例: 
# [Found] file1.txt: line 10
# [Found] file2.txt: line 5
```

## 実行方法
```bash
cd exercise
go run main.go [検索ワード]
```
