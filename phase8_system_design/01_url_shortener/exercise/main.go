package main

import (
	"fmt"
	"sync"
	"time"
)

// 演習: 以下の機能を追加してURL短縮サービスを拡張してください

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func base62Encode(id uint64) string {
	if id == 0 {
		return string(base62Chars[0])
	}
	result := []byte{}
	for id > 0 {
		result = append([]byte{base62Chars[id%62]}, result...)
		id /= 62
	}
	return string(result)
}

// URLEntry はURLの情報を保持する構造体
type URLEntry struct {
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   time.Time     // 演習1: 有効期限
	AccessCount int           // 演習2: アクセスカウンター
	CustomCode  string        // 演習3: カスタムURL
}

// URLStore はURL短縮ストア
type URLStore struct {
	mu      sync.RWMutex
	urls    map[string]*URLEntry
	counter uint64
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]*URLEntry),
	}
}

func main() {
	store := NewURLStore()
	_ = store

	// 演習1: 有効期限付きURLの実装
	// Shorten時にTTL（有効期限）を指定できるようにしてください
	// Resolve時に有効期限切れなら「expired」を返す
	fmt.Println("=== 演習1: 有効期限付きURL ===")
	// code := store.ShortenWithTTL("https://example.com", 5*time.Second)
	// fmt.Printf("[TTL] 短縮コード: %s\n", code)
	// url, _ := store.Resolve(code) // → 有効
	// time.Sleep(6 * time.Second)
	// url, _ = store.Resolve(code) // → 期限切れ

	// 演習2: アクセスカウンターの実装
	// Resolve時にアクセス回数をインクリメントし、統計情報を返す関数を追加
	fmt.Println("\n=== 演習2: アクセスカウンター ===")
	// stats := store.GetStats(code)
	// fmt.Printf("[Stats] アクセス回数: %d\n", stats.AccessCount)

	// 演習3: カスタムURLの対応
	// ユーザーが短縮コードを指定できる ShortenCustom 関数を実装
	// 既に使われているコードの場合はエラーを返す
	fmt.Println("\n=== 演習3: カスタムURL ===")
	// err := store.ShortenCustom("mysite", "https://example.com")
	// if err != nil { fmt.Println(err) }
}

// 演習1: ShortenWithTTL を実装してください
// func (s *URLStore) ShortenWithTTL(url string, ttl time.Duration) string { }

// 演習1: Resolve を実装してください（有効期限チェック付き）
// func (s *URLStore) Resolve(code string) (string, error) { }

// 演習2: GetStats を実装してください
// func (s *URLStore) GetStats(code string) *URLEntry { }

// 演習3: ShortenCustom を実装してください
// func (s *URLStore) ShortenCustom(code, url string) error { }
