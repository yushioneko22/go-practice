package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// base62Encode はIDをBase62文字列に変換する
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

// URLStore はインメモリのURL短縮ストア
type URLStore struct {
	mu       sync.RWMutex
	urls     map[string]string // shortCode → originalURL
	counter  uint64
}

func NewURLStore() *URLStore {
	return &URLStore{
		urls: make(map[string]string),
	}
}

func (s *URLStore) Shorten(originalURL string) string {
	id := atomic.AddUint64(&s.counter, 1)
	code := base62Encode(id)

	s.mu.Lock()
	s.urls[code] = originalURL
	s.mu.Unlock()

	return code
}

func (s *URLStore) Resolve(code string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, ok := s.urls[code]
	return url, ok
}

func main() {
	store := NewURLStore()

	// デモ: URL短縮の動作を表示
	fmt.Println("=== [URLShortener] URL短縮サービス ===")
	testURLs := []string{
		"https://example.com/very/long/url/path",
		"https://google.com/search?q=golang+interview",
		"https://github.com/golang/go",
	}

	for _, url := range testURLs {
		code := store.Shorten(url)
		fmt.Printf("[URLShortener] %s → /%s\n", url, code)
	}

	// リダイレクト確認
	fmt.Println("\n=== [URLShortener] リダイレクト確認 ===")
	for code := range []int{1, 2, 3} {
		resolved, ok := store.Resolve(base62Encode(uint64(code + 1)))
		if ok {
			fmt.Printf("[URLShortener] /%s → %s\n", base62Encode(uint64(code+1)), resolved)
		}
	}

	// HTTPサーバー起動
	fmt.Println("\n=== [URLShortener] HTTPサーバー起動 (localhost:8080) ===")
	fmt.Println("[URLShortener] POST /shorten で短縮URL生成")
	fmt.Println("[URLShortener] GET /{code} でリダイレクト")

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only", http.StatusMethodNotAllowed)
			return
		}
		var req struct{ URL string `json:"url"` }
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		code := store.Shorten(req.URL)
		resp := map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/%s", code)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[1:]
		if url, ok := store.Resolve(code); ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":8080", nil)
}
