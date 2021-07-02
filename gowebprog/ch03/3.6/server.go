package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
} 

func main() {
	// ハンドラとはServeHTTPというメソッドを持ったインターフェースのこと
	handler := MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:3000",
		Handler: &handler,
	}

	server.ListenAndServe()
}