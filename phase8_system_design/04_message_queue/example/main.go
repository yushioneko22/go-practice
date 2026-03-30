package main

import (
	"fmt"
	"sync"
	"time"
)

// Message はキューに送るメッセージ
type Message struct {
	Topic string
	Body  string
}

// Broker はトピックベースのPub/Subブローカー
type Broker struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Message
}

func NewBroker() *Broker {
	return &Broker{
		subscribers: make(map[string][]chan Message),
	}
}

// Subscribe はトピックを購読し、メッセージを受信するchannelを返す
func (b *Broker) Subscribe(topic string) chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan Message, 10) // バッファ付き
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

// Publish はトピックにメッセージを送信する
func (b *Broker) Publish(topic, body string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	msg := Message{Topic: topic, Body: body}
	for _, ch := range b.subscribers[topic] {
		// ノンブロッキング送信（バッファが満杯なら破棄）
		select {
		case ch <- msg:
		default:
			fmt.Printf("[Broker] WARNING: subscriber buffer full for topic '%s'\n", topic)
		}
	}
}

func main() {
	fmt.Println("=== [MQ] メッセージキュー (Pub/Sub) ===")
	broker := NewBroker()

	// Subscriber 1: "orders" トピック
	orderCh := broker.Subscribe("orders")
	go func() {
		for msg := range orderCh {
			fmt.Printf("[Subscriber-Orders] 受信: %s\n", msg.Body)
		}
	}()

	// Subscriber 2: "logs" トピック
	logCh := broker.Subscribe("logs")
	go func() {
		for msg := range logCh {
			fmt.Printf("[Subscriber-Logs] 受信: %s\n", msg.Body)
		}
	}()

	// Publisher
	time.Sleep(100 * time.Millisecond) // subscriber起動待ち
	broker.Publish("orders", "注文#1001: Goの本")
	broker.Publish("orders", "注文#1002: キーボード")
	broker.Publish("logs", "INFO: サーバー起動")
	broker.Publish("logs", "WARN: メモリ使用率80%")

	time.Sleep(500 * time.Millisecond) // メッセージ処理待ち
	fmt.Println("[MQ] 完了")
}
