package main

import "fmt"

func main() {
	// 演習1: トピックベースのルーティング
	// 複数トピックの購読・配信ができるBrokerを実装してください
	// ヒント: map[string][]chan Message で管理
	fmt.Println("=== 演習1: トピックベースBroker ===")
	// broker := NewBroker()
	// ch := broker.Subscribe("events")
	// broker.Publish("events", "ユーザー登録")

	// 演習2: メッセージの永続化（ファイル保存）
	// 送信されたメッセージをファイルに追記保存し、再起動後に復元できるようにする
	// ヒント: os.OpenFile で追記モード、json.Encoder で書き込み
	fmt.Println("\n=== 演習2: メッセージの永続化 ===")
	// pBroker := NewPersistentBroker("messages.log")
	// pBroker.Publish("orders", "注文#1")
	// pBroker.Replay() // ファイルからメッセージを再生
}
