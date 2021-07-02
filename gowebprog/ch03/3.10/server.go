package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called -" + name)
		h(w,r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:3000",
	}

	fmt.Printf("%T\n", hello) // func(http.ResponseWriter, *http.Request)
	// http.HandlerFuncの型と一致するためlogの引数として渡せる
	// https://golang.org/pkg/net/http/#HandlerFunc

	// 関数helloをハンドラに変換して、そのハンドラをDefaultServeMuxに登録する

	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}