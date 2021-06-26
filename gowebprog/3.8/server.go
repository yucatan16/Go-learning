package main

import (
	"fmt"
	"net/http"
)

// type HelloHandler struct{}

// func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello!")
// } 

// type WorldHandler struct{}

// func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "World!")
// } 

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "world!!")
}

func main() {
	// hello := HelloHandler{}
	// world := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	// http.Handleを呼び出すと実際にはDefaultServeMuxのHandleメソッドを呼ぶことになる
	// http.Handle("/hello", &hello)
	// http.Handle("/world", &world)

	// 関数helloをハンドラに変換して、そのハンドラをDefaultServeMuxに登録する
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	// http.HandleFuncの定義
	// func HanfleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	// 	DefaultServeMux.HandleFunc(pattern, handler)
	// }

	//ServeMuxのメソッドHandleFuncの定義
	// HandlerFuncは型なのでhandlerをキャストしてる
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	// 	mux.Handle(pattern, HandlerFunc(handler))
	// }
	
	server.ListenAndServe()
}