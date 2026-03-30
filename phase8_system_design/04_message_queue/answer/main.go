package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Message struct {
	Topic string `json:"topic"`
	Body  string `json:"body"`
}

// === 演習1: トピックベースBroker ===

type Broker struct {
	mu          sync.RWMutex
	subscribers map[string][]chan Message
}

func NewBroker() *Broker {
	return &Broker{subscribers: make(map[string][]chan Message)}
}

func (b *Broker) Subscribe(topic string) chan Message {
	b.mu.Lock()
	defer b.mu.Unlock()
	ch := make(chan Message, 10)
	b.subscribers[topic] = append(b.subscribers[topic], ch)
	return ch
}

func (b *Broker) Publish(topic, body string) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	msg := Message{Topic: topic, Body: body}
	for _, ch := range b.subscribers[topic] {
		select {
		case ch <- msg:
		default:
		}
	}
}

// === 演習2: 永続化Broker ===

type PersistentBroker struct {
	broker  *Broker
	logFile string
	mu      sync.Mutex
}

func NewPersistentBroker(logFile string) *PersistentBroker {
	return &PersistentBroker{
		broker:  NewBroker(),
		logFile: logFile,
	}
}

func (pb *PersistentBroker) Publish(topic, body string) {
	msg := Message{Topic: topic, Body: body}

	// ファイルに追記
	pb.mu.Lock()
	f, err := os.OpenFile(pb.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		json.NewEncoder(f).Encode(msg)
		f.Close()
	}
	pb.mu.Unlock()

	// 通常のPublish
	pb.broker.Publish(topic, body)
}

func (pb *PersistentBroker) Replay() []Message {
	f, err := os.Open(pb.logFile)
	if err != nil {
		return nil
	}
	defer f.Close()

	var messages []Message
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var msg Message
		if err := json.Unmarshal(scanner.Bytes(), &msg); err == nil {
			messages = append(messages, msg)
		}
	}
	return messages
}

func main() {
	fmt.Println("=== 演習1: トピックベースBroker ===")
	broker := NewBroker()
	ch := broker.Subscribe("events")
	go func() {
		for msg := range ch {
			fmt.Printf("[Events] %s: %s\n", msg.Topic, msg.Body)
		}
	}()
	broker.Publish("events", "ユーザー登録")
	broker.Publish("events", "ログイン")
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n=== 演習2: メッセージの永続化 ===")
	logFile := "/tmp/mq_messages.log"
	os.Remove(logFile)
	pBroker := NewPersistentBroker(logFile)
	pBroker.Publish("orders", "注文#1001")
	pBroker.Publish("orders", "注文#1002")
	pBroker.Publish("logs", "INFO: 起動")

	messages := pBroker.Replay()
	fmt.Printf("[Replay] %d件のメッセージを復元:\n", len(messages))
	for _, msg := range messages {
		fmt.Printf("[Replay] [%s] %s\n", msg.Topic, msg.Body)
	}
}
