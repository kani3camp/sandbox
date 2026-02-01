package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/a-h/templ"
)

var (
	count = 0
	mu    sync.Mutex
)

func main() {
	// 1. TOPページ：Layoutの中にCounterを入れて表示
	http.Handle("/", templ.Handler(layout(counter(0))))

	// 2. カウントアップ処理：htmxからのリクエストに応答
	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		mu.Lock()
		count++
		current := count
		mu.Unlock()

		// Counter部分だけのHTMLを返却
		templ.Handler(counter(current)).ServeHTTP(w, r)
	})

	fmt.Println("サーバー起動中: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
