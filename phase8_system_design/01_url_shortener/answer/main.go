package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

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

type URLEntry struct {
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	AccessCount int
}

type URLStore struct {
	mu      sync.RWMutex
	urls    map[string]*URLEntry
	counter uint64
}

func NewURLStore() *URLStore {
	return &URLStore{urls: make(map[string]*URLEntry)}
}

func (s *URLStore) ShortenWithTTL(url string, ttl time.Duration) string {
	id := atomic.AddUint64(&s.counter, 1)
	code := base62Encode(id)
	now := time.Now()

	s.mu.Lock()
	s.urls[code] = &URLEntry{
		OriginalURL: url,
		CreatedAt:   now,
		ExpiresAt:   now.Add(ttl),
	}
	s.mu.Unlock()
	return code
}

func (s *URLStore) Resolve(code string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	entry, ok := s.urls[code]
	if !ok {
		return "", errors.New("not found")
	}
	if !entry.ExpiresAt.IsZero() && time.Now().After(entry.ExpiresAt) {
		delete(s.urls, code)
		return "", errors.New("expired")
	}
	entry.AccessCount++
	return entry.OriginalURL, nil
}

func (s *URLStore) GetStats(code string) (*URLEntry, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	entry, ok := s.urls[code]
	if !ok {
		return nil, errors.New("not found")
	}
	return entry, nil
}

func (s *URLStore) ShortenCustom(code, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.urls[code]; exists {
		return fmt.Errorf("code '%s' is already in use", code)
	}
	s.urls[code] = &URLEntry{
		OriginalURL: url,
		CreatedAt:   time.Now(),
	}
	return nil
}

func main() {
	store := NewURLStore()

	fmt.Println("=== 演習1: 有効期限付きURL ===")
	code := store.ShortenWithTTL("https://example.com", 2*time.Second)
	fmt.Printf("[TTL] 短縮コード: %s\n", code)
	url, err := store.Resolve(code)
	fmt.Printf("[TTL] 即座にResolve: url=%s, err=%v\n", url, err)
	time.Sleep(3 * time.Second)
	url, err = store.Resolve(code)
	fmt.Printf("[TTL] 3秒後にResolve: url=%s, err=%v\n", url, err)

	fmt.Println("\n=== 演習2: アクセスカウンター ===")
	code2 := store.ShortenWithTTL("https://google.com", 1*time.Hour)
	store.Resolve(code2)
	store.Resolve(code2)
	store.Resolve(code2)
	stats, _ := store.GetStats(code2)
	fmt.Printf("[Stats] %s のアクセス回数: %d\n", code2, stats.AccessCount)

	fmt.Println("\n=== 演習3: カスタムURL ===")
	err = store.ShortenCustom("mysite", "https://mysite.com")
	fmt.Printf("[Custom] mysite → err=%v\n", err)
	err = store.ShortenCustom("mysite", "https://other.com")
	fmt.Printf("[Custom] mysite(重複) → err=%v\n", err)
	url, _ = store.Resolve("mysite")
	fmt.Printf("[Custom] Resolve(mysite) → %s\n", url)
}
